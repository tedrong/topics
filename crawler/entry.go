package crawler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
	"github.com/topics/sysexec"
)

type Crawler struct {
	Mutex           sync.Mutex
	URL             string
	SeleniumPath    string
	GeckoDriverPath string
	Port            int
	WebDriver       *selenium.WebDriver
}

var c *Crawler

func StartWebInstance() (*selenium.Service, *Crawler) {
	// Check if there is a instance running, kill it
	if pid := sysexec.FindWebDriverPID(os.Getenv("WEB_INSTANCE_PORT")); pid != nil {
		sysexec.KillWebDriver(pid)
	}

	port, err := strconv.Atoi(os.Getenv("WEB_INSTANCE_PORT"))
	if err != nil {
		log.Fatal(errors.Wrap(err, "WebDriver can't get correct port number"))
	}
	crawler := Crawler{
		URL:             "https://www.google.com",
		SeleniumPath:    os.Getenv("SELENIUM"),
		GeckoDriverPath: os.Getenv("GECKO_DRIVER"),
		Port:            port,
	}
	// Start a Selenium WebDriver server instance (if one is not already running).
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),                   // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(crawler.GeckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		// selenium.Output(os.Stderr),              // Output debug information to STDERR.
	}

	// selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(crawler.SeleniumPath, crawler.Port, opts...)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Selenium WebInstance start"))
		return nil, nil
	}
	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	firefoxCaps := firefox.Capabilities{
		Args: []string{
			"--headless",
		},
	}
	caps.AddFirefox(firefoxCaps)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", crawler.Port))
	if err != nil {
		log.Fatal(errors.Wrap(err, "Connect to WebDriver"))
	}
	crawler.WebDriver = &wd
	c = &crawler
	return service, c
}

func Get() *Crawler {
	return c
}

func (c *Crawler) GOTO() {
	if err := (*c.WebDriver).Get(c.URL); err != nil {
		log.Panic(errors.Wrap(err, fmt.Sprintf("Connect to %s", c.URL)))
	}
}

func getElenentText(element *selenium.WebElement) string {
	value, err := (*element).Text()
	if err != nil {
		return ""
	}
	return value
}
