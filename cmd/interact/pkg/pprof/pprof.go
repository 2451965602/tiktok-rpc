package pprof

import (
	"log"
	"net/http"
	"os"
	"runtime"

	_ "net/http/pprof"
)

func Load() {
	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(`:9003`, nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
}
