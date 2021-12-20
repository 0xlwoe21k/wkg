interface Children{
	key: string
	path: string
	title: string
}

interface Menu {
	key?: string
	path?: string
	icon?: string
	title?: string
	children?:Array<Children>
}

export const menuList:Array<Menu> = [
	{
		key: '1',
		path: '/index/dashboard',
		icon: 'dashboard',
		title: 'Dashboard',
	},
	{
		key: '2',
		title: '在线扫描',
		children: [{
			key: '2.1',
			path: '/index/agentManage',
			title: '漏洞扫描'
		},
		{
			key: '2.2',
			path: '/index/agentManage',
			title: '域名扫描'
		}]
	},
	{
		key: '3',
		path: '/index/taskManage',
		title: 'SRC管理',
		children: [{
			key: '3.1',
			path: '/index/src/gather',
			title: '资产总览'

		},
		{
			key: '3.2',
			path: '/index/src/company',
			title: '公司管理'
		},
		{
			key: '3.3',
			path: '/index/src/domain',
			title: '域名管理'
		},
		{
			key: '3.4',
			path: '/index/src/website',
			title: '站点管理'
		},
		{
			key: '3.5',
			path: '/index/src/ips',
			title: 'IP管理'
		},
		{
			key: '3.6',
			path: '/index/src/service',
			title: '服务管理'
		},
		{
			key: '3.7',
			path: '/index/src/littleProgram',
			title: '小程序管理'
		},
		{
			key: '3.8',
			path: '/index/src/news',
			title: '资讯管理'
		},
		{
			key: '3.9',
			path: '/index/src/webChatAccount',
			title: '微信公众号管理'
		}]
	},
	{
		key: '4',
		path: '/index/dashboard',
		title: '信息动态',
	},
	{
		key: '5',
		path: '/index/task',
		title: '任务管理',
		children: [{
			key: '5.1',
			path: '/index/task/taskManage',
			title: '任务列表'
		},]
	},
	{
		key: '6',
		path: '/index/dashboard',
		title: '漏洞管理',
		children: [{
			key: '6.1',
			path: '/index/vuln/poc',
			title: 'POC管理'
		}]
	},
	{
		key: '7',
		path: '/knowledge',
		title: '知识库',
		children: [{
			key: '7.1',
			path: '/index/knowledge/knowledge',
			title: '漏洞知识库'
		},
		{
			key: '7.2',
			path: '/index/knowledge/config',
			title: '配置'
		}]
	},
	{
		key: '8',
		path: '/index/dashboard',
		title: '系统设置',
		children: [{
			key: '8.1',
			path: '/index/agentManage',
			title: '扫描配置'
		},
		{
			key: '8.1',
			path: '/index/agentManage',
			title: '通知配置'
		}]
	}
];