package main

import(
    "encoding/json"
    "fmt"
)

type Person struct{
	FirstName string `json: "firstname"`
	LastName string `json: "lastname"`
}

func main (){
	in := `{"firstname" :"John", "lastname":"Dow"}`
	bytes := []byte(in)

	var p Person
	err := json.Unmarshal(bytes, &p)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%+v", p)
}