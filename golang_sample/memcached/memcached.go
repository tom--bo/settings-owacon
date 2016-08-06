package main

// memcached OFFICIAL DOCUMENTS
// https://godoc.org/github.com/bradfitz/gomemcache/memcache

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
    // without UNIX socket
	// mc := memcache.New("localhost:11211")
	mc := memcache.New("/tmp/memcached.sock")

	// 下でも良い
	//mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	item1 := &memcache.Item{
		Key:   "foo",
		Value: []byte("bar"),
	}
	mc.Set(item1)

	item2 := &memcache.Item{
		Key:   "hoge",
		Value: []byte("fuga"),
	}
	mc.Set(item2)

	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println("file create err:", err)
		return
	}

	fmt.Printf("%s\n", it.Key)
	fmt.Printf("%s\n", it.Value)

	// 複数同時取得
	ret, err2 := mc.GetMulti([]string{"foo", "hoge"})
	if err2 != nil {
        fmt.Println(err)
		return
	}

	for k, v := range ret {
		fmt.Printf("ret: %s => %s\n", k, v.Value)
	}
}
