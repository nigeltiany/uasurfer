package uasurfer

import (
	"strings"

	"github.com/nigeltiany/uasurfer/types/browser"
)

// Browser struct contains the lowercase name of the browser, along
// with its browser version number. Browser are grouped together without
// consideration for device. For example, Chrome (Chrome/43.0) and Chrome for iOS
// (CriOS/43.0) would both return as "chrome" (name) and 43.0 (version). Similarly
// Internet Explorer 11 and Edge 12 would return as "ie" and "11" or "12", respectively.
// type Browser struct {
// 		Name    BrowserName
// 		Version struct {
// 				Major int
// 			Minor int
// 			Patch int
// 		}
// }

// Retrieve browser name from UA strings
func (u *UserAgent) evalBrowserName(ua string) bool {
	// Blackberry goes first because it reads as MSIE & Safari
	if strings.Contains(ua, "blackberry") || strings.Contains(ua, "playbook") || strings.Contains(ua, "bb10") || strings.Contains(ua, "rim ") {
		u.Browser.Type = browser.BrowserType_Blackberry
		return u.maybeBot()
	}

	if strings.Contains(ua, "applewebkit") {
		switch {
		case strings.Contains(ua, "googlebot"):
			u.Browser.Type = browser.BrowserType_GoogleBot

		case strings.Contains(ua, "qq/") || strings.Contains(ua, "qqbrowser/"):
			u.Browser.Type = browser.BrowserType_QQ

		case strings.Contains(ua, "opr/") || strings.Contains(ua, "opios/"):
			u.Browser.Type = browser.BrowserType_Opera

		case strings.Contains(ua, "silk/"):
			u.Browser.Type = browser.BrowserType_Silk

		case strings.Contains(ua, "edg/") || strings.Contains(ua, "edgios/") || strings.Contains(ua, "edga/")|| strings.Contains(ua, "edge/") || strings.Contains(ua, "iemobile/") || strings.Contains(ua, "msie "):
			u.Browser.Type = browser.BrowserType_Internet_Explorer

		case strings.Contains(ua, "ucbrowser/") || strings.Contains(ua, "ucweb/"):
			u.Browser.Type = browser.BrowserType_UC_Browser

		case strings.Contains(ua, "nintendobrowser/"):
			u.Browser.Type = browser.BrowserType_Nintendo

		case strings.Contains(ua, "samsungbrowser/"):
			u.Browser.Type = browser.BrowserType_Samsung

		case strings.Contains(ua, "coc_coc_browser/"):
			u.Browser.Type = browser.BrowserType_Coc_Coc

		case strings.Contains(ua, "yabrowser/"):
			u.Browser.Type = browser.BrowserType_Yandex

		// Edge, Silk and other chrome-identifying browsers must evaluate before chrome, unless we want to add more overhead
		case strings.Contains(ua, "chrome/") || strings.Contains(ua, "crios/") || strings.Contains(ua, "chromium/") || strings.Contains(ua, "crmo/"):
			u.Browser.Type = browser.BrowserType_Chrome

		case strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") && !strings.Contains(ua, "like android"):
			// Android WebView on Android >= 4.4 is purposefully being identified as Chrome above -- https://developer.chrome.com/multidevice/webview/overview
			u.Browser.Type = browser.BrowserType_Android

		case strings.Contains(ua, "fxios"):
			u.Browser.Type = browser.BrowserType_Firefox

		case strings.Contains(ua, " spotify/"):
			u.Browser.Type = browser.BrowserType_Spotify

		// AppleBot uses webkit signature as well
		case strings.Contains(ua, "applebot"):
			u.Browser.Type = browser.BrowserType_AppleBot

		// presume it's safari unless an esoteric browser is being specified (webOSBrowser, SamsungBrowser, etc.)
		case strings.Contains(ua, "like gecko") && strings.Contains(ua, "mozilla/") && strings.Contains(ua, "safari/") && !strings.Contains(ua, "linux") && !strings.Contains(ua, "android") && !strings.Contains(ua, "browser/") && !strings.Contains(ua, "os/") && !strings.Contains(ua, "yabrowser/"):
			u.Browser.Type = browser.BrowserType_Safari

		// if we got this far and the device is iPhone or iPad, assume safari. Some agents don't actually contain the word "safari"
		case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
			u.Browser.Type = browser.BrowserType_Safari

		// Google's search app on iPhone, leverages native Safari rather than Chrome
		case strings.Contains(ua, " gsa/"):
			u.Browser.Type = browser.BrowserType_Safari

		default:
			goto notwebkit

		}
		return u.maybeBot()
	}

notwebkit:
	switch {
	case strings.Contains(ua, "qq/") || strings.Contains(ua, "qqbrowser/"):
		u.Browser.Type = browser.BrowserType_QQ

	case strings.Contains(ua, "msie") || strings.Contains(ua, "trident"):
		u.Browser.Type = browser.BrowserType_Internet_Explorer

	case strings.Contains(ua, "gecko") && (strings.Contains(ua, "firefox") || strings.Contains(ua, "iceweasel") || strings.Contains(ua, "seamonkey") || strings.Contains(ua, "icecat")):
		u.Browser.Type = browser.BrowserType_Firefox

	case strings.Contains(ua, "presto") || strings.Contains(ua, "opera"):
		u.Browser.Type = browser.BrowserType_Opera

	case strings.Contains(ua, "ucbrowser"):
		u.Browser.Type = browser.BrowserType_UC_Browser

	case strings.Contains(ua, "applebot"):
		u.Browser.Type = browser.BrowserType_AppleBot

	case strings.Contains(ua, "baiduspider"):
		u.Browser.Type = browser.BrowserType_BaiduBot

	case strings.Contains(ua, "adidxbot") || strings.Contains(ua, "bingbot") || strings.Contains(ua, "bingpreview"):
		u.Browser.Type = browser.BrowserType_BingBot

	case strings.Contains(ua, "duckduckbot"):
		u.Browser.Type = browser.BrowserType_DuckDuckGoBot

	case strings.Contains(ua, "facebot") || strings.Contains(ua, "facebookexternalhit"):
		u.Browser.Type = browser.BrowserType_FacebookBot

	case strings.Contains(ua, "googlebot"):
		u.Browser.Type = browser.BrowserType_GoogleBot

	case strings.Contains(ua, "linkedinbot"):
		u.Browser.Type = browser.BrowserType_LinkedInBot

	case strings.Contains(ua, "msnbot"):
		u.Browser.Type = browser.BrowserType_MsnBot

	case strings.Contains(ua, "pingdom.com_bot"):
		u.Browser.Type = browser.BrowserType_PingdomBot

	case strings.Contains(ua, "twitterbot"):
		u.Browser.Type = browser.BrowserType_TwitterBot

	case strings.Contains(ua, "yandex") || strings.Contains(ua, "yadirectfetcher"):
		u.Browser.Type = browser.BrowserType_YandexBot

	case strings.Contains(ua, "yahoo"):
		u.Browser.Type = browser.BrowserType_YahooBot

	case strings.Contains(ua, "coccocbot"):
		u.Browser.Type = browser.BrowserType_CocCocBot

	case strings.Contains(ua, "phantomjs"):
		u.Browser.Type = browser.BrowserType_Bot

	default:
		u.Browser.Type = browser.BrowserType_unknown

	}

	return u.maybeBot()
}

// Retrieve browser version
// Methods used in order:
// 1st: look for generic version/#
// 2nd: look for browser-specific instructions (e.g. chrome/34)
// 3rd: infer from OS (iOS only)
func (u *UserAgent) evalBrowserVersion(ua string) {
	// if there is a 'version/#' attribute with numeric version, use it -- except for Chrome since Android vendors sometimes hijack version/#
	if u.Browser.Type != browser.BrowserType_Chrome && u.Browser.Version.findVersionNumber(ua, "version/") {
		return
	}

	switch u.Browser.Type {
	case browser.BrowserType_Chrome:
		// match both chrome and crios
		_ = u.Browser.Version.findVersionNumber(ua, "chrome/") || u.Browser.Version.findVersionNumber(ua, "crios/") || u.Browser.Version.findVersionNumber(ua, "crmo/")
	case browser.BrowserType_Yandex:
		_ = u.Browser.Version.findVersionNumber(ua, "yabrowser/")
	case browser.BrowserType_QQ:
		if u.Browser.Version.findVersionNumber(ua, "qq/") {
			return
		}
		_ = u.Browser.Version.findVersionNumber(ua, "qqbrowser/")
	case browser.BrowserType_Internet_Explorer:
		if u.Browser.Version.findVersionNumber(ua, "msie ") || u.Browser.Version.findVersionNumber(ua, "edge/") || u.Browser.Version.findVersionNumber(ua, "edgios/") || u.Browser.Version.findVersionNumber(ua, "edga/") || u.Browser.Version.findVersionNumber(ua, "edg/") {
			return
		}

		// get MSIE version from trident version https://en.wikipedia.org/wiki/Trident_(layout_engine)
		if u.Browser.Version.findVersionNumber(ua, "trident/") {
			// convert trident versions 3-7 to MSIE version
			if (u.Browser.Version.Major >= 3) && (u.Browser.Version.Major <= 7) {
				u.Browser.Version.Major += 4
			}
		}

	case browser.BrowserType_Firefox:
		_ = u.Browser.Version.findVersionNumber(ua, "firefox/") || u.Browser.Version.findVersionNumber(ua, "fxios/")

	case browser.BrowserType_Safari: // executes typically if we're on iOS and not using a familiar browser
		u.Browser.Version = u.OS.Version
		// early Safari used a version number +1 to OS version
		if (u.Browser.Version.Major <= 3) && (u.Browser.Version.Major >= 1) {
			u.Browser.Version.Major++
		}

	case browser.BrowserType_UC_Browser:
		_ = u.Browser.Version.findVersionNumber(ua, "ucbrowser/")

	case browser.BrowserType_Opera:
		_ = u.Browser.Version.findVersionNumber(ua, "opr/") || u.Browser.Version.findVersionNumber(ua, "opios/") || u.Browser.Version.findVersionNumber(ua, "opera/")

	case browser.BrowserType_Silk:
		_ = u.Browser.Version.findVersionNumber(ua, "silk/")

	case browser.BrowserType_Spotify:
		_ = u.Browser.Version.findVersionNumber(ua, "spotify/")

	case browser.BrowserType_Coc_Coc:
		_ = u.Browser.Version.findVersionNumber(ua, "coc_coc_browser/")
	}
}
