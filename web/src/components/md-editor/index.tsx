import { defineComponent } from "vue";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { generateId } from "./config";

export default defineComponent({
  name: "MdEditor",
  setup() {
    return () => {
      return (
        <MdEditor
          toolbarsExclude={["github", "fullscreen"]}
          previewTheme="cyanosis"
          codeTheme="github"
          showCodeRowNumber
          mdHeadingId={generateId}
        />
      );
    };
  },
});
