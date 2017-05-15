package formatters

import (
	"testing"

	"github.com/behance/go-logrus"
)

func TestKVEntryStringEmpty(t *testing.T) {
	entry := &logrus.Entry{}
	s := KVEntryString(entry)
	expected := "MESSAGE=''"
	if s != expected {
		t.Errorf("s was %+v; not equal to expected value: %+v", s, expected)
	}
}

func TestKVEntryStringSimple(t *testing.T) {
	entry := &logrus.Entry{Message: "Hello"}
	s := KVEntryString(entry)
	expected := "MESSAGE='Hello'"
	if s != expected {
		t.Errorf("s was %+v; not equal to expected value: %+v", s, expected)
	}
}

func TestKVEntryStringData(t *testing.T) {
	entry := &logrus.Entry{
		Message: "Hello",
		Data:    map[string]interface{}{"foo": "bar"},
	}
	s := KVEntryString(entry)
	expected := "MESSAGE='Hello' FOO='bar'"
	if s != expected {
		t.Errorf("s was %+v; not equal to expected value: %+v", s, expected)
	}
}

func TestKVEntryStringSingleQuoteInData(t *testing.T) {
	entry := &logrus.Entry{
		Message: "Hello",
		Data: map[string]interface{}{
			"status": "It's all good",
		},
	}
	s := KVEntryString(entry)
	expected := "MESSAGE='Hello' STATUS='It\\'s all good'"
	if s != expected {
		t.Errorf("s was %+v; not equal to expected value: %+v", s, expected)
	}
}

func TestKVEntryStringSingleQuoteInMessage(t *testing.T) {
	entry := &logrus.Entry{Message: "Hello, it's me"}
	s := KVEntryString(entry)
	expected := "MESSAGE='Hello, it\\'s me'"
	if s != expected {
		t.Errorf("s was %+v; not equal to expected value: %+v", s, expected)
	}
}

func TestKVEntryStringPercentInData(t *testing.T) {
	entry := &logrus.Entry{
		Message: "Hello",
		Data: map[string]interface{}{
			"path":   "%2Fpath%2Fto%2Ffile",
			"effort": "110%",
		},
	}
	s := KVEntryString(entry)
	expected := "MESSAGE='Hello' PATH='%2Fpath%2Fto%2Ffile' EFFORT='110%'"
	if s != expected {
		t.Errorf("s was %+v; not equal to expected value: %+v", s, expected)
	}
}

func TestKVEntryStringPercentInMessage(t *testing.T) {
	entry := &logrus.Entry{Message: "Test coverage is 100%!"}
	s := KVEntryString(entry)
	expected := "MESSAGE='Test coverage is 100%!'"
	if s != expected {
		t.Errorf("s was %+v; not equal to expected value: %+v", s, expected)
	}
}
