package model

type UserAccessConf struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}
type Statement struct {
	Effect    string    `json:"Effect"`
	Action    []string  `json:"Action"`
	Resource  []string  `json:"Resource"`
	Condition Condition `json:"Condition,omitempty"`
}
type Condition struct {
	StringLike `json:"StringLike"`
}
type StringLike struct {
	S3Prefix []string `json:"s3:prefix"`
}