
// 消息队列
/*
	Go调度器
		客户端过来消息之后，会分配到一个消息队列，由该队列的worker进行处理，优先寻找空的调度器，
		如果所有的消息队列和worker都满了，则形成队列。

	1. 创建一个消息队列
		MsgHandler消息管理模块
		增加属性
			消息队列
				TaskQueue []chan ziface.IRequest

			worker工作池的数量（可以从全局配置中获取，也可以从配置中由用户设置）
				WorkerPoolSize uint32

	2. 创建多任务worker的工作池并启动
		创建工作池
			func (mh *MsgHandler) StartWorkerPool()

			根据WorkerPoolSize的数量去创建Worker
				func (mh *MsgHandler) StartOneWorker(workerID int, taskQueue chan ziface.IRequest)

			每个Worker都应该用一个Go去承载
				1.阻塞等待与当前Worker对应的channel的消息
				2.一旦有消息到来，Worker应该处理当前消息对应的业务，调用DoMsgHandler()

	3. 将之前的发送消息，全部改成把消息发送给消息队列和worker工作池来处理
		定义一个方法，将消息发送给消息队列工作池的方法
			func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest)

			1.保证每个Worker所受到的request任务是均衡的(平均分配)，让哪个Worker处理只需要将这个request请求
			发送给对应的taskQueue即可。
			
			2.将消息直接发送给对应的channel

	将消息队列机制集成到框架
		1.开启并调用消息队列及Worker工作池，保证WorkerPool只有一个，应该在创建server模块的时候开启(在server listen之前添加)

		2.将从客户端处理的消息，发送给当前的Worker工作池来处理(在已经处理完拆包，得到了request请求，交给工作池处理)
*/