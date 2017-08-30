/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package routing

import (
	"fmt"

	mock "github.com/stretchr/testify/mock"

	http "net/http"
)

// RouterMock mock
type RouterMock struct {
	mock.Mock
}

// NewRouterMock ...
func NewRouterMock() *RouterMock {
	return &RouterMock{}
}

// Connect mocked method
func (m *RouterMock) Connect(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Delete mocked method
func (m *RouterMock) Delete(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Get mocked method
func (m *RouterMock) Get(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Group mocked method
func (m *RouterMock) Group(p0 func(r Router)) Router {

	ret := m.Called(p0)

	var r0 Router
	switch res := ret.Get(0).(type) {
	case nil:
	case Router:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// Handle mocked method
func (m *RouterMock) Handle(p0 string, p1 http.Handler) {

	m.Called(p0, p1)

}

// HandleFunc mocked method
func (m *RouterMock) HandleFunc(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Head mocked method
func (m *RouterMock) Head(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Load mocked method
func (m *RouterMock) Load(p0 RouteLoader) error {

	ret := m.Called(p0)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// Method mocked method
func (m *RouterMock) Method(p0 string, p1 string, p2 http.Handler) {

	m.Called(p0, p1, p2)

}

// MethodFunc mocked method
func (m *RouterMock) MethodFunc(p0 string, p1 string, p2 http.HandlerFunc) {

	m.Called(p0, p1, p2)

}

// MethodNotAllowed mocked method
func (m *RouterMock) MethodNotAllowed(p0 http.HandlerFunc) {

	m.Called(p0)

}

// Mount mocked method
func (m *RouterMock) Mount(p0 string, p1 http.Handler) {

	m.Called(p0, p1)

}

// NotFound mocked method
func (m *RouterMock) NotFound(p0 http.HandlerFunc) {

	m.Called(p0)

}

// Options mocked method
func (m *RouterMock) Options(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Patch mocked method
func (m *RouterMock) Patch(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Post mocked method
func (m *RouterMock) Post(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Put mocked method
func (m *RouterMock) Put(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Route mocked method
func (m *RouterMock) Route(p0 string, p1 func(r Router)) Router {

	ret := m.Called(p0, p1)

	var r0 Router
	switch res := ret.Get(0).(type) {
	case nil:
	case Router:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ServeHTTP mocked method
func (m *RouterMock) ServeHTTP(p0 http.ResponseWriter, p1 *http.Request) {

	m.Called(p0, p1)

}

// Trace mocked method
func (m *RouterMock) Trace(p0 string, p1 http.HandlerFunc) {

	m.Called(p0, p1)

}

// Use mocked method
func (m *RouterMock) Use(p0 ...func(http.Handler) http.Handler) {

	m.Called(p0)

}

// With mocked method
func (m *RouterMock) With(p0 ...func(http.Handler) http.Handler) Router {

	ret := m.Called(p0)

	var r0 Router
	switch res := ret.Get(0).(type) {
	case nil:
	case Router:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}
