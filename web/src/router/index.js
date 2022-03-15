import {createRouter, createWebHistory} from "vue-router"
import Cookies from 'js-cookie'
import Home from "../views/Home.vue"
import Login from "../views/Login.vue"
import NoRoute from "../views/NoRoute.vue"
import Admin from "../views/Admin.vue"
import AdminDash from "../views/admin/AdminDash.vue"
import AdminUsers from "../views/admin/AdminUsers.vue"
import AdminBalance from "../views/admin/AdminBalance.vue"
import AdminTransfer from "../views/admin/AdminTransfer.vue"

import User from "../views/User.vue"
import UserDash from "../views/user/UserDash.vue"

const router = createRouter({
    history: createWebHistory(),
    routes: [{
        path: '',
        component: Home,
        meta: {
            title: 'Home - EEA Cash',
        },
    }, {
        path: '/login',
        component: Login,
        meta: {
            title: 'Login',
        },
    }, {
        path: '/admin',
        component: Admin,
        children: [
            {
                path: 'dashboard',
                component: AdminDash,
                meta: {
                    title: 'Dashboard',
                },
            },
            {
                path: 'users',
                component: AdminUsers,
                meta: {
                    title: 'Users',
                },
            },
            {
                path: 'balance',
                component: AdminBalance,
                meta: {
                    title: 'AdminBalance',
                },
            },
            {
                path: 'transfer',
                component: AdminTransfer,
                meta: {
                    title: 'AdminTransfer',
                },
            }
        ],
    }, {
        path: '/user',
        component: User,
        children: [
            {
                path: '',
                component: UserDash,
                meta: {
                    title: 'Dashboard',
                },
            },
        ],
    }, {
        path: '/:pathMatch(.*)*',
        component: NoRoute,
        meta: {
            title: '404 Not Found',
        }
    }],
})

function checkPermission(role) {
    let userString = Cookies.get('user_info')
    console.log("userstring:", userString)
    if (userString) {
        let user = JSON.parse(userString)
        if (role == user.role) {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }

    if (to.path == '/login') {
        if (checkPermission('admin')) {
            next('/admin')
        } else if (checkPermission('user')) {
            next('/user')
        } else {
            next()
        }
    } else if (to.path.startsWith('/admin')) {
        if (checkPermission('admin')) {
            next()
        } else {
            next('/login')
        }
    } else if (to.path.startsWith('/user')) {
        if (checkPermission('user')) {
            next()
        } else {
            next('/login')
        }
    } else {
        next()
    }
})

export default router
