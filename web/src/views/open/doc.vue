<template>
  <div class="page-open-document">
    <md-preview v-if="docType === 'md'" class="md-view" :content="content" />
    <open-api v-if="docType === 'openApi'" :content="content" mixUrl></open-api>
    <div v-if="docType === 'error'" class="error-view">文档加载失败</div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import MdPreview from "@/components/md-editor/preview";
import OpenApi from "@/components/open-api/index.vue";
import { useRoute } from "vue-router";
import axios from "axios";
import { host } from "@/config";

const hostUrl = ref("");
const content = ref("");
const docType = ref("");

onMounted(() => {
  hostUrl.value = process.env.NODE_ENV === "production" ? location.origin : host;
  axios
    .get(hostUrl.value + "/api/open/doc/get/" + useRoute().query.id, {
      responseType: "json",
    })
    .then((res) => {
      if (res.data.code === 200) {
        content.value = res.data.data.content;
        docType.value = res.data.data.type;
      } else {
        console.error(res);
        docType.value = "error";
      }
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
    height: 100%;
    width: 100%;
    overflow: auto;
  }
  .error-view {
    font-size: 16px;
    color: #595959;
    user-select: none;
  }
}
@media (max-width: 720px) {
  .page-open-document {
    .md-view {
      .catalog-view {
        display: none;
      }
    }
  }
}
</style>
