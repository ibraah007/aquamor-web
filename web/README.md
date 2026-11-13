# Aquamor website (front end)

The public one-page site for **Aquamor Purified Drinking Water**, Musindi Road, Ngara, Nairobi. Plain HTML, CSS and JavaScript. There is no build step and no framework, so it can be opened from a folder or hosted on any static host.

The repository-level README explains the Go API and how the pieces fit together. This file covers the front end. For the client's change-request sheet see [`../docs/how-to-request-a-change.md`](../docs/how-to-request-a-change.md); for the maintainer guide (editing text, prices and photos, previewing and publishing) see [`../docs/how-to-update-your-site.md`](../docs/how-to-update-your-site.md).

## Features

| Feature | What it does |
|---|---|
| Order builder | Customer picks refills, bottled water and new 20L bottles, sees a running total, and sends the order to WhatsApp, already written out; the last order can be restored with one tap |
| Online payment | When the server has Flutterwave configured, a Pay online button collects the customer's name, phone, receipt email and area, saves the order and opens the hosted checkout; a return page confirms the result |
| Refill planner | Customer enters how many people are at home and picks 5L, 10L or 20L; the site estimates weekly refills and cost, with a WhatsApp button |
| Price card | A standalone A4 page ([price-card.html](price-card.html)) for the shop counter, with the confirmed refill and bottle prices, free-delivery line and WhatsApp number; use the Print button or the browser's print dialog |
| How it works | Three steps (choose your water, order on WhatsApp, delivery or pickup) and a "Who we deliver to" strip for homes, offices, schools, restaurants, clinics and hotels |
| Label designer | Customer picks a Wedding, Church, Graduation or Corporate preset, types an event name and message, picks a colour, a rectangular or round label and a bottle size, sees it on a bottle, and sends the request to WhatsApp |
| Dispenser cleaning | Service section with before / cleaning / after photos and a WhatsApp booking button |
| Gallery | 26 real photos and three short video clips, with filters (Bottles, Events, Refills, Team) and a tap-to-enlarge view |
| Fading backgrounds | Photo slideshows behind the hero, order, branded-bottles and contact sections |
| Mobile menu | A keyboard-accessible menu below 1000px |
| English / Swahili | One-tap language switch, remembered between visits; WhatsApp messages are written in the chosen language |
| Open now | Live "Open now" / "Closed now" line using Nairobi time, refreshed every minute |
| Paybill | Copy buttons for the Paybill and account numbers |
| FAQ | Tap-to-open questions |
| Testimonials | Built, hidden until real quotes are added |
| Search data | Title, description, social preview image and structured business data |

## Project structure

```
web/
├── index.html          # All page content (sections, text, markup)
├── price-card.html     # Standalone printable A4 price card (inline styles)
├── payment.html        # Hosted-checkout return page (noindex)
├── css/
│   ├── styles.css      # Layout, colours, light/dark mode, components, mobile nav
│   └── photos.css      # One class per photo (.p-refill, .p-team, ...)
├── js/
│   ├── order.js        # Order builder + label designer
│   ├── planner.js      # Refill planner (weekly refills and cost)
│   ├── gallery.js      # Fading backgrounds, gallery filters, lightbox
│   ├── site.js         # Mobile navigation toggle
│   ├── payment.js      # Confirms the payment on payment.html
│   └── i18n.js         # Swahili switch, open/closed status, copy buttons, testimonials
├── assets/img/         # Photos (JPG + WebP, plus 480w/800w variants), og.jpg, favicons
├── assets/video/       # Short muted MP4 clips (event setup, event bottles, dispenser cleaning)
└── README.md
```

Scripts load with `defer`, in this order: `order.js`, `planner.js`, `gallery.js`, `site.js`, `i18n.js`. Keep that order, because `i18n.js` translates the rows that `order.js` creates and dispatches the language event that redraws the planner.

## How the logic works

**Order builder (`js/order.js`).** `items` holds the products shown in each tab. `PRICE` maps a product name to its price in Ksh. Each press of + or − updates the `qty` object, and the summary list and total are redrawn from it. On "Send order", the script builds a text message and opens `https://wa.me/254712140006?text=...` with the message URL-encoded. Delivery or pickup and the area are added. Product names and the delivery mode are translated when the site is in Swahili. Items without a listed price are shown separately and are not included in the estimated total.

**Online payment (`js/order.js`, `js/payment.js`).** The Pay online button is hidden until `GET /api/v1/payments/config` reports that the server has Flutterwave keys. Only refills and 20L bottles have confirmed prices, so bottled-water orders stay on WhatsApp. The dialog collects the customer's details, posts the order to `/api/v1/payments/checkout`, and redirects to the provider's secure page. When the provider returns to `payment.html`, `payment.js` calls `/api/v1/payments/verify` and shows paid, failed, pending or an error. The server prices the order again and checks the amount and currency; the browser cannot set the price.

**Reorder (`js/order.js`).** When an order is sent, its quantities, delivery mode and area are saved in `localStorage` under `aquamor-last-order`. On the next visit a "Reorder my last order" button appears if that record exists; it restores the quantities, opens the tab of the first item, sets delivery or pickup and fills the area. Items that no longer exist are skipped, and a note says so. If nothing matches, the note says the items are unavailable. The WhatsApp handoff is still the final step.

**Refill planner (`js/planner.js`).** `SIZES` holds the three refill sizes and their confirmed prices, and `PER_PERSON_DAY` is the planning assumption (2 litres). Weekly litres are the number of people times the assumption times seven; the refills are rounded up to full containers, and the cost is refills times the size price. The results and the WhatsApp message are redrawn on each change and when the language changes.

**Label designer (`js/order.js`).** The bottle is an SVG. The event text, message and label colour are SVG elements that the inputs update live. Presets fill the event name, message and colour with sample values in the current language; they do not add products or prices. The label style switches between a rectangle (`#lbl`) and a circle (`#lblRound`). `fit()` measures the rendered text and reduces the font size until it fits inside the label, using smaller limits for the round style; it runs again when the web font finishes loading, when the style changes and when the language changes. "Request this design" builds a WhatsApp message from the same values.

**Fading backgrounds (`js/gallery.js`).** Any element with class `bg` holds `<picture>` children with a `<source>` for WebP and a JPEG fallback. The script shows one at a time and fades to the next every 6 seconds. The first hero photo is eager and high priority; the other slides and below-the-fold banners are lazy-loaded. With "reduced motion" turned on in the visitor's phone, it shows only the first photo.

**Gallery filters and lightbox (`js/gallery.js`).** Each tile has `data-c` with its categories (for example `"events bottles"`). A filter button hides tiles that do not include its category. Tapping a tile opens a `<dialog>` with the photo enlarged, using the tile's caption. Gallery photos are `<picture>` elements with a WebP source and a JPEG fallback, responsive `srcset` variants (480w, 800w and the original), and `loading="lazy"`. The lightbox uses the full-size photo from `css/photos.css`, so enlarging does not show a small variant. The three short video clips sit below the grid and use `preload="none"` with a poster image, so they download only when played.

**Mobile menu (`js/site.js`).** The toggle updates `aria-expanded`, opens the menu, closes on Escape, on an outside click, or when a link is chosen.

**Open now (`js/i18n.js`).** The script reads the current day and time in the `Africa/Nairobi` time zone and compares them with the `H` table (minutes since midnight for each weekday). It writes "Open now until 8:30pm" or "Closed now, opens ..." into every element with class `js-open`, and refreshes it every minute.

**English / Swahili (`js/i18n.js`).** The `SW` dictionary maps each English sentence to its Swahili version. When Swahili is switched on, the script walks every text node on the page, looks up its exact English text, and swaps it. The original is stored so it can be switched back. It also translates `aria-label`s and placeholders. A `MutationObserver` translates text added later, such as the order rows. The chosen language is saved in `localStorage`.

**Photos (`css/photos.css`).** Each photo is a class such as `.p-refill`, using `image-set()` to prefer WebP and fall back to JPEG. These classes are used by the lightbox and the service photos; the gallery and banners use `<picture>` with explicit `srcset` so the browser can pick a smaller variant on phones. The 480w and 800w variants sit next to the original in `assets/img/`.

## Common edits

| To change | Where |
|---|---|
| Refill or bottle prices | `PRICE` in `js/order.js`, `SIZES` in `js/planner.js`, the order note and FAQ text in `index.html`, the Swahili text in `SW`, and the repeat prices in [price-card.html](price-card.html); keep `backend/internal/catalog/catalog.go` in sync |
| Phone, WhatsApp, email | `index.html` (contact section and the `wa.me` links), `js/order.js` and `js/i18n.js` (`254712140006`), and the structured data in `<head>` |
| Opening hours | `H` table in `js/i18n.js`, the contact section and FAQ in `index.html`, and the structured data in `<head>` |
| Paybill and account | Contact section and FAQ in `index.html` |
| Add a testimonial | `TESTIMONIALS` at the top of `js/i18n.js`: `{quote:"...", name:"...", place:"..."}`. The section appears when the list is not empty. |
| Add a photo | Put the JPG and a WebP copy in `assets/img/`, add `.p-name{...}` to `css/photos.css` (for the lightbox and service photos), then add a gallery tile in `index.html` with `data-c`, `data-k="name"` and a `<picture>` that includes the 480w/800w `srcset` entries. Create the variants with `convert name.jpg -resize '480x>' -strip -quality 82 name-480.jpg` and `-quality 80 name-480.webp`, then repeat for 800. |
| Add a social link | Add a button next to the Facebook and TikTok ones in the contact section of `index.html`, and add the URL to `sameAs` in `<head>` |
| Edit Swahili | `SW` in `js/i18n.js` (the English text must match `index.html` exactly) |
| Colours | The `:root` variables at the top of `css/styles.css` |
| Social preview / canonical | The `og:` tags, `twitter:` tags and `<link rel="canonical">` in `<head>`; update the `aquamor.co.ke` placeholder when the domain is confirmed |

## Run it locally

```sh
cd web
python3 -m http.server 8000
```

Then visit <http://localhost:8000>. The Pay online button stays hidden because a static server has no payment API. Or run the Go server from the repository root and open <http://localhost:8080>; set the payment variables from [`../deploy/aquamor.env.example`](../deploy/aquamor.env.example) to enable checkout.

## Deploy

Any static host works. For Netlify, Cloudflare Pages or Vercel, set the publish/output directory to `web/` and leave the build command empty. For GitHub Pages, use an Actions workflow that uploads `web/` (Pages cannot serve a subfolder directly without one).

For a custom domain (for example `aquamor.co.ke`), add it in the host's dashboard and set the DNS records the provider asks for. Then run `./scripts/set-domain.sh aquamor.co.ke` to update the canonical and `og:` URLs, `robots.txt`, `sitemap.xml` and the Caddy example.

For the full app (Go API, SQLite orders, TLS and backups), follow [`../docs/deployment.md`](../docs/deployment.md). The Go server binds to loopback by default and is meant to sit behind Caddy.

## Content rules used in this project

- Only facts the client has stated, or that are visible in their own photos and flyers, appear on the site.
- No certification, health or "100% pure" claims have been added. A KEBS permit number and purification details are still needed.
- No names or job titles are given to people in photos.
- The client confirmed permission to show the people in the photos (staff and customers, including the newer event, shop and delivery photos) on 12 November 2026.
- Customer labels (Wave360 and church labels) should still be confirmed with the client before the site goes public.

## Open issues

See [`../docs/status.md`](../docs/status.md) for the full list of what is done and what is still needed from the client.

## Ownership

The code, text and photos are for the Aquamor website. Agree with the client who owns the repository and the photos before making it public.
