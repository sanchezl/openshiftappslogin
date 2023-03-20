package pkg

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
)

func RetrieveBearerToken(url, username, password string, log func(v ...any)) string {
	browser := rod.New().DefaultDevice(ChromeOnLinux).MustConnect()
	defer browser.MustClose()

	log("Requesting oauth token.")
	page := browser.MustPage(url)
	page.MustWaitLoad()

	log("Selecting RedHat_Internal_SSO option.")
	var link *rod.Element
	for _, a := range page.MustElements("a") {
		if a.MustText() == "RedHat_Internal_SSO" {
			link = a
			break
		}
	}
	if link == nil {
		panic("RedHat_Internal_SSO option not found")
	}
	link.MustClick()
	page.MustWaitLoad()

	log("Logging into Red Hat Internal.")
	page.MustElement("input#username").MustInput(username)

	page.MustElement("input#password").MustInput(password)
	page.MustElement("input#submit").MustClick()
	page.MustWaitLoad()

	log("Requesting token display.")
	page.MustElement("button").MustClick()
	page.MustWaitLoad()

	log("Extracting token.")
	return page.MustElement("code").MustText()
}

var ChromeOnLinux = devices.Device{
	Title:          "Latest Chrome on Linux",
	Capabilities:   []string{},
	UserAgent:      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
	AcceptLanguage: "en",
	Screen:         devices.Screen{DevicePixelRatio: 1, Horizontal: devices.ScreenSize{Width: 1280, Height: 800}, Vertical: devices.ScreenSize{Width: 800, Height: 1280}},
}
