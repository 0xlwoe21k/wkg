<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
    setup(){
    const asideMenu =[
    {
        path:'/',
        name: 'home',
        label: '首页',
         icon: 's-home'
    },]
    
    return{
        asideMenu,
        }
    }
});
</script>