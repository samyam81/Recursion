package main

const MOD = 1000000007
func countGoodNumbers(n int64) int {
    var ed int64
	var od int64
	if n&1==1{
		od=n/2
		ed=(n+1)/2
	} else {
		od=n/2
		ed=n/2
	}

	return int(((power(5,ed)%MOD*power(4,od)%MOD))%MOD)
}

func power(a,b int64) int64{
	if b==0 {
		return 1
	} 
	halfPower:=power(a,b/2)
	if b%2==0 {
		return (halfPower*halfPower)%MOD
	} else{
		return (halfPower*halfPower)%MOD *(a%MOD) %MOD
	}
}