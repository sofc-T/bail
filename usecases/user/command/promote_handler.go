package usercmd

import irepo "bail/usecases/core/i_repo"




type PromoteHandler struct {
	Userrepo           irepo.User
}

// Promote Config holds the configuration for creating a Promote Handler.
type PromoteConfig struct {
	UserRepo           irepo.User
}


