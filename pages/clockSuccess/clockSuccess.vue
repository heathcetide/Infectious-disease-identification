<script setup>
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <view class="clock-info">
      <view class="clock-text">
        <text>您已打卡</text>
        <image src="../../static/icon-health.png" mode="aspectFit" class="health-icon"></image>
      </view>
      <text>今日<text style="font-size: 24px; color: #285AC8;">健康</text></text>
    </view>
    <image src="../../static/train-back.png" mode="widthFix" class="train-image"></image>
    <uni-custom-tabbar></uni-custom-tabbar>
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

.clock-info {
  position: fixed;
  bottom: 65vh;
  font-size: 20px;
  color: #646464;
  left: 5vh;
}

.clock-text {
  width: 100%;
  display: flex; /* 使用flex布局，使文字和图标在同一行 */
  align-items: center; /* 垂直居中对齐 */
  margin-bottom: -10vh;
}

.train-image {
  width: 133%;
  position: absolute;
  right: 0;
  left: 0;
  top: 36vh;
  z-index: 0;
}

.health-icon {
  width: 24px;
  margin-left: 8px; /* 添加左边距使图标和文字之间有间隔 */
}
</style>
