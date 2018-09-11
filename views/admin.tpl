<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>用户管理</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="/static/css/mycss/login.css" />
    <link rel="stylesheet" type="text/css" media="screen" href="/static/css/mycss/delete.css" />
</head>
<body>
    时间:{{.time}}
<div>
    <div class="center">
        <form action="/admin" method="GET">
            <div>
                <input placeholder="UserName" type="text" name="UserName">
                <input placeholder="Id" type="text" name="Id">
                <input placeholder="Mail" type="text" name="Mail">
                <input placeholder="Role" type="text" name="Role">
                <input type="submit" value="搜索">
            </div>
        </form>
        <script>document.write("<table class=\"tabl\" border=\"1px\">"+"<tr>"+"<td>UserName</td>  <td>Id</td>  <td>Pwd</td>  <td>Email</td>  <td>Role</td>  <td>Delete</td>  "+"</tr>");</script>
        {{range $keys,$values:=.users}}
        <script>
            document.write("<tr>")
                document.write("<td>"+"{{$values.UserName}}"+"</td>")
                document.write("<td>"+"{{$values.Id}}"+"</td>")
                document.write("<td>"+"{{$values.Pwd}}"+"</td>")
                document.write("<td>"+"{{$values.Mail}}"+"</td>")
                document.write("<td>"+"{{$values.Role}}"+"</td>")
                document.write("<td>"+"<button class=\"delete\" onclick=\"javascript:if(confirm('确定要删除吗?'))window.location.href='/admin/delete/{{$values.Id}}'\">删除</button>"+"</td>")
            document.write("</tr>")
        </script>
        {{end}}
        <script>document.write("</table>")</script>
        <div>
            <button class="add" onclick="javascript:window.location.href='/admin/add'">添加</button>
            <button class="add" onclick="javascript:window.location.href='/userInformation'">查看信息</button>
        </div>
    </div>
</div> 
</body>
</html>