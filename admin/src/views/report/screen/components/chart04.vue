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
      chartInstance: null,
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
  async mounted() {
    await this.$nextTick();
    if (!this.$refs.chartContainer) return;
    this.chartInstance = echarts.init(this.$refs.chartContainer);
    this.chartInstance.resize();
    await this.updateChart();
    window.addEventListener("resize", this.handleResize);
  },
  beforeUnmount() {
    window.removeEventListener("resize", this.handleResize);
    this.chartInstance?.dispose();
  },
  methods: {
    handleResize() {
      this.chartInstance?.resize();
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

      try {
        const response = await getCheckInDailyNumber(dateList);
        if (response.data.code !== 1) {
          console.error("获取数据失败:", response.data.msg);
          return;
        }

        const rows = Array.isArray(response.data.data)
          ? response.data.data
          : [];
        const xAxisData = rows.map((item) => item.date);
        const seriesData = rows.map((item) => Number(item.count) || 0);

        if (!this.chartInstance) return;
        const tc = screenChartColors();
        this.chartInstance.setOption({
          color: ["#285AC8"],
          grid: {
            top: "12%",
            bottom: this.timeRange === "最近一年" ? "22%" : "16%",
            left: "14%",
            right: "4%",
          },
          tooltip: {
            trigger: "axis",
            axisPointer: { type: "shadow" },
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
            data: xAxisData.length ? xAxisData : ["暂无数据"],
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
            axisLabel: { show: true, color: tc.muted },
            axisLine: { show: false },
            axisTick: { show: false },
            splitLine: { show: false },
            minInterval: 1,
          },
          series: [
            {
              name: "打卡人数",
              data: seriesData.length ? seriesData : [0],
              type: "line",
              smooth: true,
              symbolSize: this.timeRange === "最近一年" ? 4 : 8,
              itemStyle: { color: "#285AC8" },
              lineStyle: { width: 3 },
              areaStyle: {
                color: {
                  type: "linear",
                  x: 0,
                  y: 0,
                  x2: 0,
                  y2: 1,
                  colorStops: [
                    { offset: 0, color: "rgba(40,90,200,0.28)" },
                    { offset: 1, color: "rgba(40,90,200,0.02)" },
                  ],
                },
              },
            },
          ],
        }, { notMerge: true });
        this.$nextTick(() => this.chartInstance?.resize());
      } catch (error) {
        console.error("请求出错:", error);
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
