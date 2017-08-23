package tonic

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ApplicationIndex",
		"GET",
		"/applications",
		ApplicationIndex,
	},
	Route{
		"ApplicationCreate",
		"POST",
		"/applications",
		ApplicationCreate,
	},
	Route{
		"ApplicationUpdate",
		"PUT",
		"/applications",
		ApplicationUpdate,
	},
	Route{
		"ApplicationShow",
		"GET",
		"/applications/{applicationId}",
		ApplicationShow,
	},
	Route{
		"ApplicationDestroy",
		"GET",
		"/applications/{applicationId}",
		ApplicationDestroy,
	},
	/////////////////////
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}
