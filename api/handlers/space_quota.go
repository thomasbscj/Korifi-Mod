package handlers

import (
	"code.cloudfoundry.org/korifi/api/repositories"
	"code.cloudfoundry.org/korifi/api/repositories/include"
	"code.cloudfoundry.org/korifi/api/routing"
)

type CFSpaceQuotaRepository interface{
	CreateSpaceQuota()()
	GetSpaceQuota()()
	ListSpaceQuota()()
	UpdateSpaceQuota()()
	DeleteSpaceQuota()()
	ApplySpaceQuotaToSpace()()
	RemoveSpaceQuotaFromSpace()()
}

type SpaceQuota struct {
	spaceQuotaRepo   CFSpaceQuotaRepository
	spaceRepo        CFSpaceRepository
	requestValidator RequestValidator
	includeResolver  *include.IncludeResolver[
		[]repositories.SpaceRecord,
		repositories.SpaceRecord,
	]
}

func (h *SpaceQuota)create()(){}
func (h *SpaceQuota)list()(){}
func (h *SpaceQuota)get()(){}
func (h *SpaceQuota)update()(){}
func (h *SpaceQuota)delete()(){}
func (h *SpaceQuota)apply()(){}
func (h *SpaceQuota)remove()(){}

const(
	SpaceQuotasPath = "/v3/space_quotas"
	SpaceQuotaPath = "v3/space_quotas/{guid}" //quota guid
	SpaceQuotasWithRelationshipSpace = "/v3/space_quotas/{guid}/relationships/spaces" // quota guid
	SpaceQuotaWithRelationshipSpace = "/v3/space_quotas/{guid}/relationships/spaces/{guid}" //quota_guid - space guid
)

func (h *SpaceQuota) AuthenticatedRoutes() []routing.Route {
	return []routing.Route{
		{Method: "POST", Pattern: SpaceQuotasPath, Handler: h.create},
		{Method: "GET", Pattern: SpaceQuotasPath, Handler: h.list},
		{Method: "GET", Pattern: SpaceQuotaPath, Handler: h.get},
		{Method: "PATCH", Pattern: SpaceQuotaPath, Handler: h.update},
		{Method: "DELETE", Pattern: SpaceQuotaPath, Handler: h.delete},
		{Method: "POST", Pattern: SpaceQuotasWithRelationshipSpace, Handler: h.apply},
		{Method: "DELETE", Pattern: SpaceQuotaWithRelationshipSpace, Handler: h.remove},
	}
}