table2excel
===========
For excel haters like me. WIP

### Notes
```go
// You can edit this code!
// Click here and start typing.
package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"abc", 23}
	tmpl, _ := template.New("test").Parse(`
	<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<title>Go Web</title>
	</head>
	<body>
		Name: {{.Name}} Age: {{.Age}}
	</body>
</html>
`)
	_ = tmpl.Execute(os.Stdout, p)

}
```
