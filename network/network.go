package network

import (
	"fmt"
	"net"

	"github.com/Eria-Project/logger"

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
	logger.Module("network").WithFields(logger.Fields{"addr": "0.0.0.0", "port": _port}).Info("Connecting...")

	// open socket (connection)
	context := fmt.Sprintf("0.0.0.0:%d", _port)
	_conn, err := reuseport.ListenPacket("udp4", context)
	if err != nil {
		logger.Module("network").WithError(err).WithFields(logger.Fields{"addr": "0.0.0.0", "port": _port}).Fatal("Cannot open UDP4 socket")
	}
	logger.Module("network").WithFields(logger.Fields{"addr": "0.0.0.0", "port": _port}).Info("Connected")

	// join multicast address
	logger.Module("network").WithField("multicastaddr", _address).Info("Joining Multicast Group...")
	group := net.ParseIP(_address)
	_pc = ipv4.NewPacketConn(_conn)
	_dst = &net.UDPAddr{IP: group, Port: int(_port)} // Set the destination address
	ifaces := getIPv4Interfaces()
	for _, iface := range ifaces {
		if err := _pc.JoinGroup(iface, _dst); err != nil {
			logger.Module("network").WithError(err).WithField("iface", iface.Name).Warn("Cannot join multicat group")
		}
		logger.Module("network").WithField("iface", iface.Name).Info("Joined Multicast group")
	}

	if err := _pc.SetControlMessage(ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
		_conn.Close()
		logger.Module("network").WithError(err).Fatal("Cannot set connection flags")
	}
	//	_pc.SetTTL(128)
	_stateConnected = true
}

func getIPv4Interfaces() map[string]*net.Interface {
	candidateInterfaces := map[string]*net.Interface{}
	ifaces, _ := net.Interfaces()
	for i, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagMulticast == 0 {
			continue // not multicast interface
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // not loopback interface
		}
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagBroadcast != 0 { // Loopback or Broadcast
			addrs, err := iface.Addrs()
			if err != nil {
				logger.Module("network").WithError(err).Fatal("Cannot list iface addresses")
			}
			if len(addrs) > 0 {
				for _, addr := range addrs {

					var ip net.IP
					switch v := addr.(type) {
					case *net.IPAddr:
					case *net.IPNet:
						ip = v.IP
						break
					}
					if ip == nil {
						continue
					}
					if ip.To4() == nil {
						continue // not an ipv4 address
					}
					logger.Module("network").WithFields(logger.Fields{"iface": iface.Name, "ip": ip.String()}).Info("Found interface")
					candidateInterfaces[iface.Name] = &(ifaces[i]) // https://blogger.omri.io/golang-sneaky-range-pointer/
				}
			}
		}
	}
	return candidateInterfaces
}

/*Disconnect : Disconnect the network */
func Disconnect() {
	logger.Module("network").Info("Disconnecting socket")
	_stateConnected = false
	_conn.Close()
}

/*IsConnected : Check is the network is connected */
func IsConnected() bool {
	return _stateConnected
}

func receive() ([]byte, error) {
	logger.Module("network").Trace("UDP: reading bytes...")
	packt := make([]byte, 10000)
	n, cm, _, err := _pc.ReadFrom(packt)
	if err != nil {
		return nil, fmt.Errorf("UDP: ReadFrom: error %v", err)
	}
	// make a copy because we will overwrite buf
	b := make([]byte, n)
	copy(b, packt)
	logger.Module("network").WithFields(logger.Fields{"size": n, "from": cm.Src, "to": cm.Dst}).Trace("UDP: recv bytes")

	return packt[:n], nil // We resize packt to the received lenght
}

func send(data []byte) error {
	n, err := _pc.WriteTo(data, nil, _dst)
	if err != nil {
		return fmt.Errorf("UDP: WriteTo: error %v", err)
	}
	logger.Module("network").WithFields(logger.Fields{"size": n, "to": _dst.IP}).Trace("UDP: send bytes")
	return nil
}

/*GetData : TODO DESCRIPTION */
func GetData() []byte {
	data, err := receive()
	if err != nil {
		logger.Module("network").WithError(err).Error("Cannot receive data")
	}
	return data
}

/*SendData : TODO DESCRIPTION */
func SendData(data []byte) {
	err := send(data)
	if err != nil {
		logger.Module("network").WithError(err).Fatal("Cannot send data")
	}
}
