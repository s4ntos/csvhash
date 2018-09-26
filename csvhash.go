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
	columns := []int{}
	switch fileType := os.Args[1]; fileType {
	case "csv3":
		columns = []int{2, 3, 4}
	case "csv2":
		columns = []int{2, 3}
	default:
		columns = []int{2}
	}
	salt := os.Args[2]
	fin, err := os.Open(os.Args[3])
	check(err)
	fout, err := os.Create(os.Args[4])
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
		for _, c := range columns {
			h.Write([]byte(salt + record[c]))
			record[c] = string(base64.StdEncoding.EncodeToString(h.Sum(nil)))
		}
		w.Write(record)
		if err := w.Error(); err != nil {
			log.Fatalln("error writing csv:", err)
		}
	}
	w.Flush()
}
