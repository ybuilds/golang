package main

import "fmt"

func displayValues(value any) {
	switch value.(type) {
	case int:
		{
			fmt.Printf("Integer: %v\n", value)
			actual, ok := value.(int)
			if ok {
				fmt.Printf("Integer + 1: %v\n", actual+1)
			}
		}
	case float64:
		{
			fmt.Printf("Float: %v\n", value)
			actual, ok := value.(float64)
			if ok {
				fmt.Printf("Float64 + 5: %v\n", actual+5.0)
			}
		}
	case string:
		{
			fmt.Printf("String: %v\n", value)
		}
	default:
		{
			fmt.Printf("Type not recognised\n")
		}
	}
}

func add[T int | int32 | int64 | float32 | float64](a, b T) T {
	return a + b
}

type User struct {
	Name string
	Age  int
}

type FetchAge interface {
	getAge() int
}

func (user *User) getAge() int {
	return user.Age
}

func addAge[T FetchAge](a, b T) int {
	return a.getAge() + b.getAge()
}

func main() {
	yashas := &User{"Yashas", 25}
	megha := &User{"Megha", 25}

	displayValues(1)
	displayValues(10.0)
	displayValues("Yashas")
	displayValues(yashas)

	res := add(10, 12)
	fmt.Println(res)

	age := addAge(yashas, megha)
	fmt.Println(age)
}
