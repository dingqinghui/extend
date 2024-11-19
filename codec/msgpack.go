/**
 * @Author: dingQingHui
 * @Description:
 * @File: msgpack
 * @Version: 1.0.0
 * @Date: 2024/11/19 18:20
 */

package codec

import "github.com/vmihailenco/msgpack/v5"

type MsgPackCodec struct {
}

func (p *MsgPackCodec) Decode(data []byte, msg interface{}) error {
	err := msgpack.Unmarshal(data, msg)
	return err
}

func (p *MsgPackCodec) Encode(msg interface{}) ([]byte, error) {
	data, err := msgpack.Marshal(msg)
	return data, err
}
