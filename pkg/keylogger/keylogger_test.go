package keylogger

import "testing"

func TestKeyLogger_Listen(t *testing.T) {
	tcs := map[string]string{
		"ok": "",
	}

	for n, tc := range tcs {
		tc := tc
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			_ = tc

			l, err := New()
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			l.Listen()
		})
	}
}
