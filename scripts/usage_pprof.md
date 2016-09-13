# how to use pprof

## add scripts
`
	import _ "net/http/pprof"

	go func() {
 		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
`

## run pprof
`go tool pprof app --second=90`
