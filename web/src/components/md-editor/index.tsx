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
import "./index.scss";
import { lineNumbers } from "@codemirror/view";

// 编辑器配置
MdEditorV3.config({
  editorConfig: {
    renderDelay: 0,
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
  codeMirrorExtensions(_theme, extensions) {
    return [...extensions, lineNumbers()];
  },
});

export default defineComponent({
  name: "MdEditor",
  setup() {
    const generateId = (text: any, level: any, index: any) => `heading-${index}`;
    return () => {
      return (
        <MdEditorV3
          toolbarsExclude={["github", "pageFullscreen", "fullscreen"]}
          previewTheme="cyanosis"
          codeTheme="github"
          showCodeRowNumber
          mdHeadingId={generateId}
        />
      );
    };
  },
});
