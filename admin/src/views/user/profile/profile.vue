<template>
  <div class="profile-page">
    <aside class="profile-nav">
      <div class="nav-user">
        <div class="avatar">{{ nameInitial }}</div>
        <div class="nav-user-meta">
          <div class="name">{{ displayName }}</div>
          <div class="role">{{ role || '—' }}</div>
        </div>
      </div>
      <nav class="nav-list">
        <button
          v-for="item in navItems"
          :key="item.key"
          type="button"
          class="nav-item"
          :class="{ active: activeTab === item.key }"
          @click="activeTab = item.key"
        >
          <span class="nav-dot" />
          {{ item.label }}
        </button>
      </nav>
      <button type="button" class="logout-link" @click="handleLogout">退出登录</button>
    </aside>

    <section class="profile-main">
      <header class="main-head">
        <h1>{{ currentNav.label }}</h1>
        <p>{{ currentNav.desc }}</p>
      </header>

      <div v-show="activeTab === 'info'" class="panel panel-form">
        <el-form
          ref="infoFormRef"
          :model="infoForm"
          :rules="infoRules"
          label-position="top"
          class="info-form"
          @submit.prevent
        >
          <el-form-item label="姓名" prop="name">
            <el-input v-model="infoForm.name" maxlength="32" placeholder="请输入姓名" />
          </el-form-item>
          <el-form-item label="手机号" prop="phoneNumber">
            <el-input
              v-model="infoForm.phoneNumber"
              maxlength="11"
              placeholder="请输入手机号"
            />
          </el-form-item>
          <el-form-item label="角色">
            <el-input :model-value="role || '—'" disabled />
          </el-form-item>
          <el-form-item label="账号状态">
            <el-input model-value="正常" disabled />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :loading="infoSubmitting" @click="submitInfo">
              保存资料
            </el-button>
          </el-form-item>
        </el-form>
      </div>

      <div v-show="activeTab === 'theme'" class="panel">
        <div class="theme-options">
          <button
            type="button"
            class="theme-card"
            :class="{ active: themeMode === 'light' }"
            @click="onThemeChange('light')"
          >
            <div class="theme-preview light">
              <span class="bar" /><span class="block" /><span class="block short" />
            </div>
            <div class="theme-label">亮色</div>
          </button>
          <button
            type="button"
            class="theme-card"
            :class="{ active: themeMode === 'dark' }"
            @click="onThemeChange('dark')"
          >
            <div class="theme-preview dark">
              <span class="bar" /><span class="block" /><span class="block short" />
            </div>
            <div class="theme-label">暗色</div>
          </button>
        </div>
      </div>

      <div v-show="activeTab === 'password'" class="panel panel-form">
        <el-form
          ref="pwdFormRef"
          :model="pwdForm"
          :rules="pwdRules"
          label-position="top"
          class="pwd-form"
          @submit.prevent
        >
          <el-form-item label="原密码" prop="oldPassword">
            <el-input
              v-model="pwdForm.oldPassword"
              type="password"
              show-password
              maxlength="16"
              placeholder="请输入原密码"
              autocomplete="current-password"
            />
          </el-form-item>
          <el-form-item label="新密码" prop="newPassword">
            <el-input
              v-model="pwdForm.newPassword"
              type="password"
              show-password
              maxlength="16"
              placeholder="8-16位，含大小写字母、数字和特殊符号"
              autocomplete="new-password"
            />
          </el-form-item>
          <el-form-item label="确认新密码" prop="confirmPassword">
            <el-input
              v-model="pwdForm.confirmPassword"
              type="password"
              show-password
              maxlength="16"
              placeholder="再次输入新密码"
              autocomplete="new-password"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :loading="submitting" @click="submitPassword">
              保存新密码
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </section>
  </div>
</template>

<script>
import { useUserStore } from '@/store/userrole.js';
import { ElMessage, ElMessageBox } from 'element-plus';
import { changePassword, updateProfile } from '@/api/login.js';
import { useThemeStore } from '@/store/theme.js';

const passwordPattern =
  /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^A-Za-z0-9]).{8,16}$/;

const NAV_ITEMS = [
  { key: 'info', label: '基本信息', desc: '编辑当前账号资料' },
  { key: 'theme', label: '外观设置', desc: '切换亮色 / 暗色模式，设置会保存在本机' },
  { key: 'password', label: '修改密码', desc: '更新登录密码，需符合安全规则' },
];

export default {
  name: 'Profile',
  data() {
    const validateConfirm = (_rule, value, callback) => {
      if (!value) {
        callback(new Error('请再次输入新密码'));
      } else if (value !== this.pwdForm.newPassword) {
        callback(new Error('两次密码不一致'));
      } else {
        callback();
      }
    };
    return {
      name: '',
      phoneNumber: '',
      role: '',
      themeMode: 'light',
      activeTab: 'info',
      navItems: NAV_ITEMS,
      submitting: false,
      infoSubmitting: false,
      infoForm: {
        name: '',
        phoneNumber: '',
      },
      infoRules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
        phoneNumber: [
          { required: true, message: '请输入手机号', trigger: 'blur' },
          {
            pattern: /^1\d{10}$/,
            message: '请输入正确的11位手机号',
            trigger: 'blur',
          },
        ],
      },
      pwdForm: {
        oldPassword: '',
        newPassword: '',
        confirmPassword: '',
      },
      pwdRules: {
        oldPassword: [{ required: true, message: '请输入原密码', trigger: 'blur' }],
        newPassword: [
          { required: true, message: '请输入新密码', trigger: 'blur' },
          {
            validator: (_r, v, cb) => {
              if (!passwordPattern.test(v || '')) {
                cb(new Error('密码必须为8-16位，含大小写字母、数字和特殊符号'));
              } else {
                cb();
              }
            },
            trigger: 'blur',
          },
        ],
        confirmPassword: [{ required: true, validator: validateConfirm, trigger: 'blur' }],
      },
    };
  },
  computed: {
    displayName() {
      return this.infoForm.name || this.name || '未命名用户';
    },
    nameInitial() {
      const n = this.displayName.trim();
      return n ? n.slice(0, 1) : '管';
    },
    currentNav() {
      return this.navItems.find((i) => i.key === this.activeTab) || this.navItems[0];
    },
  },
  methods: {
    loadUser() {
      const store = useUserStore();
      this.name = store.name || '';
      this.phoneNumber = store.phoneNumber || '';
      this.role = store.role || '';
      this.infoForm.name = store.name || '';
      this.infoForm.phoneNumber = store.phoneNumber || '';
      this.themeMode = useThemeStore().mode;
    },
    onThemeChange(mode) {
      const next = mode === 'dark' ? 'dark' : 'light';
      this.themeMode = next;
      useThemeStore().setMode(next);
      ElMessage.success(next === 'dark' ? '已切换暗色模式' : '已切换亮色模式');
    },
    async submitInfo() {
      const form = this.$refs.infoFormRef;
      if (!form) return;
      try {
        await form.validate();
      } catch {
        return;
      }
      this.infoSubmitting = true;
      try {
        const response = await updateProfile({
          name: this.infoForm.name.trim(),
          phoneNumber: this.infoForm.phoneNumber.trim(),
        });
        const { code, msg, message } = response.data || {};
        if (code === 1) {
          const store = useUserStore();
          store.name = this.infoForm.name.trim();
          store.phoneNumber = this.infoForm.phoneNumber.trim();
          this.name = store.name;
          this.phoneNumber = store.phoneNumber;
          try {
            const raw = JSON.parse(localStorage.getItem('user') || '{}');
            raw.name = store.name;
            raw.phoneNumber = store.phoneNumber;
            localStorage.setItem('user', JSON.stringify(raw));
          } catch {
            /* ignore */
          }
          ElMessage.success('资料已保存');
        } else {
          ElMessage.error(message || msg || '保存失败');
        }
      } catch {
        ElMessage.error('网络错误或服务器异常，请稍后再试');
      } finally {
        this.infoSubmitting = false;
      }
    },
    async submitPassword() {
      const form = this.$refs.pwdFormRef;
      if (!form) return;
      try {
        await form.validate();
      } catch {
        return;
      }
      this.submitting = true;
      try {
        const response = await changePassword({
          oldPassword: this.pwdForm.oldPassword,
          newPassword: this.pwdForm.newPassword,
          confirmPassword: this.pwdForm.confirmPassword,
        });
        const { code, msg, message } = response.data || {};
        if (code === 1) {
          ElMessage.success('密码修改成功');
          this.pwdForm.oldPassword = '';
          this.pwdForm.newPassword = '';
          this.pwdForm.confirmPassword = '';
          form.clearValidate();
          const store = useUserStore();
          if (store.mustChangePassword) {
            store.mustChangePassword = false;
          }
        } else {
          ElMessage.error(message || msg || '密码修改失败');
        }
      } catch {
        ElMessage.error('网络错误或服务器异常，请稍后再试');
      } finally {
        this.submitting = false;
      }
    },
    async handleLogout() {
      try {
        await ElMessageBox.confirm('您确定要退出登录吗？', '确认退出', {
          confirmButtonText: '确认',
          cancelButtonText: '取消',
          type: 'warning',
        });
        const userStore = useUserStore();
        userStore.$reset();
        localStorage.removeItem('user');
        ElMessage.success('退出登录成功');
        this.$router.push('/login');
      } catch {
        // cancelled
      }
    },
  },
  mounted() {
    this.loadUser();
    const tab = this.$route.query.tab;
    if (tab === 'password' || tab === 'theme' || tab === 'info') {
      this.activeTab = tab;
    }
  },
};
</script>

<style scoped>
.profile-page {
  min-height: calc(100vh - 140px);
  display: grid;
  grid-template-columns: 220px 1fr;
  gap: 0;
  background: var(--el-bg-color, #fff);
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid var(--el-border-color-light, #ebeef5);
}

.profile-nav {
  display: flex;
  flex-direction: column;
  padding: 20px 14px;
  background: var(--admin-card-bg, #f7f8fa);
  border-right: 1px solid var(--el-border-color-light, #ebeef5);
}

.nav-user {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 4px 8px 18px;
  border-bottom: 1px solid var(--el-border-color-light, #ebeef5);
  margin-bottom: 12px;
}

.avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  background: var(--el-color-primary, #285ac8);
  color: #fff;
  font-size: 18px;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.nav-user-meta .name {
  font-size: 14px;
  font-weight: 600;
  color: var(--el-text-color-primary, #303133);
  line-height: 1.3;
}

.nav-user-meta .role {
  margin-top: 2px;
  font-size: 12px;
  color: var(--el-text-color-secondary, #909399);
}

.nav-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  border: 0;
  background: transparent;
  text-align: left;
  padding: 10px 12px;
  border-radius: 8px;
  font-size: 14px;
  color: var(--el-text-color-regular, #606266);
  cursor: pointer;
  transition: background 0.15s, color 0.15s;
}

.nav-item:hover {
  background: var(--el-fill-color-light, #f0f2f5);
  color: var(--el-text-color-primary, #303133);
}

.nav-item.active {
  background: color-mix(in srgb, var(--el-color-primary, #285ac8) 12%, transparent);
  color: var(--el-color-primary, #285ac8);
  font-weight: 600;
}

.nav-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: currentColor;
  opacity: 0.35;
  flex-shrink: 0;
}

.nav-item.active .nav-dot {
  opacity: 1;
}

.logout-link {
  margin-top: 12px;
  border: 0;
  background: transparent;
  color: var(--el-text-color-secondary, #909399);
  font-size: 13px;
  text-align: left;
  padding: 10px 12px;
  cursor: pointer;
  border-radius: 8px;
}

.logout-link:hover {
  color: #f56c6c;
  background: rgba(245, 108, 108, 0.08);
}

.profile-main {
  padding: 28px 36px 36px;
  background: var(--el-bg-color, #fff);
}

.main-head {
  margin-bottom: 24px;
}

.main-head h1 {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  color: var(--el-text-color-primary, #303133);
}

.main-head p {
  margin: 6px 0 0;
  font-size: 13px;
  color: var(--el-text-color-secondary, #909399);
}

.panel {
  max-width: 480px;
}

.theme-options {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

.theme-card {
  width: 168px;
  border: 1px solid var(--el-border-color, #dcdfe6);
  border-radius: 10px;
  background: var(--el-bg-color, #fff);
  padding: 12px;
  cursor: pointer;
  text-align: left;
  transition: border-color 0.15s, box-shadow 0.15s;
}

.theme-card:hover {
  border-color: var(--el-color-primary-light-5, #a0cfff);
}

.theme-card.active {
  border-color: var(--el-color-primary, #285ac8);
  box-shadow: 0 0 0 1px var(--el-color-primary, #285ac8);
}

.theme-preview {
  height: 88px;
  border-radius: 6px;
  padding: 10px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 10px;
}

.theme-preview.light {
  background: #f2f3f5;
}

.theme-preview.dark {
  background: #1a1a1a;
}

.theme-preview .bar {
  display: block;
  height: 10px;
  width: 42%;
  border-radius: 4px;
  background: #285ac8;
}

.theme-preview .block {
  display: block;
  height: 18px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.9);
}

.theme-preview.dark .block {
  background: #2a2a2a;
}

.theme-preview .block.short {
  width: 68%;
}

.theme-label {
  font-size: 14px;
  font-weight: 500;
  color: var(--el-text-color-primary, #303133);
}

.info-form :deep(.el-form-item__label),
.pwd-form :deep(.el-form-item__label) {
  color: var(--el-text-color-primary, #333);
  font-weight: 500;
}

@media (max-width: 860px) {
  .profile-page {
    grid-template-columns: 1fr;
  }

  .profile-nav {
    border-right: 0;
    border-bottom: 1px solid var(--el-border-color-light, #ebeef5);
  }

  .nav-list {
    flex-direction: row;
    flex-wrap: wrap;
  }

  .nav-item {
    width: auto;
  }

  .profile-main {
    padding: 20px 16px 28px;
  }
}
</style>
