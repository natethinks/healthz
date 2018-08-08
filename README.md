# healthz

Healthz serves a single http endpoint for service health monitoring by kubernetes. Users of the package submit errors which are stored in unexported package variables. If any fatal errors have been reported, /healthz will return a status 500, otherwise it will return a status 200. All errors will be returned in the body of the response in either case.

```
healthz.Serve("localhost:8080", "/healthz") // params shown here are redundant, these are the package defaults
healthError := healthz.HealthError{Description: "Something went wrong!"}
healthz.NewFatalError(healthError)
```

Examples are found in the examples directory of the project
