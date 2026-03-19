<template>
  <div class="page-document">
    <book @change="bookChange" @books="booksFetch" :onlyPreview="onlyPreview" :isStretch="isStretch" :loading="mdLoading"></book>
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
      <open-api v-if="onlyPreview" :content="currentDoc.content"></open-api>
      <template v-else>
        <div class="codemirror-toolbar">
          <div class="icon-outer" title="保存" @click="saveDoc(currentDoc.content)">
            <save-icon name="save" class="icon-save" />
          </div>
          <div class="icon-outer" title="导出" @click="exportOpenApi(currentDoc.name, currentDoc.content)">
            <download-icon name="download" class="icon-download" />
          </div>
        </div>
        <div class="codemirror-inner">
          <codemirror-editor
            :style="{ visibility: codemirrorVisibility }"
            ref="codemirrorRef"
            v-model="currentDoc.content"
            :disabled="onlyPreview || mdLoading"
            noRadius
            @save="saveDoc(currentDoc.content)"
            @ready="codemirrorReday"
          />
        </div>
      </template>
    </div>
    <md-editor v-else v-model="currentDoc.content" :preview="onlyPreview" @save="saveDoc" :uploadImage="uploadImage"></md-editor>
  </div>
</template>

<script lang="ts" setup>
import { ref, Ref, onMounted, onBeforeUnmount, nextTick, computed, watch } from "vue";
import MdEditor from "@/components/md-editor";
import CodemirrorEditor from "@/components/codemirror-editor";
import OpenApi from "@/components/open-api/index.vue";
import { uploadPicture } from "../picture/util";
import Book from "./components/book.vue";
import Doc from "./components/doc.vue";
import DocCache from "@/store/doc-cache";
import Token from "@/store/token";
import { host } from "@/config";
import crypto from "crypto-js";
import { exportOpenApi } from "./util";
import saveIcon from "@/icons/save.svg";
import downloadIcon from "@/icons/download.svg";

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
const hostUrl = ref("");
const books: Ref<Book[]> = ref([]);
const currentBookId = ref("");
const currentDoc: Ref<CurrentDoc> = ref({
  id: "",
  name: "",
  content: "",
  originMD5: "",
  type: "",
  updateTime: "",
});
const mdLoading = ref(false);
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
  hostUrl.value = process.env.NODE_ENV === "production" ? location.origin : host;
  DocCache.getDoc().then((res) => {
    if (res) {
      currentDoc.value = res;
    }
  });
});

onBeforeUnmount(() => {
  if (Token.getAccessToken()) {
    DocCache.setDoc(currentDoc.value);
  }
});

window.onbeforeunload = () => {
  if (Token.getAccessToken()) {
    DocCache.setDoc(currentDoc.value);
  }
};

/**
 * 文档loading变化
 */
const loadingChange = (val: boolean) => {
  mdLoading.value = val;
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
const docChange = (id: string, name: string, content: string, type: string, updateTime: string, noRender?: boolean) => {
  currentDoc.value.id = id;
  currentDoc.value.name = name;
  currentDoc.value.content = content;
  currentDoc.value.type = type;
  currentDoc.value.originMD5 = crypto.MD5(content).toString();
  currentDoc.value.updateTime = updateTime;
  if (!noRender) {
    nextTick(() => {
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
const uploadImage = async (file: File, callback: (url: string, options?: object) => void) => {
  try {
    callback(hostUrl.value + (await uploadPicture(file)), { name: file.name });
  } catch (e) {}
};

/**
 * 保存文档
 */
const saveDoc = (content: string) => {
  if (mdLoading.value) {
    return;
  }
  docRef.value?.saveDoc(content);
};
</script>

<style lang="scss">
.page-document {
  display: flex;
  overflow: auto;
  .codemirror-view {
    height: 100%;
    flex: 1;
    min-width: 720px;
    overflow: hidden;
  }
  .codemirror-toolbar {
    height: 34px;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    border: #e6e6e6 1px solid;
    border-bottom: none;
    padding-right: 10px;
    .icon-outer {
      width: 30px;
      height: 24px;
      color: #3f4a54;
      cursor: pointer;
      .icon-save {
        width: 16px;
        height: 16px;
        margin: 4px 7px;
        display: block;
        fill: currentColor;
      }
      .icon-download {
        width: 20px;
        height: 20px;
        margin: 2px 5px;
        display: block;
        fill: currentColor;
      }
    }
    .icon-outer:hover {
      background: #f6f6f6;
    }
  }
  .codemirror-inner {
    height: calc(100% - 35px);
    overflow: hidden;
  }
}
</style>
