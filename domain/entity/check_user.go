package entity

import "fmt"

func Test_user() {
	var u User
	u.SetPassword("heyyou")
	err := u.ComparePassword("heyyou")
	fmt.Println("THIS ISS A TESSSSSSSSS OT OF USEEERRRR")
	fmt.Println(err)
	fmt.Println("THIS ISS A TESSSSSSSSS OT OF USEEERRRR")
}
