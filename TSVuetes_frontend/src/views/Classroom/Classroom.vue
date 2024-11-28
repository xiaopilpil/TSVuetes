<template>
    <div v-if="loading">
      <h1>Loading...</h1>
    </div>
    <div v-else-if="classroom">
      <h1>Classroom: {{ classroom.room_id }}</h1>
      <p>Status: Approved</p>
  
      <div>
        <h2>Upload and Process Image</h2>
        <input type="file" @change="handleFileChange" accept="image/*" />
        <button @click="uploadImage" :disabled="!selectedFile">Upload</button>
        <div v-if="uploading">Processing image...</div>
        <div v-if="processedImage">
          <h3>Processed Image:</h3>
          <img :src="processedImage" alt="Processed Image" />
        </div>
        <div v-if="uploadError" class="error">{{ uploadError }}</div>
      </div>
    </div>
    <div v-else-if="error">
      <h1>Error</h1>
      <p>{{ error }}</p>
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref } from 'vue';
  import { useRoute } from 'vue-router';
  import { validateApprovedClassroom } from './Classroom'; // 引入教室验证函数
  import { detectImage } from './Picture'; // 引入 Flask 的接口函数
  
  // 路由参数
  const route = useRoute();
  const roomId = route.params.id as string;
  
  // 状态管理
  const classroom = ref<{ room_id: string; state: number } | null>(null);
  const error = ref<string | null>(null);
  const loading = ref(true); // 验证加载状态
  
  // 图片上传相关状态
  const selectedFile = ref<File | null>(null);
  const processedImage = ref<string | null>(null);
  const uploading = ref(false);
  const uploadError = ref<string | null>(null);
  
  // 验证教室
  const validateClassroom = async () => {
    try {
      loading.value = true;
      classroom.value = await validateApprovedClassroom(roomId); // 验证教室
    } catch (err: any) {
      error.value = err.message || 'Failed to validate classroom.';
    } finally {
      loading.value = false;
    }
  };
  
  // 处理文件选择
  const handleFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files[0]) {
      selectedFile.value = target.files[0];
      processedImage.value = null; // 清除之前的结果
      uploadError.value = null; // 清除之前的错误
    }
  };
  
  // 上传图片并调用 Flask 接口
  const uploadImage = async () => {
    if (!selectedFile.value) return;
  
    uploading.value = true;
    uploadError.value = null;
  
    try {
      const blob = await detectImage(selectedFile.value); // 调用 Flask 接口
      if (blob) {
        processedImage.value = URL.createObjectURL(blob); // 将 Blob 转换为 URL
      } else {
        throw new Error('No image returned from server.');
      }
    } catch (err: any) {
      uploadError.value = err.message || 'Failed to process image.';
    } finally {
      uploading.value = false;
    }
  };
  
  // 页面加载时验证教室
  validateClassroom();
  </script>
  
  <style scoped>
  .error {
    color: red;
  }
  img {
    max-width: 100%;
    height: auto;
  }
  </style>
  