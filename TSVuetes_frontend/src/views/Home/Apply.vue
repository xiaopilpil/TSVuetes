<template>
    <div>
      <h2>教室申请</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="教师名" prop="teacher_id">
          <el-input v-model="form.teacher_id" placeholder="请输入教师ID" />
        </el-form-item>
        <el-form-item label="教室名" prop="room_id">
          <el-input v-model="form.room_id" placeholder="请输入教室ID" />
        </el-form-item>
        <el-form-item label="课程名" prop="name">
          <el-input v-model="form.name" placeholder="请输入课程名称" />
        </el-form-item>
        <el-form-item label="院系" prop="Faculty">
          <el-input v-model="form.Faculty" placeholder="请输入学院" />
        </el-form-item>
        <el-form-item label="学生人数" prop="Snumber">
          <el-input-number v-model="form.Snumber" :min="1" placeholder="请输入学生人数" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">提交申请</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref } from 'vue';
  import { applycourse } from './Apply'; // 引入申请教室的接口方法
  import { ElMessage } from 'element-plus'; // 引入消息提示组件
  
  // 表单数据
  const form = ref({
    teacher_id: '',
    room_id: '',
    name: '',
    Faculty: '',
    Snumber: 1,
  });
  
  // 表单验证规则
  const rules = {
    teacher_id: [
      { required: true, message: '教师ID不能为空', trigger: 'blur' },
    ],
    room_id: [
      { required: true, message: '教室ID不能为空', trigger: 'blur' },
    ],
    name: [
      { required: true, message: '课程名称不能为空', trigger: 'blur' },
    ],
    Faculty: [
      { required: true, message: '学院不能为空', trigger: 'blur' },
    ],
    Snumber: [
      { required: true, message: '学生人数不能为空', trigger: 'change' },
      { type: 'number', min: 1, message: '学生人数必须大于0', trigger: 'change' },
    ],
  };
  
  // 引用表单
  const formRef = ref();
  
  // 提交表单
  const handleSubmit = () => {
  formRef.value?.validate(async (valid: boolean) => {
    if (valid) {
      try {
        const response = await applycourse(
          form.value.teacher_id,
          form.value.room_id,
          form.value.name,
          form.value.Faculty,
          form.value.Snumber
        );
        if (response) {
          ElMessage.success('申请成功！');
          resetForm(); // 清空表单
        } else {
          ElMessage.error('申请失败，请稍后重试！');
        }
      } catch (error: unknown) {
        // 明确处理错误类型
        if (error instanceof Error) {
          ElMessage.error(`申请失败：${error.message}`);
        } else {
          ElMessage.error('申请失败：未知错误');
        }
      }
    } else {
      ElMessage.error('请完整填写表单后提交！');
    }
  });
};

  
  // 重置表单
  const resetForm = () => {
    form.value = {
      teacher_id: '',
      room_id: '',
      name: '',
      Faculty: '',
      Snumber: 1,
    };
    formRef.value?.resetFields();
  };
  </script>
  
  <style scoped>
  h2 {
    margin-bottom: 20px;
  }
  </style>
  