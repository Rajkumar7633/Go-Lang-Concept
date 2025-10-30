package main

import (
	"fmt"
	// "image/color"
	// "os/user"

	"github.com/Rajkumar7633/podcast/auth"
	"github.com/Rajkumar7633/podcast/user"
	// "github.com/faith/color"
)






func  main()  {
	auth.LoginWithCredentials("rajcoder", "sceret")
	session := auth.GetSession()
	

	fmt.Println("session", session)

	user := user.User{
		Email : "user@email.com",
		Name : "jhon Doe",
	}

	fmt.Println(user.Email, user.Name)

	// color.Gray(user.Email)



}

