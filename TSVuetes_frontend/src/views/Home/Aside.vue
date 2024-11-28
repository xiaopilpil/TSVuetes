<template>
  <el-aside width="200px" style="background-color: #d6e8fa; text-align: center;">
    <el-scrollbar>
      <el-menu :default-openeds="['1', '3']" @select="handleSelect">
        <el-sub-menu index="1">
          <template #title>
            <el-icon><message /></el-icon>课室管理
          </template>
          <el-menu-item-group>
            <!-- 条件显示管理员课程列表 -->
            <el-menu-item
              v-if="isAdmin"
              index="acourse"
            >
              管理员课程列表
            </el-menu-item>
            <!-- 条件显示教室创建 -->
            <el-menu-item
              v-if="isAdmin"
              index="createroom"
            >
              教室创建
            </el-menu-item>
            <el-menu-item index="course">教师课程列表</el-menu-item>
            <el-menu-item index="apply">教室申请</el-menu-item>
          </el-menu-item-group>
        </el-sub-menu>
        <!-- 其他菜单项 -->
                   <!-- <el-sub-menu index="2">
            <template #title>
              <el-icon><icon-menu /></el-icon>学生管理
            </template>
            <el-menu-item-group>
              <el-menu-item index="news">Option 1</el-menu-item>
              <el-menu-item index="2-2">Option 2</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu> -->
          <!-- <el-sub-menu index="3">
            <template #title>
              <el-icon><setting /></el-icon>系统管理
            </template>
            <el-menu-item-group>
              <el-menu-item index="3-1">Option 1</el-menu-item>
              <el-menu-item index="3-2">Option 2</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu> -->
      </el-menu>
    </el-scrollbar>
  </el-aside>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { useAuthStore } from '../../store/auth'; // 假设有 auth store

const router = useRouter();

// 获取用户身份
const authStore = useAuthStore();
const isAdmin = authStore.teacherId && authStore.positionId === 1; // 假设 position 1 是管理员

const handleSelect = (key: string) => {
  router.push({ name: key.charAt(0).toUpperCase() + key.slice(1) });
};
</script>

<style scoped>
.el-aside {
  background-color: #d6e8fa;
  text-align: center;
}
</style>
