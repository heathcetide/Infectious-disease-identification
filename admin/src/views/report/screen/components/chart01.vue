<template>
  <div class="container">
    <div class="title-container">
      <div class="blue-box"></div>
      <span class="title-text">其他数据</span>
    </div>
    <div class="chart-container">
      <div class="chart-title">当日疾病人数与占比</div>
      <div ref="chart" class="chart"></div>
    </div>
  </div>
</template>

<script setup>
import * as echarts from 'echarts';
import { ref, onMounted, watch, onBeforeUnmount, nextTick } from 'vue';
import { getDiseaseDataToday } from '@/api/report/screen.js';
import { useDiseaseStore } from '@/store/diseaseStore';
import { useThemeStore } from '@/store/theme';
import { screenChartColors } from '@/utils/screenTheme';

const chart = ref(null);
const diseaseStore = useDiseaseStore();
const themeStore = useThemeStore();
let myChart = null;
let pendingData = null;
let lastData = [];

const props = defineProps({
  queryDate: { type: String, default: '' },
});

const diseaseTypes = [
  '新型冠状病毒感染', '流感', '鼠疫', '感染性腹泻', '炭疽', '结核病',
  '登革热（蚊媒传染病）', '疟疾（蚊媒传染病）', '森林脑炎（蜱媒传染病）',
  '发热伴血小板减少综合征（蜱媒传染病）', '斑疹伤寒', '流行性热出血', '健康',
];

const shortLabels = [
  '新冠', '流感', '鼠疫', '腹泻', '炭疽', '结核',
  '登革热', '疟疾', '森脑', '发热减板', '斑疹', '出血热', '健康',
];

function buildSeries(data) {
  const diseaseDataMap = Object.fromEntries(diseaseTypes.map((type) => [type, 0]));
  (data || []).forEach((item) => {
    const typeName = item?.diagnosisDiseaseTypeName || item?.diseaseTypeName || '健康';
    if (typeName in diseaseDataMap) {
      diseaseDataMap[typeName] = Number(item.count) || 0;
    }
  });
  const countData = diseaseTypes.map((type) => diseaseDataMap[type]);
  const total = countData.reduce((acc, n) => acc + n, 0);
  const percentData = countData.map((n) => (total > 0 ? Number(((n / total) * 100).toFixed(2)) : 0));
  return { countData, percentData };
}

function renderChart(data) {
  if (!myChart) {
    pendingData = data;
    return;
  }
  lastData = data || [];
  const { countData, percentData } = buildSeries(data);
  const tc = screenChartColors();
  try {
    myChart.clear();
    myChart.setOption(
      {
        animation: false,
        color: ['#285AC8', '#D43030'],
        legend: {
          top: 28,
          right: 4,
          itemWidth: 10,
          itemHeight: 6,
          textStyle: { fontSize: 10, color: tc.text },
          data: ['人数', '占比%'],
        },
        grid: { top: '38%', left: '12%', right: '12%', bottom: '18%' },
        xAxis: {
          type: 'category',
          data: shortLabels,
          axisLabel: { show: true, fontSize: 9, color: tc.muted, rotate: 35, interval: 0 },
          axisLine: { show: false },
          axisTick: { show: false },
          splitLine: { show: false },
        },
        yAxis: [
          {
            type: 'value',
            name: '人',
            nameTextStyle: { fontSize: 10, color: tc.axis },
            axisLine: { show: false },
            axisTick: { show: false },
            splitLine: { show: false },
            axisLabel: { fontSize: 9, color: tc.axis },
          },
          {
            type: 'value',
            name: '%',
            nameTextStyle: { fontSize: 10, color: tc.axis },
            axisLine: { show: false },
            axisTick: { show: false },
            splitLine: { show: false },
            axisLabel: { fontSize: 9, color: tc.axis },
            max: 100,
          },
        ],
        tooltip: {
          trigger: 'axis',
          axisPointer: { type: 'line' },
          formatter(params) {
            if (!params || !params.length) return '';
            const idx = params[0]?.dataIndex ?? 0;
            const name = diseaseTypes[idx] || '';
            const count = countData[idx] || 0;
            const pct = percentData[idx] || 0;
            return `${name}<br/>人数：${count}人<br/>占比：${pct}%`;
          },
          backgroundColor: tc.tooltipBg,
          borderColor: tc.border,
          borderWidth: 1,
          padding: [10, 10],
          textStyle: { color: tc.tooltipText, fontSize: 13 },
          extraCssText: 'border-radius: 4px;',
        },
        series: [
          {
            name: '人数',
            type: 'line',
            yAxisIndex: 0,
            data: countData,
            smooth: true,
            showSymbol: true,
            symbol: 'circle',
            symbolSize: 5,
            progressive: 0,
            lineStyle: { width: 2, color: '#285AC8' },
            itemStyle: { color: '#285AC8' },
            areaStyle: { color: 'rgba(40,90,200,0.08)' },
          },
          {
            name: '占比%',
            type: 'line',
            yAxisIndex: 1,
            data: percentData,
            smooth: true,
            showSymbol: true,
            symbol: 'circle',
            symbolSize: 5,
            progressive: 0,
            lineStyle: { width: 2, type: 'dashed', color: '#D43030' },
            itemStyle: { color: '#D43030' },
          },
        ],
      },
      { notMerge: true, lazyUpdate: false }
    );
    nextTick(() => myChart?.resize());
  } catch (err) {
    console.error('疾病图表渲染失败:', err);
  }
}

async function loadData(date) {
  if (!date) return;
  try {
    const response = await getDiseaseDataToday(date);
    if (response.data.code === 1) {
      const data = response.data.data || [];
      diseaseStore.setDiseaseData(data);
      renderChart(data);
    } else {
      renderChart([]);
    }
  } catch (e) {
    console.error(e);
    renderChart([]);
  }
}

const onResize = () => myChart?.resize();

onMounted(async () => {
  await nextTick();
  if (!chart.value) return;
  // Avoid progressive-render crashes when the panel is still 0×0.
  const el = chart.value;
  if (el.clientWidth < 8 || el.clientHeight < 8) {
    el.style.minHeight = '140px';
  }
  myChart = echarts.init(el, undefined, { renderer: 'canvas', useDirtyRect: false });
  myChart.resize();
  if (pendingData != null) {
    renderChart(pendingData);
    pendingData = null;
  }
  watch(() => props.queryDate, (d) => loadData(d), { immediate: true });
  watch(() => themeStore.mode, () => renderChart(lastData));
  window.addEventListener('resize', onResize);
});

onBeforeUnmount(() => {
  window.removeEventListener('resize', onResize);
  myChart?.dispose();
  myChart = null;
});
</script>

<style scoped>
.container {
  margin-top: 0;
  width: 100%;
  flex: 1 1 0;
  min-height: 0;
  display: flex;
  flex-direction: column;
}
.title-container {
  display: flex;
  margin-left: 0px;
  margin-bottom: 4px;
  margin-top: 0;
  flex-shrink: 0;
}
.blue-box {
  width: 6px;
  height: 14px;
  background-color: #285ac8;
  margin-right: 8px;
}
.title-text {
  font-size: 12px;
  font-weight: bold;
  color: var(--el-text-color-primary, #4a4a4a);
}
.chart-container {
  width: 100%;
  max-width: 240px;
  flex: 1 1 0;
  min-height: 140px;
  background-color: var(--admin-card-bg, #fafafa);
  border: 1px solid var(--el-border-color-lighter, transparent);
  border-radius: 5px;
  margin-left: 0;
  margin-top: 0;
  margin-bottom: 0;
  position: relative;
  padding: 6px;
  box-sizing: border-box;
}
.chart-title {
  position: absolute;
  top: 10px;
  left: 10px;
  font-size: 13px;
  font-weight: bold;
  color: var(--el-text-color-primary, #4a4a4a);
  background-color: var(--admin-card-bg, #fafafa);
  padding: 2px 5px;
  border-radius: 3px;
  z-index: 10;
}
.chart {
  width: 100%;
  height: 100%;
  min-height: 120px;
}
</style>
