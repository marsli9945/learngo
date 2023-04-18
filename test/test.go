package main

type res struct {
	code    int
	message string
}

func a(i ...interface{}) {

}

func main() {
	a(&res{})
}
