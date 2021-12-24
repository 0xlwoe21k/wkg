<template>
  <div>
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import jwt_decode from "jwt-decode";
import { useStore } from "vuex";
import { onMounted } from "vue";

export default defineComponent({
  setup() {
    const store = useStore();

    onMounted(() => {
      if (localStorage.token) {
        const decode = jwt_decode(localStorage.token);
        // store.dispatch('setUser', decode);
      }
    });
  },
});
</script>
<style lang="stylus">
#app
  font-family Avenir, Helvetica, Arial, sans-serif
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale

  color #2c3e50
</style>
