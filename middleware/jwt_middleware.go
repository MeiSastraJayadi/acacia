package middleware

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type UseJWTMiddlewarewithECDSA struct {
  middleware *UseMiddleware
  privateKey string
}

func decodeEC(privateKey string) (*ecdsa.PrivateKey, error) {
  block, _ := pem.Decode([]byte(privateKey))
  Bytes := block.Bytes
  x509Block, err := x509.ParseECPrivateKey(Bytes)
  if err != nil {
    return nil, err
  }
  return x509Block, nil
}

func veriyJWTwithEC(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
    ctx := r.Context() 
    tokenString := ctx.Value("token").(string)
    privateKey := ctx.Value("private-key").(string)
    ecdsaPrivate, decodeErr := decodeEC(privateKey)
    if decodeErr != nil {
      http.Error(w, "Failed to decode private key into *ecsa.PrivateKey", http.StatusInternalServerError)
      return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
      _, ok := token.Method.(*jwt.SigningMethodECDSA)
      if !ok {
        return nil, errors.New("Invalid type of signing method")
      }
      return &ecdsaPrivate.PublicKey, nil
    })

    if err != nil {
      http.Error(w, "Failed to decode parsing private key into jwt", http.StatusInternalServerError)
      return
    }

    if token.Valid {
      ctx := r.Context()
      ctx = context.WithValue(ctx, "token", token)
      r := r.WithContext(ctx)
      next.ServeHTTP(w, r)
    } else {
      http.Error(w, "Failed to decode parsing private key into jwt", http.StatusInternalServerError)
      return
    }
  }) 
} 

func NewJWTMiddlewarewithECDSA(privateKey string) *UseJWTMiddlewarewithECDSA {
  return &UseJWTMiddlewarewithECDSA{
    middleware: &UseMiddleware{
      function: veriyJWTwithEC,
    },
    privateKey: privateKey,
  }
}

func (ec *UseJWTMiddlewarewithECDSA) AddHandler(handler http.Handler) {
  ec.middleware.handler = handler
}

func (ec *UseJWTMiddlewarewithECDSA) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if ec.middleware.handler != nil {
    hdl := ec.middleware.function(ec.middleware.handler)
    hdl.ServeHTTP(w, r)
  }
  http.Error(w, "Handler not found", http.StatusInternalServerError)
  return
}


