// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Subscription] class.
var SubscriptionClass = _SubscriptionClass{objc.GetClass("CKSubscription")}

type _SubscriptionClass struct {
	objc.Class
}

// An interface definition for the [Subscription] class.
type ISubscription interface {
	objc.IObject
	SubscriptionType() SubscriptionType
	NotificationInfo() NotificationInfo
	SetNotificationInfo(value INotificationInfo)
	SubscriptionID() SubscriptionID
}

// An abstract base class for subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription?language=objc
type Subscription struct {
	objc.Object
}

func SubscriptionFrom(ptr unsafe.Pointer) Subscription {
	return Subscription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SubscriptionClass) Alloc() Subscription {
	rv := objc.Call[Subscription](sc, objc.Sel("alloc"))
	return rv
}

func Subscription_Alloc() Subscription {
	return SubscriptionClass.Alloc()
}

func (sc _SubscriptionClass) New() Subscription {
	rv := objc.Call[Subscription](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSubscription() Subscription {
	return SubscriptionClass.New()
}

func (s_ Subscription) Init() Subscription {
	rv := objc.Call[Subscription](s_, objc.Sel("init"))
	return rv
}

// The behavior that a subscription provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/1515250-subscriptiontype?language=objc
func (s_ Subscription) SubscriptionType() SubscriptionType {
	rv := objc.Call[SubscriptionType](s_, objc.Sel("subscriptionType"))
	return rv
}

// The configuration for a subscription’s push notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/1514948-notificationinfo?language=objc
func (s_ Subscription) NotificationInfo() NotificationInfo {
	rv := objc.Call[NotificationInfo](s_, objc.Sel("notificationInfo"))
	return rv
}

// The configuration for a subscription’s push notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/1514948-notificationinfo?language=objc
func (s_ Subscription) SetNotificationInfo(value INotificationInfo) {
	objc.Call[objc.Void](s_, objc.Sel("setNotificationInfo:"), objc.Ptr(value))
}

// The subscription’s unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/1515199-subscriptionid?language=objc
func (s_ Subscription) SubscriptionID() SubscriptionID {
	rv := objc.Call[SubscriptionID](s_, objc.Sel("subscriptionID"))
	return rv
}
