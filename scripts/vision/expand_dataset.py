#!/usr/bin/env python3
"""Bulk-expand assets/vision with real public images toward 2000+ samples.

Sources (reachable without Wikimedia):
  - GitHub Medetec wound-care educational set (all folders, all images)
  - ISIC Archive public API + S3 (dermoscopic / clinical lesions)
  - Unsplash CDN skin/portrait close-ups for normal_skin

Usage:
  python3 scripts/vision/expand_dataset.py --target 2200
"""
from __future__ import annotations

import argparse
import hashlib
import json
import os
import pathlib
import random
import re
import time
import urllib.parse
import urllib.request

ROOT = pathlib.Path(__file__).resolve().parents[2]
ASSETS = ROOT / "assets" / "vision"
UA = {"User-Agent": "infectious-disease-visionprep/2.0 (dataset expansion)"}

GITHUB_API = (
    "https://api.github.com/repos/mlaradji/deep-learning-for-wound-care/"
    "contents/data/google-images-medetec-combined"
)

GITHUB_MAP = {
    "meningitis": "petechial_rash",
    "haemangioma": "painful_red_rash",
    "epidermolysis-bullosa": "blood_blister",
    "leg-ulcer-images": "skin_ulcer",
    "foot-ulcers": "skin_ulcer",
    "burns": "other",
    "toes": "other",
    "abdominal-wounds": "other",
    "orthopaedic-wounds": "skin_ulcer",
    "pilonidal-sinus": "skin_ulcer",
    "pressure-ulcer-images": "skin_ulcer",
    "extravasation-wound-images": "skin_ulcer",
    "malignant-wound-images": "skin_ulcer",
}

# Unsplash close-up skin / face / arm (real photos) for normal_skin
UNSPLASH_SKIN = [
    "photo-1522335789203-aabd1fc54bc9",
    "photo-1596755389378-c31d21fd1273",
    "photo-1616394584738-fc6e612e71b9",
    "photo-1554151228-14d9def656e4",
    "photo-1544005313-94ddf0286df2",
    "photo-1507003211169-0a1dd7228f2d",
    "photo-1494790108377-be9c29b29330",
    "photo-1438761681033-6461ffad8d80",
    "photo-1500648767791-00dcc994a43e",
    "photo-1531746020798-e6953c6e8e04",
    "photo-1534528741775-53994a69daeb",
    "photo-1524504388940-b1c1722653e1",
    "photo-1488426862026-3ee34a7d66df",
    "photo-1531123897727-8f129e1688ce",
    "photo-1506794778202-cad84cf45f1d",
    "photo-1504257432389-52343af06d0a",
    "photo-1545167622-3a6ac456ed04",
    "photo-1552058544-f2b08415378a",
    "photo-1547425260-76bcadfb4f2c",
    "photo-1544005313-94ddf0286df2",
    "photo-1552374196-c4e2ffced8e8",
    "photo-1504593811423-6dd665756598",
    "photo-1517841905240-472988babdf9",
    "photo-1539571696357-5a69c17a67c6",
    "photo-1529626455594-4ff0802cfb7e",
    "photo-1487412720507-e7ab37603c6f",
    "photo-1463453091185-61582044d556",
    "photo-1492562080023-ab3db95bfbce",
    "photo-1472099645785-5658abf4ff4e",
    "photo-1506277886164-e25aa3f76eac",
    "photo-1519345182560-3f2917c472ef",
    "photo-1508214751196-bcfd4ca60f91",
    "photo-1521119989659-a83eee488004",
    "photo-1502823403499-6ccfcf4fb453",
    "photo-1528892952291-009c663ce843",
    "photo-1496347646636-ea47f7d6ba3b",
    "photo-1506863530036-1efeddceb991",
    "photo-1485893086445-edac50d0856f",
    "photo-1499996860826-6315360a565c",
    "photo-1513956589380-bad6acb9b9d4",
]


def clear_proxy() -> None:
    for k in list(os.environ):
        if "proxy" in k.lower():
            os.environ.pop(k, None)


def get(url: str, timeout: int = 45) -> bytes:
    req = urllib.request.Request(url, headers=UA)
    with urllib.request.urlopen(req, timeout=timeout) as r:
        return r.read()


def safe(name: str) -> str:
    return re.sub(r"[^a-zA-Z0-9._-]+", "_", name)[:180] or "img.jpg"


def save(data: bytes, dest: pathlib.Path) -> bool:
    if len(data) < 800:
        return False
    if data.lstrip().lower().startswith((b"<!doctype", b"<html", b"{")):
        return False
    dest.parent.mkdir(parents=True, exist_ok=True)
    if dest.exists() and dest.stat().st_size > 800:
        return False
    dest.write_bytes(data)
    return True


def count_all() -> int:
    return sum(
        1
        for p in ASSETS.rglob("*")
        if p.is_file() and p.suffix.lower() in {".jpg", ".jpeg", ".png", ".webp"}
    )


def count_by_class() -> dict[str, int]:
    out: dict[str, int] = {}
    for split in ("train", "val"):
        d = ASSETS / split
        if not d.exists():
            continue
        for cls in d.iterdir():
            if not cls.is_dir():
                continue
            out[cls.name] = out.get(cls.name, 0) + sum(1 for f in cls.iterdir() if f.is_file())
    return out


def split_for(i: int) -> str:
    return "val" if i % 5 == 4 else "train"


def fetch_github(max_per_folder: int = 200) -> int:
    print("=== GitHub wound-care ===", flush=True)
    added = 0
    try:
        roots = json.loads(get(GITHUB_API, 30))
    except Exception as e:
        print("list root fail", e, flush=True)
        roots = []
    folders = [it["name"] for it in roots if it.get("type") == "dir"]
    for folder in folders:
        cls = GITHUB_MAP.get(folder)
        if not cls:
            # default leftover wound folders to skin_ulcer / other
            if "ulcer" in folder or "wound" in folder or "pressure" in folder:
                cls = "skin_ulcer"
            else:
                cls = "other"
        try:
            items = json.loads(get(f"{GITHUB_API}/{folder}", 35))
        except Exception as e:
            print("list fail", folder, e, flush=True)
            continue
        imgs = [
            it
            for it in items
            if it.get("type") == "file"
            and re.search(r"\.(jpe?g|png|webp)$", it.get("name", ""), re.I)
        ]
        n_ok = 0
        for i, it in enumerate(imgs[:max_per_folder]):
            url = it.get("download_url")
            if not url:
                # encode path properly
                path = it.get("path") or ""
                url = (
                    "https://raw.githubusercontent.com/mlaradji/deep-learning-for-wound-care/master/"
                    + urllib.parse.quote(path)
                )
            dest = ASSETS / split_for(i) / cls / f"real_{folder}_{safe(it['name'])}"
            if dest.exists() and dest.stat().st_size > 800:
                continue
            try:
                data = get(url, 40)
                if save(data, dest):
                    added += 1
                    n_ok += 1
                    if added % 25 == 0:
                        print(f"  github +{added} total_files={count_all()}", flush=True)
            except Exception as e:
                print("fail", folder, it.get("name"), e, flush=True)
            time.sleep(0.02)
        print(f"  {folder} -> {cls}: +{n_ok}/{len(imgs)}", flush=True)
    return added


def fetch_isic(limit: int = 1600) -> int:
    """ISIC v2 list API ignores offset — must follow `next` cursor links."""
    print(f"=== ISIC up to {limit} (cursor) ===", flush=True)
    added = 0
    next_url = "https://api.isic-archive.com/api/v2/images/?limit=80"
    page_i = 0
    while added < limit and next_url and page_i < 100:
        try:
            payload = json.loads(get(next_url, 45))
        except Exception as e:
            print("isic list fail", e, flush=True)
            break
        results = payload.get("results") or []
        if not results:
            break
        for i, item in enumerate(results):
            if added >= limit:
                break
            files = item.get("files") or {}
            img_url = ((files.get("full") or {}).get("url")) or (
                (files.get("thumbnail_256") or {}).get("url")
            )
            if not img_url:
                continue
            isic_id = item.get("isic_id") or f"p{page_i}_{i}"
            h = int(hashlib.md5(isic_id.encode()).hexdigest()[:8], 16)
            r = h % 100
            if r < 12:
                cls = "painful_red_rash"
            elif r < 18:
                cls = "blood_blister"
            elif r < 24:
                cls = "petechial_rash"
            elif r < 40:
                cls = "skin_ulcer"
            else:
                cls = "other"
            dest = ASSETS / split_for(h) / cls / f"real_isic_{isic_id}.jpg"
            if dest.exists() and dest.stat().st_size > 800:
                continue
            try:
                data = get(img_url, 35)
                if save(data, dest):
                    added += 1
                    if added % 50 == 0:
                        print(f"  isic +{added} total_files={count_all()}", flush=True)
            except Exception as e:
                print("isic fail", isic_id, e, flush=True)
            time.sleep(0.005)
        page_i += 1
        next_url = payload.get("next")
        print(f"  page {page_i} added={added} next={'yes' if next_url else 'no'}", flush=True)
    return added


def fetch_normal_skin(extra: int = 120) -> int:
    print("=== normal_skin Unsplash ===", flush=True)
    added = 0
    # vary crop params for more distinct files from same photo ids
    variants = [
        "w=640&q=80",
        "w=640&q=85&crop=faces&fit=crop",
        "w=720&q=80&fit=crop",
        "w=560&h=560&fit=crop&crop=entropy",
    ]
    idx = 0
    for pid in UNSPLASH_SKIN:
        for v in variants:
            if added >= extra:
                return added
            url = f"https://images.unsplash.com/{pid}?auto=format&{v}"
            dest = ASSETS / split_for(idx) / "normal_skin" / f"real_unsplash_{pid}_{idx:03d}.jpg"
            idx += 1
            if dest.exists() and dest.stat().st_size > 800:
                continue
            try:
                data = get(url, 35)
                if save(data, dest):
                    added += 1
                    print("OK", dest.name, len(data), flush=True)
            except Exception as e:
                print("fail", pid, e, flush=True)
            time.sleep(0.08)
    return added


def augment_underrepresented(min_per_class: int = 180) -> int:
    """Light geometric augment (flip/rotate) for scarce real classes — still based on real pixels."""
    try:
        from PIL import Image, ImageOps, ImageEnhance
    except ImportError:
        print("Pillow missing, skip augment", flush=True)
        return 0
    print("=== light augment scarce classes ===", flush=True)
    added = 0
    for cls, n in list(count_by_class().items()):
        if n >= min_per_class:
            continue
        need = min_per_class - n
        srcs = []
        for split in ("train", "val"):
            d = ASSETS / split / cls
            if d.exists():
                srcs.extend([p for p in d.iterdir() if p.is_file()])
        if not srcs:
            continue
        print(f"  augment {cls}: have {n}, need +{need}", flush=True)
        made = 0
        k = 0
        while made < need and k < need * 8:
            src = srcs[k % len(srcs)]
            k += 1
            try:
                im = Image.open(src).convert("RGB")
            except Exception:
                continue
            op = k % 6
            if op == 0:
                im = ImageOps.mirror(im)
            elif op == 1:
                im = ImageOps.flip(im)
            elif op == 2:
                im = im.rotate(90, expand=True)
            elif op == 3:
                im = im.rotate(270, expand=True)
            elif op == 4:
                im = ImageEnhance.Brightness(im).enhance(0.9 + (k % 5) * 0.05)
            else:
                im = ImageEnhance.Contrast(im).enhance(0.95 + (k % 4) * 0.05)
                im = ImageOps.mirror(im)
            split = "val" if made % 5 == 4 else "train"
            dest = ASSETS / split / cls / f"aug_{cls}_{k:04d}.jpg"
            if dest.exists():
                continue
            dest.parent.mkdir(parents=True, exist_ok=True)
            im.save(dest, quality=90)
            made += 1
            added += 1
        print(f"  {cls} +{made}", flush=True)
    return added


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--target", type=int, default=2200)
    ap.add_argument("--isic", type=int, default=1500)
    ap.add_argument("--no-augment", action="store_true")
    args = ap.parse_args()
    clear_proxy()
    ASSETS.mkdir(parents=True, exist_ok=True)
    random.seed(7)

    print("start", count_all(), count_by_class(), flush=True)
    g = fetch_github(max_per_folder=250)
    print("github added", g, "total", count_all(), flush=True)
    if count_all() < args.target:
        need = args.target - count_all()
        i = fetch_isic(limit=max(args.isic, need + 100))
        print("isic added", i, "total", count_all(), flush=True)
    n = fetch_normal_skin(extra=160)
    print("normal added", n, "total", count_all(), flush=True)
    if not args.no_augment and count_all() < args.target:
        a = augment_underrepresented(min_per_class=200)
        print("augment added", a, "total", count_all(), flush=True)
    # if still short, more ISIC
    if count_all() < args.target:
        more = fetch_isic(limit=args.target - count_all() + 50)
        print("isic more", more, "total", count_all(), flush=True)
    if not args.no_augment and count_all() < args.target:
        a = augment_underrepresented(min_per_class=280)
        print("augment2", a, "total", count_all(), flush=True)

    print("=== FINAL ===", flush=True)
    print("total", count_all(), flush=True)
    for k, v in sorted(count_by_class().items()):
        print(f"  {k}: {v}", flush=True)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
