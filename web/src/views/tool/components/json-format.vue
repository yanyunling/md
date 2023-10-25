<template>
  <div class="page-json-format">
    <el-button type="primary" @click="formatClick">格式化</el-button>
    <div class="json-view">
      <codemirror-editor v-model="content" placeholder="请输入json" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { jsonrepair } from "jsonrepair";
import CodemirrorEditor from "@/components/codemirror-editor";
import { ElMessage } from "element-plus";

const content = ref("");

const formatClick = () => {
  try {
    content.value = JSON.stringify(JSON.parse(jsonrepair(content.value)), null, 4);
  } catch (e) {
    ElMessage.warning("json格式有误");
    console.error(e);
  }
};
</script>

<style lang="scss">
.page-json-format {
  display: flex;
  flex-direction: column;
  align-items: center;
  .json-view {
    margin-top: 5px;
    height: 80vh;
    width: 100%;
    border: 1px solid #ddd;
  }
}
</style>
