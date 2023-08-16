// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutablePostalAddress] class.
var MutablePostalAddressClass = _MutablePostalAddressClass{objc.GetClass("CNMutablePostalAddress")}

type _MutablePostalAddressClass struct {
	objc.Class
}

// An interface definition for the [MutablePostalAddress] class.
type IMutablePostalAddress interface {
	IPostalAddress
	SetState(value string)
	SetSubLocality(value string)
	SetStreet(value string)
	SetCity(value string)
	SetPostalCode(value string)
	SetISOCountryCode(value string)
	SetSubAdministrativeArea(value string)
	SetCountry(value string)
}

// A mutable representation of the postal address for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress?language=objc
type MutablePostalAddress struct {
	PostalAddress
}

func MutablePostalAddressFrom(ptr unsafe.Pointer) MutablePostalAddress {
	return MutablePostalAddress{
		PostalAddress: PostalAddressFrom(ptr),
	}
}

func (mc _MutablePostalAddressClass) Alloc() MutablePostalAddress {
	rv := objc.Call[MutablePostalAddress](mc, objc.Sel("alloc"))
	return rv
}

func MutablePostalAddress_Alloc() MutablePostalAddress {
	return MutablePostalAddressClass.Alloc()
}

func (mc _MutablePostalAddressClass) New() MutablePostalAddress {
	rv := objc.Call[MutablePostalAddress](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutablePostalAddress() MutablePostalAddress {
	return MutablePostalAddressClass.New()
}

func (m_ MutablePostalAddress) Init() MutablePostalAddress {
	rv := objc.Call[MutablePostalAddress](m_, objc.Sel("init"))
	return rv
}

// The state name of the address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/1402796-state?language=objc
func (m_ MutablePostalAddress) SetState(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setState:"), value)
}

// Additional information associated with the location, typically defined at the city or town level, in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/2799058-sublocality?language=objc
func (m_ MutablePostalAddress) SetSubLocality(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setSubLocality:"), value)
}

// The street name of the address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/1402981-street?language=objc
func (m_ MutablePostalAddress) SetStreet(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setStreet:"), value)
}

// The city name of the address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/1403395-city?language=objc
func (m_ MutablePostalAddress) SetCity(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setCity:"), value)
}

// The postal code of the address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/1403303-postalcode?language=objc
func (m_ MutablePostalAddress) SetPostalCode(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setPostalCode:"), value)
}

// The ISO country code, using the ISO 3166-1 alpha-2 standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/1403295-isocountrycode?language=objc
func (m_ MutablePostalAddress) SetISOCountryCode(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setISOCountryCode:"), value)
}

// The subadministrative area (such as a county or other region) in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/2799059-subadministrativearea?language=objc
func (m_ MutablePostalAddress) SetSubAdministrativeArea(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setSubAdministrativeArea:"), value)
}

// The country or region name of the address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablepostaladdress/1402917-country?language=objc
func (m_ MutablePostalAddress) SetCountry(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setCountry:"), value)
}
