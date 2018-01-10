<html>
<head>
       <title>Upload file AWR</title>
</head>
<body>
{{if .Attribute}}
    <p>Форма для загрузки файлов AWR.</p>
{{else}}
    <p>Вы не загрузили AWR файл. Форма для загрузки файлов AWR.</p>
{{end}}

<form enctype="multipart/form-data" action="/" method="post">
    <input type="file" name="uploadfile" />
    <input type="submit" value="Загрузить" />
</form>



</body>
</html>