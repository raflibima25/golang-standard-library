package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` // structTag
}

type Person struct {
	Name  string `required:"true" max:"10"`
	No    string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		// get value structTag
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

// membuat func validasi
func isValid(value any) (result bool) {
	result = true
	valueType := reflect.TypeOf(value)
	for i := 0; i < valueType.NumField(); i++ {
		f := valueType.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Bima"})
	readField(Person{"Bima", "", "~"})

	person := Person{
		Name:  "ada",
		No:    "ada",
		Email: "ada",
	}

	fmt.Println(isValid(person))
}
