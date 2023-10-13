# HAG stack - Web Example
> HTMX, AlpineJS, Go stack

This web app is a simple application to see the validity and performance of the HAG stack. It will slowly be expanded to include DB read/write/delete, A&A (Authentication and Authorization), responsive design, etc... 

## Prerequisites
1. Go v1.21.3 or later
2. The `templ` tool: ```go install github.com/a-h/templ/cmd/templ@latest```

## Build
There are two ways of building the app, either run the `build.sh` script, or perform the following steps:
1. From the root directory run: `cd ./internal/components/` followed by ```templ generate```
2. From the root directory run: `cd ./internal/pages` followed by ```templ generate```

These commands build the required `.go` files to properly serve the pages and compenents necessary for the server to function.

3. From the root directory run: `go build *.go` or `go run *.go`

## Changelog
(13-OCT-2023)
- Project started
- README.md created and updated