package data

import (
	"os"
	"github.com/kyani-inc/kms-object-reps/src/app/log"
	"encoding/csv"
	"Easy-GoLang-Parallel-ForEach/Parallel/Parallel"
)

var (
	Common_Passwords Parallel.ConcurrentSlice
)

func Init() {
	f, err := os.Open("Auth/Data/cp.csv")

	if err != nil {
		log.Println(err)
	}

	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()

	slice := Parallel.ToConcurrentSlice(lines[0]);

	if err != nil {
		log.Println(err)
	}

	Common_Passwords = *slice;
}
