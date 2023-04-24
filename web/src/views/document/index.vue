<template>
  <div class="page-document">
    <book @change="bookChange" @books="booksFetch" :onlyPreview="onlyPreview" :isStretch="isStretch"></book>
    <doc
      :onlyPreview="onlyPreview"
      :isStretch="isStretch"
      :currentBookId="currentBookId"
      :currentDoc="currentDoc"
      :books="books"
      @change="docChange"
      @loading="loadingChange"
      ref="docRef"
    ></doc>
    <md-preview v-if="onlyPreview" :key="'preview' + mdKey" class="editor-view" :content="currentDoc.content" />
    <md-editor
      v-else
      :key="'editor' + mdKey"
      class="editor-view"
      v-model="currentDoc.content"
      v-loading="mdloading"
      @onUploadImg="uploadImage"
      @onSave="saveDoc"
    />
  </div>
</template>

<script lang="ts" setup>
import { ref, Ref, onMounted, onBeforeUnmount } from "vue";
import MdEditor from "@/components/md-editor";
import MdPreview from "@/components/md-editor/preview.vue";
import { uploadPicture } from "../picture/util";
import Book from "./components/book.vue";
import Doc from "./components/doc.vue";
import DocCache from "@/store/doc-cache";

defineProps({
  onlyPreview: {
    type: Boolean,
    default: true,
  },
  isStretch: {
    type: Boolean,
    default: true,
  },
});

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
const mdloading = ref(false);
const mdKey = ref(0);

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
 * 文档loading变化
 */
const loadingChange = (val: boolean) => {
  mdloading.value = val;
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
const docChange = (id: string, content: string, updateTime: string, noRender: boolean) => {
  currentDoc.value.id = id;
  currentDoc.value.content = content;
  currentDoc.value.originContent = content;
  currentDoc.value.updateTime = updateTime;
  if (!noRender) {
    mdKey.value++;
  }
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
const saveDoc = (content: string) => {
  if (mdloading.value) {
    return;
  }
  docRef.value?.saveDoc(content);
};
</script>

<style lang="scss">
.page-document {
  display: flex;
  overflow: auto;
  .editor-view {
    height: 100%;
    flex: 1;
    min-width: 860px;
  }
  .editor-view.md-fullscreen {
    min-width: unset;
  }
}
@media (max-width: 860px) {
  .page-document {
    .editor-view {
      min-width: 100%;
      .catalog-view {
        display: none;
      }
    }
  }
}
</style>
