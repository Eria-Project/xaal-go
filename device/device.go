package device

// Device : xAAL internal attributes for a device
type Device struct {
	devType string // xaal devtype
	address string // xaal addr

	/*
		self.devtype = devtype          # xaal devtype
		self.address = addr             # xaal addr
		self.vendor_id = None           # vendor ID ie : ACME
		self.product_id = None          # product ID
		self.version = None             # product release
		self.url = None                 # product URL
		self.info = None                # additionnal info
		self.hw_id = None               # hardware info
		self.group_id = None            # group devices
		# Some useless attributes, only for compatibility
		self.bus_addr = config.address
		self.bus_port = config.port
		self.hops = config.hops
		# Unsupported stuffs
		self.unsupported_attributes = []
		self.unsupported_methods = []
		self.unsupported_notifications = []
		# Alive management
		self.alive_period = config.alive_timer # time in sec between two alive
		self.next_alive = 0
		# Default attributes & methods
		self.__attributes = Attributes()
		self.methods = {'getAttributes' : self._get_attributes,
						'getDescription': self._get_description }
		self.engine = engine
	*/
}

/*DevType : Return device devType */
func (d Device) DevType() string {
	return d.devType
}

/*SetDevType : Set the device devType */
func (d Device) SetDevType(devType string) {
	d.devType = devType
}

/*Address : Return device address */
func (d Device) Address() string {
	return d.address
}

/*SetAddress : Set the device address */
func (d Device) SetAddress(address string) {
	d.address = address
}
