package constants

import (
	"errors"
	"time"

	"github.com/pterm/pterm"
)

var ErrorTemplate = "{{ .Error.Error }}\n"

const Version = "v0.0.1"

const UserAgent = "Jeosgram-CLI/" + Version

const ClientID = "jeosgram-cli"
const ClientSecret = "jeosgram-cli"

const MinApiDelay = 400 * time.Millisecond

const ApiURL = "https://api.jeosgram.io"

var (
	PInfo    = pterm.FgYellow.Sprint("!")
	PError   = pterm.FgRed.Sprint(">>")
	PSuccess = pterm.FgGreen.Sprint(">")
)

const jeosgramArt2 = `
       ██╗███████╗ ██████╗ ███████╗ ██████╗ ██████╗  █████╗ ███╗   ███╗
       ██║██╔════╝██╔═══██╗██╔════╝██╔════╝ ██╔══██╗██╔══██╗████╗ ████║
       ██║█████╗  ██║   ██║███████╗██║  ███╗██████╔╝███████║██╔████╔██║
  ██   ██║██╔══╝  ██║   ██║╚════██║██║   ██║██╔══██╗██╔══██║██║╚██╔╝██║
  ╚█████╔╝███████╗╚██████╔╝███████║╚██████╔╝██║  ██║██║  ██║██║ ╚═╝ ██║
   ╚════╝ ╚══════╝ ╚═════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝
`

const JeosgramArt = jeosgramArt2 + "\nJeosgram CLI client " + Version

const (
	msgRequiredMFA  = "Multi-factor authentication required"
	msgInvalidToken = "The access token provided is invalid"
)

var (
	ErrRequiredMFA  = errors.New(msgRequiredMFA)
	ErrInvalidToken = errors.New(msgInvalidToken)
)
