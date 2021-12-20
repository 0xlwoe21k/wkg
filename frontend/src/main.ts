import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Antd from "ant-design-vue";
// import { DatePicker } from 'ant-design-vue';
// import setupInterceptors from './service/setupInterceptors';
import "ant-design-vue/dist/antd.css";


// setupInterceptors(store)

const app = createApp(App);
app.use(store);
app.use(router);
app.use(Antd);
app.mount("#app");
