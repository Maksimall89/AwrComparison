# AwrComparison

Программа для анализа AWR Oracle 12. Она анализирует AWR на поиск "не оптимальных" запросов. 
Все SQL ID записываются в InfluxDB. Конфигурация InfluxDB описана в `config.json`. 

Вначале стандартно выбираем нужную нам AWR и загружаем её:

![Upload AWR](https://github.com/Maksimall89/AwrComparison/blob/master/doc/awr_upload.jpg)

В итоге мы получаем вот такой результат:
![Result](https://github.com/Maksimall89/AwrComparison/blob/master/doc/bad.jpg)

Программа ищет не оптимальные запросы (множество like, full scan, select * from и т.п. ), так же даёт общие рекомендации по оптимизации Oracle.

# Getting started
Get the source:

`go get github.com/maksimall89/AwrComparison/...`

Compile:

`go build -o AwrComparison.exe`

Start:

`http://127.0.0.1:9090/`