package main

type C1 interface {
	~int | ~int32
	M1()
}

type T struct{}

func (T) M1() {

}

type T1 int

func (T1) M1() {

}

// 泛型函数
func foo[P C1](t P) {

}

func main() {
	var t1 T1
	// 自动类型推导 foo[T1](t1)
	foo(t1)

	// var t T
	// foo(t) 报错
}

// 泛型类型
type Vector[T any] []T

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	return nil
}

// 实例化类型
var stringTree Tree[string]
