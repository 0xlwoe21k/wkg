<template>
  <div style="margin-left: 10px">
    <label>编辑公司或项目</label>
  </div>
  <div style="margin: 50px 0 0 50px">
    <a-row type="flex" align="middle">
      <a-col :span="2">
        <label>ID:</label>
      </a-col>
      <a-col :span="1">
        <a-input v-model:value="cid" />
      </a-col>
      <a-col :span="21"></a-col>
    </a-row>
    <a-row style="margin-top: 15px" type="flex" align="middle">
      <a-col :span="2">
        <label>类型:</label>
      </a-col>
      <a-col :span="1">
        <a-select
          ref="select"
          v-model:value="ctype"
          :options="options"
          style="width: 120px"
          @focus="focus"
        ></a-select>
      </a-col>
      <a-col :span="20"></a-col>
    </a-row>
    <a-row style="margin-top: 15px" type="flex" align="middle">
      <a-col :span="2">
        <label>公司名称:</label>
      </a-col>
      <a-col :span="4">
        <a-input v-model:value="companyName" placeholder="公司名或项目名" />
      </a-col>
      <a-col :span="18"></a-col>
    </a-row>
    <a-row style="margin-top: 15px" type="flex" align="middle">
      <a-col :span="2">
        <label>根域名:</label>
      </a-col>
      <a-col :span="6">
        <a-textarea
          v-model:value="rootDomain"
          placeholder="单个或多个根域名，以'|'隔开."
          auto-size
        />
      </a-col>
      <a-col :span="16"></a-col>
    </a-row>
    <a-row style="margin-top: 15px" type="flex" align="middle">
      <a-col :span="2">
        <label>项目URL:</label>
      </a-col>
      <a-col :span="6">
        <a-input
          v-model:value="srcUrl"
          placeholder="应急响应中心或项目关联地址."
        />
      </a-col>
      <a-col :span="16"></a-col>
    </a-row>
    <a-row style="margin-top: 15px" type="flex" align="middle">
      <a-col :span="2">
        <label>项目关键字:</label>
      </a-col>
      <a-col :span="6">
        <a-textarea
          v-model:value="keyWord"
          placeholder="项目关键字，用于漏洞或资产查找和导入."
          auto-size
        />
      </a-col>
      <a-col :span="16"></a-col>
    </a-row>
    <a-row style="margin-top: 15px" type="flex" align="middle">
      <a-col :span="2">
        <label>监控状态:</label>
      </a-col>
      <a-col :span="1">
        <a-switch v-model:checked="monitorStauts" />
      </a-col>
      <a-col :span="20"></a-col>
    </a-row>
    <a-row style="margin-top: 15px" type="flex" align="middle">
      <a-col :span="2">
        <label>监控频率(小时):</label>
      </a-col>
      <a-col :span="1">
        <a-input-number
          id="inputNumber"
          v-model:value="monitorRate"
          :min="1"
          :max="200"
        />
      </a-col>
      <a-col :span="16"></a-col>
    </a-row>
    <a-row style="margin-top: 50px" type="flex" align="middle">
      <a-col :span="3">
        <a-button @click="OnUpdate" type="primary">更新</a-button>
      </a-col>
      <a-col :span="4">
        <a-button @click="OnBack">返回</a-button>
      </a-col>
      <a-col :span="16"></a-col>
    </a-row>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { message } from "ant-design-vue";
import srcService from "../../../service/src.service";

const options = ref([
  {
    value: "SRC",
    label: "SRC",
  },
  {
    value: "众测",
    label: "众测",
  },
  {
    value: "公益",
    label: "公益",
  },
  {
    value: "工作",
    label: "工作",
  },
  {
    value: "挖着玩",
    label: "挖着玩",
  },
]);

export default defineComponent({
  setup() {
    const router = useRouter();
    let cid = ref(router.currentRoute.value.params.id);
    let ctype = ref("SRC");
    let companyName = ref("");
    let rootDomain = ref("");
    let srcUrl = ref("");
    let keyWord = ref("");
    let monitorStauts = ref(true);
    let monitorRate = ref(24);

    onMounted(() => {
      console.log(cid);
      srcService.getCompanyInfoById(cid).then((res: any) => {
        if (res.data.code == 400) {
          message.warning(res.data.msg);
        } else if (res.data.code == 200) {
          let jd = JSON.parse(res.data.msg);
          cid.value = jd.id;
          ctype.value = jd.projectType;
          companyName.value = jd.companyName;
          rootDomain.value = jd.domain;
          srcUrl.value = jd.srcUrl;
          keyWord.value = jd.keyWord;
          monitorStauts.value = jd.monitorStatus;
          monitorRate.value = jd.monitorRate;
        } else if (res.data.code == 302) {
          router.push({ path: "/login" });
        }
      });
    });

    const OnUpdate = () => {
      srcService
        .UpdateComapny(
          cid,
          ctype,
          companyName,
          rootDomain,
          srcUrl,
          keyWord,
          monitorStauts,
          monitorRate
        )
        .then((res: any) => {
          if (res.data.code == 400) {
            message.error(res.data.msg);
          } else if (res.data.code == 200) {
            message.success(res.data.msg);
            router.push({ path: "/index/src/company" });
          }
        });
    };

    const OnBack = () => {
      router.push({ path: "/index/src/company" });
    };

    return {
      options,
      cid,
      ctype,
      companyName,
      rootDomain,
      srcUrl,
      keyWord,
      monitorStauts,
      monitorRate,
      OnUpdate,
      OnBack,
    };
  },
});
</script>
