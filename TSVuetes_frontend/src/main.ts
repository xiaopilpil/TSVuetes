import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia';
import { useAuthStore } from './store/auth';



const app = createApp(App)

app.use(ElementPlus)
app.use(createPinia());
app.use(router)

// 确保应用初始化时同步状态
const authStore = useAuthStore();
authStore.syncStateWithLocalStorage();

app.mount('#app')