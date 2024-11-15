/**
 * @Author: dingQingHui
 * @Description:
 * @File: pool
 * @Version: 1.0.0
 * @Date: 2024/8/15 15:04
 */

package workers

import (
	"github.com/panjf2000/ants/v2"
)

var pool *ants.Pool

func init() {
	_pool, err := ants.NewPool(1000)
	if err != nil {
		panic("ants.NewPool, error")
	}
	pool = _pool
}

func Submit(fn func(), recoverFun func(err interface{})) {
	err := pool.Submit(func() {
		Count.Add(1)
		Try(fn, recoverFun)
		Count.Add(-1)
	})
	if err != nil {
		return
	}
}

func Try(fn func(), reFun func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			PanicCount.Add(1)
			if reFun != nil {
				reFun(err)
			}
		}
	}()
	fn()
}
