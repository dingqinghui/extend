/**
 * @Author: dingQingHui
 * @Description:
 * @File: api
 * @Version: 1.0.0
 * @Date: 2024/11/19 18:10
 */

package codec

type ICodec interface {
	Decode(data []byte, msg interface{}) error
	Encode(msg interface{}) ([]byte, error)
}
