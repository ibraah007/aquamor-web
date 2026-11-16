#!/usr/bin/env bash
# Replace the current default domain (aquamorwater.co.ke) with the confirmed
# domain.
#
#   ./scripts/set-domain.sh aquamorwater.co.ke
#
# This updates the canonical/Open Graph URLs, robots.txt, sitemap.xml, the
# Caddy example and the PUBLIC_BASE_URL in the environment example. It does not
# touch the client's DNS or hosting settings. If the files already use a
# different domain, edit them directly.
set -euo pipefail

if [[ $# -ne 1 ]]; then
	echo "usage: $0 <domain>" >&2
	exit 2
fi

domain="${1#https://}"
domain="${domain#http://}"
domain="${domain#www.}"
domain="${domain%/}"

if [[ ! "$domain" =~ ^[A-Za-z0-9]([A-Za-z0-9.-]*[A-Za-z0-9])?\.[A-Za-z]{2,}$ ]]; then
	echo "invalid domain: $1" >&2
	exit 2
fi

files=(web/index.html web/robots.txt web/sitemap.xml deploy/Caddyfile deploy/aquamor.env.example)
changed=0

for file in "${files[@]}"; do
	if [[ ! -f "$file" ]]; then
		echo "skip $file (not found)"
		continue
	fi
	before="$(grep -c 'aquamorwater\.co\.ke' "$file" || true)"
	sed -i "s|aquamorwater\\.co\\.ke|${domain}|g" "$file"
	after="$(grep -c 'aquamorwater\.co\.ke' "$file" || true)"
	if [[ "$before" != "$after" ]]; then
		changed=$((changed + 1))
	fi
	echo "updated $file"
done

echo "domain set to $domain in $changed file(s)"
echo "next: update the Caddy site address if needed, then run deploy/smoke-test.sh"
