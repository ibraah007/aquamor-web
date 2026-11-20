# Go-live runbook: domain and handover

This is the ordered launch sequence for taking the site from the static preview
to `https://aquamorwater.co.ke`. [`deployment.md`](deployment.md) explains every
step in full; this file is the checklist for the day, with the owner of each step
and the check that must pass before moving on.

**Status:** the site is finished and verified (see [`status.md`](status.md)
section 29). The only external inputs left are the domain registration, a server
and the client's written approval.

Throughout this runbook, `<domain>` means the final domain the client chooses. It
is `aquamorwater.co.ke` unless they choose another; section 1 explains how to
carry a different domain through the deployment.

---

## 1. What must be in place before starting

| Input | Owner | Notes |
|---|---|---|
| Written approval of the preview | Client | One message is enough: "The site is approved as it is." |
| Final domain (`aquamorwater.co.ke` by default) registered under AQUAMOR PURIFIED DRINKING WATER | Client registers, maintainer guides | The account must be in the business name, not the maintainer's personal name. Aquamor pays the yearly renewal and keeps control. |
| Registrar login or delegated access | Client → maintainer | Use the registrar's "delegate user" feature if it has one, or a shared password in a password manager. Never a WhatsApp group. |
| Business registration certificate / KRA PIN | Client | Most `.co.ke` registrars ask for one of these to register a `.co.ke` in a business name. Confirm the exact list with the registrar. |
| A server (VPS or dedicated) | Client or maintainer (agree) | Ubuntu LTS, 1 vCPU / 1 GB RAM is enough, Go 1.22+, `gcc`, `sqlite3`, `curl`, `git`, Caddy. |
| Server IP address and SSH access | Maintainer | Keep port 8080 closed to the internet. |
| `ADMIN_TOKEN` | Maintainer | `openssl rand -hex 32`, stored in `/etc/aquamor/aquamor.env` at mode 600. |
| Swahili read-through | Client's reviewer | Recommended, not blocking. The English and Swahili pages are otherwise ready. |

For the client-facing version of this list — what to ask for on a phone call, in
plain language, with the questions he is likely to ask — use
[`client-call-script.md`](client-call-script.md).

**Decide the final domain before touching DNS.** If it stays
`aquamorwater.co.ke`, no change is needed. If the client chooses something else,
run `./scripts/set-domain.sh <domain>` in the checkout that will be deployed,
then commit and push the change. The script updates the canonical/Open Graph
URLs, `robots.txt`, `sitemap.xml`, the Caddy example and the environment example.
Deploy that commit: a separate clone, or uncommitted local edits, will not
contain the new domain. On the server, the installed
`/etc/aquamor/aquamor.env` is also separate from the checkout: once the service
is set up, set `PUBLIC_BASE_URL=https://<domain>` there and restart the service.
`git pull` updates only the example file. This value matters when online payments
are enabled later.

---

## 2. Who does what

| Phase | Client | Maintainer |
|---|---|---|
| Approve the preview | ✔ | |
| Register the domain | ✔ (maintainer sends the exact steps) | ✔ guides, checks the records |
| Set up the server and deploy | | ✔ |
| DNS records | ✔ enters them (or gives access) | ✔ sends the exact values, then verifies |
| TLS certificate | | ✔ |
| Verify the live site | ✔ looks at it on their phone | ✔ runs the smoke test and checks |
| Handover and training | ✔ receives the guides and the admin key | ✔ |
| First-week monitoring and backups | | ✔ |

---

## 3. Three days before launch (preparation)

Nothing here needs the domain.

1. **Freeze the content.** No text, price or photo changes after this point until
   the site is live, so what is reviewed is what is published.
2. **Confirm the build is green.**
   ```sh
   cd /path/to/aquamor-web
   gofmt -l . && go build ./... && go vet ./... && go test -race ./...
   for f in web/js/*.js; do node --check "$f" || exit 1; done
   for f in deploy/*.sh scripts/*.sh; do bash -n "$f" || exit 1; done
   ```
3. **Dry-run the static site.** Open the preview on a phone and a laptop:
   <https://ibraah007.github.io/aquamor-web/>.
4. **Prepare the server** (steps 2–5 of [`deployment.md`](deployment.md)): build the
   binary, create the `aquamor` user, install the environment file with the real
   `ADMIN_TOKEN`, install the systemd unit, and confirm
   `curl -fsS http://127.0.0.1:8080/api/v1/health` answers on the server itself.
5. **Write down the launch window** and tell the client when the site will switch
   over. Expect 1–2 hours of work plus up to a few hours of DNS propagation.

---

## 4. Launch day, in order

### Step 1 — Register the domain

1. The client opens an account at a KeNIC-accredited registrar, in the business
   name **AQUAMOR PURIFIED DRINKING WATER**.
2. Register `<domain>` for at least one year.
3. Turn on auto-renew and confirm the contact email is one the client reads.
4. Give the maintainer access (delegated user or shared password in a password
   manager), or have the client enter the DNS records with the values in step 2.
5. Do **not** buy hosting from the registrar unless it is the chosen server; the
   site runs on its own server behind Caddy.

**Gate:** `whois <domain>` (or the registrar dashboard) shows the domain active
and in the business name.

### Step 2 — DNS records

Set these at the registrar (or wherever the nameservers point):

| Type | Name | Value | TTL |
|---|---|---|---|
| `A` | `@` | `<server IPv4 address>` | 300 (raise to 3600 after launch) |
| `A` | `www` | `<server IPv4 address>` | 300 |
| `AAAA` | `@` and `www` | `<server IPv6 address>` | 300 |

`AAAA` only if the server actually has a public IPv6 address; remove it otherwise.
Caddy creates the certificate for both names and redirects `www` to the bare
domain.

**Gate:**
```sh
dig +short <domain>
dig +short www.<domain>
```
Both must return the server IP. With a 300-second TTL this is usually minutes.

### Step 3 — Deploy and start the service

```sh
# On the server, as the deploy user: reuse the checkout prepared in
# deployment.md section 2. Do not clone again into the non-empty directory.
cd /opt/aquamor
sudo git pull --ff-only        # picks up the set-domain.sh commit, if there is one
sudoedit /etc/aquamor/aquamor.env     # PUBLIC_BASE_URL=https://<domain>; git pull does not update this copy
CGO_ENABLED=1 go build -trimpath -ldflags="-s -w" -o "$HOME/aquamor-backend" ./backend/cmd/server
sudo install -m 755 "$HOME/aquamor-backend" /opt/aquamor/aquamor-backend

sudo cp deploy/aquamor.service /etc/systemd/system/aquamor.service
sudo systemctl daemon-reload
sudo systemctl enable --now aquamor
sudo systemctl restart aquamor        # apply the new binary and environment to a running service
sudo systemctl status aquamor --no-pager
```

**Gate:** `curl -fsS http://127.0.0.1:8080/api/v1/health` returns `"status":"ok"`,
and the chosen domain appears in `deploy/Caddyfile`, `web/robots.txt` and
`web/sitemap.xml`, which proves the deployed checkout contains it.

### Step 4 — TLS and the reverse proxy

```sh
sudo cp deploy/Caddyfile /etc/caddy/Caddyfile
sudoedit /etc/caddy/Caddyfile          # confirm the site address is <domain>
sudo caddy validate --config /etc/caddy/Caddyfile
sudo systemctl reload caddy
```

**Gate:**
```sh
curl -sI https://<domain> | head -1          # HTTP/2 200
curl -sI http://<domain> | head -1           # 308, redirects to https
curl -sI http://www.<domain> | head -1       # 308, redirects to the bare domain
```
The certificate is issued automatically once DNS resolves.

### Step 5 — Firewall

```sh
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable
sudo ss -ltnp | grep 8080        # must show 127.0.0.1:8080, never :8080
```

**Gate:** `curl -m 5 http://<server-ip>:8080/` from another machine fails.

### Step 6 — Verify the live site

```sh
cd /opt/aquamor
./deploy/smoke-test.sh https://<domain>
```

Then check by hand:

- [ ] The page loads with no certificate warning, and the address bar shows
      `https://<domain>/`.
- [ ] View source: the canonical link, `og:url` and `og:image` use the real domain.
- [ ] `https://<domain>/robots.txt` and `/sitemap.xml` show the real
      domain.
- [ ] The WhatsApp link preview looks right (Facebook Sharing Debugger, then
      refresh the cache once after the domain change).
- [ ] Place one test order from the site on a phone using a supported priced item
      (for example a 20L refill); open `/admin.html`, enter the admin token,
      confirm the order is listed with the correct total, then export the CSV.
      Bottled-water items such as 500ml are quoted on request and are not saved to
      the API (see section 6).
- [ ] `/api/v1/payments/config` returns `"enabled": false` (payments are deferred).
- [ ] The Swahili switch, order builder, planner, label designer, gallery lightbox
      and map work on the real domain.
- [ ] The A4 price card still prints on one page: open
      `https://<domain>/price-card.html` and print.
- [ ] Open the site on a real iPhone and a low-end Android: safe areas, menu and
      the WhatsApp hand-off.

### Step 7 — Search engines (optional but recommended)

1. Add the domain to Google Search Console and verify it (DNS `TXT` record is
   easiest).
2. Submit `https://<domain>/sitemap.xml`.
3. After one or two weeks, check that the preview URL
   (`ibraah007.github.io/aquamor-web`) is no longer what search results show; the
   canonical tag already points search engines to the real domain.

### Step 8 — Backups

```sh
sudo install -m 755 deploy/backup.sh /usr/local/sbin/aquamor-backup
sudo mkdir -p /var/backups/aquamor
sudo /usr/local/sbin/aquamor-backup
sudo crontab -e
```

Add the daily line from [`deployment.md`](deployment.md) section 10, then:

- [ ] Copy one backup off the server.
- [ ] Restore it into a temporary file and run `PRAGMA integrity_check;`.

**Gate:** the first backup exists and one restore has been tested.

---

## 5. Handover (same day or next)

Give the client:

- [ ] The live address and confirmation that it is on the business domain.
- [ ] [`how-to-request-a-change.md`](how-to-request-a-change.md) — how to ask for
      text, price and photo changes.
- [ ] The completed ownership table in that guide: domain owner (AQUAMOR PURIFIED
      DRINKING WATER), hosting owner, who pays the renewal (Aquamor), who
      maintains the site.
- [ ] The admin page address (`/admin.html`) and the `ADMIN_TOKEN`, sent through a
      password manager or read out on a call — never WhatsApp groups, never a
      screenshot.
- [ ] A short demonstration: how to check an order, and how to send a change
      request.
- [ ] The revision-round and response-time expectations already agreed in writing.

The maintainer keeps:

- [ ] [`how-to-update-your-site.md`](how-to-update-your-site.md) and
      [`deployment.md`](deployment.md).
- [ ] The server access and the environment file, at mode 600.
- [ ] The calendar reminder for the domain and certificate renewals (Caddy renews
      certificates automatically; the domain does not renew itself).

**Gate:** the client confirms in writing that they can open the site, see an order
and know how to request a change.

---

## 6. First week after launch

- [ ] `journalctl -u aquamor` shows one line per request and no errors; no query
      strings or customer data appear in the log.
- [ ] Orders from real customers appear in `/admin.html`, except orders that
      include a bottled-water item without a confirmed price (500ml and the other
      bottled-water sizes quoted on request). `saveOrder` skips the whole order
      when any item is missing from `PRODUCTS`, so those orders exist only in
      WhatsApp; check WhatsApp as well, because the admin page is not a complete
      order list.
- [ ] Disk space and log rotation are healthy (`journalctl --disk-usage`).
- [ ] A backup ran every night and at least one is off the server.
- [ ] Check the WhatsApp link preview again after a few days.
- [ ] Keep the `gh-pages` preview for two weeks, then delete it if it is no longer
      useful (its canonical tag already points to the real domain).

---

## 7. If something goes wrong

| Problem | What to do |
|---|---|
| DNS does not resolve | Check the record type and name (`@`, not the full domain), the TTL, and that the registrar saved it. Wait for the TTL to pass. |
| Caddy cannot get a certificate | Confirm ports 80/443 are open, DNS points at this server, and no other service is using those ports. Check `journalctl -u caddy`. |
| Wrong content is live | Web files are served from `/opt/aquamor/web`; fix the file and pull. No rebuild needed. |
| Orders stop saving | `systemctl status aquamor`, then `journalctl -u aquamor -n 100`. A missing `ADMIN_TOKEN` or a read-only database produces a clear error. |
| A bad release | Roll back the binary as in [`deployment.md`](deployment.md) section 11. Database migrations are additive, so an older binary still runs. |
| Customer data looks wrong | Stop the service and inspect the damage before restoring: a restore replaces the live database and drops every order created after the backup. Note or export those newer orders and reconcile any that can be recovered from the pre-restore copy (`aquamor.db.before-restore`, plus its `-wal` if present); that copy is kept but is not merged automatically. Use the tested restore procedure and never overwrite the backup. |

---

## 8. Launch checklist mapping

The tick-boxes for the same work live in
[`launch-checklist.md`](launch-checklist.md):

- Section A — client decisions (all confirmed; optional content still open).
- Section C — domain and hosting (this runbook, steps 1–5).
- Section D — verification on the real domain (step 6).
- Section E — data and recovery (step 8).
- Section F — handover (step 5).

Work through them in that order on the day.
