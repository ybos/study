package counter

type alertCounter int
type NoticeCounter int

func NewAlter(u int) alertCounter {
	return alertCounter(u)
}

// 该类型只公开了部分标识符
type Test struct {
	Name  string
	email string
}

type user struct {
	Name  string
	Email string
}

// 这里的 Admin 是公开的
// Rights 也是公开的
// user 是未公开的，无法被结构字面量初始化，但是创建的时候会被默认值初始化，所以可以继续使用
// 后续在使用的时候，因为内部类型的标识符会被提升到外部类型上，所以可以通过外部类型直接访问
type Admin struct {
	user
	Rights int
}
