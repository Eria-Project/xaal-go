// Package message handle the message struct
package message

import (
	"fmt"
	"log"
)

// Message : xAAL Message struct
type Message struct {
	Body      map[string]interface{} `json:"body"`   // message body
	Header    messageHeader          `json:"header"` // dict used to store msg headers
	Targets   []string
	Timestamp []int  // message timestamp
	Version   string // = config.STACK_VERSION  // message API version
}

// messageHeader : xAAL Message Header struct
type messageHeader struct {
	Source  string `json:"source"`
	DevType string `json:"devType"`
	MsgType string `json:"msgType"`
	Action  string `json:"action"`
}

/*Dump : dump log a message */
func (m Message) Dump() {
	log.Printf("== Message (%p) ======================\n", &m)
	log.Println("*****Header*****")
	if m.Header.DevType != "" {
		log.Printf("devType \t%s\n", m.DevType())
	}
	if m.Header.Action != "" {
		log.Printf("action: \t%s\n", m.Action())
	}
	if m.Header.MsgType != "" {
		log.Printf("msgType: \t%s\n", m.MsgType())
	}
	if m.Header.Source != "" {
		log.Printf("source: \t%s\n", m.Source())
		log.Printf("version: \t%s\n", m.Version)
		log.Printf("targets: \t%s\n", m.Targets)
	}
	if len(m.Body) > 0 {
		log.Println("*****Body*****")
		for k, v := range m.Body {
			value := fmt.Sprint(v)
			log.Printf("%s: \t%s\n", k, value)
		}
	}
}

/*Source : Return message header source */
func (m Message) Source() string {
	return m.Header.Source
}

/*SetSource : Set the message header source */
func (m Message) SetSource(source string) {
	m.Header.Source = source
}

/*Action : Return message header action */
func (m Message) Action() string {
	return m.Header.Action
}

/*SetAction : Set the message header action */
func (m Message) SetAction(action string) {
	m.Header.Action = action
}

/*MsgType : Return message header msgType */
func (m Message) MsgType() string {
	return m.Header.MsgType
}

/*SetMsgType : Set the message header msgType */
func (m Message) SetMsgType(msgType string) {
	m.Header.MsgType = msgType
}

/*DevType : Return message header devType */
func (m Message) DevType() string {
	return m.Header.DevType
}

/*SetDevType : Set the message header devType */
func (m Message) SetDevType(devType string) {
	m.Header.DevType = devType
}

/*IsRequest : Return true if the message is a request */
func (m Message) IsRequest() bool {
	return m.Header.MsgType == "request"
}

/*IsReply : Return true if the message is a reply */
func (m Message) IsReply() bool {
	return m.Header.MsgType == "reply"
}

/*IsNotify : Return true if the message is a notification */
func (m Message) IsNotify() bool {
	return m.Header.MsgType == "notify"
}

/*IsAlive : Return true if the message is a alive notification */
func (m Message) IsAlive() bool {
	return m.Header.MsgType == "notify" && m.Header.Action == "alive"
}

/*IsAttributesChange : Return true if the message is a attributesChange notification */
func (m Message) IsAttributesChange() bool {
	return m.Header.MsgType == "notify" && m.Header.Action == "attributesChange"
}

/*IsGetAttributeReply : Return true if the message is a getAttributes reply */
func (m Message) IsGetAttributeReply() bool {
	return m.Header.MsgType == "reply" && m.Header.Action == "getAttributes"
}

/*IsGetDescriptionReply : Return true if the message is a getDescription reply */
func (m Message) IsGetDescriptionReply() bool {
	return m.Header.MsgType == "reply" && m.Header.Action == "getDescription"
}
