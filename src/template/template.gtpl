<!DOCTYPE html>
<html>
<head>
    <title>Something here {{.PageTitle}}</title>
</head>
<body><a name="top"></a>

<p>{{.NonParseCPU}}</p>
<p>{{.ParseCPUElapsd}}</p>
<p>{{.SoftParse}}</p>
<p>{{.SharedPoolStatistics}}</p>
<p>{{.}}</p>


<p>Список SQLID тяжелых запросов:</p>
<ul>
    {{range .ListSQLText}}
        <li><a href="#{{.SQLId}}">{{.SQLId}}</a> — {{.SQLDescribe}};</li>
    {{end}}
</ul>
<p>Список запросов содержащих TABLE ACCESS - STORAGE FULL или запросы со множейством like или выборкой по всем столбцам с помощью "select *".</p>
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