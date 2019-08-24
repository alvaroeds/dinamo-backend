package commons

import (
        "crypto/sha256"
        "fmt"
)

func Coding (password string) string {

        pwd := sha256.Sum256([]byte(password))
        pass := fmt.Sprintf("%x", pwd)
        password = pass

        return pass
}
