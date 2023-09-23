import { Router, createRouter, createWebHashHistory } from "vue-router";

const Layout = () => import("@/layout/index.vue");

const router: Router = createRouter({
  history: createWebHashHistory(),
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
              component: () => import("@/views/channels/ZhiHu.vue"),
            },
          ],
        },
      ],
    },
  ],
});

export default router;
