<script>
import { tip } from '../../utils/tip'
import { getPageLearningMaterials } from '../../services/clock';

export default {
  data() {
    return {
      diseaseTypeName: '',
      loading: false,
      images: [],
      videos: [],
      pdfs: [],
      links: [],
      empty: false,
    };
  },
  onLoad(options) {
    let name = (options && options.diseaseTypeName) || '';
    try {
      name = decodeURIComponent(name);
    } catch (e) {
      /* keep raw */
    }
    this.diseaseTypeName = name;
    this.fetchLearningMaterials();
  },
  methods: {
    mediaUrl(item) {
      let url = item.link || item.filePath || ''
      if (!url) return ''
      // Prefer local infect-files for seed filenames when CDN/OSS unavailable.
      if (/\/infect-files\/learning\//.test(url) || /^https?:\/\/127\.0\.0\.1:8080\//.test(url)) {
        return url
      }
      return url
    },
    fetchLearningMaterials() {
      this.loading = true;
      getPageLearningMaterials({
        diseaseTypeName: this.diseaseTypeName || '',
        pageNo: 1,
        pageSize: 100,
      })
        .then((res) => this.getSuccess(res))
        .catch((error) => this.getError(error))
        .finally(() => {
          this.loading = false;
        });
    },
    normalizeType(t) {
      return String(t || '').trim();
    },
    getSuccess(response) {
      const ok =
        response &&
        (response.code === 1 || response.code === 200 || response.data != null);
      if (!ok) {
        this.empty = true;
        return;
      }
      const records = (response.data && (response.data.list || response.data.records)) || [];
      const mapped = records.map((item) => ({
        title: item.title,
        url: this.mediaUrl(item),
        disease: item.diseaseTypeName,
        type: this.normalizeType(item.materialType),
      }));
      this.images = mapped.filter((item) => item.type === '图片' || item.type === 'image');
      this.videos = mapped.filter((item) => item.type === '视频' || item.type === 'video');
      this.pdfs = mapped.filter(
        (item) => item.type === '文章' || item.type === 'pdf' || item.type === '文档'
      );
      this.links = mapped.filter(
        (item) =>
          item.type === '网页链接' ||
          item.type === '链接' ||
          item.type === 'link' ||
          item.type === ''
      );
      // Fallback: if typed buckets are empty but API returned rows, show as links.
      if (
        !this.images.length &&
        !this.videos.length &&
        !this.pdfs.length &&
        !this.links.length &&
        mapped.length
      ) {
        this.links = mapped.filter((item) => !!item.url);
      }
      this.empty =
        !this.images.length &&
        !this.videos.length &&
        !this.pdfs.length &&
        !this.links.length;
    },
    getError(error) {
      console.error(error);
      this.empty = true;
      tip.fail('资料加载失败');
    },
    openPdfByIndex(e) {
      const index = Number(e.currentTarget.dataset.index);
      const item = this.pdfs[index];
      if (!item || !item.url) return;
      uni.showModal({
        title: '打开资料',
        content: '是否打开该文档？',
        success: (result) => {
          if (result.confirm) {
            uni.navigateTo({
              url: `/pages/webview/webview?url=${encodeURIComponent(item.url)}`,
            });
          }
        },
      });
    },
    openLinkByIndex(e) {
      const index = Number(e.currentTarget.dataset.index);
      const item = this.links[index];
      if (!item || !item.url) return;
      uni.navigateTo({
        url: `/pages/webview/webview?url=${encodeURIComponent(item.url)}`,
      });
    },
    previewImageByIndex(e) {
      const index = Number(e.currentTarget.dataset.index);
      const item = this.images[index];
      if (!item || !item.url) return;
      uni.previewImage({ urls: [item.url], current: item.url });
    },
    videoErrorCallback() {
      tip.fail('视频播放失败');
    },
  },
};
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image" />
    <view class="content">
      <view class="page-head">
        <text class="page-title">{{ diseaseTypeName || '全部宣传资料' }}</text>
        <text class="page-sub">图文海报、短视频、宣教文稿与网页链接</text>
      </view>

      <view v-if="loading" class="state-tip">加载中…</view>
      <view v-else-if="empty" class="state-tip">暂无宣传资料</view>

      <block v-else>
        <view v-if="images.length" class="section-title">
          <text class="label">图文海报</text>
        </view>

        <view v-for="(image, index) in images" :key="'img-' + index" class="media-card">
          <view class="card-head">
            <image class="type-icon" src="/static/icon-image.svg" mode="aspectFit" />
            <text class="card-title">{{ image.title }}</text>
          </view>
          <text v-if="!diseaseTypeName && image.disease" class="card-tag">{{ image.disease }}</text>
          <image
            :src="image.url"
            mode="widthFix"
            class="media-image"
            :data-index="index"
            @tap="previewImageByIndex"
          />
        </view>

        <view v-if="pdfs.length || links.length" class="section-title">
          <text class="label">文稿与链接</text>
        </view>

        <view
          v-for="(pdf, index) in pdfs"
          :key="'pdf-' + index"
          class="link-card"
          :data-index="index"
          @tap="openPdfByIndex"
        >
          <view class="icon-box pdf">
            <image class="icon-img" src="/static/file-earmark-pdf-fill.svg" mode="aspectFit" />
          </view>
          <view class="link-body">
            <text class="link-title">{{ pdf.title }}</text>
            <text v-if="!diseaseTypeName && pdf.disease" class="card-tag">{{ pdf.disease }}</text>
            <text class="card-type">文章 / PDF</text>
          </view>
          <text class="link-action">查看</text>
        </view>

        <view
          v-for="(link, index) in links"
          :key="'link-' + index"
          class="link-card"
          :data-index="index"
          @tap="openLinkByIndex"
        >
          <view class="icon-box link">
            <image class="icon-img" src="/static/link.svg" mode="aspectFit" />
          </view>
          <view class="link-body">
            <text class="link-title">{{ link.title }}</text>
            <text v-if="!diseaseTypeName && link.disease" class="card-tag">{{ link.disease }}</text>
            <text class="card-type">网页链接</text>
          </view>
          <text class="link-action">打开</text>
        </view>

        <view v-if="videos.length" class="section-title">
          <text class="label">宣传视频</text>
        </view>
        <view v-for="(video, index) in videos" :key="'vid-' + index" class="media-card">
          <view class="card-head">
            <image class="type-icon" src="/static/icon-video.svg" mode="aspectFit" />
            <text class="card-title">{{ video.title }}</text>
          </view>
          <text v-if="!diseaseTypeName && video.disease" class="card-tag">{{ video.disease }}</text>
          <video
            class="media-video"
            :src="video.url"
            controls
            object-fit="contain"
            @error="videoErrorCallback"
          />
        </view>
      </block>
    </view>
    <uni-custom-tabbar class="tabbar" />
  </view>
</template>

<style scoped>
.container {
  width: 100%;
  min-height: 100vh;
  position: relative;
  background: #f5f7fb;
}
.background-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 0;
}
.content {
  position: relative;
  z-index: 1;
  padding: 11vh 14px 14vh;
  box-sizing: border-box;
}
.page-head {
  margin-bottom: 14px;
}
.page-title {
  display: block;
  font-size: 18px;
  font-weight: 600;
  color: #1f2a44;
}
.page-sub {
  display: block;
  margin-top: 4px;
  font-size: 12px;
  color: #8a94a6;
}
.section-title {
  margin: 16px 0 10px;
  padding-left: 8px;
  border-left: 4px solid #285ac8;
  font-weight: 600;
}
.label {
  font-size: 15px;
  color: #333;
}
.state-tip {
  text-align: center;
  color: #999;
  padding: 40px 0;
  font-size: 14px;
}
.media-card {
  background: #fff;
  border-radius: 12px;
  padding: 12px;
  margin-bottom: 12px;
  box-shadow: 0 2px 8px rgba(31, 42, 68, 0.04);
}
.card-head {
  display: flex;
  align-items: flex-start;
  gap: 8px;
  margin-bottom: 6px;
}
.type-icon {
  width: 18px;
  height: 18px;
  margin-top: 2px;
  flex-shrink: 0;
}
.card-title {
  display: block;
  flex: 1;
  font-size: 14px;
  font-weight: 600;
  color: #222;
  line-height: 1.4;
}
.card-tag {
  display: inline-block;
  font-size: 11px;
  color: #285ac8;
  background: #eef3ff;
  padding: 2px 8px;
  border-radius: 999px;
  margin-bottom: 8px;
}
.media-image {
  width: 100%;
  border-radius: 8px;
  display: block;
}
.media-video {
  width: 100%;
  height: 200px;
  border-radius: 8px;
  background: #000;
}
.link-card {
  display: flex;
  align-items: center;
  gap: 10px;
  background: #fff;
  border-radius: 12px;
  padding: 12px;
  margin-bottom: 10px;
  box-shadow: 0 2px 8px rgba(31, 42, 68, 0.04);
}
.icon-box {
  width: 42px;
  height: 42px;
  border-radius: 12px;
  background: #e8eefc;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
.icon-box.link {
  background: #e7f7f2;
}
.icon-box.pdf {
  background: #fde8e8;
}
.icon-img {
  width: 20px;
  height: 20px;
}
.card-type {
  display: block;
  margin-top: 4px;
  font-size: 11px;
  color: #94a3b8;
}
.link-body {
  flex: 1;
  min-width: 0;
}
.link-title {
  display: block;
  font-size: 14px;
  color: #222;
  line-height: 1.4;
  font-weight: 500;
}
.link-action {
  flex-shrink: 0;
  font-size: 13px;
  color: #285ac8;
  padding: 6px 10px;
  background: #eef3ff;
  border-radius: 8px;
}
.tabbar {
  z-index: 999;
}
</style>
