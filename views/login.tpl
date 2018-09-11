<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>登陆界面</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="/static/css/mycss/login.css" />
</head>
<body>
    <div>
        <div class="center">
                <form method="POST" name="form" action="/index"> 
                    <!-- {{ .xsrfdata }} -->
                    <h1>用户登录</h1><hr>
                    用户名:<input type="text" name="UserName"><br>
                    密&nbsp;&nbsp;&nbsp;码:<input type="password" name="Pwd"><hr>
                    <input type="submit" value="登录" class="submit">
                </form>
        </div>
    </div>
</body>
</html>
