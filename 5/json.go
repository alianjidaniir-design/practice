package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var JSONrecord = `{ 
  "flag":true,
  "Array":["a","b","c"],
  "Entity" : {
  "a1" : "b1",
  "a2" : "b2",
  "Value":-326,
  "Null":null
},
"Message":"Hello Ali!"

}`

func typeswitch(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("it is string ", k, c)
		case int:
			fmt.Println("it is int ", k, c)
		case float64:
			fmt.Println("it is float64 ", k, c)
		case bool:
			fmt.Println("it is bool ", k, c)
		case map[string]interface{}:
			fmt.Println("it is map ", k, c)
			typeswitch(v.(map[string]interface{}))
		default:
			fmt.Printf("... Is %v : %T!\n", k, v)
		}
	}
	return
}

func exploremap(m map[string]interface{}) {
	for k, v := range m {
		embMap, ok := v.(map[string]interface{})
		if ok {
			fmt.Printf("{\"%v\": \n", k)
			exploremap(embMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v: %v\n", k, v)
		}

	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run json.go <jsonfile>")
	} else {
		JSONrecord = os.Args[1]
	}
	JSONMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(JSONrecord), &JSONMap)
	if err != nil {
		fmt.Println(err)
	}
	exploremap(JSONMap)
	typeswitch(JSONMap)
}
