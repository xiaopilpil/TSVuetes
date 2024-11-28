import { ref } from 'vue';
import axios from '../../axios';

// 通用获取接口方法
const fetchData = async <T>(
  url: string,
  processFn?: (data: T[]) => any
): Promise<T[] | null> => {
  try {
    const response = await axios.get<T[]>(url);
    let result = response.data;

    // 如果提供了处理函数，则对数据进行处理
    if (processFn) {
      result = processFn(result);
    }

    return result;
  } catch (error: any) {
    console.error(`Failed to fetch data from ${url}:`, error.message || error);
    return null;
  }
};

// 响应式数据存储
const courses = ref<{ teacher_id:string; room_id: string; state: number; name: string }[]>([]);

// 获取课程数据并提取字段
export const fetchcourses = async () => {
  courses.value =
    (await fetchData('/api/course', (data) =>
      // 提取所需字段
      data.map((item) => ({
        teacher_id: item.teacher_id,
        room_id: item.room_id,
        state: item.state,
        name: item.name,
      }))
    )) || [];
};

// 导出响应式数据
export { courses };

//修改状态请求
export const updatecourse = async (teacher_id: string, room_id: string, state: number) =>{
    try{
        const response = await axios.put(`/api/course`,{
            teacher_id,
            room_id,
            state,
        });
        return response.data
    } catch (error: any) {
        console.error('PUT request failed:', error.message || error);
        return null; // 失败返回 null
    }
}

//修改备注信息
export const updateps = async (teacher_id: string, room_id: string, ps: string) =>{
    try{
        const response = await axios.put(`/api/ps`,{
            teacher_id,
            room_id,
            ps,
        });
        return response.data
    } catch (error: any) {
        console.error('PUT request failed:', error.message || error);
        return null; // 失败返回 null
    }
}