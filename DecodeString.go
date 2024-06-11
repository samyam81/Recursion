package main

func decodeString(s string) string {
    var i int
    return helper(&i, s)
}

func helper(idx *int, str string) string {
    var n int
    var answer string

    for *idx < len(str) {
        current := str[*idx]
        *idx++

        if current >= '0' && current <= '9' {
            n = n*10 + int(current) - '0'
        } else if current == '[' {
            temp := helper(idx, str)

            for n > 0 {
                answer += temp
                n--
            }
        } else if current == ']' {
            return answer
        } else {
            answer += string(current)
        }
    }
    return answer
}