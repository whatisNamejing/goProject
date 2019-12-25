package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

func main() {

	for i:=0;i<100 ;i++  {
		random := randomNum()
		UUID := newUUID()
		fmt.Printf("index = %d ==== uuid = %s ---- random = %d \n",i,UUID,random)
	}

	Time := time.Now().Unix()
	fmt.Println(Time*1000)

	temp := time.Unix(Time,0)
	str := temp.Format("2006-01-02 15:04:05")
	fmt.Println(str)
}

func randomNum () int  {
	rand := rand.Intn(10)
	return rand
}

func newUUID() string {
	uuId,error := uuid.NewV4()
	if error != nil {
		panic(error)
	}
	return uuId.String()
}
