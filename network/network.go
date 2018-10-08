package network

import (
	"fmt"
	"log"
	"net"

	reuseport "github.com/kavu/go_reuseport"
	"golang.org/x/net/ipv4"
)

var _ifaceName string
var _address string
var _port uint16
var _hops uint8
var _stateConnected bool
var _conn net.PacketConn
var _pc *ipv4.PacketConn
var _dst *net.UDPAddr

/*Init : init the network */
func Init(ifaceName string, address string, port uint16, hops uint8) {
	_ifaceName = ifaceName
	_address = address
	_port = port
	_hops = hops
	_stateConnected = false
}

/*Connect : connect the network */
func Connect() {
	iface, err := net.InterfaceByName(_ifaceName)
	if err != nil {
		log.Fatal(err)
	}

	context := fmt.Sprintf("0.0.0.0:%d", _port)

	log.Printf("Connecting to %s on %s\n", context, _ifaceName)

	// open socket (connection)
	_conn, err := reuseport.ListenPacket("udp4", context)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to %s on %s\n", context, _ifaceName)

	// join multicast address
	log.Printf("Join Multicast Group %s\n", _address)
	group := net.ParseIP(_address)
	_pc = ipv4.NewPacketConn(_conn)
	_dst = &net.UDPAddr{IP: group, Port: int(_port)} // Set the destination address
	if err := _pc.JoinGroup(iface, _dst); err != nil {
		_conn.Close()
		log.Fatal(err)
	}

	if err := _pc.SetControlMessage(ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
		_conn.Close()
		log.Fatal(err)
	}
	//	_pc.SetTTL(128)

	/*
		self.__sock = socket.socket(socket.AF_INET,socket.SOCK_DGRAM,socket.IPPROTO_UDP)
		self.__sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
		self.__sock.bind((self.bind_addr, self.port))
		mreq = struct.pack('4sl',socket.inet_aton(self.addr),socket.INADDR_ANY)
		self.__sock.setsockopt(socket.IPPROTO_IP,socket.IP_ADD_MEMBERSHIP,mreq)
		self.__sock.setsockopt(socket.IPPROTO_IP,socket.IP_MULTICAST_TTL,self.hops)
	*/
	_stateConnected = true
}

/*Disconnect : Disconnect the network */
func Disconnect() {
	log.Println("Disconnecting from the bus")
	_stateConnected = false
	_conn.Close()
}

/*IsConnected : Check is the network is connected */
func IsConnected() bool {
	return _stateConnected
}

func receive() ([]byte, error) {
	log.Printf("UDP: reading from '%s' on '%s'", _address, _ifaceName)
	packt := make([]byte, 10000)
	n, cm, _, err := _pc.ReadFrom(packt)
	if err != nil {
		return nil, fmt.Errorf("UDP: ReadFrom: error %v", err)
	}
	// make a copy because we will overwrite buf
	b := make([]byte, n)
	copy(b, packt)
	log.Printf("UDP: recv %d bytes from %s to %s on %s", n, cm.Src, cm.Dst, _ifaceName)

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
		log.Fatal(err)
	}
	return data
}

/*SendData : TODO DESCRIPTION */
func SendData(data []byte) {
	err := send(data)
	if err != nil {
		log.Fatal(err)
	}
}
