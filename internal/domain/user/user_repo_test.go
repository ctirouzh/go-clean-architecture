package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserRepoMock_NewMockRepository(t *testing.T) {
	repo := NewMockRepository()
	assert.NotEmpty(t, repo)
	assert.IsType(t, map[uuid.UUID]*User{}, repo.data)
	assert.Equal(t, 0, len(repo.data))
}

func TestUserRepoMock_AddUsers(t *testing.T) {
	users := []*User{}
	for i := 1; i < 5; i++ {
		// user := sample.NewFakeUserEntity(USER_TYPE_ADMIN, false, false)
		// compilor --> import cycle not allowed in test
		usr := &User{
			ID: uuid.New(),
		}
		users = append(users, usr)
	}

	repo := NewMockRepository()
	repo.AddUsers(users)
	for _, want := range users {
		got, exists := repo.data[want.ID]
		assert.True(t, exists)
		assert.Equal(t, want, got)
	}
}

func TestUserRepoMock_Get(t *testing.T) {
	usr := &User{
		ID: uuid.New(),
	}

	repo := NewMockRepository()
	repo.AddUsers([]*User{usr})

	testCases := []struct {
		name string
		id   uuid.UUID
		want error
	}{
		{
			name: "should find user",
			id:   usr.ID,
			want: nil,
		},
		{
			name: "try to find a non-existing user",
			id:   uuid.New(),
			want: ErrUserNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotUser, got := repo.Get(tc.id)
			assert.Equal(t, tc.want, got)
			if got == nil {
				assert.Equal(t, usr.ID, gotUser.ID)
			}
		})
	}
}

func TestUserRepoMock_GetByUsername(t *testing.T) {
	want := &User{
		ID:       uuid.New(),
		Username: "test_username",
	}

	repo := NewMockRepository()
	repo.AddUsers([]*User{want})

	got, err := repo.GetByUsername(want.Username)
	assert.Empty(t, err)
	assert.Equal(t, want, got)
}

func TestUserRepoMock_GetByEmail(t *testing.T) {
	want := &User{
		ID:       uuid.New(),
		Username: "test@test.io",
	}

	repo := NewMockRepository()
	repo.AddUsers([]*User{want})

	got, err := repo.GetByEmail(want.Email)
	assert.Empty(t, err)
	assert.Equal(t, want, got)
}

func TestUserRepoMock_Create(t *testing.T) {

	usr := &User{
		ID:       uuid.New(),
		Username: "test",
		Email:    "test@test.io",
	}
	repo := NewMockRepository()
	repo.AddUsers([]*User{usr})

	testCases := []struct {
		name     string
		username string
		email    string
		password string
		userType UserType
		want     error
	}{
		{
			name:     "should be created",
			username: "test1",
			email:    "test1@test.io",
			password: "secret",
			userType: USER_TYPE_STUDENT,
			want:     nil,
		},
		{
			name:     "username already taken",
			username: "test",
			email:    "test2@test.io",
			password: "secret",
			userType: USER_TYPE_STUDENT,
			want:     ErrUsernameAlreadyTaken,
		}, {
			name:     "email already taken",
			username: "test3",
			email:    "test@test.io",
			password: "secret",
			userType: USER_TYPE_STUDENT,
			want:     ErrEmailAlreadyTaken,
		}, {
			name:     "empty password",
			username: "test4",
			email:    "test4@test.io",
			password: "",
			userType: USER_TYPE_STUDENT,
			want:     ErrEmptyUserPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			usr, got := repo.Create(tc.username, tc.email, tc.password, tc.userType)
			assert.Equal(t, tc.want, got)
			if got == nil { // usr created/exists in repository/not nil
				assert.Equal(t, tc.username, usr.Username)
				assert.Equal(t, tc.email, usr.Email)
				assert.True(t, usr.IsPasswordVerified(tc.password))
			}
		})
	}
}

func TestUserRepoMock_Delete(t *testing.T) {

	usr := &User{
		ID: uuid.New(),
	}
	repo := NewMockRepository()
	repo.AddUsers([]*User{usr})

	testCases := []struct {
		name string
		id   uuid.UUID
		want error
	}{
		{
			name: "should be deleted",
			id:   usr.ID,
			want: nil,
		},
		{
			name: "try to delete a non-existing user",
			id:   uuid.New(),
			want: ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := repo.Delete(tc.id)
			assert.Equal(t, tc.want, got)
		})
	}
}
