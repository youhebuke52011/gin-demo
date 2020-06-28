package compression

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func KGDecodeAndGz(rData string, res *map[string]interface{}) error {
	var (
		decodeData []byte
		dataBytes  []byte
		err        error
		readerData *gzip.Reader
	)
	fmt.Println(1)
	if decodeData, err = base64.StdEncoding.DecodeString(rData); err != nil {
		return err
	}
	fmt.Println(2)
	if readerData, err = gzip.NewReader(bytes.NewReader(decodeData)); err != nil {
		return err
	}
	fmt.Println(3)
	defer readerData.Close()
	if dataBytes, err = ioutil.ReadAll(readerData); err != nil {
		return err
	}
	fmt.Println(4)
	if err = json.Unmarshal(dataBytes, res); err != nil {
		return err
	}
	fmt.Println(res)
	fmt.Println(123)
	return nil
}
