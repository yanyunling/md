import { createRouter, createWebHashHistory } from "vue-router";
import Token from "@/store/token";

const routes = [
  { path: "/login", name: "login", component: () => import("@/views/login/index.vue"), children: null },
  {
    path: "/",
    name: "layout",
    component: () => import("@/views/layout/index.vue"),
    redirect: "/document",
    children: [
      { path: "/document", name: "document", component: () => import("@/views/document/index.vue") },
      { path: "/picture", name: "picture", component: () => import("@/views/picture/index.vue") },
      { path: "/tool", name: "tool", component: () => import("@/views/tool/index.vue") },
      { path: "/tool/cron", name: "cron", component: () => import("@/views/tool/cron/index.vue") },
      { path: "/tool/docker-convert", name: "docker-convert", component: () => import("@/views/tool/docker-convert/index.vue") },
      { path: "/tool/encryption", name: "encryption", component: () => import("@/views/tool/encryption/index.vue") },
      { path: "/tool/generate", name: "generate", component: () => import("@/views/tool/generate/index.vue") },
      { path: "/tool/image-compression", name: "image-compression", component: () => import("@/views/tool/image-compression/index.vue") },
      { path: "/tool/json-format", name: "json-format", component: () => import("@/views/tool/json-format/index.vue") },
      { path: "/tool/qr-code", name: "qr-code", component: () => import("@/views/tool/qr-code/index.vue") },
      { path: "/tool/regular", name: "regular", component: () => import("@/views/tool/regular/index.vue") },
      { path: "/tool/rsa", name: "rsa", component: () => import("@/views/tool/rsa/index.vue") },
    ],
  },
  { path: "/open/document", name: "openDocument", component: () => import("@/views/open/doc.vue"), children: null },
  { path: "/open/publish", name: "openPublish", component: () => import("@/views/open/publish.vue"), children: null },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from) => {
  if (to.name === "openDocument" || to.name === "openPublish") {
    return;
  }
  if (!Token.getAccessToken() && to.name !== "login") {
    router.push({ name: "login" });
  }
});

export default router;
