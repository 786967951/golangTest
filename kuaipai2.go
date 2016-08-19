package main
import (
"fmt"
"time"
)
var x,y,z,i,j int

func main() {
     values := make([]int,10000001,10000001)
     x := 0
     for x<10000001{
     	 values[x] = 10000000-x
     	 x=x+1
     }
     fmt.Println(time.Now()) 
     qSort(values,0,10000000) 
     fmt.Println(time.Now())
     fmt.Print(values[999]) 
     //fmt.Print(values)    
}

func qSort(a []int,p int,q int) {
	   if p<q {
	    mainSort(a, p, q)
	    qSort(a,p,y-1)
	    qSort(a,y+1,q)
	   }
}

func mainSort(a []int,p int,q int){
	   i=p
	   j=p+1
	   z=(q+p)/2
	   if p != z{
	      a[p],a[z]=a[z],a[p]
	   }
	   x=a[p]
	   for  j<=q {
	   	 if a[j] < x{
	   	 	  i=i+1
	   	 	  a[i],a[j] = a[j],a[i]
	   	 	  j=j+1
	     } else {
	     	  j=j+1
	     }
	   }
	   a[p], a[i] = a[i], a[p] 
	   y=i  
}