import { defineComponent } from "vue";
import MdEditorV3 from "md-editor-v3";
import "md-editor-v3/lib/style.css";

// 编辑器配置
MdEditorV3.config({
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
  name: "MdEditor",
  setup() {
    return () => {
      return <MdEditorV3 toolbarsExclude={["github", "fullscreen"]} previewTheme="cyanosis" codeTheme="github" showCodeRowNumber />;
    };
  },
});
