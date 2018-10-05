package network

import (
	"errors"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/ipv4"
)

var _ifaceName string
var _address string
var _port uint16
var _hops uint8
var _bindAddr = "0.0.0.0"
var _stateConnected bool
var _conn net.PacketConn
var _pc *ipv4.PacketConn

/*Init : init the network */
func Init(ifaceName string, address string, port uint16, hops uint8, bindAddr string) {
	_ifaceName = ifaceName
	_address = address
	_port = port
	_hops = hops
	if bindAddr != "" {
		_bindAddr = bindAddr
	}

	_stateConnected = false
}

/*Connect : connect the network */
func Connect() {
	iface, err := net.InterfaceByName(_ifaceName)
	if err != nil {
		log.Fatal(err)
	}

	context := fmt.Sprintf("%s:%d", _bindAddr, _port)

	log.Printf("Connecting to %s on %s\n", context, _ifaceName)

	// open socket (connection)
	_conn, err := net.ListenPacket("udp", context)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to %s on %s\n", context, _ifaceName)

	// join multicast address
	log.Printf("Join Multicast Group %s\n", _address)
	group := net.ParseIP(_address)
	_pc = ipv4.NewPacketConn(_conn)
	if err := _pc.JoinGroup(iface, &net.UDPAddr{IP: group, Port: int(_port)}); err != nil {
		_conn.Close()
		log.Fatal(err)
	}

	if err := _pc.SetControlMessage(ipv4.FlagDst, true); err != nil {
		_conn.Close()
		log.Fatal(err)
	}
	// TODO	_pc.SetTTL(128)

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
	log.Printf("udpReader: reading from '%s' on '%s'", _address, _ifaceName)
	packt := make([]byte, 10000)
	n, cm, _, err := _pc.ReadFrom(packt)
	if err != nil {
		e := fmt.Sprintf("udpReader: ReadFrom: error %v", err)
		return nil, errors.New(e)
	}
	// make a copy because we will overwrite buf
	b := make([]byte, n)
	copy(b, packt)
	log.Printf("udpReader: recv %d bytes from %s to %s on %s", n, cm.Src, cm.Dst, _ifaceName)

	return packt[:n], nil // We resize packt to the received lenght
}

/*GetData : TODO DESCRIPTION */
func GetData() []byte {
	if !IsConnected() {
		Connect()
	}

	data, err := receive()
	if err != nil {
		log.Fatal(err)
	}
	return data
}
