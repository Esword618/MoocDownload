<template>
  <div class="container">
    <el-scrollbar height="690px">
      <div
        class="container-progress"
        v-for="(progress, key) in progressInfo"
        :key="key"
      >
        <el-progress
          :text-inside="true"
          :stroke-width="26"
          :percentage="progress.Percentage"
          :status="progress.Status"
        />
        <li>
          {{ progress.Name }}
        </li>
      </div>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { EventsOn } from "../../../wailsjs/runtime";
import { onMounted, ref, reactive, watch, computed } from "vue";
import progressStore from "@/stores/progressStore";

const ProgressStore = progressStore();

const progressInfo = computed(() => {
  return Object.values(ProgressStore.getProgresses());
});
</script>

<style lang="scss" scoped>
.container {

  .container-progress {
    overflow: hidden;
  }
  .container-progress .el-progress--line {
    margin-bottom: 15px;
    width: 350px;
  }
}

</style>
