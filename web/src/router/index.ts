import { createRouter, createWebHashHistory } from "vue-router";
import Layout from "@/views/layout/index.vue";
import Login from "@/views/login/index.vue";
import Token from "@/store/token";

const routes = [
  { path: "/login", name: "login", component: Login },
  {
    path: "/",
    name: "layout",
    component: Layout,
    children: [{ path: "/user", name: "user", component: () => import("@/views/user/index.vue") }],
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from) => {
  if (!Token.getToken() && to.name !== "login") {
    router.push({ name: "login" });
  }
});

export default router;
