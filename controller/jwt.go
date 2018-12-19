package controller

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Kumi-comb/Daikon-BE/app"
	"github.com/Kumi-comb/Daikon-BE/security"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	uuid "github.com/satori/go.uuid"
)

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead.
//
// Note: the code below assumes the example is compiled against the master branch of goa.
// If compiling against goa v1 the call to jwt.New needs to be:
//
//    middleware := jwt.New(keys, ForceFail(), app.NewJWTSecurity())
func NewJWTMiddleware() (goa.Middleware, error) {
	keys, err := security.LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return jwt.New(jwt.NewSimpleResolver(keys), security.ForceFail(), app.NewJWTSecurity()), nil
}

// JWTController implements the jwt resource.
type JWTController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewJWTController creates a jwt controller.
func NewJWTController(service *goa.Service) (*JWTController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &JWTController{
		Controller: service.NewController("JWTController"),
		privateKey: privKey,
	}, nil
}

// Signin runs the signin action.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// JWTController_Signin: start_implement

	// Put your logic here
	// Generate JWT

	token := jwtgo.New(jwtgo.SigningMethodRS512)
	expire := time.Now().Add(time.Hour * 2).Unix()
	token.Claims = jwtgo.MapClaims{
		"iss":                   "kumi-comb.uths.xyz",             // who creates the token and signs it
		"aud":                   "kumi-comb.uths.xyz",             // to whom the token is intended to be sent
		"exp":                   expire,                           // time when the token will expire (10 minutes from now)
		"jti":                   uuid.Must(uuid.NewV4()).String(), // a unique identifier for the token
		"iat":                   time.Now().Unix(),                // when the token was issued/created (now)
		"nbf":                   2,                                // time before which the token is not yet valid (2 minutes ago)
		"sub":                   "AccessToken",                    // the subject/principal is whom the token is about
		"scopes":                "api:access",                     // token scope - not a standard claim
		"xyz.uths.kumi-comb.id": ctx.Payload.UserScreenName,
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	return ctx.NoContent()
	// JWTController_Signin: end_implement
}
