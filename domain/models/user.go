package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents the base struct for shared fields and methods.
type User struct {
	id                  uuid.UUID
	name                string
	email               string
	salary              float64
	age                 int
	role                string
	coSignerName        string
	codeNumber          string
	coSignerDocument    []byte 
	educationalDocument []byte
	createdAt           time.Time
	updatedAt           time.Time
	password            string
	branchCode          string
}

type UserConfig struct {
	ID                  uuid.UUID
	Name                string
	Email               string
	Salary              float64
	Age                 int
	Role                string
	CoSignerName        string
	CodeNumber          string
	CoSignerDocument    []byte
	EducationalDocument []byte
	UpdatedAt           time.Time
	CreatedAt           time.Time
	Password            string
	branchCode		  string
}

func (u *User) UpdateTimestamps() {
	u.updatedAt = time.Now()
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Salary() float64 {
	return u.salary
}

func (u *User) Age() int {
	return u.age
}

func (u *User) Role() string {
	return u.role
}

func (u *User) CoSignerName() string {
	return u.coSignerName
}

func (u *User) BranchCode() string {
	return u.branchCode
}

func (u *User) CodeNumber() string {
	return u.codeNumber
}

func (u *User) CoSignerDocument() []byte {
	return u.coSignerDocument
}

func (u *User) EducationalDocument() []byte {
	return u.educationalDocument
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetID(id uuid.UUID) {
	u.id = id
	u.UpdateTimestamps()
}

func (u *User) SetName(name string) {
	u.name = name
	u.UpdateTimestamps()
}

func (u *User) SetBranchCode(branchCode string) {
	u.branchCode = branchCode
	u.UpdateTimestamps()
}

func (u *User) SetEmail(email string) {
	u.email = email
	u.UpdateTimestamps()
}

func (u *User) SetSalary(salary float64) {
	u.salary = salary
	u.UpdateTimestamps()
}

func (u *User) SetAge(age int) {
	u.age = age
	u.UpdateTimestamps()
}

func (u *User) SetRole(role string) {
	u.role = role
	u.UpdateTimestamps()
}

func (u *User) SetCoSignerName(coSignerName string) {
	u.coSignerName = coSignerName
	u.UpdateTimestamps()
}

func (u *User) SetCodeNumber(codeNumber string) {
	u.codeNumber = codeNumber
	u.UpdateTimestamps()
}

func (u *User) SetCoSignerDocument(coSignerDocument []byte) {
	u.coSignerDocument = coSignerDocument
	u.UpdateTimestamps()
}

func (u *User) SetEducationalDocument(educationalDocument []byte) {
	u.educationalDocument = educationalDocument
	u.UpdateTimestamps()
}

func (u *User) SetPassword(password string) {
	u.password = password
	u.UpdateTimestamps()
}

func NewUser(config UserConfig) *User {
	return &User{
		id:                  uuid.New(),
		name:                config.Name,
		email:               config.Email,
		salary:              config.Salary,
		age:                 config.Age,
		role:                config.Role,
		coSignerName:        config.CoSignerName,
		codeNumber:          config.CodeNumber,
		coSignerDocument:    config.CoSignerDocument,
		educationalDocument: config.EducationalDocument,
		createdAt:           config.CreatedAt,
		updatedAt:           config.UpdatedAt,
		password:            config.Password,
		branchCode:          config.branchCode,
	}
}

func MapUser(config UserConfig) *User {
	u := &User{
		id:                  config.ID,
		name:                config.Name,
		email:               config.Email,
		salary:              config.Salary,
		age:                 config.Age,
		role:                config.Role,
		coSignerName:        config.CoSignerName,
		codeNumber:          config.CodeNumber,
		coSignerDocument:    config.CoSignerDocument,
		educationalDocument: config.EducationalDocument,
		updatedAt:           config.UpdatedAt,
		createdAt:           config.CreatedAt,
		password:            config.Password,
		branchCode:          config.branchCode,
	}
	u.UpdateTimestamps()
	return u
}
