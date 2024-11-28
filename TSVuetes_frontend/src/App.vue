<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <!-- 一级路由串口组件 -->
  <RouterView />
</template>

<style scoped>
.layout-container-demo {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
</style>
