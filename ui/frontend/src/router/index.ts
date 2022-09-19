import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
 
// 2. 配置路由
const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    component: () => import("../components/view/Home.vue"),
    name: 'main'
  },
  {
    path: "/home",
    component: () => import("../components/view/Home.vue"),
    name: 'home'
  },
  {
    path: "/about",
    component: () => import("../components/view/About.vue"),
    name: 'about'
  },
  {
    path: "/download",
    component: () => import("../components/view/Download.vue"),
    name: 'download'
  },
  
];
// 1.返回一个 router 实列，为函数，里面有配置项（对象） history
const router = createRouter({
  history: createWebHistory(),
  routes,
});
 
// 3导出路由   然后去 main.ts 注册 router.ts
export default router