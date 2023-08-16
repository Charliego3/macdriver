// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Location] class.
var LocationClass = _LocationClass{objc.GetClass("CLLocation")}

type _LocationClass struct {
	objc.Class
}

// An interface definition for the [Location] class.
type ILocation interface {
	objc.IObject
	DistanceFromLocation(location ILocation) LocationDistance
	Floor() Floor
	HorizontalAccuracy() LocationAccuracy
	Coordinate() LocationCoordinate2D
	CourseAccuracy() LocationDirectionAccuracy
	Course() LocationDirection
	Speed() LocationSpeed
	EllipsoidalAltitude() LocationDistance
	SpeedAccuracy() LocationSpeedAccuracy
	Altitude() LocationDistance
	VerticalAccuracy() LocationAccuracy
	Timestamp() foundation.Date
	SourceInformation() LocationSourceInformation
}

// The latitude, longitude, and course information reported by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation?language=objc
type Location struct {
	objc.Object
}

func LocationFrom(ptr unsafe.Pointer) Location {
	return Location{
		Object: objc.ObjectFrom(ptr),
	}
}

func (l_ Location) InitWithLatitudeLongitude(latitude LocationDegrees, longitude LocationDegrees) Location {
	rv := objc.Call[Location](l_, objc.Sel("initWithLatitude:longitude:"), latitude, longitude)
	return rv
}

// Creates a location object with the specified latitude and longitude. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423660-initwithlatitude?language=objc
func Location_InitWithLatitudeLongitude(latitude LocationDegrees, longitude LocationDegrees) Location {
	return LocationClass.Alloc().InitWithLatitudeLongitude(latitude, longitude)
}

func (l_ Location) InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyTimestamp(coordinate LocationCoordinate2D, altitude LocationDistance, hAccuracy LocationAccuracy, vAccuracy LocationAccuracy, timestamp foundation.IDate) Location {
	rv := objc.Call[Location](l_, objc.Sel("initWithCoordinate:altitude:horizontalAccuracy:verticalAccuracy:timestamp:"), coordinate, altitude, hAccuracy, vAccuracy, objc.Ptr(timestamp))
	return rv
}

// Creates a location object with the specified coordinate and altitude information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423666-initwithcoordinate?language=objc
func Location_InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyTimestamp(coordinate LocationCoordinate2D, altitude LocationDistance, hAccuracy LocationAccuracy, vAccuracy LocationAccuracy, timestamp foundation.IDate) Location {
	return LocationClass.Alloc().InitWithCoordinateAltitudeHorizontalAccuracyVerticalAccuracyTimestamp(coordinate, altitude, hAccuracy, vAccuracy, timestamp)
}

func (lc _LocationClass) Alloc() Location {
	rv := objc.Call[Location](lc, objc.Sel("alloc"))
	return rv
}

func Location_Alloc() Location {
	return LocationClass.Alloc()
}

func (lc _LocationClass) New() Location {
	rv := objc.Call[Location](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLocation() Location {
	return LocationClass.New()
}

func (l_ Location) Init() Location {
	rv := objc.Call[Location](l_, objc.Sel("init"))
	return rv
}

// Returns the distance (measured in meters) from the current object’s location to the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423689-distancefromlocation?language=objc
func (l_ Location) DistanceFromLocation(location ILocation) LocationDistance {
	rv := objc.Call[LocationDistance](l_, objc.Sel("distanceFromLocation:"), objc.Ptr(location))
	return rv
}

// The logical floor of the building in which the user is located. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1616762-floor?language=objc
func (l_ Location) Floor() Floor {
	rv := objc.Call[Floor](l_, objc.Sel("floor"))
	return rv
}

// The radius of uncertainty for the location, measured in meters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423599-horizontalaccuracy?language=objc
func (l_ Location) HorizontalAccuracy() LocationAccuracy {
	rv := objc.Call[LocationAccuracy](l_, objc.Sel("horizontalAccuracy"))
	return rv
}

// The geographical coordinate information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423504-coordinate?language=objc
func (l_ Location) Coordinate() LocationCoordinate2D {
	rv := objc.Call[LocationCoordinate2D](l_, objc.Sel("coordinate"))
	return rv
}

// The accuracy of the course value, measured in degrees. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/3524338-courseaccuracy?language=objc
func (l_ Location) CourseAccuracy() LocationDirectionAccuracy {
	rv := objc.Call[LocationDirectionAccuracy](l_, objc.Sel("courseAccuracy"))
	return rv
}

// The direction in which the device is traveling, measured in degrees and relative to due north. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423832-course?language=objc
func (l_ Location) Course() LocationDirection {
	rv := objc.Call[LocationDirection](l_, objc.Sel("course"))
	return rv
}

// The instantaneous speed of the device, measured in meters per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423798-speed?language=objc
func (l_ Location) Speed() LocationSpeed {
	rv := objc.Call[LocationSpeed](l_, objc.Sel("speed"))
	return rv
}

// The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/3861801-ellipsoidalaltitude?language=objc
func (l_ Location) EllipsoidalAltitude() LocationDistance {
	rv := objc.Call[LocationDistance](l_, objc.Sel("ellipsoidalAltitude"))
	return rv
}

// The accuracy of the speed value, measured in meters per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/3524340-speedaccuracy?language=objc
func (l_ Location) SpeedAccuracy() LocationSpeedAccuracy {
	rv := objc.Call[LocationSpeedAccuracy](l_, objc.Sel("speedAccuracy"))
	return rv
}

// The altitude above mean sea level associated with a location, measured in meters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423820-altitude?language=objc
func (l_ Location) Altitude() LocationDistance {
	rv := objc.Call[LocationDistance](l_, objc.Sel("altitude"))
	return rv
}

// The validity of the altitude values, and their estimated uncertainty, measured in meters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423550-verticalaccuracy?language=objc
func (l_ Location) VerticalAccuracy() LocationAccuracy {
	rv := objc.Call[LocationAccuracy](l_, objc.Sel("verticalAccuracy"))
	return rv
}

// The time at which this location was determined. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/1423589-timestamp?language=objc
func (l_ Location) Timestamp() foundation.Date {
	rv := objc.Call[foundation.Date](l_, objc.Sel("timestamp"))
	return rv
}

// Information about the source that provides the location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/cllocation/3861803-sourceinformation?language=objc
func (l_ Location) SourceInformation() LocationSourceInformation {
	rv := objc.Call[LocationSourceInformation](l_, objc.Sel("sourceInformation"))
	return rv
}
