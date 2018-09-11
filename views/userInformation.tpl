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
                <h1>UserName：{{.UserName}}</h1>
                    <br>
                <h1>Email:{{.Email}}</h1>
                <button onclick="javascript:window.location.href='/changeMail'" class="submit">修改</button>
                <button onclick="javascript:window.location.href='/changePwd'" class="submit">重置密码</button>
                <script>
                        if ("{{.Role}}"==1){
                            document.write("<button onclick=\"javascript:window.location.href='/admin'\" class=\"submit\">返回</button>")
                        }
                </script>
                <button onclick="javascript:window.location.href='/exit'" class="submit">退出</button>
        </div>
    </div>
    
</body>
</html>