import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import axios from '../axios';

export const useAuthStore = defineStore('auth', () => {
  // 状态
  const token = ref<string | null>(null);
  const teacherId = ref<string | null>(null);
  const positionId = ref<number>(0);

  // 在应用加载时同步状态
  const syncStateWithLocalStorage = () => {
    token.value = localStorage.getItem('token');
    teacherId.value = localStorage.getItem('teacherId');
    positionId.value = parseInt(localStorage.getItem('positionId') || '0', 10);
  };

  // 初始化时同步状态
  syncStateWithLocalStorage();

  // 是否已登录
  const isAuthenticated = computed(() => !!token.value);

  // 登录
  const login = async (teacher_id: string, password: string) => {
    try {
      // 调用登录接口
      const loginResponse = await axios.post('/auth/login', { teacher_id, password });

      token.value = loginResponse.data.token;
      teacherId.value = teacher_id;
      localStorage.setItem('token', token.value || '');
      localStorage.setItem('teacherId', teacherId.value || '');

      // 获取用户信息
      const userInfoResponse = await axios.post('/api/userinf', { teacher_id });
      positionId.value = userInfoResponse.data.position || 0;
      localStorage.setItem('positionId', String(positionId.value));
    } catch (error: any) {
      throw new Error(`登录失败: ${error.response?.data?.error || error.message}`);
    }
  };

  // 注册
  const register = async (
    teacher_id: string,
    password: string,
    name: string,
    gender: string,
    faculty: string,
    position: number
  ) => {
    try {
      const registerResponse = await axios.post('/auth/register', {
        teacher_id,
        password,
        name,
        gender,
        faculty,
        position,
      });

      token.value = registerResponse.data.token;
      teacherId.value = teacher_id;
      positionId.value = position;

      localStorage.setItem('token', token.value || '');
      localStorage.setItem('teacherId', teacherId.value || '');
      localStorage.setItem('positionId', String(positionId.value));
    } catch (error: any) {
      throw new Error(`注册失败: ${error.response?.data?.error || error.message}`);
    }
  };

  // 注销
  const logout = () => {
    token.value = null;
    teacherId.value = null;
    positionId.value = 0;

    localStorage.removeItem('token');
    localStorage.removeItem('teacherId');
    localStorage.removeItem('positionId');
  };

  return {
    token,
    teacherId,
    positionId,
    isAuthenticated,
    login,
    register,
    logout,
    syncStateWithLocalStorage, // 如果需要手动同步，可以暴露这个方法
  };
});
