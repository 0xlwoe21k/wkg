import axiosInstance from './api';

class AgentService {
[x: string]: any;
  getAgentStatus() {
    return axiosInstance.get('/api/v1/agent/getStatus',);
  }

}

const agentService = new AgentService()

export default agentService;