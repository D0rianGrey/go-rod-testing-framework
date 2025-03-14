package browser

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/D0rianGrey/go-rod-testing-framework/config"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
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

	// Создаем директорию, если она не существует
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Делаем скриншот и сохраняем в файл
	data, err := b.Page.Screenshot(false, nil)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
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

	// Устанавливаем контекст и ждем навигации
	b.Page.Context(ctx)
	b.Page.MustWaitNavigation()
}
