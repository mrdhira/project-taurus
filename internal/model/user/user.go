package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	FullName  string     `db:"full_name" json:"full_name"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"password"`
	PIN       string     `db:"pin" json:"pin"`
	Status    Status     `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (u *User) HashPassword() error {
	// Generate a bcrypt hash of the password with a default cost of bcrypt.DefaultCost (which is 10)
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the User.Password field to the generated hash
	u.Password = string(hash)

	return nil
}

func (u *User) ComparePassword(inputPassword string) error {
	// Compare the stored hashed password with the password the user submits
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword))
}

func (u *User) HashPIN() error {
	// Generate a bcrypt hash of the PIN with a default cost of bcrypt.DefaultCost (which is 10)
	hash, err := bcrypt.GenerateFromPassword([]byte(u.PIN), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the User.PIN field to the generated hash
	u.PIN = string(hash)

	return nil
}

func (u *User) ComparePIN(inputPIN string) error {
	// Compare the stored hashed pin with the pin the user submits
	return bcrypt.CompareHashAndPassword([]byte(u.PIN), []byte(inputPIN))
}

func (u *User) StatusToString() string {
	return string(u.Status)
}
