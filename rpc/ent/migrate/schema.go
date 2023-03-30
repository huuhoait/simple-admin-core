// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysApisColumns holds the columns for the "sys_apis" table.
	SysApisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "path", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "api_group", Type: field.TypeString},
		{Name: "method", Type: field.TypeString, Default: "POST"},
	}
	// SysApisTable holds the schema information for the "sys_apis" table.
	SysApisTable = &schema.Table{
		Name:       "sys_apis",
		Columns:    SysApisColumns,
		PrimaryKey: []*schema.Column{SysApisColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "api_path_method",
				Unique:  true,
				Columns: []*schema.Column{SysApisColumns[5], SysApisColumns[8]},
			},
		},
	}
	// SysAuditsColumns holds the columns for the "sys_audits" table.
	SysAuditsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "object_name", Type: field.TypeString},
		{Name: "action_name", Type: field.TypeString},
		{Name: "changed_data", Type: field.TypeString},
	}
	// SysAuditsTable holds the schema information for the "sys_audits" table.
	SysAuditsTable = &schema.Table{
		Name:       "sys_audits",
		Columns:    SysAuditsColumns,
		PrimaryKey: []*schema.Column{SysAuditsColumns[0]},
	}
	// SysDepartmentsColumns holds the columns for the "sys_departments" table.
	SysDepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "ancestors", Type: field.TypeString},
		{Name: "leader", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "remark", Type: field.TypeString},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true, Default: 0},
	}
	// SysDepartmentsTable holds the schema information for the "sys_departments" table.
	SysDepartmentsTable = &schema.Table{
		Name:       "sys_departments",
		Columns:    SysDepartmentsColumns,
		PrimaryKey: []*schema.Column{SysDepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_departments_sys_departments_children",
				Columns:    []*schema.Column{SysDepartmentsColumns[13]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysDictionariesColumns holds the columns for the "sys_dictionaries" table.
	SysDictionariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "title", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "desc", Type: field.TypeString},
	}
	// SysDictionariesTable holds the schema information for the "sys_dictionaries" table.
	SysDictionariesTable = &schema.Table{
		Name:       "sys_dictionaries",
		Columns:    SysDictionariesColumns,
		PrimaryKey: []*schema.Column{SysDictionariesColumns[0]},
	}
	// SysDictionaryDetailsColumns holds the columns for the "sys_dictionary_details" table.
	SysDictionaryDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "title", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "dictionary_id", Type: field.TypeUint64, Nullable: true},
	}
	// SysDictionaryDetailsTable holds the schema information for the "sys_dictionary_details" table.
	SysDictionaryDetailsTable = &schema.Table{
		Name:       "sys_dictionary_details",
		Columns:    SysDictionaryDetailsColumns,
		PrimaryKey: []*schema.Column{SysDictionaryDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_dictionary_details_sys_dictionaries_dictionary_details",
				Columns:    []*schema.Column{SysDictionaryDetailsColumns[10]},
				RefColumns: []*schema.Column{SysDictionariesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "menu_level", Type: field.TypeUint32},
		{Name: "menu_type", Type: field.TypeUint32},
		{Name: "path", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "name", Type: field.TypeString},
		{Name: "redirect", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "component", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "title", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString},
		{Name: "hide_menu", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "hide_breadcrumb", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "ignore_keep_alive", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "hide_tab", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "frame_src", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "carry_param", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "hide_children_in_menu", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "affix", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "dynamic_level", Type: field.TypeUint32, Nullable: true, Default: 20},
		{Name: "real_path", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true, Default: 100000},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_children",
				Columns:    []*schema.Column{SysMenusColumns[25]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenuParamsColumns holds the columns for the "sys_menu_params" table.
	SysMenuParamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "menu_id", Type: field.TypeUint64, Nullable: true},
	}
	// SysMenuParamsTable holds the schema information for the "sys_menu_params" table.
	SysMenuParamsTable = &schema.Table{
		Name:       "sys_menu_params",
		Columns:    SysMenuParamsColumns,
		PrimaryKey: []*schema.Column{SysMenuParamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menu_params_sys_menus_params",
				Columns:    []*schema.Column{SysMenuParamsColumns[8]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysOauthProvidersColumns holds the columns for the "sys_oauth_providers" table.
	SysOauthProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "client_id", Type: field.TypeString},
		{Name: "client_secret", Type: field.TypeString},
		{Name: "redirect_url", Type: field.TypeString},
		{Name: "scopes", Type: field.TypeString},
		{Name: "auth_url", Type: field.TypeString},
		{Name: "token_url", Type: field.TypeString},
		{Name: "auth_style", Type: field.TypeUint64},
		{Name: "info_url", Type: field.TypeString},
	}
	// SysOauthProvidersTable holds the schema information for the "sys_oauth_providers" table.
	SysOauthProvidersTable = &schema.Table{
		Name:       "sys_oauth_providers",
		Columns:    SysOauthProvidersColumns,
		PrimaryKey: []*schema.Column{SysOauthProvidersColumns[0]},
	}
	// SysPositionsColumns holds the columns for the "sys_positions" table.
	SysPositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "remark", Type: field.TypeString},
	}
	// SysPositionsTable holds the schema information for the "sys_positions" table.
	SysPositionsTable = &schema.Table{
		Name:       "sys_positions",
		Columns:    SysPositionsColumns,
		PrimaryKey: []*schema.Column{SysPositionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "position_code",
				Unique:  true,
				Columns: []*schema.Column{SysPositionsColumns[8]},
			},
		},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "default_router", Type: field.TypeString, Default: "dashboard"},
		{Name: "remark", Type: field.TypeString, Default: ""},
		{Name: "sort", Type: field.TypeUint32, Default: 0},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_code",
				Unique:  true,
				Columns: []*schema.Column{SysRolesColumns[7]},
			},
		},
	}
	// SysTokensColumns holds the columns for the "sys_tokens" table.
	SysTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "token", Type: field.TypeString},
		{Name: "source", Type: field.TypeString},
		{Name: "expired_at", Type: field.TypeTime},
	}
	// SysTokensTable holds the schema information for the "sys_tokens" table.
	SysTokensTable = &schema.Table{
		Name:       "sys_tokens",
		Columns:    SysTokensColumns,
		PrimaryKey: []*schema.Column{SysTokensColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "token_uuid",
				Unique:  false,
				Columns: []*schema.Column{SysTokensColumns[4]},
			},
			{
				Name:    "token_expired_at",
				Unique:  false,
				Columns: []*schema.Column{SysTokensColumns[7]},
			},
		},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "home_path", Type: field.TypeString, Default: "/dashboard"},
		{Name: "mobile", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Default: "", SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "department_id", Type: field.TypeUint64, Nullable: true, Default: 1},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:       "sys_users",
		Columns:    SysUsersColumns,
		PrimaryKey: []*schema.Column{SysUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_users_sys_departments_departments",
				Columns:    []*schema.Column{SysUsersColumns[12]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[4], SysUsersColumns[10]},
			},
		},
	}
	// RoleMenusColumns holds the columns for the "role_menus" table.
	RoleMenusColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeUint64},
		{Name: "menu_id", Type: field.TypeUint64},
	}
	// RoleMenusTable holds the schema information for the "role_menus" table.
	RoleMenusTable = &schema.Table{
		Name:       "role_menus",
		Columns:    RoleMenusColumns,
		PrimaryKey: []*schema.Column{RoleMenusColumns[0], RoleMenusColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_menus_role_id",
				Columns:    []*schema.Column{RoleMenusColumns[0]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_menus_menu_id",
				Columns:    []*schema.Column{RoleMenusColumns[1]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPositionsColumns holds the columns for the "user_positions" table.
	UserPositionsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "position_id", Type: field.TypeUint64},
	}
	// UserPositionsTable holds the schema information for the "user_positions" table.
	UserPositionsTable = &schema.Table{
		Name:       "user_positions",
		Columns:    UserPositionsColumns,
		PrimaryKey: []*schema.Column{UserPositionsColumns[0], UserPositionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_positions_user_id",
				Columns:    []*schema.Column{UserPositionsColumns[0]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_positions_position_id",
				Columns:    []*schema.Column{UserPositionsColumns[1]},
				RefColumns: []*schema.Column{SysPositionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUint64},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysApisTable,
		SysAuditsTable,
		SysDepartmentsTable,
		SysDictionariesTable,
		SysDictionaryDetailsTable,
		SysMenusTable,
		SysMenuParamsTable,
		SysOauthProvidersTable,
		SysPositionsTable,
		SysRolesTable,
		SysTokensTable,
		SysUsersTable,
		RoleMenusTable,
		UserPositionsTable,
		UserRolesTable,
	}
)

func init() {
	SysApisTable.Annotation = &entsql.Annotation{
		Table: "sys_apis",
	}
	SysAuditsTable.Annotation = &entsql.Annotation{
		Table: "sys_audits",
	}
	SysDepartmentsTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysDepartmentsTable.Annotation = &entsql.Annotation{
		Table: "sys_departments",
	}
	SysDictionariesTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionaries",
	}
	SysDictionaryDetailsTable.ForeignKeys[0].RefTable = SysDictionariesTable
	SysDictionaryDetailsTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionary_details",
	}
	SysMenusTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenusTable.Annotation = &entsql.Annotation{
		Table: "sys_menus",
	}
	SysMenuParamsTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenuParamsTable.Annotation = &entsql.Annotation{
		Table: "sys_menu_params",
	}
	SysOauthProvidersTable.Annotation = &entsql.Annotation{
		Table: "sys_oauth_providers",
	}
	SysPositionsTable.Annotation = &entsql.Annotation{
		Table: "sys_positions",
	}
	SysRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_roles",
	}
	SysTokensTable.Annotation = &entsql.Annotation{
		Table: "sys_tokens",
	}
	SysUsersTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysUsersTable.Annotation = &entsql.Annotation{
		Table: "sys_users",
	}
	RoleMenusTable.ForeignKeys[0].RefTable = SysRolesTable
	RoleMenusTable.ForeignKeys[1].RefTable = SysMenusTable
	UserPositionsTable.ForeignKeys[0].RefTable = SysUsersTable
	UserPositionsTable.ForeignKeys[1].RefTable = SysPositionsTable
	UserRolesTable.ForeignKeys[0].RefTable = SysUsersTable
	UserRolesTable.ForeignKeys[1].RefTable = SysRolesTable
}
