package images

import (
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/credentials"
        "github.com/aws/aws-sdk-go/aws/session"
)

func InitSession() (*session.Session, error){
        //bucket := "prueba-dinamo01"

        sess, err := session.NewSession(&aws.Config{
                Region: aws.String("us-east-2"),
                Credentials: credentials.NewStaticCredentials("AKIA36DW2ILGUDLYT2XS", "nrz/kura0znp4Wm8f6ro6AvvEbVjg2UdEF/TVe5K", ""),
        })
        if err != nil {
                return nil, err
        }

        return sess, nil
}

