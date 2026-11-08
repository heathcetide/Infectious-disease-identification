<template>
  <div class="container">
    <div class="title-row">
      <div class="title-container">
        <div class="blue-box"></div>
        <span class="title-text">打卡数据监控大屏</span>
      </div>
      <div class="title-extra">
        <slot name="extra"></slot>
      </div>
    </div>
    <div class="card-container">
        <div class="card" v-for="(item, index) in cardData" :key="index">
            <div class="card-title">{{ item.title }}</div>
            <div class="card-bottom">
                <div class="card-data" :id="'card-data-' + index">{{ item.data }}</div>
                <img :src="item.imgSrc" alt=""/>
            </div>
        </div>
    </div>
  </div>
</template>

<script>
import { CountUp } from 'countup.js';
import allusernumber from '@/assets/screenimgs/allusernumber.svg';
import allcheckinnumber from '@/assets/screenimgs/allcheckinnumber.svg';
import todaycheckin from '@/assets/screenimgs/todaycheckin.svg';
import todayhealth from '@/assets/screenimgs/todayhealth.svg';
import todaydisease from '@/assets/screenimgs/todaydisease.svg';
import todaynotcheckin from '@/assets/screenimgs/todaynotcheckin.svg';
import { getStatisticsData } from '@/api/report/screen.js';

export default {
  props: {
    queryDate: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      cardData: [
        { title: "职工总人数 (人)", data: 0, imgSrc: allusernumber },
        { title: "打卡总次数 (次)", data: 0, imgSrc: allcheckinnumber },
        { title: "当日打卡人数 (人)", data: 0, imgSrc: todaycheckin },
        { title: "当日健康人数 (人)", data: 0, imgSrc: todayhealth },
        { title: "当日患病人数 (人)", data: 0, imgSrc: todaydisease },
        { title: "当日未打卡人数 (人)", data: 0, imgSrc: todaynotcheckin }
      ]
    };
  },
  methods: {
    startCountAnimation() {
      this.cardData.forEach((item, index) => {
        const countUp = new CountUp(`card-data-${index}`, item.data);
        countUp.start();
      });
    },

    async fetchStatisticsData(date) {
      if (!date) return;
      try {
        const response = await getStatisticsData(date);
        if (response.data.code === 1) {
          const statistics = response.data.data;
          this.cardData[0].data = statistics.userNumber;
          this.cardData[1].data = statistics.checkInNumber;
          this.cardData[2].data = statistics.userHealthNumber + statistics.userDiseaseNumber;
          this.cardData[3].data = statistics.userHealthNumber;
          this.cardData[4].data = statistics.userDiseaseNumber;
          this.cardData[5].data = Math.max(
            0,
            statistics.userNumber - (statistics.userHealthNumber + statistics.userDiseaseNumber)
          );
          this.startCountAnimation();
        } else {
          this.$message.error("获取统计数据失败，请重试！");
        }
      } catch (error) {
        console.error("Error fetching statistics data:", error);
        this.$message.error("获取统计数据失败，请重试！");
      }
    }
  },
  watch: {
    queryDate: {
      immediate: true,
      handler(newDate) {
        if (newDate) {
          this.fetchStatisticsData(newDate);
        }
      },
    },
  },
};
</script>

<style scoped>
.container {
  margin-top: 4px;
  width: 100%;
  height: auto;
}
.title-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  padding: 2px 10px 0;
  box-sizing: border-box;
}

.title-container {
  display: flex;
  align-items: center;
  margin: 0;
  flex-shrink: 0;
}

.title-extra {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  min-width: 0;
}

.blue-box {
  width: 6px;
  height: 18px;
  background-color: #285ac8;
  margin-right: 10px;
}

.title-text {
  font-size: 12px;
  font-weight: bold;
  color: var(--el-text-color-primary, #4a4a4a);
}
.card-container {
  display: flex;
  flex-wrap: nowrap;
  gap: 6px;
  padding: 0 10px 4px;
  box-sizing: border-box;
}

.card {
  background: var(--admin-card-bg, #fafafa);
  border: 1px solid var(--el-border-color-lighter, transparent);
  border-radius: 8px;
  height: 72px;
  flex: 1 1 0;
  min-width: 0;
  margin: 0;
}

.card-title {
  margin: 8px 10px 2px;
  font-size: 11px;
  font-weight: 550;
  color: var(--el-text-color-regular, #4a4a4a);
}

.card-bottom {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-right: 16px;
}

img {
  height: 28px;
  width: 28px;
  margin-left: 0;
  margin-top: 0;
  flex-shrink: 0;
}

.card-data {
  min-width: 48px;
  font-size: 18px;
  margin-top: 2px;
  margin-left: 10px;
  font-weight: bold;
  color: var(--el-text-color-primary, #333333);
}
</style>
