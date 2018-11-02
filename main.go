// Go project by arthur
// blihUI
// 2018

package main

import (
	"blihUI/pkg/blih"

	"blihUI/pkg/user"
)

const (
	email = ""
	baseURL = "https://blih.epitech.eu/"
	baseUserAgent = "blih-1.7"
)

func main() {
	u := user.New(email)
	_ = blih.New(baseURL, u, false, baseUserAgent)
}
