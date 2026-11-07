<script>
export default {
  name: 'ReportUpload',
  props: {
    files: {
      type: Array,
      default: () => [],
    },
    limit: {
      type: Number,
      default: 3,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    tip: {
      type: String,
      default: '支持图片或 PDF，最多 3 个',
    },
  },
  emits: ['update:files', 'change'],
  methods: {
    isImage(file) {
      const src = String((file && (file.path || file.url || file.name)) || '')
      const ext = src.split('?')[0].split('.').pop().toLowerCase()
      return ['jpg', 'jpeg', 'png', 'gif', 'webp', 'bmp', 'heic'].includes(ext)
    },
    previewSrc(file) {
      return (file && (file.path || file.url)) || ''
    },
    emitFiles(next) {
      this.$emit('update:files', next)
      this.$emit('change', next)
    },
    removeAt(index) {
      if (this.disabled) return
      const next = this.files.slice()
      next.splice(index, 1)
      this.emitFiles(next)
    },
    previewImage(file) {
      const urls = this.files
        .filter((f) => this.isImage(f))
        .map((f) => this.previewSrc(f))
        .filter(Boolean)
      const current = this.previewSrc(file)
      if (!urls.length || !current) return
      uni.previewImage({ urls, current })
    },
    addFiles() {
      if (this.disabled) return
      const remain = this.limit - this.files.length
      if (remain <= 0) {
        uni.showToast({ title: `最多上传${this.limit}个文件`, icon: 'none' })
        return
      }
      uni.showActionSheet({
        itemList: ['拍照', '从相册选择', '从聊天记录选择'],
        success: (res) => {
          if (res.tapIndex === 0) {
            this.pickImages(['camera'], remain)
          } else if (res.tapIndex === 1) {
            this.pickImages(['album'], remain)
          } else if (res.tapIndex === 2) {
            this.pickMessageFiles(remain)
          }
        },
      })
    },
    pickImages(sourceType, count) {
      uni.chooseImage({
        count,
        sizeType: ['compressed'],
        sourceType,
        success: (res) => {
          const picked = (res.tempFilePaths || []).map((path, i) => ({
            path,
            name: (res.tempFiles && res.tempFiles[i] && res.tempFiles[i].name) || `image_${Date.now()}_${i}.jpg`,
          }))
          this.emitFiles([...this.files, ...picked].slice(0, this.limit))
        },
      })
    },
    pickMessageFiles(count) {
      // #ifdef MP-WEIXIN
      wx.chooseMessageFile({
        count,
        type: 'file',
        extension: ['pdf', 'jpg', 'jpeg', 'png'],
        success: (res) => {
          const picked = (res.tempFiles || []).map((file) => ({
            path: file.path,
            name: file.name || 'file',
          }))
          this.emitFiles([...this.files, ...picked].slice(0, this.limit))
        },
      })
      // #endif
      // #ifndef MP-WEIXIN
      this.pickImages(['album'], count)
      // #endif
    },
  },
}
</script>

<template>
  <view class="upload">
    <view class="grid">
      <view class="item" v-for="(file, index) in files" :key="index">
        <image
          v-if="isImage(file)"
          class="thumb"
          :src="previewSrc(file)"
          mode="aspectFill"
          @tap="previewImage(file)"
        />
        <view v-else class="file-chip" @tap.stop>
          <text class="file-name">{{ file.name || '文件' }}</text>
        </view>
        <view v-if="!disabled" class="remove" @tap.stop="removeAt(index)">×</view>
      </view>
      <view
        v-if="!disabled && files.length < limit"
        class="item add"
        @tap="addFiles"
      >
        <text class="add-plus">＋</text>
        <text class="add-text">上传</text>
      </view>
    </view>
    <text class="hint">{{ tip }}</text>
  </view>
</template>

<style scoped>
.upload {
  width: 100%;
}
.grid {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
.item {
  position: relative;
  width: 96px;
  height: 96px;
  border-radius: 12px;
  overflow: hidden;
  background: #eef2f8;
}
.thumb {
  width: 100%;
  height: 100%;
}
.file-chip {
  width: 100%;
  height: 100%;
  padding: 10px;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #e8eef9;
}
.file-name {
  font-size: 11px;
  color: #285ac8;
  text-align: center;
  word-break: break-all;
  line-height: 1.3;
}
.remove {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.55);
  color: #fff;
  font-size: 14px;
  line-height: 20px;
  text-align: center;
}
.add {
  border: 1px dashed #b7c5df;
  background: #f7f9fc;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.add-plus {
  font-size: 26px;
  color: #285ac8;
  line-height: 1;
}
.add-text {
  margin-top: 4px;
  font-size: 12px;
  color: #718096;
}
.hint {
  display: block;
  margin-top: 8px;
  font-size: 12px;
  color: #94a3b8;
}
</style>
