# Maintenance log

A dated record of client decisions, requests and site changes. Add an entry whenever the client confirms a fact, asks for a change, or we make a decision that affects the site.

Keep it short. Link to the issue, PR or file when there is one.

---

## 14 November 2026

### Cross-machine sync and verification

No client-facing change. The work merged on 13 November (PRs #32–#34) was fetched and checked on a second machine:

- Fast-forwarded `main` to `0d13173`; working tree clean and in sync with `origin/main`.
- `gofmt`, `go build ./...`, `go vet ./...` and `go test -race ./...` pass.
- `node --check` on all scripts and `bash -n` on all deploy/scripts pass.
- `deploy/smoke-test.sh` passes all checks against a local server, and a test order saved through `POST /api/v1/orders` appears in the admin list with the server-priced total.
- Removed the orphaned `aquamor.db-shm` / `aquamor.db-wal` files left in the repository root (git-ignored, no database beside them).
- Deleted the merged local and remote feature branches; only `main` remains.

Payments remain off until `FLUTTERWAVE_SECRET_KEY`, `FLUTTERWAVE_WEBHOOK_SECRET` and `PUBLIC_BASE_URL` are set, so the WhatsApp order path is unchanged.

---

## 13 November 2026

### Logo: settled

Decision: use the logo that is already on the site. The header mark (`icon.jpg` / `icon.webp`), `favicon-32.png`, `apple-touch-icon.png` and `icon-512.png` are the approved versions. No original logo file is needed from the client.

Still open from the same item: the official business name and tagline (the client's material has shown more than one version). The site currently uses "Aquamor Purified Drinking Water" and the existing labels.

### Requested from the client (WhatsApp)

Asked for:

- KEBS permit number and a photo of the certificate, if they have one;
- the free-delivery rules (areas and any minimum order).

Message sent:

> Helllo, do you have a KEBS permit?? if yes please send me the permit number and the photo of the certificate and maybe the delivery rules

**Do not add a KEBS number, certification claim, delivery area or minimum order until the client replies and the details are checked.** When the reply arrives:

1. update `docs/status.md` sections 5A and 5D;
2. if a KEBS number is confirmed, add it as a plain fact (no health claims) and keep a copy of the certificate on file;
3. update the delivery copy in `web/index.html` and the Swahili string in `web/js/i18n.js`;
4. rebuild the shop price card only if a price changes;
5. note the change here.

### Client reply (later on 13 November 2026)

Reply received:

> The KEBS is with our supplier we are working on our own next year February we have already gone to the Kebs and we are now in the process,,and also if there is prices for small bottlled water(300ml.500ml,1l,5l)  Branded-bottle pricing and minimum order.
> -  Dispenser-cleaning price.
> [Free-delivery areas and any minimum order
> Oh yes we do them but for now we deal with branded water alone
> Which ranges from 500-800 depending on the type of the bottle

Decisions:

- **KEBS:** the client does not have their own permit or certificate yet; the current certification is with their supplier. Their own application is in process and expected in February next year. Do not add a KEBS number or certification claim. A supplier's certificate is not ours to publish.
- **Prices:** do not publish the Ksh 500–800 range yet. The reply does not say whether it is per bottle or per pack, whether it covers small bottled water or custom-branded bottles, or what the minimum order is.
- **Dispenser cleaning and free delivery:** the client's "yes we do them" is not tied to one question. The cleaning price and the delivery areas/minimum order are still missing.

Follow-up questions sent to the client (13 November 2026); waiting for the reply:

1. Is Ksh 500–800 per bottle or per pack/carton?
2. Which sizes does the Ksh 500–800 range cover?
3. If the range applies to small bottled water, what is the price for each size: 300ml, 500ml, 1L and 1.5L? (The message sent also asked about 5L.)
4. What is the minimum order for branded bottles?
5. What is the dispenser-cleaning price?
6. Which areas are free delivery, and is there a minimum order?
7. What is the official business name and tagline?
8. What are the purification steps, if they want them published?

### Earlier decisions this day

- The client has no Instagram account; no Instagram link will be added (issue #5, closed).
- All remaining media was sent and permission to show the people in the photos was confirmed (issue #6, closed).
- There are no real customer testimonials, so the testimonials section stays hidden (issue #8, closed).
- Update model: managed updates. The client sends requests; the maintainer edits, checks and publishes. No self-service CMS in this phase (issue #16).

### Site changes merged

- #11 refill planner, #12 designer presets and round labels, #13 one-tap reorder, #14 how it works and delivery strip, #15 printable price card, #16 handover documents, #17 quality pass, #18 CI checks.
- Server hardening and launch tooling: environment-based config, loopback default, timeouts, graceful shutdown, security/cache headers, health check, Go tests, CI test step, Caddy/systemd/env/backup/smoke-test examples, and the deployment and launch-checklist docs.
- Responsive 480w/800w JPG/WebP variants for gallery and banner images, with `srcset`/`sizes` and lazy-loaded below-the-fold banners. The full-size lightbox photos are unchanged.
- Removed the unused legacy dashboard and the leftover Next.js `tsconfig.json`.
- CI now runs Go build/vet/tests, gofmt, and JavaScript and shell-script syntax checks on every pull request.

### Site changes: 3D label preview, map and acknowledgment

- The label designer bottle now has a 3D look and can be turned by dragging or with a slider; presets, colours, round/rect styles, text fitting and the WhatsApp hand-off are unchanged.
- The contact section now shows a Google Maps embed using the same location as the existing "Get directions" link; the link is unchanged and the CSP allows Google Maps frames.
- Any WhatsApp or email link now shows an on-page thank-you acknowledgment in English and Swahili. This is not an automated email or WhatsApp reply.
- Open: online payment and real automated replies need a provider, merchant account and credentials from the client (see status section 5A item 4).

### Online payments: hosted checkout (13 November 2026)

Decision: the client chose a hosted checkout. The site now uses Flutterwave hosted checkout; no card or M-Pesa data touches the Aquamor server.

- The order summary shows a Pay online button only when payment is configured. The dialog collects name, phone, receipt email and delivery area or pickup name, saves the order, then redirects to Flutterwave's secure page.
- `web/payment.html` verifies the payment when the customer returns and shows paid, failed, pending or an error.
- The server re-prices the order, checks the provider reference, currency and amount, and verifies the webhook `verif-hash` before marking an order paid.
- New order fields are added automatically at startup: `customer_email`, `payment_provider`, `payment_reference`, `mpesa_receipt_number` and `paid_at`.
- Settings are `FLUTTERWAVE_SECRET_KEY`, `FLUTTERWAVE_WEBHOOK_SECRET` and `PUBLIC_BASE_URL`. Without all three, the button stays hidden and WhatsApp ordering is unchanged.
- Tests cover checkout creation, provider errors, verification, amount/currency checks, webhook signatures and the unconfigured state.
- Open: the client must supply an approved Flutterwave merchant account and credentials. Online payment is limited to refills and 20L bottles because other bottled-water prices are not confirmed.

### Order capture, payment fixes and admin view (13 November 2026)

- The order builder now saves each order to `POST /api/v1/orders` before the WhatsApp hand-off. Saves are only remembered after a successful response, and two different pending orders can be saved without blocking each other.
- Checkout keeps the quantities and delivery mode as they were at submission, so reorder data matches the order that was actually paid for.
- Payment verification leaves an order `pending` while the provider is still deciding, and does not overwrite a `paid` order that a webhook confirmed while verification was in flight.
- Added the authenticated admin orders page (`web/admin.html`, `web/js/admin.js`, `web/css/admin.css`). It asks for `ADMIN_TOKEN`, keeps it in `sessionStorage` for the tab only, shows orders with line items and payment state, filters and searches them, and exports CSV. It is `noindex`, disallowed in `robots.txt`, served with `X-Robots-Tag: noindex` and `no-store`, and not linked from the public site.
- The smoke test checks `/admin.html` returns 200 and is `noindex`. `node --check`, Go build, vet and tests pass.
- Docs updated: this log, `status.md` section 24, `launch-checklist.md`, `web/README.md` and the maintainer guide.

### Self-hosted fonts and write-endpoint rate limits (13 November 2026)

- Bricolage Grotesque and Figtree are now served from `web/assets/fonts/` as Latin variable WOFF2 files, declared with `@font-face` in `web/css/styles.css`. The Google Fonts links were removed from `index.html`, `admin.html` and `payment.html`, and those pages preload the local fonts instead.
- The CSP no longer allows Google Fonts: `style-src` drops `fonts.googleapis.com` and `font-src` is `'self'`. Google is still allowed for the map frame only.
- Added an in-memory fixed-window rate limiter (`backend/internal/ratelimit`): 20 requests/minute per IP on `POST /api/v1/orders` and 10 on `POST /api/v1/payments/checkout`, returning `429` with `Retry-After`. It uses the last `X-Forwarded-For` value appended by Caddy and falls back to the connection address.
- Added unit tests for the limiter and documented the change in `status.md` section 25 and the launch checklist. Build, vet, tests, the font/CSP checks and the local smoke test pass.

### Request logging (13 November 2026)

- Added `backend/internal/requestlog`, an outer middleware that writes one line per request with method, path, status, response size, duration and client IP.
- Query strings and headers are not logged, and customer fields never appear. The admin token, admin search terms and payment parameters therefore stay out of the log.
- Wired around the security and cache middleware, so rate-limited `429` responses are recorded too. Tests cover status/size, the `200` default, query-string exclusion and the forwarded client address.
- Updated the deployment security notes and the launch checklist; the local smoke test and a privacy check of the log output pass.

---

## Open items waiting on the client

| Item | Issue | Status |
|---|---|---|
| Small bottled-water prices (300ml, 500ml, 1L, 1.5L) | #4 | Reply ambiguous; confirm unit |
| Branded-bottle pricing and minimum order | #4 | 500–800 quoted; unit/minimum unclear |
| Dispenser-cleaning price | #4 | Waiting |
| Free-delivery areas and minimum order | #4 | Requested 13 Nov; not answered |
| KEBS permit number and a photo of the certificate | #4 | Answered: with supplier; own due Feb |
| Purification steps (only if the client has them) | #4 | Requested 13 Nov; waiting |
| Official business name and tagline | #4 | Waiting |
| Facebook page is a business page, not a personal profile | #10 | Waiting |
| Confirm bottled-water sizes (flyer vs client text) | #10 | Waiting |
| Native-speaker Swahili review | #10 | Waiting |
| Regular 20L bottle: is water included? | #10 | Waiting |
| Original cleaning before/after stills | #9 | Waiting |
| More gallery content (shop front, graduation/corporate) | #7 | Waiting |
