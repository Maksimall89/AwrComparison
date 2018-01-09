<html>
<head>
    <title>Something here {{.PageTitle}}</title>
</head>
<body>
<h1>{{.PageTitle}}<h1>
<table class="table" border=1>
<caption>List SQL Text</caption>
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
            <td>{{.SQLId}}</td>
            <td>{{.SQLDescribe}}</td>
            <td>{{.SQLText}}</td>
        </tr>
      {{end}}
  </tbody>
</table>


<h2><a href="/">Back to the main page.</a></h2>

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