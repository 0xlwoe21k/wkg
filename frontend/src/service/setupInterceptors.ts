// import axiosInstance from "./api";
// import { notification } from 'ant-design-vue';
// import { useRouter } from "vue-router";

// const router = useRouter()

// axiosInstance.interceptors.request.use((req: any) => {
// 	if (localStorage.token) {
// 		const token = localStorage.getItem('token')
// 		if (token){
// 			router.push("/login");
// 		}
// 		req.headers['token'] = token;
// 	}
// 	return req;
// },
// 	(error) => {
// 		return Promise.reject(error);
// 	}
// );

// axiosInstance.interceptors.response.use(
// 	(response) => {
// 		return response;
// 	},
// 	// 请求错误
// 	(error) => {
// 		const { status } = error.response;
// 		if (status == 401) {
// 			notification["warning"]({
// 				message: 'Notification Title',
// 				description:
// 					'This is the content of the notification. This is the content of the notification. This is the content of the notification.',
// 			});
// 			localStorage.removeItem("token");
// 			router.push("/login");

// 		} 
// 		return Promise.reject(error);
// 	}
// );
