package models

import "encoding/json"

// UnmarshalToken Unmarshalls bytes to a Token struct
func UnmarshalToken(data []byte) (Token, error) {
	var r Token
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal marshals Token structs to bytes
func (r *Token) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Token is an access token issued by Azure
type Token struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	EXTExpiresIn string `json:"ext_expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	PwdExp       string `json:"pwd_exp"`
	PwdURL       string `json:"pwd_url"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}
