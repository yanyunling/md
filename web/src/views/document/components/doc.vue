<template>
  <div class="page-doc" v-loading="docLoading"></div>
</template>

<script lang="ts">
import { defineComponent, ref, Ref, onMounted, watch } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { Plus, Tools, CircleCheckFilled, CircleCloseFilled } from "@element-plus/icons-vue";
import TextTip from "@/components/text-tip";
import DocumentApi from "@/api/document";
export default defineComponent({
  components: {
    TextTip,
    Tools,
    CircleCheckFilled,
    CircleCloseFilled,
  },
  emits: ["change"],
  setup(props, { emit }) {
    const docLoading = ref(false);

    return { docLoading };
  },
});
</script>

<style lang="scss">
.page-doc {
  height: 100%;
  min-width: 300px;
  width: 300px;
  background: #fafafa;
}
</style>
