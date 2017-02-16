package jwt

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// key - JWT encrypt key
var key []byte

// method - jwt siging method
var method = jwt.SigningMethodHS256

// timeZone -
var timeZone *time.Location

// jwtToken - token for jwt
type jwtToken struct {
	t      *jwt.Token
	token  string
	result interface{}
	err    error
}

// New - new token struct
func New() Token {
	return &jwtToken{}
}

func (t *jwtToken) CreateWithCliams(body jwt.Claims) Token {
	t.t = jwt.NewWithClaims(method, body)
	t.t.SignedString(key)
	t.token, t.err = t.t.SignedString(key)
	return t
}

func (t *jwtToken) ParseWithClaims(tokenJWT string, body jwt.Claims) Token {
	t.t, t.err = jwt.ParseWithClaims(tokenJWT, body, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Failed signing method")
		}
		return key, nil
	})

	if _, ok := t.t.Claims.(jwt.Claims); ok && t.t.Valid {
	} else {
		t.err = errors.New("Failed Valid")
	}

	return t
}

func (t *jwtToken) Create(body jwt.MapClaims) Token {
	t.t = jwt.NewWithClaims(method, body)
	t.t.SignedString(key)
	t.token, t.err = t.t.SignedString(key)
	return t
}

func (t *jwtToken) Parse(tokenJWT string) Token {
	t.t, t.err = jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Failed signing method")
		}
		return key, nil
	})
	if t.err != nil {
		return t
	}
	if claims, ok := t.t.Claims.(jwt.MapClaims); ok && t.t.Valid {
		t.result = claims
	} else {
		t.err = errors.New("Valid Failed")
	}
	return t
}

func (t *jwtToken) Token() string {
	return t.token
}

func (t *jwtToken) Result() interface{} {
	return t.result
}

func (t *jwtToken) ToMap() map[string]string {
	if t.result != nil {
		result := make(map[string]string)
		for i, v := range t.result.(jwt.MapClaims) {
			result[i] = toString(v)
		}
		return result
	}
	t.err = errors.New("Faild Change")
	return nil
}

func (t *jwtToken) Error() error {
	return t.err
}

// Encode - 加密Token
func (t *jwtToken) Encode(jwtToken string) string {
	m := md5.New()
	m.Write([]byte(jwtToken))
	return hex.EncodeToString(m.Sum(nil))
}

//SetKey - set the token key
func SetKey(authkey string) {
	key = []byte(authkey)
}

//SetTimeZone - set the time zone
func SetTimeZone(tz *time.Location) {
	timeZone = tz
}

// toString - 轉換為字串
func toString(args interface{}) string {
	switch args.(type) {
	case string:
		return args.(string)
	case float64:
		return strconv.FormatFloat(args.(float64), 'f', 0, 64)
	default:
		return ""
	}
}
