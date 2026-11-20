#!/usr/bin/env bash
# Publish web/ to the gh-pages branch so GitHub Pages serves the latest site.
# Usage: ./scripts/publish-preview.sh ["commit message"]
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

domain="aquamorwater.co.ke"
source_dir="web"
message="${1:-chore: refresh static preview from main}"

if [ "$(git rev-parse --abbrev-ref HEAD)" != "main" ]; then
  echo "error: run this from the main branch" >&2
  exit 1
fi

if [ -n "$(git status --porcelain)" ]; then
  echo "error: commit or stash your changes first (working tree is not clean)" >&2
  exit 1
fi

orig_branch="$(git rev-parse --abbrev-ref HEAD)"
tmp="$(mktemp -d)"
cleanup() {
  rm -rf "$tmp"
  git checkout "$orig_branch" >/dev/null 2>&1 || true
}
trap cleanup EXIT

# Snapshot web/ before switching branches, because gh-pages does not have it.
cp -a "$source_dir"/. "$tmp"/

git fetch origin --quiet
git checkout gh-pages
git pull --ff-only --quiet origin gh-pages

rsync -a --delete --exclude='.git' --exclude='CNAME' "$tmp"/ ./
printf '%s\n' "$domain" > CNAME

git add -A
if git diff --cached --quiet; then
  echo "No changes to publish; the preview is already up to date."
  exit 0
fi

git commit -m "$message"
git push origin gh-pages
echo "Published. GitHub Pages will rebuild in about a minute at https://$domain/"
