package user

// StoredUserCollection is the result type of the users service list method.
type StoredUserCollection []*StoredUser

// ShowPayload is the payload type of the users service show method.
type ShowPayload struct {
	// Email of user to show
	Email string
	// View to render
	View *string
}

// StoredUser is the result type of the users service show method.
type StoredUser struct {
	// Email of the user
	Email string
	// First Name of the user
	Firstname string
	// Last Name of user
	Lastname string
	// Is user active.
	Isactive bool
	// user role
	Role string
}

// User is the payload type of the users service add method.
type User struct {
	// Email of the user
	Email string
	// First Name of the user
	Firstname string
	// Last Name of user
	Lastname string
	// user role
	Role string
	// Is user active.
	Isactive bool
}
