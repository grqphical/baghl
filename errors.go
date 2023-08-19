package baghl

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HandleHTTPError(code int) string {
	ret, err := LoadError("./handlers/error.lua", code)
	HandleError(err)

	return ret.(*lua.LTable).RawGetString("body").String()
}
