package decoding

import (
	"encoding/base64"
	"os"
	"strings"
)


func Base64ToFile(b64 string, f *os.File) error {
	var dec []byte
	var err error

	b64data := b64[strings.IndexByte(b64, ',')+1:]
	// 解码
	dec, err = base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return err
	}
	// 写入流文件
    if _, err := f.Write(dec); err != nil {
        return err
    }

	// 刷新同步到流文件
    if err := f.Sync(); err != nil {
        return err
    }

	// 重置文件读写位置
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	
	return nil
}