<template>
  <a-layout>
    <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible>
      <div class="logo">
        <img class="logoimg" src="../assets/logo.png" />
      </div>
      <a-menu v-model:selectedKeys="selectedKeys" @click="menuClcik" theme="dark" mode="inline">
        <template v-for="(item ,index) in menuList" :key="index">
          <template v-if="item.children">
            <a-sub-menu :key="item.key">
              <template #title>{{ item.title }}</template>
              <a-menu-item v-for="subitem in item.children" :key="subitem.key">
                <router-link :to="subitem.path">
                  <span>{{ subitem.title }}</span>
                </router-link>
              </a-menu-item>
            </a-sub-menu>
          </template>
          <template v-else>
            <a-menu-item :key="item.key">
              <router-link :to="item.path">
                <template #icon>{{ item.icon }}</template>
                <span>{{ item.title }}</span>
              </router-link>
            </a-menu-item>
          </template>
        </template>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <menu-unfold-outlined
          v-if="collapsed"
          class="trigger"
          @click="() => (collapsed = !collapsed)"
        />
        <menu-fold-outlined v-else class="trigger" @click="() => (collapsed = !collapsed)" />
        <a-dropdown placement="bottomCenter">
          <a-avatar class="user">
            <template #icon>
              <user-outlined @click="BntUser"></user-outlined>
            </template>
          </a-avatar>

          <template #overlay>
            <a-menu>
              <a-menu-item>
                <a target="_blank" rel="noopener noreferrer">个人设置</a>
              </a-menu-item>
              <a-menu-item style="text-align: center;">
                <a target="_blank" @click="OnLogout" rel="noopener noreferrer">登出</a>
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </a-layout-header>

      <a-layout-content>
        <a-breadcrumb style="margin-left:1% ; padding: 5px; font-size: 14px;">
          <a-breadcrumb-item v-for="(item, index) in routes" :key="item.name">
            <router-link
              v-if=" index !== routes.length -1 "
              :to="{ path: item.path === '' ? '/' : item.path }"
            >{{ item.meta.title }}</router-link>
            <span v-else>{{ item.meta.title }}</span>
          </a-breadcrumb-item>

        </a-breadcrumb>
        <div
          :style="{
            margin: '0px 16px',
            padding: '24px',
            background: '#fff',
            minHeight: '700px',
          }"
        >
          <router-view />
        </div>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script lang="ts">
import {
  UserOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  // DashboardOutlined,
} from "@ant-design/icons-vue";
import { defineComponent, ref, computed } from "vue";
import { useRouter } from "vue-router";
import { menuList } from "../common/menu"



export default defineComponent({
  components: {
    UserOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    // DashboardOutlined,
  },
  setup() {
    const router = useRouter();
    const routes = computed(() => { return router.currentRoute.value.matched.filter(item => item.meta.title) })

    const menuClcik = (item: any, key: any, keyPath: any) => {
      if (item.key == 1) {
        // console.log(router.getRoutes())
      }
    };
    const BntUser = () => {
      console.log(123)
    }

    const OnLogout = () => {
      localStorage.removeItem('token')
      router.push("/login");
    }


    return {
      selectedKeys: ref<string[]>(["1"]),
      collapsed: ref<boolean>(false),
      routes,
      menuList,
      basePath: '',
      OnLogout,
      menuClcik,
      BntUser,
    };
  },
});



</script>

<style>
.logoimg {
  height: 44px;
  vertical-align: top;
  margin-right: 16px;
  border-style: none;
}

.user {
  font-size: 18px;
  background-color: #1890ff;
  /* line-height: 64px; */
  /* padding: 0 24px; */
  /* cursor: pointer; */
  transition: color 0.3s;
  margin-left: 92%;
}

.trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 0px;
  cursor: pointer;
  transition: color 0.3s;
  /* margin-top: 1px; */
  /* margin-left: 90%; */
}

.trigger:hover {
  color: #1890ff;
}

#components-layout-demo-custom-trigger .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.3);
  margin: 16px;
}

.site-layout .site-layout-background {
  background: #fff;
}
</style>
