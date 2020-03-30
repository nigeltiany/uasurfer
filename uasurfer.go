// Package uasurfer provides fast and reliable abstraction
// of HTTP User-Agent strings. The philosophy is to identify
// technologies that holds >1% market share, and to avoid
// expending resources and accuracy on guessing at esoteric UA
// strings.
package uasurfer

import (
	"strings"

	"github.com/nigeltiany/uasurfer/types/browser"
	"github.com/nigeltiany/uasurfer/types/device"
	"github.com/nigeltiany/uasurfer/types/os"
	"github.com/nigeltiany/uasurfer/types/platform"
)

type Version struct {
	Major int           `json:"major"`
	Minor int           `json:"minor"`
	Patch int           `json:"patch"`
}

func (v Version) Less(c Version) bool {
	if v.Major < c.Major {
		return true
	}

	if v.Major > c.Major {
		return false
	}

	if v.Minor < c.Minor {
		return true
	}

	if v.Minor > c.Minor {
		return false
	}

	return v.Patch < c.Patch
}

type UserAgent struct {
	Browser    Browser              `json:"browser"`
	OS         OS                   `json:"os"`
	Device     device.DeviceType    `json:"device"`
}

type Browser struct {
	Type    browser.BrowserType     `json:"type"`
	Version Version                 `json:"version"`
}

type OS struct {
	Platform platform.Platform      `json:"platform"`
	Type     os.OS_Type             `json:"type"`
	Version  Version                `json:"version"`
}

// Reset resets the UserAgent to it's zero value
func (ua *UserAgent) Reset() {
	ua.Browser = Browser{}
	ua.OS = OS{}
	ua.Device = device.DeviceType_unknown
}

// IsBot returns true if the UserAgent represent a bot
func (ua *UserAgent) IsBot() bool {
	// If withing bot reserve
	if ua.Browser.Type >= 20 && ua.Browser.Type <= 60 {
		return true
	}
	if ua.OS.Type == os.OS_Type_Bot {
		return true
	}
	if ua.OS.Platform == platform.Platform_Bot {
		return true
	}
	return false
}

// Parse accepts a raw user agent (string) and returns the UserAgent.
func Parse(ua string) *UserAgent {
	dest := new(UserAgent)
	parse(ua, dest)
	return dest
}

// ParseUserAgent is the same as Parse, but populates the supplied UserAgent.
// It is the caller's responsibility to call Reset() on the UserAgent before
// passing it to this function.
func ParseUserAgent(ua string, dest *UserAgent) {
	parse(ua, dest)
}

func parse(ua string, dest *UserAgent) {
	ua = normalise(ua)
	switch {
	case len(ua) == 0:
		// dest.OS.Platform = PlatformUnknown
		dest.OS.Type = os.OS_Type_unknown
		dest.Browser.Type = browser.BrowserType_unknown
		dest.Device = device.DeviceType_unknown

	// stop on on first case returning true
	case dest.evalOS(ua):
	case dest.evalBrowserName(ua):
	default:
		dest.evalBrowserVersion(ua)
		dest.evalDevice(ua)
	}
}

// normalise normalises the user supplied agent string so that
// we can more easily parse it.
func normalise(ua string) string {
	if len(ua) <= 1024 {
		var buf [1024]byte
		ascii := copyLower(buf[:len(ua)], ua)
		if !ascii {
			// Fall back for non ascii characters
			return strings.ToLower(ua)
		}
		return string(buf[:len(ua)])
	}
	// Fallback for unusually long strings
	return strings.ToLower(ua)
}

// copyLower copies a lowercase version of s to b. It assumes s contains only single byte characters
// and will panic if b is nil or is not long enough to contain all the bytes from s.
// It returns early with false if any characters were non ascii.
func copyLower(b []byte, s string) bool {
	for j := 0; j < len(s); j++ {
		c := s[j]
		if c > 127 {
			return false
		}

		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}

		b[j] = c
	}
	return true
}
