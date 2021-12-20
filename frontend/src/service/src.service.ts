import { AntdIconType } from '@ant-design/icons-vue/lib/components/AntdIcon';
import axiosInstance from './api';
import axios from "axios";

class SrcService {
    getCompanyInfo(page:number,pagesize:number) {
        const data = { "page": page, "pagesize": pagesize }
        return axiosInstance.post('/api/v1/getCompanyInfo', data,)
    }

    searchCompanyInfo(page:number,pagesize:number,type:string,keyword:string){
        const data = { "page": page, "pagesize": pagesize ,"type":type,"keyword":keyword}
        return axiosInstance.post('/api/v1/getCompanyInfo', data,)
    }

    getCompanyInfoById(cid:any){
        const data = {"id":cid.value}
        return axiosInstance.post('/api/v1/getCompanyByid', data,)
    }

    UpdateComapny(cid:any,ctype:any,companyName:any,rootDomain:any,srcUrl:any,keyWord:any,monitorStauts:any,monitorRate:any){
        const data = {"id":cid.value, "projectType":ctype.value,"companyName":companyName.value,"domain":rootDomain.value,
        "srcUrl":srcUrl.value,"keyWord":keyWord.value,"monitorStatus":monitorStauts.value,
        "monitorRate":monitorRate.value,"vulnScanStatus":false,"vulnScanRate":24}

        return axiosInstance.post('/api/v1/updateCompanyInfo', data,)
    }
    
    NewComapny(ctype:any,companyName:any,rootDomain:any,srcUrl:any,keyWord:any,monitorStauts:any,monitorRate:any){
        const data ={"projectType":ctype.value,"companyName":companyName.value,"domain":rootDomain.value,
        "srcUrl":srcUrl.value,"keyWord":keyWord.value,"monitorStatus":monitorStauts.value,
        "monitorRate":monitorRate.value,"vulnScanStatus":false,"vulnScanRate":24}

        return axiosInstance.post('/api/v1/newCompany', data,)
    }

    DelCompanyByid(id:any){
        const data = {"id":id}
        return axiosInstance.post('/api/v1/delCompanyByid', data,)
    }

    ScanCompanyDomain(id:any){
        const data = {"id":id}
        return axiosInstance.post('/api/v1/scanCompanyInfo', data,)
    }

    GetSelectOption(){
        return axiosInstance.get('/api/v1/getSelectOption',)
    }


    ExportByCid(cid:number){
        const data = {"id":cid}
        return axiosInstance.post('/api/v1/export', data,{responseType:'blob'})
    }


}
const srcSvc = new SrcService()

export default srcSvc;