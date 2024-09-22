package util

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	httpRegisterRoutes []string
	httpRegisterMutex  sync.RWMutex
)

func HttpRegisterRoute(mux *http.ServeMux, method, pattern string, handler http.HandlerFunc) {
	httpRegisterMutex.Lock()
	defer httpRegisterMutex.Unlock()

	route := fmt.Sprintf("%s %s", method, pattern)

	httpRegisterRoutes = append(httpRegisterRoutes, route)
	mux.HandleFunc(route, handler)
}
