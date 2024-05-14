package main

import "fmt"

func main() {
	// 初始化一个长度为3的整数数组
	a := [3]int{0, 1, 2}

	// 创建一个切片s，引用数组a的第0个元素到第1个元素（不包括索引2的元素）
	s := a[0:2]

	// 修改切片s的第一个元素，由于s是a的引用，a的第一个元素也会被修改
	s[0] = 11

	// 向切片s追加元素12，如果容量足够，直接追加，否则分配新的数组并复制现有元素
	s = append(s, 12)

	// 继续向切片s追加元素13，可能会导致s指向一个新的底层数组
	s = append(s, 13)

	// 再次修改切片s的第一个元素，此时不会影响原数组a
	s[0] = 21

	// 打印原数组a，查看其是否受到切片s的修改的影响
	fmt.Println(a) // 输出：[11 1 2]

	// 打印切片s，展示最终的元素
	fmt.Println(s) // 输出：[21 1 12 13]
}
