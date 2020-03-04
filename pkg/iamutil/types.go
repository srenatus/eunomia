package iamutil

// PolicyDocument represents an IAM policy document
type PolicyDocument struct {
	Version   string
	Statement []Statement
}

// Statement represents a single-action statement in an IAM policy
type Statement struct {
	Sid      string
	Effect   string
	Action   interface{}
	Resource interface{}
}
