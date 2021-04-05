package api

import (
	go_mojoauth "github.com/mojoauth/go-sdk"
	"github.com/mojoauth/go-sdk/httprutils"
)

type Mojoauth struct {
	Client *go_mojoauth.Mojoauth
}

func (mojo Mojoauth) VerifyEmailOTP(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"otp": true,
	}
	validatedQueries, err := httprutils.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = mojo.Client.Context.ApiKey

	req := mojo.Client.NewGetReq("/users/magiclink/verify", validatedQueries)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
func (mojo Mojoauth) PingStatus(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"guid": true,
	}
	validatedQueries, err := httprutils.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = mojo.Client.Context.ApiKey

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
	request, err := mojo.Client.NewPostReq("/users/magiclink", body)
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

func (mojo Mojoauth) verifyToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := mojo.Client.NewPostReqWithToken("/token/verify", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"jwttoken": true,
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

