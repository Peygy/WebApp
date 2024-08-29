package managers

import (
	"database/sql"
	"errors"

	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type IRoleManager interface {
	AddRolesToUser(userId string, roles ...string) error

	DeleteRolesFromUser(userId string, roles ...string) error
}

var roles = []string{
	"user",
	"admin",
}

type roleManger struct {
	db  *sql.DB
	log logger.ILogger
}

func NewRoleManager(db *sql.DB, log logger.ILogger) IRoleManager {
	for _, role := range roles {
		query := `INSERT INTO roles (role_name) VALUES ($1) ON CONFLICT (role_name) DO NOTHING`

		_, err := db.Exec(query, role)
		if err != nil {
			log.Errorf("Error during role %s insertion: %v", role, err.Error())
			return nil
		}

		log.Infof("Role %s inserts successful", role)
	}

	return &roleManger{db: db, log: log}
}

func (rm *roleManger) addRoleToUser(userId, role string) error {
	roleId, ok := rm.checkRoleExists(role)
	if !ok {
		return errors.New("managers-role: role " + role + " not exists in database")
	}

	query := `INSERT INTO users_roles (user_id, role_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := rm.db.Exec(query, userId, roleId)
	if err != nil {
		rm.log.Errorf("Can't inserts role %s to user %s: ", role, userId, err)
		return errors.New("managers-role: can't add role " + role + " to the " + userId)
	}

	return nil
}

func (rm *roleManger) AddRolesToUser(userId string, roles ...string) error {
	tx, err := rm.db.Begin()
	if err != nil {
		rm.log.Errorf("Can't starts transaction: %v", err)
		return errors.New("managers-role: can't starts transaction for adding roles to user")
	}
	defer tx.Rollback()

	for _, role := range roles {
		err := rm.addRoleToUser(userId, role)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		rm.log.Errorf("Can't commits transaction: %v", err)
		return errors.New("managers-role: can't commits transaction for adding role to user")
	}

	return nil
}

func (rm *roleManger) deleteRoleFromUser(userId, role string) error {
	roleId, ok := rm.checkRoleExists(role)
	if !ok {
		return errors.New("managers-role: role " + role + " not exists in database")
	}

	query := `DELETE FROM users_roles WHERE user_id = $1 AND role_id = $2`
	_, err := rm.db.Exec(query, userId, roleId)
	if err != nil {
		rm.log.Errorf("Can't deletes role %s from user %s: %v", role, userId, err)
		return errors.New("managers-role: can't delete role " + role + " from the " + userId)
	}

	return nil
}

func (rm *roleManger) DeleteRolesFromUser(userId string, roles ...string) error {
	tx, err := rm.db.Begin()
	if err != nil {
		rm.log.Errorf("Can't starts transaction: %v", err)
		return errors.New("managers-role: can't starts transaction for deleting roles from user")
	}
	defer tx.Rollback()

	for _, role := range roles {
		err := rm.deleteRoleFromUser(userId, role)
		if err != nil {
			rm.log.Errorf("Can't delete role %s from user %s", role, userId)
			return errors.New("managers-role: can't delete role " + role + " from user " + userId)
		}
	}

	if err = tx.Commit(); err != nil {
		rm.log.Errorf("Can't commits transaction: %v", err)
		return errors.New("managers-role: can't commits transaction for deleting roles from user")
	}

	return nil
}

func (rm *roleManger) checkRoleExists(role string) (string, bool) {
	var roleId string
	query := `SELECT id FROM roles WHERE role_name = $1`

	if err := rm.db.QueryRow(query).Scan(&roleId); err != nil {
		if err == sql.ErrNoRows {
			rm.log.Warnf("Not find role %s in table roles", role)
		} else {
			rm.log.Warnf("Error during find role %s: %v", role, err)
		}

		return "", false
	}

	return roleId, true
}
