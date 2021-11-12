import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import axios from 'axios';

const app =createApp(App)

app.use(ElementPlus)
app.use(Antd);
app.use(router)
app.mount('#app')


