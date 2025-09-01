package main

import (
	"github.com/philipslstwoyears/todo/internal/dataIO"
	"github.com/philipslstwoyears/todo/internal/flags"
	"log"
)

func main() {
	flag, err := flags.New()
	if err != nil {
		log.Fatal(err)
	}

	data, err := dataIO.Read(flag.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
	//uniqData := uniq.Uniq(data, flag)
	//err = dataIO.Write(uniqData, flag.OutputFile)
}
