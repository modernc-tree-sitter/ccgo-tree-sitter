package grammar

import (
	"testing"

	"modernc.org/libc"
)

func TestExecuteMatchesNilGuards(t *testing.T) {
	tls := libc.NewTLS()
	t.Cleanup(tls.Close)

	// Non-zero ptr with real TLS: NewCursor can allocate, but root checks
	// must prevent ts_query_cursor_exec on a fake query pointer.
	fakeQ := &Query{ptr: 1, tls: tls}

	cases := []struct {
		name string
		q    *Query
		root *Node
	}{
		{"nil query", nil, nil},
		{"zero query", &Query{}, nil},
		{"zero query zero node", &Query{}, &Node{}},
		{"unusable query nil root", fakeQ, nil},
		{"unusable query zero node", fakeQ, &Node{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.q.ExecuteMatches(tc.root, nil)
			if len(got) != 0 {
				t.Fatalf("len(matches)=%d want 0", len(got))
			}
		})
	}
}
