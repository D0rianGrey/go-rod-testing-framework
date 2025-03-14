package config

import (
    "os"
    "time"
)

// Config содержит настройки для тестового фреймворка
type Config struct {
    BaseURL            string
    Browser            string
    Headless           bool
    Timeout            time.Duration
    ImplicitWaitTime   time.Duration
    ScreenshotsEnabled bool
    ScreenshotsPath    string
}

// NewConfig создает и возвращает новую конфигурацию с значениями по умолчанию и из переменных окружения
func NewConfig() *Config {
    config := &Config{
        BaseURL:            getEnv("BASE_URL", "https://www.saucedemo.com"),
        Browser:            getEnv("BROWSER", "chrome"),
        Headless:           getEnvBool("HEADLESS", true),
        Timeout:            getEnvDuration("TIMEOUT", 30*time.Second),
        ImplicitWaitTime:   getEnvDuration("IMPLICIT_WAIT_TIME", 5*time.Second),
        ScreenshotsEnabled: getEnvBool("SCREENSHOTS_ENABLED", true),
        ScreenshotsPath:    getEnv("SCREENSHOTS_PATH", "./reports/screenshots"),
    }
    return config
}

// Вспомогательные функции для чтения переменных окружения
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
    if value, exists := os.LookupEnv(key); exists {
        return value == "true" || value == "1" || value == "yes"
    }
    return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
    if value, exists := os.LookupEnv(key); exists {
        if duration, err := time.ParseDuration(value); err == nil {
            return duration
        }
    }
    return defaultValue
}