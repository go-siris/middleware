package jwt_test

// Unlike the other middleware, this middleware was cloned from external source: https://github.com/auth0/go-jwt-middleware
// (because it used "context" to define the user but we don't need that so a simple siris.ToHandler wouldn't work as expected.)
// jwt_test.go also didn't created by me:
// 28 Jul 2016
// @heralight heralight add jwt unit test.
//
// So if this doesn't works for you just try other net/http compatible middleware and bind it via `siris.ToHandler(myHandlerWithNext)`,
// It's here for your learning curve.

import (
	"testing"

	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
	"github.com/go-siris/siris/httptest"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/go-siris/middleware/jwt"
)

type Response struct {
	Text string `json:"text"`
}

func TestBasicJwt(t *testing.T) {
	var (
		api             = siris.New()
		myJwtMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("My Secret"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})
	)

	securedPingHandler := func(ctx context.Context) {
		userToken := myJwtMiddleware.Get(ctx)
		var claimTestedValue string
		if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
			claimTestedValue = claims["foo"].(string)
		} else {
			claimTestedValue = "Claims Failed"
		}

		response := Response{"Iauthenticated" + claimTestedValue}
		// get the *jwt.Token which contains user information using:
		// user:= myJwtMiddleware.Get(ctx) or context.Get("jwt").(*jwt.Token)

		ctx.JSON(response)
	}

	api.Get("/secured/ping", myJwtMiddleware.Serve, securedPingHandler)
	e := httptest.New(api, t)

	e.GET("/secured/ping").Expect().Status(siris.StatusUnauthorized)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("My Secret"))

	e.GET("/secured/ping").WithHeader("Authorization", "Bearer "+tokenString).
		Expect().Status(siris.StatusOK).Body().Contains("Iauthenticated").Contains("bar")
}
