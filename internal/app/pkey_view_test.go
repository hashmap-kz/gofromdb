package app

import (
	"strings"
	"testing"
)

func field(dbName, fieldName, fieldType string) TableToStructFieldInfo {
	return TableToStructFieldInfo{DbFieldName: dbName, FieldName: fieldName, FieldType: fieldType}
}

func TestPkURLPath(t *testing.T) {
	cases := []struct {
		fields []TableToStructFieldInfo
		want   string
	}{
		{nil, ""},
		{[]TableToStructFieldInfo{field("id", "ID", "int64")}, "{id}"},
		{
			[]TableToStructFieldInfo{
				field("user_id", "UserID", "int64"),
				field("role_id", "RoleID", "int32"),
			},
			"{user_id}/{role_id}",
		},
	}
	for _, c := range cases {
		if got := pkURLPath(c.fields); got != c.want {
			t.Errorf("pkURLPath(%v) = %q, want %q", c.fields, got, c.want)
		}
	}
}

func TestPkParams(t *testing.T) {
	cases := []struct {
		fields []TableToStructFieldInfo
		want   string
	}{
		{nil, ""},
		{[]TableToStructFieldInfo{field("id", "ID", "int64")}, "pkID int64"},
		{
			[]TableToStructFieldInfo{
				field("user_id", "UserID", "int64"),
				field("role_id", "RoleID", "int32"),
			},
			"pkUserID int64, pkRoleID int32",
		},
	}
	for _, c := range cases {
		if got := pkParams(c.fields); got != c.want {
			t.Errorf("pkParams = %q, want %q", got, c.want)
		}
	}
}

func TestPkArgs(t *testing.T) {
	cases := []struct {
		fields []TableToStructFieldInfo
		want   string
	}{
		{nil, ""},
		{[]TableToStructFieldInfo{field("id", "ID", "int64")}, "pkID"},
		{
			[]TableToStructFieldInfo{
				field("user_id", "UserID", "int64"),
				field("role_id", "RoleID", "int32"),
			},
			"pkUserID, pkRoleID",
		},
	}
	for _, c := range cases {
		if got := pkArgs(c.fields); got != c.want {
			t.Errorf("pkArgs = %q, want %q", got, c.want)
		}
	}
}

func TestPkWhereClause(t *testing.T) {
	cases := []struct {
		fields []TableToStructFieldInfo
		want   string
	}{
		{nil, ""},
		{[]TableToStructFieldInfo{field("id", "ID", "int64")}, "id = $1"},
		{
			[]TableToStructFieldInfo{
				field("user_id", "UserID", "int64"),
				field("role_id", "RoleID", "int32"),
			},
			"user_id = $1 and role_id = $2",
		},
	}
	for _, c := range cases {
		if got := pkWhereClause(c.fields); got != c.want {
			t.Errorf("pkWhereClause = %q, want %q", got, c.want)
		}
	}
}

func TestPkOrderClause(t *testing.T) {
	cases := []struct {
		fields []TableToStructFieldInfo
		want   string
	}{
		{nil, ""},
		{[]TableToStructFieldInfo{field("id", "ID", "int64")}, "id"},
		{
			[]TableToStructFieldInfo{
				field("user_id", "UserID", "int64"),
				field("role_id", "RoleID", "int32"),
			},
			"user_id, role_id",
		},
	}
	for _, c := range cases {
		if got := pkOrderClause(c.fields); got != c.want {
			t.Errorf("pkOrderClause = %q, want %q", got, c.want)
		}
	}
}

func TestPkOrderSQL(t *testing.T) {
	cases := []struct{ in, want string }{
		{"", ""},
		{"id", "order by id"},
		{"user_id, role_id", "order by user_id, role_id"},
	}
	for _, c := range cases {
		if got := pkOrderSQL(c.in); got != c.want {
			t.Errorf("pkOrderSQL(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPathValueParser(t *testing.T) {
	cases := []struct {
		in   string
		want string
		ok   bool
	}{
		{"int", "httputils.PathValueI32", true},
		{"int32", "httputils.PathValueI32", true},
		{"int16", "httputils.PathValueI16", true},
		{"int64", "httputils.PathValueI64", true},
		{"uint32", "httputils.PathValueU32", true},
		{"string", "httputils.PathValueString", true},
		{"float64", "", false},
		{"bool", "", false},
		{"", "", false},
	}
	for _, c := range cases {
		got, ok := pathValueParser(c.in)
		if ok != c.ok || got != c.want {
			t.Errorf("pathValueParser(%q) = (%q, %v), want (%q, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}

func TestPkPathRead_ContainsFieldName(t *testing.T) {
	fields := []TableToStructFieldInfo{field("user_id", "UserID", "int64")}
	got, err := pkPathRead(fields)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "pkUserID") {
		t.Errorf("pkPathRead output missing pkUserID: %s", got)
	}
	if !strings.Contains(got, `"user_id"`) {
		t.Errorf("pkPathRead output missing db field name: %s", got)
	}
	if !strings.Contains(got, "httputils.PathValueI64") {
		t.Errorf("pkPathRead output missing parser func: %s", got)
	}
}

func TestPkPathRead_Composite(t *testing.T) {
	fields := []TableToStructFieldInfo{
		field("user_id", "UserID", "int64"),
		field("code", "Code", "string"),
	}
	got, err := pkPathRead(fields)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "pkUserID") || !strings.Contains(got, "pkCode") {
		t.Errorf("pkPathRead composite missing expected vars: %s", got)
	}
}

func TestNewPrimaryKeyView(t *testing.T) {
	fields := []TableToStructFieldInfo{field("id", "ID", "int64")}
	pk, err := NewPrimaryKeyView(fields)
	if err != nil {
		t.Fatal(err)
	}

	if pk.URLPath != "{id}" {
		t.Errorf("URLPath = %q, want %q", pk.URLPath, "{id}")
	}
	if pk.WhereClause != "id = $1" {
		t.Errorf("WhereClause = %q, want %q", pk.WhereClause, "id = $1")
	}
	if pk.Params != "pkID int64" {
		t.Errorf("Params = %q, want %q", pk.Params, "pkID int64")
	}
	if pk.Args != "pkID" {
		t.Errorf("Args = %q, want %q", pk.Args, "pkID")
	}
	if pk.OrderClause != "id" {
		t.Errorf("OrderClause = %q, want %q", pk.OrderClause, "id")
	}
	if pk.OrderSQL != "order by id" {
		t.Errorf("OrderSQL = %q, want %q", pk.OrderSQL, "order by id")
	}
}

func TestNewPrimaryKeyView_Empty(t *testing.T) {
	pk, err := NewPrimaryKeyView(nil)
	if err != nil {
		t.Fatal(err)
	}
	if pk.URLPath != "" || pk.WhereClause != "" || pk.Params != "" || pk.Args != "" {
		t.Errorf("NewPrimaryKeyView(nil) should produce empty strings: %+v", pk)
	}
	if pk.OrderSQL != "" {
		t.Errorf("OrderSQL should be empty for no fields, got %q", pk.OrderSQL)
	}
}
