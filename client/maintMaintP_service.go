package client

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateMaintP(name string, description string, maintMaintPAttr models.MaintPAttributes) (*models.MaintP, error) {
	rn := fmt.Sprintf("maintpol-%s", name)
	parentDn := fmt.Sprintf("uni/fabric")
	maintMaintP := models.NewMaintP(rn, parentDn, description, maintMaintPAttr)
	err := sm.Save(maintMaintP)
	return maintMaintP, err
}

func (sm *ServiceManager) ReadMaintP(name string) (*models.MaintP, error) {
	dn := fmt.Sprintf("uni/fabric/maintpol-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	maintMaintP := models.MaintPFromContainer(cont)
	return maintMaintP, nil
}

func (sm *ServiceManager) DeleteMaintP(name string) error {
	dn := fmt.Sprintf("uni/fabric/maintpol-%s", name)
	return sm.DeleteByDn(dn, models.MaintMaintPClassName)
}

func (sm *ServiceManager) UpdateMaintP(name string, description string, maintMaintPAttr models.MaintPAttributes) (*models.MaintP, error) {
	rn := fmt.Sprintf("maintpol-%s", name)
	parentDn := fmt.Sprintf("uni/fabric")
	maintMaintP := models.NewMaintP(rn, parentDn, description, maintMaintPAttr)
	maintMaintP.Status = "modified"
	err := sm.Save(maintMaintP)
	return maintMaintP, err

}

func (sm *ServiceManager) ListMaintP() ([]*models.MaintP, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/maintMaintP.json", baseurlStr)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.MaintPListFromContainer(cont)
	return list, err
}

func (sm *ServiceManager) CreateRelationmaintRsPolSchedulerFromMaintP(parentDn, tnTrigSchedPName string) error {
	dn := fmt.Sprintf("%s/rspolScheduler", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"tnTrigSchedPName": "%s"				
			}
		}
	}`, "maintRsPolScheduler", dn, tnTrigSchedPName))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) ReadRelationmaintRsPolSchedulerFromMaintP(parentDn string) (interface{}, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/%s/%s.json", baseurlStr, parentDn, "maintRsPolScheduler")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont, "maintRsPolScheduler")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tnTrigSchedPName")
		return dat, err
	} else {
		return nil, err
	}
}
