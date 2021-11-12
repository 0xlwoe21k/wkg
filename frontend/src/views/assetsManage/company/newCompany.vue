<template>
 <div style="margin-top: 30px;">
  <el-row :gutter="0">
      <el-col :span="4">
        <span>类型：</span>
      </el-col>
      <el-col :span="3">
        <el-select v-model="ctype" filterable placeholder="Select">
            <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
            >   
        </el-option>
    </el-select>
      </el-col>
  </el-row>
    <el-row :gutter="2">
      <el-col :span="4">
        <span>公司名称：</span>
      </el-col>
      <el-col :span="10">
        <el-input
          v-model="companyName"
          placeholder="Please input"
        />
      </el-col>
  </el-row>
    <el-row :gutter="20">
      <el-col :span="4">
        <span>根域名：</span>
      </el-col>
      <el-col :span="10">
        <el-input
          :autosize="{ minRows: 2, maxRows: 6 }"
          type="textarea"
          v-model="rootDomain"
          placeholder="Please input"
        />
      </el-col>
  </el-row>
    <el-row :gutter="20">
      <el-col :span="4">
        <span>SRC首页：</span>
      </el-col>
      <el-col :span="10">
        <el-input v-model="srcUrl" placeholder="Please input">
            <template #prepend>Http://</template>
        </el-input>
      </el-col>
  </el-row>
  <el-row :gutter="20">
      <el-col :span="4">
        <span>关键字：</span>
      </el-col>
      <el-col :span="10">
        <el-input
          :autosize="{ minRows: 2, maxRows: 6 }"
          v-model="keyWord"
          type="textarea"
          placeholder="Please input"
        />
      </el-col>
  </el-row>
  <el-row :gutter="20">
      <el-col :span="4">
        <span>监控状态：</span>
      </el-col>
      <el-col :span="3">
        <el-switch
            v-model="monitorStauts"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="打开"
            inactive-text="关闭"
            :active-value="true"
            :width="50"
            :inactive-value="false"
            @change="switchTest"
        />
      </el-col>
  </el-row>
  <el-row :gutter="20">
      <el-col :span="4">
        <span>监控频率：</span>
      </el-col>
      <el-col :span="10">
        <el-input-number
          v-model="monitorRate"
          placeholder="小时为单位"
        />
      </el-col>
  </el-row>
  <el-row :gutter="20">
      <el-col :span="4">
        <span>漏扫状态：</span>
      </el-col>
      <el-col :span="3">
        <el-switch
            v-model="vulnScanStatus"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="打开"
            inactive-text="关闭"
            :width="50"
            :active-value="true"
            :inactive-value="false"
            @change="switchTest"
        />
      </el-col>
  </el-row>
    <el-row :gutter="20">
      <el-col :span="4">
        <span>漏扫频率：</span>
      </el-col>
      <el-col :span="1">
        <el-input-number
          v-model="vulnScanRate"
          placeholder="小时为单位"
        />
      </el-col>
  </el-row>
  <el-row :gutter="20">
      <el-col :span="1">
    <el-button type="primary" @click="BtnNewComapny">添加</el-button>
      </el-col>
      <el-col :span="6">
        <el-button type="warning" @click="BtnCancel">取消</el-button>
      </el-col>
  </el-row>
  </div>
</template>

<script lang="ts">
import axios from "axios";
import { defineComponent, onMounted,ref } from "vue-demi";

import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus'

export default defineComponent({
    setup(){
        const router = useRouter()
        const options = ref([{
            value:"SRC",
            label:"SRC"
            },
            {
            value:"公益",
            label:"公益"
            },
            {
            value:"工作",
            label:"工作"
            },
            {
            value:"挖着玩",
            label:"挖着玩"
            }
        ])
        let ctype = ref('')
        let companyName = ref('')
        let rootDomain = ref('')
        let srcUrl = ref('')
        let keyWord = ref('')
        let monitorStauts = ref(true)
        let monitorRate = ref(24)
        let vulnScanStatus = ref(true)
        let vulnScanRate = ref(24)


        const BtnNewComapny = () =>{
            const data ={"projectType":ctype.value,"companyName":companyName.value,"domain":rootDomain.value,
                         "srcUrl":srcUrl.value,"keyWord":keyWord.value,"monitorStatus":monitorStauts.value,
                         "monitorRate":monitorRate.value,"vulnScanStatus":vulnScanStatus.value,"vulnScanRate":vulnScanRate.value}
            axios({
                url: '/api/v1/newCompany',
                method: 'POST',
                data: data,
            }).then((res) => {
                if (res.data.code == 400){
                    ElMessage({
                        message: res.data.msg,
                        type: 'error',
                    })
                }else if (res.data.code == 200){
                     ElMessage({
                        message: res.data.msg,
                        type: 'success',
                    })
                    router.push({path:"/index/company"})
                }else if (res.data.code == 302){
                    router.push({ path: '/login' })
                }
            });

        }
        const BtnCancel = () =>{
            router.push({path:"/index/company"})

        }

        onMounted (() =>{
            
        })
        const switchTest = (value:any) =>{
            console.log(value)
        }
        // router.getRoutes().
        return{
          ctype,
          companyName,
          srcUrl,
          keyWord,
          vulnScanStatus,
          vulnScanRate,
          rootDomain,
          monitorRate,
          options,
          monitorStauts,
          switchTest,
          BtnCancel,
          BtnNewComapny,
        }
    }
})

</script>


<style lang="scss">
.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}
</style>