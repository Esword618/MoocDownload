<template>
  <div class="setting-container">
    <el-tabs
      type="border-card"
      v-model="activeName"
      class="demo-tabs"
      @tab-click="handleClick"
    >
      <el-tab-pane label="MOOC" name="mooc">
        <el-input
          v-model="setting.moocCookie"
          :autosize="{ minRows: 4, maxRows: 4 }"
          type="textarea"
          placeholder="请把cookie放入此处"
          resize="none"
        />
      </el-tab-pane>
      <el-tab-pane label="XTZX" name="xtzx">
        <div>xtzx开发中...</div>
      </el-tab-pane>
      <el-tab-pane label="DY" name="dy">
        <div>dy开发中...</div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import {ref, reactive, watch, onMounted} from "vue";
import type { TabsPaneContext } from "element-plus";
import userStore from "@/stores/userStore";
import {EventsOff, EventsOn} from "../../../wailsjs/runtime";
import {ElMessage} from "element-plus";

const UserStore = userStore();
const activeName = ref("mooc");
const setting = reactive({
  moocCookie: "",
  // d:"",
});

watch(setting, (newValue, oldValue) => {
  setting.moocCookie = newValue.moocCookie;
  // console.log(newValue.moocCookie);
  // console.log(newValue.d);
  UserStore.moocCookie = setting.moocCookie;
});

onMounted(() => {
  // 获取设置信息
  // EventsOn("default_setting", (defaultSetting: any) => {
  //   console.log(defaultSetting);
  //   setting.d = defaultSetting;
  //   // console.log("------");
  // });
  setting.moocCookie = UserStore.getMoocCookie();
});

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};
</script>

<style lang="scss" scoped>
.bug {
  //background-color: #ec2e2e;
}
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
