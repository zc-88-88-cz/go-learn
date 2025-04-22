package main

import "testing"

// 测试基础示例
func TestAddition(t *testing.T) {
    result := 1 + 1
    if result != 2 {
        t.Errorf("1 + 1 预期得到 2, 但实际得到 %d", result)
    }
}

// 表格驱动测试示例
func TestSubtraction(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{{
        name:     "正数相减",
        a:        5,
        b:        3,
        expected: 2,
    }, {
        name:     "负数相减",
        a:        -1,
        b:        -2,
        expected: 1,
    }}

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.a - tt.b
            if result != tt.expected {
                t.Errorf("%s: 预期 %d, 实际 %d", tt.name, tt.expected, result)
            }
        })
    }
}

// 基准测试示例
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fibonacci(20)
    }
}

func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}