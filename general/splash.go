package general

import "github.com/fatih/color"

func SplashScreen(includeDesc bool) string {
	// Color the splash screen
	c := color.New(color.FgBlue)
	c2 := color.New(color.FgCyan)

	splash := c.Sprint(`
  _     _         _          _ _ 
 | |__ | |__   __| |     ___| (_)
 | '_ \| '_ \ / _' |___ / __| | |
 | | | | |_) | (_| |___| (__| | |
 |_| |_|_.__/ \__,_|    \___|_|_|
 `)

	if includeDesc {
		splash += c2.Sprint(`
 Welcome to hbd-cli! This is a CLI tool to manage 
 birthday reminders using an HBD backend.

 Don't have an HBD backend? you can self-host 
 your own instance: https://github.com/dreth/hbd

 or use our instance: https://hbd.lotiguere.com
 
 If you encounter any issues or have any suggestions, 
 feel free to open an issue: https://github.com/dreth/hbd-cli`)
	}

	return splash
}
