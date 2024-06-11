package main

func myPow(x float64, n int) float64 {
	if n==0 {
		return 1
	}
	if n<0 {
		return 1/myPow(x,-n)
	}

	halfPower:=myPow(x,n/2)
	if n%2==0 {
		return halfPower*halfPower
	} else {
		return halfPower*halfPower*x
	}
}
