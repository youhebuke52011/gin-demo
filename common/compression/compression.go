package compression

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

func KGDecodeAndGz(rData string, res *map[string]interface{}) error {
	var (
		decodeData []byte
		dataBytes  []byte
		err        error
		readerData *gzip.Reader
	)
	if decodeData, err = base64.StdEncoding.DecodeString(rData); err != nil {
		return err
	}
	if readerData, err = gzip.NewReader(bytes.NewReader(decodeData)); err != nil {
		return err
	}
	defer readerData.Close()
	if dataBytes, err = ioutil.ReadAll(readerData); err != nil {
		return err
	}
	if err = json.Unmarshal(dataBytes, res); err != nil {
		return err
	}
	//fmt.Println(res)
	return nil
}
