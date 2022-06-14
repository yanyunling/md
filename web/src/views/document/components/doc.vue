<template>
  <div class="page-doc" v-loading="docLoading">
    <el-popover :visible="addDocVisible" placement="bottom" trigger="click" width="200px">
      <el-input v-model="newDocName" placeholder="请输入文档名称" style="margin-right: 10px"></el-input>
      <div style="display: flex; margin-top: 8px; justify-content: flex-end">
        <el-button @click="addDocCancel" size="small">取消</el-button>
        <el-button @click="addDocSave" type="primary" size="small">确定</el-button>
      </div>
      <template #reference>
        <el-button class="create-button" type="primary" size="large" link :icon="Plus" @click="addDocVisible = true">创建文档</el-button>
      </template>
    </el-popover>
    <el-scrollbar class="scroll-view">
      <div class="item-view" :class="currentDoc.id === item.id ? 'selected' : ''" v-for="item in docs" :key="item.id" @click="docClick(item)">
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
    <el-dialog
      v-model="dialog.visible"
      :title="dialog.isAdd ? '创建文档' : '更新文档信息'"
      width="400px"
      :show-close="false"
      :before-close="dialogClose"
    >
      <el-input v-model="dialog.condition.name" size="large" placeholder="请输入文档名称" style="width: 100%"></el-input>
      <el-select v-model="dialog.condition.bookId" size="large" style="width: 100%; margin-top: 10px">
        <el-option v-for="item in books" :key="item.id" :label="item.name" :value="item.id"></el-option>
      </el-select>
      <template #footer>
        <span class="dialog-footer">
          <el-button :loading="dialog.loading" @click="dialogClose">取消</el-button>
          <el-button type="primary" :loading="dialog.loading" @click="dialogSave">保存</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, Ref, onMounted, watch, PropType } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { Plus, Tools } from "@element-plus/icons-vue";
import TextTip from "@/components/text-tip";
import DocumentApi from "@/api/document";

const docs: Ref<Doc[]> = ref([]);
const docLoading = ref(false);
const addDocVisible = ref(false);
const newDocName = ref("");
const dialog = ref({
  isAdd: true,
  loading: false,
  visible: false,
  condition: {
    id: "",
    name: "",
    content: "",
    bookId: "",
  },
});

const emit = defineEmits(["change"]);

const props = defineProps({
  currentBookId: {
    type: String,
    default: "",
  },
  currentDoc: {
    type: Object as PropType<CurrentDoc>,
    default: {
      id: "",
      content: "",
      originContent: "",
      updateTime: "",
    },
  },
  books: {
    type: Array as PropType<Book[]>,
    default: [],
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
 * 校验文档变化
 */
const checkDocChange = () => {
  return new Promise((resolve, reject) => {
    if (props.currentDoc.content !== props.currentDoc.originContent) {
      ElMessageBox.confirm("文档未保存，是否继续？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          resolve(null);
        })
        .catch(() => {
          reject();
        });
    } else {
      resolve(null);
    }
  });
};

/**
 * 回调文档信息
 */
const emitDoc = (id: string, content: string, updateTime: string) => {
  emit("change", id, content, updateTime);
};

/**
 * 点击文档
 */
const docClick = (doc: Doc) => {
  checkDocChange().then(() => {
    docLoading.value = true;
    DocumentApi.get(doc.id)
      .then((res) => {
        emitDoc(res.data.id, res.data.content, String(res.data.updateTime));
      })
      .finally(() => {
        docLoading.value = false;
      });
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
  checkDocChange().then(() => {
    docLoading.value = true;
    DocumentApi.add({ id: "", name: name, content: "", bookId: props.currentBookId })
      .then((res) => {
        ElMessage.success("创建成功");
        emitDoc(res.data.id, res.data.content, String(res.data.updateTime));
        queryDocs(props.currentBookId);
      })
      .catch(() => {
        docLoading.value = false;
      });
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
const updateDocClick = (doc: Doc) => {
  dialog.value.condition.id = doc.id;
  dialog.value.condition.name = doc.name;
  dialog.value.condition.content = "";
  dialog.value.condition.bookId = doc.bookId;
  dialog.value.isAdd = false;
  dialog.value.visible = true;
};

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
      if (props.currentDoc.id === doc.id) {
        emitDoc("", "", "");
      }
      queryDocs(props.currentBookId);
    });
  });
};

/**
 * 弹窗关闭
 */
const dialogClose = () => {
  if (dialog.value.loading) {
    return;
  }
  dialog.value.condition.id = "";
  dialog.value.condition.name = "";
  dialog.value.condition.content = "";
  dialog.value.condition.bookId = "";
  dialog.value.isAdd = true;
  dialog.value.visible = false;
};

/**
 * 弹窗保存
 */
const dialogSave = () => {
  if (dialog.value.isAdd) {
    // 新增文档
    docLoading.value = true;
    DocumentApi.add(dialog.value.condition)
      .then((res) => {
        ElMessage.success("创建成功");
        emitDoc(res.data.id, res.data.content, String(res.data.updateTime));
        docLoading.value = false;
        dialogClose();
        queryDocs(props.currentBookId);
      })
      .catch(() => {
        docLoading.value = false;
      });
  } else {
    // 更新基本信息
    let name = String(dialog.value.condition.name).trim();
    if (!name) {
      ElMessage.warning("请填写文档名称");
      return;
    }
    dialog.value.condition.name = name;
    dialog.value.loading = true;
    DocumentApi.update(dialog.value.condition)
      .then(() => {
        ElMessage.success("更新成功");
        dialog.value.loading = false;
        dialogClose();
        queryDocs(props.currentBookId);
      })
      .catch(() => {
        dialog.value.loading = false;
      });
  }
};

/**
 * 保存文档
 */
const saveDoc = (content: string) => {
  if (props.currentDoc.id !== "") {
    // 更新文档内容
    docLoading.value = true;
    DocumentApi.updateContent({ id: props.currentDoc.id, name: "", content: content, bookId: "" })
      .then((res) => {
        ElMessage.success("保存成功");
        emitDoc(res.data.id, res.data.content, String(res.data.updateTime));
        queryDocs(props.currentBookId);
      })
      .catch(() => {
        docLoading.value = false;
      });
  } else {
    // 新增
    dialog.value.condition.id = "";
    dialog.value.condition.name = "";
    dialog.value.condition.content = content;
    dialog.value.condition.bookId = "";
    dialog.value.isAdd = true;
    dialog.value.visible = true;
  }
};

defineExpose({ saveDoc });
</script>

<style lang="scss">
.page-doc {
  height: 100%;
  min-width: 300px;
  width: 300px;
  background: #fafafa;
  display: flex;
  flex-direction: column;
  .create-button {
    height: 60px;
    border-bottom: 1px solid #e6e6e6 !important;
  }
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
      padding: 15px;
      cursor: pointer;
      border-left: 3px #fafafa solid;
      transition: 0.1s;
      border-bottom: 1px solid #eaeaea;
      .update-view {
        display: flex;
        align-items: center;
      }
    }
    .item-view:hover {
      background: #e6e6e6;
      border-left-color: #e6e6e6;
    }
    .item-view.selected {
      background: #e6e6e6;
      border-left-color: #0094c1;
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
