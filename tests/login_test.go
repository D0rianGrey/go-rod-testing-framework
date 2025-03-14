package tests

import (
	"testing"

	"github.com/D0rianGrey/go-rod-testing-framework/config"
	"github.com/D0rianGrey/go-rod-testing-framework/pkg/browser"
)

func TestLoginPage(t *testing.T) {
	// Создаем конфигурацию
	cfg := config.NewConfig()

	// Создаем браузер
	browser := browser.New(cfg)
	defer browser.Close()

	// Открываем страницу
	page := browser.OpenPage(cfg.BaseURL)

	// Проверяем, что страница загрузилась успешно
	if page.MustInfo().URL != cfg.BaseURL {
		t.Errorf("Ожидался URL %s, получен %s", cfg.BaseURL, page.MustInfo().URL)
	}

	t.Log("Страница логина успешно загружена")
}
