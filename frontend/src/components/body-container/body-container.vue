<template>
  <div class="container">
    <div class="search-container">
      <el-card shadow="hover">
        <el-input
          size="large"
          v-model="link"
          placeholder="请输入链接:"
          class="input-with-select"
        >
          <template #prepend>
            <el-select
              size="large"
              v-model="select"
              placeholder="选择平台"
              style="width: 115px"
            >
              <div v-for="(item, index) in platformList" :key="index">
                <el-option :label="item.label" :value="item.value" />
              </div>
            </el-select>
          </template>
          <template #append>
            <el-button size="large" :icon="Search" @click="onClickSearch" />
          </template>
        </el-input>
      </el-card>
    </div>
    <el-divider border-style="dashed" />
    <div class="info-container">
      <div class="custom-tree-container">
        <el-card
          shadow="always"
          style="height: 440px"
          v-loading="loading"
          :element-loading-svg="svg"
          element-loading-text="Loading..."
          element-loading-svg-view-box="-10, -10, 50, 50"
        >
          <el-scrollbar height="370px">
            <el-tree-v2
              :height="355"
              :props="props"
              :data="dataSource"
              show-checkbox
              node-key="id"
              accordion
              :expand-on-click-node="false"
              @check="handleNodeClick"
              ref="trees"
            >
              <template #default="{ node }">
                <span class="custom-tree-node">
                  <span>{{ node.label }}</span>
                </span>
              </template>
            </el-tree-v2>
          </el-scrollbar>
          <div class="download-btn">
            <el-row :gutter="20">
              <el-col :offset="20" v-if="downloadBtn">
                <el-button type="primary" @click="onDownload">
                  <el-icon class="el-icon--left"><Download /></el-icon>
                </el-button>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </div>
    </div>
    <div class="setting-download-container">
      <!-- 设置 -->
      <div class="setting-container">
        <el-button type="primary" @click="settingDrawer = true">
          <el-icon><Setting /></el-icon>
        </el-button>
        <el-drawer
          class="setting-drawer"
          v-model="settingDrawer"
          :show-close="false"
          size="40%"
        >
          <template #header="{ titleId, titleClass }">
            <h4 :id="titleId" :class="titleClass">设置</h4>
            <el-button type="primary" @click="onClickSave" plain>
              <el-icon>
                <svg
                  t="1677836313342"
                  class="icon"
                  viewBox="0 0 1024 1024"
                  version="1.1"
                  xmlns="http://www.w3.org/2000/svg"
                  p-id="3006"
                  width="200"
                  height="200"
                >
                  <path
                    d="M563.93554 122.611368a48.609619 48.609619 0 0 0-47.970018 49.313179v30.636852a48.609619 48.609619 0 0 0 47.970018 49.313179c26.47945 0 47.970019-22.130169 47.970019-49.313179v-30.636852a48.609619 48.609619 0 0 0-47.970019-49.313179z"
                    fill="#030000"
                    p-id="3007"
                  ></path>
                  <path
                    d="M991.700187 277.266708c0-2.046721-0.89544-3.837601-1.15128-5.820362a48.929419 48.929419 0 0 0-13.623485-40.998376l-215.737165-215.673204a49.377139 49.377139 0 0 0-36.073454-14.454966c-0.51168 0-0.89544-0.25584-1.34316-0.25584h-38.248095L685.331668 0l-0.44772 0.06396H339.116052L338.924172 0H81.612992l-0.70356 0.12792L80.269831 0a46.498938 46.498938 0 0 0-30.444972 12.024485c-0.9594 0.89544-2.302561 1.343161-3.198001 2.302561-1.08732 1.02336-1.599001 2.494441-2.494441 3.645721a45.859338 45.859338 0 0 0-11.832604 29.997252l0.12792 0.6396-0.12792 0.70356v924.734041l0.12792 0.831481-0.12792 0.57564c0 12.024485 4.924922 22.641849 12.344284 30.956652 0.70356 0.83148 1.15128 1.854841 1.854841 2.622361 1.2792 1.343161 2.942161 1.982761 4.349282 3.134041a47.074578 47.074578 0 0 0 29.421611 11.193005l0.575641-0.12792 0.57564 0.12792h861.157776l0.6396-0.12792 0.511681 0.12792a46.434978 46.434978 0 0 0 29.357651-11.256965c1.343161-1.15128 3.134041-1.790881 4.349282-3.134041 0.76752-0.76752 1.21524-1.790881 1.854841-2.622361a46.818738 46.818738 0 0 0 12.280324-30.956652l-0.12792-0.6396 0.12792-0.767521 0.12792-696.716552zM386.766271 95.940037h250.467458v193.223236H386.766271V95.940037z m352.483698 831.480325H284.813991v-250.851218l62.233105-62.233104h329.905808l62.297065 62.361024v250.723298z m156.510181 0h-60.570144v-267.097064c0-0.6396-0.3198-1.2792-0.38376-1.790881a48.865459 48.865459 0 0 0-14.391006-36.393254l-89.096314-88.904435a49.313179 49.313179 0 0 0-36.009494-14.454965c-0.57564 0-1.08732-0.3198-1.535041-0.3198H330.865209c-0.6396 0-1.21524 0.3198-1.85484 0.38376a49.249219 49.249219 0 0 0-36.457215 14.327045l-89.032354 88.968395a49.888819 49.888819 0 0 0-14.327046 36.585134c0 0.51168-0.3198 1.02336-0.3198 1.599001V927.420362H128.23985v-831.480325h162.586384v239.850094l0.12792 0.70356-0.12792 0.639601c0 11.640725 4.797002 21.938289 11.960524 30.253092 0.83148 1.08732 1.343161 2.430481 2.366521 3.389881 1.02336 1.02336 2.302561 1.471081 3.389882 2.366521a46.690818 46.690818 0 0 0 30.317051 11.960524l0.703561-0.12792 0.57564 0.12792h343.657214l0.76752-0.12792 0.703561 0.12792a46.371018 46.371018 0 0 0 30.317052-11.960524c1.15128-0.89544 2.366521-1.343161 3.389881-2.366521 1.08732-0.9594 1.471081-2.366521 2.366521-3.389881a46.626858 46.626858 0 0 0 12.024484-30.253092l-0.12792-0.639601 0.12792-0.70356V122.419488L895.76015 284.941911V927.420362z"
                    fill="#030000"
                    p-id="3008"
                  ></path>
                </svg>
              </el-icon>
            </el-button>
          </template>
          <setting-page />
        </el-drawer>
      </div>
      <!-- 下载 -->
      <div class="download-container">
        <el-badge :value="badge" class="item">
          <el-button type="primary" @click="downloadDrawer = true">
            <el-icon><Download /></el-icon>
          </el-button>
        </el-badge>
        <el-drawer v-model="downloadDrawer" title="下载" size="50%">
          <download-page />
        </el-drawer>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { EventsOn, EventsEmit, EventsOff } from "../../../wailsjs/runtime";
import { Search, Setting, Download } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import SettingPage from "@/components/setting-page/setting-page.vue";
import DownloadPage from "@/components/download-page/download-page.vue";
import { ref, reactive, onMounted, watchEffect, watch } from "vue";
import progressStore from "@/stores/progressStore";
import userStore from "@/stores/userStore";

const UserStore = userStore();
const ProgressStore = progressStore();
const link = ref("");
const select = ref("");
const dataSource = ref<Tree[]>();
const loading = ref(false);
const downloadBtn = ref(false);
const settingDrawer = ref(false);
const uuidList = ref<string[]>([]);
const uuidListDone = ref<string[]>([]);
const uuidSetDone = new Set<string>();
const uuidSet = new Set<string>();
const downloadDrawer = ref(false);
const platformList = reactive([
  {
    value: "1",
    label: "mooc",
  },
  {
    value: "2",
    label: "xtzx",
  },
  {
    value: "3",
    label: "dy",
  },
]);
const badge = ref(0);
const svg = `
        <path class="path" d="
          M 30 15
          L 28 17
          M 25.61 25.61
          A 15 15, 0, 0, 1, 15 30
          A 15 15, 0, 1, 1, 27.99 7.5
          L 15 15
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)"/>
      `;

const trees = ref(); //树形结构
const oldCheckedKeysLength = ref(0);

//选择节点时触发的函数
const handleNodeClick = (e: any, row: any) => {
  const checked = ((row.checkedKeys.length-oldCheckedKeysLength.value) > 0);
  // console.log("checked",checked);
  // console.log("row.checkedKeys.length", row.checkedKeys.length);
  if (e.uuid) {
    handleUuid(checked, e.uuid);
  } else {
    const childUuids = findMinChildUuids(e.children);
    for (const uuid of childUuids) {
      handleUuid(checked, uuid);
    }
  }
  oldCheckedKeysLength.value = row.checkedKeys.length;
  // console.log(uuidSet);
  // console.log(uuidSet.size);
  // selectionData is the selected values
};

function handleUuid(checked: boolean, uuid: string) {
  if (checked) {
    // console.log("---->选择", uuid);
    uuidSet.add(uuid);
  } else {
    // console.log("---->取消", uuid);
    uuidSet.delete(uuid);
  }
}

function findMinChildUuids(nodes: Tree[]): string[] {
  let uuids: string[] = [];
  for (const node of nodes) {
    if (node.children && node.children.length > 0) {
      const childUuids = findMinChildUuids(node.children);
      uuids = [...uuids, ...childUuids];
    } else {
      uuids.push(node.uuid);
    }
  }
  return uuids;
}

// 解析link
const onClickSearch = () => {
  loading.value = !loading.value;
  EventsEmit("parse_link", link.value);
};

// 下载
const onDownload = () => {
  // 下载
  uuidList.value.push(...uuidSet);
  EventsEmit("download", uuidList.value);
  // console.log("--------------------------------");
  // checkedNodes.value = CheckedNodes;
  // const minChildUuids = findMinChildUuids(checkedNodes.value);
  // console.log(minChildUuids);
};

// -------------------------

const onClickSave = () => {
  // console.log(UserStore.getMoocCookie())
  EventsEmit("update_setting", UserStore.getUserState());
  ElMessage({
    message: "保存成功！",
    grouping: true,
    type: "success",
    duration: 1500,
  });
};
// ----------1------------
interface progress {
  Uuid: string;
  Name: string;
  Percentage: number;
  Status: string;
  IsPause: boolean;
}

watchEffect(() => {
  // 获取 go 发送下载信息
  for (let i = 0; i < uuidList.value.length; i++) {
    const uuid = uuidList.value[i];
    EventsOn(uuid, (newProgress: progress) => {
      ProgressStore.setProgress(newProgress);
      console.log("-->newProgress", newProgress.Percentage);
      if (newProgress.Percentage == 100) {
        EventsOff(uuid);
        // uuidList.value.slice(i, 1);
        uuidListDone.value.push(uuid);
      }
      badge.value = uuidList.value.length - uuidListDone.value.length;
    });
  }
});

onMounted(() => {
  // 获取课程信息
  EventsOn("parse_link_result", (courseInfo) => {
    // console.log(courseInfo);
    // uuidList.value = courseInfo.UuidList;
    dataSource.value = buildTree(courseInfo.Units);
    loading.value = !loading.value;
    if (courseInfo.CourseName) {
      downloadBtn.value = !downloadBtn.value;
    } else {
      ElMessage({
        message: "解析失败！",
        grouping: true,
        type: "error",
        duration: 1500,
      });
    }
  });
});

interface Tree {
  id: string;
  label: string;
  uuid: string;
  children?: Tree[];
}

const props = {
  value: "id",
  label: "label",
  children: "children",
};

function buildTree(units: any[]): Tree[] {
  const tree: Tree[] = [];

  // Group units by chapter
  const chapters: { [key: string]: any[] } = units.reduce((acc, unit) => {
    if (!acc[unit.ChapterName]) {
      acc[unit.ChapterName] = [];
    }
    acc[unit.ChapterName].push(unit);
    return acc;
  }, {});

  // Build chapter nodes
  let chapterId = 1;
  for (const chapterName in chapters) {
    const chapterUnits = chapters[chapterName];
    const chapterNode: Tree = {
      id: `chapter-${chapterId}`,
      label: chapterName,
      uuid: "",
      children: [],
    };
    tree.push(chapterNode);
    chapterId++;

    // Group units by lesson
    const lessons: { [key: string]: any[] } = chapterUnits.reduce(
      (acc, unit) => {
        if (!acc[unit.LessonName]) {
          acc[unit.LessonName] = [];
        }
        acc[unit.LessonName].push(unit);
        return acc;
      },
      {}
    );

    // Build lesson nodes
    let lessonId = 1;
    for (const lessonName in lessons) {
      const lessonUnits = lessons[lessonName];
      const lessonNode: Tree = {
        id: `lesson-${chapterId}-${lessonId}`,
        label: lessonName,
        uuid: "",
        children: [],
      };
      chapterNode.children!.push(lessonNode);
      lessonId++;

      // Build unit nodes
      for (const unit of lessonUnits) {
        const unitNode: Tree = {
          id: `unit-${chapterId}-${lessonId}-${unit.UnitName}`,
          uuid: unit.Uuid,
          label: unit.UnitName,
        };
        lessonNode.children!.push(unitNode);
      }
    }
  }
  return tree;
}
</script>

<style lang="scss" scoped>
.container {
  .search-container {
    width: 80%;
    margin: 0 auto;
  }
  .info-container {
    width: 80%;
    margin: 0 auto;
    .custom-tree-container {
      //.download-btn {
      //  margin-right: 0;
      //}
      .custom-tree-node {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-size: 14px;
        padding-right: 8px;
      }
    }
  }
  .setting-download-container {
    .setting-container {
      .setting-drawer {
      }
    }
    .download-container {
      margin-top: 10px;
    }
  }
}
</style>
