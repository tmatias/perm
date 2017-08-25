package perm

import (
	"fmt"
)

// Description returns an human-friendly description of unix file permission bits
func Description(bits int) (string, error) {
	switch bits {
	case 0:
		return "cannot do anything", nil
	case 1:
		return "can execute it", nil
	case 2:
		return "can write to it", nil
	case 3:
		return "can execute and write to it", nil
	case 4:
		return "can read from it", nil
	case 5:
		return "can read and execute it", nil
	case 6:
		return "can read and write to it", nil
	case 7:
		return "can read, write and execute it", nil
	default:
		return "", fmt.Errorf("Invalid bits %d", bits)
	}
}

// Permissions return three integers containing permissions for the owner, group and everyone else
// from an int representing unix file permissions bits
func Permissions(bits int) (owner int, group int, other int) {
	owner = bits >> 6 & 7
	group = bits >> 3 & 7
	other = bits & 7
	return
}
