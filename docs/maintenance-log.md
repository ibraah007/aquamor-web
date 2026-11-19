# Maintenance log

A dated record of client decisions, requests and site changes. Add an entry whenever the client confirms a fact, asks for a change, or we make a decision that affects the site.

Keep it short. Link to the issue, PR or file when there is one.

---

## 18 November 2026

### Content corrections: hours, prices wording, photos and label designer

- **Opening hours** are now every day, 8:30am to 9:00pm. The contact section,
  FAQ, structured data, price card and the live "Open now" table were updated.
- **New-bottle prices include the bottle and the water.** The 16 November entry
  below said water was sold separately; the site now describes the listed price
  as including both. The prices themselves did not change.
- **Location** now leads with "Ngara, Nairobi" before the Musindi Road directions.
- **Gallery captions** were corrected to match the photos.
- **Label designer:** eight colours (black, yellow, light green, dark green, red,
  ocean blue, white, blue), a new oval style, dark text on light labels, and an
  outline on the white swatch. The WhatsApp request includes the label style.

See [`status.md`](status.md) section 33 for details.

### New photos, colour-blind help and the full wordmark

- Added 13 of the client's new photos and a 24-second delivery clip to the
  gallery and the rotating backgrounds. Images have JPG + WebP copies and
  480/800px variants (#40, #41).
- The hero now rotates through nine photos; the order, design and contact bands
  have seven, seven and six.
- Every label-colour swatch now shows its name, the selected colour is written
  out, and a help panel shows the client's photo of the shop's bottle caps
  (#42).
- The header wordmark now reads Aquamor Purified Drinking Water (#43).
- Corrected the two photo captions to the client's wording: staff cleaning a
  bottle, and staff wiping a bottle after refilling (#39).

See [`status.md`](status.md) section 34 for details.

### More colour and photos across the page

- The client wanted the site to feel more marketable, with colour in more
  sections and fading photos all over. The slideshow now changes every
  5 seconds, and the How it works, cleaning and purification sections have
  their own photo backgrounds.
- The planner, gallery, about and FAQ now use tinted backgrounds with aqua,
  sun and leaf glows; step numbers, card edges, tabs and buttons carry the same
  accents. Added a slim scroll-progress line and hover feedback on cards/photos.
- Tracked in #44 and described in [`status.md`](status.md) section 35.
- Follow-up the same day: added the four benefit cards under the hero, a bright
  closing call-to-action band, a lighter hero overlay and `?v=` version queries
  on every CSS/JS link so cached files cannot hold the old design. Tracked in
  #45 and described in section 36.

---

## 17 November 2026

### Final QA on the published preview

Ran a full pass on <https://ibraah007.github.io/aquamor-web/> and on the local Go
server. No functional defects were found. Results are recorded in
[`status.md`](status.md) section 29.

- Lighthouse (mobile): accessibility 100, SEO 100, best practices 96, performance
  84. The single best-practices deduction is the `/api/v1/payments/config` 404,
  which only happens on the static preview because there is no Go API there.
- Every width from 320px to 1600px scrolls only vertically; the mobile menu,
  language switch, order builder, reorder, planner, designer, gallery filters,
  lightbox, FAQ and Copy buttons all work, in both languages.
- `deploy/smoke-test.sh` passes against the local server, the Go tests pass, and
  the A4 price card prints on one page.
- **Change:** banner slides after the first now keep their image sources in
  `data-src`/`data-srcset` and load when the slideshow is about to show them.
  First-load image bytes at 390px fell from 626 KB to 287 KB.
- The `gh-pages` preview was refreshed so it matches `web/` again.
- **Docs:** added [`go-live-runbook.md`](go-live-runbook.md) (the ordered
  domain/DNS/TLS/verification/handover sequence with an owner for each step) and
  [`client-review-message.md`](client-review-message.md) (a ready-to-send email and
  WhatsApp message with the preview link and the remaining client questions).
- **Review corrections (PR #38):** the runbook now uses the client's chosen domain
  throughout, reuses the prepared `/opt/aquamor` checkout, checks every JS and
  shell file separately, names a supported item for the test order, and records
  two limits: orders containing an unpriced item (such as 500ml) stay only in
  WhatsApp, and a backup restore drops newer orders unless they are reconciled
  first. The client message calls the preview public and explains that the order
  draft is addressed to the Aquamor business number. The deployment steps also
  update the installed Caddyfile and `PUBLIC_BASE_URL`, restart the service after
  a new binary, and preserve the SQLite WAL before a restore.

### Scroll reveal animations

Added a quiet scroll-reveal pass: section headings, cards, steps, gallery photos
and contact/map blocks fade and rise in as they reach the viewport, with a short
stagger and a small aqua-to-gold accent under each heading. The hidden state is
only applied in JavaScript to elements that start below the fold, so the page
stays readable without JavaScript and nothing flashes on load. Reduced-motion is
respected and every element is unobserved after it appears. Recorded in
`status.md` section 31 and the front-end README.

### Plain-language launch ask and client call script

The client is expected to phone rather than reply in writing. Added
[`client-call-script.md`](client-call-script.md): each of the four launch items
(approval, domain, business details, hosting) explained in the client's terms,
with the questions he is likely to ask, short answers, a "what happens after each
item" table, a six-fact cheat sheet for the call, and a written summary to send
afterwards.

The client-facing documents were aligned with it:

- [`client-information-needed.md`](client-information-needed.md) now includes
  hosting as a launch item (three ownership options, cost guide, warning against
  e-commerce packages) and the reply form asks for the hosting choice and monthly
  manager. The PDF was regenerated.
- [`client-review-message.md`](client-review-message.md): the email now says
  Aquamor registers the domain with our guidance instead of "we will register it",
  both email and WhatsApp versions ask about hosting, and the Swahili message was
  updated to match.
- [`go-live-runbook.md`](go-live-runbook.md) points to the call script and records
  the server owner as agreed between client and maintainer.

## 16 November 2026

### Client confirmed all outstanding details

The client sent the confirmed business details in writing. Changes made:

- **Business name and tagline:** AQUAMOR PURIFIED DRINKING WATER; "Hydrated? Worry no more! Aquamor has got you covered. Your satisfaction is our priority." Added to the hero, About section and structured data.
- **Bottled water:** prices vary by quantity, so the site now asks for a quotation instead of showing fixed prices. Sizes: 200ml, 300ml, 500ml, 1L, 1.5L, 2L, 3L, 5L, 10L and 20L.
- **Refills:** 5L Ksh 50, 10L Ksh 100, 20L Ksh 200; the price includes water.
- **New bottles (bottle purchase; water refills separate):** 20L regular Ksh 450–550, 20L hard Ksh 1,000–1,700, 10L Ksh 200–250, 5L Ksh 100.
- **Customised bottles:** churches, weddings, graduations, anniversaries, corporate events and other occasions; customisation Ksh 500–1,000; minimum order 10 bales.
- **Dispenser cleaning:** Ksh 500–1,000 depending on the dispenser and cleaning required.
- **Delivery:** Ngara, Parklands, Pangani, Gikomba, Westlands, Nairobi CBD and parts of Eastleigh. Free delivery depends on distance and quantity; customers are asked to confirm.
- **Pumps and dispensers:** manual and electric pumps for 10L and 20L bottles; manual and electric dispensers.
- **Purification process:** eight steps published without individual sediment-filter micron sizes.
- **Domain:** preferred `aquamorwater.co.ke`, to be registered under AQUAMOR PURIFIED DRINKING WATER with Aquamor keeping ownership and control and paying renewal. The site, `robots.txt`, sitemap, Caddyfile and environment example now use it.
- **Contact:** existing Gmail `aquamor2020@gmail.com`; Facebook page is AQUAMOR NGARA; TikTok link kept.

Updated `web/index.html`, `web/js/order.js`, `web/js/i18n.js`, `web/price-card.html`, the hard-bottle catalog name, and the project docs. Bottled-water prices stay as quotations, as requested. New-bottle ranges use the lower bound for the on-page estimate only.

**Online payments are deferred to a later phase** at the project owner's request. The Flutterwave code stays in place but no merchant account, API keys or webhook secret are needed for launch; the Pay online button stays hidden. The two remaining permissions (other customers' labels and the Paybill display) were cleared for this build by the project owner. The Swahili copy was reviewed and awkward phrases refined.

### Static preview published

A static copy of `web/` is now published from the `gh-pages` branch at <https://ibraah007.github.io/aquamor-web/> using GitHub Pages. It is a preview of the public pages only: the Go API, admin orders page and online payment pages are not available there. The production site still needs the domain registration, DNS/TLS and a server for the Go API.

### Open items after this update

Only the domain registration/handover and a native-speaker Swahili read-through remain. Online payments are deferred.

## 14 November 2026

### Cross-machine sync and verification

No client-facing change. The work merged on 13 November (PRs #32–#34) was fetched and checked on a second machine:

- Fast-forwarded `main` to `0d13173`; working tree clean and in sync with `origin/main`.
- `gofmt`, `go build ./...`, `go vet ./...` and `go test -race ./...` pass.
- `node --check` on all scripts and `bash -n` on all deploy/scripts pass.
- `deploy/smoke-test.sh` passes all checks against a local server, and a test order saved through `POST /api/v1/orders` appears in the admin list with the server-priced total.
- Removed the orphaned `aquamor.db-shm` / `aquamor.db-wal` files left in the repository root (git-ignored, no database beside them).
- Deleted the merged local and remote feature branches; only `main` remains.

### WhatsApp icon on the "WhatsApp us" buttons

Every WhatsApp call-to-action button now shows the WhatsApp glyph before its label: the floating "WhatsApp us" button, Send order, Ask about this plan (refill planner), Request this design (label designer), Book a cleaning, Chat on WhatsApp in the contact section, and the "WhatsApp us" button on the payment status page. The icon is inline SVG, inherits the button text colour (so it works in light and dark mode), is hidden from screen readers, and adds no extra file request. The labels are unchanged, so the existing Swahili translation still applies; the buttons are flex rows so the icon stays beside the text when a label wraps on a narrow phone.

Payments remain off until `FLUTTERWAVE_SECRET_KEY`, `FLUTTERWAVE_WEBHOOK_SECRET` and `PUBLIC_BASE_URL` are set, so the WhatsApp order path is unchanged.

### Consolidated questions for the client

The 13 November follow-ups and everything else still open in `status.md` section 5 were consolidated into one message:

1. Ksh 500–800 — per bottle or per pack/carton?
2. Does that range cover small bottled water or branded bottles?
3. If small bottled water, the price per size: 300ml, 500ml, 1L, 1.5L (and 5L if sold).
4. Minimum order for branded bottles.
5. Dispenser-cleaning price.
6. Free-delivery areas.
7. Minimum order for free delivery.
8. Official business name and tagline.
9. Purification steps, if they want them published.
10. Facebook page is a business page, not a personal profile.
11. Which sizes are actually sold (flyer shows 500ml, 5L, 10L, 20L; their text also lists 300ml, 1L, 1.5L).
12. Regular 20L bottle at Ksh 450 — is water included?
13. Are the water pumps for sale?

A Swahili version was prepared alongside it so the client can read the questions in their preferred language. Nothing is published from these answers until the client replies in writing (see the open-items table below).

The English sheet for the client is [`client-information-needed.md`](client-information-needed.md): the same questions as a plain document, with the domain and online-payment items, a "not published until confirmed" table and a copy-and-fill reply form. It is also exported to [`client-information-needed.pdf`](client-information-needed.pdf) to attach to an email. **Update (16 November 2026):** the sheet and PDF now list the confirmed details and only the remaining domain, payment, Swahili and permission items.

Client email: **Aquamor2020@gmail.com** (already on the site contact section, the printable price card and the structured data). It is the address to send the sheet to. It is a Gmail address, so there is no domain email to preserve if `aquamor.co.ke` is registered; a business address such as `info@aquamor.co.ke` can be added later if they want one.

`docs/screenshots/` was regenerated from the current build so the preview sent with the sheet matches the live code: `desktop.png` (1280×1000), `mobile.png` (390×1000) and `mobile-menu.png` (390×1000, menu open). All three show the current WhatsApp icons.

### Domain

Plan agreed: register a `.co.ke` domain in the **client's name**, with the client paying the yearly renewal; the maintainer only holds DNS/technical access. Do not register it in the maintainer's personal account and do not leave ownership `[to agree]` past handover — this is the item that causes the most disputes later.

- KeNIC is the registry; register through a KeNIC-accredited registrar. A `.co.ke` needs no supporting documents for most registrations (unlike `.go.ke`/`.ac.ke`).
- Cost: roughly Ksh 1,000–1,500 for the first year and Ksh 1,200–1,350 to renew on Kenyan registrars (checked November 2026). Check the renewal price, not just the first-year offer, and enable auto-renew.
- Availability on 14 November 2026 (no DNS delegation): `aquamor.co.ke`, `aquamor.ke`, `aquamorwater.co.ke`, `aquamorpurifiedwater.co.ke`. Confirm at the registrar's checkout before promising one.
- Preferred: `aquamor.co.ke` (already the placeholder in the site and docs). Fallbacks: `aquamor.ke`, `aquamorwater.co.ke`.
- Ask the client for: the chosen name, the registrant name/email/phone to use, who pays, and whether they already own a domain or use email on one (copy the `MX` records before changing nameservers).
- When confirmed: point `A`/`AAAA` at the server, run `scripts/set-domain.sh <domain>`, set the Caddy address and `PUBLIC_BASE_URL`, then work through `launch-checklist.md` section C on the real domain.

**Update (16 November 2026):** the client confirmed `aquamorwater.co.ke` and answered the questions above; see the 16 November entry at the top.

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

### Manual slides, lit hovers and Google reviews (19 November 2026)

- Each photo section now has a manual control pill (previous, counter, next). The arrows work by mouse, touch and keyboard, and a sideways touch swipe also changes the photo. The automatic 5-second timer pauses while the section is hovered, focused or touched, and resumes afterwards.
- Slide changes fade in from a brighter, slightly larger frame and settle. Hovering or focusing a photo section flashes it brighter, then holds a gentler lift while the interaction lasts. Off-screen sections stop cycling and only load the next photo when it is needed.
- Gallery tiles brighten, lift and pass a band of light across the photo on hover or keyboard focus. Reduced-motion keeps the manual controls but disables the automatic slideshow, flash and transitions.
- Replaced the hidden review placeholder with three real reviews from the business's public Google listing, the 4.9/5 rating from 24 reviews (November 2026) and a link to the full listing. Reviewer names are not shown; each card is credited "Google review · month year". Wording and dates come from Google; the text was lightly edited for spelling, punctuation and spacing, with the meaning unchanged.
- The contact section now has a "Follow us" card with the Facebook (AQUAMOR NGARA) and TikTok links and their brand icons. No "main page" claim or follower counts are shown; the links are left to speak for themselves.
- All new labels have Swahili translations. `node --check` and a headless Chromium pass confirmed the controls, the manual photo change, the Swahili labels, the three review cards and both social icons at 390px and 1300px.
- Cache versions are now `?v=20261119a` in `index.html`, `admin.html` and `payment.html`.
- Docs updated: `status.md` section 37, this log and `web/README.md`.

### Background photos automatic again (19 November 2026)

- The client did not want manual sliding options on the fading background images. Removed the previous / counter / next controls, the sideways swipe, the keyboard arrow handling and the pause-on-hover behaviour.
- The photos now fade through automatically every 5 seconds, with the same brighten-in effect. Hovering a section still brightens it but does not stop or change the slideshow. Reduced-motion shows only the first photo.
- Removed the unused control styles and Swahili labels. Gallery hover and reviews are unchanged. The social card now shows only the Facebook and TikTok links, without the "main page" wording.
- Cache versions are now `?v=20261119c`. `node --check` passes and a headless Chromium pass confirms the controls are gone and the slides advance automatically.
- Docs updated: `status.md` section 38, this log, the root README and `web/README.md`.

---

## Open items waiting on the client

| Item | Issue | Status |
|---|---|---|
| Small bottled-water prices (300ml, 500ml, 1L, 1.5L) | #4 | Resolved 16 Nov: quoted by quantity, no fixed prices |
| Branded-bottle pricing and minimum order | #4 | Resolved 16 Nov: customisation Ksh 500–1,000; minimum 10 bales |
| Dispenser-cleaning price | #4 | Resolved 16 Nov: Ksh 500–1,000 |
| Free-delivery areas and minimum order | #4 | Resolved 16 Nov: listed areas; free delivery depends on distance and quantity |
| KEBS permit number and a photo of the certificate | #4 | Answered: with supplier; own due Feb; no claim published |
| Purification steps (only if the client has them) | #4 | Resolved 16 Nov: eight steps published without micron sizes |
| Official business name and tagline | #4 | Resolved 16 Nov: AQUAMOR PURIFIED DRINKING WATER; tagline supplied |
| Facebook page is a business page, not a personal profile | #10 | Resolved 16 Nov: AQUAMOR NGARA business page |
| Confirm bottled-water sizes (flyer vs client text) | #10 | Resolved 16 Nov: ten sizes from 200ml to 20L |
| Native-speaker Swahili review | #10 | Waiting |
| Regular 20L bottle: is water included? | #10 | Resolved 16 Nov: bottle purchase; water refill separate |
| Domain registration and handover (`aquamorwater.co.ke`) | #4 | Preferred domain set in the site; registration/handover still to complete |
| Original cleaning before/after stills | #9 | Waiting |
| More gallery content (shop front, graduation/corporate) | #7 | Waiting |
