package listops

type IntList []int

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

func (list *IntList) Append(list2 []int) []int {
	for i := len(list); i < (len(list) + len(list2)); i++ {
		list[i] = list2[i-len(list)]
	}
	return list
}
