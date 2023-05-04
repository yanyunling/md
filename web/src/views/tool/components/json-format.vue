<template>
  <div class="page-json-format">
    <el-input v-model="content" type="textarea" :rows="5" resize="none" placeholder="请输入JSON文本"></el-input>
    <json-viewer class="json-view" :value="jsonContent" copyable sort previewMode expanded theme="light" />
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from "vue";
import { jsonrepair } from "jsonrepair";
import { JsonViewer } from "vue3-json-viewer";
import "vue3-json-viewer/dist/index.css";

const content = ref("");
const jsonContent = ref("");

watch(
  () => content.value,
  (val) => {
    try {
      jsonContent.value = JSON.parse(jsonrepair(val));
    } catch (e) {
      jsonContent.value = val;
      console.error(e);
    }
  }
);
</script>

<style lang="scss">
.page-json-format {
  display: flex;
  flex-direction: column;
  align-items: center;
  .json-view {
    margin-top: 10px;
    width: 100%;
    height: 70vh;
    overflow: auto;
    box-shadow: 0 0 0 1px #dcdfe6 inset;
    border-radius: 6px;
  }
}
</style>
