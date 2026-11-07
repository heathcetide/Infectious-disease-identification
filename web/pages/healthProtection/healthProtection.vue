<script>
import { getPageLearningMaterials } from '../../services/clock'

const FALLBACK_DISEASES = [
  '新型冠状病毒感染',
  '流感',
  '鼠疫',
  '感染性腹泻',
  '炭疽',
  '结核病',
  '登革热（蚊媒传染病）',
  '疟疾（蚊媒传染病）',
  '森林脑炎（蜱媒传染病）',
  '发热伴血小板减少综合征（蜱媒传染病）',
  '斑疹伤寒',
  '流行性热出血',
]

export default {
  data() {
    return {
      diseases: [],
      loading: false,
    }
  },
  onShow() {
    this.loadDiseases()
  },
  methods: {
    shortName(name) {
      return String(name || '')
        .replace('新型冠状病毒感染', '新型冠状病毒')
        .replace('（蚊媒传染病）', '')
        .replace('（蜱媒传染病）', '')
    },
    loadDiseases() {
      this.loading = true
      getPageLearningMaterials({
        diseaseTypeName: '',
        pageNo: 1,
        pageSize: 200,
      })
        .then((res) => {
          const records =
            (res && res.data && (res.data.list || res.data.records)) || []
          const names = []
          const seen = new Set()
          records.forEach((item) => {
            const n = item.diseaseTypeName
            if (n && !seen.has(n)) {
              seen.add(n)
              names.push(n)
            }
          })
          this.diseases = names.length ? names : FALLBACK_DISEASES
        })
        .catch(() => {
          this.diseases = FALLBACK_DISEASES
        })
        .finally(() => {
          this.loading = false
        })
    },
    navigateToDiseasePage(diseaseTypeName) {
      const name = String(diseaseTypeName || '').trim()
      if (!name) return
      uni.navigateTo({
        url:
          '/pages/publicityMaterial/publicityMaterial?diseaseTypeName=' +
          encodeURIComponent(name),
      })
    },
  },
}
</script>

<template>
  <view class="page">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <view class="scroll">
      <view class="hero">
        <view class="hero-text">
          <text class="hero-kicker">健康防护</text>
          <text class="hero-title">疾病知识</text>
          <text class="hero-desc">选择病种查看宣教材料与防护要点</text>
        </view>
        <image class="hero-art" src="/static/bacteria2.png" mode="aspectFit" />
      </view>

      <view v-if="loading" class="loading">加载中…</view>

      <view class="list">
        <view
          v-for="(name, index) in diseases"
          :key="name"
          class="item"
          hover-class="item-hover"
          @tap="navigateToDiseasePage(name)"
        >
          <view class="item-left">
            <text class="idx">{{ String(index + 1).padStart(2, '0') }}</text>
            <text class="name">{{ shortName(name) }}</text>
          </view>
          <text class="more">查看</text>
        </view>
      </view>
      <view class="bottom-space" />
    </view>
    <uni-custom-tabbar class="tabbar" />
  </view>
</template>

<style scoped>
.page {
  position: relative;
  width: 100%;
  height: 100vh;
  overflow: hidden;
  background: #eef3fb;
}
.bg {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 42vh;
  z-index: 0;
  opacity: 0.95;
}
.scroll {
  position: relative;
  z-index: 1;
  height: 100vh;
  box-sizing: border-box;
  padding: 11vh 16px 0;
  overflow-y: auto;
  overflow-x: hidden;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none;
  -ms-overflow-style: none;
}
.scroll::-webkit-scrollbar {
  width: 0 !important;
  height: 0 !important;
  display: none !important;
  background: transparent;
}
.hero {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  margin-bottom: 14px;
  padding: 4px 4px 0;
}
.hero-text {
  flex: 1;
  padding-right: 8px;
}
.hero-kicker {
  display: block;
  font-size: 12px;
  color: #5b6b85;
  margin-bottom: 4px;
}
.hero-title {
  display: block;
  font-size: 26px;
  font-weight: 700;
  color: #1f2a44;
  margin-bottom: 6px;
}
.hero-desc {
  display: block;
  font-size: 13px;
  color: #718096;
}
.hero-art {
  width: 96px;
  height: 96px;
  flex-shrink: 0;
}
.loading {
  text-align: center;
  color: #999;
  padding: 16px;
}
.item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 14px;
  margin-bottom: 10px;
  border-radius: 14px;
  background: #fff;
  box-shadow: 0 4px 14px rgba(40, 90, 200, 0.06);
}
.item-hover {
  background: #f3f7ff;
}
.item-left {
  display: flex;
  align-items: center;
  min-width: 0;
  flex: 1;
  padding-right: 10px;
}
.idx {
  width: 34px;
  flex-shrink: 0;
  font-size: 15px;
  font-weight: 700;
  color: #285ac8;
  font-variant-numeric: tabular-nums;
}
.name {
  font-size: 15px;
  color: #1a202c;
  font-weight: 500;
}
.more {
  flex-shrink: 0;
  font-size: 12px;
  color: #285ac8;
  padding: 4px 10px;
  border-radius: 999px;
  background: rgba(40, 90, 200, 0.08);
}
.bottom-space {
  height: 14vh;
}
.tabbar {
  z-index: 999;
}
</style>
