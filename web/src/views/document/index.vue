<template>
  <div class="page-document">
    <book @change="bookChange"></book>
    <doc :currentBookId="currentBookId" ref="doc"></doc>
    <md-editor class="editor-view" v-model="content" @onUploadImg="uploadImage" @onSave="saveDoc" />
  </div>
</template>

<script lang="ts" setup>
import { ref, Ref, watch, computed, onMounted, onBeforeUnmount, nextTick } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import MdEditor from "@/components/md-editor";
import { uploadPicture } from "../picture/util";
import Book from "./components/book.vue";
import Doc from "./components/doc.vue";
import DocumentApi from "@/api/document";

const doc = ref<InstanceType<typeof Doc>>();
const hostUrl = ref(location.origin);
const content = ref("");
const originContent = ref("");
const currentBookId = ref("");
const currentDocId = ref("");

/**
 * 文集选择变化
 */
const bookChange = (bookId: string) => {
  currentBookId.value = bookId;
};

/**
 * 文档选择变化
 */
const docChange = (docId: string) => {
  currentDocId.value = docId;
};

/**
 * 上传图片
 * @param files
 * @param callback
 */
const uploadImage = async (files: File[], callback: (urls: string[]) => void) => {
  const pathList: string[] = [];
  for (let file of files) {
    try {
      pathList.push(hostUrl.value + (await uploadPicture(file)));
    } catch (e) {}
  }
  callback(pathList);
};

/**
 * 保存文档
 */
const saveDoc = (text: string) => {
  console.log(text, content);
};
</script>

<style lang="scss">
.page-document {
  display: flex;
  overflow: auto;
  .editor-view {
    height: 100%;
    flex: 1;
    min-width: 900px;
  }
}
</style>
