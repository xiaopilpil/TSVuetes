import axios from '../../axios';

// 创建教室 API 方法
export const createroom = async (room_id: string, device_id: string) => {
  try {
    const response = await axios.post(`/api/classroom`, {
      room_id,
      device_id,
    });
    return response.data;
  } catch (error: any) {
    console.error('POST 请求失败:', error.message || error);
    throw error; // 抛出错误，供前端捕获和处理
  }
};
