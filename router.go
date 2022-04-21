
// 基础路由模块
/*
	Request请求封装
		将连接和数据绑定在一起
		
		属性：
			连接IConnection
			请求数据

		方法：
			得到当前连接
			得到当前数据
			新建一个Request请求

	Router模块
		抽象的IRouter
			处理业务之前的方法
				PreHandle(request IRequest)

			处理业务的主方法
				Handle(request IRequest)

			处理业务之后的方法
				PostHandle(request IRequest)

		具体的BaseRouter
			处理业务之前的方法
				func (br *BaseRouter) PreHandle(request ziface.IRequest)

			处理业务的主方法
				func (br *BaseRouter) Handle(request ziface.IRequest)

			处理业务之后的方法
				func (br *BaseRouter) PostHandle(request ziface.IRequest)
*/