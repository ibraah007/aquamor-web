#!/usr/bin/env bash
# Create a consistent online backup of the Aquamor SQLite database.
#
# Install as /usr/local/sbin/aquamor-backup and run it from cron, for example:
#   15 2 * * * root DB_PATH=/var/lib/aquamor/aquamor.db /usr/local/sbin/aquamor-backup
set -euo pipefail

DB_PATH="${DB_PATH:-/var/lib/aquamor/aquamor.db}"
BACKUP_DIR="${BACKUP_DIR:-/var/backups/aquamor}"
KEEP_DAYS="${KEEP_DAYS:-14}"

if ! command -v sqlite3 >/dev/null 2>&1; then
	echo "backup: sqlite3 is not installed" >&2
	exit 1
fi
if [[ ! -r "$DB_PATH" ]]; then
	echo "backup: cannot read $DB_PATH" >&2
	exit 1
fi

umask 077
mkdir -p "$BACKUP_DIR"

stamp="$(date -u +%Y%m%dT%H%M%SZ)"
tmp="$BACKUP_DIR/.aquamor-$stamp.db"
final="$BACKUP_DIR/aquamor-$stamp.db.gz"

# .backup uses SQLite's online backup API, so it is safe while the server is
# running and includes committed WAL contents.
sqlite3 "$DB_PATH" ".backup '$tmp'"
gzip -9 "$tmp"
mv -- "$tmp.gz" "$final"

# Delete old local backups. Keep copies off the server as well.
find "$BACKUP_DIR" -type f -name 'aquamor-*.db.gz' -mtime +"$KEEP_DAYS" -delete

echo "backup: $final"
