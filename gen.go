// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"bufio"
	"os"
	"io/ioutil"
	"log" 
	"strings"
	"encoding/json"
)

type GenRoute struct {
	Method string `json:"method"`
	Route string `json:"route"`
	Middleware string `json:"middleware"`
	Function string 
}

func main() {
	d, _ := os.Getwd()
	path := d + "/controllers"

	files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
	}

	os.Remove(d + "/routes/web.go")
	
	createPath := d + "/routes/web.go"
	routeFile, _ := os.Create(createPath)
	defer routeFile.Close()
	w := bufio.NewWriter(routeFile)
	defer w.Flush()
	var middlewareString string
	finalStringToWrite := "package routes\n\nimport (\n\tflameroutes \"github.com/flame/routes\"\n\t\"github.com/app/controllers\"\n" + middlewareString + ")\n\nvar Routes = []flameroutes.Route{\n "


	//Some great spaghetti code. This really needs to be refactored.
    for _, f := range files {
		if !f.IsDir() {
			f, _ := os.Open(path + "/" + f.Name())

			scanner := bufio.NewScanner(f)
			scanner.Split(bufio.ScanLines) 
			
			for scanner.Scan() {
				line := scanner.Text()
				
				if strings.HasPrefix(line, "//") {
					slice := strings.Split(line, "routemeta:")
					
					if len(slice) > 1 {
						var route GenRoute
						if err := json.Unmarshal([]byte(strings.Trim(slice[1], " ")), &route); err != nil {
							panic(err)
						} else {
							scanner.Scan()
							nextLine := scanner.Text()
							if strings.HasPrefix(nextLine, "func") {
								nextLine = strings.TrimPrefix(nextLine, "func ")
								splitArr := strings.Split(nextLine, "(")
								if len(splitArr) > 1 {
									nextLine = splitArr[0]
									route.Function = nextLine

									// Need to implement middleware
									// if route.Middleware != "" {
									// 	middlewareString = "\t\"github.com/app/middleware\"\n"
									// }

									stringToWrite := generateRoute(route)
									//w.WriteString(stringToWrite)
									finalStringToWrite += stringToWrite
								}
							}
						}
						
					}
				}
			}
			
			f.Close()
		}
	}
	
	w.WriteString(finalStringToWrite + "}")
}

func generateRoute(route GenRoute) string {

	var routeString string

	switch os := strings.ToUpper(route.Method); os {
	case "GET":
		routeString = "\tflameroutes.Get()"
	case "POST":
		routeString = "\tflameroutes.Post()"
	default:
		routeString = "\tflameroutes.Get()"
	}

	routeString += `.Define("` + route.Route + `", controllers.` + route.Function + `),` + "\n"

	return routeString

}