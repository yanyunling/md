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
          initEditor();
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
      if (cherryInstance) {
        cherryInstance.destroy();
        cherryInstance = null;
      }
    });

    /**
     * 初始化编辑器
     */
    const initEditor = () => {
      if (cherryInstance) {
        cherryInstance.destroy();
      }
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
          toolbarRight: ["save", "export", "|", "codeTheme", "theme"],
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
        editor: {
          defaultModel: props.preview ? "previewOnly" : "edit&preview",
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
          enablePreviewerBubble: !props.preview,
        },
        themeSettings: {
          mainTheme: "light",
          codeBlockTheme: "coy",
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

    return () => {
      return <div id="md-editor" class="md-editor"></div>;
    };
  },
});
