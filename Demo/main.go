package main

import (
	"fmt"
	"log"

	go_mojoauth "github.com/mojoauth/go-sdk"
	"github.com/mojoauth/go-sdk/api"
	"github.com/mojoauth/go-sdk/mojoerror"
)

func main() {
	PasswordlessAuth()
}

func PasswordlessAuth() {
	var errors string
	//respCode := 200

	cfg := go_mojoauth.Config{
		ApiKey: "<Enter ApiKey>",
	}

	lrclient, err := go_mojoauth.NewMojoAuth(&cfg)
	if err != nil {
		errors = errors + err.(mojoerror.Error).OrigErr().Error()
		//		respCode = 500
	}

	res, err := api.Mojoauth{lrclient}.SigninWithMagicLink(
		[]byte(`{"email":"at@mail7.io"}`),
	)
	if err != nil {
		errors = errors + err.(mojoerror.Error).OrigErr().Error()
		//		respCode = 500
	}

	if errors != "" {
		log.Printf(errors)

		return
	}
	fmt.Println(res.Body)

}
