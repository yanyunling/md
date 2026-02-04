<template>
  <div class="page-login">
    <text-rain class="text-rain-background"></text-rain>
    <div class="title-view" title="查看公开文档" @click="publishClick">云文档</div>
    <transition name="fade">
      <div class="content-view" v-if="isLogin">
        <form>
          <el-input
            class="input-view"
            v-model.trim="signInData.name"
            size="large"
            clearable
            placeholder="请输入用户名"
            @keyup.enter.native="loginClick"
          ></el-input>
          <el-input
            class="input-view"
            v-model.trim="signInData.password"
            size="large"
            type="password"
            clearable
            placeholder="请输入密码"
            @keyup.enter.native="loginClick"
          ></el-input>
        </form>
        <el-button class="register-button" type="primary" link text size="small" @click="registerClick" :disabled="loading">注册</el-button>
        <el-button class="login-button" size="large" type="primary" @click="loginClick" :disabled="loading">登录</el-button>
      </div>
      <div class="content-view" v-else>
        <form>
          <el-input class="input-view" v-model.trim="signInData.name" size="large" clearable placeholder="请输入用户名"></el-input>
          <el-input class="input-view" v-model.trim="signInData.password" size="large" type="password" clearable placeholder="请输入密码"></el-input>
          <el-input class="input-view" v-model.trim="confirmPassword" size="large" type="password" clearable placeholder="请再次输入密码"></el-input>
        </form>
        <el-button class="register-button" type="primary" link text size="small" @click="registerClick" :disabled="loading">返回登录</el-button>
        <el-button class="login-button" size="large" type="primary" @click="loginClick" :disabled="loading">注册</el-button>
      </div>
    </transition>
    <el-dialog
      :model-value="captchaVisible"
      append-to-body
      destroy-on-close
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :show-close="false"
      width="360"
      center
    >
      <gocaptcha-slide
        v-loading="captchaLoading"
        :config="{ iconSize: 0 }"
        :data="captchaData"
        :events="{
          confirm: captchaConfirm,
        }"
      />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, Ref, onMounted } from "vue";
import Token from "@/store/token";
import TokenApi from "@/api/token";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import TextRain from "@/components/text-rain/index.vue";
import { SlidePoint } from "go-captcha-vue/dist/components/slide/meta/data";

const hostUrl = ref(location.origin);
const router = useRouter();
const loading = ref(false);
const captchaLoading = ref(false);
const captchaVisible = ref(false);
// 登录/注册
const isLogin = ref(true);
// 登录/注册数据
const confirmPassword = ref("");
const signInData: Ref<SignIn> = ref({
  name: "",
  password: "",
  captchaId: "",
  captchaX: 0,
  captchaY: 0,
});
// 验证码数据
const captchaData = ref({
  thumbY: 0,
  thumbWidth: 0,
  thumbHeight: 0,
  image: "",
  thumb: "",
});

onMounted(() => {
  let nameCache = Token.getName();
  if (nameCache) {
    signInData.value.name = nameCache;
  }
});

/**
 * 点击注册切换按钮
 */
const registerClick = () => {
  signInData.value.name = "";
  signInData.value.password = "";
  confirmPassword.value = "";
  isLogin.value = !isLogin.value;
};

/**
 * 点击登录按钮
 */
const loginClick = () => {
  if (!signInData.value.name) {
    ElMessage.warning("请输入用户名");
    return;
  }
  if (!signInData.value.password) {
    ElMessage.warning("请输入密码");
    return;
  }
  if (!isLogin.value) {
    if (signInData.value.password !== confirmPassword.value) {
      ElMessage.warning("两次密码不一致");
      loading.value = false;
      return;
    }
  }
  // 打开验证码窗口
  captchaGet();
};

/**
 * 获取验证码
 */
const captchaGet = () => {
  loading.value = true;
  captchaLoading.value = true;
  TokenApi.captcha(signInData.value.captchaId)
    .then((res) => {
      signInData.value.captchaId = res.data.captchaId;
      captchaData.value.thumbY = res.data.y;
      captchaData.value.thumbWidth = res.data.width;
      captchaData.value.thumbHeight = res.data.width;
      captchaData.value.image = res.data.image;
      captchaData.value.thumb = res.data.tile;
      captchaVisible.value = true;
    })
    .catch(() => {
      loading.value = false;
      captchaVisible.value = false;
    })
    .finally(() => {
      captchaLoading.value = false;
    });
};

/**
 * 验证码滑动完成回调
 * @param point
 * @param reset
 */
const captchaConfirm = (point: SlidePoint, reset: () => void) => {
  signInData.value.captchaX = point.x;
  signInData.value.captchaY = point.y;
  captchaVisible.value = true;
  TokenApi.validateCaptcha(signInData.value)
    .then((res) => {
      // 验证成功
      if (res.data) {
        signIn();
      } else {
        ElMessage.warning("验证失败");
        captchaGet();
      }
      reset();
    })
    .finally(() => {
      captchaVisible.value = false;
    });
};

/**
 * 调用登录/注册接口
 */
const signIn = () => {
  if (isLogin.value) {
    // 登录
    loading.value = true;
    TokenApi.signIn(signInData.value)
      .then((res) => {
        ElMessage.success("登录成功");
        Token.setToken(res.data);
        router.push({ name: "layout" });
      })
      .finally(() => {
        loading.value = false;
      });
  } else {
    // 注册
    loading.value = true;
    TokenApi.signUp(signInData.value)
      .then(() => {
        ElMessage.success("注册成功");
        signInData.value.password = "";
        confirmPassword.value = "";
        isLogin.value = true;
      })
      .finally(() => {
        loading.value = false;
      });
  }
};

/**
 * 点击公开文档
 */
const publishClick = () => {
  let url = hostUrl.value + "/#/open/publish";
  window.open(url, "_blank");
};
</script>

<style lang="scss" scoped>
.page-login {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  background: linear-gradient(0deg, rgba(255, 238, 213, 1) 0%, rgba(148, 210, 233, 1) 70%);
  display: flex;
  flex-direction: column;
  align-items: center;
  user-select: none;
  overflow: auto;
  .title-view {
    margin-top: 10vh;
    font-size: 24px;
    font-weight: bold;
    color: #3f3f3f;
    cursor: pointer;
  }
  .content-view {
    margin-top: 20px;
    width: 300px;
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    background: rgba(255, 255, 255, 0.2);
    padding: 30px 20px;
    .input-view {
      width: 100%;
      margin-top: 10px;
    }
    .register-button {
      margin: 5px 0;
    }
    .login-button {
      width: 100%;
    }
    .register-button + .login-button {
      margin-left: 0;
    }
  }
  .text-rain-background {
    position: absolute;
    z-index: -1;
  }
}
</style>
