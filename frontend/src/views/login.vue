<template>
  <div class="login-container">
    <!-- <h2 class="login-title">WKG</h2> -->
    <a-form
      class="login-form"
      ref="formRef"
      :model="formState"
      style="margin-left: 35%; padding: 20px 90px"
    >
      <!-- <h2 class="title">WKG</h2> -->
      <a-form-item>
        <a-input class="inputBox" v-model:value="formState.username"></a-input>
      </a-form-item>
      <a-form-item>
        <a-input-password
          class="inputBox"
          v-model:value="formState.password"
        ></a-input-password>
      </a-form-item>

      <a-form-item>
        <a-button class="submit" type="primary" @click="onSubmit"
          >登录</a-button
        >
      </a-form-item>
    </a-form>
  </div>
</template>

<script lang="ts">
import { defineComponent, toRaw, ref, reactive } from "vue";
import type { UnwrapRef } from "vue";
import axios from "axios";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import auth =  require("../service/auth.service.js")

interface FormState {
  username: string;
  password: string;
}

export default defineComponent({
  components: {},
  setup() {
    const formRef = ref();
    const router = useRouter();
    const store = useStore();
    const formState: UnwrapRef<FormState> = reactive({
      username: "gelen",
      password: "gelen",
    });

    const onSubmit = () => {
      axios({
        url: "/api/login",
        method: "POST",
        data: toRaw(formState),
      }).then((res) => {
        const token = res.data.data.token;
        localStorage.setItem("token", token);
        router.push("/index/dashboard");
      });
    };

    return {
      formRef,
      formState,
      onSubmit,
    };
  },
});
</script>

<style>
.login-form {
  width: 560px;
  height: 372px;
  margin: 550px auto;
  /* background: url("../assets/bg1.png"); */
}

/* 背景 */
.login-container {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: rgb(122, 120, 116);
  background: url("../assets/bg.png");
}

/* Log */
.login-title {
  color: #fff;
  text-align: center;
  margin: 110px 0;
  font-size: 48px;
  font-family: Microsoft Yahei;
}
/* 登陆按钮 */
.submit {
  width: 100%;
  height: 45px;
  font-size: 16px;
}
/* 用户登陆标题 */
.title {
  margin-bottom: 50px;
  color: #fff;
  font-weight: 700;
  font-size: 24px;
  font-family: Microsoft Yahei;
}
/* 输入框 */
.inputBox {
  height: 45px;
}
/* 输入框内左边距50px */
.ant-input-affix-wrapper .ant-input:not(:first-child) {
  padding-left: 50px;
}
</style>
