<!DOCTYPE html>
<html>
<head>
    <title>Something here {{.PageTitle}}</title>
</head>
<body><a name="top"></a>
<h2>Общая информация</h2>
<h3>Информация о системе</h3>
<ul>
    <li></li>
</ul>
<h3>Instance Efficiency Percentages</h3>
<p>{{.NonParseCPU}}</p>
<p>{{.ParseCPUElapsd}}</p>
<p>{{.SoftParse}}</p>
<h3>Shared Pool Statistics</h3>
<p>{{.SharedPoolStatistics}}</p>

<p><a href="#top">Наверх</a></p>
<h2>Самые долгие запросы</h2>
<h3>Top 10 Foreground Events By Total Wait Time</h3>
<p>В данном разделе представлены события, замедляющие работу базы данных больше всех. Прежде чем начинать оптимизацию событий советую, вначале определится с названием события и понять стоят ли усилия на его оптимизацию потраченного времени. </p>
<p>Если вы видите в топе <b>log file sync</b>, то не надо сразу его бежать оптмиизировать в руководстве Oracle Reference Manual, сказанно: "Когда пользовательский сеанс фиксирует транзакцию, информация повторного выполнения должна быть сброшена в файл журнала повторного выполнения. Пользовательский сеанс выдает задание процессу LGWR на запись буфера журнала повторного выполнения в файл журнала. Когда процесс LGWR завершит запись, он уведомляетобэтомпользовательскийсеанс. Wait Time: время ожидания включает время записи буфера журнала и время уведомления."</p>
<p>Теперь, когда понятно, чего именно пришлось ждать, можно придумать, как от этого ожидания избавиться. Когда ожидается синхронизация файла журнала, надо настраивать работу процесса LGWR. Чтобы уменьшить время ожидания можно использовать более быстрые диски, генерировать меньше информации повторного выполнения, снизить конфликты доступа к дискам, содержащим журналы, и т.д. Найти причину ожидания — одно дело, устранить ее — совсем другое. В Oracle измеряется время ожидания более 200 событий, причем ни для одного из них нет простого способа сократить время ожидания.</p>
<p>Не стоит забывать, что ждать чего-нибудь придется всегда. Если устранить одно препятствие, появится другое. Нельзя вообще избавиться от длительного ожидания событий — всегда придетсячего-тождать. Настройка "для максимально быстрой работы" может продолжаться бесконечно. Всегда можно сделать так, чтобы скорость работы возросла на один процент, но время, которое необходимо затратить на обеспечение каждого последующего процента прироста производительности, растет экспоненциально. Настройкой надо заниматься при наличии конкретной конечной цели. Если нельзя сказать, что настройка закончена, если достигнут показатель X, где X можно измерить, значит, вы напрасно тратите время.</p>
<table class="table" border=1 bgcolor="#f1b888">
  <thead>
	<tr>
		<th>Event</th>
		<th>Waits</th>
		<th>Total Wait Time (sec)</th>
		<th>Wait Avg(ms)</th>
		<th>% DB time</th>
		<th>Wait Class</th>
	</tr>
  </thead>
  <tbody>
     {{range .TopForegroundEventsByTotalWaitTime}}
        <tr>
            <td><a name="{{.Event}}"></a>{{.Event}}</td>
            <td>{{.Waits}}</td>
            <td>{{.TotalWaitTime}}</td>
            <td>{{.WaitAvg}}</td>
            <td>{{.PerDBTime}}</td>
            <td>{{.WaitClass}}</td>
        </tr>
      {{end}}
  </tbody>
</table>

<p><b><a href="#ForegroundWaitClass">Foreground Wait Class</a></b> и <b><a href="#ForegroundWaitEvents">Foreground Wait Events</a></b> показывают классы, которые провели в ожидании большего всего и список всех клиентов, которые также ожидали. Данный раздел является более подробных продолжение предыдущего и как провело его можно игнорировать если только вы не занимаетесь тонкой настройкой кластера т.к. например тоже ожидание <b> SQL*Net message from client</b> показывает время в рамках, которого клиент не обращался к базе данных с запросами.</p>
<p><b><a href="#BackgroundWaitEvents">Background Wait Events</a></b> опять же показывает время простоя событий, но теперь фоновыми процессами (LGWR, DBWR и т.д.).</p>

<p><a name="ForegroundWaitClass">Foreground Wait Class</a></p>
<table class="table" border=1 bgcolor="#f1b822">
  <thead>
	<tr>
		<th>WaitClass</th>
		<th>Waits</th>
		<th>%Time -outs</th>
		<th>Total Wait Time (s)</th>
		<th>Avg wait (ms)</th>
		<th>%DB time</th>
	</tr>
  </thead>
  <tbody>
     {{range .WaitEventsStatistics.ForegroundWaitClass}}
        <tr>
            <td>{{.WaitClass}}</td>
            <td>{{.Waits}}</td>
            <td>{{.PerTime}}</td>
            <td>{{.TotalWaitTime}}</td>
            <td>{{.AvgWait}}</td>
            <td>{{.PerDBTime}}</td>
        </tr>
      {{end}}
  </tbody>
</table>

<p><a name="ForegroundWaitEvents">Foreground Wait Events</a></p>
<table class="table" border=1 bgcolor="#f2b82f">
  <thead>
	<tr>
		<th>Event</th>
		<th>Waits</th>
		<th>%Time -outs</th>
		<th>Total Wait Time (s)</th>
		<th>Avg wait (ms)</th>
		<th>Waits /txn</th>
		<th>%DB time</th>
	</tr>
  </thead>
  <tbody>
     {{range .WaitEventsStatistics.ForegroundWaitEvents}}
        <tr>
            <td>{{.Event}}</td>
            <td>{{.Waits}}</td>
            <td>{{.PerTime}}</td>
            <td>{{.TotalWaitTime}}</td>
            <td>{{.AvgWait}}</td>
            <td>{{.WaitsTxn}}</td>
            <td>{{.PerDBTime}}</td>
        </tr>
      {{end}}
  </tbody>
</table>

<p><a name="BackgroundWaitEvents">Background Wait Events</a></p>
<table class="table" border=1 bgcolor="#f2bf2f">
  <thead>
	<tr>
		<th>Event</th>
		<th>Waits</th>
		<th>%Time -outs</th>
		<th>Total Wait Time (s)</th>
		<th>Avg wait (ms)</th>
		<th>Waits /txn</th>
		<th>% bg time</th>
	</tr>
  </thead>
  <tbody>
     {{range .WaitEventsStatistics.BackgroundWaitEvents}}
        <tr>
            <td>{{.Event}}</td>
            <td>{{.Waits}}</td>
            <td>{{.PerTime}}</td>
            <td>{{.TotalWaitTime}}</td>
            <td>{{.AvgWait}}</td>
            <td>{{.WaitsTxn}}</td>
            <td>{{.PerbgTime}}</td>
        </tr>
      {{end}}
  </tbody>
</table>
<p><a href="#top">Наверх</a></p>

<h2>Работа с блоками</h2>
<p><b>SQL ordered by Gets</b> в этом разделе представленные запросы к БД упорядоченные по убыванию логических операций ввода/ввыода. При анализе стоит учитывать, что для PL/SQL процедур их количество прочитанных Buffer Gets будет состоять из суммых всех запросов в рамках данной процедуры.</p>
<p><b>SQL ordered by Reads</b> данный раздел является схожим с предыдущим, в нём указываются все операции ввода/вывода наиболее активно физически считывающие данные с жетского диска. Именно на эти запросы и процессы надо обратить внимание, если система не справляется с объемом ввода/вывода.</p>
<p><b>SQL ordered by Executions</b> наиболее часто выполняемы запросы.</p>
<p><b>SQL ordered by Version Count</b> показано количество SQL-операторов экземпляров одного и того же оператора в разделяемом пуле. Появление дублей обусловлено: 1. Под разными пользователями выполняли один и тот же SQL-оператор, но обращался он к разным при этом таблицам. 2. Запрос исполнялся в другой среде. 3. Используется механизм тщательного контроля доступа (Fine Grained Access Control). 4.Клиент использует связываемые переменные разных типов или размеров: одна программа связывает запрос с текстовой строкой длиной 10 символов, а другая — со строкой длиной 20 символов. В результате тоже получается новая версия SQLоператора.</p>
<p><b>Buffer Pool Statistics</b> Если используется поддержка нескольких буферных пулов, в этом разделе представляются данные по каждому из них. В нашем случае просто повторяется общая информация, представленная в начале отчета.</p>

<p><a href="#top">Наверх</a></p>
<h2>Список SQLID тяжелых запросов, которые следует оптимизировать</h2>
<p><b>Top SQL with Top Row Sources</b>TopSQLWithTopRowSources</p>
<p><b>Top SQL with Top Events</b>TopSQLWithTopEvents</p>
<p><b>SQL ordered by Elapsed Time</b>SQL ordered by Elapsed Time</p>
<p><b>SQL ordered by CPU Time</b>SQL ordered by CPU Time</p>
<ul>
    {{range .ListSQLText}}
        <li><a href="#{{.SQLId}}">{{.SQLId}}</a> — {{.SQLDescribe}};</li>
    {{end}}
</ul></p>
<p>Список запросов содержащих TABLE ACCESS - STORAGE FULL или запросы со множейством like или выборкой по всем столбцам с помощью "select * from".</p>
<table class="table" border=1 bgcolor="#71bc78">
  <thead>
	<tr>
		<th>SQLId</th>
		<th>SQLDescribe</th>
		<th>SQLText</th>
	</tr>
  </thead>
  <tbody>
     {{range .ListSQLText}}
        <tr>
            <td><a name="{{.SQLId}}"></a>{{.SQLId}}</td>
            <td>{{.SQLDescribe}}</td>
            <td>{{.SQLText}}</td>
        </tr>
      {{end}}
  </tbody>
</table>

<p><a href="#top">Наверх</a></p>
<p><a href="/">Back to the main page.</a></p>
</body>
</html>