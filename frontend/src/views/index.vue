<template>
  <el-container style="height: 100%; border: 1px solid #eee" direction="vertical">
      <el-header style="text-align: right; font-size: 12px">
        <el-dropdown>
          <i class="el-icon-setting" style="margin-right: 15px"></i>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>View</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
        <span>{{username}}</span>
      </el-header>
      
    <el-container style="height: 100%; border: 1px solid rgb(0, 0, 0)" direction="horizontal">
    
    <el-aside width="200px" >
              
      <el-menu 
        :default-openeds="['/index/dashboard']" 
        :router="true"
        active-text-color="#ffd04b"
        background-color="#545c64"
        text-color="#fff"
        class="el-menu-vertical-demo"
        :unique-opened="true"
        >
        <el-menu-item index="/index/dashboard">
          <template #title>
            <i class="el-icon-setting"></i>dashboard
            </template>
        </el-menu-item>
        <el-sub-menu index="2">
          <template #title>
            <i class="el-icon-menu"></i>扫描模块</template>
            <el-menu-item index="/index/domainscan">域名暴破</el-menu-item>
            <el-menu-item index="/index/vulnscansingle">漏扫（单）</el-menu-item>
            <el-menu-item index="/index/vulnscanmulti">漏扫（多）</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="3">
          <template #title><i class="el-icon-search"></i>资产管理</template>
              <el-menu-item index="/index/company">公司管理</el-menu-item>
              <el-menu-item index="/index/domain">域名管理</el-menu-item>
              <el-menu-item index="/index/ip">IP管理</el-menu-item>
              <el-menu-item index="/index/website">网站管理</el-menu-item>
              <el-menu-item index="/index/service">服务管理</el-menu-item>
              <el-menu-item index="/index/littleprogram">小程序管理</el-menu-item>
              <el-menu-item index="/index/wechataccout">微信公众号管理</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="4">
          <template #title>
            <i class="el-icon-setting"></i>系统设置
          </template>
            <el-menu-item index="/index/scanConfig">扫描设置</el-menu-item>
            <el-menu-item index="/index/noticeConfig">通知配置</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    <el-container>
      <el-main>
       <router-view />
      </el-main>
      <el-footer>design by gelen v1.0</el-footer>
    </el-container>
      </el-container>
  </el-container>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup() {
    let username:string = "gelen"

    return {
      username,
    }
  },
})
</script>

<style>
.el-header {
  background-color: #8caac7;
  color: var(--el-text-color-primary);
  line-height: 60px;
}

.el-aside {
  color: var(--el-text-color-primary);
  background-color: rgb(145, 123, 123);
}

.el-menu{
  height: 100%;
} 

.el-main{
  margin-top: 10px;
  margin-left: 10px;
  border: 2px solid var(--el-border-color-base);
  height: 800px;
}
 </style> 
