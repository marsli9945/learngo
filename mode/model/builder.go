package model

type PersonBuilder struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func (pb *PersonBuilder) WithName(name string) *PersonBuilder {
	pb.Name = name
	return pb
}
func (pb *PersonBuilder) WithAge(age int) *PersonBuilder {
	pb.Age = age
	return pb
}
func (pb *PersonBuilder) WithAddress(address string) *PersonBuilder {
	pb.Address = address
	return pb
}

func (pb *PersonBuilder) Build() *Person {
	return &Person{
		Name:    pb.Name,
		Age:     pb.Age,
		Address: pb.Address,
	}
}
