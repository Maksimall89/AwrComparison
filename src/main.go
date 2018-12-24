package main

/*
	Doc:
	http://nonfunctionaltestingtools.blogspot.ru/2015/04/steps-to-analyze-awr-report-in-oracle.html
	https://habrahabr.ru/post/189574/
	http://www.sql.ru/blogs/oracleandsql/2097
	http://www.dbas-oracle.com/2013/05/10-steps-to-analyze-awr-report-in-oracle.html
	https://studfiles.net/preview/2426969/page:12/
*/

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// reading text from a file
func readFile(name string) (string, error) {
	var body string          // all text awr html
	fi, err := os.Open(name) // open file for read
	if err != nil {
		return "", err
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() { // read all html into body
		body += scanner.Text() //+ "\n"
	}
	return body, nil
}

// TODO доделать конвертацию даты в нормальный формат, чтобы дата в бд соотвествовала, TS когда сняли AWR отчёт.
// Для этого надо разобраться с unix как переводить время из 11-Дек-17 20:30:47 в timeshamp
func parseTimeStamp(utime string) (string, error) {
	i, err := strconv.ParseInt(utime, 10, 64)
	if err != nil {
		return "", err
	}
	t := time.Unix(i, 0)
	fmt.Println(t)
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	return t.Format(time.UnixDate), nil
}

func fixDot(str string) float64 {
	// replace , and .
	// convert type from string to float64
	if str == "&#160;" {
		return 0
	}
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "M", "048576", -1) // TODO доделать умножение 1 048 576
	str = strings.Replace(str, "K", "024", -1)    // 1024

	val, err := strconv.ParseFloat(str, 64)

	if err != nil {
		log.Println(err)
	}
	return val
}

func compare(data *PageData, conf *Config) {
	var str string
	counter := 0
	for _, item := range data.ListSQLText {

		str = ""

		GetDBinfo(conf, item.SQLId)

		for x, pair := range conf.Results[0].Series {
			str += fmt.Sprintf("запрос встречался ранее в AWR за участок от %s до %s, ", pair.Values[x][1], pair.Values[x][1])
		}

		if str == "" {
			data.ListSQLText[counter].TextUI = "запрос не встречался ранее"
		} else {
			data.ListSQLText[counter].TextUI = str
		}

		counter++
	}
}
func parser(conf *MainTable, maps map[string]string) {
	var textBody []string  // text section
	var textTable []string // text section
	var strArr []string    // text line
	var i int              // count line

	if value, ok := maps["Work Info"]; ok {
		textTable = regexp.MustCompile(`<table border="0" `).Split(value, -1) // split to table
		// This table displays database instance information
		textBody = strings.Split(textTable[1], `<tr><td scope="row"`) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(` class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 { // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.WorkInfo.WIDatabaseInstanceInformation = append(conf.WorkInfo.WIDatabaseInstanceInformation, WIDatabaseInstanceInformation{
				DBName:      strArr[1],
				DBId:        fixDot(strArr[2]),
				Instance:    strArr[3],
				Instnum:     fixDot(strArr[4]),
				StartupTime: strArr[5],
				Release:     strArr[6],
				RAC:         strArr[7],
			})
		}
		// This table displays host information
		textBody = strings.Split(textTable[2], `<tr><td scope="row"`) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(` class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                                                                                                                                                                                                                          // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.WorkInfo.WIHostInformation = append(conf.WorkInfo.WIHostInformation, WIHostInformation{
				HostName: strArr[1],
				Platform: strArr[2],
				CPUs:     fixDot(strArr[3]),
				Cores:    fixDot(strArr[4]),
				Sockets:  fixDot(strArr[5]),
				MemoryGB: fixDot(strArr[6]),
			})
		}
		// This table displays snapshot information
		textBody = strings.Split(textTable[3], `<tr><td scope="row"`) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(` class='\w+'>(.*?)</td><td[ align="right"]* class='\w+'>(.*?)</td><td[ align="center"]* class='\w+'>(.*?)</td><td[ align="right"]* class='\w+'>(.*?)</td><td[ align="right"]* class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                                                                                                                                                                                                             // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.WorkInfo.WISnapshotInformation = append(conf.WorkInfo.WISnapshotInformation, WISnapshotInformation{
				Name:           strArr[1],
				SnapId:         fixDot(strArr[2]),
				SnapTime:       strArr[3],
				Sessions:       fixDot(strArr[4]),
				CursorsSession: fixDot(strArr[5]),
			})
		}
	}
	if value, ok := maps["Complete List of SQL Text"]; ok {
		textBody = regexp.MustCompile(`<a class="awr" name=".+?"></a>`).Split(value, -1) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`(.+?)</td><td class='\w+'>(.+?)</td>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                                        // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.CompleteListOfSQLText = append(conf.CompleteListOfSQLText, CompleteListOfSQLText{
				SQLID:   strArr[1],
				SQLText: strArr[2],
			})
		}
	}
	if value, ok := maps["Foreground Wait Class"]; ok {
		textBody = strings.Split(value, `<tr><td scope="row" `) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                                                                                                                                                                                                                                       // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.WaitEventsStatistics.ForegroundWaitClass = append(conf.WaitEventsStatistics.ForegroundWaitClass, ForegroundWaitClass{
				WaitClass:     strArr[1],
				Waits:         fixDot(strArr[2]),
				PerTime:       fixDot(strArr[3]),
				TotalWaitTime: fixDot(strArr[4]),
				AvgWait:       fixDot(strArr[5]),
				PerDBTime:     fixDot(strArr[6]),
			})
		}
	}
	if value, ok := maps["Foreground Wait Events"]; ok {
		textBody = strings.Split(value, `<tr><td scope="row" `) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                                                                                                                                                                                                                                                                               // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.WaitEventsStatistics.ForegroundWaitEvents = append(conf.WaitEventsStatistics.ForegroundWaitEvents, ForegroundWaitEvents{
				Event:         strArr[1],
				Waits:         fixDot(strArr[2]),
				PerTime:       fixDot(strArr[3]),
				TotalWaitTime: fixDot(strArr[4]),
				AvgWait:       fixDot(strArr[5]),
				WaitsTxn:      fixDot(strArr[5]),
				PerDBTime:     fixDot(strArr[6]),
			})
		}
	}
	if value, ok := maps["Background Wait Events"]; ok {
		textBody = strings.Split(value, `<tr><td scope="row" `) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                                                                                                                                                                                                                                                                               // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.WaitEventsStatistics.BackgroundWaitEvents = append(conf.WaitEventsStatistics.BackgroundWaitEvents, BackgroundWaitEvents{
				Event:         strArr[1],
				Waits:         fixDot(strArr[2]),
				PerTime:       fixDot(strArr[3]),
				TotalWaitTime: fixDot(strArr[4]),
				AvgWait:       fixDot(strArr[5]),
				WaitsTxn:      fixDot(strArr[5]),
				PerbgTime:     fixDot(strArr[6]),
			})
		}
		/*
			for _, xx := range conf.WaitEventsStatistics.BackgroundWaitEvents{
				fmt.Println(xx.Event)
			}
		*/
	}
	if value, ok := maps["SQL ordered by Elapsed Time"]; ok {
		textBody = strings.Split(value, `<tr><td align="right" `) // split line

		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td scope="row" class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 { // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.SQLOrderByElapsedTime = append(conf.SQLOrderByElapsedTime, SQLOrderByElapsedTime{
				ElapsedTime:        fixDot(strArr[1]),
				Executions:         fixDot(strArr[2]),
				ElapsedTimePerExec: fixDot(strArr[3]),
				Total:              fixDot(strArr[4]),
				CPU:                fixDot(strArr[5]),
				IO:                 fixDot(strArr[6]),
				SQLID:              strArr[7],
				SQLModule:          strArr[8],
				SQLText:            strArr[9],
			})
		}
	}
	if value, ok := maps["SQL ordered by CPU Time"]; ok {
		textBody = strings.Split(value, `<tr><td align="right" `) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td scope="row" class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 { // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.SQLOrderedByCPUTime = append(conf.SQLOrderedByCPUTime, SQLOrderedByCPUTime{
				CPUTime:     fixDot(strArr[1]),
				Executions:  fixDot(strArr[2]),
				CPUPerExec:  fixDot(strArr[3]),
				Total:       fixDot(strArr[4]),
				ElapsedTime: fixDot(strArr[5]),
				CPU:         fixDot(strArr[6]),
				IO:          fixDot(strArr[7]),
				SQLID:       strArr[8],
				SQLModule:   strArr[9],
				SQLText:     strArr[10],
			})
		}
	}
	if value, ok := maps["SQL ordered by User I/O Wait Time"]; ok {
		textBody = strings.Split(value, `<tr><td align="right" `) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td scope="row" class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 { // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.SQLOrderedByUserIOWaitTime = append(conf.SQLOrderedByUserIOWaitTime, SQLOrderedByUserIOWaitTime{
				UserIOTime:  fixDot(strArr[1]),
				Executions:  fixDot(strArr[2]),
				UIOPerExec:  fixDot(strArr[3]),
				Total:       fixDot(strArr[4]),
				ElapsedTime: fixDot(strArr[5]),
				CPU:         fixDot(strArr[6]),
				IO:          fixDot(strArr[7]),
				SQLID:       strArr[8],
				SQLModule:   strArr[9],
				SQLText:     strArr[10],
			})
		}
	}
	if value, ok := maps["Top SQL with Top Events"]; ok {
		textBody = strings.Split(value, `<tr><td align="right" `) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row

			if len(strArr) == 0 { // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.TopSQLWithTopEvents = append(conf.TopSQLWithTopEvents, TopSQLWithTopEvents{
				SQLID:        strArr[1],
				PlanHash:     fixDot(strArr[2]),
				Executions:   fixDot(strArr[3]),
				Activity:     fixDot(strArr[4]),
				Event:        strArr[5],
				EventPer:     fixDot(strArr[6]),
				RowSource:    strArr[7],
				RowSourcePer: fixDot(strArr[8]),
				SQLText:      strArr[9],
			})
		}
	}
	if value, ok := maps["Top SQL with Top Row Sources"]; ok {
		textBody = strings.Split(value, `<tr><td align="right" `) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'><a class="awr" href=".*?">(.*?)</a></td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                                                                                                                                                                                                                                                                                                                                                   // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.TopSQLWithTopRowSources = append(conf.TopSQLWithTopRowSources, TopSQLWithTopRowSources{
				SQLID:        strArr[1],
				PlanHash:     fixDot(strArr[2]),
				Executions:   fixDot(strArr[3]),
				Activity:     fixDot(strArr[4]),
				RowSource:    strArr[5],
				RowSourcePer: fixDot(strArr[6]),
				TopEvent:     strArr[7],
				EventPer:     fixDot(strArr[8]),
				SQLText:      strArr[9],
			})
		}
	}
	if value, ok := maps["Operating System Statistics"]; ok {
		textBody = strings.Split(value, `<tr><td scope="row" `) // split line
		for _, iter := range textBody {
			strArr = regexp.MustCompile(`class='\w+'>(.+?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.+?)</td></tr>`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                                                                                                               // if we can't select to next line
				continue
			}
			// fill in our struct
			conf.OperatingSystemStatistics = append(conf.OperatingSystemStatistics, OperatingSystemStatistics{
				Statistic: strArr[1],
				Value:     fixDot(strArr[2]),
				EndValue:  fixDot(strArr[3]),
			})
		}
	}
	if value, ok := maps["Report Summary"]; ok {
		var textBodyTwo []string
		var val string
		textBody = regexp.MustCompile(`<p />(.*?)<p /><`).Split(value, -1) // split line

		for _, iter := range textBody {
			strArr = regexp.MustCompile(`summary="(.*?)"`).FindStringSubmatch(iter) // select item from row
			if len(strArr) == 0 {                                                   // if we can't select to next line
				continue
			}

			switch strArr[1] {
			case "This table displays top ADDM findings by average active sessions":
				textBodyTwo = regexp.MustCompile(`<tr><td class='\w+'>`).Split(iter, -1) // split line
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td scope="row" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                                                                                                            // if we can't select to next line
						continue
					}
					conf.ReportSummary.TopADDMFindingsByAverageActiveSessions = append(conf.ReportSummary.TopADDMFindingsByAverageActiveSessions, TopADDMFindingsByAverageActiveSessions{
						FindingName:              strArr[1],
						AvgActiveSessionsTask:    fixDot(strArr[2]),
						PerActiveSessionsFinding: fixDot(strArr[3]),
						TaskName:                 strArr[4],
						BeginSnapTime:            strArr[5],
						EndSnapTime:              strArr[6],
					})

				}
			case "This table displays load profile":
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1) // split line
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?):</td><td( align="right")* class='\w+'>\s*(.*?)</td><td( align="right")* class='\w+'>\s*(.*?)</td><td( align="right")* class='\w+'>\s*(.*?)</td><td( align="right")* class='\w+'>\s*(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                                                                                                                           // if we can't select to next line
						continue
					}
					conf.ReportSummary.LoadProfile = append(conf.ReportSummary.LoadProfile, LoadProfile{
						Name:           strArr[1],
						PerSecond:      fixDot(strArr[3]),
						PerTransaction: fixDot(strArr[5]),
						PerExec:        fixDot(strArr[7]),
						PerCall:        fixDot(strArr[9]),
					})
				}
			case "This table displays instance efficiency percentages":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1)                           // split line
				conf.ReportSummary.InstanceEfficiencyPercentages = make([]InstanceEfficiencyPercentages, len(textBodyTwo)*2-3) // -3 because last second item not contain information
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>\s*(.*?)</td>(<td class='\w+'>(.*?):</td><td align="right" class='\w+'>\s*(.*?)</td>)*</tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                                                       // if we can't select to next line
						continue
					}
					conf.ReportSummary.InstanceEfficiencyPercentages[i].Name = strArr[1]
					conf.ReportSummary.InstanceEfficiencyPercentages[i].Value = fixDot(strArr[2])
					i++
					if strArr[4] == "" { // last line without content
						break
					}
					conf.ReportSummary.InstanceEfficiencyPercentages[i].Name = strArr[4]
					conf.ReportSummary.InstanceEfficiencyPercentages[i].Value = fixDot(strArr[5])
					i++
				}
			case "This table displays top 10 wait events by total wait time": // TODO This table displays top 5, 6, 10 wait events by total wait time
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1) // split line
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                                                                                                                            // if we can't select to next line
						continue
					}
					conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime = append(conf.ReportSummary.Top10ForegroundEventsByTotalWaitTime, TopForegroundEventsByTotalWaitTime{
						Event:         strArr[1],
						Waits:         fixDot(strArr[2]),
						TotalWaitTime: fixDot(strArr[3]),
						WaitAvg:       fixDot(strArr[4]),
						PerDBTime:     fixDot(strArr[5]),
						WaitClass:     strArr[6],
					})
				}
			case "This table displays wait class statistics ordered by total wait time":
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1) // split line
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                                                                                                                                          // if we can't select to next line
						continue
					}
					conf.ReportSummary.WaitClassesByTotalWaitTime = append(conf.ReportSummary.WaitClassesByTotalWaitTime, WaitClassesByTotalWaitTime{
						WaitClass:         strArr[1],
						Waits:             fixDot(strArr[2]),
						TotalWaitTime:     fixDot(strArr[3]),
						AvgWait:           fixDot(strArr[4]),
						PerDBTime:         fixDot(strArr[5]),
						AvgActiveSessions: fixDot(strArr[6]),
					})
				}
			case "This table displays system load statistics":
				strArr = regexp.MustCompile(`<tr><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
				conf.ReportSummary.HostCPU = make([]HostCPU, 1)                                                                                                                                                                                                                                                                                                                                                                                                    // everytime we have only one line
				if len(strArr) == 10 {
					conf.ReportSummary.HostCPU[0].CPUs = fixDot(strArr[1])
					conf.ReportSummary.HostCPU[0].Cores = fixDot(strArr[2])
					conf.ReportSummary.HostCPU[0].Sockets = fixDot(strArr[3])
					conf.ReportSummary.HostCPU[0].LABegin = fixDot(strArr[4])
					conf.ReportSummary.HostCPU[0].LAEnd = fixDot(strArr[5])
					conf.ReportSummary.HostCPU[0].PerUser = fixDot(strArr[6])
					conf.ReportSummary.HostCPU[0].PerSystem = fixDot(strArr[7])
					conf.ReportSummary.HostCPU[0].PerWIO = fixDot(strArr[8])
					conf.ReportSummary.HostCPU[0].PerIDLE = fixDot(strArr[9])
				}
			case "This table displays CPU usage and wait statistics":
				strArr = regexp.MustCompile(`<td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td><td align="right" class='awrc'>(.*?)</td></tr>`).FindStringSubmatch(iter) // select item from row
				conf.ReportSummary.InstanceCPU = make([]InstanceCPU, 1)                                                                                                                                  // everytime we have only one line
				if len(strArr) == 4 {
					conf.ReportSummary.InstanceCPU[0].PerTotalCPU = fixDot(strArr[1])
					conf.ReportSummary.InstanceCPU[0].PerBysuCPU = fixDot(strArr[2])
					conf.ReportSummary.InstanceCPU[0].PerDBTimeWaiting = fixDot(strArr[3])
				}
			case "This table displays IO profile":
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1) // split line
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                                                           // if we can't select to next line
						continue
					}
					conf.ReportSummary.IOProfile = append(conf.ReportSummary.IOProfile, IOProfile{
						Name:           strArr[1],
						RWPerSecond:    fixDot(strArr[2]),
						ReadPerSecond:  fixDot(strArr[3]),
						WritePerSecond: fixDot(strArr[4]),
					})
				}
			case "This table displays memory statistics":
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1) // split line
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                   // if we can't select to next line
						continue
					}
					conf.ReportSummary.MemoryStatistics = append(conf.ReportSummary.MemoryStatistics, MemoryStatistics{
						Name:  strArr[1],
						Begin: fixDot(strArr[2]),
						End:   fixDot(strArr[3]),
					})
				}
			case "This table displays cache sizes and other statistics for                     different types of cache":
				i = 0
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1) // split line
				conf.ReportSummary.CacheSizes = make([]CacheSizes, len(textBodyTwo)*2-3)             // -3 because last second item not contain information
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td>(<td class='\w+'>(.*?):</td><td align="right" class='\w+'>(.*?)</td>)*</tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                                                                                         // if we can't select to next line
						continue
					}
					conf.ReportSummary.CacheSizes[i].Name = strArr[1]
					conf.ReportSummary.CacheSizes[i].Begin = fixDot(strArr[2])
					conf.ReportSummary.CacheSizes[i].End = fixDot(strArr[3])

					if strArr[5] == "" { // last line
						break
					}

					i++
					conf.ReportSummary.CacheSizes[i].Name = strArr[5]
					conf.ReportSummary.CacheSizes[i].Begin = fixDot(strArr[6])
					i++
				}
			case "This table displays shared pool statistics":
				textBodyTwo = regexp.MustCompile(`<tr><td scope="row" class='\w+'>`).Split(iter, -1) // split line
				for _, val = range textBodyTwo {
					strArr = regexp.MustCompile(`(.*?):</td><td align="right" class='\w+'>(.*?)</td><td align="right" class='\w+'>(.*?)</td></tr>`).FindStringSubmatch(val) // select item from row
					if len(strArr) == 0 {                                                                                                                                   // if we can't select to next line
						continue
					}
					conf.ReportSummary.SharedPoolStatistics = append(conf.ReportSummary.SharedPoolStatistics, SharedPoolStatistics{
						Name:  strArr[1],
						Begin: fixDot(strArr[2]),
						End:   fixDot(strArr[3]),
					})
				}

			default:
				continue
			}
		}
	}
}

// create maps with element
func createMaps(textInput string, maps map[string]string) error {
	textBody := strings.Split(textInput, `<h3 class="awr">`)
	for _, text := range textBody {
		if reg, _ := regexp.MatchString(`(.*?)</h3>([\D|\d]*)`, string(text)); reg {
			s := regexp.MustCompile(`(.*?a>)*(.*?)</h3>([\D|\d]+)`).FindStringSubmatch(string(text))
			maps[s[2]] = s[3]
		}
	}
	if len(maps) == 0 {
		return errors.New("Not found elements map in the AWR.")
	}
	// firts text about server
	maps["Work Info"] = textBody[0]
	return nil
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {

	var str string
	data := PageData{}

	data.AttributeUploadFile = true

	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/upload.gtpl")
		t.Execute(w, data)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			log.Printf("ERROR: %s", err)
			t, _ := template.ParseFiles("template/upload.gtpl")
			data.AttributeUploadFile = false
			t.Execute(w, data)
			return
		}
		defer file.Close()

		// check for upload folder
		_, err = os.Stat("upload")
		if os.IsNotExist(err) {
			os.MkdirAll("upload", 0666)
		}

		str = "upload/" + handler.Filename
		f, err := os.OpenFile(str, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
		defer os.Remove("upload/" + handler.Filename) // delete file

		io.Copy(f, file)
		data.PageTitle = handler.Filename
		log.Printf("File %s upload.", handler.Filename)

		worker(str, &data)
		log.Printf("File %s is processed.", handler.Filename)

		// read config file
		// Config influxdb
		configuration := Config{}
		configuration.Init()

		// Create a new HTTPClient
		configuration.Client, err = client.NewHTTPClient(client.HTTPConfig{
			Addr:     configuration.UrlDB,
			Username: configuration.Username,
			Password: configuration.Password,
		})
		if err != nil {
			log.Fatal(err)
		}

		if configuration.Debug == "true" {
			configuration.DeleteDB()
			configuration.CreateDB()
		}

		SentDB(&configuration, &data)
		log.Printf("File %s upload to DB %s.", handler.Filename, configuration.NameDB)

		compare(&data, &configuration)
		//log.Printf("File %s processed.", handler.Filename)

		// work with html
		t := template.Must(template.ParseFiles("template/template.gtpl"))
		t.Execute(w, data) // merge.
		log.Printf("File %s printed.", handler.Filename)
	}
	data = PageData{} // TODO clear struct
}
func worker(filename string, dataStruct *PageData) {

	var attribute bool
	work := MainTable{}

	// read file
	// TODO check body file
	body, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// create map
	maps := make(map[string]string)
	err = createMaps(body, maps)
	if err != nil {
		log.Fatal(err)
	}
	// fill in struct
	parser(&work, maps)

	// TODO добавить информацию о системе - дата снятия метрик, информация о бд

	// search TABLE ACCESS - STORAGE FULL
	for _, sqlText := range work.TopSQLWithTopRowSources {
		if sqlText.RowSource == "TABLE ACCESS - STORAGE FULL" {
			for _, iter := range work.CompleteListOfSQLText {
				if iter.SQLID == sqlText.SQLID {
					dataStruct.ListSQLText = append(dataStruct.ListSQLText, ListSQLText{
						SQLId:       sqlText.SQLID,
						SQLDescribe: sqlText.RowSource,
						SQLText:     iter.SQLText,
					})
				}
			}
		}
	}
	for _, sqlText := range work.TopSQLWithTopEvents {
		if sqlText.RowSource == "TABLE ACCESS - STORAGE FULL" {
			attribute = true
			for _, iter := range dataStruct.ListSQLText { // if the second item
				if iter.SQLId == sqlText.SQLID {
					attribute = false
					break
				}
			}
			for _, iter := range work.CompleteListOfSQLText {
				if (iter.SQLID == sqlText.SQLID) && attribute {
					dataStruct.ListSQLText = append(dataStruct.ListSQLText, ListSQLText{
						SQLId:       sqlText.SQLID,
						SQLDescribe: sqlText.RowSource,
						SQLText:     iter.SQLText,
					})
				}
			}
		}
	}
	for _, sqlText := range work.CompleteListOfSQLText {
		attribute = true
		// more like 10
		if strings.Count(strings.ToLower(sqlText.SQLText), " like ") > 9 {
			for _, iter := range dataStruct.ListSQLText { // if the second item
				if iter.SQLId == sqlText.SQLID {
					attribute = false
					break
				}
			}
			if attribute {
				dataStruct.ListSQLText = append(dataStruct.ListSQLText, ListSQLText{
					SQLId:       sqlText.SQLID,
					SQLDescribe: "More like then 10.",
					SQLText:     sqlText.SQLText,
				})
				continue
			}
		}
		// search select * from
		if strings.Contains(strings.ToLower(sqlText.SQLText), "select * from ") {
			for _, iter := range dataStruct.ListSQLText { // if the second item
				if iter.SQLId == sqlText.SQLID {
					attribute = false
					break
				}
			}
			if attribute {
				dataStruct.ListSQLText = append(dataStruct.ListSQLText, ListSQLText{
					SQLId:       sqlText.SQLID,
					SQLDescribe: `Use: "Select * from"`,
					SQLText:     sqlText.SQLText,
				})
				continue
			}
		}
	}
	// Instance Efficiency Percentages (Target 100%)
	for _, iter := range work.ReportSummary.InstanceEfficiencyPercentages {
		if iter.Name == "% Non-Parse CPU" {
			if iter.Value >= 90 {
				dataStruct.NonParseCPU = fmt.Sprintf("Большинство ресурсов ЦП %v процентов используется в различных операциях IO, почти отсутсвует парсинг(hard, soft, soft cursor cache hit), что говорит о правильной работе базы данных.", iter.Value)
			} else {
				dataStruct.NonParseCPU = fmt.Sprintf("Большинство ресурсов ЦП %v процентов тратится на парстинг.", iter.Value)
			}
			continue
		}
		if iter.Name == "Parse CPU to Parse Elapsd %" {
			if iter.Value >= 90 {
				dataStruct.ParseCPUElapsd = fmt.Sprintf("ЦП %v процентов не ожидает ресурсов, что говорит о правильной работе базы данных.", iter.Value)
			} else {
				dataStruct.ParseCPUElapsd = fmt.Sprintf("Большинство ресурсов ЦП %v процентов тратится на ожидание ресурсов.", iter.Value)
			}
			continue
		}
		if iter.Name == "Soft Parse %" {
			dataStruct.SoftParse = fmt.Sprintf(`Вы используйте Soft Parse на уровне %v процентов. Если же вы делаете один Hard Parse, а затем последующие execute идут уже без парсинга, то данный показатель будет очень низкий.`, iter.Value)
			continue
		}
	}
	// Shared Pool Statistics
	for _, iter := range work.ReportSummary.SharedPoolStatistics {
		if iter.Name == "Memory Usage %" {
			if (iter.Begin >= 75) && (iter.End <= 90) {
				dataStruct.SharedPoolStatistics = fmt.Sprintf("Процент использование разделяемого пулан находится в рамках %v - %v процентов, что говорит о правильной работе базы данных", iter.Begin, iter.End)
				continue
			}
			if (iter.Begin < 75) || (iter.End < 75) {
				dataStruct.SharedPoolStatistics = fmt.Sprintf("Процент использования памяти слишком низкий - %v - %v процентов. Память тратится напрасно.", iter.Begin, iter.End)
				continue
			}
			if (iter.Begin > 90) || (iter.End > 90) {
				dataStruct.SharedPoolStatistics = fmt.Sprintf("Процент использования памяти слишком высокий - %v - %v процентов. Происходит вытеснение компонентов разделяемого пула как устаревшийх файлов, что приводит к жесткому разбору (hard parse) SQL-операторов при их повторном выполнении.", iter.Begin, iter.End)
				continue
			}
		}
	}
	// Top 10 Foreground Events by Total Wait Time
	for _, sqlText := range work.ReportSummary.Top10ForegroundEventsByTotalWaitTime {
		dataStruct.TopForegroundEventsByTotalWaitTime = append(dataStruct.TopForegroundEventsByTotalWaitTime, TopForegroundEventsByTotalWaitTime{
			Event:         sqlText.Event,
			Waits:         sqlText.Waits,
			TotalWaitTime: sqlText.TotalWaitTime,
			WaitAvg:       sqlText.WaitAvg,
			PerDBTime:     sqlText.PerDBTime,
			WaitClass:     sqlText.WaitClass,
		})
	}
	// Foreground Wait Class
	for _, sqlText := range work.WaitEventsStatistics.ForegroundWaitClass {
		dataStruct.WaitEventsStatistics.ForegroundWaitClass = append(dataStruct.WaitEventsStatistics.ForegroundWaitClass, ForegroundWaitClass{
			WaitClass:     sqlText.WaitClass,
			Waits:         sqlText.Waits,
			PerTime:       sqlText.PerTime,
			TotalWaitTime: sqlText.TotalWaitTime,
			AvgWait:       sqlText.AvgWait,
			PerDBTime:     sqlText.PerDBTime,
		})
	}
	// Foreground Wait Events
	for _, sqlText := range work.WaitEventsStatistics.ForegroundWaitEvents {
		dataStruct.WaitEventsStatistics.ForegroundWaitEvents = append(dataStruct.WaitEventsStatistics.ForegroundWaitEvents, ForegroundWaitEvents{
			Event:         sqlText.Event,
			Waits:         sqlText.Waits,
			PerTime:       sqlText.PerTime,
			TotalWaitTime: sqlText.TotalWaitTime,
			AvgWait:       sqlText.AvgWait,
			WaitsTxn:      sqlText.WaitsTxn,
			PerDBTime:     sqlText.PerDBTime,
		})
	}
	// Background Wait Events
	for _, sqlText := range work.WaitEventsStatistics.BackgroundWaitEvents {
		dataStruct.WaitEventsStatistics.BackgroundWaitEvents = append(dataStruct.WaitEventsStatistics.BackgroundWaitEvents, BackgroundWaitEvents{
			Event:         sqlText.Event,
			Waits:         sqlText.Waits,
			PerTime:       sqlText.PerTime,
			TotalWaitTime: sqlText.TotalWaitTime,
			AvgWait:       sqlText.AvgWait,
			WaitsTxn:      sqlText.WaitsTxn,
			PerbgTime:     sqlText.PerbgTime,
		})
	}
	// Top SQL with Top Row Sources
	for _, sqlText := range work.TopSQLWithTopRowSources {
		dataStruct.TopSQLWithTopRowSources = append(dataStruct.TopSQLWithTopRowSources, TopSQLWithTopRowSources{
			SQLID:        sqlText.SQLID,
			PlanHash:     sqlText.PlanHash,
			Executions:   sqlText.Executions,
			Activity:     sqlText.Activity,
			RowSource:    sqlText.RowSource,
			RowSourcePer: sqlText.RowSourcePer,
			TopEvent:     sqlText.TopEvent,
			EventPer:     sqlText.EventPer,
			SQLText:      sqlText.SQLText,
		})
	}
	// Top SQL with Top Events
	for _, sqlText := range work.TopSQLWithTopEvents {
		dataStruct.TopSQLWithTopEvents = append(dataStruct.TopSQLWithTopEvents, TopSQLWithTopEvents{
			SQLID:        sqlText.SQLID,
			PlanHash:     sqlText.PlanHash,
			Executions:   sqlText.Executions,
			Activity:     sqlText.Activity,
			RowSource:    sqlText.RowSource,
			RowSourcePer: sqlText.RowSourcePer,
			Event:        sqlText.Event,
			EventPer:     sqlText.EventPer,
			SQLText:      sqlText.SQLText,
		})
	}
	// SQL ordered by Elapsed Time	SQLOrderByElapsedTime
	for _, sqlText := range work.SQLOrderByElapsedTime {
		dataStruct.SQLOrderByElapsedTime = append(dataStruct.SQLOrderByElapsedTime, SQLOrderByElapsedTime{
			ElapsedTime:        sqlText.ElapsedTime,
			Executions:         sqlText.Executions,
			ElapsedTimePerExec: sqlText.ElapsedTimePerExec,
			Total:              sqlText.Total,
			CPU:                sqlText.CPU,
			IO:                 sqlText.IO,
			SQLID:              sqlText.SQLID,
			SQLModule:          sqlText.SQLModule,
			SQLText:            sqlText.SQLText,
		})
	}
	// SQL ordered by CPU Time	SQLOrderedByCPUTime
	for _, sqlText := range work.SQLOrderedByCPUTime {
		dataStruct.SQLOrderedByCPUTime = append(dataStruct.SQLOrderedByCPUTime, SQLOrderedByCPUTime{
			CPUTime:     sqlText.CPUTime,
			Executions:  sqlText.Executions,
			CPUPerExec:  sqlText.CPUPerExec,
			Total:       sqlText.Total,
			ElapsedTime: sqlText.ElapsedTime,
			CPU:         sqlText.CPU,
			IO:          sqlText.IO,
			SQLID:       sqlText.SQLID,
			SQLModule:   sqlText.SQLModule,
			SQLText:     sqlText.SQLText,
		})
	}
	// SQL ordered by Gets
	for _, sqlText := range work.SQLOrderedByGets {
		fmt.Println("sfdsf")
		fmt.Println(sqlText.BufferGets)
		dataStruct.SQLOrderedByGets = append(dataStruct.SQLOrderedByGets, SQLOrderedByGets{
			BufferGets:  sqlText.BufferGets,
			Executions:  sqlText.Executions,
			GetsPerExec: sqlText.GetsPerExec,
			Total:       sqlText.Total,
			ElapsedTime: sqlText.ElapsedTime,
			CPU:         sqlText.CPU,
			IO:          sqlText.IO,
			SQLID:       sqlText.SQLID,
			SQLModule:   sqlText.SQLModule,
			SQLText:     sqlText.SQLText,
		})
	}
	// This table displays database instance information
	for _, sqlText := range work.WorkInfo.WIDatabaseInstanceInformation {
		dataStruct.WorkInfo.WIDatabaseInstanceInformation = append(dataStruct.WorkInfo.WIDatabaseInstanceInformation, WIDatabaseInstanceInformation{
			DBName:      sqlText.DBName,
			DBId:        sqlText.DBId,
			Instance:    sqlText.Instance,
			Instnum:     sqlText.Instnum,
			StartupTime: sqlText.StartupTime,
			Release:     sqlText.Release,
			RAC:         sqlText.RAC,
		})
	}
	// This table displays host information
	for _, sqlText := range work.WorkInfo.WIHostInformation {
		dataStruct.WorkInfo.WIHostInformation = append(dataStruct.WorkInfo.WIHostInformation, WIHostInformation{
			HostName: sqlText.HostName,
			Platform: sqlText.Platform,
			CPUs:     sqlText.CPUs,
			Cores:    sqlText.Cores,
			Sockets:  sqlText.Sockets,
			MemoryGB: sqlText.MemoryGB,
		})
	}
	// This table displays snapshot information
	for _, sqlText := range work.WorkInfo.WISnapshotInformation {
		dataStruct.WorkInfo.WISnapshotInformation = append(dataStruct.WorkInfo.WISnapshotInformation, WISnapshotInformation{
			Name:           sqlText.Name,
			SnapId:         sqlText.SnapId,
			SnapTime:       sqlText.SnapTime,
			Sessions:       sqlText.Sessions,
			CursorsSession: sqlText.CursorsSession,
		})
	}

}
func main() {

	// configurator for logger
	var str = "log" // name folder for logs

	// check what folder log is exist
	_, err := os.Stat(str)
	if os.IsNotExist(err) {
		os.MkdirAll(str, 0666)
	}
	str = fmt.Sprintf("%s/%d-%02d-%02d-%02d-%02d-%02d-logFile.log", str, time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	// open a file
	f, err := os.OpenFile(str, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f) // TODO config logs
	log.SetPrefix("AWRcompar ")

	// start server
	http.HandleFunc("/", upload) // setting router rule

	port := flag.String("port", "9090", "the port value")
	flag.Parse()

	log.Printf("Start work. Port: %s", *port)
	defer log.Println("Stop work.")

	err = http.ListenAndServe(":"+*port, nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
