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
const courses = ref<{ teacher_id: string; room_id: string; state: number; name: string; ps: string }[]>([]);

// 获取课程数据并提取字段，支持通过 teacherID 动态加载数据
export const fetchcourses = async (teacherID?: string) => {
  const authStore = useAuthStore(); // 获取 Auth Store
  const id = teacherID || authStore.teacherId; // 优先使用传入的 teacherID，否则从 store 中获取

  if (!id) {
    console.error('No teacherID provided.');
    return;
  }

  // 动态构建 URL
  const url = `/api/course/${id}`;

  // 调用 fetchData 获取数据
  courses.value =
    (await fetchData(url, (data) =>
      // 提取所需字段
      data.map((item) => ({
        teacher_id: item.teacher_id,
        room_id: item.room_id,
        state: item.state,
        name: item.name,
        ps: item.ps
      }))
    )) || [];
};

// 导出响应式数据
export { courses };
