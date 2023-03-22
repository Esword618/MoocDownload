<template>
  <el-config-provider>
<!--    <main-container />-->
    <home-view />
  </el-config-provider>
</template>

<script setup lang="ts">
// import MainContainer from "./components/main-container/main-container.vue";
import HomeView from "./views/HomeView/HomeView.vue"
// document监听
// document.body.addEventListener("click", function (event) {
//   event.preventDefault();
// });
</script>

<style lang="scss">
html {
  width: 100%;
  height: 100%;
  overflow: hidden;
}
::-webkit-scrollbar {
  width: 0 !important;
}
::-webkit-scrollbar {
  width: 0 !important;
  height: 0;
}
</style>
