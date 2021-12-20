import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";


const routes: Array<RouteRecordRaw> = [
  { path: "/", redirect: "/login" },
  {
    path: "/login",
    name: "login",
    component: ()=> import("../views/login.vue"),
  },
  {
    path: "/index",
    name: "index",
    component: () => import("../views/index.vue"),
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: () =>import("../views/dashboard.vue"),
      },
      {
        path: "src/company",
        name: "公司管理",
        component: () => import("../views/src/company/company.vue"),
      },
      {
        path: "src/company/edit/:id",
        name: "editCompany",
        component: () => import("../views/src/company/edit.vue"),
      },
      {
        path: "src/company/new",
        name: "newCompany",
        component: () => import("../views/src/company/new.vue"),
      },
      {
        path: "src/gather",
        name: "gather",
        component: () => import("../views/src/gather.vue"),
      },
      {
        path: "src/domain",
        name: "domain",
        component: () => import("../views/src/domain.vue"),
      },
      {
        path: "src/website",
        name: "website",
        component: () => import("../views/src/website.vue"),
      },
      {
        path: "src/ips",
        name: "ip",
        component: () => import("../views/src/ips.vue"),
      },
      {
        path: "src/news",
        name: "news",
        component: () => import("../views/src/news.vue"),
      },
      {
        path: "src/littleProgram",
        name: "littleProgram",
        component: () => import("../views/src/littleProgram.vue"),
      },
      {
        path: "src/service",
        name: "service",
        component: () => import("../views/src/service.vue"),
      },
      {
        path: "src/webChatAccount",
        name: "webChatAccount",
        component: () => import("../views/src/webChatAccount.vue"),
      },
      {
        path: "vuln/poc",
        name: "poc",
        component: () => import("../views/vuln/poc.vue"),
      },
      {
        path: "knowledge/knowledge",
        name: "knowledge",
        component: () => import("../views/knowledge/knowledge.vue"),
      },
      {
        path: "knowledge/config",
        name: "config",
        component: () => import("../views/knowledge/config.vue"),
      },
      {
        path: "task/taskManage",
        name: "taskManage",
        component: () => import("../views/task/taskManage.vue"),
      }
    ],
  },
  
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  // 判断有没有登录
  const isLogin = localStorage.token ? true : false;
  // 如果当前访问的是登录页面或者注册页面可以让它进入
  if (to.path == "/login" || to.path == "/register") {
    next();
  } else {
    // 如果isLogin为true表示已经登录了;就让它正常进入就可以 ,如果没登录就让他进入登录页面
    isLogin ? next() : next("/");
  }
});

export default router;