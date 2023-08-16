// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSControl subclasses to respond to editing actions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate?language=objc
type PControlTextEditingDelegate interface {
	// optional
	ControlTextDidChange(obj foundation.Notification)
	HasControlTextDidChange() bool

	// optional
	ControlTextDidBeginEditing(obj foundation.Notification)
	HasControlTextDidBeginEditing() bool

	// optional
	ControlTextDidEndEditing(obj foundation.Notification)
	HasControlTextDidEndEditing() bool

	// optional
	ControlTextShouldEndEditing(control Control, fieldEditor Text) bool
	HasControlTextShouldEndEditing() bool
}

// A delegate implementation builder for the [PControlTextEditingDelegate] protocol.
type ControlTextEditingDelegate struct {
	_ControlTextDidChange        func(obj foundation.Notification)
	_ControlTextDidBeginEditing  func(obj foundation.Notification)
	_ControlTextDidEndEditing    func(obj foundation.Notification)
	_ControlTextShouldEndEditing func(control Control, fieldEditor Text) bool
}

func (di *ControlTextEditingDelegate) HasControlTextDidChange() bool {
	return di._ControlTextDidChange != nil
}

// Tells the delegate that the control made changes to its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005177-controltextdidchange?language=objc
func (di *ControlTextEditingDelegate) SetControlTextDidChange(f func(obj foundation.Notification)) {
	di._ControlTextDidChange = f
}

// Tells the delegate that the control made changes to its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005177-controltextdidchange?language=objc
func (di *ControlTextEditingDelegate) ControlTextDidChange(obj foundation.Notification) {
	di._ControlTextDidChange(obj)
}
func (di *ControlTextEditingDelegate) HasControlTextDidBeginEditing() bool {
	return di._ControlTextDidBeginEditing != nil
}

// Tells the delegate that the control started editing its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005176-controltextdidbeginediting?language=objc
func (di *ControlTextEditingDelegate) SetControlTextDidBeginEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidBeginEditing = f
}

// Tells the delegate that the control started editing its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005176-controltextdidbeginediting?language=objc
func (di *ControlTextEditingDelegate) ControlTextDidBeginEditing(obj foundation.Notification) {
	di._ControlTextDidBeginEditing(obj)
}
func (di *ControlTextEditingDelegate) HasControlTextDidEndEditing() bool {
	return di._ControlTextDidEndEditing != nil
}

// Tells the delegate that the control finished editing its text content and committed the changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005178-controltextdidendediting?language=objc
func (di *ControlTextEditingDelegate) SetControlTextDidEndEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidEndEditing = f
}

// Tells the delegate that the control finished editing its text content and committed the changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005178-controltextdidendediting?language=objc
func (di *ControlTextEditingDelegate) ControlTextDidEndEditing(obj foundation.Notification) {
	di._ControlTextDidEndEditing(obj)
}
func (di *ControlTextEditingDelegate) HasControlTextShouldEndEditing() bool {
	return di._ControlTextShouldEndEditing != nil
}

// Invoked when the insertion point tries to leave a cell of the control that has been edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428984-control?language=objc
func (di *ControlTextEditingDelegate) SetControlTextShouldEndEditing(f func(control Control, fieldEditor Text) bool) {
	di._ControlTextShouldEndEditing = f
}

// Invoked when the insertion point tries to leave a cell of the control that has been edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428984-control?language=objc
func (di *ControlTextEditingDelegate) ControlTextShouldEndEditing(control Control, fieldEditor Text) bool {
	return di._ControlTextShouldEndEditing(control, fieldEditor)
}

// A concrete type wrapper for the [PControlTextEditingDelegate] protocol.
type ControlTextEditingDelegateWrapper struct {
	objc.Object
}

func (c_ ControlTextEditingDelegateWrapper) HasControlTextDidChange() bool {
	return c_.RespondsToSelector(objc.Sel("controlTextDidChange:"))
}

// Tells the delegate that the control made changes to its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005177-controltextdidchange?language=objc
func (c_ ControlTextEditingDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("controlTextDidChange:"), objc.Ptr(obj))
}

func (c_ ControlTextEditingDelegateWrapper) HasControlTextDidBeginEditing() bool {
	return c_.RespondsToSelector(objc.Sel("controlTextDidBeginEditing:"))
}

// Tells the delegate that the control started editing its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005176-controltextdidbeginediting?language=objc
func (c_ ControlTextEditingDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("controlTextDidBeginEditing:"), objc.Ptr(obj))
}

func (c_ ControlTextEditingDelegateWrapper) HasControlTextDidEndEditing() bool {
	return c_.RespondsToSelector(objc.Sel("controlTextDidEndEditing:"))
}

// Tells the delegate that the control finished editing its text content and committed the changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005178-controltextdidendediting?language=objc
func (c_ ControlTextEditingDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("controlTextDidEndEditing:"), objc.Ptr(obj))
}

func (c_ ControlTextEditingDelegateWrapper) HasControlTextShouldEndEditing() bool {
	return c_.RespondsToSelector(objc.Sel("control:textShouldEndEditing:"))
}

// Invoked when the insertion point tries to leave a cell of the control that has been edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428984-control?language=objc
func (c_ ControlTextEditingDelegateWrapper) ControlTextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.Call[bool](c_, objc.Sel("control:textShouldEndEditing:"), objc.Ptr(control), objc.Ptr(fieldEditor))
	return rv
}
