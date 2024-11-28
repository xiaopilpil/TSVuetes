<template>
    <el-container class="register-container">
      <el-header class="register-header">
        <h1>学生行为检测管理系统</h1>
      </el-header>
      <el-main class="register-main">
        <el-card class="register-card">
          <h2>账户注册</h2>
          <el-form :model="registerForm" :rules="rules" ref="registerFormRef" label-width="60px">
            <el-form-item label="账号" prop="teacher_id">
              <el-input v-model="registerForm.teacher_id" placeholder="请输入账号"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input type="password" v-model="registerForm.password" placeholder="请输入密码"></el-input>
            </el-form-item>
            <el-form-item label="姓名" prop="name">
              <el-input v-model="registerForm.name" placeholder="请输入姓名"></el-input>
            </el-form-item>
            <el-form-item label="性别" prop="gender">
              <el-select v-model="registerForm.gender" placeholder="请选择性别">
                <el-option label="男" value="男" />
                <el-option label="女" value="女" />
              </el-select>
            </el-form-item>
            <el-form-item label="学院" prop="faculty">
              <el-input v-model="registerForm.faculty" placeholder="请输入学院"></el-input>
            </el-form-item>
            <el-form-item label="角色" prop="position">
              <el-radio-group v-model="registerForm.position">
                <el-radio :label="0">普通用户</el-radio>
                <el-radio :label="1">管理员</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item>
              <el-checkbox-group v-model="registerForm.agreeTerms">
                <el-checkbox>我已同意隐私条款和服务条款</el-checkbox>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm">立即注册</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-main>
    </el-container>
  </template>
  
  <script lang="ts" setup>
  import { ref } from 'vue';
  import { ElMessage, ElMessageBox } from 'element-plus';
  import { useRouter } from 'vue-router';
  import { useAuthStore } from '../../store/auth';
  
  const router = useRouter();
  const authStore = useAuthStore();
  
  const registerForm = ref({
    teacher_id: '',
    password: '',
    name: '',
    gender: '',
    faculty: '',
    position: 0, // 默认角色：普通用户
    agreeTerms: [], // 用户是否同意条款
  });
  
  const rules = ref({
    teacher_id: [
      { required: true, message: '请输入账号', trigger: 'blur' },
    ],
    password: [
      { required: true, message: '请输入密码', trigger: 'blur' },
      { min: 6, max: 14, message: '密码长度为6-14个字符', trigger: 'blur' },
    ],
    name: [
      { required: true, message: '请输入姓名', trigger: 'blur' },
    ],
    gender: [
      { required: true, message: '请选择性别', trigger: 'change' },
    ],
    faculty: [
      { required: true, message: '请输入学院', trigger: 'blur' },
    ],
    position: [
      { required: true, message: '请选择角色', trigger: 'change' },
    ],
    agreeTerms: [
      { type: 'array', required: true, message: '请同意隐私条款和服务条款', trigger: 'change' },
    ],
  });
  
  const registerFormRef = ref();
  
  const submitForm = async () => {
    try {
      await registerFormRef.value?.validate();
      // 调用注册逻辑
      await authStore.register(
        registerForm.value.teacher_id,
        registerForm.value.password,
        registerForm.value.name,
        registerForm.value.gender,
        registerForm.value.faculty,
        registerForm.value.position
      );
      ElMessage.success('注册成功！即将跳转至登录页面');
      router.push({ name: 'Home' });
    } catch (error: any) {
      if (error?.response?.data?.message) {
        ElMessageBox.alert(error.response.data.message, '注册失败', {
          type: 'error',
        });
      } else {
        ElMessageBox.alert('注册失败，请重试', '错误', {
          type: 'error',
        });
      }
    }
  };
  </script>
  
  <style scoped>
  .register-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #e8f5fa;
  }
  
  .register-header {
    text-align: center;
    font-size: 2rem;
    margin-bottom: 20px;
  }
  
  .register-main {
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .register-card {
    width: 400px;
    padding: 20px;
    box-shadow: 0px 0px 20px rgba(0, 0, 0, 0.1);
    text-align: center;
  }
  
  h2 {
    margin-bottom: 20px;
  }
  
  .el-input {
    width: 100%;
  }
  
  .el-form-item {
    margin-bottom: 20px;
  }
  
  .el-button {
    width: 100%;
  }
  </style>
  