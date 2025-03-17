// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	io "io"

	ftp "github.com/jlaffaye/ftp"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: path
func (_m *Client) Delete(path string) error {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Client_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - path string
func (_e *Client_Expecter) Delete(path interface{}) *Client_Delete_Call {
	return &Client_Delete_Call{Call: _e.mock.On("Delete", path)}
}

func (_c *Client_Delete_Call) Run(run func(path string)) *Client_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_Delete_Call) Return(_a0 error) *Client_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Delete_Call) RunAndReturn(run func(string) error) *Client_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetEntry provides a mock function with given fields: p
func (_m *Client) GetEntry(p string) (*ftp.Entry, error) {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for GetEntry")
	}

	var r0 *ftp.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*ftp.Entry, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func(string) *ftp.Entry); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ftp.Entry)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetEntry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEntry'
type Client_GetEntry_Call struct {
	*mock.Call
}

// GetEntry is a helper method to define mock.On call
//   - p string
func (_e *Client_Expecter) GetEntry(p interface{}) *Client_GetEntry_Call {
	return &Client_GetEntry_Call{Call: _e.mock.On("GetEntry", p)}
}

func (_c *Client_GetEntry_Call) Run(run func(p string)) *Client_GetEntry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_GetEntry_Call) Return(_a0 *ftp.Entry, _a1 error) *Client_GetEntry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetEntry_Call) RunAndReturn(run func(string) (*ftp.Entry, error)) *Client_GetEntry_Call {
	_c.Call.Return(run)
	return _c
}

// IsSetTimeSupported provides a mock function with no fields
func (_m *Client) IsSetTimeSupported() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSetTimeSupported")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Client_IsSetTimeSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSetTimeSupported'
type Client_IsSetTimeSupported_Call struct {
	*mock.Call
}

// IsSetTimeSupported is a helper method to define mock.On call
func (_e *Client_Expecter) IsSetTimeSupported() *Client_IsSetTimeSupported_Call {
	return &Client_IsSetTimeSupported_Call{Call: _e.mock.On("IsSetTimeSupported")}
}

func (_c *Client_IsSetTimeSupported_Call) Run(run func()) *Client_IsSetTimeSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_IsSetTimeSupported_Call) Return(_a0 bool) *Client_IsSetTimeSupported_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_IsSetTimeSupported_Call) RunAndReturn(run func() bool) *Client_IsSetTimeSupported_Call {
	_c.Call.Return(run)
	return _c
}

// IsTimePreciseInList provides a mock function with no fields
func (_m *Client) IsTimePreciseInList() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsTimePreciseInList")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Client_IsTimePreciseInList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsTimePreciseInList'
type Client_IsTimePreciseInList_Call struct {
	*mock.Call
}

// IsTimePreciseInList is a helper method to define mock.On call
func (_e *Client_Expecter) IsTimePreciseInList() *Client_IsTimePreciseInList_Call {
	return &Client_IsTimePreciseInList_Call{Call: _e.mock.On("IsTimePreciseInList")}
}

func (_c *Client_IsTimePreciseInList_Call) Run(run func()) *Client_IsTimePreciseInList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_IsTimePreciseInList_Call) Return(_a0 bool) *Client_IsTimePreciseInList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_IsTimePreciseInList_Call) RunAndReturn(run func() bool) *Client_IsTimePreciseInList_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: p
func (_m *Client) List(p string) ([]*ftp.Entry, error) {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*ftp.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*ftp.Entry, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func(string) []*ftp.Entry); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ftp.Entry)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type Client_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - p string
func (_e *Client_Expecter) List(p interface{}) *Client_List_Call {
	return &Client_List_Call{Call: _e.mock.On("List", p)}
}

func (_c *Client_List_Call) Run(run func(p string)) *Client_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_List_Call) Return(_a0 []*ftp.Entry, _a1 error) *Client_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_List_Call) RunAndReturn(run func(string) ([]*ftp.Entry, error)) *Client_List_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: user, password
func (_m *Client) Login(user string, password string) error {
	ret := _m.Called(user, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(user, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type Client_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - user string
//   - password string
func (_e *Client_Expecter) Login(user interface{}, password interface{}) *Client_Login_Call {
	return &Client_Login_Call{Call: _e.mock.On("Login", user, password)}
}

func (_c *Client_Login_Call) Run(run func(user string, password string)) *Client_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Client_Login_Call) Return(_a0 error) *Client_Login_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Login_Call) RunAndReturn(run func(string, string) error) *Client_Login_Call {
	_c.Call.Return(run)
	return _c
}

// MakeDir provides a mock function with given fields: path
func (_m *Client) MakeDir(path string) error {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for MakeDir")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_MakeDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MakeDir'
type Client_MakeDir_Call struct {
	*mock.Call
}

// MakeDir is a helper method to define mock.On call
//   - path string
func (_e *Client_Expecter) MakeDir(path interface{}) *Client_MakeDir_Call {
	return &Client_MakeDir_Call{Call: _e.mock.On("MakeDir", path)}
}

func (_c *Client_MakeDir_Call) Run(run func(path string)) *Client_MakeDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_MakeDir_Call) Return(_a0 error) *Client_MakeDir_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_MakeDir_Call) RunAndReturn(run func(string) error) *Client_MakeDir_Call {
	_c.Call.Return(run)
	return _c
}

// Quit provides a mock function with no fields
func (_m *Client) Quit() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Quit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Quit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Quit'
type Client_Quit_Call struct {
	*mock.Call
}

// Quit is a helper method to define mock.On call
func (_e *Client_Expecter) Quit() *Client_Quit_Call {
	return &Client_Quit_Call{Call: _e.mock.On("Quit")}
}

func (_c *Client_Quit_Call) Run(run func()) *Client_Quit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_Quit_Call) Return(_a0 error) *Client_Quit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Quit_Call) RunAndReturn(run func() error) *Client_Quit_Call {
	_c.Call.Return(run)
	return _c
}

// Rename provides a mock function with given fields: from, to
func (_m *Client) Rename(from string, to string) error {
	ret := _m.Called(from, to)

	if len(ret) == 0 {
		panic("no return value specified for Rename")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(from, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Rename_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rename'
type Client_Rename_Call struct {
	*mock.Call
}

// Rename is a helper method to define mock.On call
//   - from string
//   - to string
func (_e *Client_Expecter) Rename(from interface{}, to interface{}) *Client_Rename_Call {
	return &Client_Rename_Call{Call: _e.mock.On("Rename", from, to)}
}

func (_c *Client_Rename_Call) Run(run func(from string, to string)) *Client_Rename_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Client_Rename_Call) Return(_a0 error) *Client_Rename_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Rename_Call) RunAndReturn(run func(string, string) error) *Client_Rename_Call {
	_c.Call.Return(run)
	return _c
}

// RetrFrom provides a mock function with given fields: path, offset
func (_m *Client) RetrFrom(path string, offset uint64) (*ftp.Response, error) {
	ret := _m.Called(path, offset)

	if len(ret) == 0 {
		panic("no return value specified for RetrFrom")
	}

	var r0 *ftp.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uint64) (*ftp.Response, error)); ok {
		return rf(path, offset)
	}
	if rf, ok := ret.Get(0).(func(string, uint64) *ftp.Response); ok {
		r0 = rf(path, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ftp.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, uint64) error); ok {
		r1 = rf(path, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_RetrFrom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrFrom'
type Client_RetrFrom_Call struct {
	*mock.Call
}

// RetrFrom is a helper method to define mock.On call
//   - path string
//   - offset uint64
func (_e *Client_Expecter) RetrFrom(path interface{}, offset interface{}) *Client_RetrFrom_Call {
	return &Client_RetrFrom_Call{Call: _e.mock.On("RetrFrom", path, offset)}
}

func (_c *Client_RetrFrom_Call) Run(run func(path string, offset uint64)) *Client_RetrFrom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint64))
	})
	return _c
}

func (_c *Client_RetrFrom_Call) Return(_a0 *ftp.Response, _a1 error) *Client_RetrFrom_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_RetrFrom_Call) RunAndReturn(run func(string, uint64) (*ftp.Response, error)) *Client_RetrFrom_Call {
	_c.Call.Return(run)
	return _c
}

// SetTime provides a mock function with given fields: path, t
func (_m *Client) SetTime(path string, t time.Time) error {
	ret := _m.Called(path, t)

	if len(ret) == 0 {
		panic("no return value specified for SetTime")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time) error); ok {
		r0 = rf(path, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_SetTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTime'
type Client_SetTime_Call struct {
	*mock.Call
}

// SetTime is a helper method to define mock.On call
//   - path string
//   - t time.Time
func (_e *Client_Expecter) SetTime(path interface{}, t interface{}) *Client_SetTime_Call {
	return &Client_SetTime_Call{Call: _e.mock.On("SetTime", path, t)}
}

func (_c *Client_SetTime_Call) Run(run func(path string, t time.Time)) *Client_SetTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(time.Time))
	})
	return _c
}

func (_c *Client_SetTime_Call) Return(_a0 error) *Client_SetTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_SetTime_Call) RunAndReturn(run func(string, time.Time) error) *Client_SetTime_Call {
	_c.Call.Return(run)
	return _c
}

// StorFrom provides a mock function with given fields: path, r, offset
func (_m *Client) StorFrom(path string, r io.Reader, offset uint64) error {
	ret := _m.Called(path, r, offset)

	if len(ret) == 0 {
		panic("no return value specified for StorFrom")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader, uint64) error); ok {
		r0 = rf(path, r, offset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_StorFrom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StorFrom'
type Client_StorFrom_Call struct {
	*mock.Call
}

// StorFrom is a helper method to define mock.On call
//   - path string
//   - r io.Reader
//   - offset uint64
func (_e *Client_Expecter) StorFrom(path interface{}, r interface{}, offset interface{}) *Client_StorFrom_Call {
	return &Client_StorFrom_Call{Call: _e.mock.On("StorFrom", path, r, offset)}
}

func (_c *Client_StorFrom_Call) Run(run func(path string, r io.Reader, offset uint64)) *Client_StorFrom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(io.Reader), args[2].(uint64))
	})
	return _c
}

func (_c *Client_StorFrom_Call) Return(_a0 error) *Client_StorFrom_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_StorFrom_Call) RunAndReturn(run func(string, io.Reader, uint64) error) *Client_StorFrom_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
