<template>
  <div>
    <el-divider content-position="center">Bcrypt加密</el-divider>
    <el-form class="page-bcrypt" label-width="100px">
      <el-form-item label="成本">
        <el-select v-model="cost" class="select-view" :disabled="loading">
          <el-option :key="item" v-for="item in costList" :label="item" :value="item"></el-option>
        </el-select>
        <el-button type="primary" :disabled="loading" @click="encryptClick">加密</el-button>
        <el-button type="warning" :disabled="loading" @click="verifyClick">验证</el-button>
      </el-form-item>
      <el-form-item label="待加密文本">
        <el-input v-model="content" type="textarea" :rows="2" resize="none" :readonly="loading" placeholder="待加密文本"></el-input>
      </el-form-item>
      <el-form-item label="密文">
        <el-input v-model="password" type="textarea" :rows="2" resize="none" :readonly="loading" placeholder="加密后的密文"></el-input>
      </el-form-item>
      <el-form-item label="验证结果">
        <el-input v-model="verify" readonly placeholder="待加密文本与密文是否匹配"></el-input>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import bcrypt from "bcryptjs";

const cost = ref(10);
const costList = ref([8, 9, 10, 11, 12, 13, 14, 15, 16]);
const content = ref("");
const password = ref("");
const verify = ref("");
const loading = ref(false);

/**
 * 加密
 */
const encryptClick = () => {
  loading.value = true;
  setTimeout(() => {
    const salt = bcrypt.genSaltSync(cost.value);
    const hash = bcrypt.hashSync(content.value, salt);
    password.value = hash;
    loading.value = false;
  }, 100);
};

/**
 * 验证
 */
const verifyClick = () => {
  loading.value = true;
  setTimeout(() => {
    const pass = bcrypt.compareSync(content.value, password.value);
    if (pass) {
      verify.value = "成功";
    } else {
      verify.value = "失败";
    }
    loading.value = false;
  }, 100);
};
</script>

<style lang="scss">
.page-bcrypt {
  width: calc(100% - 40px);
  margin: 0 20px;
  .select-view {
    width: 100px;
    margin: 1px 0;
  }
  .el-button {
    margin: 1px 10px;
  }
  .el-button + .el-button {
    margin-left: 0;
  }
}
</style>
