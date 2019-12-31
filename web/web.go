package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func Download(url string) error {
	filename := strings.Split(url, "/")[len(strings.Split(url, "/"))-1]

	aout, err := os.Create(filename + ".tmp")
	if err != nil {
		return err
	}
	defer aout.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(aout, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	return os.Rename(filename+".tmp", filename)
}
