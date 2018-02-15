// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/airking05/go-auto-trader/models"

// ExchangePrivateRepository is an autogenerated mock type for the ExchangePrivateRepository type
type ExchangePrivateRepository struct {
	mock.Mock
}

// ActiveOrders provides a mock function with given fields:
func (_m *ExchangePrivateRepository) ActiveOrders() ([]*models.Order, error) {
	ret := _m.Called()

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func() []*models.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Address provides a mock function with given fields: c
func (_m *ExchangePrivateRepository) Address(c string) (string, error) {
	ret := _m.Called(c)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Balances provides a mock function with given fields:
func (_m *ExchangePrivateRepository) Balances() (map[string]float64, error) {
	ret := _m.Called()

	var r0 map[string]float64
	if rf, ok := ret.Get(0).(func() map[string]float64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]float64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelOrder provides a mock function with given fields: orderNumber, productCode
func (_m *ExchangePrivateRepository) CancelOrder(orderNumber string, productCode string) error {
	ret := _m.Called(orderNumber, productCode)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(orderNumber, productCode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteBalances provides a mock function with given fields:
func (_m *ExchangePrivateRepository) CompleteBalances() (map[string]*models.Balance, error) {
	ret := _m.Called()

	var r0 map[string]*models.Balance
	if rf, ok := ret.Get(0).(func() map[string]*models.Balance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*models.Balance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Order provides a mock function with given fields: trading, settlement, ordertype, price, amount
func (_m *ExchangePrivateRepository) Order(trading string, settlement string, ordertype models.OrderType, price float64, amount float64) (string, error) {
	ret := _m.Called(trading, settlement, ordertype, price, amount)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, models.OrderType, float64, float64) string); ok {
		r0 = rf(trading, settlement, ordertype, price, amount)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, models.OrderType, float64, float64) error); ok {
		r1 = rf(trading, settlement, ordertype, price, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurchaseFeeRate provides a mock function with given fields:
func (_m *ExchangePrivateRepository) PurchaseFeeRate() (float64, error) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SellFeeRate provides a mock function with given fields:
func (_m *ExchangePrivateRepository) SellFeeRate() (float64, error) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: typ, addr, amount, additionalFee
func (_m *ExchangePrivateRepository) Transfer(typ string, addr string, amount float64, additionalFee float64) error {
	ret := _m.Called(typ, addr, amount, additionalFee)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, float64, float64) error); ok {
		r0 = rf(typ, addr, amount, additionalFee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransferFee provides a mock function with given fields:
func (_m *ExchangePrivateRepository) TransferFee() (map[string]float64, error) {
	ret := _m.Called()

	var r0 map[string]float64
	if rf, ok := ret.Get(0).(func() map[string]float64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]float64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
