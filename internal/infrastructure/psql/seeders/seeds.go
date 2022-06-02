package seeders

import "go_service/internal/domain"

func CreateSeed() (members []*domain.Members, organizations []*domain.Organizations, comments []*domain.Comments) {

	zendit := &domain.Organizations{
		ID:   "b15bbbf9-4213-4451-93ca-a81fba03f6b0",
		Name: "Zendit",
	}

	seiko := &domain.Organizations{
		ID:   "0b174a6c-2847-4060-a53a-a83ee2393ed3",
		Name: "Seiko",
	}

	organizations = append(organizations, zendit)
	organizations = append(organizations, seiko)

	memberSatu := &domain.Members{
		ID:             "53432b76-9ee6-4507-86f1-53af94b8bec4",
		OrganizationId: zendit.ID,
		Username:       "Anton Wibisono",
		Password:       "12345",
		AvatarUrl:      "www.google.com/pic1.jpg",
		Followers:      100,
		Following:      5,
	}

	memberDua := &domain.Members{
		ID:             "3ca71c95-8fe2-4b18-8fde-54212c474ab7",
		OrganizationId: zendit.ID,
		Username:       "John Doe",
		Password:       "12345",
		AvatarUrl:      "www.google.com/pic2.jpg",
		Followers:      75,
		Following:      15,
	}

	members = append(members, memberSatu)
	members = append(members, memberDua)

	commentSatu := &domain.Comments{
		ID:             "2cab18e2-48ef-4f3b-8a50-657aa1a0eecb",
		OrganizationId: zendit.ID,
		MemberId:       memberSatu.ID,
		Comment:        "Looking to hire SEA Developer!",
	}

	commentDua := &domain.Comments{
		ID:             "dc1c2a1d-e995-41fe-a6e1-df99dfef6440",
		OrganizationId: zendit.ID,
		MemberId:       memberDua.ID,
		Comment:        "Good Day !",
	}

	commentTiga := &domain.Comments{
		ID:             "dc1c2a1d-4c5b-41fe-a6e1-a81fba03f6b0",
		OrganizationId: seiko.ID,
		MemberId:       memberDua.ID,
		Comment:        "Comment for Seiko",
	}

	comments = append(comments, commentSatu)
	comments = append(comments, commentDua)
	comments = append(comments, commentTiga)
	return
}
