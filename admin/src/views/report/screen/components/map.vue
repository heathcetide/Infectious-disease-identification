<template>
  <div class="container">
    <div class="chart-container">
      <div ref="mapEl" class="chart"></div>
      <el-button class="reset-btn" size="small" @click="resetMapView">重置视角</el-button>
      <div v-if="statusText" class="status-text">{{ statusText }}</div>
    </div>
  </div>
</template>

<script setup>
import { onBeforeUnmount, onMounted, ref, watch } from 'vue';
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
let map = null;
let heatmap = null;
let mapReady = false;
let disposed = false;
let AMapRef = null;
let lastBounds = null;
let loadSeq = 0;

function resetMapView() {
  if (!map || !AMapRef) return;
  if (lastBounds) {
    map.setBounds(lastBounds);
  } else {
    map.setZoomAndCenter(5.2, [94.5, 33.8]);
  }
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
    const raw = (response.data.data || []).filter(
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
  background-color: #fafafa;
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
  background: rgba(255, 255, 255, 0.92) !important;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.12);
}

.status-text {
  position: absolute;
  left: 50%;
  top: 48%;
  transform: translate(-50%, -50%);
  color: #888;
  font-size: 13px;
  pointer-events: none;
  z-index: 2;
  background: rgba(250, 250, 250, 0.88);
  padding: 6px 12px;
  border-radius: 4px;
}
</style>
