// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebViewConfiguration] class.
var WebViewConfigurationClass = _WebViewConfigurationClass{objc.GetClass("WKWebViewConfiguration")}

type _WebViewConfigurationClass struct {
	objc.Class
}

// An interface definition for the [WebViewConfiguration] class.
type IWebViewConfiguration interface {
	objc.IObject
	UrlSchemeHandlerForURLScheme(urlScheme string) URLSchemeHandlerWrapper
	SetURLSchemeHandlerForURLScheme(urlSchemeHandler PURLSchemeHandler, urlScheme string)
	SetURLSchemeHandlerObjectForURLScheme(urlSchemeHandlerObject objc.IObject, urlScheme string)
	LimitsNavigationsToAppBoundDomains() bool
	SetLimitsNavigationsToAppBoundDomains(value bool)
	UpgradeKnownHostsToHTTPS() bool
	SetUpgradeKnownHostsToHTTPS(value bool)
	UserContentController() UserContentController
	SetUserContentController(value IUserContentController)
	DefaultWebpagePreferences() WebpagePreferences
	SetDefaultWebpagePreferences(value IWebpagePreferences)
	WebsiteDataStore() WebsiteDataStore
	SetWebsiteDataStore(value IWebsiteDataStore)
	MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes
	SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes)
	UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy
	SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy)
	SuppressesIncrementalRendering() bool
	SetSuppressesIncrementalRendering(value bool)
	ProcessPool() ProcessPool
	SetProcessPool(value IProcessPool)
	AllowsAirPlayForMediaPlayback() bool
	SetAllowsAirPlayForMediaPlayback(value bool)
	ApplicationNameForUserAgent() string
	SetApplicationNameForUserAgent(value string)
	Preferences() Preferences
	SetPreferences(value IPreferences)
}

// A collection of properties that you use to initialize a web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration?language=objc
type WebViewConfiguration struct {
	objc.Object
}

func WebViewConfigurationFrom(ptr unsafe.Pointer) WebViewConfiguration {
	return WebViewConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebViewConfigurationClass) Alloc() WebViewConfiguration {
	rv := objc.Call[WebViewConfiguration](wc, objc.Sel("alloc"))
	return rv
}

func WebViewConfiguration_Alloc() WebViewConfiguration {
	return WebViewConfigurationClass.Alloc()
}

func (wc _WebViewConfigurationClass) New() WebViewConfiguration {
	rv := objc.Call[WebViewConfiguration](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebViewConfiguration() WebViewConfiguration {
	return WebViewConfigurationClass.New()
}

func (w_ WebViewConfiguration) Init() WebViewConfiguration {
	rv := objc.Call[WebViewConfiguration](w_, objc.Sel("init"))
	return rv
}

// Returns the currently registered handler object for the specified URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2875767-urlschemehandlerforurlscheme?language=objc
func (w_ WebViewConfiguration) UrlSchemeHandlerForURLScheme(urlScheme string) URLSchemeHandlerWrapper {
	rv := objc.Call[URLSchemeHandlerWrapper](w_, objc.Sel("urlSchemeHandlerForURLScheme:"), urlScheme)
	return rv
}

// Registers an object to load resources associated with the specified URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2875766-seturlschemehandler?language=objc
func (w_ WebViewConfiguration) SetURLSchemeHandlerForURLScheme(urlSchemeHandler PURLSchemeHandler, urlScheme string) {
	po0 := objc.WrapAsProtocol("WKURLSchemeHandler", urlSchemeHandler)
	objc.Call[objc.Void](w_, objc.Sel("setURLSchemeHandler:forURLScheme:"), po0, urlScheme)
}

// Registers an object to load resources associated with the specified URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2875766-seturlschemehandler?language=objc
func (w_ WebViewConfiguration) SetURLSchemeHandlerObjectForURLScheme(urlSchemeHandlerObject objc.IObject, urlScheme string) {
	objc.Call[objc.Void](w_, objc.Sel("setURLSchemeHandler:forURLScheme:"), objc.Ptr(urlSchemeHandlerObject), urlScheme)
}

// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3585117-limitsnavigationstoappbounddomai?language=objc
func (w_ WebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	rv := objc.Call[bool](w_, objc.Sel("limitsNavigationsToAppBoundDomains"))
	return rv
}

// A Boolean value that indicates whether the web view limits navigation to pages within the app’s domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3585117-limitsnavigationstoappbounddomai?language=objc
func (w_ WebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setLimitsNavigationsToAppBoundDomains:"), value)
}

// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3752243-upgradeknownhoststohttps?language=objc
func (w_ WebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	rv := objc.Call[bool](w_, objc.Sel("upgradeKnownHostsToHTTPS"))
	return rv
}

// A Boolean value that indicates whether the web view should automatically upgrade supported HTTP requests to HTTPS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3752243-upgradeknownhoststohttps?language=objc
func (w_ WebViewConfiguration) SetUpgradeKnownHostsToHTTPS(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setUpgradeKnownHostsToHTTPS:"), value)
}

// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395668-usercontentcontroller?language=objc
func (w_ WebViewConfiguration) UserContentController() UserContentController {
	rv := objc.Call[UserContentController](w_, objc.Sel("userContentController"))
	return rv
}

// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395668-usercontentcontroller?language=objc
func (w_ WebViewConfiguration) SetUserContentController(value IUserContentController) {
	objc.Call[objc.Void](w_, objc.Sel("setUserContentController:"), objc.Ptr(value))
}

// The default preferences to use when loading and rendering content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3194420-defaultwebpagepreferences?language=objc
func (w_ WebViewConfiguration) DefaultWebpagePreferences() WebpagePreferences {
	rv := objc.Call[WebpagePreferences](w_, objc.Sel("defaultWebpagePreferences"))
	return rv
}

// The default preferences to use when loading and rendering content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3194420-defaultwebpagepreferences?language=objc
func (w_ WebViewConfiguration) SetDefaultWebpagePreferences(value IWebpagePreferences) {
	objc.Call[objc.Void](w_, objc.Sel("setDefaultWebpagePreferences:"), objc.Ptr(value))
}

// The object you use to get and set the site’s cookies and to track the cached data objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395661-websitedatastore?language=objc
func (w_ WebViewConfiguration) WebsiteDataStore() WebsiteDataStore {
	rv := objc.Call[WebsiteDataStore](w_, objc.Sel("websiteDataStore"))
	return rv
}

// The object you use to get and set the site’s cookies and to track the cached data objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395661-websitedatastore?language=objc
func (w_ WebViewConfiguration) SetWebsiteDataStore(value IWebsiteDataStore) {
	objc.Call[objc.Void](w_, objc.Sel("setWebsiteDataStore:"), objc.Ptr(value))
}

// The media types that require a user gesture to begin playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1851524-mediatypesrequiringuseractionfor?language=objc
func (w_ WebViewConfiguration) MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes {
	rv := objc.Call[AudiovisualMediaTypes](w_, objc.Sel("mediaTypesRequiringUserActionForPlayback"))
	return rv
}

// The media types that require a user gesture to begin playing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1851524-mediatypesrequiringuseractionfor?language=objc
func (w_ WebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes) {
	objc.Call[objc.Void](w_, objc.Sel("setMediaTypesRequiringUserActionForPlayback:"), value)
}

// The directionality of user interface elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1690322-userinterfacedirectionpolicy?language=objc
func (w_ WebViewConfiguration) UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy {
	rv := objc.Call[UserInterfaceDirectionPolicy](w_, objc.Sel("userInterfaceDirectionPolicy"))
	return rv
}

// The directionality of user interface elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1690322-userinterfacedirectionpolicy?language=objc
func (w_ WebViewConfiguration) SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy) {
	objc.Call[objc.Void](w_, objc.Sel("setUserInterfaceDirectionPolicy:"), value)
}

// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395663-suppressesincrementalrendering?language=objc
func (w_ WebViewConfiguration) SuppressesIncrementalRendering() bool {
	rv := objc.Call[bool](w_, objc.Sel("suppressesIncrementalRendering"))
	return rv
}

// A Boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395663-suppressesincrementalrendering?language=objc
func (w_ WebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setSuppressesIncrementalRendering:"), value)
}

// The object that coordinates the processes the web view uses to render its web content and execute scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395659-processpool?language=objc
func (w_ WebViewConfiguration) ProcessPool() ProcessPool {
	rv := objc.Call[ProcessPool](w_, objc.Sel("processPool"))
	return rv
}

// The object that coordinates the processes the web view uses to render its web content and execute scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395659-processpool?language=objc
func (w_ WebViewConfiguration) SetProcessPool(value IProcessPool) {
	objc.Call[objc.Void](w_, objc.Sel("setProcessPool:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the web view allows media playback over AirPlay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395673-allowsairplayformediaplayback?language=objc
func (w_ WebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	rv := objc.Call[bool](w_, objc.Sel("allowsAirPlayForMediaPlayback"))
	return rv
}

// A Boolean value that indicates whether the web view allows media playback over AirPlay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395673-allowsairplayformediaplayback?language=objc
func (w_ WebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAllowsAirPlayForMediaPlayback:"), value)
}

// The app name that appears in the user agent string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395665-applicationnameforuseragent?language=objc
func (w_ WebViewConfiguration) ApplicationNameForUserAgent() string {
	rv := objc.Call[string](w_, objc.Sel("applicationNameForUserAgent"))
	return rv
}

// The app name that appears in the user agent string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395665-applicationnameforuseragent?language=objc
func (w_ WebViewConfiguration) SetApplicationNameForUserAgent(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setApplicationNameForUserAgent:"), value)
}

// The object that manages the preference-related settings for the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395666-preferences?language=objc
func (w_ WebViewConfiguration) Preferences() Preferences {
	rv := objc.Call[Preferences](w_, objc.Sel("preferences"))
	return rv
}

// The object that manages the preference-related settings for the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395666-preferences?language=objc
func (w_ WebViewConfiguration) SetPreferences(value IPreferences) {
	objc.Call[objc.Void](w_, objc.Sel("setPreferences:"), objc.Ptr(value))
}
