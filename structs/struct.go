/**
 * @Author: dingQingHui
 * @Description:
 * @File: struct
 * @Version: 1.0.0
 * @Date: 2024/11/21 15:42
 */

package structs

import "github.com/dingqinghui/extend/codec"

func DeepCopy(src, dst interface{}) error {
	buf, err := codec.Json.Encode(src)
	if err != nil {
		return err
	}
	if err := codec.Json.Decode(buf, dst); err != nil {
		return err
	}
	return nil
}
