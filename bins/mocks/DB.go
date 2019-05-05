// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import hooks_apimodels "github.com/hugocortes/hooks-api/models"
import mock "github.com/stretchr/testify/mock"
import models "github.com/hugocortes/hooks-api/bins/models"

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// Create provides a mock function with given fields: accountID, bin
func (_m *DB) Create(accountID string, bin *models.Bin) error {
	ret := _m.Called(accountID, bin)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *models.Bin) error); ok {
		r0 = rf(accountID, bin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: accountID, ID
func (_m *DB) Delete(accountID string, ID string) (int, error) {
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
func (_m *DB) Destroy(accountID string) (int, error) {
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
func (_m *DB) Get(accountID string, ID string) (*models.Bin, error) {
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
func (_m *DB) GetAll(accountID string, opts *hooks_apimodels.QueryOpts) ([]*models.Bin, error) {
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
func (_m *DB) Update(accountID string, ID string, bin *models.Bin) (int, error) {
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
