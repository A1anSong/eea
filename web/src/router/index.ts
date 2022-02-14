import {createRouter, createWebHistory} from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import NoRoute from "../views/NoRoute.vue";
import UserHome from "../views/user/Home.vue";
import AdminHome from "../views/admin/Home.vue";
import AdminDashboard from "../components/admin/Dashboard.vue";

const router = createRouter({
        history: createWebHistory(),
        routes: [{
            path: '/',
            component: Home,
            meta: {
                title: 'Home - EEA Cash'
            }
        }, {
            path: '/login',
            component: Login,
            meta: {
                title: 'Login'
            }
        }, {
            path: '/user',
            component: UserHome,
            children: []
        }, {
            path: '/admin',
            component: AdminHome,
            children: [{
                path: '',
                component: AdminDashboard
            }]
        }, {
            path: '/:pathMatch(.*)*',
            component: NoRoute,
            meta: {
                title: '404 Not Found'
            }
        }]
    }
)

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        // @ts-ignore
        document.title = to.meta.title
    }
    next()
})

export default router
