package main

import (
	ES "gopkg.in/olivere/elastic.v3"

	"github.com/pcman312/darwin_test/es"
	PROC "github.com/shirou/gopsutil/process"
)

func main() {
	ES.NewClient()
	client := es.ESClient{}
	client.DoESStuff()
	PROC.NewProcess(12345)
}
