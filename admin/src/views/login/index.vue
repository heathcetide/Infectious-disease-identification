<template>
  <div class="login-page">
    <aside class="visual" aria-hidden="true">
      <img class="visual-img" :src="heroImg" alt="" />
      <div class="visual-shade" />
      <p class="visual-caption">突发传染病症状筛查与预警系统</p>
    </aside>

    <section class="panel">
      <div class="panel-inner">
        <header class="card-head">
          <h1>管理端登录</h1>
          <p class="sub">请输入已开通账号的手机号与密码</p>
        </header>

        <form class="form" @submit.prevent="handleSubmit">
          <label class="field">
            <span class="label">手机号</span>
            <input
              v-model="phoneNumber"
              type="text"
              inputmode="numeric"
              maxlength="11"
              autocomplete="username"
              placeholder="请输入手机号"
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
              <button
                class="eye"
                type="button"
                aria-label="显示或隐藏密码"
                @click="togglePasswordVisibility"
              >
                <img
                  :src="passwordFieldType === 'password' ? passwordClose : passwordOpen"
                  alt=""
                />
              </button>
            </div>
            <span v-if="passwordError" class="err">{{ passwordError }}</span>
          </label>

          <button class="submit" type="submit" :disabled="submitting">
            {{ submitting ? '登录中…' : '登录' }}
          </button>

          <button class="linkish" type="button" @click="resetPassword">
            忘记密码？使用手机号后六位
          </button>
        </form>
      </div>
    </section>
  </div>
</template>

<script>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { loginUser } from "@/api/login.js";
import passwordOpen from "../../assets/password_close.png";
import passwordClose from "../../assets/password_open.png";
import heroImg from "../images/login_hero.png";
import { useUserStore } from "@/store/userrole";
import { useTagsStore } from "@/store/tags.js";

function homePathForRole(role) {
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
      heroImg,
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
  --ink: #1f2a37;
  --muted: #667085;
  --line: #d0d5dd;
  --accent: #175cd3;
  --panel: #ffffff;
  --field-bg: #ffffff;
  min-height: 100vh;
  height: 100vh;
  display: grid;
  grid-template-columns: minmax(0, 1.2fr) minmax(0, 420px);
  font-family: "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Noto Sans SC", sans-serif;
  color: var(--ink);
  overflow: hidden;
  background: var(--panel);
}

.visual {
  position: relative;
  min-height: 100vh;
  background: #e8eef4;
}

.visual-img {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
}

.visual-shade {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(15, 23, 42, 0.08) 0%, rgba(15, 23, 42, 0.45) 100%);
  pointer-events: none;
}

.visual-caption {
  position: absolute;
  left: 28px;
  bottom: 28px;
  z-index: 1;
  margin: 0;
  max-width: 70%;
  color: #fff;
  font-size: 14px;
  font-weight: 500;
  letter-spacing: 0.02em;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.35);
}

.panel {
  min-height: 100vh;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--panel);
  border-left: 1px solid #eaecf0;
}

.panel-inner {
  width: min(340px, calc(100% - 56px));
}

.card-head h1 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  line-height: 1.3;
}

.card-head .sub {
  margin: 8px 0 0;
  color: var(--muted);
  font-size: 13px;
  line-height: 1.5;
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
  gap: 6px;
}

.label {
  font-size: 13px;
  font-weight: 500;
  color: #344054;
}

.field input {
  height: 40px;
  border: 1px solid var(--line);
  border-radius: 6px;
  padding: 0 12px;
  font-size: 14px;
  background: var(--field-bg);
  outline: none;
  box-sizing: border-box;
  width: 100%;
}

.field input::placeholder {
  color: #98a2b3;
}

.field input:focus {
  border-color: var(--accent);
  box-shadow: 0 0 0 3px rgba(23, 92, 211, 0.12);
}

.pwd-wrap {
  position: relative;
}

.pwd-wrap input {
  padding-right: 40px;
}

.eye {
  position: absolute;
  right: 4px;
  top: 50%;
  transform: translateY(-50%);
  width: 32px;
  height: 32px;
  border: 0;
  background: transparent;
  cursor: pointer;
  padding: 0;
  display: grid;
  place-items: center;
}

.eye img {
  width: 15px;
  height: 15px;
  opacity: 0.5;
}

.err {
  color: #d92d20;
  font-size: 12px;
}

.submit {
  margin-top: 4px;
  height: 40px;
  border: 0;
  border-radius: 6px;
  background: var(--accent);
  color: #fff;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
}

.submit:hover:not(:disabled) {
  background: #1570ef;
}

.submit:disabled {
  opacity: 0.65;
  cursor: not-allowed;
}

.linkish {
  border: 0;
  background: transparent;
  color: var(--accent);
  font-size: 13px;
  cursor: pointer;
  padding: 2px 0;
  align-self: flex-start;
}

.linkish:hover {
  text-decoration: underline;
  text-underline-offset: 2px;
}

@media (max-width: 900px) {
  .login-page {
    grid-template-columns: 1fr;
    height: auto;
    min-height: 100dvh;
    overflow-x: hidden;
    overflow-y: auto;
  }

  .visual {
    min-height: 36vw;
    max-height: none;
    height: min(42vw, 220px);
  }

  .visual-caption {
    left: 16px;
    right: 16px;
    bottom: 14px;
    max-width: none;
    font-size: 13px;
  }

  .panel {
    min-height: auto;
    align-items: flex-start;
    padding: 28px 0 calc(40px + env(safe-area-inset-bottom, 0px));
    border-left: 0;
  }

  .panel-inner {
    width: min(360px, calc(100% - 40px));
  }

  .card-head h1 {
    font-size: 22px;
  }

  .form {
    margin-top: 22px;
  }

  .field input,
  .submit {
    height: 44px;
    font-size: 16px; /* avoid iOS zoom on focus */
  }
}

@media (max-width: 480px) {
  .visual {
    height: 160px;
  }

  .panel {
    padding-top: 22px;
  }

  .panel-inner {
    width: calc(100% - 32px);
  }

  .visual-caption {
    font-size: 12px;
  }
}
</style>
