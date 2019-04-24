package commons

import (
        "crypto/rsa"
        "github.com/alvaroenriqueds/dinamo-backend/models"
        "github.com/dgrijalva/jwt-go"
        "io/ioutil"
        "log"
        "time"
)

var (
        privateKey *rsa.PrivateKey
        publicKey *rsa.PublicKey
)

func init()  {
        privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
        if err != nil {
                log.Fatal("No se pudo leer el archivo privado")
        }

        publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
        if err != nil {
                log.Fatal("No se pudo leer el archivo publico")
        }

        privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
        if err != nil {
                log.Fatal("No se pudo hacer el parse a privateKey")
        }

        publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
        if err != nil {
                log.Fatal("No se pudo hacer el parse a publicKey")
        }
}

func GenerateJWT(user models.User) string{
        claims := models.Claim{
                User : user,
                StandardClaims: jwt.StandardClaims{
                        ExpiresAt: time.Now().Add(time.Hour*1).Unix(),
                        //Objetivo
                        Issuer: "Generar JWT",
                },
        }

        //convirtiendo el claim a token
        token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
        //convirtiendo el token a string
        result, err := token.SignedString(privateKey)
        if err != nil {
                log.Fatal("No se pudo firmar el token")
        }

        //retornamos el token en un string
        return result
}
