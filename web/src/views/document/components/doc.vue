<template>
  <div class="page-doc" v-loading="docLoading">
    <el-popover :visible="addDocVisible" placement="bottom" trigger="click" width="200px">
      <el-input v-model="newDocName" placeholder="请输入文档名称" style="margin-right: 10px"></el-input>
      <div style="display: flex; margin-top: 8px; justify-content: flex-end">
        <el-button @click="addDocCancel" size="small">取消</el-button>
        <el-button @click="addDocSave" type="primary" size="small">确定</el-button>
      </div>
      <template #reference>
        <el-button type="primary" size="large" link :icon="Plus" @click="addDocVisible = true">创建文档</el-button>
      </template>
    </el-popover>
    <el-scrollbar class="scroll-view">
      <div class="item-view" :class="currentDocId === item.id ? 'selected' : ''" v-for="item in docs" :key="item.id" @click="docClick(item)">
        <text-tip :content="item.name"></text-tip>
        <el-dropdown trigger="click" v-if="item.id">
          <el-icon class="setting-button" @click.stop="() => {}"><Tools /></el-icon>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item style="user-select: none" @click="updateDocClick(item)">修改文档</el-dropdown-item>
              <el-dropdown-item style="user-select: none" @click="deleteDocClick(item)">删除文档</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-scrollbar>
  </div>
</template>

<script lang="ts" setup>
import { defineProps, ref, Ref, onMounted, watch } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { Plus, Tools } from "@element-plus/icons-vue";
import TextTip from "@/components/text-tip";
import DocumentApi from "@/api/document";

const docs: Ref<Doc[]> = ref([]);
const docLoading = ref(false);
const currentDocId = ref("");
const addDocVisible = ref(false);
const newDocName = ref("");

const props = defineProps({
  currentBookId: {
    type: String,
    default: "",
  },
});

watch(
  () => props.currentBookId,
  (val) => {
    queryDocs(val);
  }
);

onMounted(() => {
  queryDocs(props.currentBookId);
});

/**
 * 查询文档列表
 */
const queryDocs = (bookId: string) => {
  addDocCancel();
  docLoading.value = true;
  DocumentApi.list(bookId)
    .then((res) => {
      docs.value = res.data;
    })
    .finally(() => {
      docLoading.value = false;
    });
};

/**
 * 点击文档
 */
const docClick = (doc: Doc) => {
  currentDocId.value = doc.id;
  docLoading.value = true;
  DocumentApi.get(currentDocId.value)
    .then((res) => {
      console.log(res);
    })
    .finally(() => {
      docLoading.value = false;
    });
};

/**
 * 点击添加文档保存
 */
const addDocSave = () => {
  let name = String(newDocName.value).trim();
  if (!name) {
    ElMessage.warning("请填写文档名称");
    return;
  }
  docLoading.value = true;
  DocumentApi.add({ id: "", name: name, content: "", bookId: props.currentBookId })
    .then(() => {
      ElMessage.success("创建成功");
      queryDocs(props.currentBookId);
    })
    .catch(() => {
      docLoading.value = false;
    });
};

/**
 * 点击添加文档取消
 */
const addDocCancel = () => {
  addDocVisible.value = false;
  newDocName.value = "";
};

/**
 * 点击修改文档
 */
const updateDocClick = (doc: Doc) => {};

/**
 * 点击删除文档
 */
const deleteDocClick = (doc: Doc) => {
  ElMessageBox.confirm("是否删除文档：" + doc.name + "？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    DocumentApi.delete(doc.id).then(() => {
      ElMessage.success("删除成功");
      if (currentDocId.value === doc.id) {
        currentDocId.value = "";
      }
      queryDocs(props.currentBookId);
    });
  });
};
</script>

<style lang="scss">
.page-doc {
  height: 100%;
  min-width: 300px;
  width: 300px;
  background: #fafafa;
  display: flex;
  flex-direction: column;
  .el-button--large [class*="el-icon"] + span {
    margin-left: 3px;
  }
  .scroll-view {
    color: #595959;
    font-size: 15px;
    .item-view {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 10px;
      cursor: pointer;
      border-left: 3px #fafafa solid;
      transition: 0.1s;
      .update-view {
        display: flex;
        align-items: center;
      }
    }
    .item-view:hover {
      background: #e6e6e6;
      border-color: #e6e6e6;
    }
    .item-view.selected {
      background: #e6e6e6;
      border-color: #0094c1;
    }
    .setting-button {
      margin-left: 10px;
      color: #595959;
    }
    .setting-button:hover {
      color: #777;
    }
  }
  .el-loading-mask {
    background: #fafafa;
  }
}
</style>
