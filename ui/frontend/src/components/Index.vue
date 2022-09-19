<template>
  <el-row style="height: 100%">
    <!-- 主页，下载，关于，设置 -->
    <el-col class="aside-box" :span="3">
      <!-- 主页，下载，关于 -->
      <el-menu
        class="menu-box"
        router="true"
        default-active="1"
        text-color="#fff"
        active-text-color="#23efef"
        style="
          max-height: calc(100vh - 50px);
          min-height: calc(100vh - 50px);
          overflow: hidden auto;
        "
      >
        <el-avatar
          style="text-align: center"
          :size="60"
          :src="logo"
        ></el-avatar>
        <el-menu-item index="1" route="/home">
          <template #title>
            <el-icon><House /></el-icon>主页
          </template>
        </el-menu-item>
        <el-menu-item index="2" route="/download">
          <template #title>
            <el-icon><Download /></el-icon>下载
          </template>
        </el-menu-item>
        <el-menu-item index="3" route="/about">
          <template #title>
            <el-icon><Warning /></el-icon>关于
          </template>
        </el-menu-item>
      </el-menu>
      <!-- 底部 设置 按钮 -->
      <div class="setting-box">
        <!-- 设置按钮 -->
        <el-button
          class="setting-btn"
          @click="settingDrawerVisible = true"
          type="primary"
          :icon="Setting"
        >
          设&nbsp;&nbsp;&nbsp;&nbsp;置
        </el-button>
        <!-- 设置按钮触发的抽屉 -->
        <el-drawer
          v-model="settingDrawerVisible"
          title="设置"
          size="35%"
          style="height: 100vh"
        >
          <el-tabs>
            <!-- mooc -->
            <el-tab-pane label="Mooc">
              <el-form
                label-width="100px"
                :model="settingData"
                style="max-width: 460px"
                label-position="left"
              >
                <el-form-item label="Cookie">
                  <el-input
                    :autosize="{ minRows: 1, maxRows: 4 }"
                    type="textarea"
                    v-model="settingData.moocCookie"
                  />
                </el-form-item>
                <el-divider />
                <el-form-item label="爱课程账号">
                  <el-input v-model="settingData.moocAccount" />
                </el-form-item>
                <el-form-item label="爱课程密码">
                  <el-input
                    v-model="settingData.moocPassword"
                    type="password"
                  />
                </el-form-item>
              </el-form>
              <span class="dialog-footer">
                <el-button
                  type="primary"
                  @click="settingMoocDialogVisible = true"
                  >扫码登录</el-button
                >
                <el-button type="primary" @click="open">登录</el-button>
              </span>
              <!-- 登录对话框 -->
              <el-dialog
                v-model="settingMoocDialogVisible"
                title="扫码登录"
                width="25%"
                center
              >
                <div style="text-align: center">
                  <el-image class="mooc-qr-img" :src="url" fit="cover" />
                </div>
                <div style="text-align: center">
                  打开 中国大学MOOC App - 首页左上角 - 扫一扫登录
                  4.2.0以上版本可扫码登录
                </div>
              </el-dialog>
            </el-tab-pane>
            <!-- xtzx -->
            <el-tab-pane label="Xtzx">
              <el-form
                label-width="100px"
                :model="settingData"
                style="max-width: 460px"
                label-position="left"
              >
                <el-form-item label="Cookie">
                  <el-input
                    :autosize="{ minRows: 1, maxRows: 4 }"
                    type="textarea"
                    v-model="settingData.xtzxCookie"
                  />
                </el-form-item>
              </el-form>
            </el-tab-pane>
            <!-- bilibil -->
            <el-tab-pane label="Bilibili">
              <el-form
                label-width="100px"
                :model="settingData"
                style="max-width: 460px"
                label-position="left"
              >
                <el-form-item label="Cookie">
                  <el-input
                    :autosize="{ minRows: 1, maxRows: 4 }"
                    type="textarea"
                    v-model="settingData.bilibiliCookie"
                  />
                </el-form-item>
              </el-form>
            </el-tab-pane>
          </el-tabs>
        </el-drawer>
      </div>
    </el-col>
    <el-col :span="21" class="main-box">
      <router-view></router-view>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import { Download, House, Setting, Warning } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import moocqr from "../assets/images/moocqr.png";
import logo from "../assets/images/logo.png";
// 抽屉可见
const settingDrawerVisible = ref(false);
// 对话框架可见
const settingMoocDialogVisible = ref(false);
// 账号登录
const constsettingMoocAccountBtn = ref(false);
const open = () => {
  ElMessage("this is a message.");
};
const url =
  "https://cdn.staticaly.com/gh/Esword56/blogImg@main/img/image.4f55783tzm00.jpg";

const settingData = reactive({
  moocCookie: "",
  moocAccount: "",
  moocPassword: "",
  xtzxCookie: "",
  bilibiliCookie: "",
});
</script>

<style scoped>
.aside-box {
  background-color: #191a23;
}
.menu-box {
  background-color: #191a23;
}
.setting-box {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 2px 4px;
}
.setting-btn {
  width: 105px;
  margin: 5px 0;
  font-size: 9px;
  color: #fff;
  background-color: #5fa6ed;
  border-color: #23efef;
}
.main-box {
  background-color: #f1ecec;
}
.dialog-footer button:first-child {
  margin-right: 10px;
}
.mooc-qr-img {
  width: 200px;
  height: 200px;
  margin-bottom: 20px;
  box-shadow: 0px 0px 3px 3px rgb(55 85 218 / 20%);
  border: 2px solid #acd7dc;
  border-radius: 8pxrgb (194, 201, 201);
}
</style>
