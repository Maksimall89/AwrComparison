<html>
<head>
    <title>Something here {{.PageTitle}}</title>
</head>
<body>
<h1>{{.PageTitle}}<h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
<h2><a href="/">Back to the main page.</a></h2>

</body>
</html>