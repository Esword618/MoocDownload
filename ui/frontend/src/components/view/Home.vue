<template>
  <h1 cless="title" style="color: black">主页</h1>
  <el-row>
    <el-col :span="1"></el-col>
    <!-- 搜索框 -->
    <el-col :span="10">
      <div>
        <el-input
          clearable
          size="default"
          v-model="inputUrl"
          placeholder="请把链接粘贴此处"
          class="input-with-select"
        >
          <template #prepend>
            <el-button :icon="Search" />
          </template>
          <template #append>
            <el-select
              v-model="selectPlatform"
              placeholder="Select"
              style="width: 115px"
            >
              <el-option label="mooc" value="1" />
              <el-option label="xtzx" value="2" />
              <el-option label="bilibili" value="3" />
            </el-select>
          </template>
        </el-input>
      </div>
      <div>
        <el-button size="large" @click="searchFunc()" round>查询</el-button>
      </div>
    </el-col>
    <el-col :span="1"></el-col>
    <!-- 选择框 -->
    <el-col :span="10" class="tree-box">
      <el-scrollbar :height="treeHeightStr">
        <el-tree :data="data" :props="defaultProps" show-checkbox />
        <el-tree
          class="tree-box"
          :data="data"
          :props="defaultProps"
          show-checkbox
        />
        <el-tree
          class="tree-box"
          :data="data"
          :props="defaultProps"
          default-expand-all="true"
          show-checkbox
        />
      </el-scrollbar>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { Search } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { onMounted, ref, reactive } from "vue";
const inputUrl = ref("");
const selectPlatform = ref("");
let databool = ref(false);
const data = reactive([{}]);
const defaultProps = {
  children: "children",
  label: "label",
};
let treeHeight = 600;
let treeHeightStr = "600px";
const data1 = [
  {
    id: 1,
    label: "第一章",
    children: [
      {
        id: 4,
        label: "Level two 1-1",
        children: [
          {
            id: 9,
            label: "Level three 1-1-1",
          },
          {
            id: 10,
            label: "Level three 1-1-2",
          },
        ],
      },
    ],
  },
  {
    id: 2,
    label: "Level one 2",
    children: [
      {
        id: 5,
        label: "Level two 2-1",
      },
      {
        id: 6,
        label: "Level two 2-2",
      },
    ],
  },
  {
    id: 3,
    label: "Level one 3",
    children: [
      {
        id: 7,
        label: "Level two 3-1",
      },
      {
        id: 8,
        label: "Level two 3-2",
      },
    ],
  },
  {
    id: 4,
    label: "Level one 3",
    children: [
      {
        id: 9,
        label: "Level two 3-1",
      },
      {
        id: 10,
        label: "Level two 3-2",
      },
    ],
  },
];

const searchFunc = () => {
  if (inputUrl.value == "") {
    ElMessage({
      message: "链接不能为空!",
      type: "warning",
    });
    return;
  }
  if (selectPlatform.value == "") {
    ElMessage({
      message: "平台不能为空!",
      type: "warning",
    });
    return;
  }
  ElMessage({
    message: "查询成功",
    type: "success",
  });

  Object.assign(data, data1);

  console.log(data);
};

onMounted(() => {
  window.addEventListener("resize", () => {
    // 监听页面尺寸变化
    treeHeight = (document.documentElement.clientHeight * 600) / 768;
    treeHeightStr = treeHeight.toString() + "px";
  });
});
</script>

<style>
.tree-box {
  background-color: rgb(249, 249, 249);
}
</style>
