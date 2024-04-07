// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetNumberSentinelSlavesInMemory provides a mock function with given fields: ip
func (_m *Client) GetNumberSentinelSlavesInMemory(ip string) (int32, error) {
	ret := _m.Called(ip)

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int32, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(string) int32); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberSentinelsInMemory provides a mock function with given fields: ip
func (_m *Client) GetNumberSentinelsInMemory(ip string) (int32, error) {
	ret := _m.Called(ip)

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int32, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(string) int32); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSentinelMonitor provides a mock function with given fields: ip
func (_m *Client) GetSentinelMonitor(ip string) (string, string, error) {
	ret := _m.Called(ip)

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (string, string, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(ip)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSlaveOf provides a mock function with given fields: ip, port, password
func (_m *Client) GetSlaveOf(ip string, port string, password string) (string, error) {
	ret := _m.Called(ip, port, password)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(ip, port, password)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(ip, port, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(ip, port, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsMaster provides a mock function with given fields: ip, port, password
func (_m *Client) IsMaster(ip string, port string, password string) (bool, error) {
	ret := _m.Called(ip, port, password)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (bool, error)); ok {
		return rf(ip, port, password)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(ip, port, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(ip, port, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeMaster provides a mock function with given fields: ip, port, password
func (_m *Client) MakeMaster(ip string, port string, password string) error {
	ret := _m.Called(ip, port, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(ip, port, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MakeSlaveOf provides a mock function with given fields: ip, masterIP, password
func (_m *Client) MakeSlaveOf(ip string, masterIP string, password string) error {
	ret := _m.Called(ip, masterIP, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(ip, masterIP, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MakeSlaveOfWithPort provides a mock function with given fields: ip, masterIP, masterPort, password
func (_m *Client) MakeSlaveOfWithPort(ip string, masterIP string, masterPort string, password string) error {
	ret := _m.Called(ip, masterIP, masterPort, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(ip, masterIP, masterPort, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MonitorRedis provides a mock function with given fields: ip, monitor, quorum, password
func (_m *Client) MonitorRedis(ip string, monitor string, quorum string, password string) error {
	ret := _m.Called(ip, monitor, quorum, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(ip, monitor, quorum, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MonitorRedisWithPort provides a mock function with given fields: ip, monitor, port, quorum, password
func (_m *Client) MonitorRedisWithPort(ip string, monitor string, port string, quorum string, password string) error {
	ret := _m.Called(ip, monitor, port, quorum, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) error); ok {
		r0 = rf(ip, monitor, port, quorum, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetSentinel provides a mock function with given fields: ip
func (_m *Client) ResetSentinel(ip string) error {
	ret := _m.Called(ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SentinelCheckQuorum provides a mock function with given fields: ip
func (_m *Client) SentinelCheckQuorum(ip string) error {
	ret := _m.Called(ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCustomRedisConfig provides a mock function with given fields: ip, port, configs, password
func (_m *Client) SetCustomRedisConfig(ip string, port string, configs []string, password string) error {
	ret := _m.Called(ip, port, configs, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string, string) error); ok {
		r0 = rf(ip, port, configs, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCustomSentinelConfig provides a mock function with given fields: ip, configs
func (_m *Client) SetCustomSentinelConfig(ip string, configs []string) error {
	ret := _m.Called(ip, configs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(ip, configs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SlaveIsReady provides a mock function with given fields: ip, port, password
func (_m *Client) SlaveIsReady(ip string, port string, password string) (bool, error) {
	ret := _m.Called(ip, port, password)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (bool, error)); ok {
		return rf(ip, port, password)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(ip, port, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(ip, port, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
