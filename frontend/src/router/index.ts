import { createRouter,createWebHashHistory} from "vue-router";


const routes = [
    {path:"/",redirect:"/login"},
    {
        path:"/index",
        name:"index",
        component:() => import("../views/index.vue"),
        children:[
          
            {
                path: 'domainscan',
                name:'domainscan',
                component: () => import ('../views/scan/domainscan.vue')
            },
            {
                path: 'dashboard',
                name:'dashboard',
                component: () => import ('../views/dashboard.vue')
            },
            {
                path: 'vulnscansingle',
                name:'vulnscansingle',
                component: () => import ('../views/scan/vulnscansingle.vue')
            },
            {
                path: 'vulnscanmulti',
                name:'vulnscanmulti',
                component: () => import ('../views/scan/vulnscanmulti.vue')
            },
            {
                path: 'company',
                name:'company',
                component: () => import ('../views/assetsManage/company/company.vue'),
            },
            {
                path: 'company/edit/:id',
                name:'editcompany',
                component: () => import ('../views/assetsManage/company/editCompany.vue'),
            },
            {
                path: 'company/new',
                name:'newcompany',
                component: () => import ('../views/assetsManage/company/newCompany.vue'),
            },
            {
                path: 'domain',
                name: 'domain',
                component: () => import ('../views/assetsManage/domain/domain.vue'),
            }
        ]
    },
    {
        path:"/login",
        name:"login",
        component:() => import("../views/login.vue")
    },
    {
        path:"/scan",
        name:"scan",
        component:() => import("../views/login.vue")
    }
]


const router = createRouter({
    history:createWebHashHistory(),
    routes:routes
})

export default router