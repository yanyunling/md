<template>
  <div class="page-document">
    <book @change="bookChange" @books="booksFetch"></book>
    <doc :currentBookId="currentBookId" :currentDoc="currentDoc" :books="books" @change="docChange" ref="docRef"></doc>
    <md-editor class="editor-view" v-model="currentDoc.content" @onUploadImg="uploadImage" @onSave="saveDoc" />
  </div>
</template>

<script lang="ts" setup>
import { ref, Ref, watch, computed, onMounted, onBeforeUnmount, nextTick } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import MdEditor from "@/components/md-editor";
import { uploadPicture } from "../picture/util";
import Book from "./components/book.vue";
import Doc from "./components/doc.vue";
import DocCache from "@/store/doc-cache";

const docRef = ref<InstanceType<typeof Doc>>();
const hostUrl = ref(location.origin);
const books: Ref<Book[]> = ref([]);
const currentBookId = ref("");
const currentDoc: Ref<CurrentDoc> = ref({
  id: "",
  content: "",
  originContent: "",
  updateTime: "",
});

onMounted(() => {
  currentDoc.value = DocCache.getDoc();
});

onBeforeUnmount(() => {
  DocCache.setDoc(currentDoc.value);
});

window.onbeforeunload = () => {
  DocCache.setDoc(currentDoc.value);
};

/**
 * 文集选择变化
 */
const bookChange = (bookId: string) => {
  currentBookId.value = bookId;
};

/**
 * 文集列表变化
 */
const booksFetch = (bookList: Book[]) => {
  books.value = bookList;
};

/**
 * 文档选择变化
 */
const docChange = (id: string, content: string, updateTime: string) => {
  currentDoc.value.id = id;
  currentDoc.value.content = content;
  currentDoc.value.originContent = content;
  currentDoc.value.updateTime = updateTime;
};

/**
 * 上传图片
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
  console.log(text, currentDoc.value);
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
