package chromecontrol

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework ScriptingBridge -framework Foundation
#import <Foundation/Foundation.h>
#import "GoogleChrome.h"

void
enterPresentationMode(void) {
    GoogleChromeApplication *chrome = [SBApplication applicationWithBundleIdentifier:@"com.google.Chrome"];

    GoogleChromeWindow *window = chrome.windows.firstObject;

    // Create new window if no window exist
    if (!window) {
        window = [[[chrome classForScriptingClass:@"window"] alloc] init];
        [chrome.windows addObject:window];
    }

    [window enterPresentationMode];
}

*/
import "C"

func EnterPresentationMode() {
	C.enterPresentationMode()
}
