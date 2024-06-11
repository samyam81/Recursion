package main

func lastRemaining(n int) int {
    if n==1 {
		return 1
	} 
	var ord int=1
	var answer int=1
	remain:=n
	
	var left bool=false

	for remain>1{
		if left==false || remain%2==1 {
			answer=answer+ord
		}
		ord=ord*2
		remain=remain/2
		left=!left
	}
	return answer
}

// func lastRemaining(n int) int {
//     if n == 1 {
//         return 1
//     }
//     return remainingRecursive(n, 1, 1, false)
// }

// func remainingRecursive(n, ord, answer int, left bool) int {
//     if n == 1 {
//         return answer
//     }
//     if left == false || n%2 == 1 {
//         answer = answer + ord
//     }
//     return remainingRecursive(n/2, ord*2, answer, !left)
// }
