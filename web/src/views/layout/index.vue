<template>
  <div class="layout-view">
    <div class="top-view">
      <div class="top-title"><svg-icon name="md" style="width: 20px; height: 20px; margin: 5px 5px 0 0"></svg-icon><span>云文档</span></div>
      <el-button class="top-button" type="text">退出登录</el-button>
    </div>
    <router-view class="content-view"></router-view>
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
    margin-left: 25px;
    cursor: pointer;
    display: flex;
    align-items: center;
  }
  .top-button {
    margin-right: 50px;
  }
}
.content-view {
  width: 100%;
  height: calc(100% - 60px);
  overflow: auto;
  background-color: #fcfcfc;
}
</style>