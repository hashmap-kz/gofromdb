package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exp = "select * from users"

func TestPrettifierSpaces(t *testing.T) {
	pretty := PrettySQL("select    *    from    users")
	assert.Equal(t, exp, pretty)
}

func TestPrettifierNewlines(t *testing.T) {
	pretty := PrettySQL(`
		select *
		from users
	`)
	assert.Equal(t, exp, pretty)
}

func TestPrettifierTabs(t *testing.T) {
	pretty := PrettySQL(`
		select		*
			from	users
	`)
	assert.Equal(t, exp, pretty)
}
