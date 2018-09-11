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
                <div class="conts">
                        <form method="POST" name="form" action="/changePwd"> 
                            <h1 class="mail">&nbsp;&nbsp;&nbsp;原密码:</h1><input type="password" class="Mail" name="Pwd"><br>
                            <h1 class="mail">&nbsp;&nbsp;&nbsp;新密码:</h1><input type="password" class="Mail" name="NewPwd"><br>
                            <h1 class="mail">再次输入:</h1><input type="password" class="Mail"><br>
                             <input type="submit" value="修改" class="submit">
                         </form>
                </div>
        </div>
    </div>
</body>
</html>