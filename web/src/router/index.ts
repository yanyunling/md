import { createRouter, createWebHashHistory } from "vue-router";
import Layout from "@/views/layout/index.vue";
import Login from "@/views/login/index.vue";
import Document from "@/views/document/index.vue";
import Picture from "@/views/picture/index.vue";
import Tool from "@/views/tool/index.vue";
import OpenDocument from "@/views/open/doc.vue";
import Token from "@/store/token";

const routes = [
  { path: "/login", name: "login", component: Login },
  {
    path: "/",
    name: "layout",
    component: Layout,
    redirect: "/document",
    children: [
      { path: "/document", name: "document", component: Document },
      { path: "/picture", name: "picture", component: Picture },
      { path: "/tool", name: "tool", component: Tool },
    ],
  },
  { path: "/open/document", name: "openDocument", component: OpenDocument },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from) => {
  if (to.name === "openDocument") {
    return;
  }
  if (!Token.getAccessToken() && to.name !== "login") {
    router.push({ name: "login" });
  }
});

export default router;
