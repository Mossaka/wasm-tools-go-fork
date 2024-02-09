package wit

import (
	"reflect"
	"testing"

	"github.com/coreos/go-semver/semver"
)

func TestIdent(t *testing.T) {
	tests := []struct {
		s       string
		want    Ident
		wantErr bool
	}{
		{"wasi:io", Ident{Namespace: "wasi", Package: "io"}, false},
		{"wasi:io@0.2.0", Ident{Namespace: "wasi", Package: "io", Version: semver.New("0.2.0")}, false},
		{"wasi:io/streams", Ident{Namespace: "wasi", Package: "io", Extension: "streams"}, false},
		{"wasi:io/streams@0.2.0", Ident{Namespace: "wasi", Package: "io", Extension: "streams", Version: semver.New("0.2.0")}, false},

		// Errors
		{"", Ident{}, true},
		{":", Ident{}, true},
		{":/", Ident{}, true},
		{":/@", Ident{}, true},
		{"wasi", Ident{Namespace: "wasi"}, true},
		{"wasi:", Ident{Namespace: "wasi"}, true},
		{"wasi:/", Ident{}, true},
		{"wasi:clocks@", Ident{}, true},
		{"wasi:clocks/wall-clock@", Ident{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, err := ParseIdent(tt.s)
			if tt.wantErr && err == nil {
				t.Errorf("ParseIdent(%q): expected error, got nil error", tt.s)
			} else if !tt.wantErr && err != nil {
				t.Errorf("ParseIdent(%q): expected no error, got error: %v", tt.s, err)
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseIdent(%q): %v, expected %v", tt.s, got, tt.want)
			}
		})
	}
}
