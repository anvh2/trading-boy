// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cachemock

import (
	"github.com/anvh2/trading-bot/internal/cache"
	"github.com/anvh2/trading-bot/internal/cache/market"
	"sync"
)

// Ensure, that MarketMock does implement cache.Market.
// If this is not the case, regenerate this file with moq.
var _ cache.Market = &MarketMock{}

// MarketMock is a mock implementation of cache.Market.
//
// 	func TestSomethingThatUsesMarket(t *testing.T) {
//
// 		// make and configure a mocked cache.Market
// 		mockedMarket := &MarketMock{
// 			ChartFunc: func(symbol string) (*market.Chart, error) {
// 				panic("mock out the Chart method")
// 			},
// 			CreateChartFunc: func(symbol string) *market.Chart {
// 				panic("mock out the CreateChart method")
// 			},
// 			UpdateChartFunc: func(symbol string) *market.Chart {
// 				panic("mock out the UpdateChart method")
// 			},
// 		}
//
// 		// use mockedMarket in code that requires cache.Market
// 		// and then make assertions.
//
// 	}
type MarketMock struct {
	// ChartFunc mocks the Chart method.
	ChartFunc func(symbol string) (*market.Chart, error)

	// CreateChartFunc mocks the CreateChart method.
	CreateChartFunc func(symbol string) *market.Chart

	// UpdateChartFunc mocks the UpdateChart method.
	UpdateChartFunc func(symbol string) *market.Chart

	// calls tracks calls to the methods.
	calls struct {
		// Chart holds details about calls to the Chart method.
		Chart []struct {
			// Symbol is the symbol argument value.
			Symbol string
		}
		// CreateChart holds details about calls to the CreateChart method.
		CreateChart []struct {
			// Symbol is the symbol argument value.
			Symbol string
		}
		// UpdateChart holds details about calls to the UpdateChart method.
		UpdateChart []struct {
			// Symbol is the symbol argument value.
			Symbol string
		}
	}
	lockChart       sync.RWMutex
	lockCreateChart sync.RWMutex
	lockUpdateChart sync.RWMutex
}

// Chart calls ChartFunc.
func (mock *MarketMock) Chart(symbol string) (*market.Chart, error) {
	if mock.ChartFunc == nil {
		panic("MarketMock.ChartFunc: method is nil but Market.Chart was just called")
	}
	callInfo := struct {
		Symbol string
	}{
		Symbol: symbol,
	}
	mock.lockChart.Lock()
	mock.calls.Chart = append(mock.calls.Chart, callInfo)
	mock.lockChart.Unlock()
	return mock.ChartFunc(symbol)
}

// ChartCalls gets all the calls that were made to Chart.
// Check the length with:
//     len(mockedMarket.ChartCalls())
func (mock *MarketMock) ChartCalls() []struct {
	Symbol string
} {
	var calls []struct {
		Symbol string
	}
	mock.lockChart.RLock()
	calls = mock.calls.Chart
	mock.lockChart.RUnlock()
	return calls
}

// CreateChart calls CreateChartFunc.
func (mock *MarketMock) CreateChart(symbol string) *market.Chart {
	if mock.CreateChartFunc == nil {
		panic("MarketMock.CreateChartFunc: method is nil but Market.CreateChart was just called")
	}
	callInfo := struct {
		Symbol string
	}{
		Symbol: symbol,
	}
	mock.lockCreateChart.Lock()
	mock.calls.CreateChart = append(mock.calls.CreateChart, callInfo)
	mock.lockCreateChart.Unlock()
	return mock.CreateChartFunc(symbol)
}

// CreateChartCalls gets all the calls that were made to CreateChart.
// Check the length with:
//     len(mockedMarket.CreateChartCalls())
func (mock *MarketMock) CreateChartCalls() []struct {
	Symbol string
} {
	var calls []struct {
		Symbol string
	}
	mock.lockCreateChart.RLock()
	calls = mock.calls.CreateChart
	mock.lockCreateChart.RUnlock()
	return calls
}

// UpdateChart calls UpdateChartFunc.
func (mock *MarketMock) UpdateChart(symbol string) *market.Chart {
	if mock.UpdateChartFunc == nil {
		panic("MarketMock.UpdateChartFunc: method is nil but Market.UpdateChart was just called")
	}
	callInfo := struct {
		Symbol string
	}{
		Symbol: symbol,
	}
	mock.lockUpdateChart.Lock()
	mock.calls.UpdateChart = append(mock.calls.UpdateChart, callInfo)
	mock.lockUpdateChart.Unlock()
	return mock.UpdateChartFunc(symbol)
}

// UpdateChartCalls gets all the calls that were made to UpdateChart.
// Check the length with:
//     len(mockedMarket.UpdateChartCalls())
func (mock *MarketMock) UpdateChartCalls() []struct {
	Symbol string
} {
	var calls []struct {
		Symbol string
	}
	mock.lockUpdateChart.RLock()
	calls = mock.calls.UpdateChart
	mock.lockUpdateChart.RUnlock()
	return calls
}