package main

import (
	"fmt"
)

type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

// 这里类型约束使用了空接口，代表的意思是所有类型都可以用来实例化泛型类型 Queue[T] (关于接口在后半部分会详细介绍）
type Queue[T interface{}] struct {
	elements []T
}

// 将数据放入队列尾部
func (q *Queue[T]) Put(value T) {
	q.elements = append(q.elements, value)
}

// 从队列头部取出并从头部删除对应数据
func (q *Queue[T]) Pop() (T, bool) {
	var value T
	if len(q.elements) == 0 {
		return value, true
	}

	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, len(q.elements) == 0
}

// 队列大小
func (q Queue[T]) Size() int {
	return len(q.elements)
}

func main() {
	type Slice[T int | float32 | float64] []T
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T", a) //输出：Type Name: Slice[int]
	fmt.Println()
	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T", b) //输出：Type Name: Slice[float32]

	fmt.Println()
	fmt.Println("===============================")

	type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

	// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
	var aMap MyMap[string, float64] = map[string]float64{
		"jack_score": 9.6,
		"bob_score":  8.4,
	}
	fmt.Println(aMap)

	fmt.Println("===============================")

	m := MySlice[int]{1, 2}.Sum()
	fmt.Println(m)

	fmt.Println("===============================")

	var q1 Queue[int] // 可存放int类型数据的队列
	q1.Put(1)
	q1.Put(2)
	q1.Put(3)
	fmt.Println(q1.Pop())
	fmt.Println(q1.Pop())
	fmt.Println(q1.Pop())

	var q2 Queue[string] // 可存放string类型数据的队列
	q2.Put("A")
	q2.Put("B")
	q2.Put("C")
	fmt.Println(q2.Pop())
	fmt.Println(q2.Pop())
	fmt.Println(q2.Pop())

	//var q3 Queue[struct{ Name string }]
	//var q4 Queue[[]int]    // 可存放[]int切片的队列
	//var q5 Queue[chan int] // 可存放int通道的队列
	//var q6 Queue[io.Reader]

}
