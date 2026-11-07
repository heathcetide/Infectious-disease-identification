<template>
  <div class="container">
    <div class="tags-wrap">
      <div class="tags-container">
        <el-tag
          v-for="tag in tagsList"
          :key="tag.path"
          :closable="true"
          :class="{ active: tag.path === currentPath }"
          @close="removeTag(tag.path)"
          @click="selectTag(tag)"
        >
          {{ tag.title }}
        </el-tag>
      </div>
    </div>
    <div class="actions">
      <el-button link type="primary" class="action-btn" @click="goProfile">
        个人中心
      </el-button>
      <el-button class="logout-btn" @click="handleLogout">登出</el-button>
    </div>
  </div>
</template>

<script>
import { mapState } from "pinia";
import { useTagsStore } from "@/store/tags";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/userrole.js";
import { ElMessageBox, ElMessage } from "element-plus";

export default {
  name: "NavHeader",

  computed: {
    ...mapState(useTagsStore, ["tagsList"]),
    currentPath() {
      return this.$router.currentRoute.value.path;
    },
  },

  methods: {
    goProfile() {
      this.$store.addTag({ title: "个人中心", path: "/user/profile" });
      this.$router.push("/user/profile");
    },
    selectTag(tag) {
      this.$router.push(tag.path);
    },

    removeTag(tagPath) {
      const currentTagIndex = this.tagsList.findIndex(
        (tag) => tag.path === tagPath
      );

      this.$store.removeTag(tagPath);

      if (tagPath === this.currentPath) {
        const previousTag = this.$store.getPreviousTag(currentTagIndex);
        if (previousTag) {
          this.$router.push(previousTag.path);
        } else {
          this.$router.push("/report/screen");
        }
      }
    },

    async handleLogout() {
      const userStore = useUserStore();

      ElMessageBox.confirm("您确定要退出登录吗？", "确认退出", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          userStore.$reset();
          localStorage.removeItem("user");
          ElMessage({
            type: "success",
            message: "退出登录成功",
          });
          this.$router.push("/login");
        })
        .catch(() => {
          ElMessage({
            type: "info",
            message: "已取消登出操作",
          });
        });
    },
  },

  created() {
    this.$store = useTagsStore();
    this.$router = useRouter();
  },
};
</script>

<style scoped>
.container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  background: var(--admin-header-bg, #ffffff);
  height: 100%;
  padding: 0 16px 0 8px;
  box-sizing: border-box;
}

.tags-wrap {
  flex: 1;
  min-width: 0;
  overflow: hidden;
}

.tags-container {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 6px 0;
  overflow-x: auto;
  cursor: pointer;
}

.actions {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  gap: 4px;
}

.action-btn {
  font-size: 13px;
  padding: 0 8px;
}

.logout-btn {
  font-size: 12px;
}

.el-tag {
  color: var(--el-text-color-secondary, #a6a6a6);
  border: 1px solid var(--el-border-color, #a6a6a6);
  background-color: var(--admin-card-bg, #fafafa);
  height: 30px;
  line-height: 30px;
  padding: 0 12px;
}

.el-tag.active {
  color: var(--el-color-primary, #285ac8);
  border: 1px solid var(--el-color-primary, #285ac8);
  background-color: var(--admin-card-bg, #fafafa);
}
</style>
