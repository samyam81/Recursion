package main

import (
	"strconv"
)

func diffWaysToCompute(expression string) []int {
    var answer[] int
	var pureNum bool=true
	for i := 0; i < len(expression); i++ {
		if expression[i]<'0'|| expression[i]>'9'{
			pureNum=false
			L:=diffWaysToCompute(expression[:i])
			R := diffWaysToCompute(expression[i+1:])

			for _, l := range L {
				for _, r := range R {
					if expression[i]=='+'{
						answer=append(answer, l+r)
					} else if expression[i]=='-' {
						answer = append(answer, l-r)						
					} else if expression[i]=='*' {
						answer = append(answer, l*r)
					}
				}
			}
		}
	}

	if pureNum {
		number, _ := strconv.Atoi(expression)
		answer = append(answer, number)	
	}
	return answer
}