package formatters

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/behance/go-logrus"
)

// KVFormatter - takes entries and flattens them into a K=V format
type KVFormatter struct{}

// KVEntryString - flattens a logrus.Entry into a K=V formatted string
func KVEntryString(entry *logrus.Entry) string {
	var keys = make([]string, 0, len(entry.Data))

	for k := range entry.Data {
		keys = append(keys, k)
	}

	strentry := "MESSAGE=" + singleQuoteString(entry.Message)
	for _, k := range keys {
		v := entry.Data[k]
		s, ok := v.(string)
		if ok {
			v = s
		}
		strentry = strentry + " " + strings.ToUpper(k) + "=" + singleQuoteString(s)
	}

	return strentry
}

func singleQuoteString(s string) string {
	return "'" + strings.Replace(s, "'", "\\'", -1) + "'"
}

// Format - See logrus.Formatter.Format for docs
func (f KVFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}

	fmt.Fprintf(b, KVEntryString(entry))
	fmt.Fprintln(b)

	return b.Bytes(), nil
}
