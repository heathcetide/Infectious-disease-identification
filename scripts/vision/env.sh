#!/usr/bin/env bash
# YOLO classify weights + infer helpers for the Go API.
# Usage: source scripts/vision/env.sh && go run cmd/server/main.go
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
export SYMPTOM_YOLO_MODEL="${SYMPTOM_YOLO_MODEL:-$ROOT/runs/classify/symptom-cls/weights/best.pt}"
export SYMPTOM_YOLO_BIN="${SYMPTOM_YOLO_BIN:-python3}"
export SYMPTOM_YOLO_SCRIPT="${SYMPTOM_YOLO_SCRIPT:-$ROOT/scripts/vision/yolo_infer.py}"
echo "SYMPTOM_YOLO_MODEL=$SYMPTOM_YOLO_MODEL"
