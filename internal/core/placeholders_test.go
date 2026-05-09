package core

import (
	"strings"
	"testing"
)

func TestCreatePlaceholders(t *testing.T) {
	cases := []struct {
		cnt  int
		want []string
	}{
		{0, []string{}},
		{1, []string{"$1"}},
		{3, []string{"$1", "$2", "$3"}},
		{5, []string{"$1", "$2", "$3", "$4", "$5"}},
	}
	for _, c := range cases {
		got := CreatePlaceholders(c.cnt)
		if len(got) != len(c.want) {
			t.Errorf("CreatePlaceholders(%d) len = %d, want %d", c.cnt, len(got), len(c.want))
			continue
		}
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("CreatePlaceholders(%d)[%d] = %q, want %q", c.cnt, i, got[i], c.want[i])
			}
		}
	}
}

func TestMaxFieldNameLen(t *testing.T) {
	cases := []struct {
		fields []TableToStructFieldInfo
		want   int
	}{
		{nil, 0},
		{[]TableToStructFieldInfo{{DbFieldName: "id"}}, 2},
		{[]TableToStructFieldInfo{
			{DbFieldName: "id"},
			{DbFieldName: "username"},
			{DbFieldName: "email"},
		}, 8}, // "username"
	}
	for _, c := range cases {
		if got := maxFieldNameLen(c.fields); got != c.want {
			t.Errorf("maxFieldNameLen = %d, want %d", got, c.want)
		}
	}
}

func TestGenUpdateSets_PlaceholderIndex(t *testing.T) {
	// 1 pkey: first update field gets $2
	fields := []TableToStructFieldInfo{{DbFieldName: "name", DbNullifRhs: "''::text"}}
	got := GenUpdateSets(fields, 1)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if !strings.Contains(got[0], "$2") {
		t.Errorf("expected $2 in %q", got[0])
	}
}

func TestGenUpdateSets_TwoPkeys(t *testing.T) {
	// 2 pkeys: first update field gets $3
	fields := []TableToStructFieldInfo{{DbFieldName: "x", DbNullifRhs: "0"}}
	got := GenUpdateSets(fields, 2)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	want := "x = coalesce(nullif($3, 0), x)"
	if got[0] != want {
		t.Errorf("got %q, want %q", got[0], want)
	}
}

func TestGenUpdateSets_MultipleFields(t *testing.T) {
	fields := []TableToStructFieldInfo{
		{DbFieldName: "name", DbNullifRhs: "''::text"},
		{DbFieldName: "age", DbNullifRhs: "0::int4"},
	}
	got := GenUpdateSets(fields, 1)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	// sequential placeholders: $2, $3
	if !strings.Contains(got[0], "$2") {
		t.Errorf("first field should use $2, got %q", got[0])
	}
	if !strings.Contains(got[1], "$3") {
		t.Errorf("second field should use $3, got %q", got[1])
	}
}

func TestGenUpdateSets_CoalesceNullifPattern(t *testing.T) {
	fields := []TableToStructFieldInfo{{DbFieldName: "email", DbNullifRhs: "''::text"}}
	got := GenUpdateSets(fields, 1)
	if !strings.HasPrefix(got[0], "email") {
		t.Errorf("should start with field name: %q", got[0])
	}
	if !strings.Contains(got[0], "coalesce(nullif(") {
		t.Errorf("should contain coalesce(nullif(: %q", got[0])
	}
	if !strings.Contains(got[0], "''::text") {
		t.Errorf("should contain nullif rhs: %q", got[0])
	}
	if !strings.HasSuffix(strings.TrimRight(got[0], " "), ", email)") {
		t.Errorf("should end with field name as fallback: %q", got[0])
	}
}

func TestGenUpdateSets_Empty(t *testing.T) {
	got := GenUpdateSets(nil, 1)
	if len(got) != 0 {
		t.Errorf("GenUpdateSets(nil) = %v, want empty", got)
	}
}
