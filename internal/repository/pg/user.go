package pg

import (
	"context"
	"database/sql"
	"fmt"

	"task5/internal/repository/pg/entity"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IUserRepository interface {
	GetUser(ctx context.Context, id int) (*entity.User, error)
	GetUsers(ctx context.Context) (entity.UserSlice, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) (*entity.User, error)
	GetUserByLogPass(ctx context.Context, login, password string) (*entity.User, error)
	CreateUserToken(ctx context.Context, token *entity.Token) (*entity.Token, error)
	GetRole(ctx context.Context, user_id int) (*entity.Role, error)
	GetPermission(ctx context.Context, role_id int, method, url string) (*entity.Permission, error)
	CreateRole(ctx context.Context, role *entity.Role) (*entity.Role, error)
	CreatePermission(ctx context.Context, permission *entity.Permission) (*entity.Permission, error)
	SetUserRole(ctx context.Context, role_id, user_id int) error
	SetRolePermission(ctx context.Context, role_id, permission_id int) error
	DeleteRole(ctx context.Context, role_id int) (*entity.Role, error)
	ResetUsersRoles(ctx context.Context, role_id, user_id int) error
	DeletePermission(ctx context.Context, permission_id int) (*entity.Permission, error)
	ResetRolesPermissions(ctx context.Context, role_id, permission_id int) (error)
	GetRefresh(ctx context.Context, token string) (*entity.Token, error)
	GetTokenByRefresh(ctx context.Context, refresh string) (*entity.Token, error)
}

type UserRepository struct {
	conn *sql.DB
}

func (u *UserRepository) GetUser(ctx context.Context, id int) (*entity.User, error) {
	user, err := entity.Users(entity.UserWhere.ID.EQ(int64(id))).One(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetUsers(ctx context.Context) (entity.UserSlice, error) {
	users, err := entity.Users().All(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	qCheckUser := `SELECT EXISTS (SELECT * FROM users WHERE login=$1);`
	var checkUser bool
	err := u.conn.QueryRow(qCheckUser, user.Login).Scan(&checkUser)
	if err != nil {
		return nil, err
	}
	if checkUser {
		return nil, fmt.Errorf("input other login")
	}
	err = user.Insert(ctx, u.conn, boil.Infer())
	if err != nil {
		return nil, err
	} 
	return user, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := entity.Users(entity.UserWhere.ID.EQ(user.ID)).UpdateAll(ctx, u.conn, entity.M{
		"firstname": user.Firstname,
		"lastname": user.Lastname,
		"email": user.Email,
		"login": user.Login,
		"password": user.Password,
		"haspremium": user.Haspremium,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id int) (*entity.User, error) {
	user, err := u.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = entity.Users(entity.UserWhere.ID.EQ(int64(id))).DeleteAll(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByLogPass(ctx context.Context, login, password string) (*entity.User, error) {
	user := entity.User{}

	q := "SELECT id, login, password FROM users WHERE login=$1 AND password=$2;"
	err := u.conn.QueryRow(q, login, []byte(password)).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("not correct query log pass")
	}
	return &user, nil
}

func (u *UserRepository) CreateUserToken(ctx context.Context, token *entity.Token) (*entity.Token, error) {
	qCheckToken := `SELECT EXISTS (SELECT id FROM tokens WHERE user_id=$1);`
	var checkToken bool
	err := u.conn.QueryRow(qCheckToken, token.UserID).Scan(&checkToken)
	if err != nil {
		return nil, err
	}
	if !checkToken {
		err := token.Insert(ctx, u.conn, boil.Infer())
		if err != nil {
			return nil, err
		}
	} else {
		_, err := entity.Tokens(entity.TokenWhere.UserID.EQ(token.UserID)).UpdateAll(ctx, u.conn, entity.M{
			"token": token.Token,
			"refresh": token.Refresh,
			"expires_at": token.ExpiresAt,
		})
		if err != nil {
			return nil, err
		}
	}
	return token, nil
}

func (u *UserRepository) GetRole(ctx context.Context, user_id int) (*entity.Role, error) {
	role := entity.Role{}
	q := `SELECT id, name, created_at, updated_at, deleted_at FROM roles 
	WHERE id = (SELECT role_id FROM users_roles 
		WHERE user_id=$1 
		ORDER BY role_id DESC 
		LIMIT 1)`
	err := u.conn.QueryRow(q, user_id).Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (u *UserRepository) GetPermission(ctx context.Context, role_id int, method, url string) (*entity.Permission, error) {
	permission := entity.Permission{}
	methodUrl := method + " " + url
	q := `SELECT permissions.id, permissions.name, permissions.created_at, permissions.updated_at, permissions.deleted_at FROM permissions
	INNER JOIN roles_permissions AS rp ON permissions.id=rp.permission_id
	INNER JOIN users_roles AS ur ON rp.role_id=ur.role_id
	WHERE ur.role_id=$1 AND permissions.name=$2`
	err := u.conn.QueryRow(q, role_id, methodUrl).Scan(&permission.ID, &permission.Name, &permission.CreatedAt, &permission.UpdatedAt, &permission.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (u *UserRepository) CreateRole(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	err := role.Insert(ctx, u.conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (u *UserRepository) CreatePermission(ctx context.Context, permission *entity.Permission) (*entity.Permission, error) {
	err := permission.Insert(ctx, u.conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (u *UserRepository) SetUserRole(ctx context.Context, role_id, user_id int) error {
	q := `INSERT INTO users_roles(user_id, role_id) VALUES($1, $2)`
	_, err := u.conn.Exec(q, user_id, role_id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) SetRolePermission(ctx context.Context, role_id, permission_id int) error {
	q := `INSERT INTO roles_permissions(role_id, permission_id) VALUES($1, $2)`
	_, err := u.conn.Exec(q, role_id, permission_id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteRole(ctx context.Context, role_id int) (*entity.Role, error) {
	role, err := entity.Roles(entity.RoleWhere.ID.EQ(int64(role_id))).One(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	_, err = entity.Roles(entity.RoleWhere.ID.EQ(int64(role_id))).DeleteAll(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (u *UserRepository) ResetUsersRoles(ctx context.Context, role_id, user_id int) error {
	q := `DELETE FROM users_roles WHERE user_id=$1 AND role_id=$2;`
	_, err := u.conn.Exec(q, user_id, role_id)
	if err != nil {
		return err
	}
	return nil
}

func(u *UserRepository) DeletePermission(ctx context.Context, permission_id int) (*entity.Permission, error) {
	permission, err := entity.Permissions(entity.PermissionWhere.ID.EQ(int64(permission_id))).One(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	_, err = entity.Permissions(entity.PermissionWhere.ID.EQ(int64(permission_id))).DeleteAll(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func(u *UserRepository) ResetRolesPermissions(ctx context.Context, role_id, permission_id int) (error) {
	q := `DELETE FROM roles_permissions WHERE role_id=$1 AND permission_id=$2`
	_, err := u.conn.Exec(q, role_id, permission_id)
	if err != nil {
		return err
	}
	return nil
}

func(u *UserRepository) GetRefresh(ctx context.Context, token string) (*entity.Token, error) {
	t, err := entity.Tokens(entity.TokenWhere.Token.EQ([]byte(token))).One(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func(u *UserRepository) GetTokenByRefresh(ctx context.Context, refresh string) (*entity.Token, error) {
	token, err := entity.Tokens(entity.TokenWhere.Refresh.EQ([]byte(refresh))).One(ctx, u.conn)
	if err != nil {
		return nil, err
	}
	return token, nil
}
	

func NewUserRepository(conn *sql.DB) IUserRepository {
	return &UserRepository{
		conn: conn,
	}
}
