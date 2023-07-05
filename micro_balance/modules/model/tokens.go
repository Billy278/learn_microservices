package model

type TokenType string

const (
	ID_TOKEN      TokenType = "id_token"
	ACCESS_TOKEN  TokenType = "access_token"
	REFRESH_TOKEN TokenType = "refresh_token"
)

type Tokens struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type DefaultClaim struct {
	Expired   int       `json:"exp"`
	NotBefore int       `json:"nbf"`
	IssuedAt  int       `json:"iat"`
	Issuer    string    `json:"iss"`
	Audience  string    `json:"aud"`
	JTI       string    `json:"jti"`
	Type      TokenType `json:"typ"`
}

type IDclaim struct {
	IdUser   string `json:"preferrend_id"`
	Username string `json:"preferrend_username"`
	Email    string `json:"preferrend_email"`
}
type AccessClaim struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
