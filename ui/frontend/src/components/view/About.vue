<template>
  <h1 cless="title" style="color: black">关于</h1>
  <el-row>
    <!-- 时间线 -->
    <el-col :span="8">
      <div class="timeline-box">
        <el-divider />
        <el-timeline>
          <el-timeline-item
            v-for="(activity, index) in activities"
            :key="index"
            :icon="activity.icon"
            :type="activity.type"
            :color="activity.color"
            :size="activity.size"
            :timestamp="activity.timestamp"
          >
            {{ activity.content }}
          </el-timeline-item>
        </el-timeline>
      </div>
    </el-col>
    <el-col :span="2"></el-col>
    <!-- 作者 -->
    <el-col :span="11">
      <el-divider />
      <el-card shadow="always">
        <div>
          <el-avatar
            style="text-align: center"
            :size="60"
            :src="avatar"
          ></el-avatar>
        </div>
        <!-- 作者 -->
        <div>
          <text> 作者:Esword </text>
        </div>
        <!-- 邮箱链接 -->
        <div>
          <text> 邮箱:jyj349559953@qq.com </text>
        </div>
        <!-- github链接 -->
        <div>
          <el-link
            @click="openLink('https://github.com/Esword618')"
            type="success"
            >github</el-link
          >
        </div>
        <!-- 博客链接 -->
        <div>
          <el-link
            @click="openLink('https://eswordblog.esword.space/#/')"
            type="success"
            >博客</el-link
          >
        </div>
        <!-- bilibil链接 -->
        <div>
          <el-link
            @click="openLink('https://space.bilibili.com/481314897')"
            type="success"
            >bilibili</el-link
          >
        </div>
        <!-- 公众号链接 -->
        <div>
          <el-link @click="gzhImgVisible = true" type="success">公众号</el-link>
          <el-dialog
            v-model="gzhImgVisible"
            title="扫码关注"
            width="20%"
            center
          >
            <div style="text-align: center">
              <el-image class="gzh-img" :src="gzh" fit="cover" />
            </div>
          </el-dialog>
        </div>
      </el-card>
      <el-divider border-style="dashed" />
      <text style="color: black">
        如果你想获取最新版，可以加入交流群，交流群在github上就可以获得群二维码
      </text>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { Check, CircleClose, Close } from "@element-plus/icons-vue";
import color from "element-plus/es/components/color-picker/src/color";
import { ElMessage } from "element-plus";
import gzh from "../../assets/images/gzh.png";
import avatar from "../../assets/images/logo.png";
import { BrowserOpenURL } from "../../../wailsjs/runtime/runtime.js";
// 公众号二维码
const gzhImgVisible = ref(false);
const url =
  "https://cdn.staticaly.com/gh/Esword56/blogImg@main/img/image.4f55783tzm00.jpg";

const openLink = (link: string) => {
  BrowserOpenURL(link);
};

const activities = [
  {
    content: "python版初次提交",
    timestamp: "2021-07-12",
    size: "large",
    type: "primary",
    icon: Check,
  },
  {
    content: "加入视频下载",
    timestamp: "2021-09-09",
    color: "#0bbd87",
    icon: Check,
  },
  {
    content: "加入测验试卷下载",
    timestamp: "2022-02-20",
    size: "large",
    color: "#25e0da",
    icon: Check,
  },
  {
    content: "go版本初次提交",
    timestamp: "2022-04-06",
    type: "primary",
    icon: Check,
    color: "#e02525",
  },
  {
    content: "更新目录结果，加密码本地化运行",
    timestamp: "2022-07-17",
    icon: Check,
    color: "#5725e0",
  },
  {
    content: "支持其它网课平台",
    timestamp: "unknown",
    icon: Close,
    size: "large",
    color: "#9b0e38",
  },
];
</script>

<style>
.timeline-box {
  margin-left: 25px;
  /* font-size: 15px; */
}
.gzh-img {
  /* width: 200px;
    height: 200px; */
  margin-bottom: 20px;
  box-shadow: 0px 0px 3px 3px rgb(55 85 218 / 20%);
  border: 2px solid #5bebe2;
  border-radius: 8pxrgb (194, 201, 201);
}
</style>
