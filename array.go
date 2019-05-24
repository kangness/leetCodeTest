/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
    "fmt"
)

func main()  {
    primes := [8]int{2,3,5,7,9,11,13,15,} //定义一个数组
    fmt.Printf("`primes`数组的值：%d\n",primes)
    var  sum []int = primes[1:4]   //定义一个切片
    fmt.Printf("`sum`切片的值:%d\n",sum)
    fmt.Printf("`sum[0]`所对应的内存地址是：%x\n",&sum[0])
    fmt.Printf("`primes[1]`所对应的内存地址是：%x\n",&primes[1])
    sum[0] = 5
    fmt.Printf("`sum`切片的值:%d\n",sum)
    fmt.Printf("`primes`切片的值:%d\n",primes)
    fmt.Printf("`sum[0]`所对应的内存地址是：%x\n",&sum[0])
    fmt.Printf("`primes[1]`所对应的内存地址是：%x\n",&primes[1])

    var  s1 []int
    s1 = sum
    fmt.Printf("`s1`切片对应的值为：%d\n",s1)
    fmt.Printf("s1[0] == sum[0]为：%v\n",&s1[0] == &sum[0])
}
