import { config } from "md-editor-v3";
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

config({
  editorConfig: {
    renderDelay: 500,
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
  codeMirrorExtensions(_theme: any, extensions: any) {
    return [...extensions, lineNumbers()];
  },
});

export const generateId = (text: any, level: any, index: any) => `heading-${index}`;
