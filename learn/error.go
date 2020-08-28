// 关于 Golang 错误处理的一些思考​ https://mp.weixin.qq.com/s/KPrzPP797efFUKOTTfY1Ow
package learn

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

func getSql() error {
	return errors.Wrap(sql.ErrNoRows, "GetSql failed")
}

func call() error {
	return errors.WithMessage(getSql(), "bar failed")
}

func withStack() {
	err := call()
	if errors.Cause(err) == sql.ErrNoRows {
		log.Printf("data not found, %v\n", err)
		log.Printf("%+v\n", err)
		return
	}
	if err != nil {
		log.Printf("got err,  %+v\n", err)
	}
}
