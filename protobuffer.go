
// protobuffer
/*
	json
		优势：可读性比较强
		劣势：编解码比较耗时
		场景：web领域

	xml
		基于标签
		前段/网页

	protobuf
		优势：
			1.序列化后体积相比Json和XML很小，适合网络传输
			2.支持跨平台多语言
			3.消息格式升级和兼容性还不错
			4.序列化反序列化速度很快，快于Json的处理速度
		
		劣势：
			1.应用不够广(相比xml和json)
			2.二进制格式导致可读性差
			3.缺乏子描述
		可读性不强，传输过程中不是明文，二进制（序列化好的）
		场景：后端应用/微服务/服务器
*/