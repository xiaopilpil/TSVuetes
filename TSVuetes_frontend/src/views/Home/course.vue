<template>
  <div>

  <el-table v-if="!isLoading" :data="filterTableData" style="width: 100%">
    <!-- 普通列 -->
    <el-table-column label="Teacher ID" prop="teacher_id" />
    <el-table-column label="Name" prop="name" />
    <el-table-column label="Room ID" prop="room_id" />
    <el-table-column label="Remarks">
      <template #default="{ row }">
        <template v-if="row.state === 2">
          <!-- 状态为已拒绝时显示备注信息 -->
          <span>拒绝理由:{{ row.ps || '被管理员拒绝' }}</span>
        </template>
        <template v-else-if="row.state === 1">
          <!-- 状态为已通过时显示编辑按钮 -->
          <el-button size="small" @click="handleEdit(row)">
            Edit
          </el-button>
        </template>
        <template v-else>
          <!-- 状态为申请中 -->
          <span>申请中</span>
        </template>
      </template>
    </el-table-column>
    <el-table-column align="right">
      <template #header>
        <el-input v-model="search" size="small" placeholder="Type to search" />
      </template>
      <template #default="{ row }">
        <!-- 状态为已通过时无按钮 -->
        <template v-if="row.state === 0 || row.state === 2">
          <!-- 对于申请中或已拒绝，不显示按钮 -->
          <span></span>
        </template>
      </template>
    </el-table-column>
  </el-table>
    <div v-else class="loading-container">
    <el-spinner />
      <p>加载中，请稍候...</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import { fetchcourses, courses } from './course';
import { useRouter } from 'vue-router'; // 引入 vue-router
import { useAuthStore } from '../../store/auth';

const router = useRouter(); // 路由实例
const authStore = useAuthStore();
const isLoading = ref(false)
// 搜索框绑定值
const search = ref('');

// 加载数据
// const loadData = async () => {
//   await fetchcourses(); // 加载课程数据
// };
const loadData = async () => {
  courses.value = []; // 清空数据
  isLoading.value = true; // 显示加载动画
  try {
    await fetchcourses(); // 加载新的数据
  } finally {
    isLoading.value = false; // 加载完成
  }
};

// 初次加载时获取数据
onMounted(() => {
  authStore.syncStateWithLocalStorage(); // 确保页面加载时状态同步
});
onMounted(loadData);


// 过滤后的表格数据
const filterTableData = computed(() =>
  courses.value.filter(
    (course) =>
      !search.value ||
      course.teacher_id.toLowerCase().includes(search.value.toLowerCase()) ||
      course.name.toLowerCase().includes(search.value.toLowerCase())
  )
);

// 编辑操作
const handleEdit = (row: typeof courses.value[number]) => {
  router.push({ name: 'Classroom', params: { id: row.room_id } });
};
</script>
