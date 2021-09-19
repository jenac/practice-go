module ch10/gorm-example

go 1.17

require ch10/gorm_example/model v0.0.0

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
)

replace ch10/gorm_example/model v0.0.0 => ./model
