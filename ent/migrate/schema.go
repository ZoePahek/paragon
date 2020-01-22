// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/kcarretto/paragon/ent/credential"
	"github.com/kcarretto/paragon/ent/file"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// CredentialsColumns holds the columns for the "credentials" table.
	CredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "principal", Type: field.TypeString},
		{Name: "secret", Type: field.TypeString},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"password", "key", "certificate"}},
		{Name: "fails", Type: field.TypeInt, Default: credential.DefaultFails},
		{Name: "target_credential_id", Type: field.TypeInt, Nullable: true},
	}
	// CredentialsTable holds the schema information for the "credentials" table.
	CredentialsTable = &schema.Table{
		Name:       "credentials",
		Columns:    CredentialsColumns,
		PrimaryKey: []*schema.Column{CredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "credentials_targets_credentials",
				Columns: []*schema.Column{CredentialsColumns[5]},

				RefColumns: []*schema.Column{TargetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "creation_time", Type: field.TypeTime},
		{Name: "last_modified_time", Type: field.TypeTime},
		{Name: "size", Type: field.TypeInt, Default: file.DefaultSize},
		{Name: "content", Type: field.TypeBytes},
		{Name: "hash", Type: field.TypeString},
		{Name: "content_type", Type: field.TypeString},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:        "files",
		Columns:     FilesColumns,
		PrimaryKey:  []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// JobsColumns holds the columns for the "jobs" table.
	JobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "creation_time", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
		{Name: "prev_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// JobsTable holds the schema information for the "jobs" table.
	JobsTable = &schema.Table{
		Name:       "jobs",
		Columns:    JobsColumns,
		PrimaryKey: []*schema.Column{JobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "jobs_jobs_next",
				Columns: []*schema.Column{JobsColumns[4]},

				RefColumns: []*schema.Column{JobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:        "tags",
		Columns:     TagsColumns,
		PrimaryKey:  []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TargetsColumns holds the columns for the "targets" table.
	TargetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "primary_ip", Type: field.TypeString},
		{Name: "machine_uuid", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "public_ip", Type: field.TypeString, Nullable: true},
		{Name: "primary_mac", Type: field.TypeString, Nullable: true},
		{Name: "hostname", Type: field.TypeString, Nullable: true},
		{Name: "last_seen", Type: field.TypeTime, Nullable: true},
	}
	// TargetsTable holds the schema information for the "targets" table.
	TargetsTable = &schema.Table{
		Name:        "targets",
		Columns:     TargetsColumns,
		PrimaryKey:  []*schema.Column{TargetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "queue_time", Type: field.TypeTime},
		{Name: "last_changed_time", Type: field.TypeTime},
		{Name: "claim_time", Type: field.TypeTime, Nullable: true},
		{Name: "exec_start_time", Type: field.TypeTime, Nullable: true},
		{Name: "exec_stop_time", Type: field.TypeTime, Nullable: true},
		{Name: "content", Type: field.TypeString},
		{Name: "output", Type: field.TypeString, Nullable: true},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "session_id", Type: field.TypeString, Nullable: true},
		{Name: "job_id", Type: field.TypeInt, Nullable: true},
		{Name: "target_id", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "tasks_jobs_tasks",
				Columns: []*schema.Column{TasksColumns[10]},

				RefColumns: []*schema.Column{JobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "tasks_targets_tasks",
				Columns: []*schema.Column{TasksColumns[11]},

				RefColumns: []*schema.Column{TargetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JobTagsColumns holds the columns for the "job_tags" table.
	JobTagsColumns = []*schema.Column{
		{Name: "job_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// JobTagsTable holds the schema information for the "job_tags" table.
	JobTagsTable = &schema.Table{
		Name:       "job_tags",
		Columns:    JobTagsColumns,
		PrimaryKey: []*schema.Column{JobTagsColumns[0], JobTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "job_tags_job_id",
				Columns: []*schema.Column{JobTagsColumns[0]},

				RefColumns: []*schema.Column{JobsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "job_tags_tag_id",
				Columns: []*schema.Column{JobTagsColumns[1]},

				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TargetTagsColumns holds the columns for the "target_tags" table.
	TargetTagsColumns = []*schema.Column{
		{Name: "target_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// TargetTagsTable holds the schema information for the "target_tags" table.
	TargetTagsTable = &schema.Table{
		Name:       "target_tags",
		Columns:    TargetTagsColumns,
		PrimaryKey: []*schema.Column{TargetTagsColumns[0], TargetTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "target_tags_target_id",
				Columns: []*schema.Column{TargetTagsColumns[0]},

				RefColumns: []*schema.Column{TargetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "target_tags_tag_id",
				Columns: []*schema.Column{TargetTagsColumns[1]},

				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TaskTagsColumns holds the columns for the "task_tags" table.
	TaskTagsColumns = []*schema.Column{
		{Name: "task_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// TaskTagsTable holds the schema information for the "task_tags" table.
	TaskTagsTable = &schema.Table{
		Name:       "task_tags",
		Columns:    TaskTagsColumns,
		PrimaryKey: []*schema.Column{TaskTagsColumns[0], TaskTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "task_tags_task_id",
				Columns: []*schema.Column{TaskTagsColumns[0]},

				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "task_tags_tag_id",
				Columns: []*schema.Column{TaskTagsColumns[1]},

				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CredentialsTable,
		FilesTable,
		JobsTable,
		TagsTable,
		TargetsTable,
		TasksTable,
		JobTagsTable,
		TargetTagsTable,
		TaskTagsTable,
	}
)

func init() {
	CredentialsTable.ForeignKeys[0].RefTable = TargetsTable
	JobsTable.ForeignKeys[0].RefTable = JobsTable
	TasksTable.ForeignKeys[0].RefTable = JobsTable
	TasksTable.ForeignKeys[1].RefTable = TargetsTable
	JobTagsTable.ForeignKeys[0].RefTable = JobsTable
	JobTagsTable.ForeignKeys[1].RefTable = TagsTable
	TargetTagsTable.ForeignKeys[0].RefTable = TargetsTable
	TargetTagsTable.ForeignKeys[1].RefTable = TagsTable
	TaskTagsTable.ForeignKeys[0].RefTable = TasksTable
	TaskTagsTable.ForeignKeys[1].RefTable = TagsTable
}
