package main

import (
    "fmt"
)

// Condicional
func conditionalExample(x int) string {
    if x > 0 {
        return "Positive"
    } else if x < 0 {
        return "Negative"
    }
    return "Zero"
}

// Loop
func loopExample(n int) int {
    sum := 0
    for i := 0; i < n; i++ {
        sum += i
    }
    return sum
}

// Função chamada
func functionExample(a, b int) int {
    return a + b
}

// XOR Encryption
func xorEncryptDecrypt(input, key string) string {
    output := make([]byte, len(input))
    keyLen := len(key)
    for i := range input {
        output[i] = input[i] ^ key[i%keyLen]
    }
    return string(output)
}

func main() {
    fmt.Println("Condicional:", conditionalExample(5))
    fmt.Println("Loop:", loopExample(5))
    fmt.Println("Chamada de função:", functionExample(3, 4))
    encrypted := xorEncryptDecrypt("hello", "key")
    fmt.Println("XOR Encrypt:", encrypted)
    fmt.Println("XOR Decrypt:", xorEncryptDecrypt(encrypted, "key"))
}
