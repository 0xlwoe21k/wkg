interface Columns {
    title:string
    dataIndex?: string
    key: string
    align?: string
    fixed?:string
    width?:number
    minWidth?:number
    maxWidth?:number
    resizable?:boolean
    ellipsis?:boolean
}

class Types {
    getComapnyTableColumns() {
        let columns:Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                resizable:true,
                width: 50,
            },
            {
                title: "类型",
                dataIndex: "projectType",
                key: "2",
                align: 'center',
                minWidth:100,
                maxWidth:200,
                resizable:true,
                width: 100,
            },
            {
                title: "公司名称",
                dataIndex: "companyName",
                key: "3",
                align: 'center',
                resizable:true,

                width: 150,
                ellipsis: true,
            },
            {
                title: "域名",
                dataIndex: "domain",
                key: "4",
                align: 'center',
                resizable:true,

                width: 200,
                ellipsis: true,
            },
            {
                title: "SRC地址",
                dataIndex: "srcUrl",
                resizable:true,

                key: "5",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "关键字",
                dataIndex: "keyWord",
                key: "6",
                resizable:true,

                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "监控状态",
                resizable:true,

                dataIndex: "monitorStatus",
                key: "7",
                align: 'center',
                width: 90,
            },
            {
                title: "监控频率",
                resizable:true,

                dataIndex: "monitorRate",
                key: "8",
                align: 'center',
                width: 90,
            },
            {
                title: "创建时间",
                resizable:true,

                dataIndex: "lastUpdateTime",
                key: "9",
                align: 'center',
                width: 150,
            },
            {
                title: 'Action',
                resizable:true,

                key: 'operation',
                fixed: 'right',
                // align: 'center',
                width: 200,
            }
        ];
    }

    getDomainTableColumns() {
        let columns:Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 50,
            },
            {
                title: "Cid",
                dataIndex: "cid",
                key: "2",
                align: 'center',
                width: 100,
            },
            {
                title: "域名",
                dataIndex: "domain",
                key: "3",
                align: 'center',
                width: 300,
                ellipsis: true,
            },
            {
                title: "来源",
                dataIndex: "source",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "5",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: 'Action',
                key: 'operation',
                fixed: 'right',
                //  align: 'center',
                width: 200,
            }
        ];
    }

    getGatherDomainTableColumns() {
        let columns:Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 50,
            },
            {
                title: "Cid",
                dataIndex: "cid",
                key: "2",
                align: 'center',
                width: 100,
            },
            {
                title: "域名",
                dataIndex: "domain",
                key: "3",
                align: 'center',
                width: 300,
                ellipsis: true,
            },
            {
                title: "来源",
                dataIndex: "source",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "5",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            }
        ];
    }

    getWebsiteTableColumns() {
        let columns:Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "网站",
                dataIndex: "website",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "favicon",
                dataIndex: "favicon",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "标题",
                dataIndex: "title",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "ips",
                dataIndex: "ips",
                key: "5",
                align: 'center',
                width: 120,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 60,
                ellipsis: true,
            },
            {
                title: "指纹",
                dataIndex: "finger",
                key: "7",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            },
            {
                title: 'Action',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 200,
            }
        ];
    }

    
    getGatherWebsiteTableColumns() {
        let columns:Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "网站",
                dataIndex: "website",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "favicon",
                dataIndex: "favicon",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "标题",
                dataIndex: "title",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "ips",
                dataIndex: "ips",
                key: "5",
                align: 'center',
                width: 120,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 60,
                ellipsis: true,
            },
            {
                title: "指纹",
                dataIndex: "finger",
                key: "7",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            }
        
        ];
    }

    getIPsTableColumns() {
        let columns:Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 50,
            },
            {
                title: "ip",
                dataIndex: "ip",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "os",
                dataIndex: "os",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "NEW?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: 'Action',
                key: 'operation',
                fixed: 'right',
                 align: 'center',
                width: 200,
            }
        ];
    }

    getGatherIPsTableColumns() {
        let columns:Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 50,
            },
            {
                title: "ip",
                dataIndex: "ip",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "os",
                dataIndex: "os",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "NEW?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            }
        ];
    }
}




const types = new Types()

export default types;