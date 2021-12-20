<template>
    <div class="global">
        <a-row type="flex" justify="start">
            <a-col :span="3">
                <a-row class="dTree">
                    <a-col>
                        <a-input-search
                            v-model:value="searchValue"
                            style="margin-bottom: 8px"
                            enter-button
                            placeholder="Search"
                        />
                    </a-col>
                    <a-col>
                        <a-directory-tree
                            v-model:expandedKeys="expandedKeys"
                            v-model:selectedKeys="selectedKeys"
                            :showLine="true"
                            :tree-data="treeData"
                            :style="{ padding: '10px' }"
                        ></a-directory-tree>
                    </a-col>
                </a-row>
            </a-col>
            <a-col :span="1"></a-col>
            <a-col :span="18">
                <a-row>
                    <a-col :span="5">
                        <a-input-search
                            value="搜索"
                            placeholder="标题."
                            enter-button="搜索"
                            @search="onSearch"
                        />
                    </a-col>
                </a-row>
                <a-row style="margin-top: 10px;">
                    <a-col :span="24">
                        <a-table :columns="columns" :data-source="data">
                            <template #headerCell="{ column }">
                                <template v-if="column.key === 'name'">
                                    <span>
                                        <smile-outlined />Name
                                    </span>
                                </template>
                            </template>

                            <template #bodyCell="{ column, record }">
                                <template v-if="column.key === 'name'">
                                    <a>{{ record.name }}</a>
                                </template>
                                <template v-else-if="column.key === 'tags'">
                                    <span>
                                        <a-tag
                                            v-for="tag in record.tags"
                                            :key="tag"
                                            :color="tag === 'loser' ? 'volcano' : tag.length > 5 ? 'geekblue' : 'green'"
                                        >{{ tag.toUpperCase() }}</a-tag>
                                    </span>
                                </template>
                                <template v-else-if="column.key === 'action'">
                                    <span>
                                        <a>Invite 一 {{ record.name }}</a>
                                        <a-divider type="vertical" />
                                        <a>Delete</a>
                                        <a-divider type="vertical" />
                                        <a class="ant-dropdown-link">
                                            More actions
                                            <down-outlined />
                                        </a>
                                    </span>
                                </template>
                            </template>
                        </a-table>
                    </a-col>
                </a-row>
            </a-col>
        </a-row>
    </div>
</template>


<script lang="ts">

interface TreeChildren {
    title: string
    key: string
    isLeaf?: boolean
    children?: Array<TreeChildren>

}

interface TreeProps {
    title: string
    key: string
    isLeaf?: boolean
    children?: Array<TreeChildren>
}

const columns = [
    {
        name: 'Name',
        dataIndex: 'name',
        key: 'name',
    },
    {
        title: 'Age',
        dataIndex: 'age',
        key: 'age',
    },
    {
        title: 'Address',
        dataIndex: 'address',
        key: 'address',
    },
    {
        title: 'Tags',
        key: 'tags',
        dataIndex: 'tags',
    },
    {
        title: 'Action',
        key: 'action',
    },
];

const data = [
    {
        key: '1',
        name: 'John Brown',
        age: 32,
        address: 'New York No. 1 Lake Park',
        tags: ['nice', 'developer'],
    },
    {
        key: '2',
        name: 'Jim Green',
        age: 42,
        address: 'London No. 1 Lake Park',
        tags: ['loser'],
    },
    {
        key: '3',
        name: 'Joe Black',
        age: 32,
        address: 'Sidney No. 1 Lake Park',
        tags: ['cool', 'teacher'],
    },
];

import { defineComponent, ref } from 'vue';
import { SmileOutlined, DownOutlined } from '@ant-design/icons-vue';

export default defineComponent({
    components: {
        SmileOutlined,
        DownOutlined,
    },
    setup() {
        let searchValue = ref('')
        const expandedKeys = ref<string[]>(['0-0', '0-1']);
        const selectedKeys = ref<string[]>([]);
        const treeData: TreeProps[] = [
            {
                title: 'parent 0',
                key: '0-0',
                children: [
                    {
                        title: 'leaf 0-0',
                        key: '0-0-0',
                        children: [
                            {
                                title: 'leaf 0-0',
                                key: '0-0-0-0',
                                isLeaf: true,
                            },
                            {
                                title: 'leaf 0-1',
                                key: '0-0--0-1',
                                isLeaf: true,
                            },
                        ],
                    },
                    {
                        title: 'leaf 0-1',
                        key: '0-0-1',
                        isLeaf: true,
                    },
                ],
            },
            {
                title: 'parent 1',
                key: '0-1',
                children: [
                    {
                        title: 'leaf 1-0',
                        key: '0-1-0',
                        isLeaf: true,
                    },
                    {
                        title: 'leaf 1-1',
                        key: '0-1-1',
                        isLeaf: true,
                    },
                ],
            },
        ];
        const onSearch = () => { console.log() }
        return {
            expandedKeys,
            selectedKeys,
            treeData,
            data,
            onSearch,
            searchValue,
            columns,
        };
    },
});
</script>


<style>
/* .global {
    background: "rgb(73, 64, 52)";
    height: 100%;

} */

.dTree {
    /* width: 100%; */
    border: 2px solid rgb(240, 242, 245);
}
</style>