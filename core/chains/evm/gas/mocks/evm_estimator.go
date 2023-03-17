// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	assets "github.com/smartcontractkit/chainlink/core/assets"

	evmtypes "github.com/smartcontractkit/chainlink/core/chains/evm/types"

	gas "github.com/smartcontractkit/chainlink/core/chains/evm/gas"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/common/txmgr/types"
)

// EvmEstimator is an autogenerated mock type for the EvmEstimator type
type EvmEstimator struct {
	mock.Mock
}

// BumpDynamicFee provides a mock function with given fields: ctx, original, gasLimit, maxGasPriceWei, attempts
func (_m *EvmEstimator) BumpDynamicFee(ctx context.Context, original gas.DynamicFee, gasLimit uint32, maxGasPriceWei *assets.Wei, attempts []gas.EvmPriorAttempt) (gas.DynamicFee, uint32, error) {
	ret := _m.Called(ctx, original, gasLimit, maxGasPriceWei, attempts)

	var r0 gas.DynamicFee
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, gas.DynamicFee, uint32, *assets.Wei, []gas.EvmPriorAttempt) (gas.DynamicFee, uint32, error)); ok {
		return rf(ctx, original, gasLimit, maxGasPriceWei, attempts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, gas.DynamicFee, uint32, *assets.Wei, []gas.EvmPriorAttempt) gas.DynamicFee); ok {
		r0 = rf(ctx, original, gasLimit, maxGasPriceWei, attempts)
	} else {
		r0 = ret.Get(0).(gas.DynamicFee)
	}

	if rf, ok := ret.Get(1).(func(context.Context, gas.DynamicFee, uint32, *assets.Wei, []gas.EvmPriorAttempt) uint32); ok {
		r1 = rf(ctx, original, gasLimit, maxGasPriceWei, attempts)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, gas.DynamicFee, uint32, *assets.Wei, []gas.EvmPriorAttempt) error); ok {
		r2 = rf(ctx, original, gasLimit, maxGasPriceWei, attempts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BumpLegacyGas provides a mock function with given fields: ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts
func (_m *EvmEstimator) BumpLegacyGas(ctx context.Context, originalGasPrice *assets.Wei, gasLimit uint32, maxGasPriceWei *assets.Wei, attempts []gas.EvmPriorAttempt) (*assets.Wei, uint32, error) {
	ret := _m.Called(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)

	var r0 *assets.Wei
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *assets.Wei, uint32, *assets.Wei, []gas.EvmPriorAttempt) (*assets.Wei, uint32, error)); ok {
		return rf(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *assets.Wei, uint32, *assets.Wei, []gas.EvmPriorAttempt) *assets.Wei); ok {
		r0 = rf(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *assets.Wei, uint32, *assets.Wei, []gas.EvmPriorAttempt) uint32); ok {
		r1 = rf(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *assets.Wei, uint32, *assets.Wei, []gas.EvmPriorAttempt) error); ok {
		r2 = rf(ctx, originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Close provides a mock function with given fields:
func (_m *EvmEstimator) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDynamicFee provides a mock function with given fields: ctx, gasLimit, maxGasPriceWei
func (_m *EvmEstimator) GetDynamicFee(ctx context.Context, gasLimit uint32, maxGasPriceWei *assets.Wei) (gas.DynamicFee, uint32, error) {
	ret := _m.Called(ctx, gasLimit, maxGasPriceWei)

	var r0 gas.DynamicFee
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, *assets.Wei) (gas.DynamicFee, uint32, error)); ok {
		return rf(ctx, gasLimit, maxGasPriceWei)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, *assets.Wei) gas.DynamicFee); ok {
		r0 = rf(ctx, gasLimit, maxGasPriceWei)
	} else {
		r0 = ret.Get(0).(gas.DynamicFee)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, *assets.Wei) uint32); ok {
		r1 = rf(ctx, gasLimit, maxGasPriceWei)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint32, *assets.Wei) error); ok {
		r2 = rf(ctx, gasLimit, maxGasPriceWei)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLegacyGas provides a mock function with given fields: ctx, calldata, gasLimit, maxGasPriceWei, opts
func (_m *EvmEstimator) GetLegacyGas(ctx context.Context, calldata []byte, gasLimit uint32, maxGasPriceWei *assets.Wei, opts ...types.Opt) (*assets.Wei, uint32, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, calldata, gasLimit, maxGasPriceWei)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *assets.Wei
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint32, *assets.Wei, ...types.Opt) (*assets.Wei, uint32, error)); ok {
		return rf(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint32, *assets.Wei, ...types.Opt) *assets.Wei); ok {
		r0 = rf(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, uint32, *assets.Wei, ...types.Opt) uint32); ok {
		r1 = rf(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, []byte, uint32, *assets.Wei, ...types.Opt) error); ok {
		r2 = rf(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *EvmEstimator) OnNewLongestChain(ctx context.Context, head *evmtypes.Head) {
	_m.Called(ctx, head)
}

// Start provides a mock function with given fields: _a0
func (_m *EvmEstimator) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewEvmEstimator interface {
	mock.TestingT
	Cleanup(func())
}

// NewEvmEstimator creates a new instance of EvmEstimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEvmEstimator(t mockConstructorTestingTNewEvmEstimator) *EvmEstimator {
	mock := &EvmEstimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
