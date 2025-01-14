package runner

import (
	"hash/crc32"
	"log"
	"math/rand"
	"os"
	"pkg/file"
	"runtime"
	"time"
)

var (
	Hostname string
	Cwd      string
)

func Init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var err error
	Hostname, err = os.Hostname()
	if err != nil {
		log.Fatalln("[F] cannot get hostname")
	}

	Cwd = file.SelfDir()

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()+os.Getppid()) + int64(crc32.ChecksumIEEE([]byte(Hostname))))
}
