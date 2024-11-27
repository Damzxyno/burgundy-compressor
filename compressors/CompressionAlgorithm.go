package compressors

type CompressionAlgorithm interface {
	Encode([]byte) []byte
	Decode([]byte) []byte
}

var (
	NewRunLengthEncoding     = RunLengthEncoding{}
	NewZynoRunLengthEncoding = ZynoRunLengthEncoding{}
)
