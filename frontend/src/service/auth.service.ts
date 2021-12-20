import axiosInstance from './api';
import { useRouter } from "vue-router";
const router = useRouter()

class AuthService {
    login(username:string,password:string) {
        return axiosInstance.post('/api/login', { username, password }).then((res) => {
            if (res.data.code == 302){
                router.push("/login");
            }
            localStorage.setItem('token',res.data.data.token)
        });
    }

}
let authService = new AuthService()

export default {authService};