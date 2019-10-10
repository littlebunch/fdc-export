// exporter dumps documents from a couchbase bucket to a file
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/littlebunch/fdc-api/ds"
	"github.com/littlebunch/fdc-api/ds/cb"
	fdc "github.com/littlebunch/fdc-api/model"
)

var (
	c   = flag.String("c", "config.yml", "YAML Config file")
	l   = flag.String("l", "/tmp/export.out", "send log output to this file -- defaults to /tmp/ingest.out")
	o   = flag.String("o", "", "Output json file")
	t   = flag.String("t", "", "Export document type")
	s   = flag.Int64("s", 0, "Document offset to begin export.  Defaults to 0")
	m   = flag.Int64("m", 5000, "Max number of documents to export. Defaults to 5000")
	cnt = 0
	cs  fdc.Config
)

func init() {
	var (
		err   error
		lfile *os.File
	)

	lfile, err = os.OpenFile(*l, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", *l, ":", err)
	}
	m := io.MultiWriter(lfile, os.Stdout)
	log.SetOutput(m)
}

func main() {

	log.Print("Starting export")
	flag.Parse()
	var dt fdc.DocType
	dtype := dt.ToDocType(*t)
	if dtype == 999 {
		log.Fatalln("Valid t option is required")
	}

	var (
		cb cb.Cb
		ds ds.DataSource
	)
	cs.GetConfig(c)
	// create a datastore and connect to it
	ds = &cb
	if err := ds.ConnectDs(cs); err != nil {
		log.Fatalln("Cannot connect to datastore ", err)
	}
	// implement the Ingest interface
	if dtype == fdc.FOOD || dtype == fdc.NUTDATA {
		err := exportData(*o, ds, dtype, *s, *m)
		if err != nil {
			fmt.Printf("Error on export: %v\n", err)
		}
	} else {
		fmt.Println("Invalid t option -- must be FOOD or NUTDATA")
	}

	log.Println("Finished.")
	ds.CloseDs()
	os.Exit(0)
}
func exportData(ofile string, dc ds.DataSource, dt fdc.DocType, start int64, max int64) error {
	f, err := os.Create(ofile)
	var (
		foods []interface{}
	)
	if err != nil {
		return err
	}
	defer f.Close()
	where := fmt.Sprintf("type=\"%s\" ", dt.ToString(dt))
	start = 0

	for {

		food, err := dc.Browse(cs.CouchDb.Bucket, where, start, max, "fdcId", "asc")
		if err != nil {
			fmt.Printf("%v\n", err)
			//return err
		}
		fmt.Println("len = ", len(food))
		if len(food) == 0 {
			break
		}
		for fd := range food {
			foods = append(foods, food[fd])
		}
		start += max

	}
	b, err := json.Marshal(foods)
	n, err := f.Write(b)
	fmt.Printf("Wrote %d bytes with %v\n", n, err)
	return err
}