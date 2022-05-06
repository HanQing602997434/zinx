
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
				func (connMgr *ConnManager) Add(conn ziface.IConnection)
			删除连接
				func (connMgr *ConnManager) Remove(conn ziface.IConnection)
			根据连接ID查找对应的连接
				func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error)
			总连接个数
				func (connMgr *ConnManager) Len() int
			清理全部的连接
				func (connMgr *ConnManager) ClearConn()
	
	将连接管理模块集成到zinx框架中
		将ConnManager加入到Server模块中
			给server添加一个ConnMgr属性
			修改NewServer方法，加入ConnMgr初始化
			判断当前的连接数量是否已经超过最大值
		每次成功与客户端建立连接后，添加连接到ConnManager
			在NewConnection的时候将新的conn加入到ConnMgr中，需要给Connection加入隶属的Server属性，给Server提供一个GetConnMgr方法
		每次与客户端连接断开后，将连接从ConnManager删除
			在Conn.Stop()方法中，将当前连接从ConnMgr删除即可
			当server停止的时候应该清除所有的连接 Stop()方法中加入ConnMgr.ClearConn()

	给zinx框架提供 创建连接之后/销毁连接之前 所要处理的一些业务 提供给用户注册Hook函数
		添加的属性
			该server创建连接之后自动调用的Hook函数
			该server销毁连接之后自动调用的Hook函数

		添加的方法
*/