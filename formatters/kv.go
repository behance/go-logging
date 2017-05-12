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

	// Escape % if it appears in value, so Sprintf doesn't try to expand it
	entry.Message = strings.Replace(entry.Message, "%", "%%", -1)

	strentry := fmt.Sprintf("MESSAGE='%s'", entry.Message)
	for _, k := range keys {
		v := entry.Data[k]
		// Escape % if it appears in value, so Sprintf doesn't try to expand it
		s, ok := v.(string)
		if ok {
			s = strings.Replace(s, "%", "%%", -1)
			v = s
		}
		strentry = fmt.Sprintf("%s %s='%+v'", strentry, strings.ToUpper(k), v)
	}

	return strentry
}

// Format - See logrus.Formatter.Format for docs
func (f KVFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}

	fmt.Fprintf(b, KVEntryString(entry))
	fmt.Fprintln(b)

	return b.Bytes(), nil
}
