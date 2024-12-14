<script>
import store from '../../store/index.ts';
import { clockSituationService } from '../../services/clock';
export default {
  methods: {
	formatDate(date) {
	    const year = date.getFullYear();
	    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以加1
	    const day = String(date.getDate()).padStart(2, '0');
	    return `${year}-${month}-${day}`;
	},
    checkSuccess(response) {
      console.log("666666666", response);
      // 接收到的数据进行排序
      const sortedDiseases = response.data.diseasescoringList.sort((a, b) => b.score - a.score);
      this.diseases = sortedDiseases;
    },
	checkError(response){
		console.log(response)
	},  
  },
  mounted() {
    // 获取 store 中的 diseasescoring 数据
    const diseasescoring = store.state.diseasescoring;
    // 将 diseasescoring 转换为符合当前数据结构的形式
    const formattedDiseases = diseasescoring.map((item) => ({
      diseaseTypeName: item.diseaseTypeName,
      score: item.score,
      link: '/pages/healthProtection/healthProtection',
    }));
    // 对 formattedDiseases 进行排序
    formattedDiseases.sort((a, b) => b.score - a.score);
    // 将排序后的数据赋值给 diseases
    this.diseases = formattedDiseases;
	
	const now = new Date();
	const formattedDate = this.formatDate(now);
	// 设置到 feedBackForm 中
	this.dateForm.date = formattedDate;
	clockSituationService(this.dateForm)
		.then(res => this.checkSuccess(res))
		.catch(error => this.checkError(error));
  },
  data() {
    return {
	  dateForm:{
	  		date: '',
	  },
      diseases: [
        { diseaseTypeName: '新型冠状病毒', score: 90, link: '/pages/healthProtection/healthProtection' },
        { diseaseTypeName: '流感', score: 80, link: '/pages/healthProtection/healthProtection' },
        { diseaseTypeName: '鼠疫', score: 75, link: '/pages/healthProtection/healthProtection' },
        { diseaseTypeName: '感染性腹泻', score: 60, link: '/pages/healthProtection/healthProtection' },
        { diseaseTypeName: '炭疽', score: 50, link: '/pages/healthProtection/healthProtection' },
      ],
    };
  },
};
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <image src="../../static/virus.png" mode="aspectFit" class="virus-image"></image>
    <view class="clock-info">
      <view class="clock-text">
        <text>您已打卡</text>
        <image src="../../static/icon-Symptomatic.png" mode="aspectFit" class="symptomatic-icon"></image>
      </view>
      <text>
        <text class="clock-texts">根据您提交的症状</text>
        <text class="clock-texts">您可能的疾病是：</text>
      </text>
    </view>
    <view class="content">
      <view class="list-container">
        <view class="list-item" v-for="(item, index) in diseases" :key="index">
          <view class="disease-info">
            <view class="index-box">
              <text class="index">{{ index + 1 }}</text>
            </view>
            <view class="text-box">
              <text class="disease-name">{{ item.diseaseTypeName }}</text>
              <navigator class="health-protection" :url="item.link">
                <text>健康防护</text>
                <text>></text>
              </navigator>
            </view>
          </view>
          <view class="rate-box">
            <text class="rate">{{ item.score }}%</text>
            <text class="rate-label">患病率</text>
          </view>
        </view>
      </view>
    </view>
    <uni-custom-tabbar class="tabbar"></uni-custom-tabbar>
  </view>
</template>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  position: relative; /* 设置为相对定位，使内部元素能够相对于容器定位 */
}

.background-image {
  position: absolute; /* 设置为绝对定位，使其占据整个容器 */
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1; /* 将图片置于内容下方 */
}

.clock-info {
  position: fixed;
  bottom: 65vh; /* 调整位置，使其离下面的内容更近 */
  font-size: 20px;
  color: #646464;
  left: 5vh;
}

.clock-text {
  display: flex; /* 使用 flex 布局，使文字和图标在同一行 */
  align-items: center; /* 垂直居中对齐 */
  margin-bottom: -10vh; /* 减小底部间距 */
}

.clock-texts {
  font-size: 18px;
  display: block;
  font-weight: 800;
}

.virus-image {
  position: absolute;
  width: 170px;
  right: -20px;
  top: 12vh;
}

.content {
  position: fixed;
  bottom: 85px;
  left: 0;
  right: 0;
  width: 100%;
  height: 45vh; /* 调整高度，避免覆盖 tabbar */
  overflow-y: auto; /* 如果内容超出容器，则显示滚动条 */
  background-color: #ffffff;
  z-index: 20;
  border-top-left-radius: 40px;
  border-top-right-radius: 40px;
}

.tabbar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 11vh;
  z-index: 1000; /* 确保 tabbar 的层级高于其他内容 */
  background-color: #ffffff;
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
}

.list-container {
  padding: 16px;
}

.list-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #ffffff;
  padding: 12px;
  margin-bottom: 10px;
  border-radius: 8px;
}

.disease-info {
  display: flex;
  justify-content: space-between;
}

.disease-name {
  font-size: 16px;
  color: #333333;
  margin-bottom: 4px;
}

.health-protection {
  display: flex;
  align-items: center;
  color: #007aff;
  font-size: 14px;
  text-decoration: none;
}

.rate-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #285AC8;
  color: #ffffff;
  padding: 8px;
  border-radius: 8px;
  min-width: 39px;
  min-height: 47px;
}

.rate {
  font-size: 16px;
  font-weight: bold;
}

.rate-label {
  font-size: 12px;
}

.index-box {
  margin-right: 12px;
}

.index {
  font-size: 16px;
  color: #E5E5E5;
  font-weight: bold;
}

.symptomatic-icon {
  width: 24px;
  margin-left: 8px; /* 控制图标和文字之间的间距 */
}
</style>
