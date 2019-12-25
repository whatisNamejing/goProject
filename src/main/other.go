package main

import (
	"fmt"
	"main/sout"
)

func main() {
	a := 42
	p := &a
	fmt.Println(*p)

	person := Person{
		name:"xxx",
		age:12,
	}

	fmt.Println(person)

	person.age = 15
	fmt.Println(person)

	sout.SystemOut("ss")

	str := "afsfasfdasfasdf"
	for i:=0;i< len(str);i++ {
		fmt.Println(string(str[i]))
	}
}

type Person struct {
	name string
	age int16

}
