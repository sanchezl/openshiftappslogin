package pkg

import (
	"github.com/go-rod/rod"
)

func RetrieveBearerToken(url, username, password string, log func(v ...any)) string {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	log("Requesting oauth token.")
	page := browser.MustPage(url)
	page.MustWaitLoad()

	log("Selecting RedHat_Internal_SSO option.")
	page.MustElement("a").MustClick()
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
