# Deployment and launch runbook

This is the technical runbook for putting the site on a server. The domain name,
DNS access and final client sign-off are still needed from the client (issue
#19); everything else can be prepared before then.

The examples use:

| Item | Value |
|---|---|
| Repository and binary | `/opt/aquamor` |
| Public site files | `/opt/aquamor/web` |
| Database | `/var/lib/aquamor/aquamor.db` |
| Environment file | `/etc/aquamor/aquamor.env` |
| Service | `aquamor.service` |
| Reverse proxy | Caddy on ports 80 and 443 |

The Go server binds to `127.0.0.1:8080` by default, so only the local reverse
proxy can reach it. Do not expose port 8080 in the firewall.

## 1. Before you start

- A DNS provider account for the client's domain.
- A server running a current Ubuntu LTS release (other Linux distributions work
  with the same steps).
- Go 1.22 or newer, `gcc` (the SQLite driver uses cgo), `sqlite3`, `curl`, `git`
  and Caddy.
- A long random admin token: `openssl rand -hex 32`.

## 2. Build

```sh
sudo apt update
sudo apt install -y build-essential sqlite3 curl git

sudo git clone https://github.com/ibraah007/aquamor-web.git /opt/aquamor
cd /opt/aquamor

# The checkout is root-owned, so build into the operator's home and install
# the binary with privileges afterwards.
CGO_ENABLED=1 go build -trimpath -ldflags="-s -w" -o "$HOME/aquamor-backend" ./backend/cmd/server
sudo install -m 755 "$HOME/aquamor-backend" /opt/aquamor/aquamor-backend
```

Run the tests before building if this is the first deploy or after a code change:

```sh
go test -race ./...
```

## 3. Service account and directories

```sh
sudo useradd --system --home /opt/aquamor --shell /usr/sbin/nologin aquamor
sudo chown -R root:root /opt/aquamor
sudo chmod 755 /opt/aquamor/aquamor-backend

sudo install -d -o aquamor -g aquamor -m 750 /var/lib/aquamor
sudo install -d -m 750 /etc/aquamor
```

## 4. Environment file

```sh
sudo install -m 600 deploy/aquamor.env.example /etc/aquamor/aquamor.env
sudoedit /etc/aquamor/aquamor.env
```

Set `ADMIN_TOKEN` to the value from `openssl rand -hex 32`. Keep the other paths
as shown unless the server layout is different.

## 5. Run the service

```sh
sudo cp deploy/aquamor.service /etc/systemd/system/aquamor.service
sudo systemctl daemon-reload
sudo systemctl enable --now aquamor
sudo systemctl status aquamor --no-pager
sudo journalctl -u aquamor -n 50 --no-pager
```

The service should log the database path and the listening address. Check it
locally:

```sh
curl -fsS http://127.0.0.1:8080/api/v1/health
```

## 6. DNS and Caddy

1. Point the domain's `A` record (and `AAAA` if the server has IPv6) at the
   server's public IP address.
2. Wait for DNS to resolve: `dig +short aquamor.co.ke`.
3. Install Caddy using the instructions for your distribution, then install the
   example configuration:

```sh
sudo cp deploy/Caddyfile /etc/caddy/Caddyfile
sudoedit /etc/caddy/Caddyfile   # set the real domain
sudo caddy validate --config /etc/caddy/Caddyfile
sudo systemctl reload caddy
```

Caddy obtains and renews the certificate automatically once DNS points at the
server. It also redirects HTTP to HTTPS, compresses responses and redirects the
`www` host to the canonical host.

## 7. Firewall

```sh
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable

# Confirm the app is not public: this should show 127.0.0.1:8080, not :8080.
sudo ss -ltnp | grep 8080
```

## 8. Replace the domain placeholders

Once the domain is confirmed, update the canonical URL, social preview URLs,
`robots.txt`, `sitemap.xml` and the Caddy example:

```sh
cd /opt/aquamor
./scripts/set-domain.sh aquamor.co.ke
git diff -- web/index.html web/robots.txt web/sitemap.xml deploy/Caddyfile
```

Reload Caddy if its configuration changed. Static files do not need a rebuild.

## 9. Verify the launch

```sh
./deploy/smoke-test.sh https://aquamor.co.ke
```

The smoke test checks the home page, health and products endpoints, the
protected orders endpoint and the security headers. Also check:

- `https://aquamor.co.ke/` loads over HTTPS with no browser warnings.
- The canonical link and `og:image` use the new domain.
- The WhatsApp link preview looks correct (for example with Facebook's Sharing
  Debugger). The first preview may need a refresh after the domain change.
- Place a test order through the API and confirm it appears at
  `/api/v1/orders/all` with the admin token. Delete the test row afterwards or
  use a throwaway database.
- Test on a real iPhone (safe areas, menu, WhatsApp hand-off) and a low-end
  Android phone.
- The Swahili switch, order builder, planner, label designer and gallery
  lightbox all work on the public domain.

## 10. Backups

SQLite is a single file, but copying it while the service is writing can produce
a corrupt copy. Use SQLite's online backup command instead:

```sh
sudo install -m 755 deploy/backup.sh /usr/local/sbin/aquamor-backup
sudo mkdir -p /var/backups/aquamor
sudo /usr/local/sbin/aquamor-backup
```

Add a daily cron entry as root:

```cron
15 2 * * * DB_PATH=/var/lib/aquamor/aquamor.db BACKUP_DIR=/var/backups/aquamor /usr/local/sbin/aquamor-backup >> /var/log/aquamor-backup.log 2>&1
```

The script keeps 14 days by default (`KEEP_DAYS` changes this). Copy the backups
off the server as well, for example with `rclone` or the hosting provider's
snapshot feature. A backup that has never been restored is not a backup: restore
one into a temporary copy and open it with `sqlite3` at least once.

To restore, stop the service first:

```sh
sudo systemctl stop aquamor
sudo cp /var/lib/aquamor/aquamor.db /var/lib/aquamor/aquamor.db.before-restore
gunzip -c /var/backups/aquamor/aquamor-YYYYMMDDTHHMMSSZ.db.gz > /tmp/restore.db || {
  echo "restore failed: backup missing or corrupt; database unchanged" >&2
  exit 1
}
sqlite3 /tmp/restore.db 'PRAGMA integrity_check;' | grep -qx ok || {
  echo "restore failed: integrity check did not return ok; database unchanged" >&2
  exit 1
}
sudo install -o aquamor -g aquamor -m 600 /tmp/restore.db /var/lib/aquamor/aquamor.db
sudo rm -f /var/lib/aquamor/aquamor.db-wal /var/lib/aquamor/aquamor.db-shm
sudo systemctl start aquamor
```

## 11. Updates and rollback

```sh
cd /opt/aquamor
sudo git pull --ff-only
CGO_ENABLED=1 go build -trimpath -ldflags="-s -w" -o "$HOME/aquamor-backend.new" ./backend/cmd/server
sudo systemctl stop aquamor
sudo cp aquamor-backend aquamor-backend.prev
sudo install -m 755 "$HOME/aquamor-backend.new" aquamor-backend
sudo systemctl start aquamor
./deploy/smoke-test.sh https://aquamor.co.ke
```

Web-only changes need no rebuild or restart because the files are served from
`web/`. Images and other assets are cached for one day (`Cache-Control`), so a
replaced photo may take up to a day to appear for returning visitors unless the
filename changes.

To roll back:

```sh
sudo systemctl stop aquamor
sudo mv aquamor-backend.prev aquamor-backend
sudo systemctl start aquamor
```

There are no schema migrations yet, so a code rollback does not require a
database rollback. Restore a backup only if data was damaged.

## 12. Security notes

- Keep `/etc/aquamor/aquamor.env` at mode `600`; it contains the admin token.
- `ADMIN_TOKEN` must be set or the orders endpoint deliberately returns `503`.
  Never put the token in a URL, a screenshot or the repository.
- The API and the token must only travel over HTTPS.
- Bind the app to loopback and expose only Caddy on ports 80 and 443.
- Run one app instance per database. SQLite allows one writer at a time.
- Application logs contain errors, not customer names, phones or addresses.
  Keep it that way when adding logging.
- For remote administration, use an SSH tunnel rather than opening 8080:
  `ssh -L 8080:127.0.0.1:8080 user@server`.

## Still blocked on the client

- The domain name and DNS access.
- The official business name and tagline (used in the page title and structured
  data).
- Prices, delivery rules and KEBS details (see `docs/status.md` section 5A).
- Final content and photo permissions before the site is public.
