import { ref } from 'vue';
import axios from '../../axios';
import { useAuthStore } from '../../store/auth'; // 引入 auth store


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
  const classrooms = ref<{ room_id: string; state: number }[]>([]);
  
  // 获取课程数据并提取字段，支持通过 teacherID 动态加载数据
  export const fetchclassroom = async (teacherID?: string) => {
    const authStore = useAuthStore(); // 获取 Auth Store
    const id = teacherID || authStore.teacherId; // 优先使用传入的 teacherID，否则从 store 中获取
  
    if (!id) {
      console.error('No teacherID provided.');
      return;
    }
  
    // 动态构建 URL
    const url = `/api/course/${id}`;
  
    // 调用 fetchData 获取数据
    classrooms.value =
      (await fetchData(url, (data) =>
        // 提取所需字段
        data.map((item) => ({
          room_id: item.room_id,
          state: item.state,
        }))
      )) || [];
  };


  /**
 * 验证教室是否存在且状态为 1
 * @param roomId 教室的 ID
 * @returns Promise 返回教室信息（如果满足条件）
 */
export const validateApprovedClassroom = async (roomId: string): Promise<any> => {
    // 如果 classrooms 数据尚未加载，先加载数据
    if (classrooms.value.length === 0) {
      await fetchclassroom(); // 重新加载教室数据
    }
  
    // 查找符合条件的教室
    const classroom = classrooms.value.find(
      (classroom) => classroom.room_id === roomId && classroom.state === 1
    );
  
    if (!classroom) {
      console.error(`Classroom with ID ${roomId} does not exist or is not approved`);
      throw new Error(`Classroom with ID ${roomId} does not exist or is not approved`);
    }
  
    console.log('Classroom validation successful:', classroom);
    return classroom; // 返回教室信息
  };
  
// 导出响应式数据
export { classrooms };


//修改状态请求
export const updatecourse = async (teacher_id: string, room_id: string, state: number) =>{
    try{
        const response = await axios.post(`/api/course`,{
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

