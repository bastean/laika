package scraper

import (
	"errors"

	"github.com/bastean/laika/pkg/context/shared/domain/model"
	"github.com/playwright-community/playwright-go"
)

type Playwright struct {
	Pw         *playwright.Playwright
	Browser    playwright.Browser
	BrowserCtx playwright.BrowserContext
	Page       playwright.Page
}

type PlaywrightOptions struct {
	Headless bool
}

func (pw *Playwright) GetContent(source string) string {
	response, err := pw.Page.Goto(source)

	if err != nil {
		return ""
	}

	html, err := response.Body()

	if err != nil {
		return ""
	}

	return string(html)
}

func (pw *Playwright) GetLinks(source string) []string {
	if _, err := pw.Page.Goto(source); err != nil {
		return []string{}
	}

	locator := pw.Page.Locator("[href]")

	hrefs, err := locator.All()

	if err != nil {
		return []string{}
	}

	links := []string{}

	for _, href := range hrefs {
		link, err := href.GetAttribute("href")

		if err != nil {
			continue
		}

		links = append(links, link)
	}

	return links
}

func NewPlaywright(option *PlaywrightOptions) (model.Scraper, error) {
	var err error

	err = playwright.Install(&playwright.RunOptions{Browsers: []string{"chromium"}})

	if err != nil {
		return nil, errors.New("could not install browser: " + err.Error())
	}

	pw, err := playwright.Run()

	if err != nil {
		return nil, errors.New("could not start playwright: " + err.Error())
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &option.Headless})

	if err != nil {
		return nil, errors.New("could not launch browser: " + err.Error())
	}

	browserCtx, err := browser.NewContext()

	if err != nil {
		return nil, errors.New("could not create context: " + err.Error())
	}

	page, err := browserCtx.NewPage()

	if err != nil {
		return nil, errors.New("could not create page: " + err.Error())
	}

	return &Playwright{
		Pw:         pw,
		Browser:    browser,
		BrowserCtx: browserCtx,
		Page:       page,
	}, nil
}
