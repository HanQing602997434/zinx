
// 读写分离模型
/*
	1.添加一个Reader和Writer之间通信的Channel
	
	2.添加一个Writer Goroutine

	3.当Reader由之前直接发送客户端改成发送给通信Channel

	4.启动Reader和Writer一同工作
*/