import {createRouter, createWebHistory} from "vue-router";
import Home from "../views/Home.vue";

const router = createRouter({
        history: createWebHistory(),
        routes: [{
            path: '/',
            component: Home,
            meta: {
                title: 'Home - EEA Cash'
            },
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
