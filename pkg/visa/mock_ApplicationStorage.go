// Code generated by mockery v1.0.0. DO NOT EDIT.

package visa

import mock "github.com/stretchr/testify/mock"

// MockApplicationStorage is an autogenerated mock type for the ApplicationStorage type
type MockApplicationStorage struct {
	mock.Mock
}

// GetVisaApplication provides a mock function with given fields: id
func (_m *MockApplicationStorage) GetVisaApplication(id int) (*Application, error) {
	ret := _m.Called(id)

	var r0 *Application
	if rf, ok := ret.Get(0).(func(int) *Application); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
