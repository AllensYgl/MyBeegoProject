<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="/static/css/mycss/login.css" />
</head>
<body>
        时间:{{.time}}
    <div>
        <div class="center">
            <form method="POST" name="form" action="/changeMail"> 
                <!-- {{ .xsrfdata }} -->
                <h1>用户名:{{.UserName}}</h1><br>
                <div class="m">
                        <h1 class="mail">邮&nbsp;&nbsp;&nbsp;箱:</h1>
                        <input type="text" name="Mail" value="{{.Mail}}" class="Mail">
                </div>
                <hr>
                <input type="submit" value="提交" class="submit">
            </form>
        </div>
    </div>
</body>
</html>