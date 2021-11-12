<template>
<div style="margin-left: 50px;margin-top: 20px;">
  <el-row :gutter="5">
      <el-col :span="10">
        <div style="margin-top: 15px">
            <el-input
            v-model="SearchKeyWord"
            placeholder="指定要查询的列，输入要查询的内容，支持模糊查询."
            class="input-with-select"
            >
            <template #prepend>
                <el-select v-model="SearchCol" placeholder="Select" style="width: 110px">
                <el-option label="域名" value="domain"></el-option>
                <el-option label="标题" value="title"></el-option>
                <el-option label="IP" value="ip"></el-option>
                </el-select>
            </template>
            <template #append>
                <el-button icon="el-icon-search" @click="BtnSearch"></el-button>
            </template>
            </el-input>
        </div>
    </el-col>

    </el-row>
  <el-row :gutter="0">
    <el-col :span="21">
    <el-table :data="rolesList.list" style="width: 100%">
      <el-table-column prop="id" label="Id" width="60" />
      <el-table-column prop="cid" label="Cid" width="60" />
      <el-table-column prop="domain" label="域名" width="200" />
      <el-table-column prop="title" label="标题" width="200" />
      <el-table-column prop="type" label="类型" width="50" />
      <el-table-column prop="ip" label="IP" width="120" />
      <el-table-column prop="source" label="来源" width="60" />
      <el-table-column prop="updateTime" label="更新时间" width="200" />
      <el-table-column prop="isNew" label="新增" width="60" />
      <el-table-column fixed="right" label="Operations" width="150">
        <template #default="scope">
          <!-- <el-button type="text" size="small" @click="handleClick"
            >Detail</el-button
          > -->
          <el-button type="text" size="small" @click="BtnDelet(scope.row.id)">Delet</el-button>
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
	cid:number;
	domain:string;
	title:string;
	type:string;
	ip:string;
	source:string;
	updateTime:string;	
	isNew:boolean;
}

export default  defineComponent({
    setup(){
        let Total = ref(10)
        const rolesList: { list: cmpData[] } = reactive({ list: [] });  
        let CurPage  = ref(1)
        let SearchCol = ref('')
        let SearchKeyWord = ref('')
        let PageSize = ref(10)
        const router = useRouter();

        const handleSizeChange = (value:any) =>{

        }
        const handleCurrentChange = (value:any) =>{
            let data = {"page":value,"pagesize":PageSize.value}
            axios({
                url: '/api/v1/getDomainInfo',
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

        const BtnSearch = () => {
            let data = {"type":SearchCol.value,"keyWord":SearchKeyWord.value}
            axios({
                url: '/api/v1/getDomainInfoByKey',
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
        
        const BtnDelet = (index:number) =>{
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
                url: '/api/v1/getDomainInfo',
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
            SearchKeyWord,
            rolesList,
            SearchCol,
            // CompanyData,
            handleSizeChange,
            BtnDelet,
            BtnSearch,
            handleCurrentChange,
        }

    },
})
</script>