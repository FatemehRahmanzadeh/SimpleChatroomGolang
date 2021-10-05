package auth

import (
	"fmt"
	"time"

	"github.com/FatemehRahmanzadeh/chat_sample/models"
	"github.com/dgrijalva/jwt-go"
)

// define a secret for jwt
const hmacSecret = "WjdwZUh2dWJGdFB1UWRybg=="
const defaulExpireTime = 604800 // 1 week

// structure to parse user information
type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (c *Claims) GetId() string {
	return c.ID
}

func (c *Claims) GetUserame() string {
	return c.Username
}

// CreateJWTToken generates a JWT signed token for the given user
// using user information and our defined expire time
func CreateJWTToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":        user.GetId(),
		"Username":  user.GetUserame(),
		"ExpiresAt": time.Now().Unix() + defaulExpireTime,
	})
	tokenString, err := token.SignedString([]byte(hmacSecret))

	return tokenString, err
}

// checking user token validation if expired or not valid return an error
func ValidateToken(tokenString string) (models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		//validation of token algorithm to be as expected:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSecret), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
