<template>
  <main class="main-container">
    <!--无边框装饰-->
    <div class="frameless-container">
      <header-frameless />
    </div>
    <!--body设计-->
    <div class="body-container">
      <body-container />
    </div>
  </main>
</template>

<script setup lang="ts">
import HeaderFrameless from "@/components/header-frameless/header-frameless.vue";
import BodyContainer from "@/components/body-container/body-container.vue";

</script>

<style lang="scss" scoped>
.body-container {
  margin-top: 5%;
}
</style>
