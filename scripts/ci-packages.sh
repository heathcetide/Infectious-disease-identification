#!/usr/bin/env bash
# Print module packages reachable from cmd/ + internal/ (excludes unused pkg/* toolkit).
set -euo pipefail
go list -deps -f '{{if and (not .Standard) (not .Incomplete)}}{{.ImportPath}}{{end}}' \
  ./cmd/... ./internal/... \
  | grep '^github.com/heathcetide/infectious-disease/' \
  | sort -u
