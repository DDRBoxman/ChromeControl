package chromecontrol

/*
#cgo CFLAGS: -x objective-c -fobjc-arc 
#cgo LDFLAGS: -framework ScriptingBridge -framework Foundation
#import <Foundation/Foundation.h>
#import "GoogleChrome.h"

GoogleChromeApplication*
chrome(void) {
	return [SBApplication applicationWithBundleIdentifier:@"com.google.Chrome"];
}

GoogleChromeWindow*
activeWindow(void) {
	GoogleChromeApplication* app = chrome();

    GoogleChromeWindow *window = app.windows.firstObject;

    // Create new window if no window exist
    if (!window) {
        window = [[[app classForScriptingClass:@"window"] alloc] init];
        [app.windows addObject:window];
    }

    return window;
}

void
quit(void) {
	GoogleChromeApplication* app = chrome();
	[app quit];
}

void
closeActiveWindow(void) {
	GoogleChromeWindow *window = activeWindow();
	[window close];
}

void
openURLInActiveTab(char* url) {
	NSString *urlString =  [NSString stringWithUTF8String:url];

	GoogleChromeWindow *window = activeWindow();
	window.activeTab.URL = urlString;

	free(url);
}

void
openURLInNewTab(char* url) {
	NSString *urlString =  [NSString stringWithUTF8String:url];

	GoogleChromeApplication* app = chrome();
    GoogleChromeTab *tab = [[[app classForScriptingClass:@"tab"] alloc] init];
    GoogleChromeWindow *window = activeWindow();
    [window.tabs addObject:tab];
    tab.URL = urlString;

	free(url);
}

void
enterPresentationMode(void) {
	GoogleChromeWindow *window = activeWindow();
    [window enterPresentationMode];
}

*/
import "C"

// Quit exits the currently open Chrome Application
func Quit() {
	C.quit()
}

// CloseActiveWindow closes the currently active Chrome window
func CloseActiveWindow() {
	C.closeActiveWindow()
}

// OpenURLInActiveTab navigates to the provided url in the currently active tab
func OpenURLInActiveTab(url string) {
	C.openURLInActiveTab(C.CString(url))
}

// OpenURLInNewTab opens a new tab and navigates to the provided url
func OpenURLInNewTab(url string) {
	C.openURLInNewTab(C.CString(url))
}

// EnterPresentationMode makes the currently active Chrome app enter presentation mode
func EnterPresentationMode() {
	C.enterPresentationMode()
}
