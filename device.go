package uasurfer

import (
	"strings"

	"github.com/nigeltiany/uasurfer/types/browser"
	"github.com/nigeltiany/uasurfer/types/device"
	"github.com/nigeltiany/uasurfer/types/os"
	"github.com/nigeltiany/uasurfer/types/platform"
)

func (u *UserAgent) evalDevice(ua string) {
	switch {

	case u.OS.Platform == platform.Platform_Windows || u.OS.Platform == platform.Platform_Mac || u.OS.Type == os.OS_Type_Chrome_OS:
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			u.Device = device.DeviceType_tablet // windows rt, linux haxor tablets
			return
		}
		u.Device = device.DeviceType_computer

	case u.OS.Platform == platform.Platform_iPad || u.OS.Platform == platform.Platform_iPod || strings.Contains(ua, "tablet") || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook"):
		u.Device = device.DeviceType_tablet

	case u.OS.Platform == platform.Platform_iPhone || u.OS.Platform == platform.Platform_Blackberry || strings.Contains(ua, "phone"):
		u.Device = device.DeviceType_phone

	// long list of smarttv and tv dongle identifiers
	case strings.Contains(ua, "tv") || strings.Contains(ua, "crkey") || strings.Contains(ua, "googletv") || strings.Contains(ua, "aftb") || strings.Contains(ua, "aftt") || strings.Contains(ua, "aftm") || strings.Contains(ua, "adt-") || strings.Contains(ua, "roku") || strings.Contains(ua, "viera") || strings.Contains(ua, "aquos") || strings.Contains(ua, "dtv") || strings.Contains(ua, "appletv") || strings.Contains(ua, "smarttv") || strings.Contains(ua, "tuner") || strings.Contains(ua, "smart-tv") || strings.Contains(ua, "hbbtv") || strings.Contains(ua, "netcast") || strings.Contains(ua, "vizio"):
		u.Device = device.DeviceType_tv

	case u.OS.Type == os.OS_Type_Android:
		// android phones report as "mobile", android tablets should not but often do -- http://android-developers.blogspot.com/2010/12/android-browser-user-agent-issues.html
		if strings.Contains(ua, "mobile") {
			u.Device = device.DeviceType_phone
			return
		}

		if strings.Contains(ua, "tablet") || strings.Contains(ua, "nexus 7") || strings.Contains(ua, "nexus 9") || strings.Contains(ua, "nexus 10") || strings.Contains(ua, "xoom") ||
			strings.Contains(ua, "sm-t") || strings.Contains(ua, "; kf") || strings.Contains(ua, "; t1") || strings.Contains(ua, "lenovo tab") {
			u.Device = device.DeviceType_tablet
			return
		}

		u.Device = device.DeviceType_phone // default to phone

	case u.OS.Platform == platform.Platform_Playstation || u.OS.Platform == platform.Platform_Xbox || u.OS.Platform == platform.Platform_Nintendo:
		u.Device = device.DeviceType_console

	case strings.Contains(ua, "glass") || strings.Contains(ua, "watch") || strings.Contains(ua, "sm-v"):
		u.Device = device.DeviceType_wearable

	// specifically above "mobile" string check as Kindle Fire tablets report as "mobile"
	case u.Browser.Type == browser.BrowserType_Silk || u.OS.Type == os.OS_Type_Kindle && !strings.Contains(ua, "sd4930ur"):
		u.Device = device.DeviceType_tablet

	case strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") || strings.Contains(ua, "webos"): //anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
		u.Device = device.DeviceType_phone

	case u.OS.Type == os.OS_Type_Linux: // linux goes last since it's in so many other device types (tvs, wearables, android-based stuff)
		u.Device = device.DeviceType_computer

	default:
		u.Device = device.DeviceType_unknown
	}
}
