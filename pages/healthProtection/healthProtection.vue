<script>
export default {
  methods: {
    navigateToDiseasePage(diseaseTypeName) {
      uni.navigateTo({
            url: `/pages/publicityMaterial/publicityMaterial?diseaseTypeName=${diseaseTypeName}`
      });
    },
  },
};
</script>

<template>
	<view class="container" >
	  <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  <view class="content" v-show="!showMaterial">
	    <view class="header">
	      <text class="title">疾病</text>知识
		  <image src="../../static/knowledgeIcon.png" mode="widthFix" class="icons"></image>
	    </view>
			<image src="../../static/bacteria2.png" mode="widthFix" class="back-image" ></image>
		  <view class="diseases-list">
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('新型冠状病毒感染')">
		      <text class="disease-name"><text class="indexNumber">1</text> 新型冠状病毒</text>
			  <text class="examine">查看  ></text>
		    </view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('流感')">
		      <text class="disease-name"><text class="indexNumber">2</text> 流感</text>
			  <text class="examine">查看  ></text>
			</view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('鼠疫')">
		      <text class="disease-name"><text class="indexNumber">3</text>鼠疫</text>
			  <text class="examine">查看  ></text>
			</view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('感染性腹泻')">
		      <text class="disease-name"><text class="indexNumber">4</text> 感染性腹泻</text>
			  <text class="examine">查看  ></text>
			</view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('炭疽')">
		      <text class="disease-name"><text class="indexNumber">5</text> 炭疽</text>
			  <text class="examine">查看  ></text>
			</view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('结核病')">
		      <text class="disease-name"><text class="indexNumber">6</text>结核病</text>
			  <text class="examine">查看  ></text>
			</view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('蚊传(输入性为主)')">
		      <text class="disease-name"><text class="indexNumber">7</text>蚊传(输入性为主)</text>
			  <text class="examine">查看  ></text>
			</view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('蜱传(森林脑炎和莱姆病)')">
		      <text class="disease-name"><text class="indexNumber">8</text>蜱传(森林脑炎和莱姆病)</text>
			  <text class="examine">查看  ></text>
			</view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('斑疹伤寒')">
		      <text class="disease-name"><text class="indexNumber">9</text> 斑疹伤寒</text>
			  <text class="examine">查看  ></text>
			</view>
		    <view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('流行性出血热')">
		      <text class="disease-name"><text class="indexNumber">10</text>流行性出血热</text>
			  <text class="examine">查看  ></text>
			</view>
			<view class="disease-item" hover-class="hovered" @tap="navigateToDiseasePage('其他')">
			  <text class="disease-name"><text class="indexNumber">11</text>其他</text>
			  <text class="examine">查看  ></text>
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
.icons{
	width: 26px;
	margin-left: 7px;
}
.content{
	position: relative;
	padding-top: 11vh;
	border-top-left-radius: 40px;
	border-top-right-radius: 40px;
	padding-bottom: 13vh;
}
.tabbar {
  z-index: 999;
}
.header {
  position: relative;
  text-align: left;
  margin-bottom: 20px;
  font-size: 28px;
  font-weight: bold;
  padding: 0 10px;
}

.title {
  position: relative;
}
.title::after {
  content: "";
  position: absolute;
  bottom: -4px;
  left: 0;
  width: 100%;
  height: 6px;
  background-color: #285AC8;
}
.diseases-list {
  margin-top: 10px;
  margin-bottom: 10px;
  padding: 20px;
  margin-bottom: 30px;
}

.disease-item {
  padding-top: 20px;
  padding-bottom: 17px;
  padding-left: 16px;
  margin-bottom: 10px;
  border-radius: 10px;
  background-color: #FFFFFF;
  cursor: pointer;
  transition: all 0.2s ease-in-out;
}

.disease-item:hover,
.disease-item.hovered {
  background-color: #e0e0e0;
  border-color: #405ac8;
}

.disease-name {
  font-size: 16px;
  color: #333333;
}
.indexNumber{
	color: #285AC8;
	font-size: 20;
	margin-right: 16px;
	font-style: italic;
	font-size: 20px;
}
.examine{
	margin: 6px 0;
	font-size: 12px;
	margin-right: 15px;
	float: right;
	color: #285AC8;
}
.back-image{
	width: 36%;
	position: absolute;
	right: 0;
	top: 6vh;
	z-index: 0;
}
.section-title{
	margin-left: -3px;
	padding-left: 4px;
	border-left: 8px solid #285AC8;
	margin: 10px 0;
	font-weight: 600;
}
.label {
  font-size: 16px;
  color: #333333;
  display: block;
}
.section {
  margin-bottom: 20px;
  padding: 12px;
  border-radius: 8px;
  color: #333333;
}

.media-item {
  margin-bottom: 20px;
  padding: 10px;
}

.media-title {
  font-size: 16px;
  margin-bottom: 10px;
}

.media-image {
  width: 100%;
  height: auto;
}

.media-video {
  width: 100%;
  height: auto;
}
.media-item {
  display: flex;
  flex-direction: column; /* 垂直布局，确保标题和视频换行 */
  margin-bottom: 20px; /* 每个视频项之间的间距 */
}

.media-title {
  font-size: 16px;
  margin-bottom: 5px; /* 标题和视频之间的间距 */
}

.media-video {
  width: 100%; /* 根据需要调整视频宽度 */
}
input{
	opacity: 0.7;
}
.pdf-item {
  display: flex;
  align-items: center;
  background-color: #f5f7fa;
  border-radius: 12px;
  margin: 10px 0;
  padding: 12px;
  transition: box-shadow 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.pdf-item:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.pdf-icon-container {
  width: 50px;
  height: 50px;
  margin-right: 15px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #e0e4e8;
  border-radius: 8px;
}

.pdf-icon {
  width: 30px;
  height: 30px;
}

.pdf-info {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex-grow: 1;
}

.pdf-title {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin-bottom: 8px;
  word-wrap: break-word;
}

.open-pdf-btn {
  background-color: #285ac8;
  color: white;
  padding: 8px 12px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.3s;
}

.open-pdf-btn:hover {
  background-color: #1e4aa1;
}

</style>