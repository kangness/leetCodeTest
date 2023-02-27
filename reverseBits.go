package main

func reverseBits(num uint32) uint32 {
	var  ans uint32 =0;
	//进制的本质
	var i int =32;
	for {
		if i <= 0 {
			break
		}
		ans <<=1
		ans = ans + num&1
		num>>=1
		i = i -1
	}
	return ans;
}
