
// 连接管理
/*
	创建一个连接管理模块
		属性
			已经创建的Connection集合
				connections map[uint32] ziface.IConnection
			针对map的互斥锁
				connLock    sync.RWMutex

		方法
			添加连接
			删除连接
			根据连接ID查找对应的连接
			总连接个数
			清理全部的连接
	
	将连接管理模块集成到zinx框架中

	给zinx框架提供 创建连接之后/销毁连接之前 所要处理的一些业务 提供给用户注册Hook函数
*/