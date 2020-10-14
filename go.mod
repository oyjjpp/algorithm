module github.com/oyjjpp/algorithm

go 1.13

require (
	github.com/forgoer/openssl v0.0.0-20200331032942-ad9f8d57d8b1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/snappy v0.0.1
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/streadway/amqp v1.0.0
	github.com/yuin/goldmark v1.1.32 // indirect
	shield.zhangyue.com/golib/logs v0.0.5
	shield.zhangyue.com/golib/utils v0.0.6
)

replace (
	shield.zhangyue.com/golib/logs => ../logs
	shield.zhangyue.com/golib/utils => ../utils
)
