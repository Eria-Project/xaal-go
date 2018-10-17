package network

import (
	"fmt"
	"net"

	"xaal-go/log"

	reuseport "github.com/kavu/go_reuseport"
	"golang.org/x/net/ipv4"
)

var _address string
var _port uint16
var _hops uint8
var _stateConnected bool
var _conn net.PacketConn
var _pc *ipv4.PacketConn
var _dst *net.UDPAddr

/*Init : init the network */
func Init(address string, port uint16, hops uint8) {
	_address = address
	_port = port
	_hops = hops
	_stateConnected = false
}

/*Connect : connect the network */
func Connect() {
	log.Info("Connecting...", log.Fields{"-module": "network", "addr": "0.0.0.0", "port": _port})

	// open socket (connection)
	context := fmt.Sprintf("0.0.0.0:%d", _port)
	_conn, err := reuseport.ListenPacket("udp4", context)
	if err != nil {
		log.Fatal("Cannot open UDP4 socket", log.Fields{"-module": "network", "addr": "0.0.0.0", "port": _port, "err": err})
	}
	log.Info("Connected", log.Fields{"-module": "network", "addr": "0.0.0.0", "port": _port})

	// join multicast address
	log.Info("Joining Multicast Group...", log.Fields{"-module": "network", "multicastaddr": _address})
	group := net.ParseIP(_address)
	_pc = ipv4.NewPacketConn(_conn)
	_dst = &net.UDPAddr{IP: group, Port: int(_port)} // Set the destination address
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagMulticast == 0 {
			continue // not multicast interface
		}
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagBroadcast != 0 { // Loopback or Broadcast
			addrs, err := iface.Addrs()
			if err != nil {
				log.Fatal("Cannot list iface addresses", log.Fields{"-module": "network", "err": err})
			}
			if len(addrs) > 0 {
				for _, addr := range addrs {
					var ip net.IP
					switch v := addr.(type) {
					case *net.IPAddr:
						ip = v.IP
					case *net.IPNet:
						ip = v.IP
					}
					if ip == nil {
						continue
					}
					ip = ip.To4()
					if ip == nil {
						continue // not an ipv4 address
					}
					if err := _pc.JoinGroup(&iface, _dst); err != nil {
						log.Warn("Cannot join multicat group", log.Fields{"-module": "network", "iface": iface.Name, "err": err})
					}
					log.Info("Joined Multicast group", log.Fields{"-module": "network", "iface": iface.Name, "ip": ip})
				}
			}
		}
	}

	if err := _pc.SetControlMessage(ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
		_conn.Close()
		log.Fatal("Cannot set connection flags", log.Fields{"-module": "network", "err": err})
	}
	//	_pc.SetTTL(128)
	_stateConnected = true
}

/*Disconnect : Disconnect the network */
func Disconnect() {
	log.Info("Disconnecting socket", log.Fields{"-module": "network"})
	_stateConnected = false
	_conn.Close()
}

/*IsConnected : Check is the network is connected */
func IsConnected() bool {
	return _stateConnected
}

func receive() ([]byte, error) {
	log.Debug("UDP: reading bytes...", log.Fields{"-module": "network"})
	packt := make([]byte, 10000)
	n, cm, _, err := _pc.ReadFrom(packt)
	if err != nil {
		return nil, fmt.Errorf("UDP: ReadFrom: error %v", err)
	}
	// make a copy because we will overwrite buf
	b := make([]byte, n)
	copy(b, packt)
	log.Debug("UDP: recv bytes", log.Fields{"-module": "network", "size": n, "from": cm.Src, "to": cm.Dst})

	return packt[:n], nil // We resize packt to the received lenght
}

func send(data []byte) error {
	if _, err := _pc.WriteTo(data, nil, _dst); err != nil {
		return fmt.Errorf("UDP: WriteTo: error %v", err)
	}
	return nil
}

/*GetData : TODO DESCRIPTION */
func GetData() []byte {
	data, err := receive()
	if err != nil {
		log.Error("Cannot receive data", log.Fields{"-module": "network", "err": err})
	}
	return data
}

/*SendData : TODO DESCRIPTION */
func SendData(data []byte) {
	err := send(data)
	if err != nil {
		log.Fatal("Cannot send data", log.Fields{"-module": "network", "err": err})
	}
}
