package user

import (
	"fmt"

	storage "github.com/asaberwd/users-api/pkg/db"
	"github.com/labstack/echo/v4"
)

const DB = "boltDB"

// RoleManager ...
type RoleManager interface {
	CheckRoleExists(string) (bool, error)
}

// Manager ...
type Manager struct {
	Db   storage.Db
	role RoleManager
}

// NewManager ...
func NewManager(db storage.Db, role RoleManager) *Manager {
	return &Manager{db, role}
}

// List ...
func (m *Manager) List(ctx echo.Context) (res StoredUserCollection, err error) {
	err = m.Db.LoadAll(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Show ...
func (m *Manager) Show(email string) (res *StoredUser, err error) {
	res = &StoredUser{}
	err = m.Db.Load(email, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Add ...
func (m *Manager) Add(p *User) (err error) {
	if _, err := m.role.CheckRoleExists(p.Role); err != nil {
		return err
	}

	sb := StoredUser{
		Email:     p.Email,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Isactive:  p.Isactive,
		Role:      p.Role,
	}
	if err = m.Db.Save(p.Email, &sb); err != nil {
		return err
	}
	return nil

}

// Update ...
func (m *Manager) Update(p *User) (err error) {
	if _, err := m.role.CheckRoleExists(p.Role); err != nil {
		return err
	}

	sb := StoredUser{
		Email:     p.Email,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Isactive:  p.Isactive,
		Role:      p.Role,
	}
	if err = m.Db.Save(p.Email, &sb); err != nil {
		return err
	}
	return nil
}

// Remove ...
func (m *Manager) Remove(email string) (err error) {
	return m.Db.Delete(email) // internal error if not nil
}

// Activate ...
func (m *Manager) Activate(p []string) (err error) {
	for _, email := range p {
		res := StoredUser{}
		fmt.Printf("activate: %v\n", email)

		err := m.Db.Load(email, &res)

		res.Isactive = true

		if err = m.Db.Save(email, &res); err != nil {
			return err // internal error
		}
	}

	return nil
}
