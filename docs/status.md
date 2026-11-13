# Aquamor Purified Drinking Water: Website Project Status

**Updated:** 13 November 2026
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
| Established | 2019 |
| Address | Along Musindi Road, opposite Ngara Girls Secondary School, next to Ngara Heights building, Nairobi |
| WhatsApp | 0712 140 006 |
| Phone calls | 0113 952 100, 0736 054 182 |
| Email | Aquamor2020@gmail.com |
| Opening hours | Mon to Sat 8:30am to 8:30pm; Sun 12:00pm to 7:00pm |
| Refill prices | 5L Ksh 50, 10L Ksh 100, 20L Ksh 200 |
| New 20L bottles | Hard bottle + water Ksh 1,400; regular bottle Ksh 450 |
| Bottled water sizes | 300ml, 500ml, 1L, 1.5L |
| Services | Refills, bottled water, custom-branded bottles (churches, weddings, graduations, corporate, events), dispenser cleaning and sanitizing |
| Delivery | "Free delivery" (on the shop poster and two flyers; rules not yet known) |
| Delivers to (flyer) | Homes, offices, schools, restaurants, clinics, hotels |
| Payment | Equity Paybill 247247, account 395210 (account name: Aquamor Water) |
| Facebook | https://www.facebook.com/profile.php?id=100066493751902 (confirmed by the user as the client's page) |
| TikTok | https://www.tiktok.com/@aquamor2020 (provided by the client on 12 November 2026) |

---

## 3. What has been built

### Structure and design
- [x] One-page site, mobile-first, tested at 320, 360, 390, 768, 900, 1024, 1280 and 1600 px with no sideways scrolling
- [x] Light and dark mode support
- [x] Blue and white brand look; header with logo icon (cropped from the client's flyer)
- [x] Fading photo backgrounds in the hero, the order banner, the branded-bottles banner and the contact section (respects "reduced motion")
- [x] Sticky "WhatsApp us" button on every screen
- [x] Mobile navigation menu with keyboard and Escape-key support
- [x] Skip-to-content link and visible focus styles

### Sections and features
- [x] **Hero:** headline, buttons, key facts, live "Open now / Closed now" line (Nairobi time)
- [x] **Build your order:** refill sizes, bottled water and new 20L bottles with steppers, delivery/pickup choice, running total, one-tap WhatsApp message with the order already written, and a one-tap reorder of the last order
- [x] **Refill planner:** enter the number of people at home and choose 5L, 10L or 20L; the site estimates weekly litres, refills and cost, and links to WhatsApp with the plan already written
- [x] **Label designer:** pick a Wedding, Church, Graduation or Corporate preset, type an event name and message, pick a color, a rectangular or round label and a size, see it on a bottle, send the request on WhatsApp
- [x] **Printable price card:** a standalone A4 page for the shop counter with the confirmed refill and new-bottle prices, free-delivery line and WhatsApp number (linked from the docs; not part of the one-page site)
- [x] **How it works and who we deliver to:** three steps (choose your water, order on WhatsApp, delivery or pickup) and a strip for homes, offices, schools, restaurants, clinics and hotels
- [x] **Dispenser cleaning and sanitizing:** four service points, before / cleaning / after photos, book on WhatsApp
- [x] **Our work (gallery):** 26 real photos, three short video clips, filter buttons (Bottles, Events, Refills, Team), tap to enlarge
- [x] **About:** company story with a team photo
- [x] **FAQ:** eight tap-to-open questions, answered only with confirmed facts
- [x] **Contact:** address, hours, WhatsApp, calls (two numbers), email, Paybill with Copy buttons, Get directions (Google Maps), Facebook button
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
- Hard bottle price set to Ksh 1,400 (clear flyer and the latest flyer agree; an earlier blurry rate card looked like 1,500).
- The client's official TikTok, [@aquamor2020](https://www.tiktok.com/@aquamor2020), is linked in the contact section and in the structured data. A personal-looking TikTok account (@benzi254) was **not** linked.
- The client has no Instagram account, so no Instagram link was added (issue #5).
- The logo is settled: the existing header mark (`icon.jpg` / `icon.webp`) and the icon files (`favicon-32.png`, `apple-touch-icon.png`, `icon-512.png`) are approved; no original logo file is being pursued (13 November 2026).

---

## 5. Open issues

### A. Blocking: needed from the client
1. **Official business name and tagline** (four names and four taglines seen). The logo is settled: the existing header mark and icon files are the approved versions (13 November 2026).
2. **Prices** for small bottled water (300ml, 500ml, 1L, 1.5L), branded bottles (and minimum order), and dispenser cleaning.
3. **Free delivery rules:** areas and any minimum order (requested from the client on 13 November 2026).

### B. Permissions
- [x] Showing the staff and customers in the photos, including the new event, shop and delivery photos (confirmed by the client on 12 November 2026)
- Showing other customers' labels (Wave360 and the church labels)
- Whether they sell the water pumps (the pump photo is in the gallery)
- Showing the Paybill number online (it is on their own flyers)

### C. Content still missing
- Shop front on Musindi Road, taken from the street
- Graduation and corporate event photos (church-event and wedding-reception photos are now added)
- A close-up of a glass bottle with its label
- Two or three real customer comments
- Original before/after stills for the cleaning service (the current ones are small crops from a flyer; a short cleaning video clip is now on the site)
- A named person for the About page, only if the client wants one

### D. To verify
- **KEBS:** the client replied on 13 November 2026 that their certification is currently with their supplier; they have started the process for their own permit and expect it in February next year. No permit number or certificate is available to publish, so the site makes no certification claim. Purification steps have not been requested.
- **Branded bottles:** the client quoted Ksh 500–800 "depending on the type of the bottle" on 13 November 2026. Confirm whether that is per bottle or per pack, and whether a minimum order applies, before publishing it.
- **Small bottled water:** the same reply was ambiguous; confirm whether the Ksh 500–800 range covers small bottled water or custom-branded bottles.
- **Facebook page:** confirm it is a business page and not a personal profile.
- **Sizes:** the latest flyer shows 500ml, 5L, 10L, 20L only; the client's text also lists 300ml, 1L, 1.5L.
- **Swahili text:** have a native speaker read it once.
- **Regular 20L bottle (Ksh 450):** confirm whether water is included.

### E. Handover tasks (week 4)
- Client review and the agreed revision rounds
- Choose and register a domain (for example aquamor.co.ke) and set up hosting
- Move the site to the domain (this is when the Google search data matters most)
- Hand over, plus the client change-request sheet and the maintainer guide

---

## 6. Ideas that can be built next (no client input needed)
1. ~~Refill planner (people in the home, then suggested refills per week and cost)~~ — built (see section 12)
2. ~~Designer presets (Wedding, Church, Graduation, Corporate) and a round label style~~ — built (see section 15)
3. ~~"Reorder my last order" in one tap~~ — built (see section 16)
4. ~~"Who we deliver to" strip and a three-step "How it works"~~ — built (see section 14)
5. ~~Printable A4 price card for the shop counter~~ — built (see section 13)
6. ~~"How to update your site" guide~~ — built (see section 17)
7. Final speed, contrast and small-phone pass

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
- Choose the domain and replace the `aquamor.co.ke` placeholders in `<head>`, `robots.txt` and `sitemap.xml`.
- Put the Go server behind a TLS-terminating reverse proxy before any remote use and block direct access to port 8080; the admin bearer token and customer PII must not travel over plain HTTP.
- Get the client answers listed in section 5.
- Confirm the WhatsApp link preview after deployment.
- Test on a real iPhone (safe area, menu, WhatsApp hand-off) and a low-end Android phone.
- Decide whether to connect the order builder to `POST /api/v1/orders` (currently WhatsApp only) and add an authenticated admin view.
- Consider deferring the banner slideshow images that are below the fold; only the first image is needed at load.

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

- Added [web/price-card.html](../web/price-card.html), a standalone A4 page for the shop counter. It uses the Aquamor logo and only confirmed prices: refills 5L Ksh 50, 10L Ksh 100, 20L Ksh 200; new 20L bottles Ksh 1,400 (hard bottle + water) and Ksh 450 (regular bottle).
- Includes the free-delivery line with "tell us your area and we will confirm", the WhatsApp number, call numbers, opening hours, email, Paybill and the address.
- Excludes unconfirmed prices: small bottled water, custom-branded bottles and dispenser cleaning are noted as quoted on request.
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
- **Fonts:** the Google Fonts stylesheet is preloaded and loaded without blocking render, with a `<noscript>` fallback. Text still uses `display=swap`.
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
- Responsive gallery variants: the photos are about 1000–1200px wide for a roughly 210–372px tile; adding 480w/800w `srcset` variants would save about 356 KiB on the simulated run. They are lazy-loaded and below the fold, so this does not affect the first paint.
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
- **Still needed:** small bottled-water prices, branded-bottle unit and minimum order, dispenser-cleaning price, and free-delivery areas and minimum order. Follow-up questions are recorded in the maintenance log.
