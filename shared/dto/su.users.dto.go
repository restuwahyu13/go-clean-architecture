package dto

/**
* ===================================================
*  USERS REQUEST TERITORY
* ===================================================
 */

type (
	LoginDTO struct {
		Email    string `json:"email" db:"email" validator:"required,email"`
		Password string `json:"password" validator:"required"`
	}
)

/**
* ===================================================
*  USERS RESPONSE TERITORY
* ===================================================
 */

type (
	Login struct {
		Role    string `json:"role"`
		Token   string `json:"token"`
		Expired int    `json:"expired"`
	}

	Users struct {
		ID        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Email     string `json:"email,omitempty"`
		Status    string `json:"status,omitempty"`
		Password  string `json:"password,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		DeletedAt string `json:"deleted_at,omitempty"`
	}
)
