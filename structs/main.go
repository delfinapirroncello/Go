package main

import (
	"encoding/json"
	"fmt"

	"github.cpm/digitalhouse-tech/go-course/structs/commerce"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address, omitempty"`
	Age      uint8  `json:"age"`
	LastName string `json:"last_name"`
}

func (u User) Display() {
	v, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(v))
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	user := User{
		ID:      123,
		Name:    "Nahuel",
		Address: "Calle falsa 123",
	}
	user.Display()
	user.SetName("Azul")
	user.Display()

	p1 := commerce.Product{
		Name:  "Heladera Gafa",
		Price: 2000000,
		Type: commerce.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags:  []string{"heladera", "freezer", "gafa", "refrigerador"},
		Count: 5,
	}

	p2 := commerce.Product{
		Name:  "Naranja",
		Price: 50,
		Type: commerce.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"alimento", "verdura"},
		Count: 20,
	}

	p3 := commerce.Product{
		Name:  "Cortina",
		Price: 6000,
		Type: commerce.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"hogar", "cortina"},
		Count: 3,
	}

	car := commerce.NewCar(11312)
	car.AddProducts(p1, p2, p3)

	fmt.Println("PRODUCTS CAR")
	fmt.Println("Total Products: ", len(car.Products))
	fmt.Printf("Total Car &.2f\n", car.Total())
}
