package main
/*
	Doc:
	http://nonfunctionaltestingtools.blogspot.ru/2015/04/steps-to-analyze-awr-report-in-oracle.html
	https://habrahabr.ru/post/189574/
	http://www.sql.ru/blogs/oracleandsql/2097
	http://www.dbas-oracle.com/2013/05/10-steps-to-analyze-awr-report-in-oracle.html
*/
import (
	"log"
	"fmt"
	"os"
	"encoding/json"
	"time"
	"bufio"
	"strings"
	"regexp"
	"errors"
	"strconv"
)

const configFileName  = "config.json"

type Config struct {
	TelegramBotToken string
	OwnName          string
}
type MainTable struct {
	SQLOrderByElapsedTime      	[]SQLOrderByElapsedTime
	CompleteListOfSQLText      	[]CompleteListOfSQLText
	SQLOrderedByCPUTime        	[]SQLOrderedByCPUTime
	SQLOrderedByUserIOWaitTime 	[]SQLOrderedByUserIOWaitTime
	TopSQLWithTopEvents        	[]TopSQLWithTopEvents
	TopSQLWithTopRowSources    	[]TopSQLWithTopRowSources
	OperatingSystemStatistics	[]OperatingSystemStatistics
}
// SQL ordered by Elapsed Time
type SQLOrderByElapsedTime struct{
	ElapsedTime			float64
	Executions			float64
	ElapsedTimePerExec	float64
	Total				float64
	Cpu					float64
	IO					float64
	SQLID				string
	SQLModule			string
	SQLText				string
}
// SQL ordered by CPU Time
type SQLOrderedByCPUTime struct{
	CPUTime				float64
	Executions			float64
	CPUPerExec			float64
	Total				float64
	ElapsedTime			float64
	CPU					float64
	IO					float64
	SQLID				string
	SQLModule			string
	SQLText				string
}
// SQL ordered by User I/O Wait Time
type SQLOrderedByUserIOWaitTime struct{
	UserIOTime			float64
	Executions			float64
	UIOPerExec			float64
	Total				float64
	ElapsedTime			float64
	Cpu					float64
	IO					float64
	SQLID				string
	SQLModule			string
	SQLText				string
}
// Top SQL with Top Events
type TopSQLWithTopEvents struct{
	SQLID        string
	PlanHash     float64
	Executions   float64
	Activity     float64
	Event        string
	ElapsedTime  float64
	EventPer     float64
	RowSource    string
	RowSourcePer float64
	SQLText      string
}
// Top SQL with Top Row Sources
type TopSQLWithTopRowSources struct{
	SQLID				string
	PlanHash			float64
	Executions			float64
	Activity			float64
	RowSource			string
	RowSourcePer		float64
	TopEvent			string
	EventPer			float64
	SQLText				string
}
//Operating System Statistics
type OperatingSystemStatistics struct{
	Statistic		string
	Value			float64
	EndValue		float64
}
// Complete List of SQL Text
type CompleteListOfSQLText struct{
	SQLID		string
	SQLText		string
}
type worker interface{
	tableAnalyzer()
	sqlAnalyzer()
}
// reading text from a file
func readFile(name string) (string, error)  {
	var body string	// all text awr html
	fi, err := os.Open(name) // open file for read
	if err != nil{
		return "", err
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {	// read all html into body
		body += scanner.Text() //+ "\n"
	}
	return body, nil
}
// TODO парсер лога и запись его в структуры
func parser(conf *MainTable, maps map[string]string) ()  {
	var textBody []string	// text section
	var strArr []string	// text line
	var i int	// counter

	if value, ok := maps["Complete List of SQL Text"]; ok {
		textBody = regexp.MustCompile(`<a class="awr" name=".+?"></a>`).Split(value, -1)	// split line
		conf.CompleteListOfSQLText = make([]CompleteListOfSQLText, (len(textBody) - 3)) // -3 because last second item not contain information
		for _, iter := range textBody{
			strArr = regexp.MustCompile(`(.+?)</td><td class='\w+'>(.+?)</td>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {	// if we can't select to next line
				continue
			}
			// fill in our struct
			conf.CompleteListOfSQLText[i].SQLID = strArr[1]
			conf.CompleteListOfSQLText[i].SQLText = strArr[2]
			i++
		}
	}
	if value, ok := maps["SQL ordered by Elapsed Time"]; ok {
		i = 0
		//textBody := regexp.MustCompile(`<tr><td align`).Split(value, -1)	// split line
		textBody =  strings.Split(value, `<tr><td align="right" `)                     // split line
		conf.SQLOrderByElapsedTime = make([]SQLOrderByElapsedTime, (len(textBody) -1)) // -1 because first line not contain information

		for _, iter := range textBody{
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td scope="row" class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 {	// if we can't select to next line
				continue
			}
			// fill in our struct
			conf.SQLOrderByElapsedTime[i].ElapsedTime, _ = strconv.ParseFloat(strArr[1], 64)
			conf.SQLOrderByElapsedTime[i].Executions, _ = strconv.ParseFloat(strArr[2], 64)
			conf.SQLOrderByElapsedTime[i].ElapsedTimePerExec, _ = strconv.ParseFloat(strArr[3], 64)
			conf.SQLOrderByElapsedTime[i].Total, _ = strconv.ParseFloat(strArr[4], 64)
			conf.SQLOrderByElapsedTime[i].Cpu, _ = strconv.ParseFloat(strArr[5], 64)
			conf.SQLOrderByElapsedTime[i].IO, _ = strconv.ParseFloat(strArr[6], 64)
			conf.SQLOrderByElapsedTime[i].SQLID = strArr[7]
			conf.SQLOrderByElapsedTime[i].SQLModule = strArr[8]
			conf.SQLOrderByElapsedTime[i].SQLText = strArr[9]
			i++
		}
	}
	if value, ok := maps["SQL ordered by CPU Time"]; ok {
		i = 0
		textBody =  strings.Split(value, `<tr><td align="right" `)                 // split line
		conf.SQLOrderedByCPUTime = make([]SQLOrderedByCPUTime, (len(textBody) -1)) // -1 because first line not contain information

		for _, iter := range textBody{
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td scope="row" class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 {	// if we can't select to next line
				continue
			}
			// fill in our struct
			conf.SQLOrderedByCPUTime[i].CPUTime, _ = strconv.ParseFloat(strArr[1], 64)
			conf.SQLOrderedByCPUTime[i].Executions, _ = strconv.ParseFloat(strArr[2], 64)
			conf.SQLOrderedByCPUTime[i].CPUPerExec, _ = strconv.ParseFloat(strArr[3], 64)
			conf.SQLOrderedByCPUTime[i].Total, _ = strconv.ParseFloat(strArr[4], 64)
			conf.SQLOrderedByCPUTime[i].ElapsedTime, _ = strconv.ParseFloat(strArr[5], 64)
			conf.SQLOrderedByCPUTime[i].CPU, _ = strconv.ParseFloat(strArr[6], 64)
			conf.SQLOrderedByCPUTime[i].IO, _ = strconv.ParseFloat(strArr[7], 64)
			conf.SQLOrderedByCPUTime[i].SQLID = strArr[8]
			conf.SQLOrderedByCPUTime[i].SQLModule = strArr[9]
			conf.SQLOrderedByCPUTime[i].SQLText = strArr[10]
			i++
		}
	}
	if value, ok := maps["SQL ordered by User I/O Wait Time"]; ok {
		i = 0
		textBody =  strings.Split(value, `<tr><td align="right" `)                               // split line
		conf.SQLOrderedByUserIOWaitTime = make([]SQLOrderedByUserIOWaitTime, (len(textBody) -1)) // -1 because first line not contain information

		for _, iter := range textBody{
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td scope="row" class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 {	// if we can't select to next line
				continue
			}
			// fill in our struct
			conf.SQLOrderedByUserIOWaitTime[i].UserIOTime, _ = strconv.ParseFloat(strArr[1], 64)
			conf.SQLOrderedByUserIOWaitTime[i].Executions, _ = strconv.ParseFloat(strArr[2], 64)
			conf.SQLOrderedByUserIOWaitTime[i].UIOPerExec, _ = strconv.ParseFloat(strArr[3], 64)
			conf.SQLOrderedByUserIOWaitTime[i].Total, _ = strconv.ParseFloat(strArr[4], 64)
			conf.SQLOrderedByUserIOWaitTime[i].ElapsedTime, _ = strconv.ParseFloat(strArr[5], 64)
			conf.SQLOrderedByUserIOWaitTime[i].Cpu, _ = strconv.ParseFloat(strArr[6], 64)
			conf.SQLOrderedByUserIOWaitTime[i].IO, _ = strconv.ParseFloat(strArr[7], 64)
			conf.SQLOrderedByUserIOWaitTime[i].SQLID = strArr[8]
			conf.SQLOrderedByUserIOWaitTime[i].SQLModule = strArr[9]
			conf.SQLOrderedByUserIOWaitTime[i].SQLText = strArr[10]
			i++
		}
	}
	if value, ok := maps["Top SQL with Top Events"]; ok {
		i = 0
		textBody =  strings.Split(value, `<tr><td align="right" `)// split line
		conf.TopSQLWithTopEvents = make([]TopSQLWithTopEvents, (len(textBody) -1))  // -1 because first line not contain information

		for _, iter := range textBody{
			strArr = regexp.MustCompile(`class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 {	// if we can't select to next line
				continue
			}
			// fill in our struct
			conf.TopSQLWithTopEvents[i].SQLID = strArr[1]
			conf.TopSQLWithTopEvents[i].PlanHash, _ = strconv.ParseFloat(strArr[2], 64)
			conf.TopSQLWithTopEvents[i].Executions, _ = strconv.ParseFloat(strArr[3], 64)
			conf.TopSQLWithTopEvents[i].Activity, _ = strconv.ParseFloat(strArr[4], 64)
			conf.TopSQLWithTopEvents[i].Event = strArr[5]
			conf.TopSQLWithTopEvents[i].EventPer, _ = strconv.ParseFloat(strArr[6], 64)
			conf.TopSQLWithTopEvents[i].RowSource = strArr[7]
			conf.TopSQLWithTopEvents[i].RowSourcePer, _ = strconv.ParseFloat(strArr[8], 64)
			conf.TopSQLWithTopEvents[i].SQLText = strArr[9]
			i++
		}
	}
	if value, ok := maps["Top SQL with Top Row Sources"]; ok {
		i = 0
		textBody =  strings.Split(value, `<tr><td align="right" `)// split line
		conf.TopSQLWithTopRowSources = make([]TopSQLWithTopRowSources, (len(textBody) -1))  // -1 because first line not contain information

		for _, iter := range textBody{
			strArr = regexp.MustCompile(`class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {	// if we can't select to next line
				continue
			}
			// fill in our struct
			conf.TopSQLWithTopRowSources[i].SQLID = strArr[1]
			conf.TopSQLWithTopRowSources[i].PlanHash, _ = strconv.ParseFloat(strArr[2], 64)
			conf.TopSQLWithTopRowSources[i].Executions, _ = strconv.ParseFloat(strArr[3], 64)
			conf.TopSQLWithTopRowSources[i].Activity, _ = strconv.ParseFloat(strArr[4], 64)
			conf.TopSQLWithTopRowSources[i].RowSource = strArr[5]
			conf.TopSQLWithTopRowSources[i].RowSourcePer, _ = strconv.ParseFloat(strArr[6], 64)
			conf.TopSQLWithTopRowSources[i].TopEvent = strArr[7]
			conf.TopSQLWithTopRowSources[i].EventPer, _ = strconv.ParseFloat(strArr[8], 64)
			conf.TopSQLWithTopRowSources[i].SQLText = strArr[9]
			i++
		}
	}
	if value, ok := maps["Operating System Statistics"]; ok {
		i = 0
		textBody =  strings.Split(value, `<tr><td scope="row" `)// split line
		conf.OperatingSystemStatistics = make([]OperatingSystemStatistics, (len(textBody) -1))  // -1 because first line not contain information

		for _, iter := range textBody{
			strArr = regexp.MustCompile(`class='\w+'>(.+?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.+?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {	// if we can't select to next line
				continue
			}
			// fill in our struct
			conf.OperatingSystemStatistics[i].Statistic = strArr[1]
			conf.OperatingSystemStatistics[i].Value, _ = strconv.ParseFloat(strArr[2], 64)
			conf.OperatingSystemStatistics[i].EndValue, _ = strconv.ParseFloat(strArr[3], 64)

			i++
		}
	}
}

// TODO анализатор таблиц
func tableAnalyzer(){

}
// TODO анализатор не оптимальных запросов
func sqlAnalyzer()  {

}
// TODO web-server с загрузкой лога через веб морду и выводом информации по логу на экран
func server()  {

}
func (conf *Config) init() {
	//init configuration
	configuration := Config{}
	// open config-file
	file, err := os.Open(configFileName)
	defer file.Close()

	if err != nil {
		log.Println(err)
		// default configuration
		configuration.OwnName = "Maksimall89"
		configuration.TelegramBotToken = "3257"
	}else{

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&configuration)
		if err != nil {
			log.Println(err)
		}
	}
	return
}
// create maps with element
func createMaps(textInput string, maps map[string]string) error{
	textBody := strings.Split(textInput, `<h3 class="awr">`)
	for _, text := range  textBody{
		if reg, _ := regexp.MatchString(`(.*?)</h3>([\D|\d]*)`, string(text)); reg {
			s := regexp.MustCompile(`(.*?a>)*(.*?)</h3>([\D|\d]+)`).FindStringSubmatch(string(text))
			maps[s[2]] = s[3]
		}
	}
	if len(maps) == 0{
		return errors.New("Not found elements map in the AWR")
	}
	return nil
}

func main() {

	// configurator for logger
	var str = "log"	// name folder for logs

	// check what folder log is exist
	_, err := os.Stat(str)
	if os.IsNotExist(err) {
		os.MkdirAll(str, os.ModePerm);
	}
	str =  fmt.Sprintf("%s/%d-%02d-%02d-%02d-%02d-%02d-logFile.log", str, time.Now().Year(),time.Now().Month(),time.Now().Day(),time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	// open a file
	f, err := os.OpenFile(str, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f) // TODO config logs
	log.SetPrefix("AWRcompar ")

	// read config file
	configuration := Config{}
	configuration.init()

	log.Println("Start work.")

	work := MainTable{}

	// read file
	body, err := readFile("awr/global_awr_report_111755_111758.html")
	if err != nil{
		log.Fatal(err)
	}

	// create map
	maps := make(map[string]string)
	err = createMaps(body, maps)
	if err != nil{
		log.Fatal(err)
	}

	//log.Println(body)

	//log.Println(maps["SQL ordered by Elapsed Time"])

	parser(&work, maps)

}
