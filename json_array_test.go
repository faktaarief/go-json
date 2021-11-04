package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Fakta",
		MiddleName: "Arief",
		LastName:   "Farhanto",
		Hobbies:    []string{"Coding", "Traveling"},
	}

	bytes, _ := json.Marshal(&customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fakta","MiddleName":"Arief","LastName":"Farhanto","Address":null,"Hobbies":["Coding","Traveling"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Fakta",
		Addresses: []Address{
			{
				Street:     "Jalan Raya",
				Country:    "Indonesia",
				PostalCode: "1234",
			},
			{
				Street:     "Jalan Lintas Kota",
				Country:    "Indonesia",
				PostalCode: "4321",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fakta","MiddleName":"","LastName":"","Hobbies":null,"Addresses":[{"Street":"Jalan Raya","Country":"Indonesia","PostalCode":"1234"},{"Street":"Jalan Lintas Kota","Country":"Indonesia","PostalCode":"4321"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Raya","Country":"Indonesia","PostalCode":"1234"},{"Street":"Jalan Lintas Kota","Country":"Indonesia","PostalCode":"4321"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Raya",
			Country:    "Indonesia",
			PostalCode: "1234",
		},
		{
			Street:     "Jalan Lintas Kota",
			Country:    "Indonesia",
			PostalCode: "4321",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
