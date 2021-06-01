<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Login</title>
    <link rel="stylesheet" type="text/css" href="/static/css/style.css"/>
    <script src="/static/js/jquery-1.9.1.min.js"></script>
</head>
<body>
    <div id="login">
        <h1>Login</h1>
        <form method="post">
            <input type="text" id="name" required="required" placeholder="用户名" name="name"></input>
            <input type="password" id="pass" required="required" placeholder="密码" name="pass"></input>
            <button class="but" type="button" name="button" onclick="sub()">登录</button>
        </form>
    </div>
<script>
  function sub(){
    console.log("ok");
    $.ajax({
        type:"post",
        url:"/login",
        datatype: "json",
        data:{
          name:$("#name").val(),
          pass:$("#pass").val()
        },
        success:function(msg) {
          console.log(msg);
        }
    })
  }
</script>
</body>
</html>