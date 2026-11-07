<template>
  <div class="login-page">
    <header class="top-bar">
      <img class="logo" src="@/assets/banner.svg" alt="" />
      <span class="system-name">突发传染病症状筛查与预警系统</span>
    </header>

    <main class="login-main">
      <div class="login-card">
        <div class="card-head">
          <h1>管理端登录</h1>
          <p>请使用已开通的手机号与密码登录</p>
        </div>

        <form class="form" @submit.prevent="handleSubmit">
          <label class="field">
            <span class="label">手机号</span>
            <input
              v-model="phoneNumber"
              type="text"
              inputmode="numeric"
              maxlength="11"
              autocomplete="username"
              placeholder="请输入11位手机号"
              @blur="validateAndFormatphoneNumber"
            />
            <span v-if="phoneNumberError" class="err">{{ phoneNumberError }}</span>
          </label>

          <label class="field">
            <span class="label">密码</span>
            <div class="pwd-wrap">
              <input
                :type="passwordFieldType"
                v-model="password"
                maxlength="16"
                autocomplete="current-password"
                placeholder="请输入密码"
                @blur="validateAndFormatPassword"
              />
              <button class="eye" type="button" @click="togglePasswordVisibility">
                <img
                  :src="passwordFieldType === 'password' ? passwordClose : passwordOpen"
                  alt=""
                />
              </button>
            </div>
            <span v-if="passwordError" class="err">{{ passwordError }}</span>
          </label>

          <button class="submit" type="submit" :disabled="submitting">
            {{ submitting ? '登录中…' : '登 录' }}
          </button>

          <button class="linkish" type="button" @click="resetPassword">
            忘记密码？重置为手机号后六位
          </button>
        </form>
      </div>
    </main>

    <footer class="foot">突发传染病症状筛查与预警系统 · 管理端</footer>
  </div>
</template>

<script>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { loginUser } from "@/api/login.js";
import passwordOpen from "../../assets/password_close.png";
import passwordClose from "../../assets/password_open.png";
import { useUserStore } from "@/store/userrole";
import { useTagsStore } from "@/store/tags.js";

function homePathForRole(role) {
  // 登录后默认进入数据大屏（医护暂无大屏权限，进检查信息）
  if (role === "系统管理员" || role === "疾控中心工作人员") return "/report/screen";
  if (role === "专职医护人员") return "/check";
  return "/user/profile";
}

export default {
  setup() {
    const router = useRouter();
    const userStore = useUserStore();
    const tagStore = useTagsStore();
    const phoneNumber = ref("");
    const password = ref("");
    const passwordFieldType = ref("password");
    const phoneNumberError = ref("");
    const passwordError = ref("");
    const submitting = ref(false);

    const togglePasswordVisibility = () => {
      passwordFieldType.value =
        passwordFieldType.value === "password" ? "text" : "password";
    };

    const handleSubmit = async () => {
      if (!validateAndFormatphoneNumber() || !validateAndFormatPassword()) {
        return;
      }

      submitting.value = true;
      try {
        const response = await loginUser({
          phoneNumber: phoneNumber.value,
          password: password.value,
        });
        const { code, data, msg } = response.data;
        const roleFromBackend = {
          admin: "系统管理员",
          medical: "专职医护人员",
          cdc: "疾控中心工作人员",
        };

        if (code === 1 || code === 200) {
          const userRole =
            data.userType ||
            (data.user && roleFromBackend[data.user.role]) ||
            data.user?.role;

          if (
            userRole === "系统管理员" ||
            userRole === "专职医护人员" ||
            userRole === "疾控中心工作人员"
          ) {
            const userData = {
              phoneNumber: phoneNumber.value,
              token: data.accessToken || data.token,
              name: data.userName || data.user?.name,
              role: userRole,
            };
            userStore.login(userData);
            localStorage.setItem("user", JSON.stringify(userData));
            tagStore.removeAllTag();

            const mustChange = !!data.mustChangePassword;
            const target = mustChange
              ? { path: "/user/profile", query: { tab: "password" } }
              : homePathForRole(userRole);
            const targetPath = typeof target === "string" ? target : target.path;
            const homeTitle =
              targetPath === "/report/screen"
                ? "数据大屏"
                : targetPath === "/check"
                  ? "检查信息管理"
                  : "个人中心";
            tagStore.addTag({ title: homeTitle, path: targetPath });
            if (mustChange) {
              ElMessage.warning("当前为初始密码，请先修改密码");
            } else {
              ElMessage.success("登录成功");
            }
            router.push(target);
          } else {
            ElMessage.error("无登录权限");
          }
        } else {
          ElMessage.error(msg || "手机号或密码错误");
        }
      } catch (error) {
        const msg =
          error?.response?.data?.msg ||
          error?.response?.data?.message ||
          "网络错误或服务器异常，请稍后再试";
        ElMessage.error(msg);
      } finally {
        submitting.value = false;
      }
    };

    const resetPassword = () => {
      if (!phoneNumber.value || phoneNumber.value.length < 6) {
        ElMessage.warning("请先输入手机号");
        return;
      }
      password.value = phoneNumber.value.slice(-6);
      ElMessage.success("已填入手机号后六位，请登录后尽快修改");
    };

    const validateAndFormatphoneNumber = () => {
      const phonePattern = /^[1-9]\d{10}$/;
      if (!phonePattern.test(phoneNumber.value)) {
        phoneNumberError.value = "请输入正确的11位手机号";
        return false;
      }
      phoneNumberError.value = "";
      return true;
    };

    const validateAndFormatPassword = () => {
      if (!/^.{6,16}$/.test(password.value)) {
        passwordError.value = "密码应为6-16位字符";
        return false;
      }
      passwordError.value = "";
      return true;
    };

    return {
      phoneNumber,
      password,
      passwordFieldType,
      passwordOpen,
      passwordClose,
      phoneNumberError,
      passwordError,
      submitting,
      togglePasswordVisibility,
      handleSubmit,
      resetPassword,
      validateAndFormatphoneNumber,
      validateAndFormatPassword,
    };
  },
};
</script>

<style scoped>
.login-page {
  --ink: #1a2332;
  --muted: #6b7785;
  --line: #d9e1ec;
  --accent: #1f4f9c;
  --bg: #eef2f7;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background:
    linear-gradient(180deg, rgba(238, 242, 247, 0.92), rgba(238, 242, 247, 0.96)),
    url("../images/background_login.jpg") center / cover no-repeat;
  font-family: "PingFang SC", "Microsoft YaHei", "Noto Sans SC", sans-serif;
  color: var(--ink);
}

.top-bar {
  height: 64px;
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 0 28px;
  background: rgba(255, 255, 255, 0.92);
  border-bottom: 1px solid rgba(26, 35, 50, 0.08);
  backdrop-filter: blur(8px);
}

.logo {
  height: 36px;
  width: auto;
}

.system-name {
  font-size: 18px;
  font-weight: 650;
  letter-spacing: 0.02em;
}

.login-main {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
}

.login-card {
  width: min(420px, 100%);
  background: #fff;
  border: 1px solid var(--line);
  border-radius: 12px;
  padding: 32px 28px 24px;
  box-shadow: 0 12px 36px rgba(26, 35, 50, 0.1);
}

.card-head h1 {
  margin: 0;
  font-size: 22px;
  font-weight: 700;
}

.card-head p {
  margin: 8px 0 0;
  color: var(--muted);
  font-size: 13px;
}

.form {
  margin-top: 28px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.label {
  font-size: 13px;
  font-weight: 500;
}

.field input {
  height: 44px;
  border: 1px solid var(--line);
  border-radius: 8px;
  padding: 0 12px;
  font-size: 14px;
  background: #f8fafc;
  outline: none;
  box-sizing: border-box;
  width: 100%;
}

.field input:focus {
  border-color: var(--accent);
  background: #fff;
  box-shadow: 0 0 0 3px rgba(31, 79, 156, 0.12);
}

.pwd-wrap {
  position: relative;
}

.pwd-wrap input {
  padding-right: 42px;
}

.eye {
  position: absolute;
  right: 6px;
  top: 50%;
  transform: translateY(-50%);
  width: 32px;
  height: 32px;
  border: 0;
  background: transparent;
  cursor: pointer;
  padding: 0;
}

.eye img {
  width: 16px;
  height: 16px;
  opacity: 0.6;
}

.err {
  color: #c0392b;
  font-size: 12px;
}

.submit {
  margin-top: 8px;
  height: 44px;
  border: 0;
  border-radius: 8px;
  background: var(--accent);
  color: #fff;
  font-size: 15px;
  font-weight: 600;
  letter-spacing: 0.12em;
  cursor: pointer;
}

.submit:hover:not(:disabled) {
  background: #163a74;
}

.submit:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.linkish {
  border: 0;
  background: transparent;
  color: var(--accent);
  font-size: 13px;
  cursor: pointer;
  padding: 4px 0;
}

.linkish:hover {
  text-decoration: underline;
}

.foot {
  text-align: center;
  padding: 16px;
  color: var(--muted);
  font-size: 12px;
}
</style>
