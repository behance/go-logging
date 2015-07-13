package formatters

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/behance/go-logrus"
)

// SumologicFormatter - takes entries and flattens them into a K=V format
// with an additional APP_NAME key
type SumologicFormatter struct{}

// Format - See logrus.Formatter.Format for docs
func (f SumologicFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}

	fmt.Fprintf(
		b,
		"APP_NAME='%s' SEVERITY='%s' ",
		appName(),
		strings.ToUpper(entry.Level.String()),
	)
	// KVEntryString in the kv.go file
	fmt.Fprintf(b, KVEntryString(entry))
	fmt.Fprintln(b)

	return b.Bytes(), nil
}

func appName() string {
	appname := os.Getenv("LOG_APP_NAME")
	if appname == "" {
		return "GolangApp"
	}
	return appname
}
