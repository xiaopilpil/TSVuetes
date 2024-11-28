<template>
    <el-container class="layout-container-demo" style="height: 100vh; flex-direction: column;">
      <!-- 引入拆分的组件 -->
      <HeaderComponent />
      <el-container style="flex: 1;">
        <AsideComponent />
        <MainComponent />
      </el-container>
    </el-container>
</template>
  
<script lang="ts" setup>
  import HeaderComponent from './Header.vue'
  import AsideComponent from './Aside.vue'
  import MainComponent from './Main.vue'
</script>
  
<style scoped>
  .layout-container-demo {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }
</style>
  