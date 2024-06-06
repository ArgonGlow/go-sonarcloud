package organization

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddUserRequest Add a user to an organization.<br>Requires the 'Administer' permission on the specified organization.
type AddUserRequest struct {
	Organization string `form:"organization,omitempty"` // Organization key
	Login        string `form:"login,omitempty"`        // User login
}
