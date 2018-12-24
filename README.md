# AwrComparison

Программа для анализа AWR Oracle 12. Она анализирует AWR на поиск "не оптимальных" запросов. 
Все SQL ID записываются в InfluxDB. Конфигурация InfluxDB описана в `config.json`. 


# Getting started
Get the source:

`go get github.com/maksimall89/AwrComparison/...`

Compile:

`go build -o AwrComparison.exe`

Start:

`http://127.0.0.1:9090/`

