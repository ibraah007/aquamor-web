# Aquamor Purified Drinking Water

One-page website for **Aquamor Purified Drinking Water**, Musindi Road, Ngara, Nairobi, plus a small Go/SQLite API for products and orders.

- `web/` — the public site: plain HTML, CSS and JavaScript. No build step.
- `backend/` — Go API and SQLite persistence.
- `docs/status.md` — project status, client decisions and open issues.
- [`docs/how-to-update-your-site.md`](docs/how-to-update-your-site.md) — non-technical handover guide for editing text, prices and photos, previewing and publishing.

## Run it

### Static site only

```sh
cd web
python3 -m http.server 8000
```

Open <http://localhost:8000>. You can also open `web/index.html` directly, but a local server is better for testing.

### Full app (site + API)

Run from the repository root so the database and `web/` resolve correctly:

```sh
ADMIN_TOKEN="choose-a-long-random-token" go run ./backend/cmd/server
```

Open <http://localhost:8080>. The database is created/opened at `./aquamor.db`.

The `ADMIN_TOKEN` is only needed to read customer orders. Without it the orders endpoint deliberately refuses to expose customer data.

## Repository layout

```
aquamor-web/
├── web/
│   ├── index.html
│   ├── price-card.html      # printable A4 price card for the shop counter
│   ├── css/
│   │   ├── styles.css      # layout, colours, light/dark mode, mobile nav
│   │   └── photos.css      # one class per photo (WebP with JPEG fallback)
│   ├── js/
│   │   ├── order.js        # order builder + label designer
│   │   ├── planner.js      # weekly refill planner
│   │   ├── gallery.js      # fading banners, filters, lightbox
│   │   ├── site.js         # mobile navigation
│   │   └── i18n.js         # English/Swahili, open/closed status, copy, testimonials
│   ├── assets/img/         # JPG + WebP photos, social preview, icons
│   └── README.md           # front-end details and edit points
├── backend/
│   ├── cmd/server/main.go
│   └── internal/
│       ├── catalog/        # single source of truth for products and prices
│       ├── db/             # SQLite connection and schema
│       ├── handlers/       # products + orders HTTP handlers
│       ├── models/
│       └── web/            # legacy prototype dashboard (not served)
├── docs/
│   ├── status.md
│   └── how-to-update-your-site.md   # handover guide (text, prices, photos, publishing)
└── aquamor.db              # local development database (not for production)
```

## API

| Method | Path | Purpose |
|---|---|---|
| GET | `/api/v1/health` | Health check |
| GET | `/api/v1/products` | Available products and prices |
| POST | `/api/v1/orders` | Create an order; prices are calculated on the server |
| GET | `/api/v1/orders/all` | List orders; requires `Authorization: Bearer <ADMIN_TOKEN>` |

### Create an order

```sh
curl -X POST http://localhost:8080/api/v1/orders \
  -H 'Content-Type: application/json' \
  -d '{
    "customer_name": "Amina Wanjiku",
    "customer_phone": "0722000111",
    "delivery_estate": "Parklands",
    "delivery_address": "Oloitoktok Road",
    "items": [
      { "product_id": 1, "bottle_option": "swap", "quantity": 2 }
    ]
  }'
```

`unit_price` and `subtotal` in the request are ignored. The server reads the price from `backend/internal/catalog/catalog.go`, so a tampered client cannot lower the total.

Use `bottle_option: "swap"` or `"not_applicable"` for a refill, and `"not_applicable"` for the 20L bottles. `"new_purchase"` is rejected for refills because their new-bottle price has not been confirmed.

### Read orders

```sh
curl -H "Authorization: Bearer $ADMIN_TOKEN" http://localhost:8080/api/v1/orders/all
```

If `ADMIN_TOKEN` is unset, the endpoint returns `503` and does not expose customer names, phones or addresses. If it is set but the token is missing or wrong, it returns `401`.

**HTTPS is required for remote use.** The Go server listens with plain HTTP (`http.ListenAndServe`), which is fine for `localhost`. Before deploying it anywhere remote, put it behind a TLS-terminating reverse proxy (for example Caddy, nginx or a cloud load balancer) and block direct access to port 8080. Otherwise the `Authorization: Bearer` token and the customer names, phones and addresses returned by this endpoint travel in cleartext.

## What changed after importing the downloaded site

The original snapshot was a standalone marketing page. It is now the `web/` front end of this repository, with these fixes applied:

- **Order builder:** scoped the order tabs so the gallery filters no longer crash it (`TypeError: … reading 'map'`).
- **Mobile navigation:** added an accessible menu toggle; the nav is no longer simply hidden below 1000px.
- **Swahili messages:** product names and Delivery/Pickup are now translated in the WhatsApp text, not only in the page.
- **Estimated total:** relabelled as “Estimated total (listed items)” and an explicit note says bottled-water prices are confirmed on WhatsApp.
- **Label designer:** text is measured and shrunk to fit the label; the default preview is no longer clipped.
- **Accessibility:** `aria-live` on the cart, translated `aria-label`s, skip link, keyboard-friendly menu, no-JS fallback for ordering.
- **Social/search:** added `og:image` (1200×630), `og:url`, `og:site_name`, Twitter card, canonical URL, favicon, apple-touch icon; added `url`, `image`, `priceRange` and `areaServed` to the structured data.
- **Performance:** WebP versions of all photos with JPEG fallback; gallery photos are lazy-loaded `<picture>` elements; the first hero image is preloaded; fonts now load the 700 weight the CSS already used.
- **Phone safe areas:** the safe-area inset is applied to the sticky header/footer rather than the root element, avoiding a strip of content above the header on notched phones.
- **Status:** the open/closed line refreshes every minute, not only on page load.
- **Go server:** now serves `web/`; the API uses server-side prices, a transaction, size/quantity validation, and a protected orders list. SQLite uses WAL and a busy timeout.

## Before going live

- The canonical URL, `og:image` URL, `web/robots.txt` and `web/sitemap.xml` use `https://aquamor.co.ke/`. Update them when the domain is confirmed.
- Confirm the remaining client facts in `docs/status.md`: small bottled-water prices, branded-bottle pricing, delivery rules, logo file, photo/label permissions and social links.
- The legacy dashboard at `backend/internal/web/index.html` is **not served** and has no authentication. Do not expose it without adding a login.
- The site still sends orders to WhatsApp; the API stores orders only when a client posts to it. Connecting the two is the next product step.
- Keep `web/js/order.js` (`PRICE`), `web/js/planner.js` (`SIZES`) and `backend/internal/catalog/catalog.go` in sync until prices come from one API; the printable card in [web/price-card.html](web/price-card.html) repeats the confirmed refill and bottle prices.
- Decide whether to add privacy-friendly analytics and a short privacy note (the API stores name, phone and address if it is used).
- `tsconfig.json` is a leftover Next.js config and is not used by this project; remove it if you are not planning a Next.js front end.
