// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Beacon] class.
var BeaconClass = _BeaconClass{objc.GetClass("CLBeacon")}

type _BeaconClass struct {
	objc.Class
}

// An interface definition for the [Beacon] class.
type IBeacon interface {
	objc.IObject
	Major() foundation.Number
	Minor() foundation.Number
	Rssi() int
	Accuracy() LocationAccuracy
	UUID() foundation.UUID
	Timestamp() foundation.Date
	Proximity() Proximity
}

// Information about an observed iBeacon device and its relative distance to the user’s device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon?language=objc
type Beacon struct {
	objc.Object
}

func BeaconFrom(ptr unsafe.Pointer) Beacon {
	return Beacon{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BeaconClass) Alloc() Beacon {
	rv := objc.Call[Beacon](bc, objc.Sel("alloc"))
	return rv
}

func Beacon_Alloc() Beacon {
	return BeaconClass.Alloc()
}

func (bc _BeaconClass) New() Beacon {
	rv := objc.Call[Beacon](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBeacon() Beacon {
	return BeaconClass.New()
}

func (b_ Beacon) Init() Beacon {
	rv := objc.Call[Beacon](b_, objc.Sel("init"))
	return rv
}

// The major value that the observed beacon transmitted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/1621418-major?language=objc
func (b_ Beacon) Major() foundation.Number {
	rv := objc.Call[foundation.Number](b_, objc.Sel("major"))
	return rv
}

// The minor value that the observed beacon transmitted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/1621558-minor?language=objc
func (b_ Beacon) Minor() foundation.Number {
	rv := objc.Call[foundation.Number](b_, objc.Sel("minor"))
	return rv
}

// The received signal strength of the beacon, measured in decibels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/1621557-rssi?language=objc
func (b_ Beacon) Rssi() int {
	rv := objc.Call[int](b_, objc.Sel("rssi"))
	return rv
}

// The accuracy of the proximity value, measured in meters from the beacon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/1621551-accuracy?language=objc
func (b_ Beacon) Accuracy() LocationAccuracy {
	rv := objc.Call[LocationAccuracy](b_, objc.Sel("accuracy"))
	return rv
}

// The UUID that the observed beacon transmitted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/3183017-uuid?language=objc
func (b_ Beacon) UUID() foundation.UUID {
	rv := objc.Call[foundation.UUID](b_, objc.Sel("UUID"))
	return rv
}

// A timestamp representing when the beacon was observed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/3183021-timestamp?language=objc
func (b_ Beacon) Timestamp() foundation.Date {
	rv := objc.Call[foundation.Date](b_, objc.Sel("timestamp"))
	return rv
}

// The relative distance to the beacon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/1621554-proximity?language=objc
func (b_ Beacon) Proximity() Proximity {
	rv := objc.Call[Proximity](b_, objc.Sel("proximity"))
	return rv
}
