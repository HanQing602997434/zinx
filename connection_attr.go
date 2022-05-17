
// 连接属性
/*
	给Connection模块添加可以配置属性的功能

	新增的属性
		连接属性集合--map
			property map[string]interface{}
		保护连接属性的互斥锁
			propertyLock sync.RWMutex

	新增的方法
		设置连接属性
			func (c *Connection) SetProperty(key string, value interface{})
		获取连接属性
			func (c *Connection) GetProperty(key string) (interface{}, error)
		移除连接属性
			func (c *Connection) RemoveProperty(key string)
*/