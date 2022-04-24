
// 消息封装
/*
	定义一个消息的结构体Message
		属性：
			消息ID
			消息长度
			消息的内容

		方法：
			GetMsgId() uint32
			GetMsgLen() uint32
			GetData() []byte
			SetMsgId(uint32)
			SetDataLen(uint32)
			SetData([]byte)

	定义一个解决TCP粘包问题的封包拆包模块
		针对Message进行TLV格式的封装
			func (dp *DataPack) Pack(msg ziface.IMessage)([]byte, error)
			写message的长度
			写message的ID
			写message的内容

		针对Message进行TLV格式的拆包
			func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error)
			先读取固定长度的head --> 消息内容的长度和消息的类型
			再根据消息内容的长度，再次进行一次读写，从conn中读取消息的内容

	将消息封装机制集成到zinx框架中
		将Message添加到Request属性中
		修改连接读取数据的机制，将之前单纯的读取byte改成拆包形式的读取按照TLV形式读取
		给连接提供一个发包机制，将发送的消息进行打包再发送
*/