import axiosInstance from './api';

class IPsService {
    GetIPsInfo(page: number, pagesize: number, type: string, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
        return axiosInstance.post('/api/v1/getIPsInfo', data,)
    }

    DelIPsInfo(id: any) {
        const data = { "id": id}
        return axiosInstance.post('/api/v1/getIPsInfo', data,)
    }
}
const ipsSvc = new IPsService()

export default ipsSvc;