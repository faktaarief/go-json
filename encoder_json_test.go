package go_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName:  "Fakta",
		MiddleName: "Arief",
		LastName:   "Farhanto",
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}
