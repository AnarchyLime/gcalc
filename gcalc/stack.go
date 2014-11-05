package gcalc

// func Push(a []string, s string) []string {
// 	a = append(a, s)
// 	return a
// }

// func Pop(a []string) ([]string, string) {
// 	l := len(a) - 1
// 	s := a[l]
// 	return a[:], s
// }

// func Top(a []string) string {
// 	return a[len(a)-1]
// }

type StringStack []string

type IStringStack interface {
	Push(s string) *StringStack
	Pop() string
	Top() string
	Count() int
}

func (a *StringStack) Push(s string) *StringStack {
	*a = append(*a, s)
	return a
}

func (a *StringStack) Count() int {
	return len(*a)
}

func (a *StringStack) Pop() string {
	i := len(*a) - 1
	v := (*a)[i]
	*a = (*a)[0:i]
	return v
}

func (a *StringStack) Top() string {
	return (*a)[len(*a)-1]
}
