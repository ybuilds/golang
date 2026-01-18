package arrays

import "fmt"

type Product struct {
	name  string
	price int
}

func Driver() {
	hobbies := [3]string{
		"programming",
		"gaming",
		"photography",
	}

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	primeHobbies := hobbies[:2]
	fmt.Println(primeHobbies)

	primeHobbies = hobbies[1:]
	fmt.Println(primeHobbies)

	products := []*Product{
		{"Pen", 10},
		{"Book", 50},
	}

	products = append(products, &Product{"Tab", 15000}, &Product{"Laptop", 99000})

	for _, i := range products {
		fmt.Printf("%s: %d\n", i.name, i.price)
	}

	//make function
	username := make([]string, 2, 5) //Length = 2; Capacity = 5;
	username = append(username, "Yashas")
	username = append(username, "Megha")

	fmt.Println(username)
}
