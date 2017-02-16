package jwt

import (
	"errors"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var token string
var token2 string
var tokenwithclaims string

// BaseClaims -
type BaseClaims jwt.StandardClaims

// Standard - set the claims info
func (claims *BaseClaims) Standard() {
	claims.IssuedAt = time.Now().In(timeZone).Unix()
	claims.Issuer = "Cypress"
}

//Valid - valid the issuer
// TODO:  是否需要增加驗證給予的來源
func (claims BaseClaims) Valid() error {
	if claims.Issuer != "Cypress" {
		return errors.New("Not the Right Issuer")
	}
	return nil
}

// Info - 加上額外的 INFO
func (claims *BaseClaims) Info() {

}

// UserClaims - user token struct
type UserClaims struct {
	Account  string `json:"account" form:"account" validate:"required"`
	Owner    string `json:"owner" form:"owner" validate:"required"`
	Currency string `json:"currency" form:"currency" validate:"required"`
	Parent   string `json:"parent" form:"parent" validate:"required"`
	Nickname string `json:"nickname" form:"nickname" validate:"required"`
	UserID   string `json:"userid" form:"userid" validate:"required"`
	Type     string `json:"type" form:"type" validate:"required"`
	IP       string `json:"ip" form:"ip" validate:"required"`
	Location string `json:"location" form:"location" validate:"required"`
	BaseClaims
}

func init() {
	SetKey("123993929410293")
	token2 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBIjoiamozayIsIkFnZSI6IjE4IiwiQiI6ImpqM2siLCJDIjoiamozayIsIkQiOiJqajNrIiwiRSI6ImpqM2siLCJOYW1lIjoiSmltbXkiLCJTY2hvb2wiOiJOVVRDIiwiaWF0IjoiMjAxNy0wMi0wOVQwNzo0MDoyMS43OTUwNjA3MTMtMDQ6MDAiLCJpc3MiOiJDeXByZXNzIEpXVCJ9.ohxS68mCT9IIcwydM4f4ZVuJOVZH3InRwmnM2SnIar3"
	SetTimeZone(time.FixedZone("UTC-4", -4*60*60))
}

func Test_CreateWithClaims(t *testing.T) {
	body := &UserClaims{}
	body.Account = "12344566"
	body.Currency = "CNY"
	body.IP = "1.2.3.4"
	body.Location = "TW"
	body.Nickname = "J"
	body.Owner = "CQ9"
	body.Parent = "CQ9"
	body.Type = "Player"
	body.UserID = "jjrkeoigjalskdgji"
	body.Standard()
	token := New()
	token.CreateWithCliams(body)

	if token.Error() != nil {
		t.Error(token.Error())
	}
	tokenwithclaims = token.Token()
	t.Log(tokenwithclaims)
}

func Test_ParseWithClaims(t *testing.T) {
	token := New()
	user := &UserClaims{}
	token.ParseWithClaims(tokenwithclaims, user)
	if token.Error() != nil {
		t.Error(token.Error())
	} else {
		t.Log(user)
	}
}

func Test_Parse(t *testing.T) {
	token := New()
	token.Parse(tokenwithclaims)
	if token.Error() != nil {
		t.Error(token.Error())
	} else {
		t.Log(token.ToMap())
	}
}

func Test_ParseWithClaimsFailed(t *testing.T) {
	token := New()
	body := &UserClaims{}
	body.Standard()
	body.BaseClaims.Issuer = "xxx"

	token.CreateWithCliams(body)
	user := &UserClaims{}
	token.ParseWithClaims(token.Token(), user)
	if token.Error() == nil {
		t.Error("Valid Failed")
	} else {
		t.Log(token.Error())
	}
}

func Test_EncodeJWT(t *testing.T) {
	token := New()
	result := token.Encode(tokenwithclaims)
	t.Log(result)

}

func Benchmark_CreateTokenWithClaims(b *testing.B) {
	b.StopTimer()
	body := &UserClaims{}
	body.Account = "12344566"
	body.Currency = "CNY"
	body.IP = "1.2.3.4"
	body.Location = "TW"
	body.Nickname = "J"
	body.Owner = "CQ9"
	body.Parent = "CQ9"
	body.Type = "Player"
	body.UserID = "jjrkeoigjalskdgji"
	body.Standard()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		token := New()
		token.CreateWithCliams(body)
	}
}

func Benchmark_ParseToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		token := New()
		token.Parse(tokenwithclaims)
	}
}

func Benchmark_ParseWithClaims(b *testing.B) {
	b.StopTimer()
	user := &UserClaims{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		token := New()
		token.ParseWithClaims(tokenwithclaims, user)
	}

}
