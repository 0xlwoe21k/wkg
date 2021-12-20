<template>
  <a-row type="flex">
    <a-col :span="6">
      <a-input-search
        v-model:value="searchKey"
        placeholder="input search text"
        enter-button
        @search="btnSearch"
      />
    </a-col>
    <a-col :span="6" :order="3"></a-col>
  </a-row>
  <a-row type="flex" style="margin-top: 10px;">
    <a-col :span="24" :order="4">
      <a-table :columns="columns" :data-source="rolesList.list" :align="dbalign"></a-table>
    </a-col>
  </a-row>
</template>
<script lang="ts">
import { SmileOutlined, DownOutlined } from "@ant-design/icons-vue";
import { defineComponent, reactive, onMounted, ref } from "vue";
import axios from "axios";
import agentService  from "../../service/agent.service";

const columns = [

  {
    title: "地址",
    dataIndex: "addr",
    key: "addr",
    align: 'center',
  },
  {
    title: "AgentId",
    dataIndex: "agent_id",
    key: "agent_id",
    align: 'center',
  },
  {
    title: "主机名",
    dataIndex: "hostname",
    key: "hostname",
    align: 'center',
  },
  {
    title: "CPU",
    dataIndex: "cpu",
    key: "cpu",
    align: 'center',
  },
  {
    title: "内存",
    dataIndex: "memory",
    key: "memory",
    align: 'center',
  },
  {
    title: "源IP",
    dataIndex: "source_ip",
    key: "source_ip",
    align: 'center',
  },
  {
    title: "上线时间",
    dataIndex: "first_heartbeat_time_format",
    key: "first_heartbeat_time_format",
    align: 'center',
  }
];

interface Config {
  detail: string;
  downloadurl: Array<string>;
  name: string;
  sha256: string;
  version: string;
}

interface Plugins {
  cpu: number;
  io: number;
  name: string;
  pid: number;
  qps: number;
  rss: number;
  version: string;
}

interface Agent {
  _id: string;
  addr: string;
  agent_id: string;
  config: Config;
  config_update_time: number;
  cpu: number;
  extranet_ipv4: Array<string>;
  extranet_ipv6: Array<string>;
  first_heartbeat_time: number;
  first_heartbeat_time_format: string;
  hostname: string;
  intranet_ipv4: Array<string>;
  intranet_ipv6: Array<string>;
  io: string;
  last_heartbeat_time: number;
  last_heartbeat_time_format: string;
  memory: number;
  net_type: string;
  plugins: Plugins;
  slab: number;
  source_ip: string;
  source_port: number;
  version: string;
}

export default defineComponent({
  components: {},
  setup() {
    const rolesList: { list: Agent[] } = reactive({ list: [] });
    let searchKey = ref('')

    const btnSearch = () => {
      console.log(1)
    }
    onMounted(() => {
      console.log()

    });
    return {
      rolesList,
      searchKey,
      columns,
      dbalign: ref('center'),
      btnSearch,
    };
  },
});
</script>

<style>
.ant-table-thead > tr > th {
  text-align: center;
}
</style>