package adaptor

// 定义接口
type Adaptors interface {
	CreateUser(user string) (status bool, err error)
	DeleteUser(user string) (status bool, err error)
	Policies() (status bool, err error)
}

var (
	// 插件字典
	FactoryByName = make(map[string]func() Adaptors)
)

// 注册插件
func Register(name string, factory func() Adaptors) {
	FactoryByName[name] = factory
}
