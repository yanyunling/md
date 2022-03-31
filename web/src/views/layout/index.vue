<template>
  <div class="layout-view">
    <div class="top-view">
      <div class="top-title">消息服务管理</div>
      <el-button class="top-button" type="text">注销</el-button>
    </div>
    <div class="bottom-view">
      <el-menu class="menu-view">
        <el-menu-item index="app">
          <i class="el-icon-user"></i>
          <template #title>用户管理</template>
        </el-menu-item>
        <el-menu-item index="email">
          <i class="el-icon-message"></i>
          <template #title>邮件记录</template>
        </el-menu-item>
        <el-menu-item index="sms">
          <i class="el-icon-chat-line-square"></i>
          <template #title>短信记录</template>
        </el-menu-item>
        <el-menu-item index="smsSign">
          <i class="el-icon-notebook-1"></i>
          <template #title>短信签名</template>
        </el-menu-item>
        <el-menu-item index="smsTemplate">
          <i class="el-icon-notebook-2"></i>
          <template #title>短信模板</template>
        </el-menu-item>
      </el-menu>
      <router-view class="content-view"></router-view>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, reactive, watch, computed, onMounted, onBeforeUnmount, nextTick } from "vue";
import SvgIcon from "@/components/SvgIcon";
export default defineComponent({
  components: {
    SvgIcon,
  },
  setup(props, ctx) {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.layout-view {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
}

.top-view {
  width: 100%;
  height: 59px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  user-select: none;
  border-bottom: 1px #eee solid;

  .top-title {
    font-size: 20px;
    font-weight: bold;
    margin-left: 23px;
    cursor: pointer;
    display: flex;
    align-items: center;
  }

  .top-button {
    margin-right: 50px;
  }
}

.bottom-view {
  display: flex;
  flex-direction: row;
  height: calc(100% - 60px);
  width: 100%;

  .menu-view {
    height: 100%;
    overflow: auto;
    user-select: none;
  }

  .menu-view:not(.el-menu--collapse) {
    width: 200px;
  }

  .content-view {
    flex: 1;
    height: 100%;
    overflow: auto;
    background-color: #fcfcfc;
  }
}
</style>