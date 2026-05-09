package app

import (
	"testing"
)

func TestCapitalizeFirstLetter(t *testing.T) {
	cases := []struct{ in, want string }{
		{"", ""},
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"h", "H"},
		{"id", "Id"},
	}
	for _, c := range cases {
		if got := capitalizeFirstLetter(c.in); got != c.want {
			t.Errorf("capitalizeFirstLetter(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestLowerFirstLetter(t *testing.T) {
	cases := []struct{ in, want string }{
		{"", ""},
		{"Hello", "hello"},
		{"hello", "hello"},
		{"H", "h"},
		{"ID", "iD"},
	}
	for _, c := range cases {
		if got := LowerFirstLetter(c.in); got != c.want {
			t.Errorf("LowerFirstLetter(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMakeName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"id", "ID"},
		{"guid", "GUID"},
		{"uuid", "UUID"},
		{"name", "Name"},
		{"user_id", "UserID"},
		{"user_guid", "UserGUID"},
		{"user_uuid", "UserUUID"},
		{"first_name", "FirstName"},
		{"created_at", "CreatedAt"},
		{"some_field_name", "SomeFieldName"},
	}
	for _, c := range cases {
		if got := makeName(c.in); got != c.want {
			t.Errorf("makeName(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGetSchemaTable(t *testing.T) {
	schema, table := getSchemaTable("public.users")
	if schema != "public" {
		t.Errorf("schema = %q, want %q", schema, "public")
	}
	if table != "users" {
		t.Errorf("table = %q, want %q", table, "users")
	}
}

func TestAddPadding(t *testing.T) {
	cases := []struct{ in, want string }{
		{"line1\nline2", "\tline1\n\tline2"},
		{"single", "\tsingle"},
		{"", "\t"},
	}
	for _, c := range cases {
		if got := addPadding(c.in); got != c.want {
			t.Errorf("addPadding(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAddPadding2(t *testing.T) {
	cases := []struct{ in, want string }{
		{"line1\nline2", "\t\tline1\n\t\tline2"},
		{"single", "\t\tsingle"},
	}
	for _, c := range cases {
		if got := addPadding2(c.in); got != c.want {
			t.Errorf("addPadding2(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMakeDnsPathPluralFromDbTable(t *testing.T) {
	cases := []struct{ in, want string }{
		{"users", "users"},
		{"product", "product"},
		{"user_profiles", "user-profiles"},
		{"order_line_items", "order-line-items"},
	}
	for _, c := range cases {
		if got := makeDnsPathPluralFromDbTable(c.in); got != c.want {
			t.Errorf("makeDnsPathPluralFromDbTable(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFormatComment(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "empty",
			in:   "",
			want: "",
		},
		{
			name: "whitespace only",
			in:   "   ",
			want: "",
		},
		{
			name: "single short line",
			in:   "user identifier",
			want: "// user identifier",
		},
		{
			name: "trims surrounding whitespace",
			in:   "  trimmed  ",
			want: "// trimmed",
		},
		{
			name: "existing newline becomes separate comment line",
			in:   "first line\nsecond line",
			want: "// first line\n// second line",
		},
		{
			name: "blank line between paragraphs becomes empty comment",
			in:   "para one\n\npara two",
			want: "// para one\n//\n// para two",
		},
		{
			name: "long line is word-wrapped at 80 chars",
			in:   "This is a rather long description that will definitely exceed the eighty character limit for a single comment line",
			want: "// This is a rather long description that will definitely exceed the eighty\n// character limit for a single comment line",
		},
		{
			name: "single word longer than limit is kept on its own line",
			in:   "averylongwordthatexceedsthemaximumwidthlimitofeightycharactersallbyitself_endshere",
			want: "// averylongwordthatexceedsthemaximumwidthlimitofeightycharactersallbyitself_endshere",
		},
		{
			name: "multiline with long paragraph wraps at 80 chars inclusive",
			in:   "Short intro.\nThis second paragraph is much longer and should be wrapped because it exceeds the allowed width of eighty characters.",
			want: "// Short intro.\n// This second paragraph is much longer and should be wrapped because it exceeds\n// the allowed width of eighty characters.",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := formatComment(c.in); got != c.want {
				t.Errorf("formatComment(%q)\ngot:  %q\nwant: %q", c.in, got, c.want)
			}
		})
	}
}
