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

// 注册插件 提供给插件方注册使用，所谓的“插件”，就是一个定义为func() Class的普通函数，调用此函数，创建一个类实例，实现的工厂内部结构体会实现 Class 接口。
func Register(name string, factory func() Adaptors) {
	FactoryByName[name] = factory
}
