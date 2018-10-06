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
	"log"
	"time"
	"xaal-go/config"
	"xaal-go/device"
	"xaal-go/message"
	"xaal-go/tools"

	"golang.org/x/crypto/chacha20poly1305"
)

var _cipherKey []byte
var _config *config.XaalConfiguration

// Init : Initialise with cipher key
func Init(cipherKey string) {
	var err error
	_cipherKey, err = hex.DecodeString(cipherKey) // key encode / decode message built from passphrase
	if err != nil {
		log.Fatal(err)
	}

	_config = config.GetConfig()
}

// EncodeMsg : Apply security layer and return encode MSG in Json
// :param msg: xAAL msg instance
// :type msg: Message
// :return: return an xAAL msg ciphered and serialized in json
// :rtype: json
func EncodeMsg(msg message.Message) ([]byte, error) {
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
	nonce := buildNonce(msg.Timestamp)
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
	// TODO return codecs.encode(message)*/
	return message, nil
}

// DecodeMsg : Decode incoming Json data and De-Ciphering
// :param data: data received from the multicast bus
// :type data: json
// :return: xAAL msg
// :rtype: Message
func DecodeMsg(data []byte) (*message.Message, error) {
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
	now, _ := buildTimestamp() // test done only on seconds ...

	if int64(msgTime) < (now - int64(_config.CipherWindow)) {
		log.Fatalf("Potential replay attack, message too old: %d sec", now-int64(msgTime))
	}
	if int64(msgTime) > (now + int64(_config.CipherWindow)) {
		log.Fatalf("Potential replay attack, message too young: %d sec", now-int64(msgTime))
	}

	// Payload De-Ciphering
	ad := []byte(dataRx.Targets) // Additional Data
	nonce := buildNonce(dataRx.Timestamp)

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

	// Unpack Json
	if err := json.Unmarshal(decoded, &msg); err != nil {
		return nil, fmt.Errorf("Unable to parse JSON data in payload after decrypt: %v", err)
	}

	// Sanity check incomming message
	if !tools.IsValidAddr(msg.Header.Source) {
		return nil, fmt.Errorf("Wrong message source [%s]", msg.Header.Source)
	}

	return msg, nil
}

/*buildNonce : pack time using Big-Endian, time in seconds and time in microseconds */
func buildNonce(data []int) []byte {
	nonce := make([]byte, 12)
	binary.BigEndian.PutUint64(nonce[0:], uint64(data[0]))
	binary.BigEndian.PutUint32(nonce[8:], uint32(data[1]))
	return nonce
}

/*buildTimestamp : Return seconds since epoch, microseconds since last seconds Time = UTC+0000*/
func buildTimestamp() (int64, int64) {
	now := time.Now().UTC()
	secs := now.Unix()
	micros := (now.UnixNano() / 1000) - (secs * 1000000)
	return secs, micros
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
func BuildMsg(dev device.Device, targets []string, msgtype string, action string, body map[string]interface{}) []byte {
	message := message.New()
	message.Header.Source = dev.Address
	message.Header.DevType = dev.DevType
	message.Targets = targets
	secs, micros := buildTimestamp()
	message.Timestamp = []int{int(secs), int(micros)}

	if msgtype != "" {
		message.Header.MsgType = msgtype
	}
	if action != "" {
		message.Header.Action = action
	}
	if body != nil && len(body) > 0 {
		message.Body = body
	}
	data, err := EncodeMsg(message)
	if err != nil {
		log.Printf("EncodeMsg error: %v", err)
		return nil
	}
	return data
}

// BuildAliveFor : Build Alive message for a given device
// timeout = 0 is the minimum value
func BuildAliveFor(dev device.Device, timeout int) []byte {
	body := make(map[string]interface{})
	body["timeout"] = timeout
	message := BuildMsg(dev, []string{}, "notify", "alive", body)
	return message
}

// BuildErrorMsg : Build a Error message
func BuildErrorMsg(dev device.Device, errcode int, description string) []byte {
	body := make(map[string]interface{})
	body["code"] = errcode
	if description != "" {
		body["description"] = description
	}
	message := BuildMsg(dev, []string{}, "notify", "error", body)
	return message
}
