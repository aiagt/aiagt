package model

type Workflow struct {
	Base
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`

	Nodes []WorkflowNode `gorm:"-"`
}

type WorkflowOptional struct {
}
