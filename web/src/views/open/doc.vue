<template>
  <md-preview class="page-open-document" :content="content" />
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import MdPreview from "@/components/md-editor/preview";
import { ElMessage } from "element-plus";
import { useRoute } from "vue-router";
import axios from "axios";

const hostUrl = ref(location.origin);
const content = ref("");

onMounted(() => {
  axios
    .get(hostUrl.value + "/api/open/doc/get/" + useRoute().query.id, {
      responseType: "text",
    })
    .then((res) => {
      if (typeof res.data === "object") {
        content.value = JSON.stringify(res.data);
      } else {
        content.value = res.data + "";
      }
    })
    .catch((err) => {
      console.error(err);
      ElMessage.error("加载失败");
    });
});
</script>

<style lang="scss">
.page-open-document {
  height: 100%;
  width: 100%;
}
</style>
