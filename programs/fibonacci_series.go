package main

import "fmt"

func main(){

    var n,n1,n2,nextTerm int
	
	n1 = 0
	n2 = 1
    
    fmt.Print("Enter the number of terms : ")
    fmt.Scan(&n)
	
    fmt.Print("Fibonacci Series :")
	
    for i:=1;i<=n;i++ {
        if(i==1){
            fmt.Print(" ",n1)
            continue
        }
        if(i==2){
            fmt.Print(" ",n2)
            continue
        }
        nextTerm = n1 + n2
        n1 = n2
        n2=nextTerm
        fmt.Print(" ",nextTerm)
    }
}