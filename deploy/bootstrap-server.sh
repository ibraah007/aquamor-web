#!/usr/bin/env bash
# One-shot bootstrap for a fresh Ubuntu LTS server.
#
# Run as root on the server named in DNS (or with the DNS still pointing
# elsewhere; Caddy will retry the certificate until DNS resolves):
#
#   sudo ./bootstrap-server.sh aquamorwater.co.ke
#
# Optional environment overrides:
#   REPO_URL=https://github.com/ibraah007/aquamor-web.git
#   REPO_REF=main
#   ENABLE_UFW=1|0          (default 1; opens 22, 80 and 443)
#
# The script is safe to re-run: it keeps an existing /etc/aquamor/aquamor.env
# and only refreshes the code, binary, service and Caddyfile.
set -euo pipefail

DOMAIN="${1:-aquamorwater.co.ke}"
REPO_URL="${REPO_URL:-https://github.com/ibraah007/aquamor-web.git}"
REPO_REF="${REPO_REF:-main}"
ENABLE_UFW="${ENABLE_UFW:-1}"

INSTALL_DIR=/opt/aquamor
ENV_FILE=/etc/aquamor/aquamor.env
DB_DIR=/var/lib/aquamor

log() { printf '\n==> %s\n' "$*"; }
die() { printf 'error: %s\n' "$*" >&2; exit 1; }

[[ $EUID -eq 0 ]] || die "run this script as root (sudo $0 $DOMAIN)"
[[ "$DOMAIN" =~ ^[a-z0-9.-]+$ ]] || die "invalid domain: $DOMAIN"

export DEBIAN_FRONTEND=noninteractive

log "Installing system packages"
apt-get update
apt-get install -y --no-install-recommends \
	build-essential sqlite3 curl git ca-certificates openssl ufw \
	debian-keyring debian-archive-keyring apt-transport-https golang-go

if ! command -v go >/dev/null 2>&1; then
	die "Go was not installed; install Go 1.22+ and re-run"
fi
go_version="$(go env GOVERSION | sed 's/^go//')"
if [[ "$(printf '%s\n' 1.22 "$go_version" | sort -V | head -n1)" != "1.22" ]]; then
	die "Go 1.22 or newer is required (found $go_version)"
fi

if ! command -v caddy >/dev/null 2>&1; then
	log "Installing Caddy from the official repository"
	curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' \
		| gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
	curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' \
		> /etc/apt/sources.list.d/caddy-stable.list
	apt-get update
	apt-get install -y caddy
fi

log "Creating the service account and directories"
if ! id aquamor >/dev/null 2>&1; then
	useradd --system --home "$INSTALL_DIR" --shell /usr/sbin/nologin aquamor
fi
install -d -o aquamor -g aquamor -m 750 "$DB_DIR"
install -d -m 750 /etc/aquamor

log "Fetching $REPO_REF from $REPO_URL"
if [[ -d "$INSTALL_DIR/.git" ]]; then
	git -C "$INSTALL_DIR" fetch --prune origin
	git -C "$INSTALL_DIR" checkout "$REPO_REF"
	git -C "$INSTALL_DIR" reset --hard "origin/$REPO_REF"
else
	install -d -m 755 "$INSTALL_DIR"
	git clone --branch "$REPO_REF" "$REPO_URL" "$INSTALL_DIR"
fi
chown -R root:root "$INSTALL_DIR"

log "Building the backend"
build_tmp="$(mktemp)"
trap 'rm -f "$build_tmp"' EXIT
(
	cd "$INSTALL_DIR"
	CGO_ENABLED=1 go build -trimpath -ldflags="-s -w" -o "$build_tmp" ./backend/cmd/server
)
install -m 755 "$build_tmp" "$INSTALL_DIR/aquamor-backend"

log "Writing the environment file"
if [[ ! -f "$ENV_FILE" ]]; then
	admin_token="$(openssl rand -hex 32)"
	umask 077
	cat > "$ENV_FILE" <<EOF
ADDR=127.0.0.1:8080
DB_PATH=$DB_DIR/aquamor.db
WEB_DIR=$INSTALL_DIR/web
ADMIN_TOKEN=$admin_token
FLUTTERWAVE_SECRET_KEY=
FLUTTERWAVE_WEBHOOK_SECRET=
PUBLIC_BASE_URL=https://$DOMAIN
EOF
else
	log "Keeping the existing $ENV_FILE (admin token unchanged)"
fi
chown root:root "$ENV_FILE"
chmod 600 "$ENV_FILE"

log "Updating the domain in site files"
(cd "$INSTALL_DIR" && ./scripts/set-domain.sh "$DOMAIN")

log "Installing the systemd service"
install -m 644 "$INSTALL_DIR/deploy/aquamor.service" /etc/systemd/system/aquamor.service
systemctl daemon-reload
systemctl enable aquamor >/dev/null
systemctl restart aquamor

log "Configuring Caddy for $DOMAIN"
[[ -f /etc/caddy/Caddyfile ]] && cp -a /etc/caddy/Caddyfile "/etc/caddy/Caddyfile.bak.$(date +%s)"
sed "s/aquamorwater\.co\.ke/$DOMAIN/g" "$INSTALL_DIR/deploy/Caddyfile" > /etc/caddy/Caddyfile
caddy validate --config /etc/caddy/Caddyfile
systemctl enable caddy >/dev/null
systemctl reload caddy 2>/dev/null || systemctl restart caddy

if [[ "$ENABLE_UFW" == "1" ]]; then
	log "Configuring the firewall (22, 80, 443)"
	ufw --force default deny incoming
	ufw --force default allow outgoing
	ufw allow 22/tcp
	ufw allow 80/tcp
	ufw allow 443/tcp
	ufw --force enable
fi

log "Installing the backup job"
install -m 755 "$INSTALL_DIR/deploy/backup.sh" /usr/local/sbin/aquamor-backup
install -d -m 750 /var/backups/aquamor
cat > /etc/cron.d/aquamor-backup <<'EOF'
# Daily consistent SQLite backup, kept for 14 days.
15 2 * * * root DB_PATH=/var/lib/aquamor/aquamor.db BACKUP_DIR=/var/backups/aquamor /usr/local/sbin/aquamor-backup >> /var/log/aquamor-backup.log 2>&1
EOF
chmod 644 /etc/cron.d/aquamor-backup

log "Checking the local service"
for _ in $(seq 1 10); do
	if curl -fsS http://127.0.0.1:8080/api/v1/health >/dev/null 2>&1; then
		ok=1
		break
	fi
	sleep 1
done
[[ "${ok:-0}" == "1" ]] || {
	journalctl -u aquamor -n 30 --no-pager >&2 || true
	die "the service did not become healthy on 127.0.0.1:8080"
}

cat <<EOF

Bootstrap complete.

  Service:    systemctl status aquamor
  Health:     curl -fsS http://127.0.0.1:8080/api/v1/health
  Admin token: $ENV_FILE (mode 600, not printed)

DNS records to set at the registrar:

  A     @      $(curl -fsS https://api.ipify.org || echo '<server IPv4>')
  A     www    $(curl -fsS https://api.ipify.org || echo '<server IPv4>')

After DNS resolves, Caddy issues the certificate automatically. Then run:

  $INSTALL_DIR/deploy/smoke-test.sh https://$DOMAIN
EOF
