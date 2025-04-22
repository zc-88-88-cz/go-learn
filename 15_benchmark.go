package main

import (
	"testing"
	"time"
)

// 基准测试示例
func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = joinStrings("Hello", " ", "World")
	}
}

func joinStrings(strs ...string) string {
	var result string
	for _, s := range strs {
		result += s
	}
	return result
}

// 优化后的字符串拼接
func BenchmarkStringBuilderJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = joinStringsWithBuilder("Hello", " ", "World")
	}
}

func joinStringsWithBuilder(strs ...string) string {
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// 内存分配分析示例
func BenchmarkAllocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = make([]int, 100)
	}
}

// 并行基准测试
func BenchmarkParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = time.Now().Unix()
		}
	})
}