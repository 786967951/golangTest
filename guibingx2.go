package main
import (
"fmt"
"time"
"runtime"
)
var x,y int
var a1,a2,a3,a4 int
var c1= make(chan bool, 4)
func main() {
     runtime.GOMAXPROCS(runtime.NumCPU())
	 fmt.Println(runtime.NumCPU())
     values := make([]int,10000001,10000001)
	 temp := make([]int,10000001,10000001)
	 temp1 := make([]int,2500001,2500001)
	 temp2 := make([]int,2500001,2500001)
	 temp3 := make([]int,2500001,2500001)
	 temp4 := make([]int,2500001,2500001)
     x = 0
     for x<10000001{
     	 values[x] = 10000001 - x
     	 x=x+1
     }
     fmt.Println(time.Now())
	 a1=10000000/4
	 a2=10000000/2
	 a3=3*10000000/4
	 a4=10000000 
     go qSort(values,0,a1,temp1)
	 go qSort(values,a1+1,a2,temp2)
	 go qSort(values,a2+1,a3,temp3)  
	 go qSort(values,a3+1,a4,temp4)//if only using temp,there will be problem when concurrenting
	 y=0
	 for y<4{
	     <-c1
		 y=y+1
	 }
	 close(c1)	 
	 mergeArr(values,0,a1,a2,temp)
	 mergeArr(values,a2+1,a3,a4,temp)
	 mergeArr(values,0,a2,a4,temp)
     fmt.Println(time.Now())
	 fmt.Println(values[999])
}

func qSort(values []int,first int,last int,temp []int ) {
       var mid int 
	   if first < last {
	   mid=(first + last)/2
	   qSort(values, first, mid, temp)
	   qSort(values, mid+1, last, temp)
	   mergeArr(values,first,mid,last,temp)
	 }
}

func mergeArr(a []int,first int,mid int,last int,temp []int){
	   i:=first
	   j:=mid+1
	   m:=mid
	   n:=last
	   k:=0
	   for i<=m && j<=n {
	   	 if a[i] < a[j]{
	   	   temp[k] = a[i]
	   	   k=k+1
	   	   i=i+1
	   	} else {
	   	   temp[k] = a[j]
	   	   k=k+1
	   	   j=j+1
	   	 }
	   }
	   for i>m && j<=n {
	     temp[k] = a[j]
	     k=k+1
	   	 j=j+1
	   }
	   for j>n && i<=m {
	   	 temp[k] = a[i]
	   	 k=k+1
	   	 i=i+1
	   }
	    i=0
	    j=first
	    for i<k {
	     	a[j] = temp[i]
	    	i=i+1
			j=j+1
	    }
        if 	first==	0 && last==a1 ||first==	a1+1 && last==a2 || first==a2+1 && last==a3||first==a3+1 && last==a4 {
            c1<-true			
		}		
}