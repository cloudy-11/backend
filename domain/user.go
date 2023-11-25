package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
	ACTIVE_STATUS  = "active"
	PENDING_STATUS = "pending"
	BLOCK_STATUS   = "block"
	DANGER_STATUS  = "danger"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Handle   string             `bson:"handle" json:"handle"`
	Role     string             `bson:"role" json:"role"` //admin | user
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Status   string             `bson:"status" json:"status"`
	Password string             `bson:"password" json:"-"`
	Ghost    int32              `bson:"ghost" json:"ghost"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}

type UserUseCase interface {
	Patch(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	GetByID(c context.Context, id string) (*User, error)
}
