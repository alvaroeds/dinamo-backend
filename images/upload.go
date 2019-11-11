package images

import (
        "encoding/json"
        "fmt"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/service/s3/s3manager"
        "github.com/labstack/echo"
        "golang.org/x/net/websocket"
        "log"
)

func imageS3_Upload(c echo.Context) error {

        sess, err := InitSession()
        if err != nil {
                return c.String(500, "NO SE PUDO CONECATAR, LO SIENTO")
        }

        file, handle, err := c.Request().FormFile("myFile")
        if err != nil {
                return c.String(500, "222222222222222")
        }
        defer file.Close()
        /*
        a, err :=handle.Open()
        if err != nil {
                c.String(500, "2.55555555555")
        }
        */
        /*
        file, err := os.Open("./files/hola.jpg")
        if err != nil {
                return c.String(500, "os error")
        }
         */

        /*
        data, err := io.Reader(file)
        if err != nil {
                c.String(500, "3333333333333333")
        }
        */

        //UPLOAD
        uploader := s3manager.NewUploader(sess)
        bucket := "prueba-dinamo01"

        output, err := uploader.Upload(&s3manager.UploadInput{
                ACL:                       aws.String("public-read"),
                Body:                      file,
                Bucket:                    aws.String(bucket),
                CacheControl:              nil,
                ContentDisposition:        nil,
                ContentEncoding:           nil,
                ContentLanguage:           nil,
                ContentMD5:                nil,
                ContentType:               nil,
                Expires:                   nil,
                GrantFullControl:          nil,
                GrantRead:                 nil,
                GrantReadACP:              nil,
                GrantWriteACP:             nil,
                Key:                       aws.String("alvaro/_"+handle.Filename),
                Metadata:                  nil,
                ObjectLockLegalHoldStatus: nil,
                ObjectLockMode:            nil,
                ObjectLockRetainUntilDate: nil,
                RequestPayer:              nil,
                SSECustomerAlgorithm:      nil,
                SSECustomerKey:            nil,
                SSECustomerKeyMD5:         nil,
                SSEKMSEncryptionContext:   nil,
                SSEKMSKeyId:               nil,
                ServerSideEncryption:      nil,
                StorageClass:              nil,
                Tagging:                   nil,
                WebsiteRedirectLocation:   nil,
        })
        if err != nil {
                fmt.Println(err)
                return c.String(500, "44444444444444")
        }

        //fmt.Println(output)
        port := 5051
        origin := fmt.Sprintf("http://localhost:%d/", port)
        url := fmt.Sprintf("ws://localhost:%d/ws", port)
        ws, err := websocket.Dial(url, "", origin)
        if err != nil {
                log.Fatal(err)
        }

        //nombre := "./files/" + handle.Filename
        j, err := json.Marshal(&output.Location)
        if _, err := ws.Write(j); err != nil {
                log.Fatal(err)
        }

        return c.JSON(500, echo.Map{
                "url": output.Location,
        })
}
