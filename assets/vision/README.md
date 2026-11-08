# 症状视觉训练素材（P4）

目录约定（Ultralytics classify）：

```
assets/vision/train/<class_label>/*.jpg
assets/vision/val/<class_label>/*.jpg
```

类别：`painful_red_rash` `skin_ulcer` `itchy_rash` `petechial_rash` `blood_blister` `cyanosis` `normal_skin` `other`

## 准备真实图

```bash
# 从 GitHub 伤口公开集 / ISIC / Unsplash 拉真实图（会删掉合成占位图）
python3 scripts/vision/fetch_real.py

# 同步元数据进 vision_samples
go run ./cmd/visionprep -sync-db
```

Wikimedia 在部分网络下会 SSL 超时；`itchy_rash` / `cyanosis` 目前用 Unsplash 真实照片作补充，后续可再换临床图。

## 训练

```bash
./scripts/vision/train.sh --epochs 25
# 或
python3 scripts/vision/yolo_train.py --epochs 25
```

训完后权重默认在（gitignore，需本地训出或拷贝）：

```
runs/classify/symptom-cls/weights/best.pt
```

启动 API 前：

```bash
source scripts/vision/env.sh   # 设置 SYMPTOM_YOLO_MODEL / BIN / SCRIPT
go run cmd/server/main.go
```

未配置 `SYMPTOM_YOLO_MODEL` 时，打卡皮肤图回退多模态 LLM 分类。

管理端已嵌入后端：浏览器打开 `http://127.0.0.1:8080/` 即可（无需再单独 `bun run dev`）。更新前端后执行 `./scripts/build-admin.sh` 再重启 Go。
