package errors

import (
	"runtime"
	"strings"
)

var pkgPrefix = func() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("could not determine package of pkg_util.go")
	}
	p := strings.TrimSuffix(runtime.FuncForPC(pc).Name(), "pkg/errors.init")
	return p
}()

// namespace is called when errors are defined.
// It returns the package path of the caller (skipping the first frames of the call stack)
// and makes it relative (so for example: pkg/errors).
func namespace(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		panic("could not determine source of error")
	}
	fun := runtime.FuncForPC(pc).Name()
	slashIdx := strings.LastIndexByte(fun, '/')
	dotIdx := strings.IndexByte(fun[slashIdx:], '.')
	pkg := fun[:slashIdx+dotIdx]
	return strings.TrimPrefix(pkg, pkgPrefix)
}
