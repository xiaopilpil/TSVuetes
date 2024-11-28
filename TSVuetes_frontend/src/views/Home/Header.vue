<template>
  <el-header class="header-container">
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
    >
      <template v-if="!isAuthenticated">
        <el-menu-item index="login">Login</el-menu-item>
        <el-menu-item index="register">Register</el-menu-item>
      </template>
      <template v-else>
        <el-menu-item index="profile">{{ teacherId }}</el-menu-item>
        <el-menu-item index="logout">Logout</el-menu-item>
      </template>
    </el-menu>
  </el-header>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAuthStore } from '../../store/auth'; // 引入你的 Pinia Auth Store

const router = useRouter();
const route = useRoute();

const authStore = useAuthStore(); // 使用 Pinia store
const activeIndex = ref(route.name?.toString() || 'home');

watch(route, (newRoute) => {
  activeIndex.value = newRoute.name?.toString() || 'home';
});

const isAuthenticated = authStore.isAuthenticated; // 是否登录
const teacherId = authStore.teacherId; // 当前 teacher_id

const handleSelect = (key: string) => {
  if (key === 'logout') {
    authStore.logout();
    router.push({ name: 'Login' }); // 退出后跳转到登录页面
  } else {
    router.push({ name: key.charAt(0).toUpperCase() + key.slice(1) });
  }
};
</script>

<style scoped>
.header-container {
  display: flex;
  justify-content: flex-end; /* 将内容靠右对齐 */
  align-items: center; /* 垂直居中 */
  height: 64px;
  background-color: #ffffff;
  padding-right: 20px;
}

.el-menu {
  display: flex;
  justify-content: flex-end; /* 确保菜单项靠右 */
  flex: 1; /* 使菜单项占据所有可用空间 */
}
</style>
