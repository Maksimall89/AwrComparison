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
	ReportSummary				ReportSummary
}
// Report Summary
type ReportSummary struct{
	TopADDMFindingsByAverageActiveSessions	[]TopADDMFindingsByAverageActiveSessions
	LoadProfile								[]LoadProfile
	InstanceEfficiencyPercentages			[]InstanceEfficiencyPercentages
	Top10ForegroundEventsByTotalWaitTime	[]Top10ForegroundEventsByTotalWaitTime
	WaitClassesByTotalWaitTime				[]WaitClassesByTotalWaitTime
	HostCPU 								[]HostCPU
	InstanceCPU 							[]InstanceCPU
	IOProfile								[]IOProfile
	MemoryStatistics 						[]MemoryStatistics
	CacheSizes 								[]CacheSizes
	SharedPoolStatistics 					[]SharedPoolStatistics
}
// Top ADDM Findings by Average Active Sessions
type TopADDMFindingsByAverageActiveSessions struct{
	FindingName					string
	AvgActiveSessionsTask		float64
	PerActiveSessionsFinding	float64
	TaskName					string
	BeginSnapTime				string
	EndSnapTime					string
}
// Load Profile
type LoadProfile struct{
	Name			string
	PerSecond		float64
	PerTransaction	float64
	PerExec			float64
	PerCall			float64
}
// Instance Efficiency Percentages (Target 100%)
type InstanceEfficiencyPercentages struct{
	Name	string
	Value	float64
}
//Top 10 Foreground Events by Total Wait Time
type Top10ForegroundEventsByTotalWaitTime struct{
	Event			string
	Waits 			float64
	TotalWaitTime 	float64
	WaitAvg			float64
	PerDBTime		float64
	WaitClass		string
}
//Wait Classes by Total Wait Time
type WaitClassesByTotalWaitTime struct{
	WaitClass			string
	Waits				float64
	TotalWaitTime		float64
	AvgWait				float64
	PerDBTime			float64
	AvgActiveSessions	float64
}
// Host CPU
type HostCPU struct{
	CPUs	float64
	Cores	float64
	Sockets	float64
	LABegin	float64
	LAEnd	float64
	PerUser	float64
	PerSystem	float64
	PerWIO	float64
	PerIDLE	float64
}
// Instance CPU
type InstanceCPU struct{
	PerTotalCPU				float64
	PerBysuCPU				float64
	PerDBTimeWaiting		float64
}
// IO Profile
type IOProfile struct {
	Name			string
	RWPerSecond		float64
	ReadPerSecond	float64
	WritePerSecond	float64
}
// Memory Statistics
type MemoryStatistics struct {
	Name 	string
	Begin	float64
	End		float64
}
// Cache Sizes
type CacheSizes struct {
	Name 	string
	Begin	float64
	End	float64
}
// Shared Pool Statistics
type SharedPoolStatistics struct {
	Name	string
	Begin	float64
	End		float64
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

func fixDot(str string) float64{
	// replace , and .
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "M", "048576", -1)	// TODO доделать умножение 1 048 576
	str = strings.Replace(str, "K", "024", -1)	// 1024
	// convert type from string to float64
	val, err := strconv.ParseFloat(str, 64)
	if err != nil{
		// log.Println(err) // TODO edit
	}
	return 	val
}
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
			conf.SQLOrderByElapsedTime[i].ElapsedTime = fixDot(strArr[1])
			conf.SQLOrderByElapsedTime[i].Executions = fixDot(strArr[2])
			conf.SQLOrderByElapsedTime[i].ElapsedTimePerExec = fixDot(strArr[3])
			conf.SQLOrderByElapsedTime[i].Total = fixDot(strArr[4])
			conf.SQLOrderByElapsedTime[i].Cpu = fixDot(strArr[5])
			conf.SQLOrderByElapsedTime[i].IO = fixDot(strArr[6])
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
			conf.SQLOrderedByCPUTime[i].CPUTime = fixDot(strArr[1])
			conf.SQLOrderedByCPUTime[i].Executions = fixDot(strArr[2])
			conf.SQLOrderedByCPUTime[i].CPUPerExec = fixDot(strArr[3])
			conf.SQLOrderedByCPUTime[i].Total = fixDot(strArr[4])
			conf.SQLOrderedByCPUTime[i].ElapsedTime = fixDot(strArr[5])
			conf.SQLOrderedByCPUTime[i].CPU = fixDot(strArr[6])
			conf.SQLOrderedByCPUTime[i].IO = fixDot(strArr[7])
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
			conf.SQLOrderedByUserIOWaitTime[i].UserIOTime = fixDot(strArr[1])
			conf.SQLOrderedByUserIOWaitTime[i].Executions = fixDot(strArr[2])
			conf.SQLOrderedByUserIOWaitTime[i].UIOPerExec = fixDot(strArr[3])
			conf.SQLOrderedByUserIOWaitTime[i].Total = fixDot(strArr[4])
			conf.SQLOrderedByUserIOWaitTime[i].ElapsedTime = fixDot(strArr[5])
			conf.SQLOrderedByUserIOWaitTime[i].Cpu = fixDot(strArr[6])
			conf.SQLOrderedByUserIOWaitTime[i].IO = fixDot(strArr[7])
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
			conf.TopSQLWithTopEvents[i].PlanHash = fixDot(strArr[2])
			conf.TopSQLWithTopEvents[i].Executions = fixDot(strArr[3])
			conf.TopSQLWithTopEvents[i].Activity = fixDot(strArr[4])
			conf.TopSQLWithTopEvents[i].Event = strArr[5]
			conf.TopSQLWithTopEvents[i].EventPer = fixDot(strArr[6])
			conf.TopSQLWithTopEvents[i].RowSource = strArr[7]
			conf.TopSQLWithTopEvents[i].RowSourcePer = fixDot(strArr[8])
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
			conf.TopSQLWithTopRowSources[i].PlanHash = fixDot(strArr[2])
			conf.TopSQLWithTopRowSources[i].Executions = fixDot(strArr[3])
			conf.TopSQLWithTopRowSources[i].Activity = fixDot(strArr[4])
			conf.TopSQLWithTopRowSources[i].RowSource = strArr[5]
			conf.TopSQLWithTopRowSources[i].RowSourcePer = fixDot(strArr[6])
			conf.TopSQLWithTopRowSources[i].TopEvent = strArr[7]
			conf.TopSQLWithTopRowSources[i].EventPer = fixDot(strArr[8])
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
			conf.OperatingSystemStatistics[i].Value = fixDot(strArr[2])
			conf.OperatingSystemStatistics[i].EndValue = fixDot(strArr[3])
			i++
		}
	}
	if value, ok := maps["Report Summary"]; ok {

		var textBodyTwo []string
		var val string
		//var strArrItem []string
		textBody = regexp.MustCompile(`<p />(.*?)<p /><`).Split(value, -1)	// split line
		//conf.ReportSummary = make(ReportSummary)  // -1 because first line not contain information

		for _, iter := range textBody{
			strArr = regexp.MustCompile(`summary="(.*?)"`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {	// if we can't select to next line
				continue
			}

			switch strArr[1]{
			case "This table displays top ADDM findings by average active sessions":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.TopADDMFindingsByAverageActiveSessions = make([]TopADDMFindingsByAverageActiveSessions, (len(textBodyTwo) - 1)) // -1 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td scope="row" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.TopADDMFindingsByAverageActiveSessions[i].FindingName = strArr[1]
					conf.ReportSummary.TopADDMFindingsByAverageActiveSessions[i].AvgActiveSessionsTask = fixDot(strArr[2])
					conf.ReportSummary.TopADDMFindingsByAverageActiveSessions[i].PerActiveSessionsFinding = fixDot(strArr[3])
					conf.ReportSummary.TopADDMFindingsByAverageActiveSessions[i].TaskName = strArr[4]
					conf.ReportSummary.TopADDMFindingsByAverageActiveSessions[i].BeginSnapTime = strArr[5]
					conf.ReportSummary.TopADDMFindingsByAverageActiveSessions[i].EndSnapTime = strArr[6]
					i++
				}
			case "This table displays load profile":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.LoadProfile = make([]LoadProfile, (len(textBodyTwo) - 1)) // -1 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?):</td><td( align="right")* class='\w+'>\s*(.*?)</td><td( align="right")* class='\w+'>\s*(.*?)</td><td( align="right")* class='\w+'>\s*(.*?)</td><td( align="right")* class='\w+'>\s*(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.LoadProfile[i].Name = strArr[1]
					conf.ReportSummary.LoadProfile[i].PerSecond = fixDot(strArr[3])
					conf.ReportSummary.LoadProfile[i].PerTransaction = fixDot(strArr[5])
					conf.ReportSummary.LoadProfile[i].PerExec = fixDot(strArr[7])
					conf.ReportSummary.LoadProfile[i].PerCall = fixDot(strArr[9])
					i++
				}
			case "This table displays instance efficiency percentages":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.InstanceEfficiencyPercentages = make([]InstanceEfficiencyPercentages, (len(textBodyTwo)*2 - 3)) // -3 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>\s*(.*?)</td>(<td class='\w+'>(.*?):</td><td align="right" class='\w+'>\s*(.*?)</td>)*</tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.InstanceEfficiencyPercentages[i].Name = strArr[1]
					conf.ReportSummary.InstanceEfficiencyPercentages[i].Value = fixDot(strArr[2])
					i++
					if strArr[4] == "" {	// last line without content
						break
					}
					conf.ReportSummary.InstanceEfficiencyPercentages[i].Name = strArr[4]
					conf.ReportSummary.InstanceEfficiencyPercentages[i].Value = fixDot(strArr[5])
					i++
				}
			case "This table displays top 10 wait events by total wait time":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime = make([]Top10ForegroundEventsByTotalWaitTime, (len(textBodyTwo) - 1)) // -3 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime[i].Event = strArr[1]
					conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime[i].Waits = fixDot(strArr[2])
					conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime[i].TotalWaitTime = fixDot(strArr[3])
					conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime[i].WaitAvg = fixDot(strArr[4])
					conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime[i].PerDBTime = fixDot(strArr[5])
					conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime[i].WaitClass = strArr[6]
					i++
				}
			case "This table displays wait class statistics ordered by total wait time":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.WaitClassesByTotalWaitTime = make([]WaitClassesByTotalWaitTime, (len(textBodyTwo) - 1)) // -3 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.WaitClassesByTotalWaitTime[i].WaitClass = strArr[1]
					conf.ReportSummary.WaitClassesByTotalWaitTime[i].Waits = fixDot(strArr[2])
					conf.ReportSummary.WaitClassesByTotalWaitTime[i].TotalWaitTime= fixDot(strArr[3])
					conf.ReportSummary.WaitClassesByTotalWaitTime[i].AvgWait= fixDot(strArr[4])
					conf.ReportSummary.WaitClassesByTotalWaitTime[i].PerDBTime= fixDot(strArr[5])
					conf.ReportSummary.WaitClassesByTotalWaitTime[i].AvgActiveSessions= fixDot(strArr[6])
					i++
				}
			case "This table displays system load statistics":
				strArr = regexp.MustCompile(`<tr><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
				conf.ReportSummary.HostCPU = make([]HostCPU, 1) // everytime we have only one line
				if len(strArr) == 10 {
					conf.ReportSummary.HostCPU[0].CPUs = fixDot(strArr[1])
					conf.ReportSummary.HostCPU[0].Cores = fixDot(strArr[2])
					conf.ReportSummary.HostCPU[0].Sockets= fixDot(strArr[3])
					conf.ReportSummary.HostCPU[0].LABegin= fixDot(strArr[4])
					conf.ReportSummary.HostCPU[0].LAEnd= fixDot(strArr[5])
					conf.ReportSummary.HostCPU[0].PerUser= fixDot(strArr[6])
					conf.ReportSummary.HostCPU[0].PerSystem= fixDot(strArr[7])
					conf.ReportSummary.HostCPU[0].PerWIO = fixDot(strArr[8])
					conf.ReportSummary.HostCPU[0].PerIDLE= fixDot(strArr[9])
				}
			case "This table displays CPU usage and wait statistics":
				strArr = regexp.MustCompile(`<td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
				conf.ReportSummary.InstanceCPU = make([]InstanceCPU, 1) // everytime we have only one line
				if len(strArr) == 4 {
					conf.ReportSummary.InstanceCPU[0].PerTotalCPU = fixDot(strArr[1])
					conf.ReportSummary.InstanceCPU[0].PerBysuCPU = fixDot(strArr[2])
					conf.ReportSummary.InstanceCPU[0].PerDBTimeWaiting= fixDot(strArr[3])
				}
			case "This table displays IO profile":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.IOProfile = make([]IOProfile, (len(textBodyTwo) - 1)) // -3 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.IOProfile[i].Name = strArr[1]
					conf.ReportSummary.IOProfile[i].RWPerSecond = fixDot(strArr[2])
					conf.ReportSummary.IOProfile[i].ReadPerSecond= fixDot(strArr[3])
					conf.ReportSummary.IOProfile[i].WritePerSecond= fixDot(strArr[4])
					i++
				}
			case "This table displays memory statistics":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.MemoryStatistics = make([]MemoryStatistics, (len(textBodyTwo) - 1)) // -3 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.MemoryStatistics[i].Name = strArr[1]
					conf.ReportSummary.MemoryStatistics[i].Begin = fixDot(strArr[2])
					conf.ReportSummary.MemoryStatistics[i].End= fixDot(strArr[3])

					i++
				}
			case "This table displays cache sizes and other statistics for                     different types of cache":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.CacheSizes = make([]CacheSizes, (len(textBodyTwo)*2 - 3)) // -3 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td>(<td class='\w+'>(.*?):</td><td align="right" class='\w+'>(.*?)</td>)*</tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.CacheSizes[i].Name = strArr[1]
					conf.ReportSummary.CacheSizes[i].Begin = fixDot(strArr[2])
					conf.ReportSummary.CacheSizes[i].End= fixDot(strArr[3])

					if strArr[5] == ""{	// last line
						break
					}

					i++
					conf.ReportSummary.CacheSizes[i].Name = strArr[5]
					conf.ReportSummary.CacheSizes[i].Begin = fixDot(strArr[6])
					i++
				}
			case "This table displays shared pool statistics":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)// split line
				conf.ReportSummary.SharedPoolStatistics = make([]SharedPoolStatistics, (len(textBodyTwo) - 1)) // -3 because last second item not contain information
				for _, val = range textBodyTwo{
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {	// if we can't select to next line
						continue
					}
					conf.ReportSummary.SharedPoolStatistics[i].Name = strArr[1]
					conf.ReportSummary.SharedPoolStatistics[i].Begin = fixDot(strArr[2])
					conf.ReportSummary.SharedPoolStatistics[i].End= fixDot(strArr[3])
					i++
				}

			default : continue
			}
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
