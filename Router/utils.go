package Router

import (
	"fmt"
	"strings"
)

func (r *Router) PrintRoutes() {
	fmt.Println("Registered routes:")
	for method, paths := range r.Handlers {
		for path := range paths {
			fmt.Printf("%s\t%s\n", method, path)
		}
	}

	for _, dr := range r.DynamicRoutes {
		fmt.Printf("%s\t%s\n", dr.Method, dr.Pattern)
	}
}

func MatchPattern(pattern, path string) (map[string]string, bool) {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	params := make(map[string]string)

	for i := 0; i < len(patternParts); i++ {
		if i >= len(pathParts) {
			return nil, false
		}

		pp := patternParts[i]
		pv := pathParts[i]

		if strings.HasPrefix(pp, ":") {
			params[pp[1:]] = pv
		} else if strings.HasPrefix(pp, "*") {
			params[pp[1:]] = strings.Join(pathParts[i:], "/")
			return params, true
		} else if pp != pv {
			return nil, false
		}
	}

	if len(pathParts) != len(patternParts) {
		return nil, false
	}

	return params, true
}

func MatchRoute(route, path string) (bool, map[string]string) {
	routeParts := strings.Split(route, "/")
	pathParts := strings.Split(path, "/")

	if len(routeParts) != len(pathParts) {
		return false, nil
	}

	params := make(map[string]string)
	for i := range routeParts {
		if strings.HasPrefix(routeParts[i], ":") {
			key := strings.TrimPrefix(routeParts[i], ":")
			params[key] = pathParts[i]
		} else if routeParts[i] != pathParts[i] {
			return false, nil
		}
	}

	return true, params
}
