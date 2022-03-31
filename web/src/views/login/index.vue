<template>
  <div class="content-view">
    <div class="content-title">消息服务管理</div>
    <div class="login-view">
      <el-input class="login-input" type="password" placeholder="请输入口令换取验证码"></el-input>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, reactive, watch, computed, onMounted, onBeforeUnmount, nextTick } from "vue";
export default defineComponent({
  setup(props, ctx) {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.content-view {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  background-color: #f5f5f5;
  display: flex;
  flex-direction: column;
  align-items: center;
  user-select: none;

  .content-title {
    margin-top: 10vh;
    font-size: 24px;
    font-weight: bold;
  }

  .login-view {
    width: 350px;
    margin-top: 50px;

    .login-input {
      width: 100%;
      margin-top: 10px;
    }

    .login-button {
      width: 100%;
      margin-top: 20px;
    }
  }
}
</style>