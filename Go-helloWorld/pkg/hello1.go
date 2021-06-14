package pkg

import (
	"fmt"
	"os"
	"strings"
)
func man1(){
	fmt.Println(strings.Join(os.Args[1:]," "))
}
func man (){
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s+=sep+arg
		sep = " "
	}
	fmt.Println(s)
}
func echo(){
	for i:=1;i<len(os.Args);i++{
		fmt.Println(i)
	}
}

func Hello() string {
	//man()
	echo()
	fmt.Printf("Hello from a new package")
	var (
		s string = "Hello"
	)
	man1()
	fmt.Print("\n testing spaces \n")
	man()
	fmt.Println(os.Args[1:])
	return s
}