<template>
  <div class="page-tool">
    <el-divider content-position="center">JSON格式化</el-divider>
    <json-format class="tool-card" />
    <el-divider content-position="center">随机数、UUID、当前时间生成</el-divider>
    <generate class="tool-card" />
    <el-divider content-position="center">文本转换、编码、解码、加密</el-divider>
    <encryption class="tool-card" />
    <el-divider content-position="center">图片压缩</el-divider>
    <image-compression class="tool-card" />
  </div>
</template>

<script lang="ts" setup>
import Encryption from "./components/encryption.vue";
import Generate from "./components/generate.vue";
import JsonFormat from "./components/json-format.vue";
import ImageCompression from "./components/image-compression.vue";
</script>

<style lang="scss">
.page-tool {
  .el-divider {
    width: 96%;
    margin: 30px 2%;
    user-select: none;
    .el-divider__text {
      font-size: 16px;
      font-weight: bold;
      color: #666;
    }
  }
  .tool-card {
    width: 96%;
    margin: 0 2%;
    overflow: hidden;
  }
  .tool-card + .el-divider {
    margin-top: 100px;
  }
}
</style>
