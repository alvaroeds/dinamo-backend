package images

import "github.com/labstack/echo"

func Images_controller(e *echo.Echo){
        e.POST("/upload_s3", imageS3_Upload)
}
