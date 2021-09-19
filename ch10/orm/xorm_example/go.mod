module ch10/xorm_example

go 1.17

require ch10/xorm_example/model v0.0.0

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/go-xorm/xorm v0.7.9 // indirect
	xorm.io/builder v0.3.6 // indirect
	xorm.io/core v0.7.3 // indirect
)

replace ch10/xorm_example/model v0.0.0 => ./model
