<template>
    <a-row type="flex" justify="start" :gutter="5">
        <a-col :span="6">
            <a-input v-model:value="searchKey" placeholder="公司名称、域名、SRC地址、关键字." />
        </a-col>
        <a-col v-if="typeSelectenable" :span="2">
            <a-select
                ref="select"
                v-model:value="typeSelectVal"
                :options="typeOptions.list"
                style="width: 120px"
                @change="selectHandleChange"
            ></a-select>
        </a-col>

        <a-col :span="2">
            <a-select
                ref="select"
                v-model:value="selectVal"
                :options="selectOptions.list"
                style="width: 150px"
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
        <a-col :span="2">
            <a-button @click="OnExport" type="primary">
                <template #icon>
                    <plus-circle-outlined />
                </template>
                资产导出
            </a-button>
        </a-col>
        <a-col :span="1">
            <a-button @click="OnImport" type="primary">
                <template #icon>
                    <plus-circle-outlined />
                </template>
                资产导入
            </a-button>
        </a-col>
    </a-row>

    <a-row type="flex" style="margin-top: 10px;">
        <a-col :span="24" :order="4">
            <a-tabs v-model:activeKey="activeKey" @change="tableChange">
                <a-tab-pane :forceRender="false" key="1" tab="域名">
                    <a-table
                        sticky
                        :columns="domaincolumns"
                        :data-source="DomainList.list"
                        :align="dbalign"
                        :rowKey="(record: any, index: any) => index"
                        @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1000 }"
                        :pagination="Mypagination"
                    >

                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                            <template v-if="column.key === 'operation'"></template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="2" tab="站点" force-render>
                    <a-table
                        sticky
                        :columns="websitecolumns"
                        :data-source="WebSiteList.list"
                        :align="dbalign"
                        :rowKey="(record: any, index: any) => index"
                        @change="handleTableChange"
                        :scroll="{ x: 1500, y: 500 }"
                        :pagination="Mypagination"
                    >
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="3" tab="IP">
                    <a-table
                        sticky
                        :columns="ipcolumns"
                        :data-source="IPsList.list"
                        :align="dbalign"
                        :rowKey="(record: any, index: any) => index"
                        @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }"
                        :pagination="Mypagination"
                    >
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="4" tab="服务">
                    <a-table
                        sticky
                        :columns="columns"
                        :data-source="ServiceList.list"
                        :align="dbalign"
                        :rowKey="(record: any, index: any) => index"
                        @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }"
                        :pagination="Mypagination"
                    >
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="5" tab="小程序">
                    <a-table
                        sticky
                        :columns="columns"
                        :data-source="LittleProgramList.list"
                        :align="dbalign"
                        :rowKey="(record: any, index: any) => index"
                        @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }"
                        :pagination="Mypagination"
                    >
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                            <template v-if="column.key === 'operation'"></template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="6" tab="微信公众号">
                    <a-table
                        sticky
                        :columns="columns"
                        :data-source="WebChatAccountList.list"
                        :align="dbalign"
                        :rowKey="(record: any, index: any) => index"
                        @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }"
                        :pagination="Mypagination"
                    >
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="7" tab="资讯">
                    <a-table
                        sticky
                        :columns="columns"
                        :data-source="NewsList.list"
                        :align="dbalign"
                        :rowKey="(record: any, index: any) => index"
                        @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }"
                        :pagination="Mypagination"
                    >
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
            </a-tabs>
        </a-col>
    </a-row>
</template>
<script lang="ts">
import { EditTwoTone, SearchOutlined, PlusCircleOutlined } from '@ant-design/icons-vue';
import { defineComponent, reactive, computed, onMounted, ref, toRefs } from "vue";
import srcService from '../../service/src.service';
import { useRouter } from 'vue-router';
import types from "../../common/types"
import { message } from 'ant-design-vue';
import domainService from '@/service/domain.service';
import websiteService from '@/service/website.service';
import ipsSvc from '@/service/ips.service';
import myoption from '../../common/options'
import axios from "axios";

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
interface option {
    value: string,
    label: string
}

type Key = string | number;

export default defineComponent({
    components: { PlusCircleOutlined, SearchOutlined },
    setup() {
        const DomainList: { list: Company[] } = reactive({ list: [] });
        const WebSiteList: { list: Company[] } = reactive({ list: [] });
        const IPsList: { list: Company[] } = reactive({ list: [] });
        const ServiceList: { list: Company[] } = reactive({ list: [] });
        const LittleProgramList: { list: Company[] } = reactive({ list: [] });
        const NewsList: { list: Company[] } = reactive({ list: [] });
        const WebChatAccountList: { list: Company[] } = reactive({ list: [] });

        let typeSelectVal = ref('')
        let typeSelectenable = ref(false)
        let searchKey = ref('')
        let total = ref(1)
        let curPage = ref(1)
        let pageSize = ref(10)
        let domaincolumns = types.getGatherDomainTableColumns()
        let websitecolumns = types.getGatherWebsiteTableColumns()
        let ipcolumns = types.getGatherIPsTableColumns()
        let columns = [{}]
        let selectVal = ref('')
        let activeKey = ref('1')
        const router = useRouter();
        const selectOptions: { list: option[] } = reactive({ list: [] })
        const typeOptions: { list: option[] } = reactive({ list: [] })
        const websiteOptions = [
            {
                value: 'ips',
                label: 'ips',
            },
            {
                value: 'website',
                label: '网站',
            },
            {
                value: 'title',
                label: '标题',
            },
            {
                value: 'favicon',
                label: 'favicon',
            }
        ];


        const OnExport = () => {
            srcService.ExportByCid(parseInt(selectVal.value)).then((res: any) => {
                //获取文件名，需要在后端进行配置
                let filename = res.headers['content-disposition'].split('filename=')[1].split('; filename')[0]

                let type = res.headers['content-type'].split(';')[0]
                let blob = new Blob([res.data], { type: type })
                const a = document.createElement('a')

                // 创建URL
                const blobUrl = window.URL.createObjectURL(blob)
                a.download = filename
                a.href = blobUrl
                console.log(a)
                document.body.appendChild(a)
                // 下载文件
                a.click()
                // 释放内存
                URL.revokeObjectURL(blobUrl)
                document.body.removeChild(a)

                // if (!res) {
                //     return
                // }
                // srcService.ex()
                // const filename = res.headers['content-disposition'].split('filename=')[1].split('; filename')[0]
                // const url = window.URL.createObjectURL(res.data)
                // const link = document.createElement('a')
                // link.style.display = 'none'
                // link.href = url
                // link.setAttribute('download', '导出' + filename)
                // document.body.appendChild(link)
                // link.click()



            })
        }


        const OnImport = (cid: number) => {
            srcService.searchCompanyInfo(curPage.value, pageSize.value, selectVal.value, searchKey.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    // rolesList.list = JSON.parse(res.data.msg)
                    message.success(res.data.msg)
                }
            })
        }


        onMounted(() => {
            srcService.GetSelectOption().then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    selectOptions.list = JSON.parse(res.data.msg)
                    selectVal.value = selectOptions.list[0].value
                    // console.log(selectOptions)
                    domainService.getDomainInfoByCid(curPage.value, pageSize.value, "domain", searchKey.value, parseInt(selectVal.value)).then((res: any) => {
                        if (res.data.code == 400) {
                            message.error(res.data.msg)
                        } else if (res.data.code == 200) {
                            DomainList.list = JSON.parse(res.data.msg)
                            total.value = res.data.total
                        }
                    })
                }
            })


        })

        const getDomainInfoByCid = () => {
            domainService.getDomainInfoByCid(curPage.value, pageSize.value, "domain", searchKey.value, parseInt(selectVal.value)).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    DomainList.list = JSON.parse(res.data.msg)
                    total.value = res.data.total
                }
            })
        }

        const getWebsiteInfoByCid = () => {
            websiteService.GetWebSiteInfoByCid(curPage.value, pageSize.value, typeSelectVal.value, searchKey.value, parseInt(selectVal.value)).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    WebSiteList.list = JSON.parse(res.data.msg)
                    total.value = res.data.total
                }
            })
        }

        const getIPsInfoByCid = () => {
            ipsSvc.GetIPsInfo(curPage.value, pageSize.value, "", "").then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    IPsList.list = JSON.parse(res.data.msg)
                    total.value = res.data.total
                }
            })
        }


        const UpdateData = () => {
            switch (activeKey.value) {
                case '1'://域名
                    getDomainInfoByCid()
                    break;
                case '2'://website
                    getWebsiteInfoByCid()
                    break;
                case '3'://IPS
                    getIPsInfoByCid()
                    break;
                case '4':
                case '5':
                case '6':
                case '7':

            }
        }

        //tabls切换事件
        const tableChange = (activeKey: string) => {
            curPage.value = 1
            pageSize.value = 10
            switch (activeKey) {
                case '1':
                    typeSelectenable.value = false
                    break;
                case '2':
                    typeSelectVal.value = websiteOptions[2].value
                    typeOptions.list = websiteOptions
                    typeSelectenable.value = true
                    break;
            }
            UpdateData()
        }

        const OnSearch = () => {
            curPage.value = 1
            pageSize.value = 10
            UpdateData()
        }

        //table 翻页切换事件
        const handleTableChange = (pag: any, filters: any, sorter: any) => {
            console.log('current:' + pag.current + ' pag.pageSize:' + pag.pageSize)
            curPage.value = pag.current;
            pageSize.value = pag.pageSize;
            UpdateData()
            // srcService.getCompanyInfo(curPage.value, pageSize.value).then((res: any) => {
            //     if (res.data.code == 400) {
            //         alert(res.data.msg)
            //     } else if (res.data.code == 200) {
            //         // rolesList.list = JSON.parse(res.data.msg)
            //         total.value = res.data.total
            //     }
            // })
        }

        //select切换事件
        const selectHandleChange = () => {
            curPage.value = 1
            pageSize.value = 10
            UpdateData()
        }




        const Mypagination = computed(() => ({
            total: total.value,
            current: curPage.value,
            pageSize: pageSize.value,
            showTotal: () => `总共 ${total.value} 项`,
            defaultPageSize: 10,
            pageSizeOptions: ['5', '10', '20', '20'], // 可不设置使用默认
            showSizeChanger: true, // 是否显示pagesize选择
            showQuickJumper: true, // 是否显示跳转窗
        }));

        return {
            DomainList,
            WebSiteList,
            IPsList,
            ServiceList,
            LittleProgramList,
            NewsList,
            WebChatAccountList,
            searchKey,
            dbalign: ref('center'),
            domaincolumns,
            websitecolumns,
            ipcolumns,
            columns,
            typeSelectenable,
            typeOptions,
            typeSelectVal,
            activeKey,
            selectOptions,
            selectVal,
            Mypagination,
            OnSearch,
            OnExport,
            OnImport,
            tableChange,
            selectHandleChange,
            handleTableChange,
        };
    },
});
</script>
