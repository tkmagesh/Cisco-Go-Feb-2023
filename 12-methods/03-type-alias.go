package main

import "fmt"

//type alias
type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	str := MyString("Consectetur quis veniam consectetur id dolor consectetur tempor et cillum fugiat ipsum. Culpa proident dolore sunt aliquip ullamco id. Commodo nulla qui sunt culpa nulla est. Exercitation aliqua labore non eiusmod nostrud culpa.")
	fmt.Println(str.Length())
}
