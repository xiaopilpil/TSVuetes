<template>
    <el-main style="background-color: #eef5ff; text-align: center;">
      <el-scrollbar style="height: calc(100vh - 64px);">
        <RouterView />
      </el-scrollbar>
    </el-main>
  </template>
  
  <style scoped>
  .el-main {
    background-color: #eef5ff;
    text-align: center;
  }
  </style>
  