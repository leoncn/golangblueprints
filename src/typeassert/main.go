// typeassert project main.go
package main

import (
	"fmt"
)

func main() {

	num := 65
	i := MyInt{&num}

	var empty interface{} = &i

	fmt.Printf("%d\n", *i.ip)
	if sv, ok := empty.(Dog); ok {

		fmt.Println("Dog says", sv.Say())
	}
	fmt.Printf("%d\n", *i.ip)
	if sv, ok := empty.(Cat); ok {
		fmt.Println("Cat says", sv.Say())
	}

	ar := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sa := ar[0:0]
	for i := 0; i < 10; i++ {
		sa := ar[0 : i+1]
		for _, v := range sa {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}

	for _, v := range sa {
		fmt.Println(v)
	}

}

type Dog interface {
	Say() string
}

type Cat interface {
	Say() string
	Say2() string
}

type MyInt struct {
	ip *int
}

func (mi *MyInt) Say() string {
	j := 92
	mi.ip = &j
	return "MyInt Point"
}

//func (mi MyInt) Say() string {

//	return "my int instance"
//}
