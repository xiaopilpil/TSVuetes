import axios from '../../axios';

//申请教室
export const applycourse = async (teacher_id: string, room_id: string, name: string, Faculty: string, Snumber: number) =>{
    try{
        const response = await axios.post(`/api/course`,{
            teacher_id,
            room_id,
            name,
            Faculty,
            Snumber,
        });
        return response.data
    } catch (error: any) {
        console.error('POST request failed:', error.message || error);
        return null; // 失败返回 null
    }
}