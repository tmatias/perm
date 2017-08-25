package perm

import (
	"testing"
)

func TestPermDescription(t *testing.T) {
	tt := []struct {
		name string
		bits int
		desc string
	}{
		{"no permission", 0, "cannot do anything"},
		{"execute permission", 1, "can execute it"},
		{"write permission", 2, "can write to it"},
		{"execute and write permission", 3, "can execute and write to it"},
		{"read permission", 4, "can read from it"},
		{"read and execute permission", 5, "can read and execute it"},
		{"write permission", 6, "can read and write to it"},
		{"read, write and execute permission", 7, "can read, write and execute it"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			desc, err := Description(tc.bits)
			if err != nil {
				t.Error(err)
			}
			if desc != tc.desc {
				t.Errorf("expected description of %d to be '%s'; got '%s'", tc.bits, tc.desc, desc)
			}
		})
	}
}

func TestPermissions(t *testing.T) {
	tt := []struct {
		name  string
		bits  int
		owner int
		group int
		other int
	}{
		{"world readable and writable", 0777, 7, 7, 7},
		{"readable only by the owner", 0400, 4, 0, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			owner, group, other := Permissions(tc.bits)
			if owner != tc.owner {
				t.Errorf("expected owner permission for %d to be %d; got %d", tc.bits, tc.owner, owner)
			}
			if group != tc.group {
				t.Errorf("expected group permission for %d to be %d; got %d", tc.bits, tc.group, owner)
			}
			if other != tc.other {
				t.Errorf("expected other permission for %d to be %d; got %d", tc.bits, tc.other, owner)
			}
		})
	}
}
