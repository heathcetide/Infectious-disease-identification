<template>
  <div class="screen-page">
    <Card :queryDate="queryDate">
      <template #extra>
        <el-form :inline="true" class="date-form" @submit.prevent>
          <el-form-item label="选择日期">
            <el-date-picker
              v-model="selectedDate"
              type="date"
              placeholder="选择日期"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              :disabled-date="disabledFutureDate"
              size="small"
              @change="onConfirm"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" size="small" @click="onConfirm">确 定</el-button>
          </el-form-item>
        </el-form>
      </template>
    </Card>
    <ChartEval />
    <div class="elsedata">
      <div class="chartspartone">
        <Chart01 :queryDate="queryDate" />
        <Chart04 :queryDate="queryDate" />
      </div>
      <div class="Map">
        <Map :queryDate="queryDate" />
      </div>
    </div>
  </div>
</template>

<script>
import Card from './components/card.vue';
import Chart01 from './components/chart01.vue';
import Map from './components/map.vue';
import Chart04 from './components/chart04.vue';
import ChartEval from './components/chartEval.vue';

export default {
  name: 'Screen',
  components: {
    Card,
    Chart01,
    Map,
    Chart04,
    ChartEval,
  },
  data() {
    return {
      selectedDate: '',
      queryDate: '',
    };
  },
  methods: {
    localToday() {
      const now = new Date();
      const y = now.getFullYear();
      const m = String(now.getMonth() + 1).padStart(2, '0');
      const d = String(now.getDate()).padStart(2, '0');
      return `${y}-${m}-${d}`;
    },
    disabledFutureDate(date) {
      const end = new Date();
      end.setHours(23, 59, 59, 999);
      return date.getTime() > end.getTime();
    },
    onConfirm() {
      this.queryDate = this.selectedDate || this.localToday();
    },
  },
  mounted() {
    const today = this.localToday();
    this.selectedDate = today;
    this.queryDate = today;
  },
};
</script>

<style scoped>
.screen-page {
  background: var(--el-bg-color, #ffffff);
  width: 100%;
  height: calc(100vh - 136px);
  max-height: calc(100vh - 136px);
  border-radius: 5px;
  box-sizing: border-box;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  padding-bottom: 4px;
}

.date-form {
  margin: 0;
}

.date-form :deep(.el-form-item) {
  margin-bottom: 0;
  margin-right: 8px;
}

.elsedata {
  flex: 1 1 auto;
  min-height: 0;
  display: flex;
  flex-wrap: nowrap;
  align-items: stretch;
  gap: 10px;
  padding: 0 10px 6px;
  box-sizing: border-box;
}

.chartspartone {
  flex: 0 0 200px;
  min-width: 180px;
  max-width: 220px;
  min-height: 0;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.Map {
  flex: 1 1 auto;
  min-width: 0;
  min-height: 0;
  display: flex;
}
</style>
