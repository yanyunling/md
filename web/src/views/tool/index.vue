<template>
  <div>
    <div class="page-tool">
      <el-card shadow="hover" v-for="item in componentList" :key="item.code" @click="cardClick(item)">{{ item.name }}</el-card>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, Ref } from "vue";
import { useRouter } from "vue-router";

interface ToolComponent {
  code: string;
  name: string;
}

const router = useRouter();

const componentList: Ref<ToolComponent[]> = ref([
  { code: "json-format", name: "JSON格式化" },
  { code: "encryption", name: "文本转换" },
  { code: "bcrypt", name: "Bcrypt加密" },
  { code: "rsa", name: "RSA加密" },
  { code: "generate", name: "随机数生成" },
  { code: "regular", name: "正则表达式" },
  { code: "cron", name: "Cron表达式解析" },
  { code: "docker-convert", name: "docker-compose转换" },
  { code: "qr-code", name: "二维码生成" },
  { code: "image-compression", name: "图片压缩" },
]);

const cardClick = (item: ToolComponent) => {
  router.push({ name: item.code });
};
</script>

<style lang="scss">
.page-tool {
  width: calc(100% - 40px);
  padding: 20px;
  display: grid;
  justify-content: center;
  grid-template-columns: repeat(auto-fill, 300px);
  gap: 40px;
  height: auto !important;
  overflow: auto;
  .el-card {
    user-select: none;
    cursor: pointer;
    width: 100%;
    height: 150px;
    display: flex;
    align-items: center;
    justify-content: center;
    .el-card__body {
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}
</style>
