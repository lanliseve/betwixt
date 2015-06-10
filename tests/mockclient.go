package tests

import (
	"errors"
	. "github.com/zubairhamed/betwixt"
)

func NewMockClient() LWM2MClient {
	return &MockClient{
		enabledObjects: make(map[LWM2MObjectType]Object),
	}
}

type MockClient struct {
	enabledObjects map[LWM2MObjectType]Object
	registry       Registry
}

func (c *MockClient) AddObjectInstance(LWM2MObjectType, int) error {
	return nil
}
func (c *MockClient) AddObjectInstances(LWM2MObjectType, ...int) {}
func (c *MockClient) AddResource()                               {}
func (c *MockClient) AddObject()                                 {}
func (c *MockClient) Register(string) string {
	return ""
}
func (c *MockClient) Deregister() {}
func (c *MockClient) Update()     {}
func (c *MockClient) UseRegistry(r Registry) {
	c.registry = r
}

func (c *MockClient) EnableObject(t LWM2MObjectType, e ObjectEnabler) error {
	_, ok := c.enabledObjects[t]
	if !ok {
		c.enabledObjects[t] = NewMockObject(t, e, c.GetRegistry())

		return nil
	} else {
		return errors.New("Object already enabled")
	}
}

func (c *MockClient) SetEnabler(LWM2MObjectType, ObjectEnabler) {}
func (c *MockClient) GetRegistry() Registry {
	return c.registry
}
func (c *MockClient) GetEnabledObjects() map[LWM2MObjectType]Object {
	return c.enabledObjects
}

func (c *MockClient) Start()                          {}
func (c *MockClient) OnStartup(FnOnStartup)           {}
func (c *MockClient) OnRead(FnOnRead)                 {}
func (c *MockClient) OnWrite(FnOnWrite)               {}
func (c *MockClient) OnExecute(FnOnExecute)           {}
func (c *MockClient) OnRegistered(FnOnRegistered)     {}
func (c *MockClient) OnDeregistered(FnOnDeregistered) {}
func (c *MockClient) OnError(FnOnError)               {}