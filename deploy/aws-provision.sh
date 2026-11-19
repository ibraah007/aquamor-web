#!/usr/bin/env bash
# Create the Aquamor web server on AWS from a configured AWS CLI.
#
# Prerequisites: aws CLI v2 with credentials that allow the actions in
# deploy/aws-iam-policy.json. Run deploy/bootstrap-server.sh over SSH afterwards.
#
# Usage:
#   ./deploy/aws-provision.sh
#
# Optional environment overrides:
#   DOMAIN=aquamorwater.co.ke
#   REGION=af-south-1
#   INSTANCE_TYPE=t3.small
#   NAME=aquamor-web
#   KEY_NAME=aquamor-deploy
#   KEY_PATH=$HOME/.ssh/aquamor_deploy
#   SSH_CIDR=<your-ip>/32   (defaults to this machine's current public IP)
#   SKIP_DEPLOY=1           (only create the server, do not copy/bootstrap)
set -euo pipefail

DOMAIN="${DOMAIN:-aquamorwater.co.ke}"
REGION="${REGION:-af-south-1}"
INSTANCE_TYPE="${INSTANCE_TYPE:-t3.small}"
NAME="${NAME:-aquamor-web}"
KEY_NAME="${KEY_NAME:-aquamor-deploy}"
KEY_PATH="${KEY_PATH:-$HOME/.ssh/aquamor_deploy}"
SSH_CIDR="${SSH_CIDR:-}"
SKIP_DEPLOY="${SKIP_DEPLOY:-0}"

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

log() { printf '\n==> %s\n' "$*"; }
die() { printf 'error: %s\n' "$*" >&2; exit 1; }

command -v aws >/dev/null 2>&1 || die "aws CLI is not installed"
command -v curl >/dev/null 2>&1 || die "curl is not installed"
aws sts get-caller-identity >/dev/null 2>&1 \
	|| die "AWS credentials are not configured; run 'aws configure' first"

log "Using region $REGION"
ACCOUNT_ID="$(aws sts get-caller-identity --query Account --output text)"

if [[ -z "$SSH_CIDR" ]]; then
	SSH_CIDR="$(curl -fsS https://checkip.amazonaws.com | tr -d '[:space:]')/32"
fi
log "SSH will be allowed only from $SSH_CIDR"

log "Preparing the SSH key pair"
if aws ec2 describe-key-pairs --key-names "$KEY_NAME" --region "$REGION" >/dev/null 2>&1; then
	[[ -f "$KEY_PATH.pem" ]] || die "key pair $KEY_NAME exists in AWS but $KEY_PATH.pem is missing"
	log "Reusing existing key pair $KEY_NAME"
else
	umask 077
	aws ec2 create-key-pair --key-name "$KEY_NAME" --region "$REGION" \
		--query KeyMaterial --output text > "$KEY_PATH.pem"
	log "Created key pair $KEY_NAME at $KEY_PATH.pem"
fi
chmod 600 "$KEY_PATH.pem"

log "Preparing the security group"
SG_ID="$(aws ec2 describe-security-groups --region "$REGION" \
	--filters "Name=group-name,Values=$NAME-sg" \
	--query 'SecurityGroups[0].GroupId' --output text 2>/dev/null || true)"
if [[ -z "$SG_ID" || "$SG_ID" == "None" ]]; then
	SG_ID="$(aws ec2 create-security-group --region "$REGION" \
		--group-name "$NAME-sg" --description "Aquamor web: SSH, HTTP, HTTPS" \
		--query GroupId --output text)"
	log "Created security group $SG_ID"
else
	log "Reusing security group $SG_ID"
fi

for rule in "22 $SSH_CIDR" "80 0.0.0.0/0" "443 0.0.0.0/0"; do
	set -- $rule
	aws ec2 authorize-security-group-ingress --region "$REGION" \
		--group-id "$SG_ID" --protocol tcp --port "$1" --cidr "$2" >/dev/null 2>&1 \
		|| log "Port $1 rule already present"
done

log "Resolving the latest Ubuntu 24.04 LTS AMI"
AMI_ID="$(aws ssm get-parameters --region "$REGION" \
	--names /aws/service/canonical/ubuntu/server/24.04/stable/current/amd64/hvm/ebs-gp3/ami-id \
	--query 'Parameters[0].Value' --output text)"
[[ "$AMI_ID" == ami-* ]] || die "could not resolve an Ubuntu AMI in $REGION"

SUBNET_ID="$(aws ec2 describe-subnets --region "$REGION" \
	--filters Name=default-for-az,Values=true \
	--query 'Subnets[0].SubnetId' --output text 2>/dev/null || true)"
if [[ -z "$SUBNET_ID" || "$SUBNET_ID" == "None" ]]; then
	die "no default subnet in $REGION; create a VPC/subnet or set one manually"
fi

log "Launching $INSTANCE_TYPE ($AMI_ID)"
INSTANCE_ID="$(aws ec2 run-instances --region "$REGION" \
	--image-id "$AMI_ID" --instance-type "$INSTANCE_TYPE" \
	--key-name "$KEY_NAME" --security-group-ids "$SG_ID" --subnet-id "$SUBNET_ID" \
	--associate-public-ip-address \
	--tag-specifications "ResourceType=instance,Tags=[{Key=Name,Value=$NAME}]" \
	--query 'Instances[0].InstanceId' --output text)"

log "Waiting for $INSTANCE_ID to be running and pass status checks"
aws ec2 wait instance-running --region "$REGION" --instance-ids "$INSTANCE_ID"
aws ec2 wait instance-status-ok --region "$REGION" --instance-ids "$INSTANCE_ID"

PUBLIC_IP="$(aws ec2 describe-instances --region "$REGION" \
	--instance-ids "$INSTANCE_ID" \
	--query 'Reservations[0].Instances[0].PublicIpAddress' --output text)"

log "Server $INSTANCE_ID is up at $PUBLIC_IP"

if [[ "$SKIP_DEPLOY" != "1" ]]; then
	SSH_OPTS=(-o StrictHostKeyChecking=accept-new -o ConnectTimeout=5 -i "$KEY_PATH.pem")
	log "Waiting for SSH to become reachable"
	ssh_ready=0
	for _ in $(seq 1 30); do
		if ssh "${SSH_OPTS[@]}" "ubuntu@$PUBLIC_IP" true 2>/dev/null; then
			ssh_ready=1
			break
		fi
		sleep 5
	done
	[[ "$ssh_ready" == "1" ]] || die "SSH did not become reachable; server is still running at $PUBLIC_IP"

	log "Copying the repository and running the bootstrap"
	archive="$(mktemp --suffix=.tgz)"
	trap 'rm -f "$archive"' EXIT
	tar -czf "$archive" --exclude=.git --exclude=node_modules --exclude='*.db' -C "$REPO_ROOT" .
	scp "${SSH_OPTS[@]}" "$archive" "ubuntu@$PUBLIC_IP:/tmp/aquamor-src.tgz"
	ssh "${SSH_OPTS[@]}" "ubuntu@$PUBLIC_IP" \
		"rm -rf /tmp/aquamor-web && mkdir -p /tmp/aquamor-web && tar -xzf /tmp/aquamor-src.tgz -C /tmp/aquamor-web && sudo /tmp/aquamor-web/deploy/bootstrap-server.sh $DOMAIN"
	rm -f "$archive"
	trap - EXIT
fi

cat <<EOF

Server ready.

  Account:     $ACCOUNT_ID
  Region:      $REGION
  Instance:    $INSTANCE_ID ($INSTANCE_TYPE)
  Public IP:   $PUBLIC_IP
  SSH key:     $KEY_PATH.pem
  SSH command: ssh -i $KEY_PATH.pem ubuntu@$PUBLIC_IP

DNS records to add at Truehost:

  A  @    $PUBLIC_IP
  A  www  $PUBLIC_IP

After DNS resolves, verify from this machine:
  ./deploy/smoke-test.sh https://$DOMAIN
EOF
