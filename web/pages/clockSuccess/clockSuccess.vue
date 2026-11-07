<script setup>
</script>

<template>
  <view class="page">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />

    <view class="body">
      <view class="hero">
        <view class="hero-text">
          <view class="hero-title-row">
            <text class="hero-title">您已打卡</text>
            <image class="status-icon" src="/static/icon-health.png" mode="aspectFit" />
          </view>
          <text class="hero-desc">
            今日
            <text class="hero-em">健康</text>
          </text>
        </view>
      </view>

      <view class="panel">
        <image class="train-art" src="/static/train-back.png" mode="widthFix" />
        <text class="tip">继续保持良好作息与防护习惯</text>
      </view>
    </view>

    <uni-custom-tabbar />
  </view>
</template>

<style>
.page {
  position: relative;
  width: 100%;
  height: 100vh;
  overflow: hidden;
  background: #eef3fb;
}
.bg {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 38vh;
  z-index: 0;
}
.body {
  position: relative;
  z-index: 1;
  height: calc(100vh - 11vh);
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  padding: 9vh 0 0;
}
.hero {
  flex-shrink: 0;
  padding: 0 24px 12px;
}
.hero-title-row {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}
.hero-title {
  font-size: 22px;
  font-weight: 700;
  color: #2a3348;
}
.status-icon {
  width: 22px;
  height: 22px;
  margin-left: 8px;
}
.hero-desc {
  font-size: 20px;
  color: #4a5568;
  font-weight: 600;
}
.hero-em {
  font-size: 26px;
  color: #285ac8;
  font-weight: 700;
}
.panel {
  flex: 1;
  min-height: 0;
  background: #fff;
  border-radius: 28px 28px 0 0;
  box-shadow: 0 -6px 24px rgba(40, 90, 200, 0.08);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding: 28px 16px 16px;
  box-sizing: border-box;
  overflow: hidden;
}
.train-art {
  width: 112%;
  margin-left: -6%;
  margin-top: 4px;
}
.tip {
  margin-top: 8px;
  font-size: 13px;
  color: #94a3b8;
}
</style>
