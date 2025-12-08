import { defineComponent } from "vue";
import { MdPreview, MdCatalog } from "md-editor-v3";
import "md-editor-v3/lib/preview.css";
import { generateId } from "./config";

export default defineComponent({
  name: "MdPreview",
  props: {
    content: {
      type: String,
      default: "",
    },
  },
  setup(props) {
    return () => {
      return (
        <div class="md-preview_component">
          <MdPreview
            class="preview-view"
            modelValue={props.content}
            id="MdPreview"
            previewTheme="cyanosis"
            codeTheme="github"
            mdHeadingId={generateId}
            autoFoldThreshold={100}
          />
          <el-scrollbar class="catalog-view">
            <MdCatalog editorId="MdPreview" mdHeadingId={generateId} />
          </el-scrollbar>
        </div>
      );
    };
  },
});
