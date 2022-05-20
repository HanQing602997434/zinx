
// 多人在线游戏
/*
	协议的定义

	AOI兴趣点的算法
		AOI格子的数据类型Grid
			属性
				格子ID
					GID int
				格子的左边边界坐标
					MinX int
				格子的右边边界坐标
					MaxX int
				格子的上边边界坐标
					MinY int
				格子的下边边界坐标
					MaxY int
				当前格子内玩家或者物体成员的ID集合
					playerIDs map[int]bool
				保护当前集合的锁
					pIDLock sync.RWMutex

			方法
				初始化当前格子的方法
					func NewGrid(gID, minX, maxX, minY, maxY int) *Grid
				给格子添加一个玩家
					func (g *Grid) Add(playerID int)
				给格子删除一个玩家
					func (g *Grid) Remove(playerID int)
				得到当前格子中所有的玩家
					func (g *Grid) GetPlayerIDs() (playerIDs []int)
				调试使用——打印格子的基本信息
					func (g *Grid) String() string
		
		AOI管理格子(地图)数据类型AOIManager
			属性
				区域的左边界坐标
					MinX int
				区域的右边界坐标
					MaxX int
				X方向格子的数量
					CntsX int
				区域的上边界坐标
					MinY int
				区域的下边界坐标
					MaxY int
				Y方向格子的数量
					CntsY int
				当前区域中有哪些格子map-key = 格子ID, value = 格子对象
					grids map[int] *Grid

			方法
				初始化一个AOI管理区域模块
					func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager
				调式使用——打印当前AOI模块
					func (m *AOIManager) String() string
				得到每个格子在X轴方向的宽度
					func (m *AOIManager) gridWidth() int
				得到每个格子在Y轴方向的长度
					func (m *AOIManager) gridLength() int
				根据GID得到当前GID的九宫格的GID集合
					func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid)
				通过x，y横纵坐标得到GID格子编号
					func (m *AOIManager) GetGidByPos(x, y float32) int
				通过x，y横纵坐标得到周边九宫格全部的playerIDs
					func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int)
				添加一个PlayerID到一个格子中
					func (m *AOIManager) AddPidToGrid(pID, gID int)
				移除一个格子中的PlayerID
					func (m *AOIManager) RemovePidFromGrid(pID, gID int)
				通过GID获取全部的PlayerID
					func (m *AOIManager) GetPidsByGid(gID int) (playerIDs []int)
				通过坐标将一个Player添加到一个格子中
					func (m *AOIManager) AddToGridByPos(pID int, x, y float32)
				通过坐标将一个Player从一个格子中删除
					func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32)
	
	数据传输协议ProtocolBuffer

	玩家的业务
		玩家上线

		世界聊天

		上线位置的信息同步（玩家上线广播）

		移动位置与广播

		玩家下线

*/