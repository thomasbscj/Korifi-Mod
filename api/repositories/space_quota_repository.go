package repositories

import "time"

const(
	SpaceQuotaResourceType = "SpaceQuota"
)

type CreateSpaceQuotaMessage struct {
	Name string
	Relationship struct {
		organization struct { guid string } //org guid
		spaces struct { guids []string } //spaces guid
	}
	Apps struct { TotalMemoryMb int }
}

type ListSpaceQuotaMessage struct {
	OrgGuid string
}
type GetSpaceQuotaMessage struct {
	Guid string
}
type UpdateSpaceQuotaMessage struct {
	Name string
	Apps struct { TotalMemoryMb int }
}
type DeleteSpaceQuotaMessage struct {
	Guid string
}
type ApplySpaceQuotaMessage struct {
	Guid string
	Spaces struct { guids []string } 
}
type RemoveSpaceQuotaMessage struct {
	QuotaGuid string
	SpaceGuid string 
}

type SpaceQuotaRecord struct {
	Name string
	GUID string
	OrganizationGUID string
	SpacesGUID []string


	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type SpaceQuotaRepo struct {

}