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
	"xaal-go/message"
	"xaal-go/tools"

	"golang.org/x/crypto/chacha20poly1305"
)

/*xAALMessage : xAAL JSON Message struct */
type xAALMessage struct {
	Version   string `json:"version"`
	Targets   string `json:"targets"`
	Timestamp []int  `json:"timestamp"`
	Payload   string `json:"payload"`
}

var _cipherKey []byte
var _config *config.XaalConfiguration

/*Init : Initialise with cipher key */
func Init(cipherKey string) {
	var err error
	_cipherKey, err = hex.DecodeString(cipherKey) // key encode / decode message built from passphrase
	if err != nil {
		log.Fatal(err)
	}

	_config = config.GetConfig()
}

/*DecodeMsg : Decode incoming Json data and De-Ciphering
:param data: data received from the multicast bus
:type data: json
:return: xAAL msg
:rtype: Message
*/
func DecodeMsg(data []byte) (*message.Message, error) {
	//	fmt.Println(string(data[:len(data)])) // FOR DEBUG
	// Decode json incoming data
	var dataRx = new(xAALMessage)

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
	if !tools.IsValidAddr(msg.Source()) {
		return nil, fmt.Errorf("Wrong message source [%s]", msg.Source())
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
