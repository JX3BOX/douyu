# 斗鱼SDK


## Getting Start

```golang
package main
import "github.com/JX3BOX/douyu"
func main(){
    dy, err := douyu.New("AID", "Key")
    if err != nil {
        log.Fatalln(err)
        Fatalln
    }
    list, err:= dy.BatchGetRoomInfo(BatchGetRoomInfoParams{RIds: []int{8852876, 8889134}})
    if err == nil {
        log.Println(list)
    }
}
```