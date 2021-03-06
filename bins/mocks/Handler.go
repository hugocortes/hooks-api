// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import hooks_apimodels "github.com/hugocortes/hooks-api/models"
import mock "github.com/stretchr/testify/mock"
import models "github.com/hugocortes/hooks-api/bins/models"

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Create provides a mock function with given fields: bin
func (_m *Handler) Create(bin *models.Bin) (string, error) {
	ret := _m.Called(bin)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.Bin) string); ok {
		r0 = rf(bin)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Bin) error); ok {
		r1 = rf(bin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: accountID, ID
func (_m *Handler) Delete(accountID string, ID string) (int, error) {
	ret := _m.Called(accountID, ID)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(accountID, ID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountID, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields: accountID
func (_m *Handler) Destroy(accountID string) (int, error) {
	ret := _m.Called(accountID)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(accountID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: accountID, ID
func (_m *Handler) Get(accountID string, ID string) (*models.Bin, error) {
	ret := _m.Called(accountID, ID)

	var r0 *models.Bin
	if rf, ok := ret.Get(0).(func(string, string) *models.Bin); ok {
		r0 = rf(accountID, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Bin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountID, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: accountID, opts
func (_m *Handler) GetAll(accountID string, opts *hooks_apimodels.QueryOpts) ([]*models.Bin, error) {
	ret := _m.Called(accountID, opts)

	var r0 []*models.Bin
	if rf, ok := ret.Get(0).(func(string, *hooks_apimodels.QueryOpts) []*models.Bin); ok {
		r0 = rf(accountID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Bin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *hooks_apimodels.QueryOpts) error); ok {
		r1 = rf(accountID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: accountID, ID, bin
func (_m *Handler) Update(accountID string, ID string, bin *models.Bin) (int, error) {
	ret := _m.Called(accountID, ID, bin)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, *models.Bin) int); ok {
		r0 = rf(accountID, ID, bin)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *models.Bin) error); ok {
		r1 = rf(accountID, ID, bin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
