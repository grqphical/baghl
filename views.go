package baghl

import (
	"errors"
	"io"
	"log"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

func MakeRequestTable(L *lua.LState, request http.Request) lua.LTable {
	requestTable := L.NewTable()

	requestTable.RawSetString("URL", lua.LString(request.URL.String()))
	requestTable.RawSetString("Method", lua.LString(request.Method))

	bodyBytes, err := io.ReadAll(request.Body)
	HandleError(err)

	requestTable.RawSetString("Body", lua.LString(string(bodyBytes)))

	headerTable := L.NewTable()

	for key := range request.Header {
		headerTable.RawSetString(key, lua.LString(request.Header.Get(key)))
	}

	requestTable.RawSetString("Header", headerTable)

	return *requestTable
}

func LoadViewFunction(file string, request http.Request) (lua.LValue, error) {
	L := lua.NewState()
	defer L.Close()
	// Load our API
	L.PreloadModule("baghl", Loader)

	// Interpret the lua file
	if err := L.DoFile(file); err != nil {
		return nil, errors.New("404")
	}

	requestTable := MakeRequestTable(L, request)

	// Call the View function inside of the file
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("View"),
		NRet:    1,
		Protect: true,
	}, &requestTable); err != nil {
		log.Fatal(err)
	}

	// Fetch the results of the view function off the stack
	response := L.Get(-1)
	L.Pop(1)

	return response, nil
}

func LoadError(file string, code int) (lua.LValue, error) {
	L := lua.NewState()
	defer L.Close()
	// Load our API
	L.PreloadModule("baghl", Loader)

	// Interpret the lua file
	if err := L.DoFile(file); err != nil {
		return nil, errors.New("404")
	}

	// Call the View function inside of the file
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("Error"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(code)); err != nil {
		log.Fatal(err)
	}

	// Fetch the results of the view function off the stack
	response := L.Get(-1)
	L.Pop(1)

	return response, nil
}
