package handles

import (
	"os"

	"github.com/Vico1993/Otto-client/otto"
)

var OttoClient *otto.Client

// Initialise
func Init() {
	OttoClient = otto.NewClient(
		nil,
		os.Getenv("OTTO_API_URL"),
	)
}
