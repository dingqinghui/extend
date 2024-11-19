/**
 * @Author: dingQingHui
 * @Description:
 * @File: api
 * @Version: 1.0.0
 * @Date: 2024/11/19 18:10
 */

package codec

var (
	Json    = ICodec(new(jsonCodec))
	MsgPack = ICodec(new(msgPackCodec))
	PB      = ICodec(new(pbCodec))
)

type ICodec interface {
	Decode(data []byte, msg interface{}) error
	Encode(msg interface{}) ([]byte, error)
}
