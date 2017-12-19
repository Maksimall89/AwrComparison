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
)

const configFileName  = "config.json"

type Config struct {
	TelegramBotToken string
	OwnName          string
}
type MainTable struct {
	body 	string
	OrderByElapsedTime []OrderByElapsedTime
	CompleteListOfSQLText []CompleteListOfSQLText
}

// SQL ordered by Elapsed Time
type OrderByElapsedTime struct{
	ElapsedTime			float32
	Executions			float32
	ElapsedTimePerExec	float32
	Total				float32
	Cpu					float32
	IO					float32
	SQLID				string
	SQLModule			string
	SQLText				string
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
		body += scanner.Text() + "\n"
	}
	return body, nil
}
// TODO парсер лога и запись его в структуры
func (conf *MainTable) parser()  {
	//<a class="awr" name=".*?"><\/a>(.*?)<\/td><td class='awrc'>(.*?)<\/td>
	// reg, _ = regexp.MatchString(`<a class="awr" name=".*?"><\/a>(.*?)<\/td><td class='awrc'>(.*?)<\/td>`, string(body)) true
	// s := regexp.MustCompile(``<a class="awr" name=".*?"><\/a>(.*?)<\/td><td class='awrc'>(.*?)<\/td>``).FindStringSubmatch(string(body))

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
	file, err := os.Open("configFileName")
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
	//log.SetOutput(f) // TODO config logs
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
	textBody := strings.Split(body, `<h3 class="awr">`)
	maps := make(map[string]string)

	for _, text := range  textBody{
		if reg, _ := regexp.MatchString(`(.*?)<\/h3>([\D|\d]*)`, string(text)); reg {
			s := regexp.MustCompile(`(.*?a>)*(.*?)<\/h3>([\D|\d]+)`).FindStringSubmatch(string(text))
			maps[s[2]] = s[3]
		}
	}

	log.Println(maps["SQL ordered by Elapsed Time"])

	work.parser()
//	log.Println(str)

}
