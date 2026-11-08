#!/usr/bin/env python3
"""Train symptom skin/rash YOLO-cls with auto install + auto model download.

Simplest usage (uses assets/vision if present):

  python3 scripts/vision/yolo_train.py

Or from admin export:

  python3 scripts/vision/yolo_train.py --manifest assets/vision/manifest.json

After train, point the Go server at:

  export SYMPTOM_YOLO_MODEL=runs/classify/symptom-cls/weights/best.pt
  export SYMPTOM_YOLO_BIN=python3
  export SYMPTOM_YOLO_SCRIPT=scripts/vision/yolo_infer.py
"""
from __future__ import annotations

import argparse
import json
import os
import shutil
import subprocess
import sys
import urllib.request
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
DEFAULT_ASSETS = ROOT / "assets" / "vision"
DEFAULT_MODEL = "yolov8n-cls.pt"
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


def ensure_ultralytics() -> None:
    try:
        import ultralytics  # noqa: F401
        return
    except ImportError:
        pass
    print("ultralytics 未安装，正在 pip install …", flush=True)
    base = [sys.executable, "-m", "pip", "install", "-U", "--default-timeout=120", "ultralytics"]
    mirrors = [
        base + ["-i", "https://pypi.org/simple", "--trusted-host", "pypi.org", "--trusted-host", "files.pythonhosted.org"],
        base + ["-i", "https://pypi.tuna.tsinghua.edu.cn/simple", "--trusted-host", "pypi.tuna.tsinghua.edu.cn"],
        base + ["-i", "https://mirrors.aliyun.com/pypi/simple/", "--trusted-host", "mirrors.aliyun.com"],
        base,
    ]
    env = os.environ.copy()
    for k in ("http_proxy", "https_proxy", "HTTP_PROXY", "HTTPS_PROXY", "ALL_PROXY", "all_proxy"):
        env.pop(k, None)
    for c in mirrors:
        last = subprocess.run(c, cwd=str(ROOT), env=env)
        if last.returncode == 0:
            return
    raise SystemExit("pip install ultralytics 失败，请手动安装后重试")


def ensure_model(model_name: str, model_dir: Path) -> Path:
    """Return local path to base weights; download if missing."""
    model_dir.mkdir(parents=True, exist_ok=True)
    p = Path(model_name)
    if p.is_file():
        return p.resolve()
    candidates = [
        model_dir / Path(model_name).name,
        ROOT / Path(model_name).name,
        ROOT / "models" / "yolo" / Path(model_name).name,
    ]
    for local in candidates:
        if local.is_file() and local.stat().st_size > 1000:
            print(f"使用已有基座模型: {local}", flush=True)
            return local

    # Ultralytics downloads on YOLO(name) when missing.
    print(f"本地无 {Path(model_name).name}，将由 ultralytics 自动下载（首次较慢）…", flush=True)
    return Path(model_name)


def copy_or_fetch(src: str, dest: Path) -> None:
    dest.parent.mkdir(parents=True, exist_ok=True)
    if src.startswith("http://") or src.startswith("https://"):
        urllib.request.urlretrieve(src, dest)
        return
    sp = Path(src)
    if not sp.is_absolute():
        cand = ROOT / sp
        if cand.exists():
            sp = cand
    if not sp.exists():
        raise FileNotFoundError(src)
    if sp.resolve() != dest.resolve():
        shutil.copy2(sp, dest)


def materialize_from_manifest(manifest_path: Path, data_root: Path) -> list[str]:
    payload = json.loads(manifest_path.read_text(encoding="utf-8"))
    items = payload.get("items") or payload.get("data", {}).get("items") or []
    classes = list(payload.get("classNames") or DEFAULT_CLASSES)
    for split in ("train", "val", "test"):
        for c in classes:
            (data_root / split / c).mkdir(parents=True, exist_ok=True)
    for i, it in enumerate(items):
        label = (it.get("classLabel") or "other").strip()
        if label not in classes:
            classes.append(label)
            for split in ("train", "val", "test"):
                (data_root / split / label).mkdir(parents=True, exist_ok=True)
        split = (it.get("split") or "train").strip() or "train"
        if split not in ("train", "val", "test"):
            split = "train"
        name = it.get("fileName") or f"{it.get('sampleId', i)}.jpg"
        dest = data_root / split / label / name
        src = it.get("publicUrl") or it.get("filePath") or ""
        if not src:
            continue
        try:
            copy_or_fetch(src, dest)
        except Exception as e:  # noqa: BLE001
            print(f"skip {src}: {e}", file=sys.stderr)
    return classes


def use_assets_tree(assets: Path, data_root: Path) -> list[str]:
    """Symlink/copy assets/vision/{train,val}/<class> into data_root."""
    if data_root.resolve() == assets.resolve():
        classes = sorted(
            d.name for d in (assets / "train").iterdir() if d.is_dir()
        ) if (assets / "train").exists() else list(DEFAULT_CLASSES)
        return classes

    if data_root.exists():
        shutil.rmtree(data_root)
    data_root.mkdir(parents=True, exist_ok=True)
    classes: set[str] = set()
    for split in ("train", "val", "test"):
        src_split = assets / split
        if not src_split.exists():
            continue
        for cls_dir in src_split.iterdir():
            if not cls_dir.is_dir():
                continue
            classes.add(cls_dir.name)
            dest = data_root / split / cls_dir.name
            if dest.exists():
                shutil.rmtree(dest)
            shutil.copytree(cls_dir, dest)
    return sorted(classes) or list(DEFAULT_CLASSES)


def count_images(data_root: Path) -> int:
    n = 0
    for p in data_root.rglob("*"):
        if p.suffix.lower() in {".jpg", ".jpeg", ".png", ".webp", ".bmp"}:
            n += 1
    return n


def main() -> int:
    ap = argparse.ArgumentParser(description="Train symptom YOLO classifier")
    ap.add_argument("--manifest", default="", help="optional admin export JSON")
    ap.add_argument("--assets", default=str(DEFAULT_ASSETS), help="local assets/vision tree")
    ap.add_argument("--data-root", default=str(ROOT / "data" / "symptom-cls"))
    ap.add_argument("--model", default=DEFAULT_MODEL, help="base weights; auto-download if missing")
    ap.add_argument("--model-dir", default=str(ROOT / "models" / "yolo"), help="cache dir for .pt")
    ap.add_argument("--epochs", type=int, default=15)
    ap.add_argument("--imgsz", type=int, default=224)
    ap.add_argument("--batch", type=int, default=8)
    ap.add_argument("--name", default="symptom-cls")
    ap.add_argument("--prepare-only", action="store_true")
    args = ap.parse_args()

    os.chdir(ROOT)
    data_root = Path(args.data_root)
    assets = Path(args.assets)

    if args.manifest:
        classes = materialize_from_manifest(Path(args.manifest), data_root)
    elif assets.exists() and (assets / "train").exists():
        classes = use_assets_tree(assets, data_root)
    else:
        print("找不到训练数据：请先 go run ./cmd/visionprep -synthetic 或上传样本", file=sys.stderr)
        return 1

    n = count_images(data_root)
    print(json.dumps({"dataRoot": str(data_root), "classes": classes, "images": n}, ensure_ascii=False))
    if n < 8:
        print("样本过少，训练效果会很差；先补充 assets/vision 再训", file=sys.stderr)
    if args.prepare_only:
        return 0

    ensure_ultralytics()
    from ultralytics import YOLO

    model_path = ensure_model(args.model, Path(args.model_dir))
    # Force cache dir so downloaded weights land in repo models/yolo when possible.
    os.environ.setdefault("YOLO_CONFIG_DIR", str(Path(args.model_dir) / ".ultralytics"))

    model = YOLO(str(model_path))
    project = os.environ.get("YOLO_PROJECT", str(ROOT / "runs" / "classify"))
    results = model.train(
        data=str(data_root),
        epochs=args.epochs,
        imgsz=args.imgsz,
        batch=args.batch,
        name=args.name,
        project=project,
        exist_ok=True,
        patience=8,
        verbose=True,
    )
    best = Path(project) / args.name / "weights" / "best.pt"
    # Ultralytics may nest runs; also check save_dir from results.
    save_dir = Path(getattr(results, "save_dir", best.parent.parent))
    cand = save_dir / "weights" / "best.pt"
    if cand.is_file():
        best = cand
    print("")
    print("训练完成。")
    print(f"  best weights: {best}")
    print("服务端请设置：")
    print(f"  export SYMPTOM_YOLO_MODEL={best}")
    print("  export SYMPTOM_YOLO_BIN=python3")
    print("  export SYMPTOM_YOLO_SCRIPT=scripts/vision/yolo_infer.py")
    return 0


if __name__ == "__main__":
    sys.exit(main())
