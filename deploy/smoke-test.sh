#!/usr/bin/env bash
# Check a running Aquamor deployment.
#
#   ./deploy/smoke-test.sh http://127.0.0.1:8080
#   ./deploy/smoke-test.sh https://aquamor.co.ke
#
# If ADMIN_TOKEN is set in the environment, the protected orders endpoint is
# also checked with that token. Otherwise it must answer 401 or 503.
set -euo pipefail

base="${1:-http://127.0.0.1:8080}"
base="${base%/}"
body="$(mktemp)"
trap 'rm -f "$body"' EXIT
failures=0

status() {
	curl -sS -o "$body" -w '%{http_code}' "$@"
}

pass() {
	echo "ok   $1"
}

fail() {
	echo "FAIL $1" >&2
	failures=$((failures + 1))
}

expect_status() {
	local description="$1" want="$2" got="$3"
	if [[ "$got" == "$want" ]]; then
		pass "$description ($got)"
	else
		fail "$description: got $got, want $want"
	fi
}

code="$(status "$base/")" || { echo "FAIL cannot reach $base" >&2; exit 1; }
expect_status "home page" "200" "$code"
if ! grep -qi "Aquamor" "$body"; then
	fail "home page does not mention Aquamor"
else
	pass "home page contains the business name"
fi

code="$(status "$base/api/v1/health")"
expect_status "health endpoint" "200" "$code"
if ! grep -q '"status":"ok"' "$body"; then
	fail "health endpoint did not report status ok"
else
	pass "health endpoint reports status ok"
fi

code="$(status "$base/api/v1/products")"
expect_status "products endpoint" "200" "$code"
if ! grep -q '"price_swap"' "$body"; then
	fail "products endpoint did not return catalog data"
else
	pass "products endpoint returns catalog data"
fi

code="$(status "$base/payment.html")"
expect_status "payment status page" "200" "$code"
if ! grep -qi "Aquamor" "$body"; then
	fail "payment status page does not mention Aquamor"
else
	pass "payment status page contains the business name"
fi

code="$(status "$base/admin.html")"
expect_status "admin orders page" "200" "$code"
if ! grep -qi "noindex" "$body"; then
	fail "admin orders page is not marked noindex"
else
	pass "admin orders page is noindex"
fi

code="$(status "$base/api/v1/payments/config")"
expect_status "payments config endpoint" "200" "$code"
if ! grep -q '"enabled"' "$body"; then
	fail "payments config endpoint did not report enabled state"
else
	pass "payments config endpoint reports enabled state"
fi

code="$(status "$base/api/v1/orders/all")"
if [[ -n "${ADMIN_TOKEN:-}" ]]; then
	# Never send the bearer token to a remote plain-HTTP URL. The token would
	# leave the machine before any HTTP-to-HTTPS redirect could happen.
	if [[ "$base" != https://* && ! "$base" =~ ^http://(127\.0\.0\.1|localhost|\[::1\])(:[0-9]+)?$ ]]; then
		echo "refusing to send ADMIN_TOKEN over plain HTTP to $base" >&2
		exit 1
	fi
	code="$(status -H "Authorization: Bearer $ADMIN_TOKEN" "$base/api/v1/orders/all")"
	expect_status "orders endpoint with ADMIN_TOKEN" "200" "$code"
else
	if [[ "$code" == "401" || "$code" == "503" ]]; then
		pass "orders endpoint is protected ($code)"
	else
		fail "orders endpoint returned $code without a token; expected 401 or 503"
	fi
fi

# Security headers are set by the Go server, with HSTS added by Caddy over TLS.
headers="$(curl -sSI "$base/" | tr -d '\r')"
for header in "x-content-type-options: nosniff" "x-frame-options: DENY" "content-security-policy:"; do
	if grep -qi "^$header" <<<"$headers"; then
		pass "security header $header"
	else
		fail "missing security header $header"
	fi
done
if [[ "$base" == https://* ]]; then
	if grep -qi '^strict-transport-security:' <<<"$headers"; then
		pass "HSTS header present"
	else
		fail "missing HSTS header"
	fi
fi

if [[ "$failures" -ne 0 ]]; then
	echo "$failures check(s) failed" >&2
	exit 1
fi
echo "all checks passed"
