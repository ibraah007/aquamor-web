# Launch checklist

Use this with [`deployment.md`](deployment.md). Tick items only when they have
been checked on the real domain, not just locally.

## A. Client decisions (blocked)

- [ ] Official business name and tagline confirmed in writing.
- [ ] Small bottled-water prices confirmed (300ml, 500ml, 1L, 1.5L).
- [ ] Branded-bottle pricing and minimum order confirmed.
- [ ] Dispenser-cleaning price confirmed.
- [ ] Free-delivery areas and any minimum order confirmed.
- [ ] KEBS permit number and certificate photo supplied, if the client has them.
- [ ] Purification steps supplied, if the client wants them published.
- [ ] Permission to show other customers' labels (Wave360 and the church labels).
- [ ] Confirm whether the water pumps are for sale.
- [ ] Permission to display the Paybill number online.
- [ ] Confirm the Facebook page is a business page, not a personal profile.
- [ ] Confirm the bottled-water sizes shown on the flyer and in the client's text.
- [ ] Confirm whether the regular 20L bottle price includes water.
- [ ] Native-speaker Swahili review.
- [ ] Real customer testimonials, if any.
- [ ] Original cleaning before/after stills, if supplied.

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
- [x] CI runs build, vet, `go test -race` and JavaScript/script syntax checks.
- [x] Deployment files: systemd unit, Caddy example, env example, backup script,
      smoke test and domain-placement script.
- [x] Responsive gallery and banner images (`srcset`/`sizes`).
- [x] Below-the-fold banner images lazy-loaded.
- [x] Unused legacy dashboard and Next.js config removed.
- [x] No certification, health or "100% pure" claims.
- [x] English/Swahili copy in place; Swahili still needs a native-speaker check.
- [x] Accessibility and SEO checks pass locally (Lighthouse 100 / 100).
- [x] No broken internal documentation links.

## C. Domain and hosting (blocked on the client)

- [ ] Domain chosen, registered and paid for.
- [ ] DNS `A`/`AAAA` records point at the server.
- [ ] `scripts/set-domain.sh` run with the final domain.
- [ ] Caddy site address updated and certificate issued.
- [ ] HTTP redirects to HTTPS; `www` redirects to the canonical host.
- [ ] Firewall allows only 22, 80 and 443; port 8080 is not public.
- [ ] `ADMIN_TOKEN` is strong, stored at mode 600 and not in git.
- [ ] The service starts on boot (`systemctl is-enabled aquamor`).

## D. Verification on the real domain

- [ ] `deploy/smoke-test.sh https://<domain>` passes.
- [ ] `https://<domain>/api/v1/orders/all` without a token returns 401 or 503.
- [ ] A test order is stored and can be read with the admin token.
- [ ] Canonical URL, `og:url`, `og:image`, `robots.txt` and `sitemap.xml` use the
      real domain.
- [ ] WhatsApp/social link preview shows the correct image and title.
- [ ] Real iPhone test: safe areas, menu, WhatsApp hand-off, no horizontal
      scrolling.
- [ ] Low-end Android test: page loads, menu, order builder and gallery.
- [ ] Swahili switch works on every section and in WhatsApp messages.
- [ ] Print the A4 price card and check it fits one page.
- [ ] Home page, price card, planner, label designer and gallery lightbox all
      work with the final domain.

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
