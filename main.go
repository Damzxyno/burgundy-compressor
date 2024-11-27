package main

import (
	"fmt"
	"github.com/Damzxyno/burgundy-compressor/compressors"
	"os"
)

func main() {
	h := int8(-1)
	fmt.Println(h)
	niceCatInBytes, err := os.ReadFile("./landscape.bmp")
	if err != nil {
		fmt.Println(err)
	}

	rle := compressors.NewRunLengthEncoding
	zrle := compressors.NewZynoRunLengthEncoding
	rleEncodedByte := rle.Encode(niceCatInBytes)
	zrleEncodedByte := zrle.Encode(niceCatInBytes)
	fmt.Printf("Original Size => %d, RLE Size => %d, ZRLE Size => %d\n", len(niceCatInBytes), len(rleEncodedByte), len(zrleEncodedByte))
	//decodedByte := rle.Decode(rleEncodedByte)
	fmt.Println(zrleEncodedByte[:100])
	fmt.Println(niceCatInBytes[:100])
	//file, err := os.Create("./landscape1.bmp")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file.Close()
	//file.Write(zrleEncodedByte)
}
