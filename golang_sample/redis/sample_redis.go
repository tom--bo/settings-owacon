package main

import "fmt"

// リファレンス
// http://godoc.org/github.com/garyburd/redigo/redis
import "github.com/garyburd/redigo/redis"

func main() {
//INIT OMIT
c, err := redis.Dial("tcp", ":6379")
if err != nil {
panic(err)
}
defer c.Close()

//set
c.Do("SET", "message1", "Hello World")

//pipe
c.Send("SET", "foo", "bar")
c.Send("GET", "foo")
c.Flush()
c.Receive() // reply from SET
v, err = c.Receive() // reply from GET


//get
world, err := redis.String(c.Do("GET", "message1"))
if err != nil {
fmt.Println("key not found")
}

fmt.Println(world)
//ENDINIT OMIT
}

