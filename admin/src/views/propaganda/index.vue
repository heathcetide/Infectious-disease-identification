<template>
  <div class="container">
    <!-- 表单部分 -->
    <el-form
      :model="queryParams"
      ref="queryForm"
      size="default"
      :inline="true"
      v-show="showSearch"
    >
      <el-form-item label="" prop="diseaseTypeName" size="default">
        <el-select
          v-model="queryParams.diseaseTypeName"
          placeholder="请选择疾病类型"
          clearable
          size="default"
          style="width: 200px; margin-right: -15px"
          @clear="handleQuery"
        >
          <el-option
            value="新型冠状病毒感染"
            label="新型冠状病毒感染"
          ></el-option>
          <el-option value="流感" label="流感"></el-option>
          <el-option value="鼠疫" label="鼠疫"></el-option>
          <el-option value="感染性腹泻" label="感染性腹泻"></el-option>
          <el-option value="炭疽" label="炭疽"></el-option>
          <el-option value="结核病" label="结核病"></el-option>
          <el-option
            value="登革热（蚊媒传染病）"
            label="登革热（蚊媒传染病）"
          ></el-option>
          <el-option
            value="疟疾（蚊媒传染病）"
            label="疟疾（蚊媒传染病）"
          ></el-option>
          <el-option
            value="森林脑炎（蜱媒传染病）"
            label="森林脑炎（蜱媒传染病）"
          ></el-option>
          <el-option
            value="发热伴血小板减少综合征（蜱媒传染病）"
            label="发热伴血小板减少综合征（蜱媒传染病）"
          ></el-option>
          <el-option value="斑疹伤寒" label="斑疹伤寒"></el-option>
          <el-option value="流行性出血热" label="流行性出血热"></el-option>
          <el-option value="其他" label="其他"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item>
        <el-button
          type="primary"
          class="custom-button"
          @click="handleQuery"
          plain
          size="default"
          style="margin-left: 5px"
          >搜索</el-button
        >
        <el-button
          type="primary"
          class="custom-button"
          @click="handleAdd"
          size="default"
          style="margin-left: 5px"
          >添加宣传材料</el-button
        >
      </el-form-item>
    </el-form>

    <!-- 表格部分 -->
    <div class="usertable">
      <el-table
        :header-cell-style="tableHeaderStyle"
        v-loading="loading"
        :data="allData"
        class="data-table"
        style="width: 100%"
        :height="tableHeight"
        :show-overflow-tooltip="true"
        :default-sort="{ prop: 'LearningNumber', order: '' }"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="SequenceNumber" label="序号" width="80" />
        <el-table-column prop="MaterialType" label="宣传材料类型" width="130">
          <template #default="scope">
            <el-tag
              size="small"
              effect="light"
              round
              :type="materialTypeTagType(scope.row.MaterialType)"
            >
              {{ scope.row.MaterialType || '未知' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="Title" label="标题" min-width="220" />
        <el-table-column label="材料预览" width="180">
          <template #default="scope">
            <div class="preview-cell">
              <el-image
                v-if="isImageType(scope.row.MaterialType, scope.row.Link)"
                :src="scope.row.Link"
                fit="cover"
                class="preview-thumb"
                :preview-src-list="[scope.row.Link]"
                preview-teleported
              />
              <div
                v-else-if="isVideoType(scope.row.MaterialType, scope.row.Link)"
                class="preview-video"
                @click="openLink(scope.row.Link)"
              >
                <span>视频</span>
                <el-button link type="primary" size="small">预览</el-button>
              </div>
              <div
                v-else-if="isDocType(scope.row.MaterialType, scope.row.Link)"
                class="preview-doc"
                @click="openLink(scope.row.Link)"
              >
                <span>{{ scope.row.MaterialType || '文档' }}</span>
                <el-button link type="primary" size="small">打开</el-button>
              </div>
              <el-button
                v-else
                link
                type="primary"
                @click="openLink(scope.row.Link)"
              >
                打开网页
              </el-button>
            </div>
          </template>
        </el-table-column>

        <el-table-column prop="PublishDate" label="发布日期" width="200" />
        <el-table-column
          prop="LearningNumber"
          label="学习总次数"
          width="140"
          sortable
        />
        <el-table-column fixed="right" label="操作" min-width="260">
          <template #default="scope">
            <el-button
              link
              type="primary"
              size="large"
              @click="handleClick(scope.row.materialId)"
              >查看 / 编辑
            </el-button>
            <el-button
              link
              :type="scope.row.isDeleted ? 'danger' : 'primary'"
              size="large"
              @click="toggleStatus(scope.row)"
            >
              {{ scope.row.isDeleted ? "已禁用" : "已启用" }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页组件 -->
      <pagination
        v-show="total > 0"
        :total="total"
        v-model:current-page="queryParams.pageNum"
        v-model:page-size="queryParams.pageSize"
        @pagination="handlePagination"
      />
      <Checkmaterials ref="Checkmaterials" :form="queryParams" />
      <Addmaterials ref="Addmaterials" @update-materails="handleQuery" />
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import Pagination from "@/components/pagination.vue";
import Checkmaterials from "./components/checkmaterials.vue";
import Addmaterials from "./components/addmaterials.vue";
import {
  pageSelectDiseaseLearningMaterials,
  startOrStopMaterial,
} from "@/api/propaganda/propaganda.js";

export default {
  components: {
    Pagination,
    Checkmaterials,
    Addmaterials,
  },

  data() {
    return {
      queryParams: {
        diseaseTypeName: "",
        pageNum: 1,
        pageSize: 15,
      },
      allData: [],
      loading: false,
      total: 0,
      showSearch: true,
    };
  },

  computed: {
    tableHeight() {
      return window.innerHeight - 300;
    },
    tableHeaderStyle() {
      return {
        height: '40px',
        background: 'var(--el-fill-color-light, #f5f7fa)',
        color: 'var(--el-text-color-primary, #333)',
      };
    },
  },

  methods: {
    materialTypeTagType(t) {
      const s = String(t || '').toLowerCase()
      if (s.includes('图') || s.includes('image') || s.includes('img') || s.includes('png') || s.includes('jpg')) {
        return 'success'
      }
      if (s.includes('视') || s.includes('video') || s.includes('mp4')) {
        return 'warning'
      }
      if (s.includes('pdf') || s.includes('文档') || s.includes('文') || s.includes('doc')) {
        return 'info'
      }
      if (s.includes('链') || s.includes('网页') || s.includes('web') || s.includes('url')) {
        return ''
      }
      return 'primary'
    },
    handleClick(materialId) {
      this.$refs.Checkmaterials.showDrawer(materialId);
    },

    // 获取宣传材料列表
    async handleQuery() {
      this.loading = true;
      try {
        const params = {
          diseaseTypeName: this.queryParams.diseaseTypeName || "",
          pageNo: this.queryParams.pageNum || 1,
          pageSize: this.queryParams.pageSize || 15,
        };

        const response = await pageSelectDiseaseLearningMaterials(params);

        if (response.data.code === 1) {
          const records =
            response.data.data.records || response.data.data.list || [];
          this.allData = records.map((item, index) => ({
            SequenceNumber: item.sequenceNumber || index + 1,
            MaterialType: item.materialType || "未知类型",
            Title: item.title || "",
            Link: item.link || "",
            PublishDate: this.formatPublishDate(item.publishDate),
            LearningNumber: item.studyCount || 0,
            isDeleted: item.isDeleted,
            materialId: item.materialId,
          }));
          this.total = response.data.data.total;
        } else {
          this.$message.error(
            "获取宣传材料数据失败，请重试！" + response.data.message
          );
        }
      } catch (error) {
        console.error("Error fetching data:", error);
        this.$message.error("获取宣传材料数据失败，请重试！");
      } finally {
        this.loading = false;
      }
    },

    handleAdd() {
      this.$refs.Addmaterials.showDrawer();
    },

    // 切换宣传材料状态
    toggleStatus(row) {
      this.$confirm("您确定要切换材料状态吗？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          try {
            const newStatus = !row.isDeleted;
            const response = await startOrStopMaterial(
              row.materialId,
              newStatus
            );

            if (response.data.code === 1) {
              row.isDeleted = newStatus;
              this.handleQuery();
              this.$message.success(
                `宣传材料已${newStatus ? "禁用" : "启用"}！`
              );
            } else {
              this.$message.error("切换状态失败：" + response.data.message);
            }
          } catch (error) {
            console.error("Error toggling status:", error);
            this.$message.error("切换状态失败，请重试！");
          }
        })
        .catch(() => {
          this.$message.info("已取消状态切换");
        });
    },

    handlePagination({ page, limit }) {
      this.queryParams.pageNum = page;
      this.queryParams.pageSize = limit;
      this.handleQuery();
    },

    openLink(link) {
      if (!link) return;
      window.open(link, "_blank");
    },
    isImageType(type, link) {
      const t = type || "";
      const l = (link || "").toLowerCase();
      return t === "图片" || /\.(png|jpe?g|gif|webp|bmp)(\?|$)/i.test(l);
    },
    isVideoType(type, link) {
      const t = type || "";
      const l = (link || "").toLowerCase();
      return t === "视频" || /\.(mp4|mov|webm|m4v)(\?|$)/i.test(l);
    },
    isDocType(type, link) {
      const t = type || "";
      const l = (link || "").toLowerCase();
      return t === "文章" || t === "PDF" || /\.pdf(\?|$)/i.test(l);
    },

    formatPublishDate(v) {
      if (!v) return "无发布日期";
      if (Array.isArray(v) && v.length >= 3) {
        const y = v[0];
        const m = String(v[1]).padStart(2, "0");
        const d = String(v[2]).padStart(2, "0");
        const hh = String(v[3] ?? 0).padStart(2, "0");
        const mm = String(v[4] ?? 0).padStart(2, "0");
        const ss = String(v[5] ?? 0).padStart(2, "0");
        return `${y}-${m}-${d} ${hh}:${mm}:${ss}`;
      }
      return String(v).replace("T", " ").replace(/\.\d+.*$/, "").slice(0, 19);
    },
  },

  mounted() {
    this.handleQuery();
  },
};
</script>

<style scoped>
.custom-button {
  margin: 0 5px;
}

.usertable {
  margin-top: 20px;
}

.el-table th,
.el-table td {
  text-align: center;
}

.container {
  padding: 16px 16px 8px;
  background-color: var(--el-bg-color, #ffffff);
  border-radius: 8px;
  border: 1px solid var(--el-border-color-light, #ebeef5);
}

.data-table {
  --el-table-border-color: var(--el-border-color-lighter, #ebeef5);
}

.preview-video,
.preview-doc {
  color: var(--el-text-color-secondary, #666);
}

.el-form-item {
  margin-bottom: 10px;
}

.preview-cell {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 56px;
}

.preview-thumb {
  width: 72px;
  height: 54px;
  border-radius: 4px;
  overflow: hidden;
  cursor: pointer;
}

.preview-video,
.preview-doc {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
  font-size: 12px;
  color: #666;
  cursor: pointer;
}

:deep(.el-image-viewer__wrapper .el-image-viewer__img) {
  max-height: 85vh !important;
  max-width: 90vw !important;
  width: auto !important;
  height: auto !important;
  object-fit: contain !important;
}
</style>
