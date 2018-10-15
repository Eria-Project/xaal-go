// Package engine : xAAL messages tx handlers
package engine

import (
	"log"
	"xaal-go/device"
	"xaal-go/messagefactory"
	"xaal-go/network"
)

var _queueMsgTx = make(chan []byte)

// processTxMsg : Process (send) message in tx queue chan called from other tasks
func processTxMsg() {
	for msg := range _queueMsgTx {
		network.SendData(msg)
	}
}

/*
def send_request(self,dev,targets,action,body = None):
"""queue a new request"""
msg = self.msg_factory.build_msg(dev, targets, 'request', action, body)
self.queue_msg(msg)

def send_reply(self, dev, targets, action, body=None):
"""queue a new reply"""
msg = self.msg_factory.build_msg(dev, targets, 'reply', action, body)
self.queue_msg(msg)

def send_error(self, dev, errcode, description=None):
"""queue a error message"""
msg = self.msg_factory.build_error_msg(dev, errcode, description)
self.queue_msg(msg)

def send_get_description(self,dev,targets):
"""queue a getDescription request"""
self.send_request(dev,targets,'getDescription')

def send_get_attributes(self,dev,targets):
"""queue a getAttributes request"""
self.send_request(dev,targets,'getAttributes')

*/

func sendNotification(dev *device.Device, action string, body map[string]interface{}) {
	msg, err := messagefactory.BuildMsg(dev, []string{}, "notify", action, body)
	if err != nil {
		log.Println("Error while building message")
	}
	_queueMsgTx <- msg
}
