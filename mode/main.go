package main

import (
	"fmt"
	"github.com/marsli9945/learngo/mode/model"
)

func main() {
	// option模式适合单次创建对象，单次处理数据,快速创建对象，快速处理数据
	person := model.NewPerson(model.WithAge(18), model.WithName("ob"), model.WithAddress("china"))
	fmt.Printf("Person: %+v\n", person)

	// builder模式适合重复使用创建对象，批量处理数据入库
	personBuilder := model.NewPersonBuilder()
	person1 := personBuilder.WithAge(15).WithName("ob").WithAddress("china").Build()
	fmt.Printf("Person1: %+v\n", person1)
	person2 := personBuilder.WithAge(11).WithName("ob2").WithAddress("china2").Build()
	fmt.Printf("Person2: %+v\n", person2)
}
