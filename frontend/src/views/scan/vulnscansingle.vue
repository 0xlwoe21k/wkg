<template>

<el-row :gutter="20">
    <el-col :span="16">
    <div class="grid-content bg-purple">
        <el-input v-model="inpTarget"   placeholder="Please input">
            <template #prepend>根域名</template>
        </el-input>
    </div></el-col>
    <el-col :span="2"><div class="grid-content bg-purple">
        <el-button type="primary" @click="onScanBtn" icon="el-icon-search">开始扫描</el-button>
    </div></el-col>
  </el-row>

    <el-row :gutter="20">
    <el-col :span="20">
        <div class="grid-content bg-purple">
            <div style="margin-top:10px;text-align:left;white-space: pre-wrap; background-color: #f0f3f7;">        
                <el-input
                    
                    :rows="30"
                    v-loading="loading"
                    element-loading-text="Loading..."
                    element-loading-spinner="el-icon-loading"
                    v-model="scanResult"
                    type="textarea"
                    placeholder="显示扫描结果"
                    
                />
            </div>
        </div>
    </el-col>
    </el-row>
</template>

<script lang="ts">

import { defineComponent,ref ,onMounted} from "vue";
import { ElMessage } from 'element-plus'

export default defineComponent({
    name:"vulnscanmulti",
    setup(){
        let resultShow = ref(false)
        let test:string = "漏洞扫描多个漏洞"
        let inpTarget = ref('')
        let ws:WebSocket;
        let scanResult = ref('')
        let loading = ref(false)

        const onScanBtn = () => {
            loading.value = true
            scanResult.value = ''
            sendMessage(inpTarget.value)

        }

         const initwebsocket = ()=>{
            ws  = new WebSocket('ws://127.0.0.1:7788/v2/VulnScanSingle')
            ws.onclose = close;
            ws.onerror = onError;
            ws.onopen = open;
            ws.onmessage = message;
         }

        const open = ()=>{
            console.log("connect success")
        }
        const message = (value:any)=>{
            if (value != ""){
                resultShow.value= true
                loading.value = false
                scanResult.value += value.data+'\n'
            }
        }

        const close = ()=>{  //关闭
            errNotice("websocket connect error.")
        }

        const onError = ()=>{
            errNotice("websocket error.")
        }
        const sendMessage = (value:any)=>{
            ws.send(value)
        }

        const errNotice = (value:string) => {
            // ElMessage.error(value)
        }

        const InfoNotice = (value:string) => {
            ElMessage({
                message: value,
                type: 'success',
            })
        }

        onMounted(()=>{
            initwebsocket();
        })
        return{
            test,
            inpTarget,
            loading,
            resultShow,
            scanResult,
            onScanBtn,
           
        }
    }
})

</script>


<style>

#tm{
    color: rgb(214, 32, 32);
}

</style>