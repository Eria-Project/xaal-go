package network

import (
	"fmt"
	"net"
	"time"

	"github.com/project-eria/logger"

	reuseport "github.com/libp2p/go-reuseport"
	"golang.org/x/net/ipv4"
)

var _address string
var _port uint16
var _hops uint8
var _stateConnected bool
var _conn net.PacketConn
var _pc *ipv4.PacketConn
var _dst *net.UDPAddr
var _ifaces map[int]*net.Interface

/*Init : init the network */
func Init(address string, port uint16, hops uint8) {
	_address = address
	_port = port
	_hops = hops
	_stateConnected = false
}

/*Connect : connect the network */
func Connect() {
	logger.Module("xaal:network").WithFields(logger.Fields{"addr": _address, "port": _port}).Info("Connecting...")

	// open socket (connection)
	context := fmt.Sprintf("%s:%d", _address, _port)
	_conn, err := reuseport.ListenPacket("udp4", context)
	if err != nil {
		logger.Module("xaal:network").WithError(err).WithFields(logger.Fields{"addr": _address, "port": _port}).Fatal("Cannot open UDP4 socket")
	}
	logger.Module("xaal:network").WithFields(logger.Fields{"addr": _address, "port": _port}).Info("Connected")

	// join multicast address
	logger.Module("xaal:network").WithField("multicastaddr", _address).Info("Joining Multicast Group...")
	_pc = ipv4.NewPacketConn(_conn)

	// Set Multicat on Loopback
	if loop, err := _pc.MulticastLoopback(); err == nil {
		logger.Module("xaal:network").WithField("status", loop).Debug("MulticastLoopback")
		if !loop {
			if err := _pc.SetMulticastLoopback(true); err != nil {
				logger.Module("xaal:network").WithError(err).Warn("SetMulticastLoopback")
			}
		}
	}
	group := net.ParseIP(_address)
	_dst = &net.UDPAddr{IP: group, Port: int(_port)} // Set the destination address

	_ifaces = getIPv4Interfaces()
	for _, iface := range _ifaces {
		if err := _pc.JoinGroup(iface, _dst); err != nil {
			logger.Module("xaal:network").WithError(err).WithField("iface", iface.Name).Warn("Cannot join multicat group")
		}
		logger.Module("xaal:network").WithField("iface", iface.Name).Info("Joined Multicast group")
	}

	if err := _pc.SetControlMessage(ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
		_conn.Close()
		logger.Module("xaal:network").WithError(err).Fatal("Cannot set connection flags")
	}
	//	_pc.SetTTL(128)
	_stateConnected = true

	go func() {
		tick := time.Tick(1 * time.Second)
		// Keep trying until we're timed out or got a result or got an error
		for {
			select {
			case <-tick:
				//				logger.Module("xaal:network").Tracef("Check network => %+v %+v", _pc, _pc.PacketConn)
			}
		}
	}()
}

func getIPv4Interfaces() map[int]*net.Interface {
	candidateInterfaces := map[int]*net.Interface{}
	ifaces, _ := net.Interfaces()
	for i, iface := range ifaces {
		logger.Module("xaal:network").WithFields(logger.Fields{"iface": iface.Name, "flags": iface.Flags}).Trace("Interface")

		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagMulticast != 0 { // Loopback or Broadcast
			addrs, err := iface.Addrs()
			if err != nil {
				logger.Module("xaal:network").WithError(err).Fatal("Cannot list iface addresses")
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
					logger.Module("xaal:network").WithFields(logger.Fields{"iface": iface.Name, "ip": ip.String()}).Info("Found interface")
					candidateInterfaces[iface.Index] = &(ifaces[i]) // https://blogger.omri.io/golang-sneaky-range-pointer/
				}
			}
		}
	}
	return candidateInterfaces
}

/*Disconnect : Disconnect the network */
func Disconnect() {
	logger.Module("xaal:network").Info("Disconnecting socket")
	_stateConnected = false
	_conn.Close()
}

/*IsConnected : Check is the network is connected */
func IsConnected() bool {
	return _stateConnected
}

func receive() ([]byte, error) {
	logger.Module("xaal:network").Trace("UDP: reading bytes...")
	packt := make([]byte, 10000)
	n, cm, _, err := _pc.ReadFrom(packt)
	if err != nil {
		return nil, fmt.Errorf("UDP: ReadFrom: error %v", err)
	}
	// make a copy because we will overwrite buf
	b := make([]byte, n)
	copy(b, packt)
	logger.Module("xaal:network").WithFields(logger.Fields{"size": n, "from": cm.Src, "via": _ifaces[cm.IfIndex].Name, "to": cm.Dst}).Trace("UDP: recv bytes")

	return packt[:n], nil // We resize packt to the received lenght
}

func send(data []byte) error {
	dst := _dst
	var (
		n   int
		err error
	)
	n, err = _pc.WriteTo(data, nil, dst)
	if err != nil {
		logger.Module("xaal:network").WithError(err).Error("UDP: Error writing connection, trying localhost")
		dst = &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: int(_port)}
		n, err = _pc.WriteTo(data, nil, dst)
		if err != nil {
			return fmt.Errorf("UDP: WriteTo: error %v", err)
		}
	}
	logger.Module("xaal:network").WithFields(logger.Fields{"size": n, "to": dst.IP, "via": _pc.PacketConn.LocalAddr()}).Trace("UDP: send bytes")
	return nil
}

/*GetData : TODO DESCRIPTION */
func GetData() []byte {
	data, err := receive()
	if err != nil {
		logger.Module("xaal:network").WithError(err).Error("Cannot receive data")
	}
	return data
}

/*SendData : TODO DESCRIPTION */
func SendData(data []byte) {
	err := send(data)
	if err != nil {
		logger.Module("xaal:network").WithError(err).Fatal("Cannot send data")
	}
}
