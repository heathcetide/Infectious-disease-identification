<script>
import { getPageLearningMaterials } from '../../services/clock';

export default {
  data() {
    return {
      diseaseTypeName: '', // Declare disease type name
      images: [],
      videos: [],
      pdfs: [],
      links: [],
    };
  },
  onLoad(options) {
    // Retrieve the diseaseTypeName parameter from the options
    this.diseaseTypeName = options.diseaseTypeName;
    this.fetchLearningMaterials();
  },
  methods: {
    fetchLearningMaterials() {
      const materialsParams = {
        diseaseTypeName: this.diseaseTypeName,
        pageNo: 1,
        pageSize: 20
      };
      getPageLearningMaterials(materialsParams)
        .then(res => this.getSuccess(res))
        .catch(error => this.getError(error));
    },
    getSuccess(response) {
	  console.log("资料信息",response)
      if (response.code === 1) {
        const records = response.data.records;
        this.images = records.filter(item => item.materialType === '图片').map(item => ({
          title: item.title,
          filePath:item.filePath,
        }));

        this.videos = records.filter(item => item.materialType === '视频').map(item => ({
          title: item.title,
          filePath:item.filePath,
        }));

        this.pdfs = records.filter(item => item.materialType === '文章').map(item => ({
          title: item.title,
          filePath:item.filePath,
        }));

        this.links = records.filter(item => item.materialType === '网页链接').map(item => ({
          title: item.title,
          filePath: item.link,
        }));
      }
    },
    getError(error) {
      console.error(error);
    },
	openPdf(pdfUrl) {
	  // 使用 WebView 打开PDF文件网址
	  uni.showModal({
	    title: '打开PDF',
	    content: '是否在浏览器中打开此PDF?',
	    success: (result) => {
	      if (result.confirm) {
	        // 唤起浏览器打开PDF文件
	        uni.navigateTo({
	          url: `/pages/webview/webview?url=${encodeURIComponent(pdfUrl)}` // 新增一个页面用于WebView加载网址
	        });
	      }
	    }
	  });
	},
    // openPdf(pdfUrl) {
    //   uni.downloadFile({
    //     url: pdfUrl,
    //     success: (res) => {
    //       const filePath = res.tempFilePath;
    //       uni.openDocument({
    //         filePath: filePath,
    //         fileType: 'pdf',
    //         success: () => {
    //           console.log('PDF 打开成功');
    //         },
    //         fail: (error) => {
    //           console.error('打开 PDF 失败', error);
    //         }
    //       });
    //     },
    //     fail: (error) => {
    //       console.error('文件下载失败', error);
    //     }
    //   });
    // },
    openLink(linkUrl) {
      // 跳转到 /pages/webview/webview 页面并传递链接
      uni.navigateTo({
        url: `/pages/webview/webview?url=${encodeURIComponent(linkUrl)}`
      });
    },
    videoErrorCallback(event) {
      console.error('视频播放失败', event);
      uni.showToast({
        title: '视频播放失败，请稍后重试',
        icon: 'none',
        duration: 2000
      });
    }
  }
};
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <view class="content">
      <view class="section-title">
        <text class="label">健康防护资料</text>
      </view>
      <view v-for="(image, index) in images" :key="index" class="media-item">
        <text class="media-title">{{ image.title }}</text>
        <image :src="image.filePath" mode="widthFix" class="media-image"></image>
      </view>
      <view v-for="(pdf, index) in pdfs" :key="index" class="pdf-item">
        <view class="pdf-icon-container">
          <image src="../../static/file-earmark-pdf-fill.svg" class="pdf-icon"></image>
        </view>
        <view class="pdf-info">
          <text class="pdf-title">{{ pdf.title }}</text>
          <button class="open-pdf-btn" @click="openPdf(pdf.filePath)">查看 PDF</button>
        </view>
      </view>
      <view v-for="(link, index) in links" :key="index" class="pdf-item">
        <view class="pdf-icon-container">
          <image src="../../static/link.svg" class="pdf-icon"></image>
        </view>
        <view class="pdf-info">
          <text class="pdf-title">{{ link.title }}</text>
          <button class="open-link-btn" @click="openLink(link.filePath)">查看链接</button>
        </view>
      </view>
      <view class="section-title">
        <text class="label">宣传视频</text>
      </view>
      <view v-for="(video, index) in videos" :key="index" class="media-item">
        <view class="media-title">
          <text>{{ video.title }}</text>
        </view>
        <view class="media-video">
          <video
            id="myVideo"
            :src="video.filePath"
            @error="videoErrorCallback"
            controls
          >
          </video>
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

.open-link-btn {
  background-color: #285ac8;
  color: white;
  padding: 8px 12px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.3s;
}

.open-link-btn:hover {
  background-color: #1e4aa1;
}

</style>