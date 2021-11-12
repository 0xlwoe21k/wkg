<template>
<div style="margin-left: 50px;margin-top: 20px;">
  <el-row :gutter="5">
      <el-col :span="1">
        <el-button type="primary" round @click="BtnNew">新建</el-button>
      </el-col>
          <el-col :span="4">
        <el-button type="primary" round @click="BtnFlush">刷新</el-button>
      </el-col>
    </el-row>
  <el-row :gutter="20">
    <el-col :span="20">
    <el-table :data="rolesList.list" style="width: 100%">
      <el-table-column prop="id" label="Id" width="60" />
      <el-table-column prop="projectType" label="类型" width="60" />
      <el-table-column prop="companyName" label="公司名称" width="120" />
      <el-table-column prop="domain" label="域名" width="120" />
      <el-table-column prop="srcUrl" label="SRC首页" width="200" />
      <el-table-column prop="keyWord" label="关键字" width="120" />
      <el-table-column prop="monitorStatus" label="监控状态" width="60" />
      <el-table-column prop="monitorRate" label="监控频率" width="60" />
      <el-table-column prop="vulnScanStatus" label="漏扫状态" width="60" />
      <el-table-column prop="vulnScanRate" label="漏扫状态" width="60" />
      <el-table-column prop="lastUpdateTime" label="创建时间" width="200" />
      <el-table-column fixed="right" label="Operations" width="200">
        <template #default="scope">
          <!-- <el-button type="text" size="small" @click="handleClick"
            >Detail</el-button
          > -->
          <el-button type="text" size="small" @click="BtnEdit(scope.row.id)">Edit</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div >
      <el-pagination
        v-model:currentPage="CurPage"
        :page-size="PageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="Total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :stripe="true"
        :border="true"
      >
      </el-pagination>
      </div>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent,onMounted,ref,reactive } from "vue-demi"
import { useRouter } from 'vue-router';
import axios from 'axios';

interface cmpData {
  id:number;
	projectType:string;
	companyName:string;
	domain:string;
	SrcUrl:string;
	keyWord:string;
	monitorStatus:boolean;
	monitorRate:number;	
	vulnScanStatus:boolean;
	vulnScanRate:number;
	lastUpdateTime:string;
}

export default  defineComponent({
    setup(){
        let Total = ref(10)
        // const dt:cmpData = null
        // let datas = reactive({
        //     tableData:[]
        // })
        const rolesList: { list: cmpData[] } = reactive({ list: [] });  
        // let CompanyData:cmpData[] = ref()
        let CurPage  = ref(1)
        let PageSize = ref(10)
        const router = useRouter();

        const handleSizeChange = (value:any) =>{

        }
        const handleCurrentChange = (value:any) =>{
            // alert(123)
            let data = {"page":value,"pagesize":PageSize.value}
            axios({
                url: '/api/v1/getCompanyInfo',
                method: 'POST',
                data: data,
            }).then((res) => {
                console.log(res)
                if (res.data.code == 400){
                    alert(res.data.msg)
                }else if (res.data.code == 200){
                    console.log(res.data.msg)
                    rolesList.list = JSON.parse(res.data.msg)
                    Total.value = res.data.total
                }
            });
        }
        const BtnNew = () => {
            router.push({ path: '/index/company/new' });
        }
        
        const BtnEdit = (index:number) =>{
            console.log(index)
            router.push({ name: 'editcompany',params:{id:index} });
        }
        const BtnFlush = () => {
            location.reload()
        }
        onMounted(()=>{
            // alert(123)
          InitData()
        })
        const InitData =()=>{
            let data = {"page":CurPage.value,"pagesize":PageSize.value}
            axios({
                url: '/api/v1/getCompanyInfo',
                method: 'POST',
                data: data,
            }).then((res) => {
                if (res.data.code == 400){
                    alert(res.data.msg)
                }else if (res.data.code == 200){
                    // console.log(res.data.msg)
                    rolesList.list = JSON.parse(res.data.msg)
                    Total.value = res.data.total
                }else if (res.data.code == 302){
                    router.push({ path: '/login' })
                }
            });
        }
        return{
            Total,
            CurPage,
            PageSize,
            rolesList,
            BtnFlush,
            // CompanyData,
            handleSizeChange,
            BtnEdit,
            BtnNew,
            handleCurrentChange,
        }

    },
})
</script>