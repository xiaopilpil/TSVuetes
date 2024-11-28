<template>
  <div>
    <el-table v-if="!isLoading" :data="filterTableData" style="width: 100%">
      <el-table-column label="Teacher ID" prop="teacher_id" />
      <el-table-column label="Name" prop="name" />
      <el-table-column label="Room ID" prop="room_id" />
      <el-table-column label="State">
        <template #default="scope">
          <!-- 状态显示逻辑 -->
          <span v-if="scope.row.state === 0">申请中</span>
          <span v-else-if="scope.row.state === 1">已通过</span>
          <span v-else>已拒绝</span>
        </template>
      </el-table-column>
      <el-table-column align="right">
        <template #header>
          <el-input v-model="search" size="small" placeholder="Type to search" />
        </template>
        <template #default="scope">
          <!-- 按钮显示逻辑 -->
          <template v-if="scope.row.state === 0">
            <!-- 申请中时显示“通过”和“拒绝”按钮 -->
            <el-button
              size="small"
              type="success"
              @click="handleApprove(scope.$index, scope.row)"
            >
            同意
            </el-button>
            <el-button
              size="small"
              type="danger"
              @click="openRejectDialog(scope.row)"
            >
            拒绝
            </el-button>
          </template>
          <template v-else>
            <!-- 已通过或已拒绝状态，显示撤回按钮 -->
            <el-button
              size="small"
              type="warning"
              @click="handleRevert(scope.row)"
            >
            撤回
            </el-button>
          </template>
        </template>
      </el-table-column>
    </el-table>
        <div v-else class="loading-container">
    <el-spinner />
      <p>加载中，请稍候...</p>
    </div>
  </div>
  
    <!-- 拒绝理由弹出框 -->
    <el-dialog
      v-model="rejectDialogVisible"
      title="拒绝理由"
      width="30%"
      :before-close="handleDialogClose"
    >
      <el-form ref="rejectFormRef">
        <el-form-item label="理由" required>
          <el-input
            type="textarea"
            v-model="rejectReason"
            placeholder="请输入拒绝理由"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="rejectDialogVisible = false">取消</el-button>
        <el-button type="danger" @click="submitRejectReason">确定</el-button>
      </template>
    </el-dialog>
    
  </template>
  
  
  
  
  <script lang="ts" setup>
  import { ref, computed, onMounted } from 'vue';
  import { fetchcourses, courses, updatecourse, updateps } from './Acourse';
  import { useAuthStore } from '../../store/auth';
  const search = ref('');
  const rejectDialogVisible = ref(false); // 控制弹出框显示
  const rejectReason = ref(''); // 拒绝理由
  const currentRow = ref<typeof courses.value[number] | null>(null); // 当前操作的行
  const authStore = useAuthStore()
  const isLoading = ref(false)

  // 加载数据
  const loadData = async () => {
  courses.value = []; // 清空数据
  isLoading.value = true; // 显示加载动画
  try {
    await fetchcourses(); // 加载新的数据
  } finally {
    isLoading.value = false; // 加载完成
  }
};

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
  
  // 审核通过
  const handleApprove = async (_index: number, row: typeof courses.value[number]) => {
    const originalState = row.state;
    row.state = 1;
  
    try {
      const success = await updatecourse(String(row.teacher_id), row.room_id, 1);
      if (!success) {
        throw new Error('Failed to update course state.');
      }
    } catch (error) {
      console.error('Approval failed:', error);
      row.state = originalState;
    }
  };
  
  // 打开拒绝弹出框
  const openRejectDialog = (row: typeof courses.value[number]) => {
    currentRow.value = row;
    rejectReason.value = ''; // 清空拒绝理由
    rejectDialogVisible.value = true;
  };
  
  // 提交拒绝理由
  const submitRejectReason = async () => {
    if (!currentRow.value) return;
  
    const originalState = currentRow.value.state;
    currentRow.value.state = 2;
  
    try {
      const success = await updatecourse(
        String(currentRow.value.teacher_id),
        currentRow.value.room_id,
        2
      );
      if (!success) {
        throw new Error('Failed to update course state.');
      }
  
      // 调用 updateps 更新备注信息
      const psSuccess = await updateps(
        String(currentRow.value.teacher_id),
        currentRow.value.room_id,
        rejectReason.value
      );
      if (!psSuccess) {
        throw new Error('Failed to update rejection reason.');
      }
  
      console.log('Rejection successful. Reason:', rejectReason.value);
    } catch (error) {
      console.error('Rejection failed:', error);
      currentRow.value.state = originalState; // 回滚状态
    } finally {
      rejectDialogVisible.value = false; // 关闭弹出框
    }
  };
  
  // 撤回操作
  const handleRevert = async (row: typeof courses.value[number]) => {
    console.log('Reverting:', row);
    const originalState = row.state;
    row.state = 0;
  
    try {
      const success = await updatecourse(String(row.teacher_id), row.room_id, 0);
      if (!success) {
        throw new Error('Failed to revert course state.');
      }
  
      // 调用 updateps 清空备注信息
      const psSuccess = await updateps(String(row.teacher_id), row.room_id, '');
      if (!psSuccess) {
        throw new Error('Failed to clear rejection reason.');
      }
  
      console.log('Revert successful.');
    } catch (error) {
      console.error('Revert failed:', error);
      row.state = originalState; // 回滚状态
    }
  };
  
  // 处理弹出框关闭
  const handleDialogClose = () => {
    rejectDialogVisible.value = false;
  };
  </script>
  
  
  