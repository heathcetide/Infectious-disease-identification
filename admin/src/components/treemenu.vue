<template>
  <el-menu
    :default-active="activeIndex"
    class="custom-menu"
    @select="handleSelect"
    unique-opened
  >
    <!-- 登录后首页：数据大屏置顶 -->
    <el-menu-item
      index="/report/screen"
      v-if="userStore.role === '系统管理员' || userStore.role === '疾控中心工作人员'"
    >
      <img src="../assets/menu_icons/screen.svg" alt="" class="menu-icons" />
      <span>数据大屏</span>
    </el-menu-item>

    <el-sub-menu index="1" v-if="userStore.role === '系统管理员'">
      <template #title>
        <img src="../assets/menu_icons/user.svg" alt="" class="menu-icons" />
        <span>用户管理</span>
      </template>
      <el-menu-item index="/user/alluser">
        <img src="../assets/menu_icons/all-users.svg" alt="" class="menu-icons" />
        <span>所有用户管理</span>
      </el-menu-item>
    </el-sub-menu>

    <!-- 系统管理员和疾控中心工作人员可见 -->
    <el-menu-item index="/work-env" v-if="userStore.role === '系统管理员' || userStore.role === '疾控中心工作人员'">
      <img src="../assets/menu_icons/work-env.svg" alt="" class="menu-icons" />
      <span>工作环境管理</span>
    </el-menu-item>

    <!-- 统计报表：系统管理员和疾控中心工作人员可见 -->
    <el-sub-menu index="/report" v-if="userStore.role === '系统管理员' || userStore.role === '疾控中心工作人员'">
      <template #title>
        <img src="../assets/menu_icons/report.svg" alt="" class="menu-icons" />
        <span>统计报表</span>
      </template>
      <el-menu-item index="/report/check-in">
        <img src="../assets/menu_icons/checkin.svg" alt="" class="menu-icons" />
        <span>职工打卡信息</span>
      </el-menu-item>
      <el-menu-item index="/report/disease-report">
        <img src="../assets/menu_icons/virus.svg" alt="" class="menu-icons" />
        <span>疾病统计报表</span>
      </el-menu-item>
      <el-menu-item index="/report/accuracy">
        <img src="../assets/menu_icons/accuracy.svg" alt="" class="menu-icons" />
        <span>疾病诊断准确率报表</span>
      </el-menu-item>
    </el-sub-menu>

    <!-- 诊断信息管理：系统管理员、专职医护人员、疾控中心工作人员可见 -->
    <el-menu-item index="/diagnosis" v-if="userStore.role === '系统管理员' || userStore.role === '专职医护人员' || userStore.role === '疾控中心工作人员'">
      <img src="../assets/menu_icons/diagnosis.svg" alt="" class="menu-icons" />
      <span>诊断信息管理</span>
    </el-menu-item>

    <!-- 检查信息管理：系统管理员、专职医护人员、疾控中心工作人员可见 -->
    <el-menu-item index="/check" v-if="userStore.role === '系统管理员' || userStore.role === '专职医护人员' || userStore.role === '疾控中心工作人员'">
      <img src="../assets/menu_icons/check.svg" alt="" class="menu-icons" />
      <span>检查信息管理</span>
    </el-menu-item>

    <!-- 系统管理：仅系统管理员可见 -->
    <el-sub-menu index="/system" v-if="userStore.role === '系统管理员'">
      <template #title>
        <img src="../assets/menu_icons/system.svg" alt="" class="menu-icons" />
        <span>系统管理</span>
      </template>
      <el-menu-item index="/system/weight">
        <img src="../assets/menu_icons/weight_management.svg" alt="" class="menu-icons" />
        <span>疾病数据与权重</span>
      </el-menu-item>
      <el-menu-item index="/system/auto-weight">
        <img src="../assets/menu_icons/auto_weight_update.svg" alt="" class="menu-icons" />
        <span>启动自动权重更新</span>
      </el-menu-item>
      <el-menu-item index="/system/expert-AI">
        <img src="../assets/menu_icons/ExpertorAI_diagnosis_switching.svg" alt="" class="menu-icons" />
        <span>专家/AI诊断切换</span>
      </el-menu-item>
      <el-menu-item index="/system/hospital">
        <img src="../assets/menu_icons/hospital.svg" alt="" class="menu-icons" />
        <span>医院基础数据管理</span>
      </el-menu-item>
      <el-menu-item index="/system/accuracy-manage">
        <img src="../assets/menu_icons/accuracy_manage.svg" alt="" class="menu-icons" />
        <span>准确率范围管理</span>
      </el-menu-item>
    </el-sub-menu>

    <!-- 用户反馈：仅系统管理员可见 -->
    <el-menu-item index="/feedback" v-if="userStore.role === '系统管理员'">
      <img src="../assets/menu_icons/user_feedback.svg" alt="" class="menu-icons" />
      <span>用户反馈</span>
    </el-menu-item>

    <!-- 宣传材料管理：系统管理员和疾控中心工作人员可见 -->
    <el-menu-item index="/propaganda" v-if="userStore.role === '系统管理员' || userStore.role === '疾控中心工作人员'">
      <img src="../assets/menu_icons/study.svg" alt="" class="menu-icons" />
      <span>宣传材料管理</span>
    </el-menu-item>

    <!-- 满意度调查：仅系统管理员可见 -->
    <el-menu-item index="/satisfaction" v-if="userStore.role === '系统管理员'">
      <img src="../assets/menu_icons/satisfaction.svg" alt="" class="menu-icons" />
      <span>满意度调查</span>
    </el-menu-item>

    <!-- 个人中心固定在侧栏底部 -->
    <el-menu-item
      index="/user/profile"
      class="profile-menu-item"
      v-if="userStore.role === '系统管理员' || userStore.role === '专职医护人员' || userStore.role === '疾控中心工作人员'"
    >
      <img src="../assets/menu_icons/user.svg" alt="" class="menu-icons" />
      <span>个人中心</span>
    </el-menu-item>
  </el-menu>
</template>

<script>
import { useUserStore } from '@/store/userrole'; // 确保路径正确
import { ref, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useTagsStore } from '@/store/tags';

export default {
  name: 'TreeMenu',
  data() {
    return {
      activeIndex: this.$route.path,
      userStore: useUserStore() // 使用用户 store
    };
  },
  methods: {
    handleSelect(key) {
      if (!key || key === this.activeIndex) {
        return;
      }
      this.activeIndex = key;
      const tag = {
        path: key,
        title: this.getTitleByPath(key)
      };
      this.$tagsStore.addTag(tag);
      this.$router.push(key).catch(() => {});
    },
    getTitleByPath(path) {
      const menuMap = {
        '/user/alluser': '所有用户管理',
        '/user/profile': '个人中心',
        '/work-env': '工作环境管理',
        '/report/check-in': '职工打卡信息',
        '/report/disease-report': '疾病统计报表',
        '/report/screen': '数据大屏',
        '/diagnosis': '诊断信息管理',
        '/check': '检查信息管理',
        '/system/dis-data': '疾病数据与权重',
        '/system/weight': '疾病数据与权重',
        '/system/auto-weight': '启动自动权重更新',
        '/system/expert-AI': '专家/AI诊断切换',
        '/system/hospital': '医院基础数据管理',
        '/feedback': '用户反馈',
        '/propaganda': '宣传材料管理',
        '/satisfaction': '满意度调查',
        '/report/accuracy' : '疾病诊断准确率报表',
        '/system/accuracy-manage' :'准确率范围管理',
      };
      return menuMap[path] || '未知标题';
    }
  },
  created() {
    this.$tagsStore = useTagsStore();
    this.$router = useRouter();
    this.$route = useRoute();

    // 监听路由路径变化以更新活动菜单项
    this.$watch(() => this.$route.path, (newPath) => {
      this.activeIndex = newPath;
    });
  }
};
</script>



<style scoped>
.el-menu {
  border-right: 1px solid var(--el-border-color-light, #ebeef5);
  font-size: 13px;
  width: 180px;
  display: flex;
  flex-direction: column;
  min-height: calc(100vh - 120px);
  box-sizing: border-box;
  background: transparent !important;
  --el-menu-bg-color: transparent;
  --el-menu-hover-bg-color: var(--admin-menu-hover, #f0f2f5);
  --el-menu-text-color: var(--el-text-color-regular, #606266);
  --el-menu-active-color: var(--el-color-primary, #285ac8);
  --el-menu-item-height: 36px;
  --el-menu-base-level-padding: 12px;
}
.profile-menu-item {
  margin-top: auto !important;
  border-top: 1px solid var(--el-border-color-light, #f0f0f0);
}
.menu-icons {
  width: 14px;
  height: 14px;
  margin-right: 6px;
  vertical-align: middle;
  opacity: 0.75;
  transition: opacity 0.2s ease, filter 0.2s ease;
}
.el-menu-item {
  margin: 1px 6px;
  width: calc(100% - 12px);
  border-radius: 6px;
  height: 36px !important;
  line-height: 36px !important;
  font-size: 13px;
}
.el-menu-item.is-active {
  background-color: var(--admin-menu-active-bg, rgba(40, 90, 200, 0.12)) !important;
  color: var(--el-color-primary, #285ac8) !important;
}
.el-menu-item.is-active span {
  color: inherit;
}
.el-menu-item.is-active .menu-icons {
  opacity: 1;
  filter: none;
}
.el-menu-item:not(.is-active):hover {
  background-color: var(--admin-menu-hover, #f0f2f5) !important;
}
.el-menu-item:not(.is-active):hover .menu-icons {
  opacity: 1;
}
:deep(.el-sub-menu__title) {
  margin: 1px 6px;
  width: calc(100% - 12px);
  border-radius: 6px;
  height: 36px !important;
  line-height: 36px !important;
  font-size: 13px;
  padding-right: 32px !important;
  position: relative;
}
:deep(.el-sub-menu__icon-arrow) {
  display: inline-flex !important;
  position: absolute !important;
  right: 12px;
  top: 50%;
  margin-top: -7px;
  width: 14px;
  height: 14px;
  font-size: 14px;
  color: var(--el-text-color-secondary, #909399) !important;
  transition: transform 0.25s ease;
}
:deep(.el-sub-menu.is-opened > .el-sub-menu__title .el-sub-menu__icon-arrow) {
  transform: rotate(180deg);
}
:deep(.el-sub-menu .el-menu-item) {
  padding-left: 40px !important;
  font-size: 12px;
  height: 34px !important;
  line-height: 34px !important;
}
</style>
