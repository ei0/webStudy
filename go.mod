module 01-init

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	
	middleware v0.0.0
	tool v0.0.0
	sql v0.0.0
	handle v0.0.0
)

replace middleware v0.0.0 => ./middleware
replace tool v0.0.0 => ./tool
replace	sql v0.0.0 => ./sql
replace handle v0.0.0 => ./handle
