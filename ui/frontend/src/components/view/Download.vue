<template>
  <h1 cless="title" style="color: black">这是下载</h1>
  <el-row>
    <el-col :span="2"></el-col>
    <el-col :span="15">
      <div class="progress-list">
        <el-progress
          :text-inside="true"
          :stroke-width="22"
          :percentage="percentage2"
          :color="colors"
          status="success"
        />
        <el-progress
          :text-inside="true"
          :stroke-width="22"
          :percentage="percentage2"
          :color="colors"
          status="warning"
        />
        <el-progress
          :text-inside="true"
          :stroke-width="22"
          :percentage="percentage2"
          :color="colors"
          status="exception"
        />
      </div>
      <!-- <div class="demo-progress">
        <el-progress :text-inside="true" :stroke-width="26" :percentage="70" />
        <el-progress
          :text-inside="true"
          :stroke-width="24"
          :percentage="100"
          status="success"
        />
        <el-progress
          :text-inside="true"
          :stroke-width="22"
          :percentage="80"
          status="warning"
        />
        <el-progress
          :text-inside="true"
          :stroke-width="20"
          :percentage="50"
          status="exception"
        />
      </div> -->
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { Minus, Plus } from "@element-plus/icons-vue";
import { EventsEmit, EventsOn } from "../../../wailsjs/runtime/runtime.js";
import { ElMessage } from "element-plus";
const percentage = ref(10);
const percentage2 = ref(10);
const colors = [
  { color: "#f56c6c", percentage: 20 },
  { color: "#e6a23c", percentage: 40 },
  { color: "#5cb87a", percentage: 60 },
  { color: "#1989fa", percentage: 80 },
  { color: "#6f7ad3", percentage: 100 },
];

const increase = () => {
  percentage.value += 10;
  if (percentage.value > 100) {
    percentage.value = 100;
  }
};
const decrease = () => {
  percentage.value -= 10;
  if (percentage.value < 0) {
    percentage.value = 0;
  }
};

function test(info: any) {
  ElMessage({
    message: "链接不能为空!",
    type: "info",
  });
  console.log("----->", info);
}

onMounted(() => {
  setInterval(() => {
    // mypercentage();
    percentage2.value = (percentage2.value % 100) + 10;
    EventsOn("test", test);
  }, 500);
});
</script>

<style>
.progress-list .el-progress--line {
  margin-bottom: 15px;
  width: 350px;
}
</style>
