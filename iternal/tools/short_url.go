package tools

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RundUrl() uint64 {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(9999999998))

	if err != nil {
		fmt.Println("Ошибка при генерации случайного числа:", err)
		return 0
	}

	randUrl := randomNumber.Uint64()

	return randUrl
}

func Base62Encode(number uint64) string {
	length := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}
