/**
 * @Author: dingQingHui
 * @Description:
 * @File: statistics
 * @Version: 1.0.0
 * @Date: 2024/8/15 15:10
 */

package workers

import (
	"sync/atomic"
)

var Count atomic.Int64
var PanicCount atomic.Uint64
