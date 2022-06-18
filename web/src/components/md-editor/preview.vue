<template>
  <div class="md-preview_component">
    <md-editor class="preview-view" ref="previewRef" v-model="content" editorId="MdPreview" preview-only />
    <el-scrollbar class="catalog-view" ref="catalogRef">
      <md-catalog editorId="MdPreview" />
    </el-scrollbar>
  </div>
</template>

<script lang="ts" setup>
import { ref, nextTick, watch } from "vue";
import MdEditor from "@/components/md-editor";
import MdEditorV3 from "md-editor-v3";
const MdCatalog = MdEditorV3.MdCatalog;

const props = defineProps({
  content: {
    type: String,
    default: "",
  },
});

const previewRef = ref();
const catalogRef = ref();

watch(
  () => props.content,
  () => {
    nextTick(() => {
      const previewContentDom = previewRef.value?.$el.getElementsByClassName("md-preview-wrapper")[0];
      if (previewContentDom) {
        previewContentDom.scrollTop = 0;
      }
      const catalogContentDom = catalogRef.value.$el.getElementsByClassName("el-scrollbar__wrap")[0];
      if (catalogContentDom) {
        catalogContentDom.scrollTop = 0;
      }
    });
  }
);
</script>

<style lang="scss">
.md-preview_component {
  display: flex;
  .preview-view {
    flex: 1;
    .md-preview-wrapper {
      padding: 0 20px;
    }
  }
  .catalog-view {
    width: 200px;
  }
}
</style>
