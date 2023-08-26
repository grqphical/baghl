package baghl

import (
	"bytes"
	"fmt"
	"html/template"
	"mime"
	"os"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

// Main API
func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"HTMLResponse": HTMLResponse,
	"Redirect":     Redirect,
	"JSONResponse": JSONResponse,
	"FileResponse": FileResponse,
	"TextResponse": TextResponse,
}

func RenderTemplate(file string, data lua.LTable) string {
	t, err := template.ParseFiles("templates/base.tpl", "templates/"+file)
	HandleError(err)

	var tmpl bytes.Buffer
	var templateData = map[string]interface{}{}

	data.ForEach(func(l1, l2 lua.LValue) {
		templateData[l1.String()] = l2
	})

	err = t.Execute(&tmpl, templateData)
	HandleError(err)
	result := tmpl.String()

	return result
}

func HTMLResponse(L *lua.LState) int {
	header := L.ToTable(1)
	file := L.ToString(2)
	data := L.ToTable(3)
	statusCode := L.ToNumber(4)

	body := RenderTemplate(file, *data)

	table := L.NewTable()

	table.RawSetString("header", header)
	table.RawSetString("body", lua.LString(body))
	table.RawSetString("statusCode", statusCode)

	L.Push(table)

	return 1
}

func Redirect(L *lua.LState) int {
	route := L.ToString(1)

	table := L.NewTable()

	table.RawSetString("redirect", lua.LBool(true))
	table.RawSetString("url", lua.LString(route))

	L.Push(table)
	return 1
}

func JSONResponse(L *lua.LState) int {
	header := L.ToTable(1)
	body := L.ToTable(2)
	statusCode := L.ToNumber(3)

	table := L.NewTable()
	jsonText := "{"

	body.ForEach(func(l1, l2 lua.LValue) {
		jsonText += fmt.Sprintf(`"%s" : "%s",`, l1.String(), l2.String())
	})

	jsonText += "}"

	header.RawSetString("Content-Type", lua.LString("application/json"))

	table.RawSetString("header", header)
	table.RawSetString("body", lua.LString(jsonText))
	table.RawSetString("statusCode", statusCode)
	table.RawSetString("json", lua.LBool(true))

	L.Push(table)

	return 1

}

func FileResponse(L *lua.LState) int {
	header := L.ToTable(1)
	fileName := L.ToString(2)
	statusCode := L.ToNumber(3)

	header.RawSetString("Content-Type", lua.LString(mime.TypeByExtension(filepath.Ext(fileName))))

	_, err := os.Stat(fileName)
	if err != nil {
		L.ArgError(2, "File not found")
	}

	data, err := os.ReadFile(fileName)
	HandleError(err)

	table := L.NewTable()

	table.RawSetString("header", header)
	table.RawSetString("body", lua.LString(string(data)))
	table.RawSetString("statusCode", statusCode)

	L.Push(table)
	return 1
}

func TextResponse(L *lua.LState) int {
	header := L.ToTable(1)
	text := L.ToString(2)
	statusCode := L.ToNumber(3)

	header.RawSetString("Content-Type", lua.LString("text/plain"))

	table := L.NewTable()

	table.RawSetString("header", header)
	table.RawSetString("body", lua.LString(text))
	table.RawSetString("statusCode", statusCode)

	L.Push(table)
	return 1
}
