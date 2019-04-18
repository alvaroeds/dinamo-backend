package commons

import (
        "crypto/rsa"
        "github.com/dgrijalva/jwt-go"
        "io/ioutil"
        "log"
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

        publicBytes, err := ioutil.ReadFile("./keys/private.rsa")
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

func GenerateJWT(){

}
