package models

type (
	StudentResponse struct {
		StudentID       int64  `json:"student_id"`
		FirstName       string `json:"first_name"`
		SecondName      string `json:"second_name"`
		Gender          string `json:"gender"`
		GroupNumber     string `json:"group_number"`
		GraduationYear  int16  `json:"graduation_year"`
		ExamScore       int16  `json:"exam_score"`
		AdditionalScore int8   `json:"additional_score"`
	}

	StudentResponseOptional struct {
		StudentID       *int64  `json:"student_id,omitempty"`
		FirstName       *string `json:"first_name,omitempty"`
		SecondName      *string `json:"second_name,omitempty"`
		Gender          *string `json:"gender,omitempty"`
		GroupNumber     *string `json:"group_number,omitempty"`
		GraduationYear  *int16  `json:"graduation_year,omitempty"`
		ExamScore       *int16  `json:"exam_score,omitempty"`
		AdditionalScore *int8   `json:"additional_score,omitempty"`
	}

	StudentInsert struct {
		FirstName       string `json:"first_name" binding:"required"`
		SecondName      string `json:"second_name" binding:"required"`
		Gender          string `json:"gender" binding:"required"`
		GroupNumber     string `json:"group_number" binding:"required"`
		GraduationYear  int16  `json:"graduation_year" binding:"required"`
		ExamScore       int16  `json:"exam_score" binding:"required"`
		AdditionalScore int8   `json:"additional_score" binding:"required"`
	}

	StudentUpdate struct {
		StudentID int64 `json:"student_id" binding:"required"`
		StudentInsert
	}

	StudentPageRequest struct {
		PageSize int64
		PageNum  int64
	}
)
