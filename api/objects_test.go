package api

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/cheekybits/is"
)

type testData struct {
	Field string `json:"field_foo,omitempty"`
}

func (t *testData) setField(k, v string) error {
	return setField(t, k, v)
}

func TestSetField(t *testing.T) {
	is := is.New(t)

	o := &testData{}
	err := o.setField("field_foo", "123")
	is.NoErr(err)
	is.Equal("123", o.Field)
}

func TestParseBlock(t *testing.T) {
	is := is.New(t)

	dat, err := ioutil.ReadFile("testdata/contact_status.dat")
	is.NoErr(err)

	a := strings.SplitAfterN(string(dat), "}", -1)
	is.Equal(2, len(a))

	lines := strings.Split(a[0], "\n")

	c := &ContactStatus{}
	parseBlock(c, "contactstatus", lines)

	// Sample a handful of fields
	is.Equal(c.ContactName, "jason")
	is.Equal(c.ModifiedAttributes, "0")
	is.Equal(c.ModifiedHostAttributes, "0")
	is.Equal(c.LastHostNotification, "1481756484")
	is.Equal(c.ServiceNotificationPeriod, "24x7")

	is.NotNil(c.CustomVariables)
	is.Equal(len(c.CustomVariables), 1)
	val, ok := c.CustomVariables["SOMECUSTOMVAR"]
	is.OK(ok)
	is.Equal(val, "http://example.com/customvar")
}
