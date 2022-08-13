package user

import (
	"testing"
)

func TestUserEntity_CheckUserType(t *testing.T) {

	type testCase struct {
		name     string
		user     User
		expected []bool
	}

	testCases := []testCase{
		{
			name:     "calling IsAdmin, IsTecher, and IsStudent on an admin user type",
			user:     User{Type: USER_TYPE_ADMIN},
			expected: []bool{true, false, false},
		}, {
			name:     "calling IsAdmin, IsTecher, and IsStudent on a teacher user type",
			user:     User{Type: USER_TYPE_TEACHER},
			expected: []bool{false, true, false},
		}, {
			name:     "calling IsAdmin, IsTecher, and IsStudent on a student user type",
			user:     User{Type: USER_TYPE_STUDENT},
			expected: []bool{false, false, true},
		}, {
			name:     "calling IsAdmin, IsTecher, and IsStudent on an unknown user type",
			user:     User{Type: -1},
			expected: []bool{false, false, false},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := []bool{tc.user.IsAdmin(), tc.user.IsTeacher(), tc.user.IsStudent()}
			for i := 0; i < len(got); i++ {
				if got[i] != tc.expected[i] {
					t.Errorf("expected %v, got %v", tc.expected, got)
				}
			}
		})
	}

}
