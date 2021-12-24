import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";


const routes: Array<RouteRecordRaw> = [
  { path: "/", redirect: "/login" },
  {
    path: "/login",
    name: "login",
    component: () => import("../views/login.vue"),
  },
  {
    path: "/index",
    name: "index",
    meta: { title: "首页" },
    redirect: {name: 'Dashboard'},
    component: () => import("../views/index.vue"),
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        meta: { title: "Dashboard" },
        component: () => import("../views/dashboard.vue"),
      },
      {
        path: "src",
        name: "SRC管理",
        meta: { title: "SRC管理" },
        redirect: {name: 'gather'},
        component: () => import("../views/src/src.vue"),
        children: [{
          path: "company",
          name: "公司管理",
          meta: { title: "公司管理" },
          component: () => import("../views/src/company/company.vue"),
        },
        {
          path: "company/edit/:id",
          name: "editCompany",
          meta: { title: "公司管理" },
          component: () => import("../views/src/company/edit.vue"),
        },
        {
          path: "company/new",
          name: "newCompany",
          meta: { title: "公司管理" },
          component: () => import("../views/src/company/new.vue"),
        },
        {
          path: "gather",
          name: "gather",
          meta: { title: "资产总览" },
          component: () => import("../views/src/gather.vue"),
        },
        {
          path: "domain",
          name: "domain",
          meta: { title: "域名管理" },
          component: () => import("../views/src/domain.vue"),
        },
        {
          path: "website",
          name: "website",
          meta: { title: "站点管理" },
          component: () => import("../views/src/website.vue"),
        },
        {
          path: "ips",
          name: "ip",
          meta: { title: "IP管理" },
          component: () => import("../views/src/ips.vue"),
        },
        {
          path: "news",
          name: "news",
          meta: { title: "资讯管理" },
          component: () => import("../views/src/news.vue"),
        },
        {
          path: "littleProgram",
          name: "littleProgram",
          meta: { title: "小程序管理" },
          component: () => import("../views/src/littleProgram.vue"),
        },
        {
          path: "service",
          name: "service",
          meta: { title: "服务管理" },
          component: () => import("../views/src/service.vue"),
        },
        {
          path: "webChatAccount",
          name: "webChatAccount",
          meta: { title: "微信公众号管理" },
          component: () => import("../views/src/webChatAccount.vue"),
        },]
      },
      {
        path: "vulns",
        name: "vulns",
        meta: { title: "漏洞管理" },
        redirect: {name:'pocs'},
        component: () => import("../views/vulns/vuln.vue"),
        children: [{
          path: "pocs",
          name: "pocs",
          meta: { title: "POC管理" },
          component: () => import("../views/vulns/poc.vue"),
        }],
      },
      {
        path: "knowledge",
        name: "knowledge",
        meta: { title: "知识库" },
        redirect: {name:'vuln'},
        component: () => import("../views/knowledge/knowledge.vue"),
        children:[{
          path: "vuln",
          name: "vuln",
          meta: { title: "漏洞知识库" },
          component: () => import("../views/knowledge/vuln.vue"),
        },
        {
          path: "config",
          name: "配置",
          meta: { title: "配置" },
          component: () => import("../views/knowledge/config.vue"),
        }]
      },
      {
        path: "task",
        name: "task",
        meta: { title: "任务管理" },
        component: () => import("../views/task/task.vue"),
        children:[{
          path: "taskList",
          name: "taskList",
          meta: { title: "任务列表" },
          component: () => import("../views/task/taskList.vue"),
        }]
      },
      {
        path: "onlineScan",
        name: "onlineScan",
        meta: { title: "在线扫描" },
        component: () => import("../views/scan/scan.vue"),
        children:[{
          path: "domainScan",
          name: "domainScan",
          meta: { title: "任务列表" },
          component: () => import("../views/scan/domainScan.vue"),
        },{
          path: "vulnScan",
          name: "vulnScan",
          meta: { title: "漏洞扫描" },
          component: () => import("../views/scan/vScan.vue"),
        }]
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