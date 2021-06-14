package api

import (
	"encoding/json"
	go_mojoauth "github.com/mojoauth/go-sdk"
	"github.com/mojoauth/go-sdk/httprutils"
	"github.com/mojoauth/go-sdk/mojobody"
	"log"

	"github.com/dgrijalva/jwt-go"

	"github.com/MicahParks/keyfunc"
)

type Mojoauth struct {
	Client *go_mojoauth.Mojoauth
}
func (mojo Mojoauth) VerifyEmailOTP(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := mojo.Client.NewPostReq("/users/emailotp/verify", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{}
		validatedQueries, err := httprutils.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

func (mojo Mojoauth) PingStatus(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"state_id": true,
	}
	validatedQueries, err := httprutils.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req := mojo.Client.NewGetReq("/users/status", validatedQueries)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
func (mojo Mojoauth) SigninWithMagicLink(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := mojo.Client.NewPostReq("/users/magiclink", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{}
		validatedQueries, err := httprutils.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
func (mojo Mojoauth) SigninWithEmailOTP(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := mojo.Client.NewPostReq("/users/emailotp", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"email": true,
		}
		validatedQueries, err := httprutils.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

//func (mojo Mojoauth) VerifyToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
//
//	request, err := mojo.Client.NewPostReqWithToken("/token/verify", body)
//	res, err := httprutils.TimeoutClient.Send(*request)
//	return res, err
//}
func (mojo Mojoauth) GetJwks() (*httprutils.Response, error) {


	req := mojo.Client.NewGetReq("/token/jwks", nil)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}



func JWTVerifier(jwtB64 string,body string)(*mojobody.TokenResponse, error) {
	var jwksJSON json.RawMessage = []byte(body)

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.New(jwksJSON)
	if err != nil {
		log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
		return nil, err
	}


	// Parse the JWT.
	token, err := jwt.Parse(jwtB64, jwks.KeyFunc)
	if err != nil {
		log.Fatalf("Failed to parse the JWT.\nError: %s", err.Error())
		return nil, err
	}

	// Check if the token is valid.
	response := mojobody.TokenResponse{IsValid: token.Valid, Token: jwtB64}

	return &response,nil

}

func (mojo Mojoauth) VerifyToken(token string)(*mojobody.TokenResponse, error){
	if mojo.Client.Context.Jwks != ""{
		res, err := JWTVerifier(token, mojo.Client.Context.Jwks)
		if err != nil{
			return nil, err
		}
		return res,nil



	}else{
		res, err := Mojoauth{mojo.Client}.GetJwks()
		if err != nil {
			return nil, err
			//		respCode = 500
		}else{
			mojo.Client.Context.Jwks = res.Body
			res, err := JWTVerifier(token, res.Body)
			if err != nil{
				return nil, err
			}
			return res,nil
		}
	}

}


