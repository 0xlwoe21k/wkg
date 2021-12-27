<template>
    <div class="global">
        <a-row type="flex" justify="start">
            <!-- <a-col :span="3"> -->
            <a-col :span="6">
                <a-row>
                    <a-col :span="18">
                        <a-input-search
                            v-model:value="searchValue"
                            style="margin-bottom: 8px"
                            placeholder="Search"
                        >
                            <template #enterButton>
                                <redo-outlined @click="OnFlush" />
                            </template>
                        </a-input-search>
                    </a-col>
                </a-row>
                <a-row type="flex" justify="start">
                    <a-col :span="11">
                        <a-button @click="BtnAddTop" type="primary">添加一级节点</a-button>
                        <a-modal
                            v-model:visible="modelOption.showTopModel"
                            title="添加一级节点"
                            ok-text="确认"
                            cancel-text="取消"
                            @ok="AddTopNode"
                        >
                            <a-row>
                                <a-col style="margin-top:5px;">
                                    <label>名称：</label>
                                </a-col>
                                <a-col>
                                    <a-input v-model:value="modelOption.topInpuVaule" placeholder />
                                </a-col>
                            </a-row>
                        </a-modal>
                    </a-col>
                    <a-col :span="8">
                        <a-button @click="BtnAddSecond" type="primary">添加二级节点</a-button>
                        <a-modal
                            v-model:visible="modelOption.showSecondModel"
                            title="添加一级节点"
                            ok-text="确认"
                            cancel-text="取消"
                            @ok="AddSecondNode"
                        >
                            <a-row>
                                <a-col>
                                    <a-select
                                        ref="select"
                                        v-model:value="modelOption.topSelectVaule"
                                        :options="modelOption.topSelectOption"
                                        style="width: 120px"
                                    ></a-select>
                                </a-col>
                                <a-col style="margin-top:5px;">
                                    <label>名称：</label>
                                </a-col>
                                <a-col>
                                    <a-input v-model:value="modelOption.SecondInpuVaule" />
                                </a-col>
                            </a-row>
                        </a-modal>
                    </a-col>
                </a-row>
                <a-row style="margin-top: 5px">
                    <a-col class="dTree">
                        <a-directory-tree
                            v-model:expandedKeys="expandedKeys"
                            v-model:selectedKeys="selectedKeys"
                            :auto-expand-parent="autoExpandParent"
                            :showLine="true"
                            :multiple="false"
                            :tree-data="tdList.list"
                            :style="{ padding: '10px' }"
                            :height="1000"
                            @expand="Onexpand"
                            @select="onSelect"
                        >
                            <template #title="{ isLeaf, key: treeKey, title }">
                                <a-dropdown :trigger="['contextmenu']">
                                    <span v-if="title.indexOf(searchValue) > -1">
                                        {{ title.substr(0, title.indexOf(searchValue)) }}
                                        <span
                                            style="color: #f50"
                                        >{{ searchValue }}</span>
                                        {{
                                            title.substr(
                                                title.indexOf(searchValue) + searchValue.length
                                            )
                                        }}
                                    </span>
                                    <span v-else>{{ title }}</span>
                                    <template #overlay>
                                        <a-menu>
                                            <a-menu-item
                                                @click="() => OnDeleteClick(isLeaf, treeKey)"
                                                key="1"
                                            >delete node</a-menu-item>
                                        </a-menu>
                                    </template>
                                </a-dropdown>
                            </template>
                        </a-directory-tree>
                    </a-col>
                </a-row>
            </a-col>
            <!-- <a-col :span="1"></a-col> -->
            <a-col :span="17">
                <a-row>
                    <a-col v-if="showInit" :span="24">
                        <a-row>
                            <a-button @click="OnAdd" type="primary">
                                <template #icon>
                                    <FileAddOutlined />
                                </template>
                                新增知识页
                            </a-button>
                        </a-row>
                        <a-row v-if="showInit">
                            <a-col :span="24" style="margin-top: 10px">
                                <a-timeline mode="alternate" style="margin: 50px 300px 0 0">
                                    <div v-for="(item, index) in tlList" :key="index">
                                        <a-timeline-item>{{ item.title }} | {{ item.updateTime }}</a-timeline-item>
                                    </div>
                                    <a-timeline-item>
                                        <template #dot>
                                            <ClockCircleOutlined style="font-size: 16px" />
                                        </template>
                                        已添加的漏洞
                                    </a-timeline-item>
                                </a-timeline>
                            </a-col>
                        </a-row>
                    </a-col>
                </a-row>
                <!-- 第一行做为工具栏显示 -->
                <a-row :gutter="16">
                    <a-col v-if="showEdit" :span="24">
                        <a-row :gutter="16">
                            <a-col>
                                <a-select
                                    ref="select"
                                    v-model:value="topCate"
                                    :options="topOptions.list"
                                    style="width: 120px"
                                    @change="topSelectHandleChange"
                                ></a-select>
                            </a-col>
                            <a-col>
                                <a-select
                                    ref="select"
                                    v-model:value="secondCate"
                                    :options="secondOptions.list"
                                    style="width: 120px"
                                ></a-select>
                                <!-- @change="secondSelectHandleChange" -->
                            </a-col>
                            <a-col>
                                <a-button @click="OnSave" type="primary">保存</a-button>
                            </a-col>
                            <a-col>
                                <a-button @click="OnCancel" type="primary">取消</a-button>
                            </a-col>
                        </a-row>
                        <a-row :gutter="16" style="margin-top: 10px; margin-bottom: 10px">
                            <a-col style="padding: 5px">标题：</a-col>
                            <a-col>
                                <a-input
                                    style="width: 500px"
                                    v-model:value="MdEdit.title"
                                    placeholder="输入标题."
                                />
                            </a-col>
                        </a-row>
                        <a-row>
                            <v-md-editor v-model="MdEdit.content" height="700px" @save="onSave"></v-md-editor>
                        </a-row>
                    </a-col>
                    <a-col v-if="showPriview" :span="24">
                        <a-row :gutter="16">
                            <a-col>
                                <a-button @click="OnEdit" type="primary">编辑</a-button>
                            </a-col>
                            <a-col>
                                <a-button @click="OnEditReturn" type="primary">返回</a-button>
                            </a-col>
                        </a-row>
                        <a-row :gutter="16" style="margin-top: 10px">
                            <a-col style="padding: 5px">标题：</a-col>
                            <a-col>
                                <a-input
                                    v-model:disabled="Editing"
                                    style="width: 500px"
                                    v-model:value="kwge.title"
                                    placeholder="输入标题."
                                />
                            </a-col>
                        </a-row>
                        <a-row>
                            <a-col :span="24">
                                <v-md-preview
                                    style="
                    width: 100%;
                    margin-top: 10px;
                    height: 700px;
                    border: 2px solid rgb(240, 242, 245);
                    box-shadow: 2px 2px 5px #000;
                  "
                                    :text="kwge.content"
                                ></v-md-preview>
                            </a-col>
                        </a-row>
                    </a-col>
                </a-row>
            </a-col>
        </a-row>
    </div>
</template>


<script lang="ts">
interface TreeProps {
    title?: string;
    key?: string;
    level?: number;
    isLeaf?: boolean;
    children?: TreeProps[];
}

interface Konwledge {
    title?: string;
    content?: string;
    updateTime?: string;
}

interface option {
    value: string;
    label: string;
}

interface NodeOption {
    topSelectOption: option[]
    topSelectVaule: string
    showTopModel: boolean
    topInpuVaule: string

    showSecondModel: boolean
    SecondInpuVaule: string


}


import { defineComponent, ref, onMounted, reactive, watch } from "vue";
import knowledgeService from "../../service/knowledge.service";
import { FileAddOutlined, ClockCircleOutlined, RedoOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";

export default defineComponent({
    components: {
        FileAddOutlined,
        RedoOutlined,
        ClockCircleOutlined,
    },
    setup() {
        let searchValue = ref("");
        const expandedKeys = ref<(string | number)[]>([]);
        const selectedKeys = ref<string[]>([]);
        // let tdList: TreeProps[] = reactive([])
        let tdList: { list: TreeProps[] } = reactive({ list: [] });

        let kwge: Konwledge = reactive({});
        let MdEdit: Konwledge = reactive({});
        let tlList: Konwledge[] = reactive([]);
        let topCate = ref("");
        let topOptions: { list: option[] } = reactive({ list: [] });
        let secondCate = ref("");
        let secondOptions: { list: option[] } = reactive({ list: [] });
        let autoExpandParent = ref<boolean>(true);
        let ShowNodeKey: (string | number)[] = [];
        let CurSelectKey: string;
        let showPriview = ref(false);
        let showInit = ref(true);
        let showEdit = ref(false);
        let newTopvisible = ref(false);
        let Editing = ref(true);
        let IsAdd = false;
        let modelOption: NodeOption = reactive({
            topSelectOption: [],
            topSelectVaule: '',
            showTopModel: false,
            topInpuVaule: '',
            showSecondModel: false,
            SecondInpuVaule: '',
        })


        const hideModal = (text: string) => {
            console.log(text);
        };


        const onSave = (text: string) => {
            console.log(text);
        };
        const onSearch = () => {
            console.log();
        };

        const Init = () => {
            knowledgeService.getTree().then((res: any) => {
                // tdList = JSON.parse(res.data.data)
                let jdata = JSON.parse(res.data.data);
                if (jdata.length < 0) {
                    return;
                }
                tdList.list = jdata;
                // console.log(tdList.list);
                tdList.list = tdList.list.sort((a: TreeProps, b: TreeProps) => {
                    return (b.title || "").length - (a.title || "").length;
                });
                // console.log(tdList.list);
            });
            knowledgeService.getSummary().then((res: any) => {
                let _kList: Konwledge = {};
                let jdata = JSON.parse(res.data.msg);
                if (jdata.length < 0) {
                    return;
                }
                for (let i = 0; i < jdata.length; i++) {
                    _kList.title = jdata[i].title;
                    _kList.updateTime = jdata[i].updateTime;
                    // console.log(_kList)
                    tlList.push(_kList);
                    _kList = {};
                }
            });
        };

        const topSelectHandleChange = (value: string) => {
            knowledgeService.getSecondSelectOption(value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    let jdata = JSON.parse(res.data.data);
                    if (jdata.length > 0) {
                        secondOptions.list = jdata;
                        secondCate.value = secondOptions.list[0].value;
                    } else {
                        secondOptions.list = [];
                        secondCate.value = "";
                    }
                }
            });
        };

        const UpdateTreeMenu = () => {
            tdList.list = [];

            knowledgeService.getTree().then((res: any) => {
                if (res.data.code == 200) {
                    let jdata = JSON.parse(res.data.data);
                    if (jdata.length < 0) {
                        return;
                    }
                    tdList.list = jdata;
                    // console.log(tdList.list);
                    tdList.list = tdList.list.sort((a: TreeProps, b: TreeProps) => {
                        return (b.title || "").length - (a.title || "").length;
                    });

                    // console.log(tdList.list);
                }
            });
        };

        onMounted(() => {
            Init();
        });

        const getParentKey = (
            key: string | number,
            tree: TreeProps[]
        ): string | number | undefined => {
            let parentKey;
            let xtree = tree || [];

            for (let i = 0; i < xtree.length; i++) {
                const node = xtree[i];
                if (node.children) {
                    if (node.children.some((item) => item.key === key)) {
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
                return;
            }
            for (let i = 0; i < node.length; i++) {
                if ((node[i].title || "").indexOf(searchValue.value) > -1) {
                    let xkey = node[i].key || "";
                    // ShowNodeKey.push(xkey)
                    let pkey = getParentKey(xkey, tdList.list) || "";
                    ShowNodeKey.push(pkey);
                }
                let _child = node[i].children || [];
                if (_child) {
                    getMatchKey(_child);
                }
            }
        };

        watch(searchValue, (value) => {
            if (value.length > 0) {
                getMatchKey(tdList.list);
                searchValue.value = value;
                expandedKeys.value = ShowNodeKey;
                autoExpandParent.value = true;
            } else {
                searchValue.value = value;
                expandedKeys.value = [];
                autoExpandParent.value = true;
            }
        });

        const Onexpand = (exkey: any) => {
            expandedKeys.value = exkey;
            autoExpandParent.value = false;
        };

        const OnEditReturn = () => {
            showInit.value = true;
            showEdit.value = false;
            showPriview.value = false;

            kwge.title = "";
            kwge.content = "";
            kwge.updateTime = "";
            MdEdit.title = "";
            MdEdit.content = "";
        };

        const OnEdit = () => {
            Editing.value = false;
            showEdit.value = true;
            showPriview.value = false;
            showInit.value = false;
            knowledgeService.getTopSelectOption().then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    topOptions.list = JSON.parse(res.data.msg);
                    topCate.value = topOptions.list[0].value;
                    knowledgeService
                        .getSecondSelectOption(topCate.value)
                        .then((res: any) => {
                            if (res.data.code == 400) {
                                message.error(res.data.msg);
                            } else if (res.data.code == 200) {
                                secondOptions.list = JSON.parse(res.data.data);
                                secondCate.value = secondOptions.list[0].value;
                            }
                        });
                }
            });
            MdEdit.title = kwge.title || "";
            MdEdit.content = kwge.content || "";
        };

        const OnAdd = () => {
            showEdit.value = true;
            showPriview.value = false;
            showInit.value = false;
            IsAdd = true;

            knowledgeService.getTopSelectOption().then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    topOptions.list = JSON.parse(res.data.msg);
                    topCate.value = topOptions.list[0].value;

                    knowledgeService
                        .getSecondSelectOption(topCate.value)
                        .then((res: any) => {
                            if (res.data.code == 400) {
                                message.error(res.data.msg);
                            } else if (res.data.code == 200) {
                                secondOptions.list = JSON.parse(res.data.data);
                                secondCate.value = secondOptions.list[0].value;
                            }
                        });
                }
            });
        };

        const onSelect = (exkey: any, expand: any) => {
            // console.log(expand.node.key)
            // console.log(expand)
            if (expand.node.isLeaf) {
                CurSelectKey = expand.node.key;
                showEdit.value = false;
                showPriview.value = true;
                showInit.value = false;
                // console.log(expand.node)
                knowledgeService.getKnowledge(CurSelectKey).then((res: any) => {
                    let jdata = JSON.parse(res.data.data);
                    kwge.title = jdata.title;
                    kwge.content = jdata.content;
                    kwge.updateTime = jdata.updateTime;
                });
            }
            // console.log(kwge)
        };

        const OnCancel = () => {
            console.log(CurSelectKey);
            showEdit.value = false;
            showPriview.value = false;
            showInit.value = true;

            kwge.title = "";
            kwge.content = "";
            kwge.updateTime = "";
            MdEdit.title = "";
            MdEdit.content = "";
        };

        const OnSave = () => {
            let topValue = topCate.value;
            let secondValue = secondCate.value;
            let title = MdEdit.title || "";
            let content = MdEdit.content || "";
            console.log(Editing.value);
            if (title == "" || content == "") {
                message.info("标题或内容未填写.");
                return;
            }
            if (IsAdd) {
                knowledgeService
                    .SaveNewKnowledge(topValue, secondValue, title, content)
                    .then((res: any) => {
                        if (res.data.code == 400) {
                            message.error(res.data.msg);
                        } else if (res.data.code == 200) {
                            message.success(res.data.msg);
                            let kdata = JSON.parse(res.data.data);
                            showPriview.value = true;
                            Editing.value = true
                            showEdit.value = false;
                            showInit.value = false;
                            let tsk:string[] = []
                            tsk.push(kdata[0].key)
                            expandedKeys.value = tsk
                            selectedKeys.value = tsk
                            CurSelectKey = kdata[0].key
                            console.log(CurSelectKey)
                            UpdateTreeMenu();
                            // console.log("k length:", kdata.length)
                            if (kdata.length > 0) {
                                knowledgeService.getKnowledge(kdata[0].key).then((res: any) => {
                                    if (res.data.code == 200) {
                                        let jdata = JSON.parse(res.data.data);
                                        // console.log("jdata:", jdata)
                                        kwge.title = jdata.title;
                                        kwge.content = jdata.content;
                                        kwge.updateTime = jdata.updateTime;


                                    }
                                });
                            }
                        }
                    });
            } else {
                knowledgeService
                    .SaveEditKnowledge(
                        topValue,
                        secondValue,
                        title,
                        content,
                        CurSelectKey
                    )
                    .then((res: any) => {
                        if (res.data.code == 400) {
                            message.error(res.data.msg);
                        } else if (res.data.code == 200) {
                            message.success(res.data.msg);
                            Editing.value = true;
                            showPriview.value = true;
                            showEdit.value = false;
                            showInit.value = false;

                            knowledgeService.getKnowledge(CurSelectKey).then((res: any) => {
                                let jdata = JSON.parse(res.data.data);
                                kwge.title = jdata.title;
                                kwge.content = jdata.content;
                                kwge.updateTime = jdata.updateTime;
                            });
                        }
                    });
            }
            IsAdd = false;
        };

        const BtnAddTop = () => {
            modelOption.showTopModel = true;
        };

        const BtnAddSecond = () => {
            modelOption.showSecondModel = true;
            knowledgeService.getTopSelectOption().then((res: any) => {
                if (res.data.code == 200) {
                    modelOption.topSelectOption = JSON.parse(res.data.msg);
                    modelOption.topSelectVaule = modelOption.topSelectOption[0].value;
                }
            })
        };

        const AddTopNode = () => {
            console.log(modelOption.topInpuVaule)
            knowledgeService.AddTopNode(modelOption.topInpuVaule)
            modelOption.showTopModel = false;
            modelOption.topInpuVaule = ''
            UpdateTreeMenu()
        };
        const OnFlush = () => { UpdateTreeMenu() }

        const AddSecondNode = () => {
            console.log(modelOption.SecondInpuVaule)
            knowledgeService.AddSecondNode(modelOption.topSelectVaule || '', modelOption.SecondInpuVaule)
            modelOption.showSecondModel = false;
            modelOption.SecondInpuVaule = ''
            UpdateTreeMenu()
        }


        const OnDeleteClick = (isLeaf: boolean, treeKey: string) => {
            console.log(isLeaf)
            console.log(treeKey)
            knowledgeService.DelTreeNode(isLeaf, treeKey)
            UpdateTreeMenu()
        }


        return {
            expandedKeys,
            selectedKeys,
            searchValue,
            tlList,
            kwge,
            MdEdit,
            Editing,
            topCate,
            topOptions,
            OnDeleteClick,
            topSelectHandleChange,
            secondCate,
            secondOptions,
            // secondSelectHandleChange,
            tdList,
            autoExpandParent,
            showPriview,
            showInit,
            showEdit,
            modelOption,
            newTopvisible,
            OnEdit,
            OnFlush,
            AddSecondNode,
            OnSave,
            OnEditReturn,
            OnAdd,
            BtnAddSecond,
            AddTopNode,
            OnCancel,
            hideModal,
            BtnAddTop,
            // xxxtreeData,
            onSearch,
            onSelect,
            Onexpand,
            onSave,
        };
    },
});
</script>


<style>
.dTree {
    width: 300px;
    border: 2px solid rgb(240, 242, 245);
}
</style>