package cipher

var _key = []byte("uZw4z6B9EhGdKQnjmVsYv2x5")

func GenerateKey(key string) {
	if key != "" {
		_key = []byte(key)
	}
}

func XOR(src []byte) []byte {
	_klen := len(_key)
	for i := 0; i < len(src); i++ {
		src[i] ^= _key[i%_klen]
	}
	return src
}
