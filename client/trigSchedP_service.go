package client

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateSchedP(name string, description string, trigSchedPAttr models.SchedPAttributes) (*models.SchedP, error) {
	rn := fmt.Sprintf("schedp-%s", name)
	parentDn := fmt.Sprintf("uni/fabric")
	trigSchedP := models.NewSchedP(rn, parentDn, description, trigSchedPAttr)
	err := sm.Save(trigSchedP)
	return trigSchedP, err
}

func (sm *ServiceManager) ReadSchedP(name string) (*models.SchedP, error) {
	dn := fmt.Sprintf("uni/fabric/schedp-%s", name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	trigSchedP := models.SchedPFromContainer(cont)
	return trigSchedP, nil
}

func (sm *ServiceManager) DeleteSchedP(name string) error {
	dn := fmt.Sprintf("uni/fabric/schedp-%s", name)
	return sm.DeleteByDn(dn, models.TrigSchedPClassName)
}

func (sm *ServiceManager) UpdateSchedP(name string, description string, trigSchedPAttr models.SchedPAttributes) (*models.SchedP, error) {
	rn := fmt.Sprintf("schedp-%s", name)
	parentDn := fmt.Sprintf("uni/fabric")
	trigSchedP := models.NewSchedP(rn, parentDn, description, trigSchedPAttr)
	trigSchedP.Status = "modified"
	err := sm.Save(trigSchedP)
	return trigSchedP, err

}

func (sm *ServiceManager) ListSchedP() ([]*models.SchedP, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/trigSchedP.json", baseurlStr)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.SchedPListFromContainer(cont)
	return list, err
}
