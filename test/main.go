package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	//t := []map[string]string{}
	//t = append(t, map[string]string{"a": "1"})
	//t = append(t, map[string]string{"b": "2"})
	//t = append(t, map[string]string{"c": "3"})
	//var wg sync.WaitGroup
	//flag := true
	//for _, tmp := range t {
	//	//data := tmp
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		if flag {
	//			tmp["a"] = "11"
	//			tmp["b"] = "22"
	//			tmp["c"] = "33"
	//			flag = false
	//		}
	//		//data["a"] = "11"
	//		//data["b"] = "22"
	//		//data["c"] = "33"
	//	}()
	//}
	//wg.Wait()
	//for _, tmp := range t {
	//	fmt.Println(tmp)
	//}

	//ib := IntToBytes(71)
	//fmt.Println(string(ib))
	//sbb := []byte("K")
	//sbb[0] = uint8(71)
	//sb := []byte("GN6q4UtZg35ztOms")
	//fmt.Println(sb)

	sbb := []byte("GN6q4UtZg35ztOms")
	fmt.Println(sbb[0] == 'G')
	sb := decode("GtGotGdainwCleWc")
	fmt.Println(string(sb))

	sbe := encodeScore("33.2_997.0_C")
	fmt.Println(sbe)
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func decode(encodeStr string) []byte {
	if encodeStr == "" {
		return []byte("")
	}
	decodeTable := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		62,
		0, 0, 0,
		63,
		52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
		0, 0, 0, 0, 0, 0, 0,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
		0, 0, 0, 0, 0, 0,
		26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
		39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	var (
		res       []byte
		sb        []byte
		nValue    int
		temp      int
		m_randNum int
		idx       int
		k         int
		tmp       uint8
	)
	sb = []byte(encodeStr)
	res = []byte(encodeStr)
	for i := 0; i < len(sb); i++ {
		/**
		int d = 0;
		temp = i/4;
		m_randNum = new Double(Math.pow(new Double((temp+353)%1024), 3)).intValue();
		m_randNum = m_randNum%1909553;

		nValue = ((((DecodeTable[(int)encodeCharArr[idx++]])+64-m_randNum)%64)+64)%64 << 18;
		nValue += ((((DecodeTable[(int)encodeCharArr[idx++]])+64-m_randNum)%64)+64)%64 << 12;
		strDecode +=(char)((nValue & 0x00FF0000) >> 16);
		char curData = encodeCharArr[idx];
		if (curData != '=')
		{
				nValue += ((((DecodeTable[(int)encodeCharArr[idx++]])+64-m_randNum)%64)+64)%64 << 6;
				strDecode +=(char)((nValue & 0x0000FF00) >> 8);
				curData = encodeCharArr[idx];
				if (curData != '=')
				{
						nValue += ((((DecodeTable[(int)encodeCharArr[idx++]])+64-m_randNum)%64)+64)%64;
						strDecode +=(char)(nValue & 0x000000FF);
				}
		}
		i += 4;
		*/
		temp = i / 4
		m_randNum = int(math.Pow(float64((temp+353)%1024), 3)) % 1909553
		nValue = ((((decodeTable[sb[idx]] + 64 - m_randNum) % 64) + 64) % 64) << 18
		idx++
		nValue += ((((decodeTable[sb[idx]] + 64 - m_randNum) % 64) + 64) % 64) << 12
		idx++
		tmp = uint8((nValue & 0x00FF0000) >> 16)
		res[k] = tmp
		k++
		if sb[idx] != '=' {
			nValue += ((((decodeTable[sb[idx]] + 64 - m_randNum) % 64) + 64) % 64) << 6
			idx++
			tmp = uint8((nValue & 0x0000FF00) >> 8)
			res[k] = tmp
			k++
			if sb[idx] != '=' {
				nValue += (((decodeTable[sb[idx]] + 64 - m_randNum) % 64) + 64) % 64
				idx++
				tmp = uint8(nValue & 0x000000FF)
				res[k] = tmp
				k++
			}
		}
		i += 4
	}
	fmt.Println(k)
	fmt.Println(len(sb))
	return res[:k]
}

func encodeScore(originStr string) string {
	/**
	final char EncodeTable[]="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/".toCharArray();

		String strEncode = "";
		char Tmp[]=new char[4];

		char[] data = originStr.toCharArray();
		int DataByte = originStr.getBytes().length;
		int idx = 0;
		int m_randNum = 0;
		for(int i=0;i<(int)(DataByte / 3);i++)
		{
			Tmp[1] = data[idx++];
			Tmp[2] = data[idx++];
			Tmp[3] = data[idx++];


			m_randNum = new Double(Math.pow(new Double((i+353)%1024), 3)).intValue();
			m_randNum = m_randNum%1909553;

			strEncode+= EncodeTable[((Tmp[1] >> 2)+m_randNum)%64];
			strEncode+= EncodeTable[((((Tmp[1] << 4) | (Tmp[2] >> 4) & 0x3F)+m_randNum)%64)];
			strEncode+= EncodeTable[((((Tmp[2] << 2) | (Tmp[3] >> 6) & 0x3F)+m_randNum)%64)];
			strEncode+= EncodeTable[(((Tmp[3] & 0x3F)+m_randNum)%64)];
		}

		int Mod=DataByte % 3;
		int temp = DataByte/3;
		m_randNum = new Double(Math.pow(new Double((temp+353)%1024), 3)).intValue();
		m_randNum = m_randNum%1909553;
		if(Mod==1)
		{
			Tmp[1] = data[idx++];
			strEncode+= EncodeTable[((((Tmp[1] & 0xFC) >> 2)+m_randNum)%64)];
			strEncode+= EncodeTable[((((Tmp[1] & 0x03) << 4)+m_randNum)%64)];
			strEncode+= "==";
		}
		else if(Mod==2)
		{
			Tmp[1] = data[idx++];
			Tmp[2] = data[idx++];
			strEncode+= EncodeTable[((((Tmp[1] & 0xFC) >> 2) +m_randNum)%64)];
			strEncode+= EncodeTable[(((((Tmp[1] & 0x03) << 4) | ((Tmp[2] & 0xF0) >> 4))+m_randNum)%64)];
			strEncode+= EncodeTable[((((Tmp[2] & 0x0F) << 2)+m_randNum)%64)];
			strEncode+= "=";
		}
		return strEncode;
	*/
	var (
		encodeTable = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
		tmp         = make([]byte, 4)
		data        = []byte(originStr)
		res         = make([]byte, 2*len(originStr))
		length      = len(data)
		idx         int
		mRandNum    int
		k           int
	)
	for i := 0; i < length/3; i++ {
		tmp[1] = data[idx]
		idx++
		tmp[2] = data[idx]
		idx++
		tmp[3] = data[idx]
		idx++

		mRandNum = int(math.Pow(float64((i+353)%1024), 3)) % 1909553
		res[k] = encodeTable[(int(tmp[1]>>2)+mRandNum)%64]
		k++
		res[k] = encodeTable[((int((tmp[1]<<4)|(tmp[2]>>4)&0x3F) + mRandNum) % 64)]
		k++
		res[k] = encodeTable[((int((tmp[2]<<2)|(tmp[3]>>6)&0x3F) + mRandNum) % 64)]
		k++
		res[k] = encodeTable[((int(tmp[3]&0x3F) + mRandNum) % 64)]
		k++
	}
	mod := length % 3
	temp := length / 3
	mRandNum = int(math.Pow(float64((temp+353)%1024), 3)) % 1909553
	if mod == 1 {
		tmp[1] = data[idx]
		idx++
		res[k] = encodeTable[((int((tmp[1] & 0xFC) >> 2)+mRandNum)%64)]
		k++
		res[k] = encodeTable[((int((tmp[1] & 0x03) << 4)+mRandNum)%64)]
		k++
		res[k] = '='
		k++
		res[k] = '='
		k++
	} else if mod == 2 {
		tmp[1] = data[idx]
		idx++
		tmp[2] = data[idx]
		idx++
		res[k] = encodeTable[((int((tmp[1] & 0xFC) >> 2) +mRandNum)%64)]
		k++
		res[k] = encodeTable[((int(((tmp[1] & 0x03) << 4) | ((tmp[2] & 0xF0) >> 4))+mRandNum)%64)]
		k++
		res[k] = encodeTable[((int((tmp[2] & 0x0F) << 2)+mRandNum)%64)]
		k++
		res[k] = '='
		k++
	}
	return string(res[:k])
}
