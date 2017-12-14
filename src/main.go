package main
/*
	Документация:
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
)
type Config struct {
	TelegramBotToken string
	OwnName          string
}

// TODO парсер лога и запись его в структуры
func parser()  {

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
	file, err := os.Open("config.json")

	if err != nil {
		log.Println(err)
		configuration.OwnName = "Maksimall89"
		configuration.TelegramBotToken = "3257"
	}else{
		defer file.Close()
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
	defer log.Println(recover())

	// assign it to the standard logger
	//log.SetOutput(f) // TODO config logs
	log.SetPrefix("AWRcompar ")

	// read config file
	configuration := Config{}
	configuration.init()

	log.Println("Start work.")




}
