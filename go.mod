module github.com/oyjjpp/algorithm

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/streadway/amqp v1.0.0
)

replace (
	shield.zhangyue.com/golib/logs => ../logs
	shield.zhangyue.com/golib/utils => ../utils
)
