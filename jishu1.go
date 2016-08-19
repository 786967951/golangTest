package main
import (
"fmt"
"time"
)
var x int
func main() {
     values := make([]int,10000001,10000001)
     count := make([]int,10000001,10000001)
     temp := make([]int,10000001,10000001)
     x = 0
     for x<10000001{
     	 values[x] = 10000000-x
     	 x=x+1
     }
     fmt.Println(time.Now()) 
     lineSort(values,10000000,count,10000000,temp) 
     fmt.Println(time.Now())
     fmt.Print(temp[999])   
}

func lineSort(a []int,p int,b []int,q int,c []int) {
     x=0
	   for x<=p{
	    b[a[x]]+=1
	    x=x+1
	   }
     x=1
	   for x<=q{
	    b[x]=b[x]+b[x-1]
	    x=x+1
	   } 
	 x=p
	   for x>=0 {
        c[b[a[x]]-1] =a[x]
        if a[x]>0  && b[a[x]]- b[a[x]-1] > 1 {
      	   b[a[x]]=b[a[x]]-1
        }else if a[x]==0 && b[a[x]] > 1 {
           b[a[x]]=b[a[x]]-1	
        }
        x=x-1 
     } 
}