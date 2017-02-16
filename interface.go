package jwt

import jwt "github.com/dgrijalva/jwt-go"

// Token - Command for token
type Token interface {
	// Token -  取得Token
	Token() string
	// CreateWithCliams - 建立Token 使用 Cliams 結構
	CreateWithCliams(body jwt.Claims) Token
	// Parse - 解析 Token
	Parse(tokenJWT string) Token
	//
	ParseWithClaims(tokenJWT string, body jwt.Claims) Token
	// Result  - 取得解析後Token 的Release
	Result() interface{}
	// ToMap - 將結果轉為 map[string] string
	ToMap() map[string]string
	// Error - 回傳 error資訊
	Error() error
	// Encode - 把Token轉為md5
	Encode(jwtToken string) string
}
