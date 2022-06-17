import { defineComponent } from "vue";
import MdEditorV3 from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import screenfull from "screenfull";
import katex from "katex";
import "katex/dist/katex.css";
import cropper from "cropperjs";
import "cropperjs/dist/cropper.css";
import highlight from "highlight.js";
import "highlight.js/styles/github.css";
import mermaid from "mermaid";

// 编辑器配置
MdEditorV3.config({
  editorConfig: {
    renderDelay: 0,
  },
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
      instance: highlight,
    },
    prettier: {
      standaloneJs: "/static/md-editor/standalone.js",
      parserMarkdownJs: "/static/md-editor/parser-markdown.js",
    },
    cropper: {
      instance: cropper,
    },
    iconfont: "/static/md-editor/iconfont.js",
    screenfull: {
      instance: screenfull,
    },
    mermaid: {
      instance: mermaid,
    },
    katex: {
      instance: katex,
    },
  },
});

export default defineComponent({
  name: "MdEditor",
  setup() {
    return () => {
      return <MdEditorV3 toolbarsExclude={["github", "fullscreen"]} previewTheme="cyanosis" codeTheme="github" showCodeRowNumber />;
    };
  },
});
