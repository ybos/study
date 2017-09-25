<!DOCTYPE html>
<html>
<head>
	<title>登陆</title>
</head>
<body>
	<form action="" method="post">
		<input type="hidden" name="token" value="{{.}}" />
		<label>Username: <input type="text" name="username"></label>
		<label>Password: <input type="password" name="password"></label>
		<button type="submit">Submit</button>
	</form>
</body>
</html>