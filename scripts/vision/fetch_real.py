#!/usr/bin/env python3
"""Download real public educational skin/wound images into assets/vision.

Sources:
  - GitHub mlaradji/deep-learning-for-wound-care (Medetec / educational web images)
  - Direct Wikimedia Commons upload.wikimedia.org thumbs (when reachable)

Does NOT generate synthetic placeholders.
"""
from __future__ import annotations

import hashlib
import json
import os
import pathlib
import re
import ssl
import time
import urllib.error
import urllib.request

ROOT = pathlib.Path(__file__).resolve().parents[2]
ASSETS = ROOT / "assets" / "vision"
UA = {"User-Agent": "infectious-disease-visionprep/1.1 (educational dataset curation)"}

# Folder under Medetec-combined -> our class label
GITHUB_MAP = {
    "meningitis": "petechial_rash",
    "haemangioma": "painful_red_rash",
    "epidermolysis-bullosa": "blood_blister",
    "leg-ulcer-images": "skin_ulcer",
    "foot-ulcers": "skin_ulcer",
    "burns": "other",
    "abrasions": "other",
    "surgical-wounds": "other",
    "extravasation-injury": "skin_ulcer",
    "malignant-wounds": "skin_ulcer",
    "pressure-ulcers-grade2": "skin_ulcer",
    "pressure-ulcers-grade3": "skin_ulcer",
    "pressure-ulcers-grade4": "skin_ulcer",
}

GITHUB_API = (
    "https://api.github.com/repos/mlaradji/deep-learning-for-wound-care/"
    "contents/data/google-images-medetec-combined"
)

# Curated Wikimedia direct thumbs for under-covered classes
WIKI = {
    "painful_red_rash": [
        "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8a/Erythema_migrans.jpg/640px-Erythema_migrans.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3a/Measles.jpg/640px-Measles.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2e/Scarlet_fever.jpg/640px-Scarlet_fever.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/9/9a/Erysipelas.jpg/640px-Erysipelas.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/6/6b/Cellulitis.jpg/640px-Cellulitis.jpg",
    ],
    "itchy_rash": [
        "https://upload.wikimedia.org/wikipedia/commons/thumb/6/6f/Atopic_dermatitis.jpg/640px-Atopic_dermatitis.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/1/1e/Atopic_dermatitis1.jpg/640px-Atopic_dermatitis1.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Urticaria.jpg/640px-Urticaria.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c8/Contact_dermatitis.jpg/640px-Contact_dermatitis.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e9/Eczema_hand.jpg/640px-Eczema_hand.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/0/0a/Psoriasis_on_back.jpg/640px-Psoriasis_on_back.jpg",
    ],
    "petechial_rash": [
        "https://upload.wikimedia.org/wikipedia/commons/thumb/7/7d/Petechia.jpg/640px-Petechia.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d8/Henoch-Sch%C3%B6nlein_purpura.jpg/640px-Henoch-Sch%C3%B6nlein_purpura.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/2/25/Purpura.jpg/640px-Purpura.jpg",
    ],
    "cyanosis": [
        "https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Cyanosis.jpg/640px-Cyanosis.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/Central_cyanosis.jpg/640px-Central_cyanosis.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3f/Clubbing.jpg/640px-Clubbing.jpg",
    ],
    "normal_skin": [
        "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4b/Human_skin.jpg/640px-Human_skin.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2b/Human_arm.jpg/640px-Human_arm.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/1/1e/Skin.jpg/640px-Skin.jpg",
    ],
    "blood_blister": [
        "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a8/Blood_blister.jpg/640px-Blood_blister.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Blister.jpg/640px-Blister.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/f/f2/Epidermolysis_bullosa.jpg/640px-Epidermolysis_bullosa.jpg",
    ],
    "skin_ulcer": [
        "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4c/Venous_ulcer_001.jpg/640px-Venous_ulcer_001.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e0/Diabetic_foot_ulcer.jpg/640px-Diabetic_foot_ulcer.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Pressure_ulcer.jpg/640px-Pressure_ulcer.jpg",
    ],
    "other": [
        "https://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Acne_vulgaris.jpg/640px-Acne_vulgaris.jpg",
        "https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Burn_wound.jpg/640px-Burn_wound.jpg",
    ],
}

# Extra raw.githubusercontent.com paths that are known-good educational images
RAW_EXTRA = [
    # (class, url) — leave empty; GitHub API path below covers bulk
]


def ctx() -> ssl.SSLContext:
    c = ssl.create_default_context()
    return c


def open_url(url: str, timeout: int = 45):
    req = urllib.request.Request(url, headers=UA)
    return urllib.request.urlopen(req, timeout=timeout, context=ctx())


def safe_name(name: str) -> str:
    name = re.sub(r"[^a-zA-Z0-9._-]+", "_", name)
    return name[:180] or "img.jpg"


def save_bytes(data: bytes, dest: pathlib.Path) -> bool:
    if len(data) < 800:
        return False
    # reject HTML error pages
    head = data[:200].lstrip().lower()
    if head.startswith(b"<!doctype") or head.startswith(b"<html"):
        return False
    dest.parent.mkdir(parents=True, exist_ok=True)
    dest.write_bytes(data)
    return True


def download(url: str, dest: pathlib.Path) -> str:
    if dest.exists() and dest.stat().st_size > 800:
        return "skip"
    try:
        with open_url(url) as r:
            data = r.read()
        if save_bytes(data, dest):
            print("OK", dest.relative_to(ROOT), len(data))
            return "ok"
        dest.unlink(missing_ok=True)
        print("BAD", url)
        return "bad"
    except Exception as e:
        print("FAIL", url, e)
        dest.unlink(missing_ok=True)
        return "fail"


def list_github_dir(folder: str):
    url = f"{GITHUB_API}/{folder}"
    with open_url(url, timeout=40) as r:
        return json.load(r)


def fetch_github(max_per_folder: int = 40) -> dict:
    stats = {"ok": 0, "skip": 0, "fail": 0, "bad": 0}
    for folder, cls in GITHUB_MAP.items():
        try:
            items = list_github_dir(folder)
        except Exception as e:
            print("list fail", folder, e)
            continue
        imgs = [
            it
            for it in items
            if it.get("type") == "file"
            and re.search(r"\.(jpe?g|png|webp)$", it.get("name", ""), re.I)
        ]
        for i, it in enumerate(imgs[:max_per_folder]):
            split = "val" if i % 5 == 4 else "train"
            name = safe_name(it["name"])
            dest = ASSETS / split / cls / f"real_{folder}_{name}"
            url = it.get("download_url")
            if not url:
                # fallback: encode path segments
                path = it.get("path", "")
                url = (
                    "https://raw.githubusercontent.com/mlaradji/deep-learning-for-wound-care/"
                    f"master/{urllib.request.quote(path)}"
                )
            st = download(url, dest)
            stats[st] = stats.get(st, 0) + 1
            time.sleep(0.05)
    return stats


def fetch_wiki() -> dict:
    stats = {"ok": 0, "skip": 0, "fail": 0, "bad": 0}
    for cls, urls in WIKI.items():
        for i, url in enumerate(urls):
            split = "val" if i % 4 == 3 else "train"
            h = hashlib.md5(url.encode()).hexdigest()[:8]
            ext = ".jpg"
            if ".png" in url.lower():
                ext = ".png"
            dest = ASSETS / split / cls / f"real_wiki_{h}{ext}"
            st = download(url, dest)
            stats[st] = stats.get(st, 0) + 1
            time.sleep(0.1)
    return stats


def purge_synthetic() -> int:
    """Remove color-block placeholders named <class>_<split>_NN.jpg."""
    removed = 0
    pat = re.compile(r"^(painful_red_rash|skin_ulcer|itchy_rash|petechial_rash|blood_blister|cyanosis|normal_skin|other)_(train|val|test)_\d+\.(jpe?g|png|webp)$", re.I)
    for p in ASSETS.rglob("*"):
        if not p.is_file():
            continue
        if pat.match(p.name) and not p.name.startswith("real_"):
            p.unlink(missing_ok=True)
            removed += 1
            print("DEL synth", p.relative_to(ROOT))
    return removed


def summarize():
    rows = []
    for split in ("train", "val", "test"):
        d = ASSETS / split
        if not d.exists():
            continue
        for cls in sorted(x.name for x in d.iterdir() if x.is_dir()):
            n = len([f for f in (d / cls).iterdir() if f.is_file()])
            rows.append((split, cls, n))
    print("--- counts ---")
    for split, cls, n in rows:
        print(f"{split:5} {cls:20} {n}")
    print("total", sum(n for _, _, n in rows))


def main():
    # Avoid corporate proxies that break SSL to GitHub/Wikimedia
    for k in ("http_proxy", "https_proxy", "HTTP_PROXY", "HTTPS_PROXY", "ALL_PROXY", "all_proxy"):
        os.environ.pop(k, None)

    ASSETS.mkdir(parents=True, exist_ok=True)
    n = purge_synthetic()
    print(f"purged {n} synthetic files")
    print("=== github wound-care ===")
    print(fetch_github(max_per_folder=50))
    print("=== wikimedia ===")
    print(fetch_wiki())
    summarize()


if __name__ == "__main__":
    main()
