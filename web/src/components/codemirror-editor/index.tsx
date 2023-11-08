import { defineComponent } from "vue";
import { Codemirror } from "vue-codemirror";
import { json } from "@codemirror/lang-json";
import { keymap } from "@codemirror/view";

export default defineComponent({
  name: "CodemirrorEditor",
  setup(_, { emit }) {
    return () => {
      return (
        <Codemirror
          style={{ width: "100%" }}
          autofocus
          indent-with-tab
          tab-size={2}
          extensions={[
            json(),
            keymap.of([
              {
                key: "Ctrl-s",
                run: (view) => {
                  emit("save");
                  return true;
                },
              },
            ]),
          ]}
        />
      );
    };
  },
});
