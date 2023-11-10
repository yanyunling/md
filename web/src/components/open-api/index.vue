<template>
  <iframe class="open-api_component" ref="openApiRef" :src="hostUrl + '/api.html'"></iframe>
</template>

<script lang="ts" setup>
import { ref, watch, onMounted } from "vue";

const props = defineProps({
  content: {
    type: String,
    default: "",
  },
});

const hostUrl = ref(location.origin);
const openApiRef = ref();

onMounted(() => {
  openApiRef.value.onload = () => {
    postMessage();
  };
});

watch(
  () => props.content,
  (val) => {
    postMessage();
  }
);

const postMessage = () => {
  openApiRef.value?.contentWindow?.postMessage(props.content);
};
</script>

<style lang="scss" scoped>
.open-api_component {
  border: none;
  width: 100%;
  height: 100%;
}
</style>
