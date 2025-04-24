package model

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Option func(*Person)

func NewPerson(opts ...Option) *Person {
	p := &Person{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithName(name string) Option {
	return func(person *Person) {
		person.Name = name
	}
}

func WithAge(age int) Option {
	return func(person *Person) {
		person.Age = age
	}
}

func WithAddress(address string) Option {
	return func(person *Person) {
		person.Address = address
	}
}
