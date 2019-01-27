// Package message handle the message struct
package message

import (
	"fmt"
	"strings"
	"time"
)

// DataMessage : xAAL JSON Message struct
type DataMessage struct {
	Version   string `json:"version"`
	Targets   string `json:"targets"`
	Timestamp []int  `json:"timestamp"`
	Payload   string `json:"payload"`
}

// Message : xAAL Message struct
type Message struct {
	Body      map[string]interface{} `json:"body,omitempty"` // message body
	Header    Header                 `json:"header"`         // dict used to store msg headers
	Targets   []string               `json:"-"`
	Timestamp []int                  `json:"-"` // message timestamp
	Version   string                 `json:"-"` // message API version
	Raw       string
}

// Header : xAAL Message Header struct
type Header struct {
	Source  string `json:"source"`
	DevType string `json:"devType"`
	MsgType string `json:"msgType"`
	Action  string `json:"action"`
}

// New : Initiate a new Message struct
func New(stackVersion string) Message {
	return Message{Version: stackVersion, Targets: []string{}}
}

// Dump : dump log a message
func (m Message) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "\n== Message (%p) ======================\n", &m)
	fmt.Fprintln(&b, m.Time())
	fmt.Fprintln(&b, "*****Header*****")
	if m.Header.DevType != "" {
		fmt.Fprintf(&b, "devType \t%s\n", m.Header.DevType)
	}
	if m.Header.Action != "" {
		fmt.Fprintf(&b, "action: \t%s\n", m.Header.Action)
	}
	if m.Header.MsgType != "" {
		fmt.Fprintf(&b, "msgType: \t%s\n", m.Header.MsgType)
	}
	if m.Header.Source != "" {
		fmt.Fprintf(&b, "source: \t%s\n", m.Header.Source)
		fmt.Fprintf(&b, "version: \t%s\n", m.Version)
		fmt.Fprintf(&b, "targets: \t%s\n", m.Targets)
	}
	if len(m.Body) > 0 {
		fmt.Fprintln(&b, "*****Body*****")
		for k, v := range m.Body {
			value := fmt.Sprint(v)
			fmt.Fprintf(&b, "%s: \t%s\n", k, value)
		}
	}
	return b.String()
}

/*IsRequest : Return true if the message is a request */
func (m *Message) IsRequest() bool {
	return m.Header.MsgType == "request"
}

/*IsReply : Return true if the message is a reply */
func (m *Message) IsReply() bool {
	return m.Header.MsgType == "reply"
}

/*IsNotify : Return true if the message is a notification */
func (m *Message) IsNotify() bool {
	return m.Header.MsgType == "notify"
}

/*IsAlive : Return true if the message is a alive notification */
func (m *Message) IsAlive() bool {
	return m.Header.MsgType == "notify" && m.Header.Action == "alive"
}

/*IsAttributesChange : Return true if the message is a attributesChange notification */
func (m *Message) IsAttributesChange() bool {
	return m.Header.MsgType == "notify" && m.Header.Action == "attributesChange"
}

/*IsGetAttributeReply : Return true if the message is a getAttributes reply */
func (m *Message) IsGetAttributeReply() bool {
	return m.Header.MsgType == "reply" && m.Header.Action == "getAttributes"
}

/*IsGetDescriptionReply : Return true if the message is a getDescription reply */
func (m *Message) IsGetDescriptionReply() bool {
	return m.Header.MsgType == "reply" && m.Header.Action == "getDescription"
}

// Time : Return timestamp in Time format
func (m *Message) Time() time.Time {
	sec := int64(m.Timestamp[0])
	nsec := int64(m.Timestamp[1] * 1000)
	return time.Unix(sec, nsec)
}

/*
def _get_parameters(self, args):
	parameters = {}
	for p in range(0, len(args)):
		if isinstance(self.body[args[p]], str):
			param = "'%s'" % self.body[args[p]]
		else:
			param = self.body[args[p]]

		parameters[args[p]] = param
	return parameters
*/

// GetParameters : request parameters are in body hash, return asis
/* TO REMOVE
func (m *Message) GetParameters() []string {
	keys := make([]string, 0, len(m.Body))
	for k := range m.Body {
		keys = append(keys, k)
	}
	return keys
}
*/
