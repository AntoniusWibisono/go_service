package comment

type Organization struct {
	Name string `json:"name"`
}

type Member struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatarUrl"`
	Followers int64  `json:"followers"`
	Following int64  `json:"following"`
}

type Comment struct {
	Comment string `json:"comment"`
}

type Request struct {
	SearchBy    string `json:"searchBy" validate:"required_with=SearchValue"`
	SearchValue string `json:"searchValue" validate:"required_with=SearchBy"`
	SortBy      string `json:"sortBy" validate:"required_with=SortType"`
	SortType    string `json:"sortType" validate:"required_with=SortBy,omitempty,oneof=asc desc"`
	Page        int64  `json:"page"`
	PerPage     int64  `json:"perPage"`
}

type ListCommentResponse struct {
	Comments []Comment `json:"comments"`
	Page     int64     `json:"page"`
	PerPage  int64     `json:"perPage"`
	Count    int64     `json:"count"`
}

type CommentRequest struct {
	OrganizationName string `json:"organizationName"`
}

type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}

type PostCommentRequest struct {
	OrganizationName string `json:"organizationName"`
	MemberId         string `json:"memberId"`
	Comment          string `json:"comment"`
}
