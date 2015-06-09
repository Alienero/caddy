package setup

import (
	"testing"

	"github.com/mholt/caddy/middleware/gzip"
)

func TestGzip(t *testing.T) {
	c := newTestController(`gzip`)

	mid, err := Gzip(c)
	if err != nil {
		t.Errorf("Expected no errors, but got: %v", err)
	}
	if mid == nil {
		t.Fatal("Expected middleware, was nil instead")
	}

	handler := mid(emptyNext)
	myHandler, ok := handler.(*gzip.Gzip)
	if !ok {
		t.Fatalf("Expected handler to be type *gzip.Gzip, got: %#v", handler)
	}

	if !sameNext(myHandler.Next, emptyNext) {
		t.Error("'Next' field of handler was not set properly")
	}

	tests := []struct {
		input     string
		shouldErr bool
	}{
		{`gzip {`, true},
		{`gzip {}`, true},
		{`gzip a b`, true},
		{`gzip a {`, true},
		{`gzip { not f } `, true},
		{`gzip { not } `, true},
		{`gzip { not /file
		 ext .html
		 level 1
		} `, false},
		{`gzip { level 9 } `, false},
		{`gzip { ext } `, true},
		{`gzip { ext /f
		} `, true},
		{`gzip { not /file
		 ext .html
		 level 1
		}
		gzip`, false},
		{`gzip { not /file
		 ext .html
		 level 1
		}
		gzip { not /file1
		 ext .htm
		 level 3
		}
		`, false},
	}
	for i, test := range tests {
		c := newTestController(test.input)
		_, err := gzipParse(c)
		if test.shouldErr && err == nil {
			t.Errorf("Text %v: Expected error but found nil", i)
		} else if !test.shouldErr && err != nil {
			t.Errorf("Text %v: Expected no error but found error: ", i, err)
		}
	}
}
