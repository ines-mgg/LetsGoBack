package Router

import (
	"encoding/json"
	"fmt"
	"os"
)

// PrintRoutes prints all registered routes in the router.
// It iterates through the Handlers map and the DynamicRoutes slice,
// printing the HTTP method and path for each route.
func (r *Router) PrintRoutes() {
	fmt.Println("Registered routes:")
	for method, paths := range r.Handlers {
		for path := range paths {
			fmt.Printf("%s\t%s\n", method, path)
		}
	}

	for _, dr := range r.DynamicRoutes {
		fmt.Printf("%s\t%s\n", dr.method, dr.pattern)
	}
}

// WriteRoutesToJsonFile writes all registered routes to a JSON file.
// It creates a slice of routeInfo structs, each containing the HTTP method and path.
// The routes are then marshaled into JSON format and written to the specified file.
// The filename is appended with ".json" to indicate the file format.
// If there is an error during marshaling or file writing, it returns the error.
// The JSON file will contain an array of objects, each representing a route with its method and path.
func (r *Router) WriteRoutesToJsonFile(filename string) error {
	type routeInfo struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}

	var routes []routeInfo

	for method, paths := range r.Handlers {
		for path := range paths {
			routes = append(routes, routeInfo{
				Method: method,
				Path:   path,
			})
		}
	}

	for _, dr := range r.DynamicRoutes {
		routes = append(routes, routeInfo{
			Method: dr.method,
			Path:   dr.pattern,
		})
	}

	data, err := json.MarshalIndent(routes, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename+".json", data, 0644)
}
