# Hello World project

The hello world project is web application that says Hello {name} on a web page. The default name is World. THe user can provide the name as well, or fetch from a web service part of the project that provides random names.


## Frontend
- Writen in modern HTML5, Javascript (with modules), and the latest CSS Standards.
- Uses Semantic tags
- There should be tests for UI display and interactions.

## Backend
- Serves the HTML and other resources from web root folder frontend/www relative to where the current directory is.
- root folder can also be specifed by the --www {pathname} flag
- port it serves on it by default 8080
- port can be specified by the --port {portnumber} flag
- Written in go.
- Uses a modular approach where the routes, handlers and business logic are seperate.
- Tests using the go test framework
- The go server also an endpoint /api/vi/name-generator that returns a name in json format.
