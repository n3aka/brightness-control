package main

import (
	"os"
	"text/template"
)

type dataCRUD struct {
	Sensor uint
	Hue    uint
}

var (
	j uint
	i uint
)

func main() {

	t := template.Must(template.New("models").Parse(head))
	x := dataCRUD{}
	f, _ := os.Create("gen.go")
	defer f.Close()
	t.Execute(f, x)
	j = 5
	for i = 700; i < 9673; i = i + 6 {
		d := dataCRUD{Sensor: j, Hue: i}
		t := template.Must(template.New("models").Parse(queue))
		t.Execute(f, d)
		j = j + 1
	}
	t = template.Must(template.New("models").Parse(end))
	t.Execute(f, x)
}

var head = `package main
import(
	"fmt"
)
func f(i uint32) ([]byte, error) {
	switch i{
`

var queue = `	case {{.Hue}}: 
		return []byte("{{.Sensor}}"), nil
		break
`
var end = `}
return []byte("0"), fmt.Errorf("something")
}`
