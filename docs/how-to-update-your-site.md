# How to update your website

A short guide for the Aquamor team and whoever maintains the site. You do not need to install a framework or run a build: the site is plain HTML, CSS and JavaScript, and the prices it shows are written directly in the files.

The day-to-day edits are in a few places:

| What you want to change | Where it lives |
|---|---|
| Text, headings, contact details, FAQ | `web/index.html` |
| Prices | `web/js/order.js`, `web/js/planner.js` and `web/price-card.html` (see the price table below) |
| Photos and captions | `web/assets/img/`, `web/css/photos.css` and the gallery in `web/index.html` |
| Swahili wording | `web/js/i18n.js` |

Everything is kept in the git repository. Make the change on a branch, open a pull request, wait for the checks, then merge. Keep a copy of the current site before a big edit.

---

## 1. Edit text

1. Open `web/index.html` in a text editor.
2. Use your editor's search (Ctrl+F or Cmd+F) to find the sentence you want to change.
3. Edit only the text between the tags, for example:

   ```html
   <h2>How it works</h2>
   <p class="lead">Three simple steps from your first message to water at your door.</p>
   ```

4. If the page has a Swahili version, add or update the matching entry in the `SW` dictionary in `web/js/i18n.js`. The dictionary is keyed by the **exact English text**, so if you change the English sentence, change the key too. Otherwise the Swahili version silently falls back to English.

   ```js
   "How it works": "Inafanyaje kazi",
   ```

5. Save the file and preview it (section 4).

**Tip:** the contact details appear in more than one place. If the phone, WhatsApp, email, hours or Paybill changes, search the whole `web/` folder for the old value and update every copy, including the structured data in the `<head>` of `index.html`.

---

## 2. Edit prices

Prices are repeated in several files on purpose: the browser needs them, and the server checks orders against its own list. Change them together, otherwise the site and the API will disagree.

| Price | File | What to edit |
|---|---|---|
| Refills in the order builder | `web/js/order.js` | `PRICE` and the matching text in `items` (for example `Ksh 50 · bring your container`) |
| Refills in the refill planner | `web/js/planner.js` | `SIZES` |
| Notes, FAQ and the order summary text | `web/index.html` | the order note and the FAQ answers |
| Printable shop card | `web/price-card.html` | the prices in the card |
| Swahili copies of the notes and FAQ | `web/js/i18n.js` | the `SW` entries that contain prices |
| Server-side prices | `backend/internal/catalog/catalog.go` | the catalog used by `POST /api/v1/orders` |

Only publish a price that the client has confirmed in writing. If a price is not confirmed, do not add it: the site currently leaves small bottled water, custom-branded bottles and dispenser cleaning as "quoted on request".

After a price change, check it on the page in English and Swahili, in the order builder, in the planner and on the price card.

---

## 3. Add a gallery photo

You need two copies of the same photo: a JPG (fallback) and a WebP (smaller, used when the browser supports it).

1. **Prepare the files.** Resize the long edge to about 1000–1200px and save a JPG (`myphoto.jpg`) and a WebP (`myphoto.webp`) with the same name. Keep the file size small; the gallery loads many photos.
2. **Put both files in `web/assets/img/`.**
3. **Add a photo class** at the end of `web/css/photos.css`, replacing `myphoto` with your file name:

   ```css
   .p-myphoto{background-image:url("../assets/img/myphoto.jpg");background-image:-webkit-image-set(url("../assets/img/myphoto.webp") type("image/webp"),url("../assets/img/myphoto.jpg") type("image/jpeg"));background-image:image-set(url("../assets/img/myphoto.webp") type("image/webp"),url("../assets/img/myphoto.jpg") type("image/jpeg"))}
   ```

4. **Add a tile** inside the gallery grid in `web/index.html`, next to the other `<button class="g">` tiles:

   ```html
   <button class="g" data-c="bottles" data-k="myphoto" aria-label="A short description of the photo">
     <picture><source srcset="assets/img/myphoto.webp" type="image/webp"><img src="assets/img/myphoto.jpg" alt="" width="1000" height="750" loading="lazy" decoding="async"></picture>
     <span>A short description of the photo</span>
   </button>
   ```

   - `data-k` must match the class name without `.p-` (`myphoto`).
   - `data-c` is the filter category: `bottles`, `events`, `refills` or `team`. You can use more than one, separated by a space.
   - `aria-label` and the `<span>` should say what the photo shows; they are also used by the enlarged view.
   - Use the real width and height of the file in the `img` tag.

5. **Add the Swahili caption** to `web/js/i18n.js`, keyed by the exact English caption:

   ```js
   "A short description of the photo": "Maelezo mafupi ya picha",
   ```

6. **Update the photo count** in `docs/status.md` (the "Our work (gallery)" line) and in `web/README.md` if the number is written there.

7. Preview the gallery, test the filter buttons and tap the photo to check the enlarged view.

Permission to show people in a photo must be confirmed by the client before it goes on the site.

---

## 4. Preview the site locally

### Static site (fastest)

```sh
cd web
python3 -m http.server 8000
```

Open <http://localhost:8000>. Stop the server with Ctrl+C.

### Full site with the Go API

From the repository root:

```sh
ADMIN_TOKEN="choose-a-long-random-token" go run ./backend/cmd/server
```

Open <http://localhost:8080>. The API serves the same `web/` folder.

---

## 5. What to check before publishing

- [ ] Prices match in the order builder, the planner, the notes, the price card and `backend/internal/catalog/catalog.go`
- [ ] English and Swahili both read correctly (the SW button at the top)
- [ ] The page looks right on a phone (about 360–390px wide) and on a desktop
- [ ] Light and dark mode (change your phone or browser theme)
- [ ] The WhatsApp buttons open a chat with the right message
- [ ] New photos load, and the gallery filters and enlarged view work
- [ ] `go build ./...`, `go vet ./...` and `node --check web/js/*.js` pass, or the pull-request checks are green
- [ ] The title, description and canonical URL in the `<head>` are correct for the live domain
- [ ] Nothing private was added: no admin token, no customer names or phone numbers, no Paybill photo that the client has not approved

---

## 6. Publishing, domain, hosting and support

The site goes live by merging to `main` and letting the hosting pull or deploy the repository. Ask the person who set up hosting for the exact deploy step if it is not automatic.

Fill this in at handover and keep it with the client:

| Item | Detail | Owner |
|---|---|---|
| Domain name | for example `aquamor.co.ke` | to agree in writing |
| Registrar account | | to agree in writing |
| Hosting account | | to agree in writing |
| Who pays the yearly domain and hosting renewals | | to agree in writing |
| TLS / HTTPS | site must be served over HTTPS; the Go API must sit behind a TLS-terminating reverse proxy and port 8080 must not be public | |
| Backups | git repository (and a copy of the database if orders are used) | |
| Support contact | | to agree in writing |

**Security note:** the admin token for the orders list is a password. Never put it in `web/`, in a screenshot, or in a shared chat. The orders endpoint must only be used over HTTPS.

---

## 7. If something breaks

1. Do not keep editing on top of a broken page. Note what changed.
2. Check the browser console (right-click → Inspect → Console) for an error message; it usually names the file and line.
3. Revert the change (in git, or restore your backup) and preview again.
4. If you cannot fix it, send the exact error text and the page you were on to the support contact above.
