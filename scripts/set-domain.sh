#!/usr/bin/env bash
# Replace the aquamor.co.ke placeholders with the confirmed domain.
#
#   ./scripts/set-domain.sh aquamor.co.ke
#
# This updates the canonical/Open Graph URLs, robots.txt, sitemap.xml and the
# Caddy example. It does not touch the client's DNS or hosting settings.
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

files=(web/index.html web/robots.txt web/sitemap.xml deploy/Caddyfile)
changed=0

for file in "${files[@]}"; do
	if [[ ! -f "$file" ]]; then
		echo "skip $file (not found)"
		continue
	fi
	before="$(grep -c 'aquamor\.co\.ke' "$file" || true)"
	sed -i "s|aquamor\\.co\\.ke|${domain}|g" "$file"
	after="$(grep -c 'aquamor\.co\.ke' "$file" || true)"
	if [[ "$before" != "$after" ]]; then
		changed=$((changed + 1))
	fi
	echo "updated $file"
done

echo "domain set to $domain in $changed file(s)"
echo "next: update the Caddy site address if needed, then run deploy/smoke-test.sh"
