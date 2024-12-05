package main

import (
	"encoding/json"
	"fmt"
)

// import "encoding/json"

type myStruct struct {
	IntValue int `json:"intVal"`
	MyWish *int // this is either int or nil (passing a reference)
	EmptyString     string    `json:"emptyString,omitempty"` // json will remove this field if the val is empty
}

func main() {
	data := map[string]interface{}{
		"myName": "Karthik", 
		"myAge": 21, 
		"favCountry": "USA", 
		"favDestination": "Positano",
		"myMap": map[int]interface{}{
			23: "karthik", 
			12: []bool{false, true},
		},
	}

	//Converting into json
	jsonData, _ := json.Marshal(data)
	//json data has to be converted to string 

	fmt.Printf("%s", jsonData) // %s does string(jsonData)

	//converting struct type into json
	var ptrVal int = 23
	data1 := myStruct{IntValue: 24, MyWish: &ptrVal, EmptyString: "ada"}

	jsonData1, _ := json.Marshal(data1)
	fmt.Println(string(jsonData1))


}

func unMarshal() {
	jsonData := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null
		}
	`

	var data map[string]interface{};
	json.Unmarshal([]byte(jsonData), &data) // pass pointer

	fmt.Println(data)

}