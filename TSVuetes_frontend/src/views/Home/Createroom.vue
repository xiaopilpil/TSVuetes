<template>
    <div>
      <h2>创建教室</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="教室名" prop="room_id">
          <el-input v-model="form.room_id" placeholder="请输入教室名"></el-input>
        </el-form-item>
        <el-form-item label="设备号" prop="device_id">
          <el-input v-model="form.device_id" placeholder="请输入设备号"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm">创建</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </template>
  
  <script lang="ts" setup>
  import { reactive, ref } from 'vue';
  import { createroom } from './Createroom';
  import { ElMessage, ElMessageBox } from 'element-plus';
  
  // 表单数据和验证规则
  const form = reactive({
    room_id: '',
    device_id: '',
  });
  
  const rules = {
    room_id: [
      { required: true, message: '教室名不能为空', trigger: 'blur' },
    ],
    device_id: [
      { required: true, message: '设备号不能为空', trigger: 'blur' },
    ],
  };
  
  const formRef = ref();
  
  const submitForm = async () => {
    // 提交前验证表单
    formRef.value?.validate(async (valid: boolean) => {
      if (valid) {
        try {
          const response = await createroom(form.room_id, form.device_id);
          ElMessage({
            message: '教室创建成功',
            type: 'success',
          });
          console.log('Response:', response);
          resetForm();
        } catch (error: any) {
          console.error('创建失败:', error);
          const errorMessage = error?.response?.data?.error || '创建教室失败，请重新尝试';
          if (errorMessage.includes('23505')) {
            ElMessageBox.alert(`教室名 "${form.room_id}" 已存在，请尝试使用其他教室名`, '创建失败', {
              type: 'error',
            });
          } else {
            ElMessageBox.alert(errorMessage, '创建失败', {
              type: 'error',
            });
          }
        }
      } else {
        console.log('表单验证失败');
      }
    });
  };
  
  const resetForm = () => {
    formRef.value?.resetFields();
  };
  </script>
  
  <style scoped>
  h2 {
    margin-bottom: 20px;
  }
  </style>
  