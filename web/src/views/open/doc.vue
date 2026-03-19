<template>
  <div class="page-open-document">
    <md-editor v-if="docType === 'md'" class="md-view" v-model="content" preview></md-editor>
    <open-api-preview v-if="docType === 'openApi'" :content="content" mixUrl></open-api-preview>
    <div v-if="docType === 'error'" class="error-view">文档加载失败</div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import MdEditor from "@/components/md-editor";
import OpenApiPreview from "@/components/open-api/index.vue";
import { useRoute } from "vue-router";
import OpenApi from "@/api/open";

const content = ref("");
const docType = ref("");

onMounted(() => {
  OpenApi.getDoc(String(useRoute().query.id))
    .then((res) => {
      content.value = res.data.content;
      docType.value = res.data.type!;
    })
    .catch((err) => {
      console.error(err);
      docType.value = "error";
    });
});
</script>

<style lang="scss">
.page-open-document {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  .md-view {
    border-right: 2px solid #ebedee;
    max-width: 1200px;
  }
  .error-view {
    font-size: 16px;
    color: #595959;
    user-select: none;
  }
}
</style>
