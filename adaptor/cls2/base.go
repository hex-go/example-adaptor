package cls2

import (
	"example/adaptor"
	"fmt"
)

// 定义 Cls2
type Cls2 struct {
	Name string
}

// 实现Class接口
func (g *Cls2) CreateUser(user string) (status bool, err error) {
	fmt.Println("Cls2 - create user: ", user)

	return true, nil
}

func (g *Cls2) DeleteUser(user string) (status bool, err error) {
	fmt.Println("Cls2 - Delete user: ", user)

	return true, nil
}

func (g *Cls2) Policies() (status bool, err error) {
	fmt.Println("Cls2 - get policies")

	return true, nil
}

func init() {
	// 在启动时注册类1工厂
	adaptor.Register("Cls2", func() adaptor.Adaptors {
		return new(Cls2)
	})
}
