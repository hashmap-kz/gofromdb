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
