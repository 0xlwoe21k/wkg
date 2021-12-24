/* eslint-disable @typescript-eslint/no-explicit-any */
<template>
 
    <a-row type="flex" justify="start" :gutter="16">
      <a-col :span="6">
        <a-input v-model:value="searchKey" placeholder="公司名称、域名、SRC地址、关键字." />
      </a-col>
      <a-col :span="2">
        <a-select
          ref="select"
          v-model:value="selectVal"
          :options="options"
          style="width: 120px"
          @change="selectHandleChange"
        ></a-select>
      </a-col>
      <a-col :span="2">
        <a-button @click="OnSearch" type="primary">
          <template #icon>
            <SearchOutlined />
          </template>
          搜索
        </a-button>
      </a-col>
      <a-col :span="1">
        <a-button @click="OnNew" type="primary">
          <template #icon>
            <plus-circle-outlined />
          </template>
          新增
        </a-button>
      </a-col>
    </a-row>
    <!-- <a-row style="margin-top: 15px;" justify="space-between">
    <a-col :span="1">
      <a-button type="primary">删除全部</a-button>
    </a-col>
    <a-col :span="2">
      <a-button typeof="primary" :disabled="!hasSelected">删除选中</a-button>
    </a-col>
    <a-col :span="20"></a-col>
    </a-row>-->
    <!-- <a-row type="flex" style="margin-top: 15px;">
    <span style="margin-left: 8px">
      <template v-if="hasSelected">
        <a-alert :message="selectRowKeyLength" type="info" show-icon />
      </template>
    </span>
    </a-row>-->
    <a-row type="flex" style="margin-top: 15px">
      <a-col :span="24" :order="4">
        <a-table
          sticky
          :row-selection="{
            selectedRowKeys: selectedRowKeys,
            onChange: onSelectChange,
          }"
          :columns="columns"
          :data-source="rolesList.list"
          :align="dbalign"
          :rowKey="(record: any, index: any) => index"
          @change="handleTableChange"
          :scroll="{ x: 1500, y: 1500 }"
          :pagination="Mypagination"
        >
          <template #bodyCell="{ column, text, record }">
            <template v-if="column.key === 'operation'">
              <a-row type="flex" justify="center">
                <a-col :span="6">
                  <a-button size="small" type="primary" @click="OnEdit(record.id)">Edit</a-button>
                </a-col>
                <a-col :span="7">
                  <a-button size="small" type="primary" @click="OnScan(record.id)">Scan</a-button>
                </a-col>
                <a-col :span="7">
                  <a-popconfirm
                    title="你确定？"
                    ok-text="Yes"
                    cancel-text="No"
                    @confirm="OnDelete(record.id)"
                  >
                    <a-button size="small" type="primary" danger>Delete</a-button>
                  </a-popconfirm>
                </a-col>
              </a-row>
            </template>
            <template v-if="column.dataIndex === 'monitorStatus'">
              <div v-if="text">
                <a-tag color="#87d068">Y</a-tag>
              </div>
              <div v-else>
                <a-tag color="#f50">N</a-tag>
              </div>
            </template>
          </template>
        </a-table>
      </a-col>
    </a-row>

</template>
<script lang="ts">
import { SearchOutlined, PlusCircleOutlined } from "@ant-design/icons-vue";
import {
  defineComponent,
  reactive,
  computed,
  onMounted,
  ref,
  toRefs,
} from "vue";
import srcService from "../../../service/src.service";
import { useRouter } from "vue-router";
import types from "../../../common/types";
import { message } from "ant-design-vue";

interface Company {
  id: number;
  projectType: string;
  companyName: string;
  domain: string;
  SrcUrl: string;
  keyWord: string;
  monitorStatus: boolean;
  monitorRate: number;
  lastUpdateTime: string;
}

type Key = string | number;

export default defineComponent({
  components: { SearchOutlined, PlusCircleOutlined },
  setup() {
    const rolesList: { list: Company[] } = reactive({ list: [] });
    let searchKey = ref("");
    let total = ref(1);
    let curPage = ref(1);
    let pageSize = ref(10);
    let columns = types.getComapnyTableColumns();
    let selectVal = ref("companyName");
    let batch = ref(false);
    const router = useRouter();

    const state = reactive<{
      selectedRowKeys: Key[];
      loading: boolean;
    }>({
      selectedRowKeys: [], // Check here to configure the default column
      loading: false,
    });

    const hasSelected = computed(() => state.selectedRowKeys.length > 0);
    const options = [
      {
        value: "companyName",
        label: "公司名称",
      },
      {
        value: "domain",
        label: "域名",
      },
      {
        value: "keyWord",
        label: "关键字",
      },
    ];

    const selectHandleChange = () => {
      console.log(1);
    };
    const OnEdit = (cid: number) => {
      console.log(cid);
      router.push({ name: "editCompany", params: { id: cid } });
    };
    const OnNew = (cid: number) => {
      console.log(cid);
      router.push("/index/src/company/new");
    };

    const OnScan = (cid: number) => {
      srcService.ScanCompanyDomain(cid).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg);
        } else if (res.data.code == 200) {
          message.success(res.data.msg);
        }
      });
    };

    const OnDelete = (cid: number) => {
      srcService.DelCompanyByid(cid).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg);
        } else if (res.data.code == 200) {
          message.success(res.data.msg);
          InitData();
        }
      });
    };

    const onSelectChange = (selectedRowKeys: Key[]) => {
      console.log("selectedRowKeys changed: ", selectedRowKeys);
      state.selectedRowKeys = selectedRowKeys;
    };

    const OnSearch = () => {
      srcService
        .searchCompanyInfo(
          curPage.value,
          pageSize.value,
          selectVal.value,
          searchKey.value
        )
        .then((res: any) => {
          if (res.data.code == 400) {
            alert(res.data.msg);
          } else if (res.data.code == 200) {
            rolesList.list = JSON.parse(res.data.msg);
            total.value = res.data.total;
          }
        });
    };

    const handleTableChange = (pag: any, filters: any, sorter: any) => {
      state.selectedRowKeys = [];
      curPage.value = pag.current;
      pageSize.value = pag.pageSize;
      srcService
        .getCompanyInfo(curPage.value, pageSize.value)
        .then((res: any) => {
          if (res.data.code == 400) {
            alert(res.data.msg);
          } else if (res.data.code == 200) {
            rolesList.list = JSON.parse(res.data.msg);
            total.value = res.data.total;
          }
        });
    };

    onMounted(() => {
      InitData();
    });

    const InitData = () => {
      srcService
        .getCompanyInfo(curPage.value, pageSize.value)
        .then((res: any) => {
          if (res.data.code == 400) {
            alert(res.data.msg);
          } else if (res.data.code == 200) {
            rolesList.list = JSON.parse(res.data.msg);
            total.value = res.data.total;
          }
        });
    };

    const Mypagination = computed(() => ({
      total: total.value,
      current: curPage.value,
      pageSize: pageSize.value,
      showTotal: () => `总共 ${total.value} 项`,
      defaultPageSize: 10,
      pageSizeOptions: ["5", "10", "20", "50"], // 可不设置使用默认
      showSizeChanger: true, // 是否显示pagesize选择
      showQuickJumper: true, // 是否显示跳转窗
    }));

    return {
      rolesList,
      searchKey,
      dbalign: ref("center"),
      columns,
      options,
      hasSelected,
      batch,
      OnScan,
      OnNew,
      // selectRowKeyLength,
      ...toRefs(state),
      OnSearch,
      selectVal,
      Mypagination,
      OnEdit,
      selectHandleChange,
      OnDelete,
      onSelectChange,
      handleTableChange,
    };
  },
});
</script>

<!-- <style>
.ant-table-thead > tr > th {
  text-align: center;
}
</style> -->
