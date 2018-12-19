package main

import (
	"encoding/json"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"os"
	"time"
)

// Init config program
func (conf *Config) Init() {
	fileConfig, err := os.Open("config.json")
	if err != nil {
		log.Println(err)
	}
	defer fileConfig.Close()
	decoder := json.NewDecoder(fileConfig)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Println(err)
	}

}
// sent data to DB
func SentDB(conf *Config, data *PageData)  {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: conf.NameDB,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Init new column  and param
	var tags = make(map[string]string)
	var fields = make(map[string]interface{})
	//var times int64
	var tm time.Time
	var counter int64

	// Read files from catalog
	for _, line := range data.ListSQLText {
		counter = 0

		tags["SQLDescribe"] = line.SQLDescribe
		tags["DateStart"] = data.WorkInfo.WISnapshotInformation[0].SnapTime
		tags["DateStop"] = data.WorkInfo.WISnapshotInformation[1].SnapTime

		fields["SQLId"] = line.SQLId
		fields["SQLText"] = line.SQLText

		// TODO Make real time
		/*
		times, err = strconv.ParseInt(line.SQLDescribe[0:10], 10, 64)
		if err != nil {
			panic(err)
		}

		*/

		// Shift time and convert time
		tm = time.Unix(time.Now().Unix(), 10000 + counter)

		// Create string on db
		pt, err := client.NewPoint(conf.Measurement, tags, fields, tm)
		if err != nil {
			log.Fatal(err)
		}

		counter++
		// Add new line in list
		bp.AddPoint(pt)
	}

	// Write the batch
	if err := conf.Client.Write(bp); err != nil {
		log.Fatal(err)
	}

	log.Printf("Date upload in DB %s, measurement %s.", data.WorkInfo.WIHostInformation[0].HostName, conf.NameDB, conf.Measurement)
}

// TODO select
func (conf *Config) GetDB() {
	var err error
	conf.Result, err = QueryDB(conf, fmt.Sprintf(`SELECT SQLId, DateStart, DateStop FROM %s`, conf.Measurement))
	if err != nil {
		log.Fatal(err)
	}
/*
	for x, _ := range conf.Result[0].Series[0].Values{
		fmt.Println(conf.Result[0].Series[0].Values[x][1]) // sqlID - 1; datastart - 2; data stop -3
	}
*/
	log.Printf("Data from %s.%s get", conf.NameDB, conf.Measurement)
}

func QueryDB(conf *Config, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: conf.NameDB,
	}
	if response, err := conf.Client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
func (conf *Config) CreateDB() {
	_, err := QueryDB(conf, fmt.Sprintf("CREATE DATABASE %s", conf.NameDB))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Database %s created.", conf.NameDB)
}
func (conf *Config) DeleteDB() {
	_, err := QueryDB(conf, fmt.Sprintf("DROP DATABASE %s", conf.NameDB))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Database %s deleted.", conf.NameDB)
}
