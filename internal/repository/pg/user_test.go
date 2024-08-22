package pg

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"task5/internal/repository/pg/entity"
)

var user = entity.User{
	Firstname:  "bibo",
	Lastname:   "bobo",
	Login:      "bobobibo",
	Password:   []byte("qwerty"),
	Email:      "bobobibo@ds.ru",
	Haspremium: false,
}

func TestGetUsers(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	repo := NewUserRepository(conn)

	_, err = repo.GetUsers(context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestCreate(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	repo := NewUserRepository(conn)

	u, err := repo.CreateUser(context.TODO(), &user)
	if err != nil {
		t.Error(err)
	}
	check := checkUser(u)
	if !check {
		t.Error(fmt.Errorf("error not correct field"))
	}

	_, err = repo.DeleteUser(context.TODO(), int(u.ID))
	if err != nil {
		t.Error(fmt.Errorf("%e ERROR DELETE USER", err))
	}
}

func TestGetUser(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	repo := NewUserRepository(conn)

	_, err = repo.CreateUser(context.TODO(), &user)
	if err != nil {
		t.Error(err)
	}
	u, err := repo.GetUser(context.TODO(), int(user.ID))
	if err != nil {
		t.Error(err)
	}
	if !checkUser(u) {
		t.Error(fmt.Errorf("not correct"))
	}
	_, err = repo.DeleteUser(context.TODO(), int(user.ID))
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUser(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	repo := NewUserRepository(conn)
	updateUser := user
	updateUser.Firstname = "otherBobo"

	createUser, err := repo.CreateUser(context.TODO(), &user)
	if err != nil {
		t.Error(err)
	}
	updateUser.ID = createUser.ID

	_, err = repo.UpdateUser(context.TODO(), &updateUser)
	if err != nil {
		t.Error(err)
	}

	u, err := repo.GetUser(context.TODO(), int(updateUser.ID))
	if err != nil {
		t.Error(err)
	}

	if u.Firstname != "otherBobo" {
		t.Error("not update name")
	}

	_, err = repo.DeleteUser(context.TODO(), int(updateUser.ID))
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	repo := NewUserRepository(conn)

	createUser, err := repo.CreateUser(context.TODO(), &user)
	if err != nil {
		t.Error(err)
	}

	_, err = repo.DeleteUser(context.TODO(), int(createUser.ID))
	if err != nil {
		t.Error(err)
	}
}

func checkUser(user *entity.User) bool {
	switch {
	case user.Firstname != "bibo":
		return false
	case user.Lastname != "bobo":
		return false
	case user.Email != "bobobibo@ds.ru":
		return false
	case user.Login != "bobobibo":
		return false
	case string(user.Password) != "qwerty":
		return false
	case user.Haspremium != false:
		return false
	}
	return true
}

func TestGetRole(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	userRepo := NewUserRepository(conn)
	
	createUser := user
	user, err := userRepo.CreateUser(ctx, &createUser)
	if err != nil {
		t.Error(err)
	}
	role, err := userRepo.CreateRole(ctx, &entity.Role{Name: "testAdmin"})
	if err != nil {
		t.Error()
	}
	err = userRepo.SetUserRole(ctx, int(role.ID), int(user.ID))
	if err != nil {
		t.Error(err)
	}
	role, err = userRepo.GetRole(context.TODO(), int(user.ID))
	if err != nil {
		t.Error(err)
	}

	if role.Name != "testAdmin" {
		t.Error(fmt.Errorf("not correct role"))
	}
	err = userRepo.ResetUsersRoles(ctx, int(role.ID), int(user.ID))
	if err != nil {
		t.Error(err)
	}
	_, err = userRepo.DeleteUser(ctx, int(user.ID))
	if err != nil {
		t.Error(err)
	}
	_, err = userRepo.DeleteRole(ctx, int(role.ID))
	if err != nil {
		t.Error(err)
	}
}

func TestGetPermission(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	userRepo := NewUserRepository(conn)
	
	createUser := user
	user, err := userRepo.CreateUser(ctx, &createUser)
	if err != nil {
		t.Error(err)
	}
	role, err := userRepo.CreateRole(ctx, &entity.Role{Name: "testAdmin"})
	if err != nil {
		t.Error()
	}
	err = userRepo.SetUserRole(ctx, int(role.ID), int(user.ID))
	if err != nil {
		t.Error(err)
	}

	permission, err := userRepo.CreatePermission(ctx, &entity.Permission{Name: "test Permission"})
	if err != nil {
		t.Error(err)
	}

	err = userRepo.SetRolePermission(ctx, int(role.ID), int(permission.ID))
	if err != nil {
		t.Error(err)
	}

	permission, err = userRepo.GetPermission(ctx, int(role.ID), "test", "Permission")
	if err != nil {
		t.Error(err)
	}

	if permission.Name != "test Permission" {
		t.Error("not correct permission")
	}

	err = userRepo.ResetUsersRoles(ctx, int(role.ID), int(user.ID))
	if err != nil {
		t.Error(err)
	}
	err = userRepo.ResetRolesPermissions(ctx, int(role.ID), int(permission.ID))
	if err != nil {
		t.Error(err)
	}
	
	_, err = userRepo.DeleteUser(ctx, int(user.ID))
	if err != nil {
		t.Error(err)
	}
	_, err = userRepo.DeleteRole(ctx, int(role.ID))
	if err != nil {
		t.Error(err)
	}
	
	_, err = userRepo.DeletePermission(ctx, int(permission.ID))
	if err != nil {
		t.Error(err)
	}
}
