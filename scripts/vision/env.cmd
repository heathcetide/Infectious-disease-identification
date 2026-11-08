@echo off
REM YOLO classify env for cmd.exe. Run from repo root:
REM   scripts\vision\env.cmd
REM   go run .\cmd\server\main.go
set "ROOT=%~dp0..\.."
if not defined SYMPTOM_YOLO_MODEL set "SYMPTOM_YOLO_MODEL=%ROOT%\runs\classify\symptom-cls\weights\best.pt"
if not defined SYMPTOM_YOLO_BIN set "SYMPTOM_YOLO_BIN=python"
if not defined SYMPTOM_YOLO_SCRIPT set "SYMPTOM_YOLO_SCRIPT=%ROOT%\scripts\vision\yolo_infer.py"
echo SYMPTOM_YOLO_MODEL=%SYMPTOM_YOLO_MODEL%
echo SYMPTOM_YOLO_BIN=%SYMPTOM_YOLO_BIN%
