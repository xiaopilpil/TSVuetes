import { createWebHistory, createRouter, RouteRecordRaw } from 'vue-router'

import Home from "../views/Home/index.vue";
import Login from "../views/Login/index.vue";
import Register from "../views/Register/index.vue";
import News from '../views/Home/news.vue';
import Course from '../views/Home/course.vue'
import Acourse from '../views/Home/Acourse.vue';
import Classroom from '../views/Classroom/Classroom.vue';
import Apply from '../views/Home/Apply.vue';
import Createroom from '../views/Home/Createroom.vue';
import { useAuthStore } from '../store/auth';


const routes: RouteRecordRaw[] = [
  { path: "/", name: "Home", component: Home, children:[
    {
      path: "News", 
      name: "News", 
      component: News 
    },
    {
      path: "Course", 
      name: "Course", 
      component: Course 
    },
    {
      path: "Acourse",
      name: "Acourse",
      component: Acourse
    },
    {
      path: "Apply",
      name: "Apply",
      component: Apply
    },
    {
      path: "Createroom",
      name: "Createroom",
      component: Createroom
    }
  ]},
  { path: "/login", name: "Login", component: Login },
  { path: "/Register", name: "Register", component: Register },
  { path: "/classroom/:id", name:"Classroom", component: Classroom,meta: { requiresAuth: true }}
];

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// 添加路由守卫
router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore();

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    // 如果页面需要登录，且未登录，则跳转到登录页面
    next({ name: 'Login', query: { redirect: to.fullPath } }); // 记录重定向路径
  } else {
    next(); // 否则允许访问
  }
});

export default router