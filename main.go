package main

import (
	"fmt"
	"github.com/Damzxyno/burgundy-compressor/compressors"
)

func main() {
	//niceCatInBytes, err := os.ReadFile("./landscape.bmp")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//rle := compressors.NewRunLengthEncoding
	//zrle := compressors.NewZynoRunLengthEncoding
	//rleEncodedByte := rle.Encode(niceCatInBytes)
	//zrleEncodedByte := zrle.Encode(niceCatInBytes)
	//fmt.Printf("Original Size => %d, RLE Size => %d, ZRLE Size => %d\n", len(niceCatInBytes), len(rleEncodedByte), len(zrleEncodedByte))
	////decodedByte := rle.Decode(rleEncodedByte)
	//fmt.Println(zrleEncodedByte[:100])
	//fmt.Println(niceCatInBytes[:100])
	////file, err := os.Create("./landscape1.bmp")
	////if err != nil {
	////	fmt.Println(err)
	////}
	////defer file.Close()
	////file.Write(zrleEncodedByte)
	// Input data
	data := "huffman encoding example"

	// Step 1: Calculate frequencies
	freqMap := make(map[rune]int)
	for _, char := range data {
		freqMap[char]++
	}
	for k, v := range freqMap {
		fmt.Printf("%c, %v\n", k, v)
	}
	// Step 2: Build the Huffman Tree
	root := compressors.BuildHuffmanTree(freqMap)

	// Step 3: Generate Huffman Codes
	codes := make(map[rune]string)
	compressors.GenerateCodes(root, "", codes)

	// Step 4: Encode the input data
	encoded := compressors.HuffmanEncode(data, codes)
	fmt.Println("Huffman Codes:", codes)
	fmt.Println("Encoded Data:", encoded)

	// Step 5: Decode the encoded data
	decoded := compressors.HuffmanDecode(encoded, root)
	fmt.Println("Decoded Data:", decoded)
}
