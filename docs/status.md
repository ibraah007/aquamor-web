# Aquamor Purified Drinking Water: Website Project Status

**Updated:** 16 November 2026
**Project start:** 10 November 2026
**Target delivery:** about 10 December 2026 (one month)
**Repository:** the site is in `web/`; the Go API is in `backend/` (see the root `README.md`)
**Preview:** run locally with `cd web && python3 -m http.server 8000`, or run the Go server and open http://localhost:8080. No public URL yet.

---

## 1. Project summary

A mobile-first, one-page website for Aquamor, a purified drinking water business on Musindi Road, Ngara, Nairobi (established 2019). It serves households, businesses, churches and organizations with water refills, bottled water, custom-branded bottles for events, and dispenser cleaning.

**Working rule:** the site states only what the client has said or what is visible in their own photos and flyers. Anything unconfirmed is left out or flagged below.

---

## 2. Confirmed business facts (source: client text, flyers, shop poster)

| Item | Detail |
|---|---|
| Official business name | AQUAMOR PURIFIED DRINKING WATER |
| Tagline | "Hydrated? Worry no more! Aquamor has got you covered. Your satisfaction is our priority." |
| Established | 2019 |
| Address | Along Musindi Road, opposite Ngara Girls Secondary School, next to Ngara Heights building, Nairobi |
| WhatsApp | 0712 140 006 |
| Phone calls | 0113 952 100, 0736 054 182 |
| Email | aquamor2020@gmail.com |
| Opening hours | Mon to Sat 8:30am to 8:30pm; Sun 12:00pm to 7:00pm |
| Refill prices | 5L Ksh 50, 10L Ksh 100, 20L Ksh 200 (includes water) |
| New bottle purchase prices | 20L regular Ksh 450–550; 20L hard Ksh 1,000–1,700; 10L Ksh 200–250; 5L Ksh 100 |
| Bottled water | 200ml, 300ml, 500ml, 1L, 1.5L, 2L, 3L, 5L, 10L and 20L; prices vary by quantity, quoted on request |
| Customised bottles | Churches, weddings, graduations, anniversaries, corporate events and other occasions; customisation Ksh 500–1,000; minimum 10 bales |
| Dispenser cleaning | Ksh 500–1,000 depending on the dispenser and cleaning required |
| Pumps and dispensers | Manual and electric pumps for 10L and 20L bottles; manual and electric dispensers |
| Services | Refills, bottled water, customised bottles, dispenser cleaning and sanitizing, pumps and dispensers |
| Delivery | Available in the listed areas; free delivery depends on distance and quantity |
| Delivers to | Ngara, Parklands, Pangani, Gikomba, Westlands, Nairobi CBD and parts of Eastleigh |
| Domain | aquamorwater.co.ke (preferred; registration under AQUAMOR PURIFIED DRINKING WATER) |
| Payment | Equity Paybill 247247, account 395210 (account name: Aquamor Water) |
| Facebook | AQUAMOR NGARA business page: https://www.facebook.com/profile.php?id=100066493751902 |
| TikTok | https://www.tiktok.com/@aquamor2020 (provided by the client on 12 November 2026) |

---

## 3. What has been built

### Structure and design
- [x] One-page site, mobile-first, tested at 320, 360, 390, 768, 900, 1024, 1280 and 1600 px with no sideways scrolling
- [x] Light and dark mode support
- [x] Blue and white brand look; header with logo icon (cropped from the client's flyer)
- [x] Fading photo backgrounds in the hero, the order banner, the branded-bottles banner and the contact section (respects "reduced motion")
- [x] Sticky "WhatsApp us" button with the WhatsApp icon on every screen
- [x] Mobile navigation menu with keyboard and Escape-key support
- [x] Skip-to-content link and visible focus styles

### Sections and features
- [x] **Hero:** headline, official tagline, buttons, key facts, live "Open now / Closed now" line (Nairobi time)
- [x] **Build your order:** refill sizes, bottled water in ten sizes (quoted by quantity) and new bottles in four sizes with steppers, delivery/pickup choice, running total, one-tap WhatsApp message with the order already written, and a one-tap reorder of the last order
- [x] **Refill planner:** enter the number of people at home and choose 5L, 10L or 20L; the site estimates weekly litres, refills and cost, and links to WhatsApp with the plan already written
- [x] **Label designer:** pick a Wedding, Church, Graduation or Corporate preset, type an event name and message, pick a color, a rectangular or round label and one of ten bottle sizes, see it on a bottle, send the request on WhatsApp; customisation and minimum-order details are shown
- [x] **Printable price card:** a standalone A4 page for the shop counter with the confirmed refill and new-bottle prices, bottled-water quotation, customisation and cleaning charges, delivery areas, tagline and WhatsApp number (linked from the docs; not part of the one-page site)
- [x] **How it works and who/where we deliver to:** three steps (choose your water, order on WhatsApp, delivery or pickup), the business types served and the delivery areas; free delivery depends on distance and quantity
- [x] **Dispenser cleaning and sanitizing:** four service points, before / cleaning / after photos, the Ksh 500–1,000 price range, book on WhatsApp
- [x] **Our work (gallery):** 26 real photos, three short video clips, filter buttons (Bottles, Events, Refills, Team), tap to enlarge
- [x] **Water purification process:** the client's eight steps, without individual micron sizes
- [x] **About:** company story, official tagline, pumps and dispensers, with a team photo
- [x] **FAQ:** ten tap-to-open questions, answered only with confirmed facts
- [x] **Contact:** address, hours, WhatsApp, calls (two numbers), email, Paybill with Copy buttons, Get directions (Google Maps), Facebook (AQUAMOR NGARA) and TikTok buttons
- [x] **Testimonials:** built but hidden until real quotes are added

### Languages and search
- [x] **English / Swahili switch** (SW / EN button), remembered between visits; WhatsApp messages are written in the chosen language, including product names and Delivery/Pickup
- [x] Page title, description, canonical URL, Open Graph and Twitter card tags, including a 1200x630 preview image
- [x] Business data (address, hours, phone, Facebook, FAQ) in structured form for search engines, with `url`, `image`, `priceRange` and `areaServed`
- [x] Favicon and Apple touch icon

### Photos used (all from the client's own material)
Refill station, glass bottle rows, custom church labels, Wave360 label bottle, shelf of "50 years" labels, bottle stacks, team photos, pump photos, dispenser cleaning before/after, wedding label designs, bottled water on the shop shelves, 20L containers at the counter, a church-event bottle setup, a wedding-reception hall, a customer carrying a refill, a team member refilling a 20L container, a close-up of a custom-labelled bottle, an event table, team members checking and washing containers, 20L bottle tops, and stacked blue containers. Each photo has a JPG and a WebP copy; the browser downloads only one. Gallery photos are lazy-loaded. Three short muted clips (event setup, custom-labelled event bottles, and dispenser cleaning) use `preload="none"` so they download only when played.

---

## 4. Decisions and corrections made

- Removed the wedding table photo from the gallery (it came from a video and may be a stock image). It remains only as a decorative background in the branded-bottles banner.
- Cropped the right side off two photos so the Paybill sign is not visible in them.
- No names or job titles are given to anyone in the photos (nobody has been confirmed as the owner).
- No health, certification or "100% pure" claims have been added. The flyers' claims about germs and bacteria were left out.
- Hard bottle price was first set to Ksh 1,400 (clear flyer and the latest flyer agree; an earlier blurry rate card looked like 1,500). On 16 November 2026 the client confirmed a range of Ksh 1,000–1,700; Ksh 1,400 remains within it and is the catalog value.
- The client's official TikTok, [@aquamor2020](https://www.tiktok.com/@aquamor2020), is linked in the contact section and in the structured data. A personal-looking TikTok account (@benzi254) was **not** linked.
- The client has no Instagram account, so no Instagram link was added (issue #5).
- The logo is settled: the existing header mark (`icon.jpg` / `icon.webp`) and the icon files (`favicon-32.png`, `apple-touch-icon.png`, `icon-512.png`) are approved; no original logo file is being pursued (13 November 2026).

---

## 5. Open issues

### A. Blocking: needed from the client
1. **Domain registration and handover:** the client prefers `aquamorwater.co.ke` and wants registration under AQUAMOR PURIFIED DRINKING WATER, with Aquamor keeping ownership and control. Confirm who completes the registration and the DNS/hosting handover.

All other client details requested on 13 November 2026 were confirmed on 16 November 2026; see section 28.

**Deferred to a later phase (not required for launch):** online payments and automated replies. The hosted-checkout code stays in place but remains switched off until Flutterwave credentials are supplied later. The site currently shows an on-page acknowledgment instead of an automated reply.

### B. Permissions
- [x] Showing the staff and customers in the photos, including the new event, shop and delivery photos (confirmed by the client on 12 November 2026)
- [x] Showing other customers' labels (Wave360 and the church labels) — cleared for this build by the project owner on 16 November 2026
- [x] Water pumps: the client confirmed on 16 November 2026 that they sell manual and electric pumps for 10L and 20L bottles, plus manual and electric dispensers
- [x] Showing the Paybill number online — cleared for this build by the project owner on 16 November 2026; the number is already on the client's flyers

### C. Content still missing
- Shop front on Musindi Road, taken from the street
- Graduation and corporate event photos (church-event and wedding-reception photos are now added)
- A close-up of a glass bottle with its label
- Two or three real customer comments
- Original before/after stills for the cleaning service (the current ones are small crops from a flyer; a short cleaning video clip is now on the site)
- A named person for the About page, only if the client wants one

### D. To verify
- **KEBS:** the client's certification is currently with their supplier; their own permit is in process and expected in February next year. No permit number or certificate is available to publish, so the site makes no certification claim. The purification steps were supplied on 16 November 2026 and are published without the individual sediment-filter micron sizes.
- **Swahili text:** internal review completed on 16 November 2026 (awkward phrases refined). A native-speaker read-through is still recommended before launch.
- **Domain ownership:** confirm that `aquamorwater.co.ke` is registered under AQUAMOR PURIFIED DRINKING WATER and that the client retains the registrar login and renewal control. Aquamor will pay the annual renewal.

### E. Handover tasks (week 4)
- Client review and the agreed revision rounds
- Register the preferred domain `aquamorwater.co.ke` under AQUAMOR PURIFIED DRINKING WATER and set up hosting using [`deployment.md`](deployment.md)
- Move the site to the domain (this is when the Google search data matters most)
- Hand over, plus the client change-request sheet and the maintainer guide
- Work through [`launch-checklist.md`](launch-checklist.md) before announcing the domain
- Online payment activation is deferred to a later phase; no launch action is needed beyond keeping the payment values empty.

---

## 6. Ideas that can be built next (no client input needed)
1. ~~Refill planner (people in the home, then suggested refills per week and cost)~~ — built (see section 12)
2. ~~Designer presets (Wedding, Church, Graduation, Corporate) and a round label style~~ — built (see section 15)
3. ~~"Reorder my last order" in one tap~~ — built (see section 16)
4. ~~"Who we deliver to" strip and a three-step "How it works"~~ — built (see section 14)
5. ~~Printable A4 price card for the shop counter~~ — built (see section 13)
6. ~~"How to update your site" guide~~ — built (see section 17)
7. ~~Final speed, contrast and small-phone pass~~ — built (see section 18)

---

## 7. Suggested timeline

| Week | Focus |
|---|---|
| 1 (10 to 16 Nov) | Core site built. Send the client the details request. Keep collecting photos. |
| 2 (17 to 23 Nov) | Add the client's answers: logo, prices, delivery rules, social links. Build the refill planner and designer upgrade. |
| 3 (24 to 30 Nov) | Polish, Swahili review, testimonials, phone testing, final photos. |
| 4 (1 to 10 Dec) | Client review and revisions, domain and hosting, handover. |

---

## 8. Agreement points to settle with the client (in writing)
- Number of revision rounds included
- Who gives final approval
- Deposit and payment dates
- Who pays for the domain and hosting
- Who handles updates after launch, the change-request process and response times
- What is included in ongoing maintenance and what is charged separately
- Whether the optional business social media setup is a separate paid service

---

## 9. Repository integration and fixes (11 November 2026)

The standalone download has been imported into this git repository.

- `web/` now holds the site; the Go server serves it at `/`.
- Fixed: gallery filters crashed the order builder (both used the `.tab` class).
- Fixed: the navigation was hidden with no replacement below 1000px; there is now a mobile menu.
- Fixed: Swahili WhatsApp messages mixed in English product names and "Delivery".
- Fixed: "Estimated total" omitted bottled water; it is now labelled and explained.
- Fixed: label-designer text was clipped; it now shrinks to fit and re-fits when fonts load or the language changes.
- Added: `og:image`, `og:url`, Twitter card, canonical URL, favicon and apple-touch icon.
- Added: WebP + JPEG fallback for photos, lazy-loaded gallery images, hero preload and the 700 font weight.
- Added: live cart announcements, translated `aria-label`s and placeholders, skip link, keyboard-friendly menu, no-JS order fallback.
- Added: the open/closed line refreshes every minute; safe-area insets are handled on the header and footer.
- Backend: prices now come from the server catalog (client prices are ignored); order and items are written in one transaction; request size and quantity limits; the orders list requires `ADMIN_TOKEN`; SQLite uses WAL and a busy timeout; the legacy dashboard is no longer served.
- Docs: root `README.md` and `web/README.md` updated; this status file updated.

### Still to do before launch
- The preferred domain `aquamorwater.co.ke` is now used in `<head>`, `robots.txt`, `sitemap.xml` and the deployment examples; register it and point DNS at the server.
- Put the Go server behind a TLS-terminating reverse proxy before any remote use and block direct access to port 8080; the admin bearer token and customer PII must not travel over plain HTTP.
- Get the client answers listed in section 5.
- Confirm the WhatsApp link preview after deployment.
- Test on a real iPhone (safe area, menu, WhatsApp hand-off) and a low-end Android phone.
- [x] Connected the order builder to `POST /api/v1/orders` (WhatsApp remains the final hand-off) and added an authenticated admin orders view (see section 24).
- [x] Deferred the banner slideshow images that are below the fold; only the first image loads eagerly.

## 10. Media and social update (12 November 2026)

- Added the client's official TikTok (https://www.tiktok.com/@aquamor2020) to the contact section and the `LocalBusiness` `sameAs` data.
- Reviewed the new client media and added six photos to the gallery: bottled water on the shop shelves, 20L containers at the counter, bottles ready for a church event, a wedding-reception hall, a customer carrying a refill and a team member refilling a 20L container. Each has JPG and WebP copies.
- The gallery now has 26 photos and three short video clips. Filter categories and English/Swahili captions were updated for the new tiles.
- The client confirmed permission on 12 November 2026 to show the people in the new photos (the customer and the team member), as well as the existing staff photos.

## 11. Additional photos and video clips (12 November 2026)

- Added six more photos from the client's media folder: a close-up of a custom-labelled bottle, an event table, a team member checking containers, a team member washing a container, a close-up of 20L bottle tops, and stacked blue 20L containers.
- Added three short muted video clips: event setup, custom-labelled event bottles, and dispenser cleaning. All use `preload="none"` and a poster image so they do not download until played.
- The gallery now has 26 photos. English and Swahili captions were added for all six new tiles.
- The TikTok link (`https://www.tiktok.com/@aquamor2020`) is already present in the contact section and the `LocalBusiness` `sameAs` data.

## 12. Refill planner (13 November 2026)

- Added a refill planner between the order builder and the label designer. The visitor enters the number of people at home and chooses 5L, 10L or 20L; the site estimates weekly litres, the number of refills (rounded up to full containers) and the weekly cost.
- Uses only confirmed refill prices: 5L Ksh 50, 10L Ksh 100, 20L Ksh 200.
- The estimate assumes about 2 litres of water per person per day; this assumption is stated under the result. No health or quality claims were added.
- Works in English and Swahili (including the numbers and the WhatsApp message), and the plan can be sent to WhatsApp in one tap.
- No backend; the logic is in `web/js/planner.js`.

## 13. Printable price card (13 November 2026)

- Added [web/price-card.html](../web/price-card.html), a standalone A4 page for the shop counter. Updated on 16 November 2026 with the client's confirmed details: refills 5L Ksh 50, 10L Ksh 100, 20L Ksh 200; new bottle purchase prices 20L regular Ksh 450–550, 20L hard Ksh 1,000–1,700, 10L Ksh 200–250, 5L Ksh 100; bottled water, customisation charges, dispenser cleaning and pumps/dispensers.
- Includes the business name and tagline, the delivery areas with the "free delivery depends on distance and quantity" wording, the WhatsApp number, call numbers, opening hours, email, Paybill and the address.
- Bottled water is quoted on request, as the client asked; customisation is Ksh 500–1,000 with a minimum order of 10 bales; dispenser cleaning is Ksh 500–1,000.
- Print styles are built in: A4, one page, with a Print button on screen. The page is `noindex` and is not linked from the public site.

## 14. How it works and delivery strip (13 November 2026)

- Added a "How it works" section between the refill planner and the label designer: choose your water, order on WhatsApp, delivery or pickup.
- Added a "Who we deliver to" strip under it: homes, offices, schools, restaurants, clinics and hotels, as listed on the client's delivery flyer.
- Delivery wording stays within confirmed facts: free delivery is mentioned, and visitors are asked for their area so the shop can confirm. No areas, minimum order or time claims were added.
- English and Swahili copy added; no JavaScript or backend changes.

## 15. Designer presets and round labels (13 November 2026)

- Added four presets to the label designer (Wedding, Church, Graduation, Corporate). A preset fills the event name, message and label colour with sample text in the current language; it adds no products or prices.
- Added a label style switch: the existing rectangular label, or a round label drawn inside the bottle preview. Text is measured and shrunk to fit the round shape.
- The WhatsApp handoff is unchanged and still sends the design details for Aquamor to confirm.
- English and Swahili labels added for the presets and style controls.

## 16. One-tap reorder (13 November 2026)

- The last submitted order (quantities, delivery or pickup, and area or name) is saved in the browser's `localStorage` under `aquamor-last-order`.
- A "Reorder my last order" button appears in the order summary when that record exists. It restores the quantities, opens the tab of the first item, sets delivery or pickup and fills the area.
- Items that are no longer available are skipped with a note; if none match, the note says the items are unavailable. Cleared or blocked storage simply hides the button.
- No account or backend is used; the customer still confirms and sends the order on WhatsApp.

## 17. Handover documents (13 November 2026)

The handover is deliberately split into two audiences, because the client does not use a terminal or edit files:

- `docs/how-to-request-a-change.md` is the one-page guide for the client: how to send a change request, what to include, how prices and photo permissions are confirmed, what happens next, and the domain/hosting ownership table to complete at handover.
- `docs/how-to-update-your-site.md` is the maintainer guide: editing text (including the exact-match Swahili keys and the `H` opening-hours table), editing every copy of a price, adding a gallery photo (JPG + WebP, photo class, tile, captions), previewing with Python or the Go server, a pre-publish checklist, publishing and support.
- Both are linked from the root `README.md` and `web/README.md`; the change-request process and maintenance scope are added to the agreement points in section 8.
- Update model agreed on 13 November 2026: managed updates. The client sends requests; the maintainer edits, checks and publishes. No self-service CMS is included in this phase.

## 18. Final performance, contrast and small-phone pass (13 November 2026)

Lighthouse 12.8.2, mobile, on the local preview:

| Category | Score |
|---|---|
| Performance | 78 (simulated mobile); 99 with throttling provided |
| Accessibility | 100 |
| Best practices | 100 |
| SEO | 100 |

The simulated mobile run reports an inflated LCP (5.4s) while the observed LCP is about 0.26s and an unthrottled run scores 99 with LCP 0.22s. The simulation is affected by the local preview server, which sends no compression or cache headers and uses HTTP/1.0; the real host should enable gzip/brotli, long cache lifetimes for images and fonts, and HTTP/2 or HTTP/3.

### Fixes made
- **Hero paint:** the first banner photo is now visible immediately instead of fading in from zero opacity, so the hero no longer waits for JavaScript plus a 1.6s transition before it can be the largest paint.
- **Fonts:** text uses `display=swap`. (The Google Fonts stylesheet was later replaced by self-hosted font files; see section 25.)
- **Video posters:** the three posters now use the existing WebP copies instead of JPG, and the unused `eventBottlesClip.jpg` was deleted. No unused image or video files remain.
- **Dark-mode contrast:** the dark-mode accent (`--deep`) was lightened from `#0a5a82` to `#1f7ea6`. White text on it is 4.57:1, and it is at least 3:1 against the dark background, card and foam surfaces; the old value was 1.87:1 against cards.
- **Accessible names:** the language button's accessible name now includes its visible `SW`/`EN` code. The gallery tiles no longer carry longer `aria-label`s that did not match their captions; the visible caption names the tile and the enlarged photo.
- **Lazy loading:** every gallery image is still `loading="lazy"` with width and height set; the three videos keep `preload="none"`, `muted` and `playsinline`.

### Checks
- Contrast: all body, muted, heading, button and band text pairs pass WCAG AA in light and dark mode (lowest text ratio 4.57:1; most above 6:1).
- Widths: no sideways scrolling at 320, 360, 390, 768, 1024 and 1280px, including the planner, designer presets, round label and reorder controls.
- Reduced motion: with `prefers-reduced-motion: reduce`, smooth scrolling is off and each banner shows a single photo.
- No JavaScript errors; all scripts pass `node --check`; Go build, vet and the CI workflow pass.

### Deferred (optional)
- Real-device test on an iPhone (safe areas, menu, WhatsApp hand-off) and a low-end Android phone, already listed under "Still to do before launch".

## 19. Client input status (13 November 2026)

- Logo: settled. Keep the existing header mark and icons; the original logo file is no longer required.
- Sent the client a WhatsApp message asking for the KEBS permit number, a photo of the certificate, and the free-delivery rules.
- No KEBS number, certification claim, delivery area or minimum order will be published until the client replies and the details are checked.
- Decisions, requests and the open-items table are kept in [`docs/maintenance-log.md`](maintenance-log.md).

## 20. Client reply (13 November 2026)

- **KEBS:** the client says their certification is currently with their supplier. They have started the process for their own permit and expect it in February next year. There is no number or certificate to publish, so no certification claim was added.
- **Prices:** the client replied about prices, but the reply does not clearly separate small bottled water from custom-branded bottles. The quoted range is Ksh 500–800 "depending on the type of the bottle"; the unit (per bottle or per pack) and any minimum order are not stated.
- **Not published yet:** the Ksh 500–800 range, dispenser-cleaning price, delivery areas and minimum order. Publishing an unclear price would mislead customers, so the site keeps its current "confirmed on WhatsApp" wording.
- **Still needed:** small bottled-water prices, branded-bottle unit and minimum order, dispenser-cleaning price, free-delivery areas and minimum order, the official business name and tagline, and the purification steps.
- Follow-up questions covering those items and the Ksh 500–800 sizes were sent to the client on 13 November 2026. The client replied with the confirmed details on 16 November 2026; see section 28. The earlier Ksh 500–800 range was superseded by Ksh 500–1,000 for customisation and a 10-bale minimum order.
- **Update (16 November 2026):** all the items listed in this section were answered. The "not published yet" decisions above no longer apply; see section 28.

## 21. Launch readiness, tests and responsive images (13 November 2026)

- **Server:** `ADDR`, `DB_PATH` and `WEB_DIR` are read from the environment; the default bind is `127.0.0.1:8080`. Added read/write/idle timeouts, graceful shutdown, a database health check, security headers (CSP, frame/nosniff/referrer/permissions, HSTS over HTTPS) and cache headers for static files.
- **Tests:** added Go tests for the confirmed catalog prices, order totals and validation, server-side pricing, admin-token protection and the products endpoint. CI now runs `go test -race ./...`, plus JavaScript and deployment-script syntax checks.
- **Deployment:** added [`deploy/Caddyfile`](../deploy/Caddyfile), a systemd unit, an environment example, an online SQLite backup script and a smoke test. [`docs/deployment.md`](deployment.md) covers DNS, TLS, firewall, backups, updates, rollback and security; [`docs/launch-checklist.md`](launch-checklist.md) tracks the launch.
- **Images:** generated 480w and 800w JPG/WebP variants for the gallery and banner photos. The gallery and banners use `srcset`/`sizes`, below-the-fold banner images are lazy-loaded, and the hero preload uses the same responsive source set. The lightbox still loads the full-size photo.
- **Cleanup:** removed the unused legacy dashboard (`backend/internal/web/`) and the leftover Next.js `tsconfig.json`.
- **Checks:** Go build/vet/test/race and `gofmt` pass; all JavaScript and shell scripts pass syntax checks; the local smoke test passes for the home page, health, products, protected orders and security headers.

## 22. 3D bottle, on-site map and reach-out acknowledgment (13 November 2026)

- **Label designer:** the bottle preview now uses glass, water and cap gradients, edge shading, highlights and a soft shadow. It can be turned by dragging it or with a slider, and the label compresses and moves around the curve of the bottle. The rectangular/round styles, colours, presets, text fitting and WhatsApp hand-off are unchanged. No idle animation, so it respects reduced motion.
- **Map:** added a lazy-loaded Google Maps embed in the contact section using the same location text as the existing "Get directions" link. The link itself is unchanged. The server CSP now allows Google Maps frames (`frame-src`).
- **Reach-out acknowledgment:** tapping any WhatsApp or email link now shows an on-page message: thanks for reaching out, send the prepared message, and Aquamor will get back to you. It is translated into Swahili. This is a UI acknowledgment, not an automated email or WhatsApp reply.
- **Still needed:** online payment and true automated replies require an external provider and credentials. Both are deferred to a later phase; see section 23 and section 5A.

## 23. Online payments with Flutterwave hosted checkout (13 November 2026)

- **Status (16 November 2026):** deferred to a later phase at the project owner's request. The code remains in place but stays switched off; no Flutterwave account, API key or webhook secret is needed for launch. Customers order through WhatsApp.
- **Decision:** the client chose a hosted checkout rather than building card or M-Pesa collection directly. Flutterwave is the first provider.
- **Button:** a Pay online button appears in the order summary only when `/api/v1/payments/config` reports that payment is configured. If the provider is not configured, the existing WhatsApp order path is unchanged.
- **Flow:** the dialog collects name, phone, receipt email and delivery area or pickup name; the site saves the order and redirects to Flutterwave's secure page. The no-JS WhatsApp path still works.
- **Return page:** [`web/payment.html`](../web/payment.html) confirms the payment with `/api/v1/payments/verify`, then shows paid, failed, pending or an error message, with links to order again or contact Aquamor on WhatsApp.
- **Security:** the server prices every order from the catalog, ignores client prices, checks the provider reference, currency and amount before marking an order paid, and verifies the `verif-hash` webhook signature. The webhook is the fallback when a customer closes the browser before returning.
- **Data:** orders can store `customer_email`, `payment_provider`, `payment_reference`, `mpesa_receipt_number` and `paid_at`; older databases receive these columns automatically at startup.
- **Tests:** Go tests cover checkout creation, provider errors, verification, amount/currency checks, the webhook signature, the paid-order update and the unconfigured state. The smoke test checks `/payment.html` and `/api/v1/payments/config`.
- **Configuration:** set `FLUTTERWAVE_SECRET_KEY`, `FLUTTERWAVE_WEBHOOK_SECRET` and `PUBLIC_BASE_URL`; see [`deployment.md`](deployment.md). Without all three, the button stays hidden.
- **Deferred to a later phase:** an approved Flutterwave merchant account and live credentials are not needed for this launch. When enabled later, online payment is limited to refills and the two 20L new-bottle catalog items because bottled water is quoted by quantity and other bottle ranges have no single payable price.

## 24. Authenticated admin orders view (13 November 2026)

- **Page:** [`web/admin.html`](../web/admin.html) with `web/js/admin.js` and `web/css/admin.css`. It is a private staff page, not linked from the public site, marked `noindex`, disallowed in `robots.txt`, and served with `X-Robots-Tag: noindex, nofollow` and `Cache-Control: no-store`.
- **Access:** the page asks for `ADMIN_TOKEN` and keeps it in `sessionStorage`, so it lasts for that browser tab only. It is never written to the page, the URL or `localStorage`. The token is sent as a bearer token to the existing `GET /api/v1/orders/all` endpoint; the server still returns 503 when no token is configured and 401 for a wrong token.
- **Content:** orders are listed newest first with customer name, phone and email, delivery area and address, order status, payment status, provider and reference, M-Pesa receipt, paid time, total and every line item. Product names come from the public catalog; an unknown product falls back to its ID.
- **Controls:** text search across order number, customer, contact, area and references; order-status and payment-status filters; refresh; sign out; and CSV export of the filtered rows (RFC 4180 quoting and formula-injection guard). All customer fields are inserted as text, not HTML.
- **Checks:** `node --check` passes; the local smoke test now checks that `/admin.html` returns 200 and is `noindex`; Go build, vet and tests still pass.
- **Access control note:** the page is only a shell. Order data remains protected by `ADMIN_TOKEN` on the API, so anyone who opens the page without the token sees nothing. The token must only be used over HTTPS.

## 25. Self-hosted fonts and write-endpoint rate limits (13 November 2026)

- **Fonts:** Bricolage Grotesque (weights 600–800) and Figtree (weights 400–600) are now served from `web/assets/fonts/` as Latin variable WOFF2 files. The Google Fonts preconnect, preload and stylesheet links were removed from `index.html`, `admin.html` and `payment.html`; each page preloads the two local files instead.
- **`@font-face`:** both faces are declared at the top of `web/css/styles.css` with the local variable-font weight ranges and `font-display: swap`. The A4 price card keeps its system-font stack.
- **CSP:** `style-src` no longer allows `fonts.googleapis.com` and `font-src` is now `'self'`. Google remains allowed only for the map frame. This removes the last third-party asset origin from the public pages.
- **Rate limiting:** a small in-memory fixed-window limiter (`backend/internal/ratelimit`) now guards the public write endpoints. `POST /api/v1/orders` allows 20 requests per minute per client IP and `POST /api/v1/payments/checkout` allows 10. Over-limit requests get `429 Too Many Requests`, a `Retry-After` header and a JSON error.
- **Client IP:** the limiter uses the last `X-Forwarded-For` value (appended by Caddy in production) and falls back to the connection address. Taking the last value ignores spoofed entries sent by the client. Expired keys are swept once per window so memory cannot grow without bound. Limits are per process and reset on restart.
- **Tests:** `backend/internal/ratelimit/ratelimit_test.go` covers the window, separate keys, concurrent access, the 429 response and client-IP parsing. The local smoke test still passes, the font files are served as `font/woff2` with a cache header, and no `fonts.googleapis.com` or `fonts.gstatic.com` references remain.

## 26. Request logging (13 November 2026)

- Added `backend/internal/requestlog`, an outer middleware that logs one line per request: method, path, status, response size, duration and client IP.
- Query strings and headers are never logged, and no customer fields are included. The logged path is `URL.Path` only, so the admin token (sent in a header), admin-page search terms (client-side) and payment parameters cannot appear in the log.
- The middleware records the status and byte count, defaults to `200` when a handler writes without calling `WriteHeader`, and keeps the response writer usable through `Unwrap`.
- It wraps the security and cache middleware in `cmd/server/main.go`, so rate-limited `429` responses are logged too.
- Tests cover the status and size, the `200` default, query-string exclusion and the forwarded client address. The local smoke test passes and the command was checked to contain no customer data or query string.

## 27. Cross-machine sync and verification (14 November 2026)

- Fetched the work merged in PRs #32–#34 on a second machine and fast-forwarded `main` to `0d13173`; the working tree is clean and matches `origin/main`.
- Re-ran every CI check locally: `gofmt`, `go build`, `go vet`, `go test -race`, `node --check` and `bash -n` all pass. GitHub CI is green for the same commits.
- [`deploy/smoke-test.sh`](../deploy/smoke-test.sh) passes all checks against a local server; a test order posted to `/api/v1/orders` is listed by `/api/v1/orders/all` with the server-priced total.
- Removed two orphaned, git-ignored SQLite files from the repository root and deleted the merged local and remote branches, leaving only `main`.
- No client-facing change. Payments stay hidden until the three `FLUTTERWAVE_*`/`PUBLIC_BASE_URL` values are set; online payments are deferred for this launch.

## 28. Client confirmed details and content update (16 November 2026)

The client sent the confirmed business details in writing. The site, price card and Swahili copy were updated.

### Confirmed details now published
- **Business name:** AQUAMOR PURIFIED DRINKING WATER.
- **Tagline:** "Hydrated? Worry no more! Aquamor has got you covered. Your satisfaction is our priority." Added to the hero and About section, and as `slogan` in the structured data.
- **Bottled water:** prices vary with quantity; the site asks visitors to contact Aquamor for a quotation. Sizes: 200ml, 300ml, 500ml, 1L, 1.5L, 2L, 3L, 5L, 10L and 20L.
- **Refills:** 5L Ksh 50, 10L Ksh 100, 20L Ksh 200; the price includes water.
- **New bottles (bottle purchase; water refills separate):** 20L regular Ksh 450–550, 20L hard Ksh 1,000–1,700, 10L Ksh 200–250, 5L Ksh 100.
- **Customised bottles:** churches, weddings, graduations, anniversaries, corporate events and other occasions; customisation Ksh 500–1,000; minimum order 10 bales; quotation on request.
- **Dispenser cleaning:** Ksh 500–1,000 depending on the dispenser and cleaning required.
- **Delivery:** Ngara, Parklands, Pangani, Gikomba, Westlands, Nairobi CBD and parts of Eastleigh. Free delivery depends on distance and quantity; customers are asked to confirm.
- **Pumps and dispensers:** manual and electric pumps for 10L and 20L bottles; manual and electric dispensers.
- **Purification process:** eight steps published in a new section, without individual sediment-filter micron sizes.
- **Domain:** preferred `aquamorwater.co.ke`; to be registered under AQUAMOR PURIFIED DRINKING WATER, with Aquamor paying renewal and keeping ownership and control.
- **Contact:** existing email `aquamor2020@gmail.com`; Gmail only, no domain email. Facebook page named AQUAMOR NGARA; the TikTok link is kept.

### Changes made
- `web/index.html`: meta and structured data updated; tagline added; bottle sizes, order notes, delivery areas, customisation charges, cleaning price, pumps, purification section, FAQ and Facebook label updated; canonical and Open Graph URLs moved to `aquamorwater.co.ke`.
- `web/js/order.js`: the bottle tab lists all ten sizes with "Price on quotation"; the new-bottle tab lists the confirmed ranges; the refill note says the 20L price includes water.
- `web/js/i18n.js`: Swahili copy added or updated for every changed string.
- `web/price-card.html`: rebuilt with the confirmed prices, the bottle-purchase/refill distinction, delivery areas and new sections.
- `backend/internal/catalog/catalog.go`: the 20L hard-bottle product now matches the bottle-purchase wording; its Ksh 1,400 price remains within the confirmed range.
- `web/robots.txt`, `web/sitemap.xml`, `deploy/Caddyfile`, `deploy/aquamor.env.example`: domain set to `aquamorwater.co.ke` with `./scripts/set-domain.sh`.
- Docs updated: this file, `maintenance-log.md`, `launch-checklist.md`, `how-to-update-your-site.md`, root `README.md` and `web/README.md`.

### Deliberate choices
- Bottled-water unit prices are not shown, because the client asked for quotations instead of fixed prices.
- New-bottle prices are shown as ranges. The order builder's estimated total uses the lower bound and is labelled "from listed prices"; online payment remains limited to refills and the two 20L bottle catalog items.
- No KEBS number, certification or health claim was added.
- The purification steps are listed exactly as supplied; individual micron sizes are omitted.

### Launch status (16 November 2026)
- **Online payments:** deferred to a later phase. No Flutterwave merchant account, API keys or webhook secret are needed for launch; the code stays hidden until configured.
- **Permissions:** the project owner cleared the existing gallery labels and the Paybill display for this build.
- **Swahili:** internal review completed and awkward phrases refined; a native-speaker read-through is still recommended.
- **Still external:** register `aquamorwater.co.ke` under AQUAMOR PURIFIED DRINKING WATER, then complete DNS/TLS/hosting handover.
- **Optional:** testimonials, original cleaning stills, shop-front/graduation/corporate photos, and a physical-device test.
