import { defineComponent, onMounted, onBeforeUnmount, watch, PropType } from "vue";
import "cherry-markdown/dist/cherry-markdown.css";
import "./index.scss";
import Cherry from "cherry-markdown";
import type { CherryFileUploadHandler } from "cherry-markdown/types/cherry";

export default defineComponent({
  name: "MdEditor",
  props: {
    modelValue: {
      type: String,
      default: "",
    },
    preview: {
      type: Boolean,
      default: false,
    },
    uploadImage: {
      type: Function as PropType<CherryFileUploadHandler>,
      default: undefined,
    },
  },
  emits: ["save", "update:modelValue"],
  setup(props, { emit }) {
    let cherryInstance: Cherry = null;

    watch(
      () => props.modelValue,
      () => {
        if (cherryInstance && cherryInstance.getValue() !== props.modelValue) {
          cherryInstance.setValue(props.modelValue);
          cherryInstance.previewer.scrollToTop(0);
          cherryInstance.refreshPreviewer(true);
        }
      },
    );

    watch(
      () => props.preview,
      () => {
        initEditor();
      },
    );

    onMounted(() => {
      initEditor();
    });

    onBeforeUnmount(() => {
      destroyEditor();
    });

    /**
     * 初始化编辑器
     */
    const initEditor = () => {
      destroyEditor();
      cherryInstance = new Cherry({
        id: "md-editor",
        value: props.modelValue,
        toolbars: {
          toc: {
            updateLocationHash: false,
            defaultModel: "pure",
            position: "absolute",
          },
          toolbar: [
            "bold",
            "italic",
            "underline",
            "strikethrough",
            "sub",
            "sup",
            "size",
            "color",
            "|",
            "header",
            "list",
            "panel",
            {
              insert: ["image", "link", "quote", "table", "code", "hr", "br", "ruby", "formula", "toc", "detail"],
            },
            "graph",
            "|",
            "undo",
            "redo",
            "shortcutKey",
            "togglePreview",
            "fullScreen",
          ],
          toolbarRight: ["save" as any, "export", "|", "codeTheme", "theme"],
          bubble: ["bold", "italic", "underline", "strikethrough", "sub", "sup", "size", "color", "|", "quote", "link"],
          float: ["h1", "h2", "h3", "|", "ol", "ul", "table", "code"],
          customMenu: {
            save: saveButton,
          },
          shortcutKeySettings: {
            isReplace: false,
            shortcutKeyMap: {
              "Control-KeyS": {
                hookName: "save",
                aliasName: "保存",
              },
            },
          },
        },
        callback: {
          afterChange: (text: string) => {
            emit("update:modelValue", text);
          },
          fileUpload: props.uploadImage,
        },
        engine: {
          syntax: {
            link: {
              target: "_blank",
            },
            autoLink: {
              target: "_blank",
              enableShortLink: false,
            },
            table: {
              enableChart: false,
            },
            header: {
              anchorStyle: "none",
            },
            mathBlock: {
              engine: "katex",
              src: "/static/katex/katex.min.js",
              css: "/static/katex/katex.min.css",
            },
            inlineMath: {
              engine: "katex",
            },
          },
        },
        previewer: {
          enablePreviewerBubble: true,
        },
        themeSettings: {
          themeList: [
            { className: "default", label: "默认" },
            { className: "light", label: "明亮" },
            { className: "dark", label: "暗黑" },
            { className: "abyss", label: "深海" },
            { className: "green", label: "清新" },
            { className: "red", label: "热情" },
            { className: "violet", label: "淡雅" },
            { className: "blue", label: "清幽" },
          ],
          mainTheme: "default",
          codeBlockTheme: "coy",
          inlineCodeTheme: "red",
        },
        isPreviewOnly: props.preview,
      });
    };

    /**
     * 自定义保存按钮
     */
    const saveButton = Cherry.createMenuHook("保存", {
      iconName: "",
      onClick: () => {
        emit("save", cherryInstance.getValue());
      },
    });

    /**
     * 销毁编辑器
     */
    const destroyEditor = () => {
      if (cherryInstance) {
        cherryInstance.destroy();
        cherryInstance = null;
      }
    };

    return () => {
      return <div id="md-editor" class="md-editor"></div>;
    };
  },
});
