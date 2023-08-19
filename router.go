package baghl

import (
	"io/fs"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	lua "github.com/yuin/gopher-lua"
)

func CreateGetRoute(path string, d fs.DirEntry, router *gin.Engine) {
	router.GET(FilesystemRouteToURLEndpoint(path), func(ctx *gin.Context) {
		response, err := LoadViewFunction(filepath.Join(path, "get.lua"), *ctx.Request)
		if err != nil {
			ctx.String(404, HandleHTTPError(404))
			return

		}

		responseTable := response.(*lua.LTable)

		if responseTable.RawGetString("redirect") == lua.LBool(true) {
			ctx.Redirect(302, responseTable.RawGetString("url").String())
			return
		}

		responseTable.RawGetString("header").(*lua.LTable).ForEach(func(l1, l2 lua.LValue) {
			ctx.Writer.Header().Add(l1.String(), l2.String())
		})

		if ctx.Writer.Header().Get("Content-Type") == "" {
			ctx.Writer.Header().Add("Content-Type", "text/html")
		}

		statusCode, err := strconv.ParseInt(responseTable.RawGetString("statusCode").String(), 10, 32)
		HandleError(err)

		ctx.String(int(statusCode), responseTable.RawGetString("body").String())
	})
}

func CreatePostRoute(path string, d fs.DirEntry, router *gin.Engine) {
	router.POST(FilesystemRouteToURLEndpoint(path), func(ctx *gin.Context) {
		response, err := LoadViewFunction(filepath.Join(path, "post.lua"), *ctx.Request)
		if err != nil {
			ctx.String(404, HandleHTTPError(404))
			return

		}

		responseTable := response.(*lua.LTable)

		if responseTable.RawGetString("redirect") == lua.LBool(true) {
			ctx.Redirect(302, responseTable.RawGetString("url").String())
			return
		}

		responseTable.RawGetString("header").(*lua.LTable).ForEach(func(l1, l2 lua.LValue) {
			ctx.Writer.Header().Add(l1.String(), l2.String())
		})

		if ctx.Writer.Header().Get("Content-Type") == "" {
			ctx.Writer.Header().Add("Content-Type", "text/html")
		}

		statusCode, err := strconv.ParseInt(responseTable.RawGetString("statusCode").String(), 10, 32)
		HandleError(err)

		ctx.String(int(statusCode), responseTable.RawGetString("body").String())
	})
}

func CreatePutRoute(path string, d fs.DirEntry, router *gin.Engine) {
	router.PUT(FilesystemRouteToURLEndpoint(path), func(ctx *gin.Context) {
		response, err := LoadViewFunction(filepath.Join(path, "put.lua"), *ctx.Request)
		if err != nil {
			ctx.String(404, HandleHTTPError(404))
			return

		}

		responseTable := response.(*lua.LTable)

		if responseTable.RawGetString("redirect") == lua.LBool(true) {
			ctx.Redirect(302, responseTable.RawGetString("url").String())
			return
		}

		responseTable.RawGetString("header").(*lua.LTable).ForEach(func(l1, l2 lua.LValue) {
			ctx.Writer.Header().Add(l1.String(), l2.String())
		})

		if ctx.Writer.Header().Get("Content-Type") == "" {
			ctx.Writer.Header().Add("Content-Type", "text/html")
		}

		statusCode, err := strconv.ParseInt(responseTable.RawGetString("statusCode").String(), 10, 32)
		HandleError(err)

		ctx.String(int(statusCode), responseTable.RawGetString("body").String())
	})
}

func CreateDeleteRoute(path string, d fs.DirEntry, router *gin.Engine) {
	router.DELETE(FilesystemRouteToURLEndpoint(path), func(ctx *gin.Context) {
		response, err := LoadViewFunction(filepath.Join(path, "delete.lua"), *ctx.Request)
		if err != nil {
			ctx.String(404, HandleHTTPError(404))
			return

		}

		responseTable := response.(*lua.LTable)

		if responseTable.RawGetString("redirect") == lua.LBool(true) {
			ctx.Redirect(302, responseTable.RawGetString("url").String())
			return
		}

		responseTable.RawGetString("header").(*lua.LTable).ForEach(func(l1, l2 lua.LValue) {
			ctx.Writer.Header().Add(l1.String(), l2.String())
		})

		if ctx.Writer.Header().Get("Content-Type") == "" {
			ctx.Writer.Header().Add("Content-Type", "text/html")
		}

		statusCode, err := strconv.ParseInt(responseTable.RawGetString("statusCode").String(), 10, 32)
		HandleError(err)

		ctx.String(int(statusCode), responseTable.RawGetString("body").String())
	})
}

func CreateOptionsRoute(path string, d fs.DirEntry, router *gin.Engine) {
	router.OPTIONS(FilesystemRouteToURLEndpoint(path), func(ctx *gin.Context) {
		response, err := LoadViewFunction(filepath.Join(path, "options.lua"), *ctx.Request)
		if err != nil {
			ctx.String(404, HandleHTTPError(404))
			return

		}

		responseTable := response.(*lua.LTable)

		if responseTable.RawGetString("redirect") == lua.LBool(true) {
			ctx.Redirect(302, responseTable.RawGetString("url").String())
			return
		}

		responseTable.RawGetString("header").(*lua.LTable).ForEach(func(l1, l2 lua.LValue) {
			ctx.Writer.Header().Add(l1.String(), l2.String())
		})

		if ctx.Writer.Header().Get("Content-Type") == "" {
			ctx.Writer.Header().Add("Content-Type", "text/html")
		}

		statusCode, err := strconv.ParseInt(responseTable.RawGetString("statusCode").String(), 10, 32)
		HandleError(err)

		ctx.String(int(statusCode), responseTable.RawGetString("body").String())
	})
}

func CreateHeadRoute(path string, d fs.DirEntry, router *gin.Engine) {
	router.HEAD(FilesystemRouteToURLEndpoint(path), func(ctx *gin.Context) {
		response, err := LoadViewFunction(filepath.Join(path, "head.lua"), *ctx.Request)
		if err != nil {
			ctx.String(404, HandleHTTPError(404))
			return

		}

		responseTable := response.(*lua.LTable)

		if responseTable.RawGetString("redirect") == lua.LBool(true) {
			ctx.Redirect(302, responseTable.RawGetString("url").String())
			return
		}

		responseTable.RawGetString("header").(*lua.LTable).ForEach(func(l1, l2 lua.LValue) {
			ctx.Writer.Header().Add(l1.String(), l2.String())
		})

		if ctx.Writer.Header().Get("Content-Type") == "" {
			ctx.Writer.Header().Add("Content-Type", "text/html")
		}

		statusCode, err := strconv.ParseInt(responseTable.RawGetString("statusCode").String(), 10, 32)
		HandleError(err)

		ctx.String(int(statusCode), responseTable.RawGetString("body").String())
	})
}

func CreatePatchRoute(path string, d fs.DirEntry, router *gin.Engine) {
	router.PATCH(FilesystemRouteToURLEndpoint(path), func(ctx *gin.Context) {
		response, err := LoadViewFunction(filepath.Join(path, "patch.lua"), *ctx.Request)
		if err != nil {
			ctx.String(404, HandleHTTPError(404))
			return

		}

		responseTable := response.(*lua.LTable)

		if responseTable.RawGetString("redirect") == lua.LBool(true) {
			ctx.Redirect(302, responseTable.RawGetString("url").String())
			return
		}

		responseTable.RawGetString("header").(*lua.LTable).ForEach(func(l1, l2 lua.LValue) {
			ctx.Writer.Header().Add(l1.String(), l2.String())
		})

		if ctx.Writer.Header().Get("Content-Type") == "" {
			ctx.Writer.Header().Add("Content-Type", "text/html")
		}

		statusCode, err := strconv.ParseInt(responseTable.RawGetString("statusCode").String(), 10, 32)
		HandleError(err)

		ctx.String(int(statusCode), responseTable.RawGetString("body").String())
	})
}

func GenerateFilesystemRoutes(router *gin.Engine) {
	// Generate the other routes in the filesystem
	filepath.WalkDir("routes/", func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			CreateGetRoute(path, d, router)
			CreatePostRoute(path, d, router)
			CreatePatchRoute(path, d, router)
			CreateDeleteRoute(path, d, router)
			CreatePutRoute(path, d, router)
			CreateHeadRoute(path, d, router)
			CreateOptionsRoute(path, d, router)
		}
		return nil
	})
}

// Create the HTTP router based on the filesystem routes
func CreateRouter() *gin.Engine {
	router := gin.Default()
	GenerateFilesystemRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "text/html")
		c.String(404, HandleHTTPError(404))
	})

	router.Static("/static", "./static")
	router.Static("/asset", "./assets")

	return router
}
