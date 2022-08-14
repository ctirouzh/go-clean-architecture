package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserType_String(t *testing.T) {
	type testCase struct {
		name     string
		user     User
		expected string
	}

	testCases := []testCase{
		{
			name:     "get the string of student type",
			user:     User{Type: USER_TYPE_STUDENT},
			expected: "STUDENT",
		}, {
			name:     "get the string of teacher type",
			user:     User{Type: USER_TYPE_TEACHER},
			expected: "TEACHER",
		}, {
			name:     "get the string of admin type",
			user:     User{Type: USER_TYPE_ADMIN},
			expected: "ADMIN",
		}, {
			name:     "calling the String method of an unknown user type returns \"UNKNOWN\"",
			user:     User{Type: -53},
			expected: "UNKNOWN",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.user.Type.String()
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestUserType_Index(t *testing.T) {
	type testCase struct {
		name     string
		user     User
		expected int
	}

	testCases := []testCase{
		{
			name:     "get the index of student type",
			user:     User{Type: USER_TYPE_STUDENT},
			expected: int(3),
		}, {
			name:     "get the index of teacher type",
			user:     User{Type: USER_TYPE_TEACHER},
			expected: int(2),
		}, {
			name:     "get the index of admin type",
			user:     User{Type: USER_TYPE_ADMIN},
			expected: int(1),
		}, {
			name:     "calling the index method of an unknown user type returns -1",
			user:     User{Type: -53},
			expected: int(-1),
		},
		{
			name:     "calling the index method of an unknown user type returns -1",
			user:     User{Type: 53},
			expected: int(-1),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.user.Type.Index()
			assert.Equal(t, tc.expected, got)
		})
	}

}
