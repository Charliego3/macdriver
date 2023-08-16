// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ViewController] class.
var ViewControllerClass = _ViewControllerClass{objc.GetClass("NSViewController")}

type _ViewControllerClass struct {
	objc.Class
}

// An interface definition for the [ViewController] class.
type IViewController interface {
	IResponder
	ViewDidLayout()
	ViewWillLayout()
	CommitEditing() bool
	DismissController(sender objc.IObject) objc.Object
	ViewWillDisappear()
	DismissViewController(viewController IViewController)
	ViewDidAppear()
	ViewWillAppear()
	UpdateViewConstraints()
	AddChildViewController(childViewController IViewController)
	PresentViewControllerAsSheet(viewController IViewController)
	TransitionFromViewControllerToViewControllerOptionsCompletionHandler(fromViewController IViewController, toViewController IViewController, options ViewControllerTransitionOptions, completion func())
	RemoveChildViewControllerAtIndex(index int)
	ViewDidLoad()
	LoadView()
	InsertChildViewControllerAtIndex(childViewController IViewController, index int)
	ViewWillTransitionToSize(newSize foundation.Size)
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer)
	PreferredContentSizeDidChangeForViewController(viewController IViewController)
	RemoveFromParentViewController()
	DiscardEditing()
	ViewDidDisappear()
	PresentViewControllerAsModalWindow(viewController IViewController)
	PresentViewControllerAnimator(viewController IViewController, animator PViewControllerPresentationAnimator)
	PresentViewControllerAnimatorObject(viewController IViewController, animatorObject objc.IObject)
	SourceItemView() View
	SetSourceItemView(value IView)
	PresentedViewControllers() []ViewController
	IsViewLoaded() bool
	NibName() NibName
	ParentViewController() ViewController
	View() View
	SetView(value IView)
	NibBundle() foundation.Bundle
	PreferredMaximumSize() foundation.Size
	PreferredMinimumSize() foundation.Size
	PreferredScreenOrigin() foundation.Point
	SetPreferredScreenOrigin(value foundation.Point)
	Storyboard() Storyboard
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.IObject)
	Title() string
	SetTitle(value string)
	PresentingViewController() ViewController
	ExtensionContext() foundation.ExtensionContext
	PreferredContentSize() foundation.Size
	SetPreferredContentSize(value foundation.Size)
	ChildViewControllers() []ViewController
	SetChildViewControllers(value []IViewController)
}

// A controller that manages a view, typically loaded from a nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller?language=objc
type ViewController struct {
	Responder
}

func ViewControllerFrom(ptr unsafe.Pointer) ViewController {
	return ViewController{
		Responder: ResponderFrom(ptr),
	}
}

func (v_ ViewController) InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) ViewController {
	rv := objc.Call[ViewController](v_, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, objc.Ptr(nibBundleOrNil))
	return rv
}

// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434481-initwithnibname?language=objc
func ViewController_InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) ViewController {
	return ViewControllerClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
}

func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.Call[ViewController](vc, objc.Sel("alloc"))
	return rv
}

func ViewController_Alloc() ViewController {
	return ViewControllerClass.Alloc()
}

func (vc _ViewControllerClass) New() ViewController {
	rv := objc.Call[ViewController](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewViewController() ViewController {
	return ViewControllerClass.New()
}

func (v_ ViewController) Init() ViewController {
	rv := objc.Call[ViewController](v_, objc.Sel("init"))
	return rv
}

// Called immediately after the layout method of the view controller's view is called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434451-viewdidlayout?language=objc
func (v_ ViewController) ViewDidLayout() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidLayout"))
}

// Called just before the layout method of the view controller's view is called. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434495-viewwilllayout?language=objc
func (v_ ViewController) ViewWillLayout() {
	objc.Call[objc.Void](v_, objc.Sel("viewWillLayout"))
}

// Returns whether the receiver was able to commit any pending edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434485-commitediting?language=objc
func (v_ ViewController) CommitEditing() bool {
	rv := objc.Call[bool](v_, objc.Sel("commitEditing"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434447-dismisscontroller?language=objc
func (v_ ViewController) DismissController(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](v_, objc.Sel("dismissController:"), sender)
	return rv
}

// Called when the view controller’s view is about to be removed from the view hierarchy in the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434483-viewwilldisappear?language=objc
func (v_ ViewController) ViewWillDisappear() {
	objc.Call[objc.Void](v_, objc.Sel("viewWillDisappear"))
}

// Dismisses a presented view controller, using the same animator that presented it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434413-dismissviewcontroller?language=objc
func (v_ ViewController) DismissViewController(viewController IViewController) {
	objc.Call[objc.Void](v_, objc.Sel("dismissViewController:"), objc.Ptr(viewController))
}

// Called when the view controller’s view is fully transitioned onto the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434455-viewdidappear?language=objc
func (v_ ViewController) ViewDidAppear() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidAppear"))
}

// Called after the view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434415-viewwillappear?language=objc
func (v_ ViewController) ViewWillAppear() {
	objc.Call[objc.Void](v_, objc.Sel("viewWillAppear"))
}

// Called during Auto Layout constraint updating to enable the view controller to mediate the process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434400-updateviewconstraints?language=objc
func (v_ ViewController) UpdateViewConstraints() {
	objc.Call[objc.Void](v_, objc.Sel("updateViewConstraints"))
}

// A convenience method for adding a child view controller at the end of the childViewControllers array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434501-addchildviewcontroller?language=objc
func (v_ ViewController) AddChildViewController(childViewController IViewController) {
	objc.Call[objc.Void](v_, objc.Sel("addChildViewController:"), objc.Ptr(childViewController))
}

// Presents another view controller as a sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434489-presentviewcontrollerassheet?language=objc
func (v_ ViewController) PresentViewControllerAsSheet(viewController IViewController) {
	objc.Call[objc.Void](v_, objc.Sel("presentViewControllerAsSheet:"), objc.Ptr(viewController))
}

// Performs a transition between two sibling child view controllers of the view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434407-transitionfromviewcontroller?language=objc
func (v_ ViewController) TransitionFromViewControllerToViewControllerOptionsCompletionHandler(fromViewController IViewController, toViewController IViewController, options ViewControllerTransitionOptions, completion func()) {
	objc.Call[objc.Void](v_, objc.Sel("transitionFromViewController:toViewController:options:completionHandler:"), objc.Ptr(fromViewController), objc.Ptr(toViewController), options, completion)
}

// Removes a specified child controller from the view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434404-removechildviewcontrolleratindex?language=objc
func (v_ ViewController) RemoveChildViewControllerAtIndex(index int) {
	objc.Call[objc.Void](v_, objc.Sel("removeChildViewControllerAtIndex:"), index)
}

// Called after the view controller’s view has been loaded into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434476-viewdidload?language=objc
func (v_ ViewController) ViewDidLoad() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidLoad"))
}

// Instantiates a view from a nib file and sets the value of the view property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434405-loadview?language=objc
func (v_ ViewController) LoadView() {
	objc.Call[objc.Void](v_, objc.Sel("loadView"))
}

// Inserts a specified child view controller into the childViewControllers array at a specified position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434437-insertchildviewcontroller?language=objc
func (v_ ViewController) InsertChildViewControllerAtIndex(childViewController IViewController, index int) {
	objc.Call[objc.Void](v_, objc.Sel("insertChildViewController:atIndex:"), objc.Ptr(childViewController), index)
}

// For a view controller that is part of an app extension, called when its view is about to be resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434443-viewwilltransitiontosize?language=objc
func (v_ ViewController) ViewWillTransitionToSize(newSize foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("viewWillTransitionToSize:"), newSize)
}

// Attempt to commit any currently edited results of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434464-commiteditingwithdelegate?language=objc
func (v_ ViewController) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](v_, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}

// Called when there is a change in value of the preferredContentSize property of a child view controller or a presented view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434434-preferredcontentsizedidchangefor?language=objc
func (v_ ViewController) PreferredContentSizeDidChangeForViewController(viewController IViewController) {
	objc.Call[objc.Void](v_, objc.Sel("preferredContentSizeDidChangeForViewController:"), objc.Ptr(viewController))
}

// Removes the called view controller from its parent view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434466-removefromparentviewcontroller?language=objc
func (v_ ViewController) RemoveFromParentViewController() {
	objc.Call[objc.Void](v_, objc.Sel("removeFromParentViewController"))
}

// Causes the receiver to discard any changes, restoring the previous values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434487-discardediting?language=objc
func (v_ ViewController) DiscardEditing() {
	objc.Call[objc.Void](v_, objc.Sel("discardEditing"))
}

// Called after the view controller’s view is removed from the view hierarchy in a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434416-viewdiddisappear?language=objc
func (v_ ViewController) ViewDidDisappear() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidDisappear"))
}

// Presents another view controller as a modal window, also known as an alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434462-presentviewcontrollerasmodalwind?language=objc
func (v_ ViewController) PresentViewControllerAsModalWindow(viewController IViewController) {
	objc.Call[objc.Void](v_, objc.Sel("presentViewControllerAsModalWindow:"), objc.Ptr(viewController))
}

// Presents another view controller using a specified, custom animator for presentation and dismissal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434431-presentviewcontroller?language=objc
func (v_ ViewController) PresentViewControllerAnimator(viewController IViewController, animator PViewControllerPresentationAnimator) {
	po1 := objc.WrapAsProtocol("NSViewControllerPresentationAnimator", animator)
	objc.Call[objc.Void](v_, objc.Sel("presentViewController:animator:"), objc.Ptr(viewController), po1)
}

// Presents another view controller using a specified, custom animator for presentation and dismissal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434431-presentviewcontroller?language=objc
func (v_ ViewController) PresentViewControllerAnimatorObject(viewController IViewController, animatorObject objc.IObject) {
	objc.Call[objc.Void](v_, objc.Sel("presentViewController:animator:"), objc.Ptr(viewController), objc.Ptr(animatorObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434479-sourceitemview?language=objc
func (v_ ViewController) SourceItemView() View {
	rv := objc.Call[View](v_, objc.Sel("sourceItemView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434479-sourceitemview?language=objc
func (v_ ViewController) SetSourceItemView(value IView) {
	objc.Call[objc.Void](v_, objc.Sel("setSourceItemView:"), objc.Ptr(value))
}

// The view controllers, if any, that are currently presented by the view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434497-presentedviewcontrollers?language=objc
func (v_ ViewController) PresentedViewControllers() []ViewController {
	rv := objc.Call[[]ViewController](v_, objc.Sel("presentedViewControllers"))
	return rv
}

// A Boolean value indicating whether the view controller’s view is loaded into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434435-viewloaded?language=objc
func (v_ ViewController) IsViewLoaded() bool {
	rv := objc.Call[bool](v_, objc.Sel("isViewLoaded"))
	return rv
}

// The name of the nib file to be loaded to instantiate the receiver’s primary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434472-nibname?language=objc
func (v_ ViewController) NibName() NibName {
	rv := objc.Call[NibName](v_, objc.Sel("nibName"))
	return rv
}

// The immediate ancestor view controller of the view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434491-parentviewcontroller?language=objc
func (v_ ViewController) ParentViewController() ViewController {
	rv := objc.Call[ViewController](v_, objc.Sel("parentViewController"))
	return rv
}

// The view controller’s primary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434401-view?language=objc
func (v_ ViewController) View() View {
	rv := objc.Call[View](v_, objc.Sel("view"))
	return rv
}

// The view controller’s primary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434401-view?language=objc
func (v_ ViewController) SetView(value IView) {
	objc.Call[objc.Void](v_, objc.Sel("setView:"), objc.Ptr(value))
}

// The nib bundle to be loaded to instantiate the receiver’s primary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434433-nibbundle?language=objc
func (v_ ViewController) NibBundle() foundation.Bundle {
	rv := objc.Call[foundation.Bundle](v_, objc.Sel("nibBundle"))
	return rv
}

// For a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434403-preferredmaximumsize?language=objc
func (v_ ViewController) PreferredMaximumSize() foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("preferredMaximumSize"))
	return rv
}

// For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434418-preferredminimumsize?language=objc
func (v_ ViewController) PreferredMinimumSize() foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("preferredMinimumSize"))
	return rv
}

// For a view controller that is part of an app extension, the preferred screen origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434468-preferredscreenorigin?language=objc
func (v_ ViewController) PreferredScreenOrigin() foundation.Point {
	rv := objc.Call[foundation.Point](v_, objc.Sel("preferredScreenOrigin"))
	return rv
}

// For a view controller that is part of an app extension, the preferred screen origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434468-preferredscreenorigin?language=objc
func (v_ ViewController) SetPreferredScreenOrigin(value foundation.Point) {
	objc.Call[objc.Void](v_, objc.Sel("setPreferredScreenOrigin:"), value)
}

// The storyboard from which the view controller was loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434499-storyboard?language=objc
func (v_ ViewController) Storyboard() Storyboard {
	rv := objc.Call[Storyboard](v_, objc.Sel("storyboard"))
	return rv
}

// The object whose value is presented in the receiver’s primary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434453-representedobject?language=objc
func (v_ ViewController) RepresentedObject() objc.Object {
	rv := objc.Call[objc.Object](v_, objc.Sel("representedObject"))
	return rv
}

// The object whose value is presented in the receiver’s primary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434453-representedobject?language=objc
func (v_ ViewController) SetRepresentedObject(value objc.IObject) {
	objc.Call[objc.Void](v_, objc.Sel("setRepresentedObject:"), value)
}

// The localized title of the receiver’s primary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434426-title?language=objc
func (v_ ViewController) Title() string {
	rv := objc.Call[string](v_, objc.Sel("title"))
	return rv
}

// The localized title of the receiver’s primary view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434426-title?language=objc
func (v_ ViewController) SetTitle(value string) {
	objc.Call[objc.Void](v_, objc.Sel("setTitle:"), value)
}

// The view controller that presented the view controller or that presented its farthest ancestor view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434439-presentingviewcontroller?language=objc
func (v_ ViewController) PresentingViewController() ViewController {
	rv := objc.Call[ViewController](v_, objc.Sel("presentingViewController"))
	return rv
}

// For a view controller that is part of an app extension, the app extension context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434457-extensioncontext?language=objc
func (v_ ViewController) ExtensionContext() foundation.ExtensionContext {
	rv := objc.Call[foundation.ExtensionContext](v_, objc.Sel("extensionContext"))
	return rv
}

// The desired size of the view controller’s view, in screen units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434409-preferredcontentsize?language=objc
func (v_ ViewController) PreferredContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("preferredContentSize"))
	return rv
}

// The desired size of the view controller’s view, in screen units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434409-preferredcontentsize?language=objc
func (v_ ViewController) SetPreferredContentSize(value foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("setPreferredContentSize:"), value)
}

// An array of view controllers that are hierarchical children of the view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434432-childviewcontrollers?language=objc
func (v_ ViewController) ChildViewControllers() []ViewController {
	rv := objc.Call[[]ViewController](v_, objc.Sel("childViewControllers"))
	return rv
}

// An array of view controllers that are hierarchical children of the view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434432-childviewcontrollers?language=objc
func (v_ ViewController) SetChildViewControllers(value []IViewController) {
	objc.Call[objc.Void](v_, objc.Sel("setChildViewControllers:"), value)
}
