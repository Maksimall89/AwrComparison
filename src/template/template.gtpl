<html>
<head>
    <title>Something here {{.PageTitle}}</title>
</head>
<body>
<a name="top"></a>
<ul>
    {{range .ListSQLText}}
        <li><a href="#{{.SQLId}}">{{.SQLId}}</a> - {{.SQLDescribe}};</li>
    {{end}}
</ul>
<table class="table" border=1>
<caption>Список запросов содержащих TABLE ACCESS - STORAGE FULL или запросы со множейством like или выборкой по всем столбцам с помощью "select *".</caption>
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

<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>

</body>
</html>