<template>
  <div class="document-view">
    <div class="book-view"></div>
    <div class="doc-view"></div>
    <md-editor
      class="editor-view"
      v-model="content"
      :toolbarsExclude="['github', 'fullscreen']"
      previewTheme="github"
      codeTheme="github"
      showCodeRowNumber
      @onUploadImg="uploadImage"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, Ref, watch, computed, onMounted, onBeforeUnmount, nextTick } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import DocumentApi from "@/api/document";
import { uploadPicture } from "../picture/util";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";

// 编辑器配置
MdEditor.config({
  markedRenderer(renderer) {
    renderer.link = (href, title, text) => {
      return `<a href="${href}" title="${title}" target="_blank">${text}</a>`;
    };
    renderer.heading = (text, level, raw) => {
      return `<h${level} id="${raw}">${text}</h${level}>`;
    };
    return renderer;
  },
  editorExtensions: {
    highlight: {
      js: "/static/js/highlight.min.js",
      css: {
        github: {
          light: "/static/css/github.min.css",
          dark: "/static/css/github-dark.min.css",
        },
      },
    },
    prettier: {
      standaloneJs: "/static/js/standalone.js",
      parserMarkdownJs: "/static/js/parser-markdown.js",
    },
    cropper: {
      js: "/static/js/cropper.min.js",
      css: "/static/css/cropper.min.css",
    },
    iconfont: "/static/js/font_2605852_pqekijay2ij.js",
    screenfull: {
      js: "/static/js/screenfull.min.js",
    },
    mermaid: {
      js: "/static/js/mermaid.min.js",
    },
    katex: {
      js: "/static/js/katex.min.js",
      css: "/static/css/katex.min.css",
    },
  },
});

export default defineComponent({
  components: {
    MdEditor,
  },
  setup() {
    const hostUrl = ref(location.origin);
    const content = ref("");

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

    return { content, uploadImage };
  },
});
</script>

<style lang="scss" scoped>
.document-view {
  display: flex;
  overflow: auto;
  .book-view {
    height: 100%;
    min-width: 200px;
    width: 240px;
    background: #404040;
  }
  .doc-view {
    height: 100%;
    min-width: 250px;
    width: 300px;
    background: #fafafa;
  }
  .editor-view {
    height: 100%;
    flex: 1;
    min-width: 900px;
  }
}
</style>
