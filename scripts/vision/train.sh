#!/usr/bin/env bash
# One-shot: train YOLO-cls on assets/vision (real images preferred).
# Prefer: python3 scripts/vision/fetch_real.py   # download real images first
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
cd "$ROOT"

if [[ ! -d assets/vision/train ]]; then
  echo "[prep] assets/vision empty — run: python3 scripts/vision/fetch_real.py"
  exit 1
fi

# Broken local proxy often breaks pip / model download
unset http_proxy https_proxy HTTP_PROXY HTTPS_PROXY ALL_PROXY all_proxy || true

exec python3 scripts/vision/yolo_train.py "$@"
