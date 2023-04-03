package partner

import "github.com/twharmon/govalid"

// Package-level singleton for the usecase model validations
var validator = govalid.New()
