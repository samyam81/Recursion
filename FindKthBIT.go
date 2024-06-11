package main

func findKthBit(n int, k int) byte {
    s := kobbie(n)
    return s[k-1]
}

func kobbie(n int) string {
    if n == 1 {
        return "0"
    }

    s1 := kobbie(n-1)
    s2 := "1"
    s3 := revInverse(s1)
    return s1 + s2 + s3
}

func revInverse(s1 string) string {
    bytes := []byte(s1)
    for i := 0; i < len(bytes); i++ {
        if bytes[i] == '1' {
            bytes[i] = '0'
        } else {
            bytes[i] = '1'
        }
    }
    return reverseString(string(bytes))
}

func reverseString(str string) string {
    runes := []rune(str)
    left := 0
    right := len(runes) - 1
    for left < right {
        temp := runes[left]
        runes[left] = runes[right]
        runes[right] = temp
        left++
        right--
    }
    return string(runes)
}
