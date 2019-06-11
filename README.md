# mt19937
Golang


# Sample
```go
package main
import(
    "fmt"
    "github.com/k-lee9575/mt19937"
  	//"time"
)
  
func main(){
    mt := mt19937.New()
    //mt.Seed(uint64(time.Now().Unix()))
    fmt.Println(mt.Random())
    
    distInt64 := mt19937.DistInt64(mt, -5, 5)
    for i:=0; i< 10; i ++{
        fmt.Println(distInt64.Int64())
    }
}
```
