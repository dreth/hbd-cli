# hbd-cli

This is a CLI for the HBD app: github.com/dreth/hbd

HBD is a birthday reminder app with a backend and a frontend. This CLI is a way to interact with the backend, so that users that prefer the terminal can use the app without having to open the browser.

Using the CLI you're able to do all the same things you can do on the frontend.

## Options

The CLI allows you to do the following:

- Authenticate or register an account and perform certain account actions
  - Login
  - Logout (Which just removes the file created with the token)

## Help output

```txt
  _     _         _          _ _ 
 | |__ | |__   __| |     ___| (_)
 | '_ \| '_ \ / _' |___ / __| | |
 | | | | |_) | (_| |___| (__| | |
 |_| |_|_.__/ \__,_|    \___|_|_|
 
 Welcome to hbd-cli! This is a CLI tool to manage 
 birthday reminders using an HBD backend.

 Don't have an HBD backend? you can self-host 
 your own instance: https://github.com/dreth/hbd

 or use our instance: https://hbd.lotiguere.com
 
 If you encounter any issues or have any suggestions, 
 feel free to open an issue: https://github.com/dreth/hbd-cli

Usage:
  hbd [command]

Available Commands:
  auth        Authentication related commands (login, register, etc.)
  birthdays   Birthday related commands (add, list, delete, modify)
  completion  Generate the autocompletion script for the specified shell
  health      Healthcheck the HBD service
  help        Help about any command

Flags:
  -h, --help      help for hbd
  -v, --version   version for hbd

Use "hbd [command] --help" for more information about a command.
```
