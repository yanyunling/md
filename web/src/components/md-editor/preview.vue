<template>
  <div class="md-preview_component">
    <md-editor class="preview-view" v-model="contentData" editorId="MdPreview" preview-only />
    <el-scrollbar class="catalog-view">
      <md-catalog editorId="MdPreview" :markedHeadingId="generateId" />
    </el-scrollbar>
  </div>
</template>

<script lang="ts" setup>
import MdEditor from "@/components/md-editor";
import MdEditorV3 from "md-editor-v3";
import { computed } from "vue";
const MdCatalog = MdEditorV3.MdCatalog;

const props = defineProps({
  content: {
    type: String,
    default: "",
  },
});

const emit = defineEmits(["update:content"]);

const contentData = computed({
  get: () => props.content,
  set: (val) => {
    emit("update:content", val);
  },
});

const generateId = (text: any, level: any, index: any) => `heading-${index}`;
</script>

<style lang="scss">
.md-preview_component {
  display: flex;
  .preview-view {
    flex: 1;
    .md-editor-content {
      display: flex;
      justify-content: center;
      .md-editor-preview-wrapper {
        padding: 0 20px;
        max-width: 1200px;
      }
    }
  }
  .catalog-view {
    width: 200px;
  }
}
</style>
