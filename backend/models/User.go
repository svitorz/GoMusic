package models

import (
	"database/sql/driver"
	"time"
)

type User struct {
	ID        uint      `json:"id"`                                                                       // Standard field for the primary key
	Name      string    `json:"name"`                                                                     // A regular string field
	Email     *string   `json:"email"`                                                                    // A pointer to a string, allowing for null values
	Password  string    `json:"-"`                                                                        // Excluded from JSON serialization for security
	Role      Role      `json:"role" gorm:"type:enum('admin', 'customer', 'creator');default:'customer'"` // Enum field with default value
	CreatedAt time.Time `json:"created_at"`                                                               // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updated_at"`                                                               // Automatically managed by GORM for update time
}

type Role string

const (
	Admin    Role = "admin"
	Customer Role = "customer"
	Creator  Role = "creator"
)

func (p *Role) Scan(value any) error {
	*p = Role(value.([]byte))
	return nil
}

func (p Role) Value() (driver.Value, error) {
	return string(p), nil
}
