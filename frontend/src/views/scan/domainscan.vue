<template>

<el-row :gutter="20">
    <el-col :span="16">
    <div class="grid-content bg-purple">
        <el-input v-model="input_rootDomain"  placeholder="Please input">
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
            <div>
                <el-checkbox v-model="passive" label="passive"></el-checkbox>
                <el-checkbox v-model="active" label="active"></el-checkbox>
            </div>
        </div>
    </el-col>
  </el-row>

    <el-row :gutter="20">
    <el-col :span="20">
        <div class="grid-content bg-purple">
            <div style="margin-top:10px;text-align:left;white-space: pre-wrap; background-color: #f0f3f7;">        
                <el-input
                    v-show="resultShow"
                    :maxlength="100"
                    v-loading="loading"
                    element-loading-text="Loading..."
                    element-loading-spinner="el-icon-loading"
                    v-model="scanResult"
                    autosize="true"
                    type="textarea"
                    
                />
            </div>
        </div>
    </el-col>
      </el-row>

</template>

<script lang="ts">

import { defineComponent ,reactive,toRaw,ref,onMounted} from "vue";
import { ElMessage } from 'element-plus'
import axios from 'axios';


interface domainScanInter {
    rootdomain:string;
    passive:boolean;
    active:boolean;
}

export default defineComponent({
    name:"domainscan",
    setup(){
        let resultShow = ref(false)
        let loading = ref(false)
        const passive = ref(true)
        const active = ref(false)
        let domainScanInfo:domainScanInter = reactive({
            rootdomain: '',
            passive: false,
            active: false
        });
        let scanResult = ref('')
        let msg:string=ref('')
        let ws:WebSocket;
        let errMsg = ref('')
        let test:string = "扫描子域名"
        let input_rootDomain:string = ref('')
        // const onScaninput = (value: string | number) => {
        //     startScan()
        // }
        const onScanBtn = () => {
            startScan()
        }

        const startScan = () => {
            
            loading.value=true
            InfoNotice('scanning...')
            scanResult.value = ''

            if (passive.value){
                domainScanInfo.passive = true
            }else{
                domainScanInfo.passive = false
            }
            if (active.value){
                domainScanInfo.active = true
            }else{
                domainScanInfo.active = false
            }

            if (input_rootDomain != ""){
                domainScanInfo.rootdomain = input_rootDomain.value
            }
            sendMessage(JSON.stringify(domainScanInfo))
        }

         const initwebsocket = ()=>{
            ws  = new WebSocket('ws://127.0.0.1:7788/v2/domainScan')
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
            ElMessage.error(value)
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
            passive,
            active,
            msg,
            initwebsocket,
            onScanBtn,
            test,
            errNotice,
            resultShow,
            input_rootDomain,
            scanResult,
            loading,
        }
    }
})

</script>