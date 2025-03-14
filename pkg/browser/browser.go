package browser

import (
	"context"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/yourusername/go-rod-testing-framework/config"
)

// Browser представляет собой обертку над браузером Rod
type Browser struct {
	Browser *rod.Browser
	Page    *rod.Page
	Config  *config.Config
}

// New создает новый экземпляр браузера
func New(cfg *config.Config) *Browser {
	url := launcher.New().
		Headless(cfg.Headless).
		MustLaunch()

	browser := rod.New().
		ControlURL(url).
		Timeout(cfg.Timeout).
		MustConnect()

	return &Browser{
		Browser: browser,
		Config:  cfg,
	}
}

// OpenPage открывает новую страницу с заданным URL
func (b *Browser) OpenPage(url string) *rod.Page {
	page := b.Browser.MustPage(url)
	b.Page = page
	return page
}

// Close закрывает браузер
func (b *Browser) Close() {
	if b.Browser != nil {
		b.Browser.MustClose()
	}
}

// TakeScreenshot делает скриншот страницы и сохраняет его в указанный путь
func (b *Browser) TakeScreenshot(path string) error {
	if b.Page == nil {
		return nil
	}
	return b.Page.Screenshot(path)
}

// WaitForLoad ожидает загрузки страницы
func (b *Browser) WaitForLoad() {
	if b.Page != nil {
		b.Page.MustWaitLoad()
	}
}

// WaitForNavigation ожидает навигации на странице
func (b *Browser) WaitForNavigation(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	b.Page.MustWaitNavigation(ctx)
}
