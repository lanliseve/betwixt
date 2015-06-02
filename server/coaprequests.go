package server

import (
	. "github.com/zubairhamed/go-commons/network"
	"github.com/zubairhamed/goap"
	"log"
)

func SetupCoapRoutes(server *DefaultServer) {
	coap := server.coapServer

	coap.NewRoute("rd", goap.POST, handleRegister(server))
	coap.NewRoute("rd/{id}", goap.PUT, handleUpdate(server))
	coap.NewRoute("rd/{id}", goap.DELETE, handleDelete(server))
}

func handleRegister(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		req := r.(*goap.CoapRequest)
		ep := req.GetUriQuery("ep")

		clientId, err := server.register(ep, req.GetAddress().String())
		if err != nil {
			log.Println("Error registering client ", ep)
		}

		msg := goap.NewMessageOfType(goap.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.AddOption(goap.OPTION_LOCATION_PATH, "rd/"+clientId)
		msg.Code = goap.COAPCODE_201_CREATED

		return goap.NewResponseWithMessage(msg)
	}
}

func handleUpdate(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		req := r.(*goap.CoapRequest)
		id := req.GetAttribute("id")

		server.update(id)

		msg := goap.NewMessageOfType(goap.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.Code = goap.COAPCODE_204_CHANGED

		return goap.NewResponseWithMessage(msg)
	}
}

func handleDelete(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		req := r.(*goap.CoapRequest)
		id := req.GetAttribute("id")

		server.delete(id)

		msg := goap.NewMessageOfType(goap.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.Code = goap.COAPCODE_202_DELETED

		return goap.NewResponseWithMessage(msg)
	}
}