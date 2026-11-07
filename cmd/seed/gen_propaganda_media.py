#!/usr/bin/env python3
"""Generate disease-education posters (PNG) and short clips (MP4) for seed upload."""

from __future__ import annotations

import os
import subprocess
import textwrap
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

OUT = Path(__file__).resolve().parents[2] / "data" / "seed_materials"
FONT = "/Library/Fonts/Arial Unicode.ttf"
if not Path(FONT).exists():
    FONT = "/System/Library/Fonts/Supplemental/Arial Unicode.ttf"

POSTERS = [
    {
        "file": "flu_vaccine_poster.png",
        "disease": "流感",
        "title": "秋冬流感防护要点",
        "source": "依据中国疾控中心流感科普整理",
        "bullets": [
            "流感不是普通感冒，可引发肺炎等重症",
            "老人、儿童、慢性病患者优先接种流感疫苗",
            "出现发热、咳嗽及时就医并减少聚集",
            "勤洗手、戴口罩、保持室内通风",
        ],
        "colors": ((36, 99, 185), (232, 242, 255)),
    },
    {
        "file": "diarrhea_hygiene_poster.png",
        "disease": "感染性腹泻",
        "title": "细菌性腹泻：病从口入怎么防",
        "source": "依据中国疾控中心肠道传染病科普整理",
        "bullets": [
            "饭前便后用流动水洗手",
            "不喝生水，食物烧熟煮透",
            "生熟分开，避免交叉污染",
            "夏秋季注意防蝇防蟑",
        ],
        "colors": ((16, 122, 94), (230, 247, 240)),
    },
    {
        "file": "plague_prevention_poster.png",
        "disease": "鼠疫",
        "title": "预防鼠疫 你需要知道这些事",
        "source": "依据中国疾控中心鼠疫科普整理",
        "bullets": [
            "不接触野生鼠类、旱獭及病死动物",
            "草原旅游穿长袖长裤，防跳蚤叮咬",
            "不私自捕猎、剥食疫源动物",
            "返回后一周监测发热、淋巴结肿痛",
        ],
        "colors": ((140, 48, 48), (255, 240, 238)),
    },
    {
        "file": "anthrax_poster.png",
        "disease": "炭疽",
        "title": "正确认识炭疽 科学预防",
        "source": "依据中国疾控中心炭疽科普整理",
        "bullets": [
            "不接触、不私自处理可疑病死畜",
            "肉类从正规渠道购买并煮熟煮透",
            "畜牧屠宰从业者做好个人防护",
            "发现牲畜突然死亡立即报告畜牧部门",
        ],
        "colors": ((92, 64, 51), (245, 239, 232)),
    },
    {
        "file": "tb_poster.png",
        "disease": "结核病",
        "title": "肺结核核心知识",
        "source": "依据中国疾控中心结核病科普整理",
        "bullets": [
            "咳嗽咳痰超过2周应及时就诊排查",
            "主要通过呼吸道飞沫传播",
            "规范全程治疗一般需6—8个月",
            "咳嗽时掩口鼻，不随地吐痰",
        ],
        "colors": ((70, 90, 140), (236, 240, 250)),
    },
    {
        "file": "dengue_poster.png",
        "disease": "登革热（蚊媒传染病）",
        "title": "登革热防蚊指引",
        "source": "依据中国疾控中心登革热科普整理",
        "bullets": [
            "伊蚊叮咬传播，白天也要注意防蚊",
            "清除花盆、瓶罐等积水孳生地",
            "外出穿浅色长袖，喷涂有效驱蚊剂",
            "高热、头痛、皮疹及时就医并告知叮咬史",
        ],
        "colors": ((196, 112, 20), (255, 246, 230)),
    },
    {
        "file": "malaria_poster.png",
        "disease": "疟疾（蚊媒传染病）",
        "title": "严防境外疟疾输入",
        "source": "依据中国疾控中心疟疾日宣传整理",
        "bullets": [
            "赴非洲、东南亚等流行区务必防蚊",
            "回国后出现寒战高热及时就医",
            "就诊时主动告知境外旅居史",
            "我国已消除本土疟疾，输入病例仍需警惕",
        ],
        "colors": ((28, 110, 130), (230, 246, 250)),
    },
    {
        "file": "sfts_poster.png",
        "disease": "发热伴血小板减少综合征（蜱媒传染病）",
        "title": "防蜱叮咬 远离发热伴血小板减少综合征",
        "source": "依据中国疾控中心蜱媒传染病科普整理",
        "bullets": [
            "野外作业穿长袖长裤，扎紧袖口裤脚",
            "避免在草地树林长时间坐卧",
            "皮肤喷涂驱避剂，回家后检查蜱虫",
            "发热、出血等症状及时就医并告知野外史",
        ],
        "colors": ((78, 108, 58), (238, 245, 230)),
    },
    {
        "file": "covid_mask_poster.png",
        "disease": "新型冠状病毒感染",
        "title": "公众科学戴口罩指引",
        "source": "依据国家卫生健康委口罩指引整理",
        "bullets": [
            "密闭场所、人员密集处建议戴口罩",
            "口罩遮住口鼻，佩戴后少触摸外表面",
            "潮湿破损及时更换，废弃口罩规范处理",
            "勤洗手、常通风仍是基础防护措施",
        ],
        "colors": ((55, 88, 160), (236, 242, 255)),
    },
    {
        "file": "tick_enceph_poster.png",
        "disease": "森林脑炎（蜱媒传染病）",
        "title": "野外作业防蜱叮咬提示",
        "source": "依据中国疾控中心月度健康提示整理",
        "bullets": [
            "林区作业前可用氯菊酯处理衣物装备",
            "浅色衣物便于发现附着蜱虫",
            "被蜱叮咬后正确取下并观察身体状况",
            "出现发热头痛及时就医",
        ],
        "colors": ((60, 95, 70), (235, 245, 236)),
    },
    {
        "file": "typhus_poster.png",
        "disease": "斑疹伤寒",
        "title": "夏秋虫媒传染病健康提示",
        "source": "依据中国疾控中心健康提示整理",
        "bullets": [
            "注意环境卫生，减少媒介孳生",
            "野外作业做好个人防护",
            "发热伴皮疹应及时排查虫媒传染病",
            "出现不适尽早就医并说明暴露史",
        ],
        "colors": ((120, 70, 110), (248, 236, 246)),
    },
    {
        "file": "hf_poster.png",
        "disease": "流行性热出血",
        "title": "流行性出血热早期识别",
        "source": "依据中国疾控中心出血热科普整理",
        "bullets": [
            "鼠类是重要传染源，注意防鼠灭鼠",
            "清理仓库杂物时戴口罩手套",
            "发热、头痛、腰痛要提高警惕",
            "早发现、早治疗可明显改善预后",
        ],
        "colors": ((150, 55, 55), (255, 238, 238)),
    },
]


def load_font(size: int) -> ImageFont.FreeTypeFont:
    return ImageFont.truetype(FONT, size)


def wrap_lines(text: str, width: int) -> list[str]:
    return textwrap.wrap(text, width=width) or [text]


def render_poster(item: dict, path: Path) -> None:
    w, h = 1080, 1440
    primary, bg = item["colors"]
    img = Image.new("RGB", (w, h), bg)
    draw = ImageDraw.Draw(img)

    draw.rectangle([0, 0, w, 220], fill=primary)
    draw.rectangle([0, h - 90, w, h], fill=primary)

    title_font = load_font(54)
    body_font = load_font(36)
    small_font = load_font(26)
    brand_font = load_font(28)

    draw.text((48, 48), "突发传染病症状筛查预警系统 · 健康宣教", font=brand_font, fill=(255, 255, 255))
    # title may wrap
    y = 110
    for line in wrap_lines(item["title"], 14):
        draw.text((48, y), line, font=title_font, fill=(255, 255, 255))
        y += 62

    draw.text((48, 250), f"疾病类型：{item['disease']}", font=body_font, fill=primary)
    draw.text((48, 310), item["source"], font=small_font, fill=(90, 90, 90))

    y = 390
    for i, bullet in enumerate(item["bullets"], 1):
        box_top = y
        lines = wrap_lines(f"{i}. {bullet}", 18)
        box_h = 36 + len(lines) * 48
        draw.rounded_rectangle([48, box_top, w - 48, box_top + box_h], radius=18, fill=(255, 255, 255))
        ty = box_top + 24
        for line in lines:
            draw.text((80, ty), line, font=body_font, fill=(40, 40, 40))
            ty += 48
        y = box_top + box_h + 28

    draw.text((48, h - 58), "资料来源：中国疾控中心 / 国家卫生健康委公开科普", font=small_font, fill=(255, 255, 255))
    img.save(path, "PNG", optimize=True)


def make_video(poster: Path, out_mp4: Path, subtitle: str) -> None:
    # 5s silent clip from poster (letterboxed 1280x720)
    cmd = [
        "ffmpeg", "-y",
        "-loop", "1", "-i", str(poster),
        "-f", "lavfi", "-i", "anullsrc=r=44100:cl=stereo",
        "-vf", "scale=1280:720:force_original_aspect_ratio=decrease,pad=1280:720:(ow-iw)/2:(oh-ih)/2",
        "-t", "5", "-r", "25",
        "-c:v", "libx264", "-pix_fmt", "yuv420p",
        "-c:a", "aac", "-shortest", "-movflags", "+faststart",
        str(out_mp4),
    ]
    subprocess.run(cmd, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)


def make_pdf_article(item: dict, path: Path) -> None:
    # Simple multi-page-like long PNG saved as PDF via Pillow
    w, h = 1240, 1754
    primary, bg = item["colors"]
    img = Image.new("RGB", (w, h), (255, 255, 255))
    draw = ImageDraw.Draw(img)
    draw.rectangle([0, 0, w, 160], fill=primary)
    title_font = load_font(48)
    body_font = load_font(32)
    small_font = load_font(24)
    draw.text((60, 50), item["title"], font=title_font, fill=(255, 255, 255))
    y = 200
    draw.text((60, y), item["source"], font=small_font, fill=(100, 100, 100))
    y += 60
    paragraphs = [
        f"【疾病】{item['disease']}",
        "以下要点整理自中国疾控中心 / 国家卫生健康委公开发布的健康科普材料，供铁路职工健康宣教使用。",
        *item["bullets"],
        "如出现相关症状，请及时通过本系统打卡上报，并就近前往定点医疗机构就诊。",
    ]
    for p in paragraphs:
        for line in wrap_lines(p, 28):
            draw.text((60, y), line, font=body_font, fill=(35, 35, 35))
            y += 48
        y += 24
    img.save(path, "PDF", resolution=150.0)


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    manifest = []
    for item in POSTERS:
        png = OUT / item["file"]
        render_poster(item, png)
        manifest.append(("图片", item["disease"], item["title"] + "（图文海报）", png.name))

        if item["file"] in {
            "flu_vaccine_poster.png",
            "plague_prevention_poster.png",
            "dengue_poster.png",
            "malaria_poster.png",
            "sfts_poster.png",
            "covid_mask_poster.png",
        }:
            mp4 = OUT / item["file"].replace(".png", ".mp4")
            make_video(png, mp4, item["title"])
            manifest.append(("视频", item["disease"], item["title"] + "（短视频）", mp4.name))

        if item["file"] in {
            "flu_vaccine_poster.png",
            "tb_poster.png",
            "diarrhea_hygiene_poster.png",
            "anthrax_poster.png",
        }:
            pdf = OUT / item["file"].replace(".png", ".pdf")
            make_pdf_article(item, pdf)
            manifest.append(("文章", item["disease"], item["title"] + "（宣教文稿）", pdf.name))

    with open(OUT / "manifest.tsv", "w", encoding="utf-8") as f:
        for typ, disease, title, name in manifest:
            f.write(f"{typ}\t{disease}\t{title}\t{name}\n")
    print(f"generated {len(manifest)} media files in {OUT}")


if __name__ == "__main__":
    main()
