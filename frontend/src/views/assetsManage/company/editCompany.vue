<template>
 <div style="margin-top: 30px;">
  <el-row :gutter="0">
      <el-col :span="4">
        <span>id：</span>
      </el-col>
      <el-col :span="2">
        <el-input
          disabled
          v-model="cid"
        />
      </el-col>
  </el-row>
  <el-row :gutter="0">
      <el-col :span="4">
        <span>类型：</span>
      </el-col>
      <el-col :span="2">
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
        />
      </el-col>
  </el-row>
  <el-row :gutter="20">
      <el-col :span="4">
        <span>监控频率：</span>
      </el-col>
      <el-col :span="1">
        <el-input-number
          v-model="monitorRate"
          placeholder="Please input"
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
          placeholder="Please input"
        />
      </el-col>
  </el-row>
  <el-row :gutter="20">
      <el-col :span="10">
    <el-button type="primary" @click="BtnUpdateComapny">修改</el-button>
      </el-col>
      <el-col :span="6">
        <el-button type="warning" @click="BtnCancel">取消</el-button>
      </el-col>
  </el-row>
  </div>
</template>

<script lang="ts">import axios from "axios";
import { ElMessage } from "element-plus";
import { computed, defineComponent, getCurrentInstance, onMounted,ref } from "vue-demi";
import { useRouter } from 'vue-router';

export default defineComponent({

    setup(){
        // const { proxy } = getCurrentInstance()
        // console.log(proxy.$root.$route.params.id)
        // let cmpId = ref(proxy.$route.params.id)
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
        const router = useRouter()
        let cmpId = ref(router.currentRoute.value.params.id)
        
        let cid = ref(0)
        let ctype = ref('')
        let companyName = ref('')
        let rootDomain = ref('')
        let srcUrl = ref('')
        let keyWord = ref('')
        let monitorStauts = ref(true)
        let monitorRate = ref(24)
        let vulnScanStatus = ref(true)
        let vulnScanRate = ref(24)

        onMounted (() =>{
            const data ={"id":cmpId.value}
            axios({
                url: '/api/v1/getCompanyByid',
                method: 'POST',
                data: data,
            }).then((res) => {
                if (res.data.code == 400){
                    ElMessage({
                        message: res.data.msg,
                        type: 'error',
                    })
                }else if (res.data.code == 200){
                     let jd = JSON.parse(res.data.msg)
                     cid.value = jd.id
                     ctype.value = jd.projectType
                     companyName.value = jd.companyName
                     rootDomain.value = jd.domain
                     srcUrl.value = jd.srcUrl
                     keyWord.value = jd.keyWord
                     monitorStauts.value = jd.monitorStatus
                     monitorRate.value = jd.monitorRate
                     vulnScanStatus.value = jd.vulnScanStatus
                     vulnScanRate.value = jd.vulnScanRate
                    
                }else if (res.data.code == 302){
                    router.push({ path: '/login' })
                }
            });

        })

        const BtnUpdateComapny =()=>{
          const data ={"id":cid.value, "projectType":ctype.value,"companyName":companyName.value,"domain":rootDomain.value,
                         "srcUrl":srcUrl.value,"keyWord":keyWord.value,"monitorStatus":monitorStauts.value,
                         "monitorRate":monitorRate.value,"vulnScanStatus":vulnScanStatus.value,"vulnScanRate":vulnScanRate.value}
            axios({
                url: '/api/v1/updateCompanyInfo',
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

        // router.getRoutes().
        return{
          cid,
          ctype,
          options,
          cmpId,
          companyName,
          srcUrl,
          keyWord,
          vulnScanStatus,
          vulnScanRate,
          rootDomain,
          monitorRate,
          monitorStauts,
          BtnCancel,
          BtnUpdateComapny,
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