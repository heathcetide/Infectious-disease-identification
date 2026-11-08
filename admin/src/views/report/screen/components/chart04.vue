<template>
  <div class="container">
    <div class="chart-container">
      <h3>打卡时间分布</h3>
      <el-radio-group
        v-model="timeRange"
        size="small"
        @change="updateChart"
        class="radio-group"
      >
        <el-radio-button value="最近一周">一周</el-radio-button>
        <el-radio-button value="最近一月">一月</el-radio-button>
        <el-radio-button value="最近一年">一年</el-radio-button>
      </el-radio-group>
      <div ref="chartContainer" class="chart"></div>
    </div>
  </div>
</template>

<script>
import * as echarts from "echarts";
import dayjs from "dayjs";
import { getCheckInDailyNumber } from "@/api/report/screen.js";
import { useThemeStore } from "@/store/theme";
import { screenChartColors } from "@/utils/screenTheme";

export default {
  props: {
    queryDate: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      // Do NOT put echarts instance in data — Vue Proxy breaks coordinateSystem.
      timeRange: "最近一周",
      themeStore: useThemeStore(),
    };
  },
  watch: {
    queryDate() {
      this.updateChart();
    },
    "themeStore.mode"() {
      this.updateChart();
    },
  },
  created() {
    this.chartInstance = null;
    this.renderSeq = 0;
    this.resizeTimer = null;
  },
  async mounted() {
    await this.$nextTick();
    const el = this.$refs.chartContainer;
    if (!el) return;
    if (el.clientWidth < 8 || el.clientHeight < 8) {
      el.style.minHeight = "140px";
    }
    this.chartInstance = echarts.init(el, undefined, {
      renderer: "canvas",
      useDirtyRect: false,
      width: el.clientWidth || 200,
      height: el.clientHeight || 140,
    });
    await this.updateChart();
    window.addEventListener("resize", this.handleResize);
  },
  beforeUnmount() {
    window.removeEventListener("resize", this.handleResize);
    clearTimeout(this.resizeTimer);
    this.renderSeq += 1;
    if (this.chartInstance && !this.chartInstance.isDisposed?.()) {
      this.chartInstance.dispose();
    }
    this.chartInstance = null;
  },
  methods: {
    handleResize() {
      clearTimeout(this.resizeTimer);
      this.resizeTimer = setTimeout(() => {
        if (!this.chartInstance || this.chartInstance.isDisposed?.()) return;
        try {
          this.chartInstance.resize();
        } catch (err) {
          console.warn("打卡图表 resize 跳过:", err);
        }
      }, 80);
    },
    async updateChart() {
      const end = this.queryDate ? dayjs(this.queryDate) : dayjs();
      let start;
      if (this.timeRange === "最近一周") {
        start = end.subtract(6, "day");
      } else if (this.timeRange === "最近一月") {
        start = end.subtract(29, "day");
      } else {
        start = end.subtract(1, "year");
      }
      const dateList = [start.format("YYYY-MM-DD"), end.format("YYYY-MM-DD")];
      const seq = ++this.renderSeq;

      try {
        const response = await getCheckInDailyNumber(dateList);
        if (seq !== this.renderSeq) return;
        if (response.data.code !== 1) {
          console.error("获取数据失败:", response.data.msg);
          return;
        }

        const rows = Array.isArray(response.data.data)
          ? response.data.data
          : [];
        const xAxisData = rows.map((item) => item.date);
        const seriesData = rows.map((item) => Number(item.count) || 0);

        if (!this.chartInstance || this.chartInstance.isDisposed?.()) return;
        const el = this.$refs.chartContainer;
        if (!el || el.clientWidth < 4 || el.clientHeight < 4) return;

        const tc = screenChartColors();
        const xs = xAxisData.length ? xAxisData : ["暂无数据"];
        const ys = seriesData.length ? seriesData : [0];

        this.chartInstance.setOption(
          {
            animation: false,
            color: ["#285AC8"],
            grid: {
              top: 12,
              bottom: this.timeRange === "最近一年" ? 36 : 28,
              left: 36,
              right: 8,
            },
            tooltip: {
              trigger: "axis",
              axisPointer: { type: "line" },
              formatter(params) {
                const p = Array.isArray(params) ? params[0] : null;
                if (!p) return "";
                return `${p.axisValue}<br/>打卡人数：${p.value}`;
              },
              backgroundColor: tc.tooltipBg,
              borderColor: tc.border,
              textStyle: { color: tc.tooltipText },
            },
            xAxis: {
              type: "category",
              data: xs,
              axisLabel: {
                show: true,
                rotate: this.timeRange === "最近一年" ? 45 : 0,
                formatter(v) {
                  return String(v).slice(5);
                },
                color: tc.muted,
                fontSize: 10,
              },
              axisLine: { show: false },
              axisTick: { show: false },
              splitLine: { show: false },
            },
            yAxis: {
              type: "value",
              axisLabel: { show: true, color: tc.muted, fontSize: 10 },
              axisLine: { show: false },
              axisTick: { show: false },
              splitLine: { show: false },
              minInterval: 1,
            },
            series: [
              {
                name: "打卡人数",
                type: "line",
                coordinateSystem: "cartesian2d",
                xAxisIndex: 0,
                yAxisIndex: 0,
                data: ys,
                smooth: true,
                progressive: 0,
                animation: false,
                showSymbol: this.timeRange !== "最近一年",
                symbol: "circle",
                symbolSize: this.timeRange === "最近一年" ? 3 : 6,
                itemStyle: { color: "#285AC8" },
                lineStyle: { width: 2, color: "#285AC8" },
                areaStyle: { color: "rgba(40,90,200,0.16)" },
              },
            ],
          },
          { notMerge: true, lazyUpdate: false }
        );
      } catch (error) {
        if (seq !== this.renderSeq) return;
        console.error("打卡时间分布图表失败:", error);
      }
    },
  },
};
</script>

<style scoped>
.container {
  background-color: transparent;
  margin-top: 0;
  margin-left: 0;
  flex: 1 1 0;
  min-height: 0;
  display: flex;
  flex-direction: column;
}

.chart-container {
  padding: 8px;
  background-color: var(--admin-card-bg, #fafafa);
  border: 1px solid var(--el-border-color-lighter, transparent);
  border-radius: 5px;
  width: 100%;
  max-width: 240px;
  flex: 1 1 0;
  min-height: 140px;
  height: auto;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
}

h3 {
  font-size: 14px;
  font-weight: bold;
  color: var(--el-text-color-primary, #4a4a4a);
  margin-bottom: 10px;
}

.radio-group {
  margin-top: 6px;
  margin-left: 0;
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.radio-group :deep(.el-radio-button__inner) {
  padding: 4px 8px;
  font-size: 11px;
  height: auto;
  line-height: 1.2;
}

.chart {
  width: 100%;
  flex: 1 1 auto;
  min-height: 100px;
  height: 100%;
  margin-top: 4px;
}
</style>
