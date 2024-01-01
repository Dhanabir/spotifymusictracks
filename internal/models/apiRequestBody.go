package models

type TokenAPIBody struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenResponseBody struct {
	AccessToken string `json:"access_token"`
	//TokenType   string `json:"token_type"`
	//ExpiresIn   int    `json:"expires_in"`
}
