package app

import (
	"testing"
)

func TestIsInternalFieldToSkip(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{"created_at", true},
		{"updated_at", true},
		{"guid", true},
		{"id", false},
		{"name", false},
		{"email", false},
		{"user_id", false},
	}
	for _, c := range cases {
		if got := isInternalFieldToSkip(c.name); got != c.want {
			t.Errorf("isInternalFieldToSkip(%q) = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestGetStructFields_InsertableOnly(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id", DbIsInsertable: false, DbIsPrimaryKey: true},
			{DbFieldName: "name", DbIsInsertable: true},
			{DbFieldName: "score", DbIsInsertable: false},
		},
	}
	got := s.GetStructFields(Filters{WithInsertableOnly: true, WithInternals: true})
	if len(got) != 1 || got[0].DbFieldName != "name" {
		t.Errorf("GetStructFields(insertable only) = %v, want [name]", got)
	}
}

func TestGetStructFields_ExcludesInternals(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id"},
			{DbFieldName: "name"},
			{DbFieldName: "created_at"},
			{DbFieldName: "updated_at"},
			{DbFieldName: "guid"},
		},
	}
	got := s.GetStructFields(Filters{WithInternals: false})
	for _, f := range got {
		if isInternalFieldToSkip(f.DbFieldName) {
			t.Errorf("result should not contain internal field %q", f.DbFieldName)
		}
	}
	if len(got) != 2 {
		t.Errorf("len = %d, want 2 (id, name)", len(got))
	}
}

func TestGetStructFields_IncludesInternals(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id"},
			{DbFieldName: "created_at"},
			{DbFieldName: "guid"},
		},
	}
	got := s.GetStructFields(Filters{WithInternals: true})
	if len(got) != 3 {
		t.Errorf("len = %d, want 3", len(got))
	}
}

func TestGetStructFields_WithoutPrimaryKeys(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id", DbIsPrimaryKey: true, DbIsInsertable: true},
			{DbFieldName: "name", DbIsPrimaryKey: false, DbIsInsertable: true},
		},
	}
	got := s.GetStructFields(Filters{WithoutPrimaryKeys: true})
	for _, f := range got {
		if f.DbIsPrimaryKey {
			t.Errorf("result should not contain primary key %q", f.DbFieldName)
		}
	}
	if len(got) != 1 || got[0].DbFieldName != "name" {
		t.Errorf("got %v, want [name]", got)
	}
}

func TestInsertFields(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id", DbIsInsertable: false, DbIsPrimaryKey: true},
			{DbFieldName: "name", DbIsInsertable: true},
			{DbFieldName: "created_at", DbIsInsertable: false},
			{DbFieldName: "score", DbIsInsertable: true},
		},
	}
	got := s.InsertFields()
	if len(got) != 2 {
		t.Fatalf("InsertFields len = %d, want 2", len(got))
	}
	names := map[string]bool{}
	for _, f := range got {
		names[f.DbFieldName] = true
	}
	if !names["name"] || !names["score"] {
		t.Errorf("InsertFields missing expected fields: %v", got)
	}
}

func TestInsertFields_ExcludesInternals(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "name", DbIsInsertable: true},
			{DbFieldName: "created_at", DbIsInsertable: true},
			{DbFieldName: "guid", DbIsInsertable: true},
		},
	}
	got := s.InsertFields()
	for _, f := range got {
		if isInternalFieldToSkip(f.DbFieldName) {
			t.Errorf("InsertFields should not contain internal field %q", f.DbFieldName)
		}
	}
}

func TestUpdateFields(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id", DbIsInsertable: true, DbIsPrimaryKey: true},
			{DbFieldName: "name", DbIsInsertable: true, DbIsPrimaryKey: false},
			{DbFieldName: "created_at", DbIsInsertable: false},
		},
	}
	got := s.UpdateFields()
	if len(got) != 1 || got[0].DbFieldName != "name" {
		t.Errorf("UpdateFields = %v, want [name]", got)
	}
}

func TestUpdateFields_ExcludesPrimaryKeysAndNonInsertable(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id", DbIsInsertable: true, DbIsPrimaryKey: true},
			{DbFieldName: "serial_col", DbIsInsertable: false, DbIsPrimaryKey: false},
			{DbFieldName: "name", DbIsInsertable: true, DbIsPrimaryKey: false},
		},
	}
	got := s.UpdateFields()
	for _, f := range got {
		if f.DbIsPrimaryKey {
			t.Errorf("UpdateFields should not include primary key %q", f.DbFieldName)
		}
		if !f.DbIsInsertable {
			t.Errorf("UpdateFields should not include non-insertable field %q", f.DbFieldName)
		}
	}
}

func TestFullFields_IncludesInternals(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id"},
			{DbFieldName: "name"},
			{DbFieldName: "created_at"},
			{DbFieldName: "updated_at"},
			{DbFieldName: "guid"},
		},
	}
	got := s.FullFields()
	if len(got) != 5 {
		t.Errorf("FullFields len = %d, want 5 (includes internals)", len(got))
	}
}

func TestHandlePkeys_Single(t *testing.T) {
	fields := []TableToStructFieldInfo{
		{DbFieldName: "id", FieldName: "ID"},
		{DbFieldName: "name", FieldName: "Name"},
	}
	got := handlePkeys(fields, []string{"id"})
	if len(got) != 1 || got[0].DbFieldName != "id" {
		t.Errorf("handlePkeys single = %v, want [id]", got)
	}
}

func TestHandlePkeys_Composite(t *testing.T) {
	fields := []TableToStructFieldInfo{
		{DbFieldName: "user_id"},
		{DbFieldName: "role_id"},
		{DbFieldName: "name"},
	}
	got := handlePkeys(fields, []string{"user_id", "role_id"})
	if len(got) != 2 {
		t.Fatalf("handlePkeys composite len = %d, want 2", len(got))
	}
	if got[0].DbFieldName != "user_id" || got[1].DbFieldName != "role_id" {
		t.Errorf("handlePkeys wrong order: %v", got)
	}
}

func TestHandlePkeys_PreservesKeyOrder(t *testing.T) {
	// fields listed in reverse order of keys
	fields := []TableToStructFieldInfo{
		{DbFieldName: "b"},
		{DbFieldName: "a"},
	}
	got := handlePkeys(fields, []string{"a", "b"})
	if len(got) != 2 {
		t.Fatalf("len = %d", len(got))
	}
	if got[0].DbFieldName != "a" || got[1].DbFieldName != "b" {
		t.Errorf("handlePkeys should follow key order, got %v", got)
	}
}

func TestHandlePkeys_Empty(t *testing.T) {
	fields := []TableToStructFieldInfo{{DbFieldName: "id"}}
	if got := handlePkeys(fields, nil); len(got) != 0 {
		t.Errorf("handlePkeys(nil keys) = %v, want empty", got)
	}
}

func TestHandlePkeys_UnknownKey(t *testing.T) {
	fields := []TableToStructFieldInfo{{DbFieldName: "id"}}
	got := handlePkeys(fields, []string{"nonexistent"})
	if len(got) != 0 {
		t.Errorf("handlePkeys with unknown key should return empty, got %v", got)
	}
}

func TestDbFieldNames(t *testing.T) {
	fields := []TableToStructFieldInfo{
		{DbFieldName: "id"},
		{DbFieldName: "name"},
		{DbFieldName: "email"},
	}
	got := dbFieldNames(fields)
	want := []string{"id", "name", "email"}
	if len(got) != len(want) {
		t.Fatalf("dbFieldNames len = %d, want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("dbFieldNames[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestScanFields_SameAsFullFields(t *testing.T) {
	s := TableToStructInfo{
		Fields: []TableToStructFieldInfo{
			{DbFieldName: "id"},
			{DbFieldName: "name"},
			{DbFieldName: "created_at"},
		},
	}
	full := s.FullFields()
	scan := s.ScanFields()
	if len(full) != len(scan) {
		t.Errorf("ScanFields len=%d != FullFields len=%d", len(scan), len(full))
	}
}
