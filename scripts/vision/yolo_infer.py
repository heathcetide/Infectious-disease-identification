#!/usr/bin/env python3
"""Symptom skin/rash classifier inference for P4.

Usage:
  python scripts/vision/yolo_infer.py --model runs/symptom-cls/weights/best.pt --image /tmp/x.jpg

Stdout JSON:
  {"labels":[{"label":"skin_ulcer","confidence":0.81}]}

Env / deps: pip install ultralytics pillow
Class names must match DB catalog (see train script).
"""
from __future__ import annotations

import argparse
import json
import sys


DEFAULT_CLASSES = [
    "painful_red_rash",
    "skin_ulcer",
    "itchy_rash",
    "petechial_rash",
    "blood_blister",
    "cyanosis",
    "normal_skin",
    "other",
]


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--model", required=True, help="Ultralytics .pt classify weights")
    ap.add_argument("--image", required=True, help="Input image path")
    ap.add_argument("--topk", type=int, default=3)
    args = ap.parse_args()

    try:
        from ultralytics import YOLO
    except ImportError:
        print(json.dumps({"error": "ultralytics not installed", "labels": []}))
        return 2

    model = YOLO(args.model)
    results = model.predict(args.image, verbose=False)
    if not results:
        print(json.dumps({"labels": [{"label": "other", "confidence": 0.0}]}))
        return 0

    r0 = results[0]
    names = getattr(r0, "names", None) or getattr(model, "names", None) or {}
    if isinstance(names, dict):
        id_to_name = {int(k): str(v) for k, v in names.items()}
    else:
        id_to_name = {i: str(n) for i, n in enumerate(names)}

    labels = []
    probs = getattr(r0, "probs", None)
    if probs is not None:
        topk = min(args.topk, len(getattr(probs, "top5", []) or list(range(len(id_to_name)))))
        # ultralytics classify: probs.top5 / top5conf
        idxs = list(getattr(probs, "top5", [])[:topk])
        confs = list(getattr(probs, "top5conf", []).tolist()[:topk]) if hasattr(probs, "top5conf") else []
        if not idxs and hasattr(probs, "data"):
            import torch

            data = probs.data
            if hasattr(data, "cpu"):
                data = data.cpu()
            vals, order = data.topk(min(topk, data.numel()))
            idxs = [int(i) for i in order]
            confs = [float(v) for v in vals]
        for i, idx in enumerate(idxs):
            name = id_to_name.get(int(idx), DEFAULT_CLASSES[int(idx)] if int(idx) < len(DEFAULT_CLASSES) else "other")
            conf = float(confs[i]) if i < len(confs) else 0.0
            labels.append({"label": name, "confidence": round(conf, 4)})
    else:
        # detect fallback: map highest-conf box class
        boxes = getattr(r0, "boxes", None)
        if boxes is not None and len(boxes):
            cls = boxes.cls.tolist()
            conf = boxes.conf.tolist()
            paired = sorted(zip(cls, conf), key=lambda x: -x[1])[: args.topk]
            for c, p in paired:
                name = id_to_name.get(int(c), "other")
                labels.append({"label": name, "confidence": round(float(p), 4)})

    if not labels:
        labels = [{"label": "other", "confidence": 0.0}]
    print(json.dumps({"labels": labels}, ensure_ascii=False))
    return 0


if __name__ == "__main__":
    sys.exit(main())
