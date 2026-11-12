# Aquamor website (front end)

The public one-page site for **Aquamor Purified Drinking Water**, Musindi Road, Ngara, Nairobi. Plain HTML, CSS and JavaScript. There is no build step and no framework, so it can be opened from a folder or hosted on any static host.

The repository-level README explains the Go API and how the pieces fit together. This file covers the front end.

## Features

| Feature | What it does |
|---|---|
| Order builder | Customer picks refills, bottled water and new 20L bottles, sees a running total, and sends the order to WhatsApp, already written out |
| Label designer | Customer types an event name and message, picks a colour and size, sees it on a bottle, and sends the request to WhatsApp |
| Dispenser cleaning | Service section with before / cleaning / after photos and a WhatsApp booking button |
| Gallery | 26 real photos and two short video clips, with filters (Bottles, Events, Refills, Team) and a tap-to-enlarge view |
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
├── css/
│   ├── styles.css      # Layout, colours, light/dark mode, components, mobile nav
│   └── photos.css      # One class per photo (.p-refill, .p-team, ...)
├── js/
│   ├── order.js        # Order builder + label designer
│   ├── gallery.js      # Fading backgrounds, gallery filters, lightbox
│   ├── site.js         # Mobile navigation toggle
│   └── i18n.js         # Swahili switch, open/closed status, copy buttons, testimonials
├── assets/img/         # Photos (JPG + WebP), og.jpg, favicons
├── assets/video/       # Short muted MP4 clips (event setup, dispenser cleaning)
└── README.md
```

Scripts load with `defer`, in this order: `order.js`, `gallery.js`, `site.js`, `i18n.js`. Keep that order, because `i18n.js` translates the rows that `order.js` creates.

## How the logic works

**Order builder (`js/order.js`).** `items` holds the products shown in each tab. `PRICE` maps a product name to its price in Ksh. Each press of + or − updates the `qty` object, and the summary list and total are redrawn from it. On "Send order", the script builds a text message and opens `https://wa.me/254712140006?text=...` with the message URL-encoded. Delivery or pickup and the area are added. Product names and the delivery mode are translated when the site is in Swahili. Items without a listed price are shown separately and are not included in the estimated total.

**Label designer (`js/order.js`).** The bottle is an SVG. The event text, message and label colour are SVG elements that the inputs update live. `fit()` measures the rendered text and reduces the font size until it fits inside the label; it runs again when the web font finishes loading and when the language changes. "Request this design" builds a WhatsApp message from the same values.

**Fading backgrounds (`js/gallery.js`).** Any element with class `bg` holds `<i>` children, each with a photo class. The script shows one at a time and fades to the next every 6 seconds. With "reduced motion" turned on in the visitor's phone, it shows only the first photo.

**Gallery filters and lightbox (`js/gallery.js`).** Each tile has `data-c` with its categories (for example `"events bottles"`). A filter button hides tiles that do not include its category. Tapping a tile opens a `<dialog>` with the photo enlarged, using the tile's caption. Gallery photos are `<picture>` elements with a WebP source and a JPEG fallback, and are lazy-loaded. The two short video clips sit below the grid and use `preload="none"` with a poster image, so they download only when played.

**Mobile menu (`js/site.js`).** The toggle updates `aria-expanded`, opens the menu, closes on Escape, on an outside click, or when a link is chosen.

**Open now (`js/i18n.js`).** The script reads the current day and time in the `Africa/Nairobi` time zone and compares them with the `H` table (minutes since midnight for each weekday). It writes "Open now until 8:30pm" or "Closed now, opens ..." into every element with class `js-open`, and refreshes it every minute.

**English / Swahili (`js/i18n.js`).** The `SW` dictionary maps each English sentence to its Swahili version. When Swahili is switched on, the script walks every text node on the page, looks up its exact English text, and swaps it. The original is stored so it can be switched back. It also translates `aria-label`s and placeholders. A `MutationObserver` translates text added later, such as the order rows. The chosen language is saved in `localStorage`.

**Photos (`css/photos.css`).** Each photo is a class such as `.p-refill`, using `image-set()` to prefer WebP and fall back to JPEG. The same photo can be used in a banner, the gallery and the enlarged view without being stored twice.

## Common edits

| To change | Where |
|---|---|
| Refill or bottle prices | `PRICE` and `items` in `js/order.js` (also the order note and FAQ text in `index.html`, and the Swahili text in `SW`); keep `backend/internal/catalog/catalog.go` in sync |
| Phone, WhatsApp, email | `index.html` (contact section and the `wa.me` links), `js/order.js` and `js/i18n.js` (`254712140006`), and the structured data in `<head>` |
| Opening hours | `H` table in `js/i18n.js`, the contact section and FAQ in `index.html`, and the structured data in `<head>` |
| Paybill and account | Contact section and FAQ in `index.html` |
| Add a testimonial | `TESTIMONIALS` at the top of `js/i18n.js`: `{quote:"...", name:"...", place:"..."}`. The section appears when the list is not empty. |
| Add a photo | Put the JPG (and a WebP copy) in `assets/img/`, add `.p-name{...}` to `css/photos.css`, then add a tile in the gallery in `index.html` with `data-c` and `data-k="name"` |
| Add a social link | Add a button next to the Facebook and TikTok ones in the contact section of `index.html`, and add the URL to `sameAs` in `<head>` |
| Edit Swahili | `SW` in `js/i18n.js` (the English text must match `index.html` exactly) |
| Colours | The `:root` variables at the top of `css/styles.css` |
| Social preview / canonical | The `og:` tags, `twitter:` tags and `<link rel="canonical">` in `<head>`; update the `aquamor.co.ke` placeholder when the domain is confirmed |

## Run it locally

```sh
cd web
python3 -m http.server 8000
```

Then visit <http://localhost:8000>. Or run the Go server from the repository root and open <http://localhost:8080>.

## Deploy

Any static host works. For Netlify, Cloudflare Pages or Vercel, set the publish/output directory to `web/` and leave the build command empty. For GitHub Pages, use an Actions workflow that uploads `web/` (Pages cannot serve a subfolder directly without one).

For a custom domain (for example `aquamor.co.ke`), add it in the host's dashboard and set the DNS records the provider asks for. Then update the canonical and `og:` URLs.

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
