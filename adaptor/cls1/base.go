package cls1

import (
	"example/adaptor"
	"fmt"
)

// 定义 Cls1
type Cls1 struct {
	Name string
}

// 实现Class接口
func (g *Cls1) CreateUser(user string) (status bool, err error) {
	fmt.Println("Cls1 - create user: ", user)

	return true, nil
}

func (g *Cls1) DeleteUser(user string) (status bool, err error) {
	fmt.Println("Cls1 - Delete user: ", user)

	return true, nil
}

func (g *Cls1) Policies() (status bool, err error) {
	fmt.Println("Cls1 - get policies")

	return true, nil
}

func init() {
	// 导入包时 注册 cls1， func() 约定的返回值为 adaptor.Adaptors，所以 new(Cls1)使得Cls1必须集成interface
	adaptor.Register("Cls1", func() adaptor.Adaptors {
		return new(Cls1)
	})
}
