<template>
  <div class="md-preview_component">
    <md-preview
      class="preview-view"
      v-model="contentData"
      editorId="MdPreview"
      previewTheme="cyanosis"
      codeTheme="github"
      :mdHeadingId="generateId"
    />
    <el-scrollbar class="catalog-view">
      <md-catalog editorId="MdPreview" :mdHeadingId="generateId" />
    </el-scrollbar>
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { MdPreview, MdCatalog } from "md-editor-v3";
import { generateId } from "./config";

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
</script>

<style lang="scss">
.md-preview_component {
  display: flex;
  .preview-view {
    flex: 1;
    display: flex;
    flex-direction: row;
    justify-content: center;
    .md-editor-preview-wrapper {
      padding: 0 20px;
      max-width: 1200px;
    }
  }
  .catalog-view {
    width: 200px;
  }
}
</style>
