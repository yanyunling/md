import { createRouter, createWebHashHistory } from "vue-router";
import Token from "@/store/token";
import Cron from "@/views/tool/cron/index.vue";
import DockerConvert from "@/views/tool/docker-convert/index.vue";
import Encryption from "@/views/tool/encryption/index.vue";
import Generate from "@/views/tool/generate/index.vue";
import ImageCompression from "@/views/tool/image-compression/index.vue";
import JsonFormat from "@/views/tool/json-format/index.vue";
import QrCode from "@/views/tool/qr-code/index.vue";
import Regular from "@/views/tool/regular/index.vue";
import Rsa from "@/views/tool/rsa/index.vue";

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
      { path: "/tool/cron", name: "cron", component: Cron },
      { path: "/tool/docker-convert", name: "docker-convert", component: DockerConvert },
      { path: "/tool/encryption", name: "encryption", component: Encryption },
      { path: "/tool/generate", name: "generate", component: Generate },
      { path: "/tool/image-compression", name: "image-compression", component: ImageCompression },
      { path: "/tool/json-format", name: "json-format", component: JsonFormat },
      { path: "/tool/qr-code", name: "qr-code", component: QrCode },
      { path: "/tool/regular", name: "regular", component: Regular },
      { path: "/tool/rsa", name: "rsa", component: Rsa },
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
