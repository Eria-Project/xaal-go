package schemas

import (
	"fmt"
	"strings"

	"github.com/project-eria/xaal-go/device"
)

// DeviceFromType return a device created using a string type schema
func DeviceFromType(schema string, addr string) (*device.Device, error) {
	schema = strings.ToLower(schema)
	switch schema {
	case "lamp.basic":
		return Lamp(addr), nil
	case "shutter.basic":
		return Shutter(addr), nil
	case "shutter.position":
		return ShutterPosition(addr), nil
	}
	return nil, fmt.Errorf("Device type '%s' not implemented", schema)
}

/*
#=====================================================================
def audiomixer(addr=None):
    """Simple audio mixer with only one general volume control"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('audiomixer.basic',addr)

    # -- Attributes --
    # Address of the device that is the audio source
    dev.new_attribute('source')
    # General volume
    dev.new_attribute('volume')

    # -- Methods --
    def default_setSource(_source):
        """Set the device that is the audio source"""
        logger.info("default_setSource(source=[%s],)" % (_source))

    def default_up():
        """Turn up the volume"""
        logger.info("default_up()")

    def default_down():
        """Turn down the volume"""
        logger.info("default_down()")

    def default_mute():
        """Mute the volume"""
        logger.info("default_mute()")

    def default_unmute():
        """Unmute the volume"""
        logger.info("default_unmute()")

    dev.add_method('setSource',default_setSource)
    dev.add_method('up',default_up)
    dev.add_method('down',default_down)
    dev.add_method('mute',default_mute)
    dev.add_method('unmute',default_unmute)

    return dev

#=====================================================================
def battery(addr=None):
    """Report on state of a battery"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('battery.basic',addr)

    # -- Attributes --
    # Battery state
    dev.new_attribute('level')
    # List of devices concerned with this battery
    dev.new_attribute('devices')
    return dev

#=====================================================================
def cache(addr=None):
    """Simple cache that can be queried about attributes of devices"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('cache.basic',addr)

    # -- Methods --
    def default_getDeviceAttribute(_device,_attribute,_value,_date):
        """Get the value of a device's attribute"""
        logger.info("default_getDeviceAttribute(device=[%s],attribute=[%s],value=[%s],date=[%s],)" % (_device,_attribute,_value,_date))

    def default_getDeviceAttributes(_device,_attributes):
        """Get all known atributes of a device"""
        logger.info("default_getDeviceAttributes(device=[%s],attributes=[%s],)" % (_device,_attributes))

    dev.add_method('getDeviceAttribute',default_getDeviceAttribute)
    dev.add_method('getDeviceAttributes',default_getDeviceAttributes)

    return dev

#=====================================================================
def co2meter(addr=None):
    """Simple CO2 meter"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('co2meter.basic',addr)

    # -- Attributes --
    # CO2
    dev.new_attribute('co2')
    return dev

#=====================================================================
def door(addr=None):
    """Simple door device"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('door.basic',addr)

    # -- Attributes --
    # Position of the door
    dev.new_attribute('position')

    # -- Methods --
    def default_open():
        """Open the door"""
        logger.info("default_open()")

    def default_close():
        """Close the door"""
        logger.info("default_close()")

    dev.add_method('open',default_open)
    dev.add_method('close',default_close)

    return dev

#=====================================================================
def falldetector(addr=None):
    """Simple fall detection device"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('falldetector.basic',addr)

    # -- Attributes --
    # List of detected falls
    dev.new_attribute('falls')
    return dev

#=====================================================================
def hmi(addr=None):
    """Basic Human Machine Interface"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('hmi.basic',addr)
    return dev

#=====================================================================
def lamp_cie1931(addr=None):
    """Lamp color in cie1931 color space"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('lamp.cie1931',addr)

    # -- Attributes --
    # State of the lamp
    dev.new_attribute('light')
    # Level of the dimmer
    dev.new_attribute('dimmer')
    # xy coordinates in the cie1931 chromacity diagram
    dev.new_attribute('xy')
    # saturation
    dev.new_attribute('saturation')

    # -- Methods --
    def default_on():
        """Switch on the lamp"""
        logger.info("default_on()")

    def default_off():
        """Switch off the lamp"""
        logger.info("default_off()")

    def default_dim(_target):
        """Change the dimmer of the lamp"""
        logger.info("default_dim(target=[%s],)" % (_target))

    def default_xy(_target):
        """Set color coordinates in cie1931 chromacity diagram"""
        logger.info("default_xy(target=[%s],)" % (_target))

    def default_saturation(_target):
        """Set saturation level"""
        logger.info("default_saturation(target=[%s],)" % (_target))

    dev.add_method('on',default_on)
    dev.add_method('off',default_off)
    dev.add_method('dim',default_dim)
    dev.add_method('xy',default_xy)
    dev.add_method('saturation',default_saturation)

    return dev

#=====================================================================
def lamp_dimmer(addr=None):
    """Lamp with a dimmer"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('lamp.dimmer',addr)

    # -- Attributes --
    # State of the lamp
    dev.new_attribute('light')
    # Level of the dimmer
    dev.new_attribute('dimmer')

    # -- Methods --
    def default_on():
        """Switch on the lamp"""
        logger.info("default_on()")

    def default_off():
        """Switch off the lamp"""
        logger.info("default_off()")

    def default_dim(_target):
        """Change the dimmer of the lamp"""
        logger.info("default_dim(target=[%s],)" % (_target))

    dev.add_method('on',default_on)
    dev.add_method('off',default_off)
    dev.add_method('dim',default_dim)

    return dev

#=====================================================================
def linkquality(addr=None):
    """Report on quality of a transmission link"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('linkquality.basic',addr)

    # -- Attributes --
    # Link quality
    dev.new_attribute('level')
    # List of devices concerned with this link
    dev.new_attribute('devices')
    return dev

#=====================================================================
def luxmeter(addr=None):
    """Simple luxmeter"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('luxmeter.basic',addr)

    # -- Attributes --
    # Measure of how much luminous flux is spread over a given area
    dev.new_attribute('illuminance')
    return dev

#=====================================================================
def mediaplayer(addr=None):
    """Generic media player"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('mediaplayer.basic',addr)

    # -- Attributes --
    # Addresses of the rendering devices (audio mixer and/or display screen)
    dev.new_attribute('destination')
    # Activity of the player
    dev.new_attribute('activity')

    # -- Methods --
    def default_setDestination(_destination):
        """Set devices that make the rendering of the medias"""
        logger.info("default_setDestination(destination=[%s],)" % (_destination))

    def default_play():
        """Play"""
        logger.info("default_play()")

    def default_pause():
        """Pause"""
        logger.info("default_pause()")

    def default_stop():
        """Stop"""
        logger.info("default_stop()")

    def default_next():
        """Next track"""
        logger.info("default_next()")

    def default_prev():
        """Previous track"""
        logger.info("default_prev()")

    def default_ff():
        """Fast forward"""
        logger.info("default_ff()")

    def default_rw():
        """Rewind"""
        logger.info("default_rw()")

    dev.add_method('setDestination',default_setDestination)
    dev.add_method('play',default_play)
    dev.add_method('pause',default_pause)
    dev.add_method('stop',default_stop)
    dev.add_method('next',default_next)
    dev.add_method('prev',default_prev)
    dev.add_method('ff',default_ff)
    dev.add_method('rw',default_rw)

    return dev

#=====================================================================
def mediaplayer_spotify(addr=None):
    """A Spotify player"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('mediaplayer.spotify',addr)

    # -- Attributes --
    # Addresses of the rendering devices (audio mixer and/or display screen)
    dev.new_attribute('destination')
    # Activity of the player
    dev.new_attribute('activity')

    # -- Methods --
    def default_setDestination(_destination):
        """Set devices that make the rendering of the medias"""
        logger.info("default_setDestination(destination=[%s],)" % (_destination))

    def default_play():
        """Play"""
        logger.info("default_play()")

    def default_pause():
        """Pause"""
        logger.info("default_pause()")

    def default_stop():
        """Stop"""
        logger.info("default_stop()")

    def default_next():
        """Next track"""
        logger.info("default_next()")

    def default_prev():
        """Previous track"""
        logger.info("default_prev()")

    def default_ff():
        """Fast forward"""
        logger.info("default_ff()")

    def default_rw():
        """Rewind"""
        logger.info("default_rw()")

    def default_playUrl(_url):
        """Play an URL"""
        logger.info("default_playUrl(url=[%s],)" % (_url))

    dev.add_method('setDestination',default_setDestination)
    dev.add_method('play',default_play)
    dev.add_method('pause',default_pause)
    dev.add_method('stop',default_stop)
    dev.add_method('next',default_next)
    dev.add_method('prev',default_prev)
    dev.add_method('ff',default_ff)
    dev.add_method('rw',default_rw)
    dev.add_method('playUrl',default_playUrl)

    return dev

#=====================================================================
def metadatadb(addr=None):
    """Simple metatdata database to manage tags associated with devices"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('metadatadb.basic',addr)

    # -- Methods --
    def default_getDevices(_key,_value,_devices):
        """Get the list of known devices"""
        logger.info("default_getDevices(key=[%s],value=[%s],devices=[%s],)" % (_key,_value,_devices))

    def default_getKeysValues(_device,_keys,_map):
        """Get keys-values associated with a device"""
        logger.info("default_getKeysValues(device=[%s],keys=[%s],map=[%s],)" % (_device,_keys,_map))

    def default_getValue(_device,_key,_value):
        """Get the value of a key of a device"""
        logger.info("default_getValue(device=[%s],key=[%s],value=[%s],)" % (_device,_key,_value))

    def default_addKeysValues(_device,_map):
        """Add non-existing keys-values on a device"""
        logger.info("default_addKeysValues(device=[%s],map=[%s],)" % (_device,_map))

    def default_updateKeysValues(_device,_map):
        """Update keys-values on a device. Use the null value to delete a key. Use a null map to delete all keys-values of a device. Devices with no more key-value are withdrawn."""
        logger.info("default_updateKeysValues(device=[%s],map=[%s],)" % (_device,_map))

    dev.add_method('getDevices',default_getDevices)
    dev.add_method('getKeysValues',default_getKeysValues)
    dev.add_method('getValue',default_getValue)
    dev.add_method('addKeysValues',default_addKeysValues)
    dev.add_method('updateKeysValues',default_updateKeysValues)

    return dev

#=====================================================================
def motion(addr=None):
    """Simple motion detector device"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('motion.basic',addr)

    # -- Attributes --
    # Motion detected
    dev.new_attribute('presence')
    return dev

#=====================================================================
def powermeter(addr=None):
    """Simple powermeter"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('powermeter.basic',addr)

    # -- Attributes --
    # Current power
    dev.new_attribute('power')
    # Energy used or produced since cycle beginning
    dev.new_attribute('energy')
    # List of devices concerned with this powermeter
    dev.new_attribute('devices')
    return dev

#=====================================================================
def powerrelay(addr=None):
    """Simple power relay device"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('powerrelay.basic',addr)

    # -- Attributes --
    # State of the relay
    dev.new_attribute('power')

    # -- Methods --
    def default_on():
        """Switch on the relay"""
        logger.info("default_on()")

    def default_off():
        """Switch off the relay"""
        logger.info("default_off()")

    dev.add_method('on',default_on)
    dev.add_method('off',default_off)

    return dev

#=====================================================================
def powerrelay_toggle(addr=None):
    """Power relay with toggle function - Note that a toggle function may leads to undefined state due to its stateful nature; its usage should be avoided."""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('powerrelay.toggle',addr)

    # -- Attributes --
    # State of the relay
    dev.new_attribute('power')

    # -- Methods --
    def default_on():
        """Switch on the relay"""
        logger.info("default_on()")

    def default_off():
        """Switch off the relay"""
        logger.info("default_off()")

    def default_toggle():
        """Toggle relay state"""
        logger.info("default_toggle()")

    dev.add_method('on',default_on)
    dev.add_method('off',default_off)
    dev.add_method('toggle',default_toggle)

    return dev

#=====================================================================
def raingauge(addr=None):
    """Simple rain gauge"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('raingauge.basic',addr)

    # -- Attributes --
    # Real-time amount of rainfall
    dev.new_attribute('rain')
    # Accumulated precipitation
    dev.new_attribute('accumulated')
    return dev

#=====================================================================
def scale(addr=None):
    """Simple scale"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('scale.basic',addr)

    # -- Attributes --
    # Weight
    dev.new_attribute('weight')
    return dev

#=====================================================================
def schemarepository(addr=None):
    """Simple schema repository that can be queried about a schema name"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('schemarepository.basic',addr)

    # -- Methods --
    def default_getSchema(_devType,_schema):
        """Get a schema"""
        logger.info("default_getSchema(devType=[%s],schema=[%s],)" % (_devType,_schema))

    def default_getExpandedSchema(_devType,_schema):
        """Get a schema expanded by including schemas it extends"""
        logger.info("default_getExpandedSchema(devType=[%s],schema=[%s],)" % (_devType,_schema))

    dev.add_method('getSchema',default_getSchema)
    dev.add_method('getExpandedSchema',default_getExpandedSchema)

    return dev


#=====================================================================
def soundmeter(addr=None):
    """Simple sound meter"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('soundmeter.basic',addr)

    # -- Attributes --
    # Sound intensity
    dev.new_attribute('sound')
    return dev

#=====================================================================
def switch(addr=None):
    """Simple switch button device"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('switch.basic',addr)

    # -- Attributes --
    # Position of the switch
    dev.new_attribute('position')
    return dev

#=====================================================================
def tts(addr=None):
    """Text-To-Speech device"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('tts.basic',addr)

    # -- Methods --
    def default_say(_msg,_lang,_voice):
        """Say message"""
        logger.info("default_say(msg=[%s],lang=[%s],voice=[%s],)" % (_msg,_lang,_voice))

    dev.add_method('say',default_say)

    return dev

#=====================================================================
def window(addr=None):
    """Simple window device"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('window.basic',addr)

    # -- Attributes --
    # Position of the window
    dev.new_attribute('position')

    # -- Methods --
    def default_open():
        """Open the window"""
        logger.info("default_open()")

    def default_close():
        """Close the window"""
        logger.info("default_close()")

    dev.add_method('open',default_open)
    dev.add_method('close',default_close)

    return dev

#=====================================================================
def worktop(addr=None):
    """Simple worktop, e.g. kitchen worktop"""
    if (addr==None):addr = tools.get_random_uuid()
    dev = Device('worktop.basic',addr)

    # -- Attributes --
    # Ongoing action of the worktop
    dev.new_attribute('action')

    # -- Methods --
    def default_up():
        """Up the worktop"""
        logger.info("default_up()")

    def default_down():
        """Down the worktop"""
        logger.info("default_down()")

    def default_stop():
        """Stop ongoing action of the shutter"""
        logger.info("default_stop()")

    dev.add_method('up',default_up)
    dev.add_method('down',default_down)
    dev.add_method('stop',default_stop)

    return dev

*/
