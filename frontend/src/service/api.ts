import axios from "axios";
import { useRouter } from "vue-router";

const router = useRouter()

const axiosInstance = axios.create({
  headers: {
    "Content-Type": "application/json",
  },

});


axiosInstance.interceptors.request.use((req: any) => {
  if (localStorage.token) {
    const token = localStorage.getItem('token')
    // console.log('goke', token)
    req.headers['token'] = token;
  }
  return req;
},
  (error) => {
    return Promise.reject(error);
  }
);

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  // 请求错误
  (error) => {
    if (error.response.status == 401) {
      // alert("token过期,请重新登录!");
      localStorage.removeItem("token");
      router.push("/login");
    } else {
      // alert(error.response.data);
    }
    return Promise.reject(error);
  }
);

export default axiosInstance;