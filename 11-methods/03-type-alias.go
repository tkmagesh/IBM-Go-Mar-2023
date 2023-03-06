package main

import "fmt"

type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	str := MyString("Laborum proident minim aliquip commodo nisi qui ea anim excepteur anim irure est. Id ipsum enim elit culpa. Mollit nulla veniam reprehenderit velit eu aliquip id commodo consequat ipsum exercitation. Aute culpa qui quis occaecat veniam ipsum laborum. Exercitation adipisicing qui aute consectetur aliqua culpa in irure aliqua occaecat minim laboris sunt. Sit elit cillum excepteur minim incididunt ut pariatur officia ullamco anim incididunt enim.")
	fmt.Println(str.Length())
}
