<template>
  <div class="login-page">
    <aside class="visual" aria-hidden="true">
      <img class="visual-img" :src="heroImg" alt="" />
      <div class="visual-shade" />
      <div class="visual-copy">
        <p class="brand">症状筛查与预警</p>
        <h2 class="headline">突发传染病管理端</h2>
        <p class="tagline">自主判别 · 现场调查 · 数据决策</p>
      </div>
    </aside>

    <section class="panel">
      <div class="panel-inner">
        <header class="card-head">
          <p class="eyebrow">ADMIN CONSOLE</p>
          <h1>登录</h1>
          <p class="sub">使用已开通的手机号与密码进入系统</p>
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
              placeholder="11 位手机号"
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
                placeholder="6–16 位密码"
                @blur="validateAndFormatPassword"
              />
              <button class="eye" type="button" aria-label="显示或隐藏密码" @click="togglePasswordVisibility">
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
            忘记密码？填入手机号后六位
          </button>
        </form>

        <footer class="foot">突发传染病症状筛查与预警系统</footer>
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
@import url("https://fonts.googleapis.com/css2?family=DM+Sans:ital,opsz,wght@0,9..40,500;0,9..40,600;0,9..40,700;1,9..40,500&family=Noto+Sans+SC:wght@400;500;600;700&display=swap");

.login-page {
  --ink: #122033;
  --muted: #5b6b7c;
  --line: #d7e0ea;
  --accent: #1a5f7a;
  --accent-deep: #0f3d4f;
  --panel: #f4f7fa;
  --field: #ffffff;
  min-height: 100vh;
  height: 100vh;
  display: grid;
  grid-template-columns: minmax(0, 1.15fr) minmax(360px, 0.85fr);
  font-family: "Noto Sans SC", "PingFang SC", "Microsoft YaHei", sans-serif;
  color: var(--ink);
  overflow: hidden;
  background: var(--panel);
}

.visual {
  position: relative;
  min-height: 100vh;
  overflow: hidden;
}

.visual-img {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  transform: scale(1.02);
  animation: drift 18s ease-in-out infinite alternate;
}

.visual-shade {
  position: absolute;
  inset: 0;
  background:
    linear-gradient(180deg, rgba(8, 22, 36, 0.28) 0%, rgba(8, 22, 36, 0.55) 55%, rgba(8, 22, 36, 0.78) 100%),
    linear-gradient(90deg, rgba(8, 22, 36, 0.15) 0%, transparent 40%);
  pointer-events: none;
}

.visual-copy {
  position: absolute;
  left: clamp(28px, 4vw, 56px);
  bottom: clamp(36px, 6vh, 72px);
  z-index: 1;
  max-width: min(420px, 80%);
  color: #f4f8fb;
  animation: rise 0.8s ease-out both;
}

.brand {
  margin: 0 0 10px;
  font-family: "DM Sans", sans-serif;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.18em;
  text-transform: uppercase;
  opacity: 0.85;
}

.headline {
  margin: 0;
  font-size: clamp(28px, 3.2vw, 40px);
  font-weight: 700;
  line-height: 1.2;
  letter-spacing: 0.02em;
}

.tagline {
  margin: 14px 0 0;
  font-size: 14px;
  letter-spacing: 0.08em;
  opacity: 0.78;
}

.panel {
  min-height: 100vh;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--panel);
  border-left: 1px solid rgba(18, 32, 51, 0.06);
  animation: panel-in 0.55s ease-out both;
}

.panel-inner {
  width: min(400px, calc(100% - 64px));
  padding: 8px 0 24px;
}

.card-head .eyebrow {
  margin: 0 0 10px;
  font-family: "DM Sans", sans-serif;
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.2em;
  color: var(--accent);
}

.card-head h1 {
  margin: 0;
  font-size: 32px;
  font-weight: 700;
  letter-spacing: 0.04em;
}

.card-head .sub {
  margin: 10px 0 0;
  color: var(--muted);
  font-size: 14px;
  line-height: 1.5;
}

.form {
  margin-top: 36px;
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.label {
  font-size: 13px;
  font-weight: 600;
  color: #314356;
}

.field input {
  height: 48px;
  border: 1px solid var(--line);
  border-radius: 10px;
  padding: 0 14px;
  font-size: 15px;
  background: var(--field);
  outline: none;
  box-sizing: border-box;
  width: 100%;
  transition: border-color 0.15s ease, box-shadow 0.15s ease, background 0.15s ease;
}

.field input::placeholder {
  color: #93a1b0;
}

.field input:hover {
  border-color: #b7c6d6;
}

.field input:focus {
  border-color: var(--accent);
  background: #fff;
  box-shadow: 0 0 0 4px rgba(26, 95, 122, 0.12);
}

.pwd-wrap {
  position: relative;
}

.pwd-wrap input {
  padding-right: 46px;
}

.eye {
  position: absolute;
  right: 8px;
  top: 50%;
  transform: translateY(-50%);
  width: 34px;
  height: 34px;
  border: 0;
  border-radius: 8px;
  background: transparent;
  cursor: pointer;
  padding: 0;
  display: grid;
  place-items: center;
}

.eye:hover {
  background: rgba(18, 32, 51, 0.05);
}

.eye img {
  width: 16px;
  height: 16px;
  opacity: 0.55;
}

.err {
  color: #b42318;
  font-size: 12px;
}

.submit {
  margin-top: 10px;
  height: 48px;
  border: 0;
  border-radius: 10px;
  background: linear-gradient(180deg, #21708e 0%, var(--accent) 100%);
  color: #fff;
  font-size: 15px;
  font-weight: 650;
  letter-spacing: 0.18em;
  cursor: pointer;
  box-shadow: 0 10px 24px rgba(26, 95, 122, 0.22);
  transition: transform 0.15s ease, box-shadow 0.15s ease, filter 0.15s ease;
}

.submit:hover:not(:disabled) {
  filter: brightness(1.04);
  transform: translateY(-1px);
  box-shadow: 0 14px 28px rgba(26, 95, 122, 0.28);
}

.submit:active:not(:disabled) {
  transform: translateY(0);
}

.submit:disabled {
  opacity: 0.72;
  cursor: not-allowed;
  box-shadow: none;
}

.linkish {
  border: 0;
  background: transparent;
  color: var(--accent);
  font-size: 13px;
  cursor: pointer;
  padding: 6px 0;
  align-self: flex-start;
}

.linkish:hover {
  color: var(--accent-deep);
  text-decoration: underline;
  text-underline-offset: 3px;
}

.foot {
  margin-top: 40px;
  color: #8a98a8;
  font-size: 12px;
  letter-spacing: 0.02em;
}

@keyframes drift {
  from { transform: scale(1.02) translate3d(0, 0, 0); }
  to { transform: scale(1.06) translate3d(-1.2%, -0.8%, 0); }
}

@keyframes rise {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes panel-in {
  from { opacity: 0; transform: translateX(12px); }
  to { opacity: 1; transform: translateX(0); }
}

@media (max-width: 900px) {
  .login-page {
    grid-template-columns: 1fr;
    height: auto;
    min-height: 100vh;
    overflow: auto;
  }

  .visual {
    min-height: 36vh;
    max-height: 42vh;
  }

  .visual-copy {
    bottom: 24px;
    left: 24px;
  }

  .headline {
    font-size: 26px;
  }

  .panel {
    min-height: auto;
    align-items: flex-start;
    padding: 36px 0 48px;
    border-left: 0;
  }
}
</style>
