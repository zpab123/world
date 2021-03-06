// /////////////////////////////////////////////////////////////////////////////
// 数据库

package db

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	dbEngine IEntityStorage // 数据库引擎
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 启动数据库
func Run() {
	// 创建引擎
	createEngine()

	go mainLoop()
}

// /////////////////////////////////////////////////////////////////////////////
// private api

// 创建引擎
func createEngine() error {
	if nil == dbEngine {
		return nil
	}

}

// 主循环
func mainLoop() {

}
