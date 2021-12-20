package caesar

func Encrypt(input string, key int) string {
	//만약 키가 음수 값이면 키값에 mod26을 하여 나온 값에 26을 더해줍니다.
	key8 := byte(key%26+26) % 26
	var outBuffer []byte
	for _, r := range input {
		newByte := byte(r)
		if 'A' <= r && r <= 'Z' {
			outBuffer = append(outBuffer, 'A'+(newByte-'A'+key8)%26)
		} else if 'a' <= r && r <= 'z' {
			outBuffer = append(outBuffer, 'a'+(newByte-'a'+key8)%26)
		} else {
			outBuffer = append(outBuffer, newByte)
		}
	}
	return string(outBuffer)
}

// Decrypt decrypts by left shift of "key" each character of "input"
func Decrypt(input string, key int) string {
	// left shift of "key" is same as right shift of 26-"key"
	return Encrypt(input, 26-key)
}

func MyEncrypt(input string, key int) string {
	var outBuffer []byte
	for _, r := range input {
		outBuffer = append(outBuffer, byte(int(r)+key))
	}
	return string(outBuffer)
}

func MyDecrypt(input string, key int) string {
	var outBuffer []byte
	key = key % 26
	for _, r := range input {
		outBuffer = append(outBuffer, byte(int(r)-key))
	}
	return string(outBuffer)
}
