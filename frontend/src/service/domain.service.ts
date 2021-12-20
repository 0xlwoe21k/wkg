import axiosInstance from './api';

class DoaminService {

  getDomainInfo(page:number,pagesize:number,type:string,keyword:string) {
    const data = {"page":page,"pagesize":pagesize,"type":type,"keyword":keyword}
    return axiosInstance.post('/api/v1/getDomainInfo',data,);
  }

  getDomainInfoByCid(page:number,pagesize:number,type:string,keyword:string,cid:number) {
    const data = {"page":page,"pagesize":pagesize,"type":type,"keyword":keyword,"cid":cid}
    return axiosInstance.post('/api/v1/getDomainInfoByCid',data,);
  }

  DelDomainInfoById(id:number) {
    const data = {"id":id}
    return axiosInstance.post('/api/v1/getDomainInfo',data,);
  }

  ReadFlagDomainInfoById(id:string) {
    const data = {"id":id}
    return axiosInstance.post('/api/v1/readFlagDomainInfoById',data,);
  }

  ReadAllFlagDomainInfo(id:string) {
    const data = {"id":id}
    return axiosInstance.post('/api/v1/readAllFlagDomainInfo',data,);
  }

}

const domainService = new DoaminService()

export default domainService;