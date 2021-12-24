<template>
    <a-input-search
        v-model:value="searchValue"
        style="margin-bottom: 8px"
        enter-button
        placeholder="Search"
    />

    <a-directory-tree
        v-model:expandedKeys="expandedKeys"
        v-model:selectedKeys="selectedKeys"
        :auto-expand-parent="autoExpandParent"
        :showLine="true"
        :multiple="false"
        :height="500"
        :tree-data="tdList"
        :style="{ padding: '10px' }"
        @expand="Onexpand"
    >
        <template #title="{ title }">
            <span v-if="title.indexOf(searchValue) > -1">
                {{ title.substr(0, title.indexOf(searchValue)) }}
                <span
                    style="color: #f50"
                >{{ searchValue }}</span>
                {{ title.substr(title.indexOf(searchValue) + searchValue.length) }}
            </span>
            <span v-else>{{ title }}</span>
        </template>
    </a-directory-tree>
</template>


<script lang="ts">
interface TreeProps {
    title?: string
    key?: string
    level?: number
    isLeaf?: boolean
    children?: TreeProps[]
}

interface Konwledge {
    title?: string
    content?: string
    updateTime?: string
}

interface option {
    value: string,
    label: string
}
import { defineComponent, ref, onMounted, reactive, watch } from 'vue';
import knowledgeService from "../../service/knowledge.service";

export default defineComponent({
    components: {

    },
    setup() {
        let searchValue = ref('')
        const expandedKeys = ref<(string | number)[]>([]);
        const selectedKeys = ref<string[]>([]);
        let tdList: TreeProps[] = reactive([])
        let tlList: Konwledge[] = reactive([])
        let autoExpandParent = ref<boolean>(true);
        let ShowNodeKey: (string | number)[] = []


        const Init = () => {
            knowledgeService.getTree().then((res: any) => {
                // tdList = JSON.parse(res.data.data)
                let td: TreeProps = {}
                let jdata = JSON.parse(res.data.data)
                if (jdata.length < 0) { return }

                for (let element of jdata) {
                    td.title = element.title
                    td.key = element.key
                    td.level = element.level
                    td.isLeaf = element.isLeaf
                    td.children = element.children
                    tdList.push(td)
                    td = {}
                }
            })
            knowledgeService.getSummary().then((res: any) => {
                let _kList: Konwledge = {}
                let jdata = JSON.parse(res.data.msg)
                if (jdata.length < 0) { return }
                for (let i = 0; i < jdata.length; i++) {
                    _kList.title = jdata[i].title
                    _kList.updateTime = jdata[i].updateTime
                    // console.log(_kList)
                    tlList.push(_kList)
                    _kList = {}
                }

            })
            console.log(tdList)
        }


        onMounted(() => {
            Init()
        })


        const getParentKey = (
            key: string | number,
            tree: TreeProps[],
        ): string | number | undefined => {
            let parentKey;
            let xtree = tree || []

            for (let i = 0; i < xtree.length; i++) {
                const node = xtree[i];
                if (node.children) {
                    if (node.children.some(item => item.key === key)) {
                        parentKey = node.key;
                    } else if (getParentKey(key, node.children)) {
                        parentKey = getParentKey(key, node.children);
                    }
                }
            }
            return parentKey;
        };

        const getMatchKey = (node: TreeProps[]) => {
            if (!node) {
                return
            }
            for (let i = 0; i < node.length; i++) {
                if ((node[i].title || '').indexOf(searchValue.value) > -1) {
                    let xkey = node[i].key || ''
                    // ShowNodeKey.push(xkey)
                    let pkey = getParentKey(xkey, tdList) || ''
                    ShowNodeKey.push(pkey)
                }
                let _child = node[i].children || []
                if (_child) {
                    getMatchKey(_child)
                }
            }
            console.log(ShowNodeKey)
        }

        watch(searchValue, value => {
            ShowNodeKey = []
            if (value.length > 0) {

                getMatchKey(tdList);
                searchValue.value = value;
                expandedKeys.value = ShowNodeKey
                autoExpandParent.value = true;
            } else {
                searchValue.value = value;
                expandedKeys.value = []
                autoExpandParent.value = true;
            }


        });

        const Onexpand = (exkey: any, expand: any) => {
            expandedKeys.value = exkey;
            autoExpandParent.value = false;
        }



        return {
            expandedKeys,
            selectedKeys,
            searchValue,
            tdList,
            autoExpandParent,
            Onexpand,
        };
    },
});
</script>
 
