<template>
  <div class="md-preview_component">
    <md-editor class="preview-view" v-model="content" editorId="MdPreview" preview-only />
    <el-scrollbar class="catalog-view">
      <md-catalog editorId="MdPreview" :markedHeadingId="generateId" />
    </el-scrollbar>
  </div>
</template>

<script lang="ts" setup>
import MdEditor from "@/components/md-editor";
import MdEditorV3 from "md-editor-v3";
const MdCatalog = MdEditorV3.MdCatalog;

defineProps({
  content: {
    type: String,
    default: "",
  },
});

const generateId = (text: any, level: any, index: any) => `heading-${index}`;
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
