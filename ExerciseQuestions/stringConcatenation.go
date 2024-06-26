package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 使用 + 连接字符串
func joinWithPlus(s1, s2 string) string {
	return s1 + s2
}

// 使用 fmt.Sprintf 连接字符串
func joinWithSprintf(s1, s2 string) string {
	return fmt.Sprintf("%s%s", s1, s2)
}

// 使用 strings.Builder 连接字符串
func joinWithBuilder(s1, s2 string) string {
	var builder strings.Builder
	builder.WriteString(s1)
	builder.WriteString(s2)
	return builder.String()
}

// 使用 strings.Join 连接字符串切片
func joinWithJoin(
	elements []string,
	sep string,
) string {
	return strings.Join(elements, sep)
}

// 使用 bytes.Buffer 连接字符串
func joinWithBuffer(s1, s2 string) string {
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	return buffer.String()
}

func main() {
	s1 := "Hello, "
	s2 := "world!"
	elements := []string{s1, s2}

	// 使用 + 连接
	// 小需求场景下使用 + 连接字符串
	fmt.Println("Join with +:", joinWithPlus(s1, s2))

	// 使用 fmt.Sprintf 连接
	fmt.Println("Join with fmt.Sprintf:", joinWithSprintf(s1, s2))

	// 使用 strings.Builder 连接
	// 此方法最高效 但是需要 Go 1.10 以上版本 是专门为高效构建和连接大量字符串而设计的
	fmt.Println("Join with strings.Builder:", joinWithBuilder(s1, s2))

	// 使用 strings.Join 连接
	fmt.Println("Join with strings.Join:", joinWithJoin(elements, ""))

	// 使用 bytes.Buffer 连接
	fmt.Println("Join with bytes.Buffer:", joinWithBuffer(s1, s2))
}
