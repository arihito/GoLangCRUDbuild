package request

type (
	CreateStudyGroupRequest struct {
		Title string
	}

	UpdateStudyGroupRequest struct {
		Title     string
		SubTitle  string
		ImagePath string
		PageUrl   string
		Published bool
	}
)
