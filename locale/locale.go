package locale

import (
	"io/ioutil"
	"os"
)

// GetLocale default locale
func GetLocale() string {
	return os.Getenv("LANG")
}

// SetLocale default locale
func SetLocale(code string) error {
	data := []byte("export LANG=" + code)
	return ioutil.WriteFile("/etc/profile.d/i18n.sh", data, 0755)
}

// ListLocale Locale
func ListLocale() []string {
	localeDir, err := ioutil.ReadDir("/usr/share/i18n/locales")
	if err != nil {
		return nil
	}
	locales := make([]string, 0)
	for _, l := range localeDir {
		locales = append(locales, l.Name())
	}
	return locales
}
