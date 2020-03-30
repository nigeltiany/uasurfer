package uasurfer

import (
	"testing"

	"github.com/nigeltiany/uasurfer/types/browser"
	"github.com/nigeltiany/uasurfer/types/device"
	"github.com/nigeltiany/uasurfer/types/os"
	"github.com/nigeltiany/uasurfer/types/platform"
	"github.com/nigeltiany/uasurfer/types/versioning"
)

var testUAVars = []struct {
	UA string
	UserAgent
}{
	// Empty
	{"",
		UserAgent{}},

	// Single char
	{"a",
		UserAgent{}},

	// Some random string
	{"some random string",
		UserAgent{}},

	// Potentially malformed ua
	{")(",
		UserAgent{}},

	// iPhone
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 8, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 8, Minor: 0, Patch: 2}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (iPhone10,3; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 8, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 8, Minor: 0, Patch: 2}}}, device.DeviceType_phone}},

	// iPad
	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 4, Minor: 0, Patch: 4}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 3, Minor: 2, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (iPad; CPU OS 9_0 like Mac OS X) AppleWebKit/601.1.17 (KHTML, like Gecko) Version/8.0 Mobile/13A175 Safari/600.1.4",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 8, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 9, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.32 (KHTML, like Gecko) Version/10.0 Mobile/14A5261v Safari/602.1",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	// Chrome
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 43, Minor: 0, Patch: 2357}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 10, Patch: 4}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/534.48.3",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 19, Minor: 0, Patch: 1084}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 5, Minor: 1, Patch: 1}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; Android 6.0; Nexus 5X Build/MDB08L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 46, Minor: 0, Patch: 2490}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	// Chromium (Chrome)
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.19 (KHTML, like Gecko) Ubuntu/11.10 Chromium/18.0.1025.142 Chrome/18.0.1025.142 Safari/535.19",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 18, Minor: 0, Patch: 1025}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 45, Minor: 0, Patch: 2454}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 11, Patch: 0}}}, device.DeviceType_computer}},

	//TODO: refactor "getVersion()" to handle this device/chrome version douchebaggery
	// {"Mozilla/5.0 (Linux; Android 4.4.2; en-gb; SAMSUNG SM-G800F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/1.6 Chrome/28.0.1500.94 Mobile Safari/537.36",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Chrome, Version{28,0,1500}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{4,4,2}}, device.DeviceType_phone}},

	// Safari
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 8, Minor: 0, Patch: 7}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 10, Patch: 4}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 3, Minor: 2, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 5, Patch: 5}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.1.32 (KHTML, like Gecko) Version/10.0 Safari/602.1.32", // macOS Sierra dev beta
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 12, Patch: 0}}}, device.DeviceType_computer}},

	// Firefox
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 1, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 8, Minor: 3, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Android 4.4; Tablet; rv:41.0) Gecko/41.0 Firefox/41.0",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 41, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 4, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (Android; Mobile; rv:40.0) Gecko/40.0 Firefox/40.0",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 40, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:38.0) Gecko/20100101 Firefox/38.0",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 38, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	// Silk
	{"Mozilla/5.0 (Linux; U; Android 4.4.3; de-de; KFTHWI Build/KTU84M) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.47 like Chrome/37.0.2026.117 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Silk, Version{versioning.Version{Major: 3, Minor: 47, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{versioning.Version{Major: 4, Minor: 4, Patch: 3}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (Linux; U; en-us; KFJWI Build/IMM76D) AppleWebKit/535.19 (KHTML like Gecko) Silk/2.4 Safari/535.19 Silk-Acceleratedtrue",
		UserAgent{
			Browser{browser.BrowserType_Silk, Version{versioning.Version{Major: 2, Minor: 4, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	// Opera
	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 OPR/18.0.1284.68",
		UserAgent{
			Browser{browser.BrowserType_Opera, Version{versioning.Version{Major: 18, Minor: 0, Patch: 1284}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) OPiOS/10.2.0.93022 Mobile/12H143 Safari/9537.53",
		UserAgent{
			Browser{browser.BrowserType_Opera, Version{versioning.Version{Major: 10, Minor: 2, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 8, Minor: 4, Patch: 0}}}, device.DeviceType_phone}},

	// Internet Explorer -- https://msdn.microsoft.com/en-us/library/hh869301(v=vs.85).aspx
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.123",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 12, Minor: 123, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 2, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 11, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 3, Patch: 0}}}, device.DeviceType_computer}},

    {"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 EdgiOS/44.3.5 Mobile/15E148 Safari/605.1.15",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 12, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 12, Minor: 3, Patch: 1}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (iPad; CPU OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 EdgiOS/44.3.2 Mobile/15E148 Safari/605.1.15",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 12, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 12, Minor: 3, Patch: 1}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (Linux; Android 9; motorola one) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.90 Mobile Safari/537.36 EdgA/42.0.2.3728",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 42, Minor: 0, Patch: 2}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 9, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3800.0 Safari/537.36 Edg/76.0.172.0",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 76, Minor: 0, Patch: 172}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3803.0 Safari/537.36 Edg/76.0.176.0",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 76, Minor: 0, Patch: 176}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 14, Patch: 5}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; DEVICE INFO) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Edge/12.123",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 12, Minor: 123, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 520) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 11, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 8, Minor: 1, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705; .NET CLR 2.0.50727)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 5, Minor: 0, Patch: 1}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 5, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; GTB6.4; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; OfficeLiveConnector.1.3; OfficeLivePatch.0.0; .NET CLR 1.1.4322)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)", //Windows Surface RT tablet
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 2, Patch: 0}}}, device.DeviceType_tablet}},

	// UC Browser
	{"Mozilla/5.0 (Linux; U; Android 2.3.4; en-US; MT11i Build/4.0.2.A.0.62) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.0.1.275 U3/0.8.0 Mobile Safari/534.31",
		UserAgent{
			Browser{browser.BrowserType_UC_Browser, Version{versioning.Version{Major: 9, Minor: 0, Patch: 1}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 3, Patch: 4}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 4.0.4; en-US; Micromax P255 Build/IMM76D) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.2.0.308 U3/0.8.0 Mobile Safari/534.31",
		UserAgent{
			Browser{browser.BrowserType_UC_Browser, Version{versioning.Version{Major: 9, Minor: 2, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 4}}}, device.DeviceType_phone}},

	{"UCWEB/2.0 (Java; U; MIDP-2.0; en-US; MicromaxQ5) U2/1.0.0 UCBrowser/9.4.0.342 U2/1.0.0 Mobile",
		UserAgent{
			Browser{browser.BrowserType_UC_Browser, Version{versioning.Version{Major: 9, Minor: 4, Patch: 0}}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	// Nokia Browser
	// {"Mozilla/5.0 (Series40; Nokia501/14.0.4/java_runtime_version=Nokia_Asha_1_2; Profile/MIDP-2.1 Configuration/CLDC-1.1) Gecko/20100401 S40OviBrowser/4.0.0.0.45",
	// 	UserAgent{
	//		Browser{browser.BrowserType_unknown, Version{4,0,0}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{0,0,0}}, device.DeviceType_phone}},

	// {"Mozilla/5.0 (Symbian/3; Series60/5.3 NokiaN8-00/111.040.1511; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/535.1 (KHTML, like Gecko) NokiaBrowser/8.3.1.4 Mobile Safari/535.1",
	// 	UserAgent{
	//		Browser{browser.BrowserType_unknown, Version{8,0,0}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{0,0,0}}, device.DeviceType_phone}},

	// {"NokiaN97/21.1.107 (SymbianOS/9.4; Series60/5.0 Mozilla/5.0; Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebkit/525 (KHTML, like Gecko) BrowserNG/7.1.4",
	// 	browser.BrowserType_unknown, Version{7,0,0}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{0,0,0}}, device.DeviceType_phone}},

	// ChromeOS
	{"Mozilla/5.0 (X11; U; CrOS i686 9.10.0; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.253.0 Safari/532.5",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 4, Minor: 0, Patch: 253}}}, OS{platform.Platform_Linux, os.OS_Type_Chrome_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	// iPod, iPod Touch
	{"mozilla/5.0 (ipod touch; cpu iphone os 9_3_3 like mac os x) applewebkit/601.1.46 (khtml, like gecko) version/9.0 mobile/13g34 safari/601.1",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 9, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPod, os.OS_Type_iOS, Version{versioning.Version{Major: 9, Minor: 3, Patch: 3}}}, device.DeviceType_tablet}},

	{"mozilla/5.0 (ipod; cpu iphone os 6_1_6 like mac os x) applewebkit/536.26 (khtml, like gecko) version/6.0 mobile/10b500 safari/8536.25",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPod, os.OS_Type_iOS, Version{versioning.Version{Major: 6, Minor: 1, Patch: 6}}}, device.DeviceType_tablet}},

	// WebOS
	{"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.0; U; de-DE) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/233.70 Safari/534.6 TouchPad/1.0",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Web_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (webOS/1.4.1.1; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.0",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 1, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Web_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	// Android WebView (Android <= 4.3)
	{"Mozilla/5.0 (Linux; U; Android 2.2; en-us; DROID2 GLOBAL Build/S273) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 2, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 4.0.3; de-ch; HTC Sensation Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari53/4.30",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 3}}}, device.DeviceType_phone}},

	// BlackBerry
	{"Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML, like Gecko) Version/7.2.1.0 Safari/536.2+",
		UserAgent{
			Browser{browser.BrowserType_Blackberry, Version{versioning.Version{Major: 7, Minor: 2, Patch: 1}}}, OS{platform.Platform_Blackberry, os.OS_Type_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (BB10; Kbd) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.2.1.1925 Mobile Safari/537.35+",
		UserAgent{
			Browser{browser.BrowserType_Blackberry, Version{versioning.Version{Major: 10, Minor: 2, Patch: 1}}}, OS{platform.Platform_Blackberry, os.OS_Type_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) BlackBerry8703e/4.1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/104",
		UserAgent{
			Browser{browser.BrowserType_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Blackberry, os.OS_Type_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	// Windows Phone
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 625; ANZ941)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 8, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 900)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 9, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 7, Minor: 5, Patch: 0}}}, device.DeviceType_phone}},

	// Kindle eReader
	{"Mozilla/5.0 (Linux; U; en-US) AppleWebKit/528.5+ (KHTML, like Gecko, Safari/528.5+) Version/4.0 Kindle/3.0 (screen 600×800; rotate)",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (X11; U; Linux armv7l like Android; en-us) AppleWebKit/531.2+ (KHTML, like Gecko) Version/5.0 Safari/533.2+ Kindle/3.0+",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 5, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	// Amazon Fire
	{"Mozilla/5.0 (Linux; U; Android 4.4.3; de-de; KFTHWI Build/KTU84M) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.67 like Chrome/39.0.2171.93 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Silk, Version{versioning.Version{Major: 3, Minor: 67, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{versioning.Version{Major: 4, Minor: 4, Patch: 3}}}, device.DeviceType_tablet}}, // Fire tablet

	{"Mozilla/5.0 (Linux; U; Android 4.2.2; en­us; KFTHWI Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.22 like Chrome/34.0.1847.137 Mobile Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Silk, Version{versioning.Version{Major: 3, Minor: 22, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{versioning.Version{Major: 4, Minor: 2, Patch: 2}}}, device.DeviceType_tablet}}, // Fire tablet, but with "Mobile"

	{"Mozilla/5.0 (Linux; Android 4.4.4; SD4930UR Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/34.0.0.0 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/35.0.0.48.273;]",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 34, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{versioning.Version{Major: 4, Minor: 4, Patch: 4}}}, device.DeviceType_phone}}, // Facebook app on Fire Phone

	{"mozilla/5.0 (linux; android 4.4.3; kfthwi build/ktu84m) applewebkit/537.36 (khtml, like gecko) version/4.0 chrome/34.0.0.0 safari/537.36 [pinterest/android]",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 34, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{versioning.Version{Major: 4, Minor: 4, Patch: 3}}}, device.DeviceType_tablet}}, // Fire tablet running pinterest

	// extra logic to identify phone when using silk has not been added
	// {"Mozilla/5.0 (Linux; Android 4.4.4; SD4930UR Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.67 like Chrome/39.0.2171.93 Mobile Safari/537.36",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Silk, Version{3,0,0}}, OS{platform.Platform_Linux, os.OS_Type_Kindle, Version{4,0,0}}, device.DeviceType_phone}}, // Silk on Fire Phone

	// Nintendo
	{"Opera/9.30 (Nintendo Wii; U; ; 2047-7; fr)",
		UserAgent{
			Browser{browser.BrowserType_Opera, Version{versioning.Version{Major: 9, Minor: 30, Patch: 0}}}, OS{platform.Platform_Nintendo, os.OS_Type_Nintendo, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_console}},

	{"Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.US",
		UserAgent{
			Browser{browser.BrowserType_Nintendo, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Nintendo, os.OS_Type_Nintendo, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_console}},

	// Xbox
	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; Xbox)", //Xbox 360
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 9, Minor: 0, Patch: 0}}}, OS{platform.Platform_Xbox, os.OS_Type_Xbox, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, device.DeviceType_console}},

	// Playstation
	{"Mozilla/5.0 (PlayStation 4 4.50) AppleWebKit/601.2 (KHTML, like Gecko)",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Playstation, os.OS_Type_Playstation, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_console}},

	{"Mozilla/5.0 (Playstation Vita 1.61) AppleWebKit/531.22.8 (KHTML, like Gecko) Silk/3.2",
		UserAgent{
			Browser{browser.BrowserType_Silk, Version{versioning.Version{Major: 3, Minor: 2, Patch: 0}}}, OS{platform.Platform_Playstation, os.OS_Type_Playstation, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_console}},

	// Smart TVs and TV dongles
	{"Mozilla/5.0 (CrKey armv7l 1.4.15250) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.0 Safari/537.36", // Chromecast
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 31, Minor: 0, Patch: 1650}}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tv}},

	{"Mozilla/5.0 (Linux; GoogleTV 3.2; VAP430 Build/MASTER) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.77 Safari/534.24", // Google TV
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 11, Minor: 0, Patch: 696}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tv}},

	{"Mozilla/5.0 (Linux; Android 5.0; ADT-1 Build/LPX13D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Mobile Safari/537.36", // Android TV
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 40, Minor: 0, Patch: 2214}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 5, Minor: 0, Patch: 0}}}, device.DeviceType_tv}},

	{"Mozilla/5.0 (Linux; Android 4.2.2; AFTB Build/JDQ39) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.173 Mobile Safari/537.22", // Amazon Fire
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 25, Minor: 0, Patch: 1364}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 2, Patch: 2}}}, device.DeviceType_tv}},

	{"Mozilla/5.0 (Unknown; Linux armv7l) AppleWebKit/537.1+ (KHTML, like Gecko) Safari/537.1+ LG Browser/6.00.00(+mouse+3D+SCREEN+TUNER; LGE; GLOBAL-PLAT5; 03.07.01; 0x00000001;); LG NetCast.TV-2013/03.17.01 (LG, GLOBAL-PLAT4, wired)", // LG TV
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tv}},

	{"Mozilla/5.0 (X11; FreeBSD; U; Viera; de-DE) AppleWebKit/537.11 (KHTML, like Gecko) Viera/3.10.0 Chrome/23.0.1271.97 Safari/537.11", // Panasonic Viera
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 23, Minor: 0, Patch: 1271}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tv}},

	// TODO: not catching "browser/" and reporting as safari -- ua string not being fully checked?
	// {"Mozilla/5.0 (DTV) AppleWebKit/531.2+ (KHTML, like Gecko) Espial/6.1.5 AQUOSBrowser/2.0 (US01DTV;V;0001;0001)", // Sharp Aquos
	// 	browser.BrowserType_unknown, Version{0,0,0}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{0,0,0}}, device.DeviceType_tv}},

	{"Roku/DVP-5.2 (025.02E03197A)", // Roku
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tv}},

	{"mozilla/5.0 (smart-tv; linux; tizen 2.3) applewebkit/538.1 (khtml, like gecko) samsungbrowser/1.0 tv safari/538.1", // Samsung SmartTV
		UserAgent{
			Browser{browser.BrowserType_Samsung, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tv}},

	{"mozilla/5.0 (linux; u) applewebkit/537.36 (khtml, like gecko) version/4.0 mobile safari/537.36 smarttv/6.0 (netcast)",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tv}},

	// Google search app (GSA) for iOS -- it's Safari in disguise as of v6
	{"Mozilla/5.0 (iPad; CPU OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) GSA/6.0.51363 Mobile/12F69 Safari/600.1.4",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 8, Minor: 3, Patch: 0}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 8, Minor: 3, Patch: 0}}}, device.DeviceType_tablet}},

	// Spotify (applicable for advertising applications)
	{"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Spotify/1.0.9.133 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Spotify, Version{versioning.Version{Major: 1, Minor: 0, Patch: 9}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 5, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Spotify/1.0.9.133 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Spotify, Version{versioning.Version{Major: 1, Minor: 0, Patch: 9}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 10, Patch: 2}}}, device.DeviceType_computer}},

	// OCSP fetchers
	{"Microsoft-CryptoAPI/10.0",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},
	{"trustd (unknown version) CFNetwork/811.7.2 Darwin/16.7.0 (x86_64)",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},
	{"ocspd (unknown version) CFNetwork/520.5.3 Darwin/11.4.2 (x86_64)(MacBookAir5%2C2)",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},
	// Bots
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)",
		UserAgent{
			Browser{browser.BrowserType_AppleBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 10, Minor: 10, Patch: 1}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
		UserAgent{
			Browser{browser.BrowserType_BaiduBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
		UserAgent{
			Browser{browser.BrowserType_BingBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"DuckDuckBot/1.0; (+http://duckduckgo.com/duckduckbot.html)",
		UserAgent{
			Browser{browser.BrowserType_DuckDuckGoBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
		UserAgent{
			Browser{browser.BrowserType_FacebookBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Facebot/1.0",
		UserAgent{
			Browser{browser.BrowserType_FacebookBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		UserAgent{
			Browser{browser.BrowserType_GoogleBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/3.1 +http://www.linkedin.com)",
		UserAgent{
			Browser{browser.BrowserType_LinkedInBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"msnbot/2.0b (+http://search.msn.com/msnbot.htm)",
		UserAgent{
			Browser{browser.BrowserType_MsnBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Pingdom.com_bot_version_1.4_(http://www.pingdom.com/)",
		UserAgent{
			Browser{browser.BrowserType_PingdomBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Twitterbot/1.0",
		UserAgent{
			Browser{browser.BrowserType_TwitterBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)",
		UserAgent{
			Browser{browser.BrowserType_YandexBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
		UserAgent{
			Browser{browser.BrowserType_YahooBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"{UA:Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)}, ua: &{Browser:{Name:BrowserGoogleBot Version:{Major:0 Minor:0 Patch:0}} OS:{Platform:platform.Platform_Bot Name:os.OS_Type_Bot Version:{Major:6 Minor:0 Patch:1}} DeviceType:device.DeviceType_computer}",
		UserAgent{
			Browser{browser.BrowserType_GoogleBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 6, Minor: 0, Patch: 1}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		UserAgent{
			Browser{browser.BrowserType_GoogleBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"mozilla/5.0 (unknown; linux x86_64) applewebkit/538.1 (khtml, like gecko) phantomjs/2.1.1 safari/538.1",
		UserAgent{
			Browser{browser.BrowserType_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	// Unknown or partially handled
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.4; en-US; rv:1.9.1b3pre) Gecko/20090223 SeaMonkey/2.0a3", //Seamonkey (~FF)
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 4, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en; rv:1.9.0.8pre) Gecko/2009022800 Camino/2.0b3pre", //Camino (~FF)
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 5, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0", //firefox OS
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 26, Minor: 0, Patch: 0}}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.45 Safari/535.19", //chrome for android having requested desktop site
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 18, Minor: 0, Patch: 1025}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		UserAgent{
			Browser{browser.BrowserType_Opera, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	// browser.BrowserType_QQ
	{"Mozilla/5.0 (Windows NT 6.2; WOW64; Trident/7.0; Touch; .NET4.0E; .NET4.0C; .NET CLR 3.5.30729; .NET CLR 2.0.50727; .NET CLR 3.0.30729; InfoPath.3; Tablet PC 2.0; QQBrowser/7.6.21433.400; rv:11.0) like Gecko",
		UserAgent{
			Browser{browser.BrowserType_QQ, Version{versioning.Version{Major: 7, Minor: 6, Patch: 21433}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 2, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36 QQBrowser/9.0.2191.400",
		UserAgent{
			Browser{browser.BrowserType_QQ, Version{versioning.Version{Major: 9, Minor: 0, Patch: 2191}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	{"mozilla/5.0 (iphone; cpu iphone os 8_1_2 like mac os x) applewebkit/600.1.4 (khtml, like gecko) mobile/12b440 qq/5.3.0.319 nettype/wifi mem/205",
		UserAgent{
			Browser{browser.BrowserType_QQ, Version{versioning.Version{Major: 5, Minor: 3, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 8, Minor: 1, Patch: 2}}}, device.DeviceType_phone}},

	// ANDROID TESTS

	{"Mozilla/5.0 (Linux; U; Android 1.0; en-us; dream) AppleWebKit/525.10+ (KHTML,like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 4}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 1.0; en-us; generic) AppleWebKit/525.10 (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 4}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 1.0.3; de-de; A80KSC Build/ECLAIR) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 0, Patch: 3}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 1.5; en-gb; T-Mobile G1 Build/CRC1) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 1, Patch: 2}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 5, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 1.5; es-; FBW1_4 Build/MASTER) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 4}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 5, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux U; Android 1.5 en-us hero) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 4}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 5, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 1.5; en-us; Opus One Build/RBE.00.00) AppleWebKit/528.18.1 (KHTML, like Gecko) Version/3.1.1 Mobile Safari/525.20.1",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 1, Patch: 1}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 5, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 1.6; ar-us; SonyEricssonX10i Build/R2BA026) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 1, Patch: 2}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 6, Patch: 0}}}, device.DeviceType_phone}},

	// TODO: support names of Android OS?
	//{"Mozilla/5.0 (Linux; U; Android Donut; de-de; HTC Tattoo 1.52.161.1 Build/Donut) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
	//	UserAgent{
	//		Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 1, Patch: 2}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 1.6; en-gb; HTC Tattoo Build/DRC79) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 4}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 6, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 1.6; ja-jp; Docomo HT-03A Build/DRD08) AppleWebKit/525.10 (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 4}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 1, Minor: 6, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 2.1; en-us; Nexus One Build/ERD62) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 1, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 2.1-update1; en-au; HTC_Desire_A8183 V1.16.841.1 Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 1, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 2.1; en-us; generic) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 4}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 1, Patch: 0}}}, device.DeviceType_phone}},

	// TODO support named versions of Android?
	{"Mozilla/5.0 (Linux; U; Android Eclair; en-us; sholes) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 4}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 2.2; en-sa; HTC_DesireHD_A9191 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 2, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 2.2.1; en-gb; HTC_DesireZ_A7272 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 2, Patch: 1}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; Sensation_4G Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Safari/533.16",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 5, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 3, Patch: 3}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 2.3.5; ko-kr; SHW-M250S Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 3, Patch: 5}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 2.3.7; ja-jp; L-02D Build/GWK74) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 2, Minor: 3, Patch: 7}}}, device.DeviceType_phone}},

	// TODO: is tablet, not phone
	{"Mozilla/5.0 (Linux; U; Android 3.0; xx-xx; Transformer TF101 Build/HRI66) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 3.0; en-us; Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 3, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (Linux; U; Android 4.0.1; en-us; sdk Build/ICS_MR0) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 1}}}, device.DeviceType_phone}},

	// TODO support "android-" version prefix
	// However, can't find reference to this naming scheme in real-world UA gathering
	// {"Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Chrome, Version{16,0,0}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{4,0,0}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Nexus S Build/JRO03E) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 1, Patch: 1}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 4.1; en-gb; Build/JRN84D) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 1, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; el-gr; MB525 Build/JRO03H; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 1, Patch: 1}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; fr-fr; MB525 Build/JRO03H; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 1, Patch: 1}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; U; Android 4.2; en-us; Nexus 10 Build/JVP15I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 2, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (Linux; U; Android 4.2; ro-ro; LT18i Build/4.1.B.0.431) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{browser.BrowserType_Android, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 2, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; Android 4.3; Nexus 7 Build/JWR66D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.111 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 27, Minor: 0, Patch: 1453}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 3, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (Linux; Android 4.4; Nexus 7 Build/KOT24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.105 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 30, Minor: 0, Patch: 1599}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 4, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (Linux; Android 4.4; Nexus 4 Build/KRT16E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.105 Mobile Safari",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 30, Minor: 0, Patch: 1599}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 4, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; Android 6.0.1; SM-G930V Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 52, Minor: 0, Patch: 2743}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 6, Minor: 0, Patch: 1}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; Android 7.0; Nexus 5X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 52, Minor: 0, Patch: 2743}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Linux; Android 7.0; Nexus 6P Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.98 Mobile Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 52, Minor: 0, Patch: 2743}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	// BLACKBERRY TESTS

	{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) BlackBerry8703e/4.1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/104",
		UserAgent{
			Browser{browser.BrowserType_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Blackberry, os.OS_Type_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.1.0.4633 Mobile Safari/537.10+",
		UserAgent{
			Browser{browser.BrowserType_Blackberry, Version{versioning.Version{Major: 10, Minor: 1, Patch: 0}}}, OS{platform.Platform_Blackberry, os.OS_Type_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (BB10; Kbd) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.2.1.1925 Mobile Safari/537.35+",
		UserAgent{
			Browser{browser.BrowserType_Blackberry, Version{versioning.Version{Major: 10, Minor: 2, Patch: 1}}}, OS{platform.Platform_Blackberry, os.OS_Type_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (PlayBook; U; RIM Tablet OS 1.0.0; en-US) AppleWebKit/534.11 (KHTML, like Gecko) Version/7.1.0.7 Safari/534.11",
		UserAgent{
			Browser{browser.BrowserType_Blackberry, Version{versioning.Version{Major: 7, Minor: 1, Patch: 0}}}, OS{platform.Platform_Blackberry, os.OS_Type_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML, like Gecko) Version/7.2.1.0 Safari/536.2+",
		UserAgent{
			Browser{browser.BrowserType_Blackberry, Version{versioning.Version{Major: 7, Minor: 2, Patch: 1}}}, OS{platform.Platform_Blackberry, os.OS_Type_Blackberry, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (X11; U; CrOS i686 9.10.0; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.253.0 Safari/532.5",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 4, Minor: 0, Patch: 253}}}, OS{platform.Platform_Linux, os.OS_Type_Chrome_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (X11; CrOS armv7l 5500.100.6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.120 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 34, Minor: 0, Patch: 1847}}}, OS{platform.Platform_Linux, os.OS_Type_Chrome_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	// {"Mozilla/5.0 (Mobile; rv:14.0) Gecko/14.0 Firefox/14.0",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Firefox, 14, OSFirefoxOS, 14}, device.DeviceType_phone}},

	// {"Mozilla/5.0 (Mobile; rv:17.0) Gecko/17.0 Firefox/17.0",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Firefox, , OSFirefoxOS}, device.DeviceType_phone}},

	// {"Mozilla/5.0 (Mobile; rv:18.1) Gecko/18.1 Firefox/18.1",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Firefox, , OSFirefoxOS}, device.DeviceType_phone}},

	// {"Mozilla/5.0 (Tablet; rv:18.1) Gecko/18.1 Firefox/18.1",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Firefox, , OSFirefoxOS}, device.DeviceType_phone}},

	// {"Mozilla/5.0 (Mobile; LG-D300; rv:18.1) Gecko/18.1 Firefox/18.1",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Firefox, , OSFirefoxOS}, device.DeviceType_phone}},

	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 4, Minor: 0, Patch: 4}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 3, Minor: 2, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 4, Minor: 0, Patch: 5}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 4, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 5, Minor: 1, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 5, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (iPad; CPU OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 5, Minor: 1, Patch: 0}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 5, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (iPad; CPU OS 7_0_2 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A501 Safari/9537.53",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{versioning.Version{Major: 7, Minor: 0, Patch: 2}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_2_1 like Mac OS X) AppleWebKit/602.4.6 (KHTML, like Gecko) Mobile/14D27 [FBAN/FBIOS;FBAV/86.0.0.48.52;FBBV/53842252;FBDV/iPhone9,1;FBMD/iPhone;FBSN/iOS;FBSV/10.2.1;FBSS/2;FBCR/Verizon;FBID/phone;FBLC/en_US;FBOP/5;FBRV/0]",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 10, Minor: 2, Patch: 1}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 10, Minor: 2, Patch: 1}}}, device.DeviceType_phone}},

	// TODO handle default browser based on iOS version
	// {"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/538.34.9 (KHTML, like Gecko) Mobile/12A4265u",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Safari, Version{8,0,0}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{8,0,0}}, device.DeviceType_phone}},

	// TODO extrapolate browser from iOS version
	// {"Mozilla/5.0 (iPad; CPU OS 8_0 like Mac OS X) AppleWebKit/538.34.9 (KHTML, like Gecko) Mobile/12A4265u",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Safari, Version{8,0,0}}, OS{platform.Platform_iPad, os.OS_Type_iOS, Version{8,0,0}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 8, Minor: 0, Patch: 0}}}, OS{platform.Platform_iPhone, os.OS_Type_iOS, Version{versioning.Version{Major: 8, Minor: 0, Patch: 2}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (X11; U; Linux x86_64; en; rv:1.9.0.14) Gecko/20080528 Ubuntu/9.10 (karmic) Epiphany/2.22 Firefox/3.0",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 3, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	// Can't parse browser due to limitation of user agent library
	{"Mozilla/5.0 (X11; U; Linux x86_64; zh-TW; rv:1.9.0.8) Gecko/2009032712 Ubuntu/8.04 (hardy) Firefox/3.0.8 GTB5",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 3, Minor: 0, Patch: 8}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; Konqueror/3.5; Linux; x86_64) KHTML/3.5.5 (like Gecko) (Debian)",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (X11; U; Linux i686; de; rv:1.9.1.5) Gecko/20091112 Iceweasel/3.5.5 (like Firefox/3.5.5; Debian-3.5.5-1)",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 3, Minor: 5, Patch: 5}}}, OS{platform.Platform_Linux, os.OS_Type_Linux, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	// TODO consider bot?
	// {"Miro/2.0.4 (http://www.getmiro.com/; Darwin 10.3.0 i386)",
	// 	UserAgent{
	//		Browser{browser.BrowserType_unknown, Version{0,0,0}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{3,0,0}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.4; en-US; rv:1.9.1b3pre) Gecko/20090223 SeaMonkey/2.0a3",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 4, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 3, Minor: 2, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 5, Patch: 5}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en; rv:1.9.0.8pre) Gecko/2009022800 Camino/2.0b3pre",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 5, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_2; en-US) AppleWebKit/533.1 (KHTML, like Gecko) Chrome/5.0.329.0 Safari/533.1",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 5, Minor: 0, Patch: 329}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 6, Patch: 2}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6 (.NET CLR 3.5.30729)",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 3, Minor: 5, Patch: 6}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 6, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 5, Minor: 1, Patch: 2}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 7, Patch: 2}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:9.0) Gecko/20111222 Thunderbird/9.0.1",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 7, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 16, Minor: 0, Patch: 912}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 7, Patch: 2}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8) AppleWebKit/535.18.5 (KHTML, like Gecko) Version/5.2 Safari/535.18.5",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 5, Minor: 2, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 8, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_8; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.249.0 Safari/532.5",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 4, Minor: 0, Patch: 249}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 8, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9) AppleWebKit/537.35.1 (KHTML, like Gecko) Version/6.1 Safari/537.35.1",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 9, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.34.48 (KHTML, like Gecko) Version/8.0 Safari/538.35.8",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 8, Minor: 0, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 10, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.32 (KHTML, like Gecko) Version/7.1 Safari/538.4",
		UserAgent{
			Browser{browser.BrowserType_Safari, Version{versioning.Version{Major: 7, Minor: 1, Patch: 0}}}, OS{platform.Platform_Mac, os.OS_Type_Mac_OSX, Version{versioning.Version{Major: 10, Minor: 10, Patch: 0}}}, device.DeviceType_computer}},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		UserAgent{
			Browser{browser.BrowserType_Opera, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		UserAgent{
			Browser{browser.BrowserType_Opera, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	// TODO: support OneBrowser? https://play.google.com/store/apps/details?id=com.tencent.ibibo.mtt&hl=en_GB
	// {"OneBrowser/3.1 (NokiaN70-1/5.0638.3.0.1)",
	// 	UserAgent{
	//		Browser{browser.BrowserType_unknown, Version{0,0,0}}, OS{platform.Platform_Unknown, os.OS_Type_unknown, Version{0,0,0}}, device.DeviceType_phone}},

	// WebOS reports itself as safari :(
	{"Mozilla/5.0 (webOS/1.0; U; en-US) AppleWebKit/525.27.1 (KHTML, like Gecko) Version/1.0 Safari/525.27.1 Pre/1.0",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 1, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Web_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (webOS/1.4.1.1; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.0",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 1, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Web_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.0; U; de-DE) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/233.70 Safari/534.6 TouchPad/1.0",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Web_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.2; U; en-US) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/234.40.1 Safari/534.6 TouchPad/1.0",
		UserAgent{
			Browser{browser.BrowserType_unknown, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Linux, os.OS_Type_Web_OS, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_tablet}},

	{"Opera/9.30 (Nintendo Wii; U; ; 2047-7; fr)",
		UserAgent{
			Browser{browser.BrowserType_Opera, Version{versioning.Version{Major: 9, Minor: 30, Patch: 0}}}, OS{platform.Platform_Nintendo, os.OS_Type_Nintendo, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_console}},

	{"Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.US",
		UserAgent{
			Browser{browser.BrowserType_Nintendo, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Nintendo, os.OS_Type_Nintendo, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_console}},

	{"Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.28 (KHTML, like Gecko) NX/3.0.3.12.6 NintendoBrowser/2.0.0.9362.US",
		UserAgent{
			Browser{browser.BrowserType_Nintendo, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Nintendo, os.OS_Type_Nintendo, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_console}},

	// TODO fails to get opera first -- but is this a real UA string or an uncommon spoof?
	// {"Mozilla/4.0 (compatible; MSIE 5.0; Windows 2000) Opera 6.0 [en]",
	// 	browser.BrowserType_Internet_Explorer, Version{5,0,0}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{4,0,0}}, device.DeviceType_computer}},

	{"Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705; .NET CLR 2.0.50727)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 5, Minor: 0, Patch: 1}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 5, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; GTB6.4; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; OfficeLiveConnector.1.3; OfficeLivePatch.0.0; .NET CLR 1.1.4322)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Windows; U; Windows NT 6.1; sk; rv:1.9.1.7) Gecko/20091221 Firefox/3.5.7",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 3, Minor: 5, Patch: 7}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 2, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) YaBrowser/1.0.1084.5402 Chrome/19.0.1084.5402 Safari/536.5",
		UserAgent{
			Browser{browser.BrowserType_Yandex, Version{versioning.Version{Major: 1, Minor: 0, Patch: 1084}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 2, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.15 (KHTML, like Gecko) Chrome/24.0.1295.0 Safari/537.15",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 24, Minor: 0, Patch: 1295}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 2, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 11, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 3, Patch: 0}}}, device.DeviceType_tablet}},

	{"Mozilla/5.0 (IE 11.0; Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 11, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 3, Patch: 0}}}, device.DeviceType_computer}},

	// {"Mozilla/4.0 (compatible; MSIE 4.01; Windows 95)",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Internet_Explorer, Version{5,0,0}}, OS{platform.Platform_Windows, os.OS_Type_Windows95, Version{5,0,0}}, device.DeviceType_computer}},

	// {"Mozilla/4.0 (compatible; MSIE 5.0; Windows 95) Opera 6.02 [en]",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Internet_Explorer, Version{5,0,0}}, OS{platform.Platform_Windows, os.OS_Type_Windows95, Version{5,0,0}}, device.DeviceType_computer}},

	// {"Mozilla/4.0 (compatible; MSIE 6.0b; Windows 98; YComp 5.0.0.0)",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Internet_Explorer, Version{6,0,0}}, OS{platform.Platform_Windows, os.OS_Type_Windows98, Version{5,0,0}}, device.DeviceType_computer}},

	// {"Mozilla/4.0 (compatible; MSIE 4.01; Windows 98)",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Internet_Explorer, Version{4,0,0}}, OS{platform.Platform_Windows, os.OS_Type_Windows98, Version{5,0,0}}, device.DeviceType_computer}},

	// {"Mozilla/5.0 (Windows; U; Windows 98; en-US; rv:1.8.1.8pre) Gecko/20071019 Firefox/2.0.0.8 Navigator/9.0.0.1",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Firefox, Version{2,0,0}}, OS{platform.Platform_Windows, os.OS_Type_Windows98, Version{5,0,0}}, device.DeviceType_computer}},

	//Can't parse due to limitation of user agent library
	// {"Mozilla/5.0 (Windows; U; Windows CE 5.1; rv:1.8.1a3) Gecko/20060610 Minimo/0.016",
	// UserAgent{
	//		Browser{	browser.BrowserType_unknown, Version{0,0,0}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{0,0,0}}, device.DeviceType_phone}},

	// {"Mozilla/4.0 (compatible; MSIE 4.01; Windows CE; 176x220)",
	// 	UserAgent{
	//		Browser{browser.BrowserType_Internet_Explorer, Version{4,0,0}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{0,0,0}}, device.DeviceType_phone}},

	// Can't parse browser due to limitation of user agent library
	// {"Mozilla/4.0 (compatible; MSIE 5.0; Windows ME) Opera 6.0 [de]",
	// 	UserAgent{
	//		Browser{browser.BrowserType_unknown, os.OS_Type_WindowsME}, device.DeviceType_computer}},

	{"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 1.1.4322; InfoPath.2; .NET CLR 3.5.21022; .NET CLR 3.5.30729; MS-RTC LM 8; OfficeLiveConnector.1.4; OfficeLivePatch.1.3; .NET CLR 3.0.30729)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 8, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Windows; U; Windows NT 5.1; cs; rv:1.9.1.8) Gecko/20100202 Firefox/3.5.8",
		UserAgent{
			Browser{browser.BrowserType_Firefox, Version{versioning.Version{Major: 3, Minor: 5, Patch: 8}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 5, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; )",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 5, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	// Can't parse due to limitation of user agent library
	{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Windows Phone 6.5.3.5)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 6, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 6, Minor: 5, Patch: 3}}}, device.DeviceType_phone}},

	// desktop mode for Windows Phone 7
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; XBLWP7; ZuneWP7)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, device.DeviceType_computer}},

	// mobile mode for Windows Phone 7
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; HTC; T8788)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 7, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 9, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 7, Minor: 5, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 920)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 10, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 8, Minor: 0, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch IEMobile/11.0; HTC; Windows Phone 8S by HTC) like Gecko",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 11, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 8, Minor: 1, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch IEMobile/11.0; NOKIA; 909) like Gecko",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 11, Minor: 0, Patch: 0}}}, OS{platform.Platform_WindowsPhone, os.OS_Type_Windows_Phone, Version{versioning.Version{Major: 8, Minor: 1, Patch: 0}}}, device.DeviceType_phone}},

	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; Xbox)",
		UserAgent{
			Browser{browser.BrowserType_Internet_Explorer, Version{versioning.Version{Major: 9, Minor: 0, Patch: 0}}}, OS{platform.Platform_Xbox, os.OS_Type_Xbox, Version{versioning.Version{Major: 6, Minor: 1, Patch: 0}}}, device.DeviceType_console}},
	{"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) coc_coc_browser/42.0 CoRom/36.0.1985.144 Chrome/36.0.1985.144 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Coc_Coc, Version{versioning.Version{Major: 42, Minor: 0, Patch: 0}}}, OS{platform.Platform_Windows, os.OS_Type_Windows, Version{versioning.Version{Major: 6, Minor: 3, Patch: 0}}}, device.DeviceType_computer}},
	{"Mozilla/5.0 (compatible; coccocbot/1.0; +http://help.coccoc.com/searchengine)",
		UserAgent{
			Browser{browser.BrowserType_CocCocBot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, OS{platform.Platform_Bot, os.OS_Type_Bot, Version{versioning.Version{Major: 0, Minor: 0, Patch: 0}}}, device.DeviceType_computer}},

	{"Mozilla/5.0 (Linux; Android 4.4.4; SM-T560 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 63, Minor: 0, Patch: 3239}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 4, Patch: 4}}}, device.DeviceType_tablet}},
	{"Mozilla/5.0 (Linux; Android 5.1.1; KFSUWI) AppleWebKit/537.36 (KHTML, like Gecko) Silk/70.4.2 like Chrome/70.0.3538.80 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Silk, Version{versioning.Version{Major: 70, Minor: 4, Patch: 2}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 5, Minor: 1, Patch: 1}}}, device.DeviceType_tablet}},
	{"Mozilla/5.0 (Linux; Android 4.4.2; T1-701u Build/HuaweiMediaPad) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.123 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 64, Minor: 0, Patch: 3282}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 4, Patch: 2}}}, device.DeviceType_tablet}},
	{"Mozilla/5.0 (Linux; Android 4.4.2; Lenovo TAB 2 A7-30F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.84 Safari/537.36",
		UserAgent{
			Browser{browser.BrowserType_Chrome, Version{versioning.Version{Major: 45, Minor: 0, Patch: 2454}}}, OS{platform.Platform_Linux, os.OS_Type_Android, Version{versioning.Version{Major: 4, Minor: 4, Patch: 2}}}, device.DeviceType_tablet}},
}

func TestAgentSurfer(t *testing.T) {
	for _, determined := range testUAVars {
		t.Run("", func(t *testing.T) {
			testFuncs := []func(string) *UserAgent{
				Parse,
				func(ua string) *UserAgent {
					u := new(UserAgent)
					ParseUserAgent(ua, u)
					return u
				},
			}

			for _, f := range testFuncs {
				ua := f(determined.UA)

				if ua.Browser.Type != determined.Browser.Type {
					t.Errorf("browserName: got %v, wanted %v", ua.Browser.Type, determined.Browser.Type)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.Browser.Version.Major != determined.Browser.Version.Major || ua.Browser.Version.Minor != determined.Browser.Version.Minor || ua.Browser.Version.Patch != determined.Browser.Version.Patch {
					t.Errorf("browser version: got %v, wanted %v", ua.Browser.Version, determined.Browser.Version)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.OS.Platform != determined.OS.Platform {
					t.Errorf("platform: got %v, wanted %v", ua.OS.Platform, determined.OS.Platform)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.OS.Type != determined.OS.Type {
					t.Errorf("os: got %s, wanted %s", ua.OS.Type, determined.OS.Type)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.OS.Version.Major != determined.OS.Version.Major || ua.OS.Version.Minor != determined.OS.Version.Minor || ua.OS.Version.Patch != determined.OS.Version.Patch {
					t.Errorf("os version: got %v, wanted %v", ua.OS.Version, determined.OS.Version)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.Device != determined.Device {
					t.Errorf("device type: got %v, wanted %v", ua.Device, determined.Device)
					t.Logf("agent: %s", determined.UA)
				}
			}
		})
	}
}

func BenchmarkAgentSurfer(b *testing.B) {
	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(testUAVars[i%num].UA)
	}
}

func BenchmarkAgentSurferReuse(b *testing.B) {
	dest := new(UserAgent)
	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest.Reset()
		ParseUserAgent(testUAVars[i%num].UA, dest)
	}
}

func BenchmarkEvalSystem(b *testing.B) {
	num := len(testUAVars)
	v := UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.evalOS(testUAVars[i%num].UA)
	}
}

func BenchmarkEvalBrowserName(b *testing.B) {
	num := len(testUAVars)
	v := UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.evalBrowserName(testUAVars[i%num].UA)
	}
}

func BenchmarkEvalBrowserVersion(b *testing.B) {
	num := len(testUAVars)
	v := UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Browser.Type = testUAVars[i%num].Browser.Type
		v.evalBrowserVersion(testUAVars[i%num].UA)
	}
}

func BenchmarkEvalDevice(b *testing.B) {
	num := len(testUAVars)
	v := UserAgent{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.OS.Type = testUAVars[i%num].OS.Type
		v.OS.Platform = testUAVars[i%num].OS.Platform
		v.Browser.Type = testUAVars[i%num].Browser.Type
		v.evalDevice(testUAVars[i%num].UA)
	}
}

// Chrome for Mac
func BenchmarkParseChromeMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36")
	}
}

// Chrome for Windows
func BenchmarkParseChromeWin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.134 Safari/537.36")
	}
}

// Chrome for Android
func BenchmarkParseChromeAndroid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (Linux; Android 4.4.2; GT-P5210 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.93 Safari/537.36")
	}
}

// Safari for Mac
func BenchmarkParseSafariMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12")
	}
}

// Safari for iPad
func BenchmarkParseSafariiPad(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (iPad; CPU OS 8_1_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B440 Safari/600.1.4")
	}
}

// func TestStringTrimPrefix(t *testing.T) {
// 	testCases := []struct {
// 		f        func() string
// 		expected string
// 	}{
// 		{
// 			f:        DeviceUnknown.StringTrimPrefix,
// 			expected: "Unknown",
// 		},
// 		{
// 			f:        browser.BrowserType_unknown.StringTrimPrefix,
// 			expected: "Unknown",
// 		},
// 		{
// 			f:        os.OS_Type_unknown.StringTrimPrefix,
// 			expected: "Unknown",
// 		},
// 		{
// 			f:        platform.Platform_Unknown.StringTrimPrefix,
// 			expected: "Unknown",
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		t.Run("", func(t *testing.T) {
// 			s := tc.f()
// 			if tc.expected != s {
// 				t.Fatalf("Expected %q, got %q", tc.expected, s)
// 			}
// 		})
// 	}
// }
