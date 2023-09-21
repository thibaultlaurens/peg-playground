package main

import (
	"encoding/json"
	"fmt"
	"pegplayground/query"
)

func main() {
	q := "item.name=laptop item.spec.ram > 8gb  ( item.spec.ssd=yes | item.spec.ssd.capacity > 512gb)  ( item.maker=asus | item.maker=coconics )"
	res, err := query.Parse("", []byte(q))
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := json.MarshalIndent(res, "", "   ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("output:\n%v\n", string(result))
}
