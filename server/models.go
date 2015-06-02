package server

import (
	"github.com/zubairhamed/go-lwm2m/api"
	"time"
)

func NewRegisteredClient(ep string, id string, addr string) api.RegisteredClient {
	return &DefaultRegisteredClient{
		name:    	ep,
		id:      	id,
		addr:    	addr,
		regDate: 	time.Now(),
		updateDate: time.Now(),
	}
}

type DefaultRegisteredClient struct {
	id          string
	name        string
	lifetime    int
	version     string
	bindingMode api.BindingMode
	smsNumber   string
	addr        string
	regDate     time.Time
	updateDate	time.Time
}

func (c *DefaultRegisteredClient) GetId() string {
	return c.id
}

func (c *DefaultRegisteredClient) GetName() string {
	return c.name
}

func (c *DefaultRegisteredClient) GetLifetime() int {
	return c.lifetime
}

func (c *DefaultRegisteredClient) GetVersion() string {
	return c.version
}

func (c *DefaultRegisteredClient) GetBindingMode() api.BindingMode {
	return c.bindingMode
}

func (c *DefaultRegisteredClient) GetSmsNumber() string {
	return c.smsNumber
}

func (c *DefaultRegisteredClient) GetRegistrationDate() time.Time {
	return c.regDate
}

func (c *DefaultRegisteredClient) Update() {
	c.updateDate = time.Now()
}

func (c *DefaultRegisteredClient) LastUpdate() time.Time {
	return c.updateDate
}