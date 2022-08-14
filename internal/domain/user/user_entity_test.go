package user

import (
	"testing"

	"github.com/google/uuid"
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

func TestUserEntity_PrepareForCreate(t *testing.T) {
	usr := User{}
	usr.PrepareForCreate()
	if _, err := uuid.Parse(usr.ID.String()); err != nil {
		t.Error("expected a valid uuid, got invalid one")
	}

	if usr.CreatedAt.IsZero() {
		t.Error("expected a timestamp, got zero CreatedAt")
	}

	if usr.UpdatedAt.IsZero() {
		t.Error("expected a timestamp, got zero UpdatedAt")
	}

	if usr.IsVerified() {
		t.Error("expected an unverified user, got a verified one")
	}
	if usr.IsBanned() {
		t.Error("expected a permitted user, got a banned one")
	}
}
