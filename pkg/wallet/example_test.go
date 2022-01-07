package wallet

import "fmt"

func ExampleFindAccountById_FoundAccount() {
	acc, _ := FindAccountById(1)
	fmt.Println(acc.ID)
	// Output:
	// 1
	//
}

func ExampleFindAccountById_NotFoundAccount() {
	_, err := FindAccountById(2)
	fmt.Println(err)
	// Output:
	// account not found
	//
}
