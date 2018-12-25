<!DOCTYPE html>
<html>
<head>
    <title>Something {{.PageTitle}}</title>
</head>
<body><a name="top"></a>
<h2>Общая информация</h2>
<h3>Информация о системе</h3>
<table class="table" border=1 bgcolor="#ffffff">
  <thead>
   <tr>
      <th>Имя БД</th>
      <th>ID БД</th>
      <th>Экземпляр БД/th>
      <th>Номер экземпляра</th>
      <th>Время запуска экземпляра БД</th>
      <th>Релиз</th>
      <th>Кластер(RAC)</th>
   </tr>
  </thead>
  <tbody>
     {{range .WorkInfo.WIDatabaseInstanceInformation}}
        <tr>
            <td>{{.DBName}}</td>
            <td>{{.DBId}}</td>
            <td>{{.Instance}}</td>
            <td>{{.Instnum}}</td>
            <td>{{.StartupTime}}</td>
            <td>{{.Release}}</td>
            <td>{{.RAC}}</td>
        </tr>
      {{end}}
  </tbody>
</table>
<br />
<table class="table" border=1 bgcolor="#fffff0">
  <thead>
   <tr>
      <th>Имя сервера</th>
      <th>ОС сервера</th>
      <th>Кол-во процессоров</th>
      <th>Кол-во ядер</th>
      <th>Гнезд</th>
      <th>Объём ОП, ГБ</th>
   </tr>
  </thead>
  <tbody>
     {{range .WorkInfo.WIHostInformation}}
        <tr>
            <td>{{.HostName}}</td>
            <td>{{.Platform}}</td>
            <td>{{.CPUs}}</td>
            <td>{{.Cores}}</td>
            <td>{{.Sockets}}</td>
            <td>{{.MemoryGB}}</td>
        </tr>
      {{end}}
  </tbody>
</table>
<br />
<table class="table" border=1 bgcolor="#fffff5">
  <thead>
   <tr>
      <th>Параметр</th>
      <th>ID snapshot</th>
      <th>Время снятия snapshot</th>
      <th>Сессии</th>
      <th>Курсоры/сессии</th>
   </tr>
  </thead>
  <tbody>
     {{range .WorkInfo.WISnapshotInformation}}
        <tr>
            <td>{{.Name}}</td>
            <td>{{.SnapId}}</td>
            <td>{{.SnapTime}}</td>
            <td>{{.Sessions}}</td>
            <td>{{.CursorsSession}}</td>
        </tr>
      {{end}}
  </tbody>
</table>
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
<p>Если вы видите в топе <b>log file sync</b>, то не надо сразу его бежать оптимизировать в руководстве Oracle Reference Manual, сказанно: "Когда пользовательский сеанс фиксирует транзакцию, информация повторного выполнения должна быть сброшена в файл журнала повторного выполнения. Пользовательский сеанс выдает задание процессу LGWR на запись буфера журнала повторного выполнения в файл журнала. Когда процесс LGWR завершит запись, он уведомляет об этом пользовательский сеанс. Wait Time: время ожидания включает время записи буфера журнала и время уведомления."</p>
<p>Теперь, когда понятно, чего именно пришлось ждать, можно придумать, как от этого ожидания избавиться. Когда ожидается синхронизация файла журнала, надо настраивать работу процесса LGWR. Чтобы уменьшить время ожидания можно использовать более быстрые диски, генерировать меньше информации повторного выполнения, снизить конфликты доступа к дискам, содержащим журналы, и т.д. Найти причину ожидания — одно дело, устранить ее — совсем другое. В Oracle измеряется время ожидания более 200 событий, причем ни для одного из них нет простого способа сократить время ожидания. </p>
<p>Не стоит забывать, что ждать чего-нибудь придется всегда. Если устранить одно препятствие, появится другое. Нельзя вообще избавиться от длительного ожидания событий — всегда придется чего-то ждать. Настройка "для максимально быстрой работы" может продолжаться бесконечно. Всегда можно сделать так, чтобы скорость работы возросла на один процент, но время, которое необходимо затратить на обеспечение каждого последующего процента прироста производительности, растет экспоненциально. Настройкой надо заниматься при наличии конкретной конечной цели. Если нельзя сказать, что настройка закончена, если достигнут показатель X, где X можно измерить, значит, вы напрасно тратите время.</p>
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
<p><a href="#top">Наверх</a></p>

<p><b><a href="#ForegroundWaitClass">Foreground Wait Class</a></b> и <b><a href="#ForegroundWaitEvents">Foreground Wait Events</a></b> показывают классы, которые провели в ожидании большего всего и список всех клиентов, которые также ожидали. Данный раздел является более подробных продолжение предыдущего и как провело его можно игнорировать если только вы не занимаетесь тонкой настройкой кластера т.к. например, тоже ожидание <b> SQL*Net message from client</b> показывает время в рамках, которого клиент не обращался к базе данных с запросами.</p>
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
<p><a href="#top">Наверх</a></p>

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
<p><a href="#top">Наверх</a></p>

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
<p><b><a href="#SQLOrderedByGets">SQL ordered by Gets</a></b> в этом разделе представленные запросы к БД упорядоченные по убыванию логических операций ввода/ввыода. При анализе стоит учитывать, что для PL/SQL процедур их количество прочитанных Buffer Gets будет состоять из суммы всех запросов в рамках данной процедуры. </p>
<p><b>SQL ordered by Reads</b> данный раздел является схожим с предыдущим, в нём указываются все операции ввода/вывода наиболее активно физически считывающие данные с жёсткого диска. Именно на эти запросы и процессы надо обратить внимание, если система не справляется с объемом ввода/вывода. </p>
<p><b>SQL ordered by Executions</b> наиболее часто выполняемы запросы.</p>
<p><b>SQL ordered by Version Count</b> показано количество SQL-операторов экземпляров одного и того же оператора в разделяемом пуле. Появление дублей обусловлено: 1. Под разными пользователями выполняли один и тот же SQL-оператор, но обращался он к разным при этом таблицам. 2. Запрос исполнялся в другой среде. 3. Используется механизм тщательного контроля доступа (Fine Grained Access Control). 4.Клиент использует связываемые переменные разных типов или размеров: одна программа связывает запрос с текстовой строкой длиной 10 символов, а другая — со строкой длиной 20 символов. В результате тоже получается новая версия SQLоператора. </p>
<p><b>Buffer Pool Statistics</b> Если используется поддержка нескольких буферных пулов, в этом разделе представляются данные по каждому из них. В нашем случае просто повторяется общая информация, представленная в начале отчета. </p>

<p><a name="SQLOrderedByGets">SQL ordered by Gets</a></p>
<table class="table" border=1 bgcolor="#02b12f">
  <thead>
   <tr>
      <th>Buffer Gets </th>
      <th>Executions</th>
      <th>Gets per Exec</th>
      <th>%Total</th>
      <th>Elapsed Time (s)</th>
      <th>%CPU</th>
      <th>%IO</th>
      <th>SQL Id</th>
      <th>SQL Module</th>
      <th>SQL Text</th>
   </tr>
  </thead>
  <tbody>
     {{range .SQLOrderedByGets}}
        <tr>
            <td>{{.BufferGets}}</td>
            <td>{{.Executions}}</td>
            <td>{{.GetsPerExec}}</td>
            <td>{{.Total}}</td>
            <td>{{.ElapsedTime}}</td>
            <td>{{.Cpu}}</td>
            <td>{{.IO}}</td>
            <td><a href="#{{.SQLID}}">{{.SQLID}}</a></td>
            <td>{{.SQLModule}}</td>
            <td>{{.SQLText}}</td>
        </tr>
      {{end}}
  </tbody>
</table>

<p><a href="#top">Наверх</a></p>

<h2>Список SQLID тяжелых запросов, которые следует оптимизировать</h2>
<p><b><a href="#TopSQLWithTopEvents">Top SQL with Top Events</a></b> — топ SQL запросов на которые приходится наибольший процент активности сесси и  больше всего ожидающие события порожденные этими операторами SQL. В столбце Sampled # of Executions показано, сколько выборочных исполнений конкретной инструкции SQL было выбрано. </p>
<p><b><a href="#TopSQLWithTopRowSources">Top SQL with Top Row Sources</a></b> — топ SQL запросов на которые приходится наибольший процент выборочной активности сеанса и их подробная информация о плане выполнения. Вы можете использовать эту информацию, чтобы определить, какая часть выполнения SQL операторов значительно повлияла на затраченное время SQL оператора. </p>
<p><b><a href="#SQLOrderByElapsedTime">SQL ordered by Elapsed Time</a></b> — топ SQL запросов по затраченному времени на их выполнение. Следует на них уделить большее внимание. </p>
<p><b><a href="#SQLOrderedByCPUTime">SQL ordered by CPU Time</a></b> — топ Sql запросов по процессорному времени. Следует на них уделить большее внимание.</p>
<p>Список того что нашли и то что было</p>
<ul>
    {{range .ListSQLText}}
        <li><a href="#{{.SQLId}}">{{.SQLId}}</a> — {{.SQLDescribe}}, <i>{{.TextUI}} </i>;</li>
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

<p><a name="SQLOrderByElapsedTime">SQL ordered by Elapsed Time</a></p>
<table class="table" border=1 bgcolor="#00bff0">
  <thead>
   <tr>
      <th>Elapsed Time (s)</th>
      <th>Executions</th>
      <th>Elapsed Time per Exec (s)</th>
      <th>% Total</th>
      <th>% CPU</th>
      <th>% IO</th>
      <th>SQL Id</th>
      <th>SQL Module</th>
      <th>SQL Text</th>
   </tr>
  </thead>
  <tbody>
     {{range .SQLOrderByElapsedTime}}
        <tr>
            <td>{{.ElapsedTime}}</td>
            <td>{{.Executions}}</td>
            <td>{{.ElapsedTimePerExec}}</td>
            <td>{{.Total}}</td>
            <td>{{.Cpu}}</td>
            <td>{{.IO}}</td>
            <td><a href="#{{.SQLID}}">{{.SQLID}}</a></td>
            <td>{{.SQLModule}}</td>
            <td>{{.SQLText}}</td>
        </tr>
      {{end}}
  </tbody>
</table>
<p><a href="#top">Наверх</a></p>

<p><a name="SQLOrderedByCPUTime">SQL ordered by CPU Time</a></p>
<table class="table" border=1 bgcolor="#00bf0f">
  <thead>
   <tr>
      <th>CPU Time (s)</th>
      <th>Executions</th>
      <th>CPU per Exec (s)</th>
      <th>% Total</th>
      <th>Elapsed Time (s)</th>
      <th>% CPU</th>
      <th>% IO</th>
      <th>SQL Id</th>
      <th>SQL Module</th>
      <th>SQL Text</th>
   </tr>
  </thead>
  <tbody>
     {{range .SQLOrderedByCPUTime}}
        <tr>
            <td>{{.CPUTime}}</td>
            <td>{{.Executions}}</td>
            <td>{{.CPUPerExec}}</td>
            <td>{{.Total}}</td>
            <td>{{.ElapsedTime}}</td>
            <td>{{.CPU}}</td>
            <td>{{.IO}}</td>
            <td><a href="#{{.SQLID}}">{{.SQLID}}</a></td>
            <td>{{.SQLModule}}</td>
            <td>{{.SQLText}}</td>

        </tr>
      {{end}}
  </tbody>
</table>
<p><a href="#top">Наверх</a></p>

<p><a name="TopSQLWithTopEvents">Top SQL with Top Events</a></p>
<table class="table" border=1 bgcolor="#00bff4">
  <thead>
   <tr>
      <th>SQL ID</th>
      <th>Plan Hash</th>
      <th>Executions</th>
      <th>% Activity</th>
      <th>Event</th>
      <th>% Event</th>
      <th>Top Row Source</th>
      <th>% Row Source</th>
      <th>SQL Text</th>
   </tr>
  </thead>
  <tbody>
     {{range .TopSQLWithTopEvents}}
        <tr>
            <td><a href="#{{.SQLID}}">{{.SQLID}}</a></td>
            <td>{{.PlanHash}}</td>
            <td>{{.Executions}}</td>
            <td>{{.Activity}}</td>
            <td>{{.Event}}</td>
            <td>{{.EventPer}}</td>
            <td>{{.RowSource}}</td>
            <td>{{.RowSourcePer}}</td>
            <td>{{.SQLText}}</td>
        </tr>
      {{end}}
  </tbody>
</table>
<p><a href="#top">Наверх</a></p>

<p><a name="TopSQLWithTopRowSources">Top SQL with Top Row Sources</a></p>
<table class="table" border=1 bgcolor="#00bfff">
  <thead>
   <tr>
      <th>SQL ID</th>
      <th>Plan Hash</th>
      <th>Executions</th>
      <th>% Activity</th>
      <th>Row Source</th>
      <th>% Row Source</th>
      <th>Top Event</th>
      <th>% Event</th>
      <th>SQL Text</th>
   </tr>
  </thead>
  <tbody>
     {{range .TopSQLWithTopRowSources}}
        <tr>
            <td><a href="#{{.SQLID}}">{{.SQLID}}</a></td>
            <td>{{.PlanHash}}</td>
            <td>{{.Executions}}</td>
            <td>{{.Activity}}</td>
            <td>{{.RowSource}}</td>
            <td>{{.RowSourcePer}}</td>
            <td>{{.TopEvent}}</td>
            <td>{{.EventPer}}</td>
            <td>{{.SQLText}}</td>
        </tr>
      {{end}}
  </tbody>
</table>
<p><a href="#top">Наверх</a></p>

<p><a href="/">Back to the main page.</a></p>
</body>
</html>