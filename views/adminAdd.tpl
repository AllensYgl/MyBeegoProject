<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Page Title</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="/static/css/mycss/login.css" />
    <link rel="stylesheet" type="text/css" media="screen" href="/static/css/mycss/delete.css" />
    <link rel="stylesheet" type="text/css" media="screen" href="/static/css/mycss/add.css" />
</head>
<body>
        时间:{{.time}}
    <div>
        <div class="center">
            <form action="/admin/add" method="POST">
                <div>
                    用户名:&nbsp;&nbsp;&nbsp;<input placeholder="UserName" type="text" name="UserName" class="cont"><br>
                    用户邮箱:&nbsp;&nbsp;<input placeholder="Mail" type="text" name="Mail" class="cont"><br>
                    用户密码:&nbsp;&nbsp;<input placeholder="Pwd" type="text" name="Pwd" class="cont"><br>
                    用户权限:
                    <select class="cont" name="Role">
                        <option value="1">管理员</option>
                        <option value="2" selected>普通用户</option>
                    </select><br>
                    <input type="submit" value="添加">
                </div>
            </form>
        </div>
    </div>
</body>
</html>