module 01-init

go 1.14

require (
	github.com/gin-contrib/zap v0.0.1 // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	handle v0.0.0
	middleware v0.0.0
	sql v0.0.0
	tool v0.0.0
)

replace middleware v0.0.0 => ./middleware

replace tool v0.0.0 => ./tool

replace sql v0.0.0 => ./sql

replace handle v0.0.0 => ./handle
