package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

func setScreenBRightnes(val []byte) error {
	err = ioutil.WriteFile("/sys/class/backlight/intel_backlight/brightness", val, 0644)
	valInt, err := strconv.Atoi(string(val))
	if err != nil {
		return err
	}
	log.Println(valInt)
	if valInt > 1000 {
		colorswitcher("")
		codeOSSswitcher("Solarized Light")
	} else {
		colorswitcher("Dark")
		codeOSSswitcher("Material Theme Palenight High Contrast")
	}
	return err
}

func colorswitcher(theme string) error {

	command := "plasma-theme"
	args := []string{"--colors", "/home/oli/.local/share/color-schemes/WhiteSur" + theme + ".colors"}
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))

	return nil
}

func codeOSSswitcher(theme string) error {

	template := `{
		"workbench.iconTheme": "material-icon-theme",
		"files.autoSave": "afterDelay",
		"workbench.colorCustomizations": {},
		"workbench.colorTheme": "` + theme + `",
		"editor.fontSize": 16,
		"editor.fontFamily": "'Fira Code Retina', 'Droid Sans Mono', 'monospace', monospace, 'Droid Sans Fallback'",
		"editor.lineHeight": 30,
		"editor.fontLigatures": true
		
	}`
	err := ioutil.WriteFile("/home/oli/.config/Code - Insiders/User/settings.json", []byte(template), 0755)
	if err != nil {
		log.Println(err)
	}

	return nil
}
