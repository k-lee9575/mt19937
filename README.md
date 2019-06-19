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
    
    distBernolli := mt19937.DistBernolli(mt, 0.25)
    for i := 0; i < 10; i++{
        fmt.Println(distBernolli.Bool())
    }

    dist01 := mt19937.Dist01(mt)
    for i := 0; i < 10; i++ {
        fmt.Println(dist01.Float64())
    }

    weightList := []int {1,2,3,4}
    distDiscrete := mt19937.DistDiscrete(mt,weightList)

    for i := 0; i < 50; i ++ {
        fmt.Println(distDiscrete.Discrete())
    }
}
```
