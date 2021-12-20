import axiosInstance from './api';

class WebsiteService {
  GetWebSiteInfo(page: number, pagesize: number,type:string,keyword:string) {
    const data = { "page": page, "pagesize": pagesize,"type":type,"keyword":keyword}
    return axiosInstance.post('/api/v1/getWebSiteInfo', data,)
  }

  GetWebSiteInfoByCid(page:number,pagesize:number,type:string,keyword:string,cid:number) {
    const data = {"page":page,"pagesize":pagesize,"type":type,"keyword":keyword,"cid":cid}
    return axiosInstance.post('/api/v1/getWebSiteInfoByCid', data,)
  }

  DelWebSiteInfo(id:any) {
    const data = {"id":id}
    return axiosInstance.post('/api/v1/DelWebSiteInfoById', data,)
  }

  ScanNew() {
    return axiosInstance.get('/api/v1/scanNew',)
  }


  
}

const websiteService = new WebsiteService()

export default websiteService;