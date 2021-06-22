package main

import "fmt"

func main()  {

	s := "hello"
	var p *string = &s
	fmt.Printf("%T %v\n",s,s)
	fmt.Printf("%T %v\n",&s,&s)
	fmt.Printf("%T %v\n",p,p)
	fmt.Printf("%T %v\n",*p,*p)





}
