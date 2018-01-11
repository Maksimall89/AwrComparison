<!DOCTYPE html>
<html>
<head>
       <title>Upload file AWR</title>
</head>
<body>
{{if .AttributeUploadFile}}
    <p>Форма для загрузки файлов AWR.</p>
{{else}}
    <p><b>Вы не загрузили AWR файл.</b> <br/> Форма для загрузки файлов AWR.</p>
{{end}}

<form enctype="multipart/form-data" action="/" method="post">
    <input type="file" name="uploadfile" />
    <input type="submit" value="Загрузить" />
</form>
</body>
</html>