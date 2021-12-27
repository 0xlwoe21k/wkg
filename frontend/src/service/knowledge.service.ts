import axiosInstance from './api';
import { message } from "ant-design-vue";


class KnowledgeService {

  getTopCategories() {
    return axiosInstance.get('/api/v1/getTopCategories',);
  }

  getSecondCategories(key: string) {
    return axiosInstance.get('/api/v1/getSecondCategories?key=' + key,);
  }

  getKnowledgeCategories(key: string) {
    return axiosInstance.get('/api/v1/getKnowledgeCategories?key=' + key,);
  }

  getKnowledge(key: string) {
    return axiosInstance.get('/api/v1/getKnowledge?key=' + key,);
  }


  getTopSelectOption() {
    return axiosInstance.get('/api/v1/getTopSelectOption',);
  }

  getSecondSelectOption(key: string) {
    return axiosInstance.get('/api/v1/getSecodSelectOption?key=' + key,);
  }


  getSummary() {
    return axiosInstance.get('/api/v1/getSummary')
  }

  getTree() {
    return axiosInstance.get('/api/v1/getTree')
  }

  SaveEditKnowledge(topValue: string, secondValue: string, title: string, content: string, key: string) {
    const data = { "topValue": topValue, "secondValue": secondValue, "title": title, "content": content, "key": key }
    return axiosInstance.post('/api/v1/saveEditKnowledge', data,)
  }

  SaveNewKnowledge(topValue: string, secondValue: string, title: string, content: string) {
    const data = { "topValue": topValue, "secondValue": secondValue, "title": title, "content": content }
    return axiosInstance.post('/api/v1/saveNewKnowledge', data,)
  }


  AddTopNode(topNode: string) {
    return axiosInstance.get('/api/v1/addTopNode?topNode=' + topNode).then((res: any) => {
      if (res.data.code == 400) {
        message.error(res.data.msg);
      } else if (res.data.code == 200) {
        message.success(res.data.msg);
      }
    })
  }

  AddSecondNode(parentKey: string, topNode: string) {
    axiosInstance.get('/api/v1/addSecondNode?topNode=' + topNode + "&parentKey=" + parentKey).then((res: any) => {
      if (res.data.code == 400) {
        message.error(res.data.msg);
      } else if (res.data.code == 200) {
        message.success(res.data.msg);
      }
    })
  }


  DelTreeNode(isLeaf:boolean,key: string) {
    const data ={"isLeaf":isLeaf,"key":key}
    axiosInstance.post('/api/v1/delTreeNode' ,data,).then((res: any) => {
      if (res.data.code == 400) {
        message.error(res.data.msg);
      } else if (res.data.code == 200) {
        message.success(res.data.msg);
      }
    })

  }
}

const knowledgeService = new KnowledgeService()

export default knowledgeService;