package jwt_test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/northwindman/sso/internal/domain/models"
	myJWT "github.com/northwindman/sso/internal/lib/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewToken(t *testing.T) {
	user := models.User{
		ID:    1, // int64
		Email: "user@example.com",
	}
	app := models.App{
		ID:     145, // int
		Secret: "secret123",
	}
	duration := time.Hour

	t.Run("successfully create token", func(t *testing.T) {

		tokenString, err := myJWT.NewToken(user, app, duration)

		assert.NoError(t, err)          // check that err is nil
		assert.NotEmpty(t, tokenString) // check that token is not empty string

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(app.Secret), nil
		})

		assert.NoError(t, err)

		claims, ok := token.Claims.(jwt.MapClaims)
		assert.True(t, ok)
		uid, err := getInt64Claims(claims, "uid")
		assert.NoError(t, err)
		assert.Equal(t, user.ID, uid)
		assert.Equal(t, user.Email, claims["email"])
		appID, err := getInt64Claims(claims, "app_id")
		assert.NoError(t, err)
		assert.Equal(t, app.ID, appID)

		exp := claims["exp"].(float64)
		assert.WithinDuration(t, time.Now().Add(duration), time.Unix(int64(exp), 0), duration)
	})
}

func getInt64Claims(claims jwt.MapClaims, key string) (int64, error) {
	if val, ok := claims[key]; ok {
		return int64(val.(float64)), nil
	}
	return 0, fmt.Errorf("claim %s is not valid float64", key)
}
