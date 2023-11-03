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
    <div class="codemirror-view" v-if="docType === 'openApi'">
      <codemirror-editor
        :style="{ visibility: codemirrorVisibility }"
        ref="codemirrorRef"
        v-model="currentDoc.content"
        :disabled="onlyPreview"
        @save="saveDoc(currentDoc.content)"
        @ready="codemirrorReday"
      />
    </div>
    <template v-else>
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
    </template>
  </div>
</template>

<script lang="ts" setup>
import { ref, Ref, onMounted, onBeforeUnmount, nextTick, computed, watch } from "vue";
import MdEditor from "@/components/md-editor";
import MdPreview from "@/components/md-editor/preview";
import CodemirrorEditor from "@/components/codemirror-editor";
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
const codemirrorRef = ref();
const hostUrl = ref(location.origin);
const books: Ref<Book[]> = ref([]);
const currentBookId = ref("");
const currentDoc: Ref<CurrentDoc> = ref({
  id: "",
  content: "",
  originContent: "",
  type: "",
  updateTime: "",
});
const mdloading = ref(false);
const mdKey = ref(0);
const codemirrorVisibility = ref("hidden");

const docType = computed(() => {
  return currentDoc.value.type;
});

watch(docType, (newVal, oldVal) => {
  if (oldVal && oldVal !== newVal && newVal === "openApi") {
    codemirrorVisibility.value = "hidden";
  }
});

onMounted(() => {
  DocCache.getDoc().then((res) => {
    if (res) {
      currentDoc.value = res;
    }
  });
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
const docChange = (id: string, content: string, type: string, updateTime: string, noRender?: boolean) => {
  currentDoc.value.id = id;
  currentDoc.value.content = content;
  currentDoc.value.type = type;
  currentDoc.value.originContent = content;
  currentDoc.value.updateTime = updateTime;
  if (!noRender) {
    nextTick(() => {
      mdKey.value++;
      if (codemirrorRef.value) {
        codemirrorRef.value.$el.getElementsByClassName("cm-scroller")[0].scrollTop = 0;
      }
    });
  }
  DocCache.setDoc(currentDoc.value);
};

/**
 * codemirror加载完成
 */
const codemirrorReday = () => {
  setTimeout(() => {
    if (codemirrorRef.value) {
      codemirrorRef.value.$el.getElementsByClassName("cm-scroller")[0].scrollTop = 0;
      codemirrorVisibility.value = "unset";
    }
  }, 100);
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
    min-width: 720px;
  }
  .editor-view.md-fullscreen {
    min-width: unset;
  }
  .codemirror-view {
    height: 100%;
    flex: 1;
    min-width: 720px;
    overflow: hidden;
  }
}
@media (max-width: 720px) {
  .page-document {
    .editor-view {
      min-width: 100%;
      .catalog-view {
        display: none;
      }
    }
    .codemirror-view {
      min-width: 100%;
    }
  }
}
</style>
