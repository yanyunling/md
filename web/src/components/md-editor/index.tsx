import { defineComponent } from "vue";
import MdEditorV3 from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import highlight from "highlight.js";
import "highlight.js/styles/github.css";
import prettier from "prettier";
import parserMarkdown from "prettier/parser-markdown";
import cropper from "cropperjs";
import "cropperjs/dist/cropper.css";
import screenfull from "screenfull";
import mermaid from "mermaid";
import katex from "katex";
import "katex/dist/katex.css";

// 编辑器配置
MdEditorV3.config({
  editorConfig: {
    renderDelay: 0,
  },
  markedRenderer(renderer) {
    renderer.link = (href, title, text) => {
      return `<a href="${href}" target="_blank">${text}</a>`;
    };
    renderer.heading = (text, level, raw, slugger, index) => {
      return `<h${level} id="heading-${index}">${text}</h${level}>`;
    };
    return renderer;
  },
  editorExtensions: {
    iconfont: "/static/md-editor/iconfont.js",
    highlight: {
      instance: highlight,
    },
    prettier: {
      prettierInstance: prettier,
      parserMarkdownInstance: parserMarkdown,
    },
    cropper: {
      instance: cropper,
    },
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
    const generateId = (text: any, level: any, index: any) => `heading-${index}`;
    return () => {
      return (
        <MdEditorV3
          toolbarsExclude={["github", "fullscreen"]}
          previewTheme="cyanosis"
          codeTheme="github"
          showCodeRowNumber
          markedHeadingId={generateId}
        />
      );
    };
  },
});
