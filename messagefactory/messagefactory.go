/*Package messagefactory :
- Build xAAL message
- Apply security layer, Ciphering/De-Ciphering chacha20 poly1305
- Serialize/Deserialize data in JSON
*/
package messagefactory

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/project-eria/xaal-go/device"
	"github.com/project-eria/xaal-go/message"
	"github.com/project-eria/xaal-go/utils"

	logger "github.com/project-eria/eria-logger"

	"golang.org/x/crypto/chacha20poly1305"
)

var (
	_cipherKey    []byte
	_cipherWindow uint16
	_stackVersion string
)

// Init : Initialise and decode the cipher key
func Init(stackVersion string, key string, cipherWindow uint16) {
	_stackVersion = stackVersion
	_cipherWindow = cipherWindow
	var err error
	_cipherKey, err = hex.DecodeString(key) // key encode / decode message built from passphrase
	if err != nil {
		logger.Module("xaal:messagefactory").WithError(err).Fatal("Cannot decode cipher key")
	}
}

// EncodeMsg : Apply security layer and return encode MSG in Json
// :param msg: xAAL msg instance
// :type msg: Message
// :return: return an xAAL msg ciphered and serialized in json
// :rtype: json
func EncodeMsg(msg message.Message) ([]byte, error) {
	return encodeMsg(msg, nowFunc)
}

func encodeMsg(msg message.Message, nowTime func() time.Time) ([]byte, error) {
	var result message.DataMessage

	// Format data msg to send
	result.Version = msg.Version
	targets, err := json.Marshal(msg.Targets)
	if err != nil {
		return nil, fmt.Errorf("Cannot encode targets to JSON: %v", err)
	}
	result.Targets = string(targets)
	result.Timestamp = msg.Timestamp

	// Format payload before ciphering
	payload, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("Cannot encode payload to JSON: %v", err)
	}

	// Payload Ciphering: ciph
	ad := []byte(result.Targets) // Additionnal Data == json serialization of the targets array
	nonce, _ := buildNonce(msg.Timestamp)
	// chacha20 ciphering
	aead, err := chacha20poly1305.New(_cipherKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate ChaCha20-Poly1305: %v", err)
	}

	ciph := aead.Seal(nil, nonce, []byte(payload), ad)

	// Add payload: base64 encoded of payload cipher
	result.Payload = base64.StdEncoding.EncodeToString(ciph)

	// Json serialization
	message, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("Cannot encode message to JSON: %v", err)
	}
	return message, nil
}

// DecodeMsg : Decode incoming Json data and De-Ciphering
// :param data: data received from the multicast bus
// :type data: json
// :return: xAAL msg
// :rtype: Message
func DecodeMsg(data []byte) (*message.Message, error) {
	return decodeMsg(data, nowFunc)
}

func decodeMsg(data []byte, nowTime func() time.Time) (*message.Message, error) {
	//	fmt.Println(string(data[:len(data)])) // FOR DEBUG
	// Decode json incoming data
	var dataRx = new(message.DataMessage)
	err := json.Unmarshal(data, &dataRx)
	if err != nil {
		return nil, errors.New("Unable to parse JSON data")
	}
	// Instanciate Message
	var msg = new(message.Message)
	// TODO we should test if fields are available
	_ = json.Unmarshal([]byte(dataRx.Targets), &msg.Targets)
	msg.Version = dataRx.Version
	msg.Timestamp = dataRx.Timestamp
	var msgTime = dataRx.Timestamp[0]

	// Replay attack, window fixed to CIPHER_WINDOW in seconds
	now, _ := buildTimestamp(nowTime) // test done only on seconds ...

	if int64(msgTime) < (now - int64(_cipherWindow)) {
		return nil, fmt.Errorf("Potential replay attack, message too old: %d sec", now-int64(msgTime))
	}
	if int64(msgTime) > (now + int64(_cipherWindow)) {
		return nil, fmt.Errorf("Potential replay attack, message too young: %d sec", now-int64(msgTime))
	}

	// Payload De-Ciphering
	ad := []byte(dataRx.Targets) // Additional Data
	nonce, _ := buildNonce(dataRx.Timestamp)

	var ciph []byte

	// base64 decoding
	if dataRx.Payload != "" {
		var err error
		ciph, err = base64.StdEncoding.DecodeString(dataRx.Payload)
		if err != nil {
			return nil, fmt.Errorf("decode error: %v", err)
		}
	} else {
		return nil, errors.New("Bad message, no payload found")
	}

	// chacha20 deciphering
	aead, err := chacha20poly1305.New(_cipherKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate ChaCha20-Poly1305: %v", err)
	}
	decoded, err := aead.Open(nil, nonce, ciph, ad)
	if err != nil {
		return nil, fmt.Errorf("Failed to decrypt or authenticate message: %v", err)
	}

	msg.Raw = string(decoded) // TO REMOVE

	// Unpack Json
	if err := json.Unmarshal(decoded, &msg); err != nil {
		return nil, fmt.Errorf("Unable to parse JSON data in payload after decrypt: %v", err)
	}

	// Sanity check incomming message
	if !utils.IsValidAddr(msg.Header.Source) {
		return nil, fmt.Errorf("Wrong message source [%s]", msg.Header.Source)
	}
	return msg, nil
}

/*buildNonce : pack time using Big-Endian, time in seconds and time in microseconds */
func buildNonce(data []int) ([]byte, error) {
	if data == nil {
		return nil, errors.New("Can't build nouce for empty data")
	}
	nonce := make([]byte, 12)
	binary.BigEndian.PutUint64(nonce[0:], uint64(data[0]))
	binary.BigEndian.PutUint32(nonce[8:], uint32(data[1]))
	return nonce, nil
}

/*buildTimestamp : Return seconds since epoch, microseconds since last seconds Time = UTC+0000*/
func buildTimestamp(nowTime func() time.Time) (int64, int64) {
	now := nowTime()
	secs := now.Unix()
	micros := (now.UnixNano() / 1000) - (secs * 1000000)
	return secs, micros
}

func nowFunc() time.Time {
	return time.Now().UTC()
}

/*************
 MSG builder
**************/

// BuildMsg : the build method takes in parameters :
// -A device
// -The list of targets of the message
// -The type of the message
// -The action of the message
// -A body if it's necessary (None if not)
// it will return a message encoded in Json and Ciphered.
func BuildMsg(dev *device.Device, targets []string, msgtype string, action string, body map[string]interface{}) ([]byte, error) {
	return buildMsg(dev, targets, msgtype, action, body, nowFunc)
}
func buildMsg(dev *device.Device, targets []string, msgtype string, action string, body map[string]interface{}, nowTime func() time.Time) ([]byte, error) {
	msg := message.New(_stackVersion)
	msg.Header.Source = dev.Address
	msg.Header.DevType = dev.DevType
	msg.Targets = targets
	secs, micros := buildTimestamp(nowTime)
	msg.Timestamp = []int{int(secs), int(micros)}

	if msgtype != "" {
		msg.Header.MsgType = msgtype
	}
	if action != "" {
		msg.Header.Action = action
	}
	if body != nil && len(body) > 0 {
		msg.Body = body
	}
	data, err := encodeMsg(msg, nowTime)
	if err != nil {
		return nil, fmt.Errorf("EncodeMsg error: %v", err)
	}
	return data, nil
}

// BuildAliveFor : Build Alive message for a given device
// timeout = 0 is the minimum value
func BuildAliveFor(dev *device.Device, timeout uint16) ([]byte, error) {
	body := make(map[string]interface{})
	body["timeout"] = timeout
	return BuildMsg(dev, []string{}, "notify", "alive", body)
}

// BuildErrorMsg : Build a Error message
func BuildErrorMsg(dev *device.Device, errcode int, description string) ([]byte, error) {
	body := make(map[string]interface{})
	body["code"] = errcode
	if description != "" {
		body["description"] = description
	}
	return BuildMsg(dev, []string{}, "notify", "error", body)
}
