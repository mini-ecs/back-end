package service

import "github.com/mini-ecs/back-end/internal/model"

const (
	Version                     = "2012-10-17"
	Allow                       = "Allow"
	Prefix                      = "arn:aws:s3:::"
	RootDir                     = Prefix + "*"
	PermissionByUserName        = "${aws:username}"
	AllFolder                   = "/*"
	CommonUserFolder            = Prefix + PermissionByUserName
	CommonUserFolderAndChildren = CommonUserFolder + AllFolder

	// Permissions
	AllPermission = "s3:*"
	ListBucket    = "s3:ListBucket"
)

// TemplateUserConf 允许使用以自己用户名创建的桶
var TemplateUserConf = model.UserAccessConf{
	Version: Version,
	Statement: []model.Statement{
		{
			Effect:   Allow,
			Action:   []string{ListBucket},
			Resource: []string{RootDir},
			Condition: model.Condition{
				StringLike: model.StringLike{S3Prefix: []string{PermissionByUserName + AllFolder}},
			},
		},
		{
			Effect:   Allow,
			Action:   []string{AllPermission},
			Resource: []string{CommonUserFolderAndChildren},
		},
	},
}
