# Launch checklist

Use this with [`deployment.md`](deployment.md). Tick items only when they have
been checked on the real domain, not just locally.

## A. Client decisions

- [x] Official business name and tagline confirmed in writing (16 November 2026).
- [x] Bottled-water pricing confirmed: quoted by quantity, no fixed prices.
- [x] Branded-bottle customisation and minimum order confirmed (Ksh 500–1,000; 10 bales).
- [x] Dispenser-cleaning price confirmed (Ksh 500–1,000).
- [x] Free-delivery areas and conditions confirmed (free depends on distance and quantity).
- [x] KEBS position confirmed: certificate is with the supplier; no number or claim published.
- [x] Purification steps supplied and published without micron sizes.
- [x] Permission to show other customers' labels (Wave360 and the church labels) — cleared for this build by the project owner on 16 November 2026.
- [x] Water pumps and dispensers confirmed for sale.
- [x] Permission to display the Paybill number online — cleared for this build by the project owner on 16 November 2026.
- [x] Facebook page confirmed as the AQUAMOR NGARA business page.
- [x] Bottled-water sizes confirmed: 200ml, 300ml, 500ml, 1L, 1.5L, 2L, 3L, 5L, 10L and 20L.
- [x] New-bottle prices distinguished from refills; water is separate.
- [x] Swahili copy reviewed on 16 November 2026; a native-speaker read-through is still recommended before launch.
- [ ] Real customer testimonials, if any (optional).
- [ ] Original cleaning before/after stills, if supplied (optional).

**Deferred to a later phase (not needed for launch):** online payments and automated replies. The Flutterwave merchant account, API keys, webhook secret and settlement details are not required now; the code stays in place but switched off.

## B. Code and content (done, re-check after changes)

- [x] Go server reads `ADDR`, `DB_PATH`, `WEB_DIR` and `ADMIN_TOKEN` from the
      environment.
- [x] Server binds to loopback by default and has read/write/idle timeouts.
- [x] Graceful shutdown on `SIGTERM`/`SIGINT`.
- [x] Security headers: CSP, `X-Content-Type-Options`, `X-Frame-Options`,
      `Referrer-Policy`, `Permissions-Policy`, and HSTS over HTTPS.
- [x] Static cache headers and `no-store` on the API.
- [x] Database health check at `/api/v1/health`.
- [x] Go tests for catalog prices, order totals/validation and admin auth.
- [x] Online payment backend: hosted checkout, return-page verification, signed
      webhook, server-side price/amount/currency checks and paid-order update.
      Built but deferred for this launch; it stays hidden until configured.
- [x] Payment tests cover checkout creation, verification, webhook signatures,
      amount checks and the unconfigured state.
- [x] Authenticated admin orders page (`web/admin.html`) with token login,
      line items, payment state, filters and CSV export; `noindex` and no-store.
- [x] CI runs build, vet, `go test -race` and JavaScript/script syntax checks.
- [x] Deployment files: systemd unit, Caddy example, env example, backup script,
      smoke test and domain-placement script.
- [x] Responsive gallery and banner images (`srcset`/`sizes`).
- [x] Below-the-fold banner images lazy-loaded.
- [x] Fonts are self-hosted; the CSP allows no third-party font origin.
- [x] Order and checkout endpoints are rate-limited per client IP.
- [x] Request logging records method/path/status/duration/IP, with no customer
      data, headers or query strings.
- [x] Unused legacy dashboard and Next.js config removed.
- [x] No certification, health or "100% pure" claims.
- [x] English/Swahili copy in place; internal Swahili review done on 16 November 2026, with a native-speaker read recommended.
- [x] Accessibility and SEO were 100/100 in the section 18 Lighthouse run; re-run on the deployed domain after the image changes.
- [x] No broken internal documentation links.

## C. Domain and hosting (blocked on the client)

- [x] Preferred domain `aquamorwater.co.ke` registered under the client's name and paid for (Truehost, 18 Nov 2026; renews 17 Nov 2027). Confirm the registrant name reads AQUAMOR PURIFIED DRINKING WATER.
- [ ] DNS `A`/`AAAA` records point at the server (need DNS access from the client).
- [x] `scripts/set-domain.sh` run with `aquamorwater.co.ke` (re-run if the final domain differs).
- [ ] Caddy site address updated and certificate issued.
- [ ] HTTP redirects to HTTPS; `www` redirects to the canonical host.
- [ ] Firewall allows only 22, 80 and 443; port 8080 is not public.
- [ ] `ADMIN_TOKEN` is strong, stored at mode 600 and not in git.
- [ ] The service starts on boot (`systemctl is-enabled aquamor`).

**Deferred with online payments:** store payment keys at mode 600, set the Flutterwave webhook URL and set `PUBLIC_BASE_URL` only when online payments are enabled later. None of these are needed for this launch.

## D. Verification on the real domain

- [ ] `deploy/smoke-test.sh https://<domain>` passes.
- [ ] `https://<domain>/api/v1/orders/all` without a token returns 401 or 503.
- [ ] A test order is stored and can be read with the admin token.
- [ ] `/admin.html` loads, accepts the admin token, shows the test order with its
      line items and payment state, and the CSV export opens correctly.
- [ ] Canonical URL, `og:url`, `og:image`, `robots.txt` and `sitemap.xml` use the
      real domain.
- [ ] WhatsApp/social link preview shows the correct image and title.
- [x] Automated mobile emulation at 320, 360, 390 and 414px: no horizontal
      scrolling, menu, order builder and gallery work (physical iPhone check
      still recommended for safe areas and the WhatsApp hand-off).
- [x] Automated low-end-width emulation at 320/360px; physical Android check
      still recommended.
- [x] Swahili switch works on every section and in WhatsApp messages (checked in
      the browser); native-speaker read still recommended.
- [ ] Print the A4 price card and check it fits one page.
- [ ] Home page, price card, planner, label designer and gallery lightbox all
      work with the final domain.
- [ ] `/api/v1/payments/config` returns `"enabled": false` (online payments deferred).

Online payment end-to-end checks are deferred with the feature until Flutterwave is supplied later.

## E. Data and recovery

- [ ] Daily backup cron installed and the first backup exists.
- [ ] Backups are copied off the server.
- [ ] A restore has been tested on a temporary copy.
- [ ] Database file and directory permissions are `aquamor:aquamor`, mode 600/750.
- [ ] Disk space and log rotation checked (`journalctl --disk-usage`).

## F. Handover

- [ ] Client has the change-request sheet
      ([`how-to-request-a-change.md`](how-to-request-a-change.md)).
- [ ] Maintainer has the update guide
      ([`how-to-update-your-site.md`](how-to-update-your-site.md)).
- [ ] Revision rounds, approval owner, payment dates and ongoing maintenance are
      agreed in writing.
- [ ] The client knows who owns the domain, hosting, repository and photos.
- [ ] Emergency contact and response expectations are agreed.
