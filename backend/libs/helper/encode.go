package helper

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/twmb/murmur3"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"hash"
	"io"
	"io/ioutil"
)


func Mmh3Hash32(raw []byte) string {
	var h32 hash.Hash32 = murmur3.New32()
	h32.Write(raw)
	//if IsUint32 {
	//	return fmt.Sprintf("%d", h32.Sum32())
	//}
	return fmt.Sprintf("%d", int32(h32.Sum32()))
}


func StandBase64(braw []byte) []byte {
	bckd := base64.StdEncoding.EncodeToString(braw)
	var buffer bytes.Buffer
	for i := 0; i < len(bckd); i++ {
		ch := bckd[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}
	buffer.WriteByte('\n')
	return buffer.Bytes()
}



func UTF8ReDecode(text string, encodeStr string) (string, error) {
	var reader io.Reader
	switch encodeStr {
	case "GB18030":
		reader = transform.NewReader(bytes.NewReader([]byte(text)), simplifiedchinese.GB18030.NewDecoder())
	case "GBK":
		reader = transform.NewReader(bytes.NewReader([]byte(text)), simplifiedchinese.GBK.NewDecoder())
	case "HZGB2312":
		reader = transform.NewReader(bytes.NewReader([]byte(text)), simplifiedchinese.HZGB2312.NewDecoder())
	default:
		reader = transform.NewReader(bytes.NewReader([]byte(text)), simplifiedchinese.GB18030.NewDecoder())
	}
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func Md5(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}