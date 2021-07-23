package crawler

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
)

type CrawlerEntry struct {
	URL             string
	SeleniumPath    string
	GeckoDriverPath string
	Port            int
	Crawler         *selenium.WebDriver
}

func (c *CrawlerEntry) StartWebInstance() (*selenium.Service, error) {
	// Start a Selenium WebDriver server instance (if one is not already running).
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(c.GeckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		// selenium.Output(os.Stderr),              // Output debug information to STDERR.
	}

	// selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(c.SeleniumPath, c.Port, opts...)
	if err != nil {
		log.Fatal(errors.Wrap(err, "WebDriver instance start"))
		return nil, err
	}
	return service, nil
}

func (c *CrawlerEntry) Init() (*selenium.WebDriver, error) {
	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", c.Port))
	if err != nil {
		log.Fatal(errors.Wrap(err, "Connect to WebDriver"))
		return nil, err
	}

	if err := wd.Get(c.URL); err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Connect to %s", c.URL)))
		return nil, err
	}
	return &wd, nil
}
