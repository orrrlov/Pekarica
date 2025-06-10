package main

import (
	"fmt"
	"html/template"
	"os/exec"
	"runtime"
)

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return err
}

func url(domain, port string) string {
	return fmt.Sprintf("http://%s:%s", domain, port)
}

func parseTemplate(page string) *template.Template {
	path := "./frontend"
	return template.Must(template.ParseFiles(fmt.Sprintf("%s/layout.html", path), fmt.Sprintf("%s/%s.html", path, page)))
}
