<template>
  <el-container class="login-container">
    <el-header class="login-header">
      <h1>学生行为检测管理系统</h1>
    </el-header>
    <el-main class="login-main">
      <el-card class="login-card">
        <h2>账户登录</h2>
        <el-form :model="loginForm" :rules="rules" ref="loginFormRef" label-width="60px">
          <el-form-item label="账号" prop="teacher_id">
            <el-input v-model="loginForm.teacher_id" placeholder="请输入账号"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input type="password" v-model="loginForm.password" placeholder="请输入密码"></el-input>
          </el-form-item>
          <el-form-item prop="agreeTerms">
            <el-checkbox-group v-model="loginForm.agreeTerms">
              <el-checkbox>我已同意隐私条款和服务条款</el-checkbox>
            </el-checkbox-group>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm()">点击登录</el-button>
          </el-form-item>
        </el-form>
        <!-- 添加忘记密码和注册账号的按钮 -->
        <div class="additional-links">
          <el-link @click="goToRegister" type="primary">注册账号</el-link>
          <el-link @click="goToForgotPassword" type="danger">忘记密码</el-link>
        </div>
      </el-card>
    </el-main>
  </el-container>
</template>

<script lang="ts" setup>
import { useAuthStore } from '../../store/auth';
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';

const loginForm = ref({
  teacher_id: '',
  password: '',
  agreeTerms: [],
});

const rules = ref({
  teacher_id: [
    { required: true, message: '请输入账号', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 14, message: '密码长度为6-14个字符', trigger: 'blur' },
  ],
  agreeTerms: [
    { required: true, message: '请同意隐私条款和服务条款', trigger: 'change' },
  ],
});

const loginFormRef = ref();
const authStore = useAuthStore();
const router = useRouter();

const submitForm = async () => {
  try {
    await authStore.login(loginForm.value.teacher_id, loginForm.value.password);
    router.push({ name: 'Home' });
  } catch {
    ElMessage.error('登录失败，请检查用户名和密码。');
  }
};

// 跳转到注册页面
const goToRegister = () => {
  router.push({ name: 'Register' });
};

// 跳转到忘记密码页面
const goToForgotPassword = () => {
  router.push({ name: 'ForgotPassword' });
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #e8f5fa;
}

.login-header {
  text-align: center;
  font-size: 2rem;
  margin-bottom: 20px;
}

.login-main {
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-card {
  width: 400px;
  padding: 20px;
  box-shadow: 0px 0px 20px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.additional-links {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
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
