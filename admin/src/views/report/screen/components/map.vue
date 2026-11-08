<template>
  <div class="container">
    <div class="chart-container">
      <div ref="mapEl" class="chart"></div>
      <el-button class="reset-btn" size="small" @click="resetMapView">重置视角</el-button>
      <div v-if="statusText" class="status-text">{{ statusText }}</div>

      <!-- Bottom-right: collapsed by default; badge when alerts exist -->
      <div class="alert-fab" v-if="regionAlerts.length" :class="{ open: alertOpen }">
        <button
          type="button"
          class="alert-fab-toggle"
          :class="worstLevel"
          @click="alertOpen = !alertOpen"
        >
          <span class="fab-label">{{ alertOpen ? '收起预警' : '地区预警' }}</span>
          <span class="fab-badge">{{ regionAlerts.length }}</span>
        </button>

        <div class="alert-panel" v-show="alertOpen">
          <div class="alert-panel-head">
            <span>超阈值地区</span>
            <em>{{ criticalCount }} 紧急 · {{ warningCount }} 预警 · {{ watchCount }} 关注</em>
          </div>
          <div class="alert-list">
            <button
              v-for="item in regionAlerts"
              :key="item.regionKey"
              type="button"
              class="alert-item"
              :class="[item.level, { active: activeKey === item.regionKey }]"
              @click="focusRegion(item)"
            >
              <div class="alert-item-top">
                <span class="level-tag">{{ levelLabel(item.level) }}</span>
                <span class="region-name">{{ item.regionName }}</span>
              </div>
              <div class="alert-meta">
                {{ item.caseCount }} 例
                <template v-if="item.highRiskCount"> · 高风险 {{ item.highRiskCount }}</template>
                <template v-if="item.topDisease"> · {{ item.topDisease }}</template>
              </div>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onBeforeUnmount, onMounted, ref, watch } from 'vue';
import { getUserStation } from '@/api/report/screen.js';
import { loadAMap } from '@/utils/amap';

const props = defineProps({
  queryDate: {
    type: String,
    required: true,
  },
});

const mapEl = ref(null);
const statusText = ref('');
const regionAlerts = ref([]);
const alertOpen = ref(false);
const activeKey = ref('');

let map = null;
let heatmap = null;
let mapReady = false;
let disposed = false;
let AMapRef = null;
let lastBounds = null;
let loadSeq = 0;
let focusMarker = null;

const criticalCount = computed(
  () => regionAlerts.value.filter((a) => a.level === 'critical').length
);
const warningCount = computed(
  () => regionAlerts.value.filter((a) => a.level === 'warning').length
);
const watchCount = computed(
  () => regionAlerts.value.filter((a) => a.level === 'watch').length
);
const worstLevel = computed(() => {
  if (criticalCount.value) return 'critical';
  if (warningCount.value) return 'warning';
  return 'watch';
});

function levelLabel(level) {
  if (level === 'critical') return '紧急';
  if (level === 'warning') return '预警';
  return '关注';
}

function clearFocusMarker() {
  if (focusMarker) {
    try {
      map?.remove(focusMarker);
    } catch (_) {
      /* ignore */
    }
    focusMarker = null;
  }
}

function resetMapView() {
  if (!map || !AMapRef) return;
  activeKey.value = '';
  clearFocusMarker();
  if (lastBounds) {
    map.setBounds(lastBounds);
  } else {
    map.setZoomAndCenter(5.2, [94.5, 33.8]);
  }
}

function focusRegion(item) {
  if (!map || !AMapRef || !item) return;
  activeKey.value = item.regionKey;
  clearFocusMarker();
  const lng = Number(item.longitude);
  const lat = Number(item.latitude);
  if (!Number.isFinite(lng) || !Number.isFinite(lat)) return;

  focusMarker = new AMapRef.Marker({
    position: [lng, lat],
    anchor: 'bottom-center',
    offset: new AMapRef.Pixel(0, -6),
    zIndex: 200,
    content: `
      <div class="focus-pin ${item.level}">
        <div class="focus-pin-title">${item.regionName}</div>
        <div class="focus-pin-sub">${item.caseCount} 例${
          item.highRiskCount ? ` · 高风险 ${item.highRiskCount}` : ''
        }</div>
      </div>
    `,
  });
  map.add(focusMarker);
  map.setZoomAndCenter(8.6, [lng, lat], false, 450);
}

function aggregateHeatPoints(points) {
  const buckets = new Map();
  for (const p of points) {
    const lng = Number(p.longitude);
    const lat = Number(p.latitude);
    if (!Number.isFinite(lng) || !Number.isFinite(lat)) continue;
    const key = `${Math.round(lng / 0.08)}_${Math.round(lat / 0.08)}`;
    const prev = buckets.get(key);
    if (prev) {
      prev.count += 1;
      prev.lng = (prev.lng * (prev.count - 1) + lng) / prev.count;
      prev.lat = (prev.lat * (prev.count - 1) + lat) / prev.count;
    } else {
      buckets.set(key, { lng, lat, count: 1 });
    }
  }
  return [...buckets.values()];
}

function normalizePayload(data) {
  if (Array.isArray(data)) {
    return { stations: data, regionAlerts: [] };
  }
  return {
    stations: data?.stations || [],
    regionAlerts: data?.regionAlerts || [],
  };
}

async function loadHeatmap(newDate) {
  if (!newDate || disposed) return;
  const seq = ++loadSeq;
  if (!heatmap) return;
  try {
    statusText.value = '';
    const response = await getUserStation(newDate);
    if (disposed || seq !== loadSeq || !heatmap) return;
    if (response.data.code !== 1) {
      statusText.value = response.data.msg || '获取数据失败';
      return;
    }
    const { stations, regionAlerts: alerts } = normalizePayload(response.data.data);
    regionAlerts.value = alerts || [];
    // Keep collapsed when data refreshes; only show badge.
    if (!regionAlerts.value.length) {
      alertOpen.value = false;
      activeKey.value = '';
      clearFocusMarker();
    }

    const raw = stations.filter(
      (item) => Number.isFinite(item.longitude) && Number.isFinite(item.latitude)
    );
    const heatmapData = aggregateHeatPoints(raw);
    const maxCount = Math.max(3, ...heatmapData.map((p) => p.count), 3);

    try {
      heatmap.setDataSet({
        data: heatmapData,
        max: maxCount,
      });
    } catch (e) {
      console.warn('热力图层更新失败:', e);
      return;
    }

    if (heatmapData.length > 0 && map && AMapRef) {
      const lngs = heatmapData.map((p) => p.lng);
      const lats = heatmapData.map((p) => p.lat);
      const padLng = Math.max(1.2, (Math.max(...lngs) - Math.min(...lngs)) * 0.12);
      const padLat = Math.max(0.9, (Math.max(...lats) - Math.min(...lats)) * 0.12);
      const sw = [Math.min(...lngs) - padLng, Math.min(...lats) - padLat];
      const ne = [Math.max(...lngs) + padLng, Math.max(...lats) + padLat];
      lastBounds = new AMapRef.Bounds(sw, ne);
      map.setBounds(lastBounds);
    } else if (!heatmapData.length) {
      statusText.value = '当日暂无打卡坐标';
    }
  } catch (error) {
    if (disposed || seq !== loadSeq) return;
    console.error('请求出错:', error);
    statusText.value = '地图数据加载失败';
  }
}

onMounted(async () => {
  disposed = false;
  try {
    statusText.value = '地图加载中…';
    AMapRef = await loadAMap();
    if (disposed) return;

    map = new AMapRef.Map(mapEl.value, {
      zoom: 5.2,
      center: [94.5, 33.8],
      viewMode: '2D',
      mapStyle: 'amap://styles/normal',
    });

    await new Promise((resolve, reject) => {
      const timer = setTimeout(() => reject(new Error('HeatMap 插件加载超时')), 15000);
      map.plugin(['AMap.HeatMap'], () => {
        clearTimeout(timer);
        resolve();
      });
    });
    if (disposed) return;

    heatmap = new AMapRef.HeatMap(map, {
      radius: 32,
      opacity: [0, 0.9],
      zooms: [3, 18],
      gradient: {
        0.15: 'rgba(0, 120, 255, 0.55)',
        0.35: 'rgba(0, 220, 180, 0.7)',
        0.55: 'rgba(255, 230, 0, 0.85)',
        0.75: 'rgba(255, 120, 0, 0.92)',
        1.0: 'rgba(220, 0, 0, 1)',
      },
    });
    mapReady = true;
    statusText.value = '';
    if (props.queryDate) {
      await loadHeatmap(props.queryDate);
    }
  } catch (error) {
    console.error('地图初始化失败:', error);
    if (!disposed) {
      statusText.value = '高德地图加载失败，请检查网络后刷新';
    }
  }
});

onBeforeUnmount(() => {
  disposed = true;
  mapReady = false;
  loadSeq += 1;
  clearFocusMarker();
  try {
    heatmap?.setMap?.(null);
  } catch (_) {
    /* ignore */
  }
  try {
    map?.destroy?.();
  } catch (_) {
    /* ignore */
  }
  heatmap = null;
  map = null;
});

watch(
  () => props.queryDate,
  async (newDate, oldDate) => {
    if (!newDate || newDate === oldDate || disposed) return;
    if (!mapReady) return;
    await loadHeatmap(newDate);
  }
);
</script>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  min-height: 0;
}

.chart-container {
  width: 100%;
  height: 100%;
  min-height: 0;
  background-color: var(--admin-card-bg, #fafafa);
  border: 1px solid var(--el-border-color-lighter, transparent);
  border-radius: 5px;
  position: relative;
  overflow: hidden;
  box-sizing: border-box;
}

.chart {
  width: 100%;
  height: 100%;
}

.reset-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  z-index: 5;
  background: var(--el-bg-color-overlay, rgba(255, 255, 255, 0.92)) !important;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.12);
}

.status-text {
  position: absolute;
  left: 50%;
  top: 48%;
  transform: translate(-50%, -50%);
  color: var(--el-text-color-secondary, #888);
  font-size: 13px;
  pointer-events: none;
  z-index: 2;
  background: color-mix(in srgb, var(--admin-card-bg, #fafafa) 88%, transparent);
  padding: 6px 12px;
  border-radius: 4px;
}

.alert-fab {
  position: absolute;
  right: 12px;
  bottom: 12px;
  z-index: 8;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 8px;
  max-width: min(320px, calc(100% - 24px));
}

.alert-fab-toggle {
  position: relative;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border: none;
  cursor: pointer;
  padding: 8px 14px;
  border-radius: 999px;
  background: #fff;
  color: #0f172a;
  font-size: 13px;
  font-weight: 650;
  box-shadow: 0 6px 18px rgba(15, 23, 42, 0.16);
  border: 1px solid rgba(15, 23, 42, 0.08);
}

.alert-fab-toggle.warning {
  border-color: rgba(217, 119, 6, 0.35);
}

.alert-fab-toggle.critical {
  border-color: rgba(220, 38, 38, 0.4);
}

.fab-badge {
  min-width: 20px;
  height: 20px;
  padding: 0 6px;
  border-radius: 999px;
  background: #dc2626;
  color: #fff;
  font-size: 11px;
  font-weight: 700;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.alert-fab-toggle.warning .fab-badge {
  background: #d97706;
}

.alert-fab-toggle.watch .fab-badge {
  background: #0284c7;
}

.alert-panel {
  width: 280px;
  max-height: min(42vh, 300px);
  background: rgba(255, 255, 255, 0.97);
  border: 1px solid rgba(15, 23, 42, 0.1);
  border-radius: 10px;
  box-shadow: 0 10px 28px rgba(15, 23, 42, 0.16);
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.alert-panel-head {
  padding: 10px 12px 8px;
  display: flex;
  justify-content: space-between;
  gap: 8px;
  align-items: baseline;
  border-bottom: 1px solid #f1f5f9;
  font-size: 13px;
  font-weight: 700;
  color: #0f172a;
}

.alert-panel-head em {
  font-style: normal;
  font-size: 11px;
  color: #64748b;
  font-weight: 500;
  white-space: nowrap;
}

.alert-list {
  overflow: auto;
  padding: 8px;
}

.alert-item {
  display: block;
  width: 100%;
  text-align: left;
  border: none;
  cursor: pointer;
  padding: 8px 9px;
  border-radius: 8px;
  margin-bottom: 6px;
  background: #f8fafc;
  border-left: 3px solid #0284c7;
}

.alert-item:last-child {
  margin-bottom: 0;
}

.alert-item:hover,
.alert-item.active {
  background: #eff6ff;
}

.alert-item.warning {
  border-left-color: #d97706;
}

.alert-item.critical {
  border-left-color: #dc2626;
}

.alert-item-top {
  display: flex;
  align-items: center;
  gap: 6px;
}

.level-tag {
  font-size: 10px;
  padding: 1px 6px;
  border-radius: 999px;
  color: #fff;
  background: #0284c7;
  font-weight: 700;
}

.alert-item.warning .level-tag {
  background: #d97706;
}

.alert-item.critical .level-tag {
  background: #dc2626;
}

.region-name {
  color: #0f172a;
  font-size: 12px;
  font-weight: 700;
}

.alert-meta {
  margin-top: 4px;
  color: #64748b;
  font-size: 11px;
  line-height: 1.35;
}

:deep(.focus-pin) {
  min-width: 110px;
  padding: 7px 9px;
  border-radius: 8px;
  background: #fff;
  border: 1px solid #cbd5e1;
  box-shadow: 0 8px 18px rgba(15, 23, 42, 0.18);
  text-align: center;
  color: #0f172a;
}

:deep(.focus-pin.watch) {
  border-color: #7dd3fc;
}

:deep(.focus-pin.warning) {
  border-color: #fbbf24;
}

:deep(.focus-pin.critical) {
  border-color: #f87171;
}

:deep(.focus-pin-title) {
  font-size: 12px;
  font-weight: 700;
}

:deep(.focus-pin-sub) {
  margin-top: 2px;
  font-size: 11px;
  color: #64748b;
}
</style>
