// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var UserFieldPathsNested = []string{
	"admin",
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"ids.email",
	"ids.user_id",
	"name",
	"password",
	"password_updated_at",
	"primary_email_address",
	"primary_email_address_validated_at",
	"profile_picture",
	"profile_picture.embedded",
	"profile_picture.embedded.data",
	"profile_picture.embedded.mime_type",
	"profile_picture.sizes",
	"require_password_update",
	"state",
	"temporary_password",
	"temporary_password_created_at",
	"temporary_password_expires_at",
	"updated_at",
}

var UserFieldPathsTopLevel = []string{
	"admin",
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"name",
	"password",
	"password_updated_at",
	"primary_email_address",
	"primary_email_address_validated_at",
	"profile_picture",
	"require_password_update",
	"state",
	"temporary_password",
	"temporary_password_created_at",
	"temporary_password_expires_at",
	"updated_at",
}
var UsersFieldPathsNested = []string{
	"users",
}

var UsersFieldPathsTopLevel = []string{
	"users",
}
var GetUserRequestFieldPathsNested = []string{
	"field_mask",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var GetUserRequestFieldPathsTopLevel = []string{
	"field_mask",
	"user_ids",
}
var CreateUserRequestFieldPathsNested = []string{
	"invitation_token",
	"user",
	"user.admin",
	"user.attributes",
	"user.contact_info",
	"user.created_at",
	"user.description",
	"user.ids",
	"user.ids.email",
	"user.ids.user_id",
	"user.name",
	"user.password",
	"user.password_updated_at",
	"user.primary_email_address",
	"user.primary_email_address_validated_at",
	"user.profile_picture",
	"user.profile_picture.embedded",
	"user.profile_picture.embedded.data",
	"user.profile_picture.embedded.mime_type",
	"user.profile_picture.sizes",
	"user.require_password_update",
	"user.state",
	"user.temporary_password",
	"user.temporary_password_created_at",
	"user.temporary_password_expires_at",
	"user.updated_at",
}

var CreateUserRequestFieldPathsTopLevel = []string{
	"invitation_token",
	"user",
}
var UpdateUserRequestFieldPathsNested = []string{
	"field_mask",
	"user",
	"user.admin",
	"user.attributes",
	"user.contact_info",
	"user.created_at",
	"user.description",
	"user.ids",
	"user.ids.email",
	"user.ids.user_id",
	"user.name",
	"user.password",
	"user.password_updated_at",
	"user.primary_email_address",
	"user.primary_email_address_validated_at",
	"user.profile_picture",
	"user.profile_picture.embedded",
	"user.profile_picture.embedded.data",
	"user.profile_picture.embedded.mime_type",
	"user.profile_picture.sizes",
	"user.require_password_update",
	"user.state",
	"user.temporary_password",
	"user.temporary_password_created_at",
	"user.temporary_password_expires_at",
	"user.updated_at",
}

var UpdateUserRequestFieldPathsTopLevel = []string{
	"field_mask",
	"user",
}
var CreateTemporaryPasswordRequestFieldPathsNested = []string{
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var CreateTemporaryPasswordRequestFieldPathsTopLevel = []string{
	"user_ids",
}
var UpdateUserPasswordRequestFieldPathsNested = []string{
	"new",
	"old",
	"revoke_all_access",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var UpdateUserPasswordRequestFieldPathsTopLevel = []string{
	"new",
	"old",
	"revoke_all_access",
	"user_ids",
}
var ListUserAPIKeysRequestFieldPathsNested = []string{
	"limit",
	"page",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var ListUserAPIKeysRequestFieldPathsTopLevel = []string{
	"limit",
	"page",
	"user_ids",
}
var GetUserAPIKeyRequestFieldPathsNested = []string{
	"key_id",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var GetUserAPIKeyRequestFieldPathsTopLevel = []string{
	"key_id",
	"user_ids",
}
var CreateUserAPIKeyRequestFieldPathsNested = []string{
	"name",
	"rights",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var CreateUserAPIKeyRequestFieldPathsTopLevel = []string{
	"name",
	"rights",
	"user_ids",
}
var UpdateUserAPIKeyRequestFieldPathsNested = []string{
	"api_key",
	"api_key.id",
	"api_key.key",
	"api_key.name",
	"api_key.rights",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var UpdateUserAPIKeyRequestFieldPathsTopLevel = []string{
	"api_key",
	"user_ids",
}
var InvitationFieldPathsNested = []string{
	"accepted_at",
	"accepted_by",
	"accepted_by.email",
	"accepted_by.user_id",
	"created_at",
	"email",
	"expires_at",
	"token",
	"updated_at",
}

var InvitationFieldPathsTopLevel = []string{
	"accepted_at",
	"accepted_by",
	"created_at",
	"email",
	"expires_at",
	"token",
	"updated_at",
}
var ListInvitationsRequestFieldPathsNested = []string{
	"limit",
	"page",
}

var ListInvitationsRequestFieldPathsTopLevel = []string{
	"limit",
	"page",
}
var InvitationsFieldPathsNested = []string{
	"invitations",
}

var InvitationsFieldPathsTopLevel = []string{
	"invitations",
}
var SendInvitationRequestFieldPathsNested = []string{
	"email",
}

var SendInvitationRequestFieldPathsTopLevel = []string{
	"email",
}
var DeleteInvitationRequestFieldPathsNested = []string{
	"email",
}

var DeleteInvitationRequestFieldPathsTopLevel = []string{
	"email",
}
var UserSessionIdentifiersFieldPathsNested = []string{
	"session_id",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var UserSessionIdentifiersFieldPathsTopLevel = []string{
	"session_id",
	"user_ids",
}
var UserSessionFieldPathsNested = []string{
	"created_at",
	"expires_at",
	"session_id",
	"updated_at",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var UserSessionFieldPathsTopLevel = []string{
	"created_at",
	"expires_at",
	"session_id",
	"updated_at",
	"user_ids",
}
var UserSessionsFieldPathsNested = []string{
	"sessions",
}

var UserSessionsFieldPathsTopLevel = []string{
	"sessions",
}
var ListUserSessionsRequestFieldPathsNested = []string{
	"limit",
	"order",
	"page",
	"user_ids",
	"user_ids.email",
	"user_ids.user_id",
}

var ListUserSessionsRequestFieldPathsTopLevel = []string{
	"limit",
	"order",
	"page",
	"user_ids",
}
