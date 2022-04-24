
// 多路由模块
/*
	消息管理模块（支持多路由业务api调度管理的）--msgHandler
		属性：
			集合-消息ID和对应的router的关系-map

		方法：
			根据msgID来索引调度路由的方法
				func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest)

			添加路由方法到map集合中
				func (mh *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter)

	消息管理模块集成到zinx框架中
		1.将server模块中的Router属性替换成MsgHandler属性
		2.将server之前的AddRouter修改，调用MsgHandler的AddRouter
		3.将connection模块Router属性替换成MsgHandler，修改初始化Connection方法
		4.Connection之前调度Router的业务替换成MsgHandler调度，修改StartHandler
*/