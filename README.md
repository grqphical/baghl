# baghl

A simple web stack that uses:

**B** [BulmaCSS](https://bulma.io).

**A** [AlpineJS](https://alpinejs.dev).

**G** [GO](https://go.dev).

**H** [HTMX](https://htmx.org).

**L** [Lua](https://lua.org).

## Setup

Make sure to install the baghl-cli with 
```bash
$   go install github.com/grqphical07/baghl-cli
```

Then to create a new project run:
```bash
$   baghl-cli create PROJECT_NAME_HERE
```

This will create a new baghl project in the current folder. To add a route to the web app just create a file called ```get.lua``` under routes. Baghl uses filesystem routing so the path of a file is it's
url. You can create multiple HTTP methods for different routes, just create a lua file called ```METHOD.lua``` with an all lowercase name.

Inside of ```get.lua``` put:
```lua
local baghl = require("baghl")

function View(request)
    return baghl.StringResponse(request.Header, "Hello World", 200)
end
```

Then finally run the go app to use your web app
```bash
$   go run .
```
