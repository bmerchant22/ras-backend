package rc

import (
	"gorm.io/gorm"
)

type RecruitmentCycleType string

const (
	PLACEMENT  RecruitmentCycleType = "Placement"
	INTERNSHIP RecruitmentCycleType = "Internship"
)

type RecruitmentCycle struct {
	gorm.Model
	IsActive            bool                 `json:"is_active" gorm:"default:true"`
	AcademicYear        string               `json:"academic_year"`
	Type                RecruitmentCycleType `json:"type"`
	StartDate           int64                `json:"start_date"`
	Phase               string               `json:"phase"`
	ApplicationCountCap uint                 `json:"application_count_cap"`
}

type RecruitmentCycleQuestionsType string

const (
	MCQ         RecruitmentCycleQuestionsType = "MCQ"
	SHORTANSWER RecruitmentCycleQuestionsType = "ShortAnswer"
	BOOLEAN     RecruitmentCycleQuestionsType = "Boolean"
)

type RecruitmentCycleQuestion struct {
	gorm.Model
	Type               RecruitmentCycleQuestionsType `json:"type"`
	Question           string                        `json:"question"`
	RecruitmentCycleID uint                          `gorm:"uniqueIndex;->;<-:create" json:"recruitment_cycle_id"`
	RecruitmentCycle   RecruitmentCycle              `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	Mandatory          bool                          `json:"mandatory" gorm:"default:false"`
	Options            string                        `json:"options"` //csv
}

type RecruitmentCycleQuestionsAnswer struct {
	gorm.Model
	RecruitmentCycleQuestionID uint                     `gorm:"index" json:"recruitment_cycle_question_id"`
	RecruitmentCycleQuestion   RecruitmentCycleQuestion `gorm:"foreignkey:RecruitmentCycleQuestionID" json:"-"`
	StudentRecruitmentCycleID  uint                     `gorm:"uniqueIndex;->;<-:create" json:"student_recruitment_cycle_id"`
	StudentRecruitmentCycle    StudentRecruitmentCycle  `gorm:"foreignkey:StudentRecruitmentCycleID" json:"-"`
	Answer                     string                   `json:"answer"`
	Comments                   string                   `json:"comments"`
	Status                     string                   `json:"status"`
}

type CompanyRecruitmentCycle struct {
	gorm.Model
	CompanyID          uint             `gorm:"index" json:"company_id"`
	CompanyName        string           `json:"company_name"`
	RecruitmentCycleID uint             `gorm:"index" json:"recruitment_cycle_id"`
	RecruitmentCycle   RecruitmentCycle `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	HR1                string           `json:"hr1"`
	HR2                string           `json:"hr2"`
	HR3                string           `json:"hr3"`
	Comments           string           `json:"comments"`
}

type Notice struct {
	gorm.Model
	RecruitmentCycleID uint             `gorm:"index" json:"recruitment_cycle_id"`
	RecruitmentCycle   RecruitmentCycle `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	Title              string           `json:"title" binding:"required"`
	Description        string           `json:"description" binding:"required"`
	Tags               string           `json:"tags" binding:"required"`
	Attachment         string           `json:"attachment"`
	CreatedBy          string           `json:"created_by"`
	LastReminderAt     int64            `json:"last_reminder_at" gorm:"default:0"`
}

type StudentRecruitmentCycleType string

const (
	PIOPPO    StudentRecruitmentCycleType = "PIO-PPO"
	RECRUITED StudentRecruitmentCycleType = "Recruited"
	AVAILABLE StudentRecruitmentCycleType = "Available"
)

type StudentRecruitmentCycle struct {
	gorm.Model
	StudentID                    uint                        `gorm:"index" json:"student_id"`
	RecruitmentCycleID           uint                        `gorm:"index" json:"recruitment_cycle_id" binding:"required"`
	RecruitmentCycle             RecruitmentCycle            `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	ProgramDepartmentID          uint                        `gorm:"index" json:"program_department_id"`
	SecondaryProgramDepartmentID uint                        `gorm:"index" json:"secondary_program_department_id"`
	CPI                          float64                     `json:"cpi"`
	Email                        string                      `grom:"index" json:"email"`
	Name                         string                      `json:"name"`
	Type                         StudentRecruitmentCycleType `json:"type" gorm:"default:Available"`
	IsFrozen                     bool                        `json:"is_frozen" gorm:"default:false"`
	IsVerified                   bool                        `json:"is_verified" gorm:"default:false"`
	Comment                      string                      `json:"comment"`
}
