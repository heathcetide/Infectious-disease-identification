<script>export default {
  methods: {
    goToWorkEnvironment() {
      // 跳转到工作环境基本信息页面
      uni.navigateTo({
        url: '/pages/environmentInfo/environmentInfo'
      });
    },
    goToMonthlyReport() {
      // 跳转到本月情况登记页面
      uni.navigateTo({
        url: '/pages/registerSituation/registerSituation'
      });
    }
  }
};
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <view class="environment-info">
      <button class="info-button" @tap="goToWorkEnvironment">工作环境基本信息</button>
      <button class="info-button" @tap="goToMonthlyReport">本月情况登记</button>
    </view>
    <image src="../../static/train-back.png" mode="widthFix" class="train-image"></image>
    <uni-custom-tabbar class="tabbar"></uni-custom-tabbar>
  </view>
</template>

<style>
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
  height: 100vh;
  z-index: -1; /* 将图片置于内容下方 */
}
.tabbar{
	z-index: 999;
}

.train-image {
  width: 133%;
  position: absolute;
  right: 0;
  left: 0;
  top: 34vh;
  z-index: 0;
}

.environment-info {
  position: fixed;
  bottom: 50%;
  left: 50%;
  transform: translate(-50%, 50%);
  width: 70%;
  z-index: 2;
}

.info-button {
  display: block;
  width: 100%;
  padding: 6px;
  margin: 20px 0;
  font-size: 18px;
  text-align: center;
  background-color: #FFFFFF;
  color: #285AC8;
  font-weight: 530;
  border-radius: 40px;
  border: none;
}

</style>
