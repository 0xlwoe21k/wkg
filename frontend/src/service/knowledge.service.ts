import axiosInstance from './api';

class KnowledgeService {

  getTopCategories() {
    return axiosInstance.get('/api/v1/getTopCategories',);
  }

  getSecondCategories(key:string) {
    return axiosInstance.get('/api/v1/getSecondCategories?key='+key,);
  }

  getKnowledgeCategories(key:string) {
    return axiosInstance.get('/api/v1/getKnowledgeCategories?key='+key,);
  }

  getKnowledge(key:string) {
    return axiosInstance.get('/api/v1/getKnowledge?key='+key,);
  }
  

  getTopSelectOption(){
    return axiosInstance.get('/api/v1/getTopSelectOption',);
  }

  getSecondSelectOption(key:string){
    return axiosInstance.get('/api/v1/getSecodSelectOption?key='+key,);
  }


  getSummary(){
    return axiosInstance.get('/api/v1/getSummary')
  }

  getTree(){
    return axiosInstance.get('/api/v1/getTree')
  }

  SaveEditKnowledge(topValue:string,secondValue:string,title:string,content:string,key:string){
    const data = {"topValue":topValue, "secondValue":secondValue,"title":title,"content":content,"key":key}
    return axiosInstance.post('/api/v1/saveEditKnowledge',data,)
  }

  SaveNewKnowledge(topValue:string,secondValue:string,title:string,content:string){
    const data = {"topValue":topValue, "secondValue":secondValue,"title":title,"content":content}
    return axiosInstance.post('/api/v1/saveNewKnowledge',data,)
  }
  
}

const knowledgeService = new KnowledgeService()

export default knowledgeService;