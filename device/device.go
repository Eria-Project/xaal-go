package device

// Device : xAAL internal attributes for a device
type Device struct {
	DevType     string // xaal devtype
	Address     string // xaal addr
	alivePeriod int    // time in sec between two alive

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
		# Default attributes & methods
		self.__attributes = Attributes()
		self.methods = {'getAttributes' : self._get_attributes,
						'getDescription': self._get_description }
		self.engine = engine
	*/
}

/*SetDevType : Set the device devType */
func (d Device) SetDevType(devType string) {
	d.DevType = devType
}

/*SetAddress : Set the device address */
func (d Device) SetAddress(address string) error {
	if d.Address == "" {
		d.Address = ""
		return nil
	}
	/* TODO
	   if not tools.is_valid_addr(value):
	       raise DeviceError("This address is not valid")
	   if value == config.XAAL_BCAST_ADDR:
	       raise DeviceError("This address is reserved")
	*/
	d.Address = address
	return nil
}

// GetTimeout : return Alive timeout used for isAlive msg
func (d Device) GetTimeout() int {
	return 2 * d.alivePeriod
}
