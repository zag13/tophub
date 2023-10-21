import {createRouter, createWebHistory, Router} from "vue-router";

const Layout = () => import("@/layout/index.vue");

const router: Router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            name: "home",
            component: Layout,
            redirect: "/c",
            children: [
                {
                    path: "/c",
                    name: "c",
                    component: () => import("@/views/channels/index.vue"),
                    redirect: "/c/zhihu",
                    children: [
                        {
                            path: "/c/zhihu",
                            name: "c-zhizhu",
                            component: () => import("@/views/channels/index.vue"),
                        },
                        {
                            path: "/c/weibo",
                            name: "c-weibo",
                            component: () => import("@/views/channels/index.vue"),
                        },
                        {
                            path: "/c/hacker-news",
                            name: "c-hacker-news",
                            component: () => import("@/views/channels/index.vue"),
                        },
                    ],
                },
            ],
        },
        {
            path: '/:pathMatch(.*)*',
            redirect: "/",
        },
    ],
});

export default router;
