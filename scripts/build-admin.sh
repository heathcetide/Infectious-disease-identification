#!/usr/bin/env bash
# Build admin SPA into internal/webui/dist for go:embed.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT/admin"
if command -v bun >/dev/null 2>&1; then
  bun install --frozen-lockfile 2>/dev/null || bun install
  bun run build
elif command -v npm >/dev/null 2>&1; then
  npm ci 2>/dev/null || npm install
  npm run build
else
  echo "need bun or npm to build admin" >&2
  exit 1
fi
rm -rf "$ROOT/internal/webui/dist"
mkdir -p "$ROOT/internal/webui/dist"
cp -R "$ROOT/admin/dist/." "$ROOT/internal/webui/dist/"
echo "admin UI → internal/webui/dist ($(du -sh "$ROOT/internal/webui/dist" | awk '{print $1}'))"
