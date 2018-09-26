package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// salt := "This is salt"
	fin, err := os.Open(os.Args[1])
	check(err)
	fout, err := os.Create(os.Args[2])
	check(err)
	r := csv.NewReader(bufio.NewReader(fin))
	w := csv.NewWriter(bufio.NewWriter(fout))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		h := sha256.New()
		h.Write([]byte(record[1]))
		record[1] = string(base64.StdEncoding.EncodeToString(h.Sum(nil)))
		w.Write(record)
		if err := w.Error(); err != nil {
			log.Fatalln("error writing csv:", err)
		}
	}
	w.Flush()
}
