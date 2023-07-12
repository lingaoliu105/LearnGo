package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg)) //like float64, etc
	fmt.Println("value: ", reflect.ValueOf(arg))
}

type User struct {
	Id      int    `info:"id" doc:"identification number" json:"id"`
	Name    string `info:"name" json:"name"`
	Age     int    `json:"age"`
	friends []User `json:"friends"`
}

func (u User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", u)
}
func main() {
	//var num float64 = 1.14514
	//reflectNum(num)
	user := User{1, "Raiden", 500, []User{}}
	//DoFieldAndMethod(user)
	//user.Call()

	//using tags
	//findTag(&user)

	//using json

	//encoding struct->json
	jsonStr, err := json.Marshal(user)
	if err != nil {
		println("error")
		return
	}
	fmt.Printf("%s\n", jsonStr) //jsonStr is uint8 array, println will only print those numbers

	//json->string
	srcJson := "{\"id\":1,\"name\":\"Raiden\",\"age\":500}"
	newUser := User{}
	err = json.Unmarshal([]byte(srcJson), &newUser)
	if err != nil {
		fmt.Println("unmarshal error")
	}
	fmt.Printf("%v\n", newUser)
}

func findTag(stru interface{}) {
	t := reflect.TypeOf(stru).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagStr := t.Field(i).Tag.Get("info")
		fmt.Println(tagStr)

	}
}

func DoFieldAndMethod(input interface{}) {
	//get input type
	inputType := reflect.TypeOf(input)
	fmt.Println(inputType.Name()) //user .Name() to get the name
	//get input value
	inputValue := reflect.ValueOf(input)
	fmt.Println(inputValue)
	//get fields of the input type
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)              //get ith field of inputtype
		value := inputValue.Field(i).Interface() //note that type.field (return structfield) is different from value.field (return value)
		fmt.Println(field.Name, field.Type, value)
	}

	//get methods from type
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println(m, m.Name, m.Type)
	}
}
