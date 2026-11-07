<template>
  <div>
    <!-- 上传组件，控制是否可编辑 -->
    <el-upload
      v-model:file-list="files"
      action="#"
      list-type="picture-card"
      :auto-upload="false"
      :on-preview="handlePictureCardPreview"
      :on-remove="handleRemove"
      :before-upload="handleBeforeUpload"
      :on-change="handleFileChange"
      :disabled="!editable"
      accept=".jpg,.jpeg,.png"
    >
      <el-icon v-if="editable"><Plus /></el-icon>
    </el-upload>
    <div class="file-names">
      <div v-for="file in files" :key="file.uid" class="file-name">
        {{ file.name }}
      </div>
    </div>

    <!-- 预览图片弹窗：限制高度，避免页面出现滚轮 -->
    <el-dialog
      v-model="dialogVisible"
      class="custom-dialog"
      :modal="true"
      width="auto"
      align-center
      destroy-on-close
    >
      <img
        v-if="dialogImageUrl"
        :src="dialogImageUrl"
        alt="图片预览"
        class="dialog-image"
      />
    </el-dialog>
  </div>
</template>

<script lang="js" setup>
import { ref, watch, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { Plus, Delete, ZoomIn } from '@element-plus/icons-vue';

const props = defineProps({
  filePath: {
    type: String,
    required: false,
  },
  editable: {
    type: Boolean,
    default: false,  // 控制是否可编辑
  },
});
const emits = defineEmits(['file-selected']);

const files = ref([]);
const dialogVisible = ref(false);
const dialogImageUrl = ref('');

// 初始化文件
onMounted(() => {
  if (props.filePath) {
    files.value.push({
      name: props.filePath.split('/').pop(),
      url: props.filePath,
    });
  }
});

// 当 filePath 变化时更新文件列表
watch(() => props.filePath, (newFilePath) => {
  if (newFilePath) {
    files.value = [
      {
        name: newFilePath.split('/').pop(),
        url: newFilePath,
      }
    ];
  }
});

// 预览图片
const handlePictureCardPreview = (file) => {
  dialogImageUrl.value = file.url || URL.createObjectURL(file.raw);
  dialogVisible.value = true;
};

// 删除文件
const handleRemove = (file, fileList) => {
  const index = files.value.findIndex(f => f.uid === file.uid);
  if (index !== -1) {
    files.value.splice(index, 1);
    ElMessage.success('图片已删除');
  } else {
    ElMessage.warning('文件未找到，无法删除');
  }
};


// 上传文件限制
const handleBeforeUpload = () => {
  if (!props.editable) {
    ElMessage.warning('请先进入编辑模式');
    return false;
  }
  return true;
};

// 选择文件后将文件对象传递到父组件，同时限制上传数量
const handleFileChange = (file, fileList) => {
  // 检查文件列表中是否已有文件
  if (files.value.length >= 1) {
    ElMessage.warning('该宣传材料只能上传一个文件');
    fileList.pop(); // 移除新添加的文件
    return;
  }

  // 更新文件列表并生成缩略图
  files.value = fileList.map(f => ({
    name: f.name,
    url: URL.createObjectURL(f.raw),
    raw: f.raw,
    uid: f.uid,
  }));

  // 将文件传递给父组件
  if (file.raw) {
    emits('file-selected', file.raw);
  }
};
</script>



<style scoped>
.image-preview {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}

.filestyle {
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  width: 100%;
  height: 100%;
}

.el-upload-list__item-thumbnail_pdf {
  width: 60%;
  height: 60%;
}

.el-upload-list__item-thumbnail_jpg {
  width: 70%;
  height: 70%;
  object-fit: contain;
  border-radius: 5px;
}

.file-names {
  margin-top: 10px;
}

.file-name {
  text-align: left;
  font-size: 14px;
}

.custom-dialog :deep(.el-dialog) {
  margin: 0 auto;
  max-width: 90vw;
  max-height: 90vh;
  overflow: hidden;
}

.custom-dialog :deep(.el-dialog__body) {
  padding: 12px;
  max-height: 80vh;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
}

.dialog-image {
  display: block;
  max-width: min(880px, 86vw);
  max-height: 78vh;
  width: auto;
  height: auto;
  object-fit: contain;
}
</style>
