import { flask } from '../../axios'; // 引入 Flask Axios 实例

/**
 * 调用 Flask 的 /detect 接口上传图片并返回处理结果
 * @param file 要上传的图片文件
 * @returns Promise 返回处理后的图片数据
 */
export const detectImage = async (file: File): Promise<Blob | null> => {
  try {
    const formData = new FormData();
    formData.append('image', file); // 添加图片到 form-data 中

    const response = await flask.post('/detect', formData, {
      headers: {
        'Content-Type': 'multipart/form-data', // 确保使用 form-data
      },
      responseType: 'blob', // 设置响应类型为 blob，用于处理二进制数据
    });

    console.log('Image detection successful:', response);
    return response.data; // 返回处理后的图片数据
  } catch (error: any) {
    console.error('Error detecting image:', error.message || error);
    throw new Error('Image detection failed');
  }
};
