import { createApp } from "vue";
import ElementPlus from "element-plus";
import locale from "element-plus/lib/locale/lang/zh-cn";
import "element-plus/dist/index.css";
import "./styles/index.scss";
import App from "./App.vue";
import router from "./router";
import "virtual:svg-icons-register";

const app = createApp(App);
app.use(router);
app.use(ElementPlus, { locale });
app.mount("#app");
