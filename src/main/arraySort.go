package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	nums := []int{12,20,5,16,15,1,30,45,23,9,99,22,33,10,0,-1,99,88,7,66,5,44,33,21,4,34,65,87,534,13,345,745,756,2342,243,546,65,2342,546,76,878,11,22,3,8,44,987,24,123,12,45,56,67,68,54,53,5,35,2,2,35,345,346,5,7,68,67,98,766,8,67,96,8,567,67867,4,7345,6,45,675,46,4,7,678,4,6,456,45,456,54,7,5,7,345,5345,765,756}

	current := time.Now().UnixNano()
	pop(nums)
	unix := time.Now().UnixNano()
	fmt.Printf("冒泡时间差为 %d \n",unix-current)
	str := "["
	str = arrayToString(nums, str)
	fmt.Printf(str)

	sendHttpTOBaiDu()

}

func sendHttpTOBaiDu() {
	url := "http://www.baidu.com"
	request,error := http.NewRequest("get",url,nil)
	if error != nil {
		panic(error)
	}
	client := http.Client{}
	response,_ := client.Do(request);
	stdOut := os.Stdout
	_,err := io.Copy(stdOut,response.Body)
	if error != nil {
		 panic(err)
	}
	fmt.Println(response.Status)
}

func arrayToString(nums []int, str string) string {
	for i := 0; i < len(nums); i++ {
		tem := strconv.Itoa(nums[i])
		if i == len(nums)-1 {
			str = str + tem + "]"
		} else {
			str = str + tem + ","
		}
	}
	return str
}

func pop(nums []int) {

	for i := 0; i < len(nums);i++  {
		for j := 0; j < len(nums) - i-1; j++ {
			if nums[j] > nums[j+1]{
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
}
