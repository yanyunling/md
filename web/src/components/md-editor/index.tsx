import { defineComponent } from "vue";
import { MdEditor, NormalToolbar } from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { generateId } from "./config";
import { Fragment } from "vue";
import DownloadIcon from "@/icons/download.svg";

export default defineComponent({
  name: "MdEditor",
  emits: ["export"],
  setup(_, { emit }) {
    const exportClick = () => {
      emit("export");
    };
    return () => {
      return (
        <MdEditor
          toolbars={[
            "bold",
            "underline",
            "italic",
            "-",
            "title",
            "strikeThrough",
            "sub",
            "sup",
            "quote",
            "unorderedList",
            "orderedList",
            "task",
            "-",
            "codeRow",
            "code",
            "link",
            "image",
            "table",
            "mermaid",
            "katex",
            "-",
            "revoke",
            "next",
            "save",
            0,
            "=",
            "prettier",
            "pageFullscreen",
            "preview",
            "htmlPreview",
            "catalog",
          ]}
          previewTheme="cyanosis"
          codeTheme="github"
          showCodeRowNumber
          mdHeadingId={generateId}
          autoFoldThreshold={100}
          defToolbars={
            <>
              <NormalToolbar
                title="导出"
                onClick={exportClick}
                trigger={
                  <div class="md-editor-icon">
                    <DownloadIcon class="icon-download" />
                  </div>
                }
              ></NormalToolbar>
            </>
          }
        />
      );
    };
  },
});
