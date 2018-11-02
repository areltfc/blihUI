// Go project by arthur
// blihUI
// 2018

package main

import (
	"blihUI/pkg/blih"
	"blihUI/pkg/repository"
	"blihUI/pkg/user"
	"fmt"
)

const email = "arthur.delattre@epitech.eu"

func main() {
	u := user.New(email)
	b := blih.New("https://blih.epitech.eu/", u, false, "blih-1.7")
	repo, err := repository.Info("Conte2017", &b)
	if err != nil {
		panic(err)
	}
	fmt.Println(repo)
}
