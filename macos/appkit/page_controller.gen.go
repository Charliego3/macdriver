// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PageController] class.
var PageControllerClass = _PageControllerClass{objc.GetClass("NSPageController")}

type _PageControllerClass struct {
	objc.Class
}

// An interface definition for the [PageController] class.
type IPageController interface {
	IViewController
	NavigateForward(sender objc.IObject) objc.Object
	NavigateForwardToObject(object objc.IObject)
	TakeSelectedIndexFrom(sender objc.IObject) objc.Object
	CompleteTransition()
	NavigateBack(sender objc.IObject) objc.Object
	SelectedViewController() ViewController
	ArrangedObjects() []objc.Object
	SetArrangedObjects(value []objc.IObject)
	Delegate() PageControllerDelegateWrapper
	SetDelegate(value PPageControllerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	TransitionStyle() PageControllerTransitionStyle
	SetTransitionStyle(value PageControllerTransitionStyle)
	SelectedIndex() int
	SetSelectedIndex(value int)
}

// An object that controls swipe navigation and animations between views or view content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller?language=objc
type PageController struct {
	ViewController
}

func PageControllerFrom(ptr unsafe.Pointer) PageController {
	return PageController{
		ViewController: ViewControllerFrom(ptr),
	}
}

func (pc _PageControllerClass) Alloc() PageController {
	rv := objc.Call[PageController](pc, objc.Sel("alloc"))
	return rv
}

func PageController_Alloc() PageController {
	return PageControllerClass.Alloc()
}

func (pc _PageControllerClass) New() PageController {
	rv := objc.Call[PageController](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPageController() PageController {
	return PageControllerClass.New()
}

func (p_ PageController) Init() PageController {
	rv := objc.Call[PageController](p_, objc.Sel("init"))
	return rv
}

func (p_ PageController) InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) PageController {
	rv := objc.Call[PageController](p_, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, objc.Ptr(nibBundleOrNil))
	return rv
}

// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434481-initwithnibname?language=objc
func PageController_InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) PageController {
	return PageControllerClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
}

// Navigates to the next object in the page controller’s arranged objects array, if appropriate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435004-navigateforward?language=objc
func (p_ PageController) NavigateForward(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("navigateForward:"), sender)
	return rv
}

// Navigates to the specific object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1434990-navigateforwardtoobject?language=objc
func (p_ PageController) NavigateForwardToObject(object objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("navigateForwardToObject:"), object)
}

// Navigates to the selected index, which is taken from the sender. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435011-takeselectedindexfrom?language=objc
func (p_ PageController) TakeSelectedIndexFrom(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("takeSelectedIndexFrom:"), sender)
	return rv
}

// Invoked when the page transition is completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1434994-completetransition?language=objc
func (p_ PageController) CompleteTransition() {
	objc.Call[objc.Void](p_, objc.Sel("completeTransition"))
}

// Navigates backwards in the page controller’s arranged objects array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435017-navigateback?language=objc
func (p_ PageController) NavigateBack(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("navigateBack:"), sender)
	return rv
}

// The view controller associated with the selected object.. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435013-selectedviewcontroller?language=objc
func (p_ PageController) SelectedViewController() ViewController {
	rv := objc.Call[ViewController](p_, objc.Sel("selectedViewController"))
	return rv
}

// An array containing the objects displayed in the page controller’s view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435001-arrangedobjects?language=objc
func (p_ PageController) ArrangedObjects() []objc.Object {
	rv := objc.Call[[]objc.Object](p_, objc.Sel("arrangedObjects"))
	return rv
}

// An array containing the objects displayed in the page controller’s view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435001-arrangedobjects?language=objc
func (p_ PageController) SetArrangedObjects(value []objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setArrangedObjects:"), value)
}

// The page controller’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435019-delegate?language=objc
func (p_ PageController) Delegate() PageControllerDelegateWrapper {
	rv := objc.Call[PageControllerDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// The page controller’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435019-delegate?language=objc
func (p_ PageController) SetDelegate(value PPageControllerDelegate) {
	po0 := objc.WrapAsProtocol("NSPageControllerDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), po0)
}

// The page controller’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1435019-delegate?language=objc
func (p_ PageController) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The transition style the page controller uses when changing pages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1434999-transitionstyle?language=objc
func (p_ PageController) TransitionStyle() PageControllerTransitionStyle {
	rv := objc.Call[PageControllerTransitionStyle](p_, objc.Sel("transitionStyle"))
	return rv
}

// The transition style the page controller uses when changing pages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1434999-transitionstyle?language=objc
func (p_ PageController) SetTransitionStyle(value PageControllerTransitionStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setTransitionStyle:"), value)
}

// The currently selected object in the arranged objects array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1434988-selectedindex?language=objc
func (p_ PageController) SelectedIndex() int {
	rv := objc.Call[int](p_, objc.Sel("selectedIndex"))
	return rv
}

// The currently selected object in the arranged objects array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontroller/1434988-selectedindex?language=objc
func (p_ PageController) SetSelectedIndex(value int) {
	objc.Call[objc.Void](p_, objc.Sel("setSelectedIndex:"), value)
}
