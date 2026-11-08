# YOLO classify env for Windows PowerShell.
# Usage (from repo root):
#   . .\scripts\vision\env.ps1
#   go run .\cmd\server\main.go
$Root = Resolve-Path (Join-Path $PSScriptRoot "..\..")
if (-not $env:SYMPTOM_YOLO_MODEL) {
  $env:SYMPTOM_YOLO_MODEL = Join-Path $Root "runs\classify\symptom-cls\weights\best.pt"
}
# Windows installs usually expose `python`, not `python3`
if (-not $env:SYMPTOM_YOLO_BIN) {
  $env:SYMPTOM_YOLO_BIN = "python"
}
if (-not $env:SYMPTOM_YOLO_SCRIPT) {
  $env:SYMPTOM_YOLO_SCRIPT = Join-Path $Root "scripts\vision\yolo_infer.py"
}
Write-Host "SYMPTOM_YOLO_MODEL=$env:SYMPTOM_YOLO_MODEL"
Write-Host "SYMPTOM_YOLO_BIN=$env:SYMPTOM_YOLO_BIN"
