package testing

import "testing"

func TestAdd(t *testing.T) {
	// t *testing.T 是测试的“控制器”，我们用它来报告测试状态

	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		// t.Errorf 会记录一个错误，但测试会继续执行
		t.Errorf("Add(2, 3) = %d; want %d", sum, expected)
	}

	// 如果想在出错时立即停止当前测试，可以使用 t.Fatalf()
}
