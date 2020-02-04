package client

import (
        "fmt"
        "github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateMaintP(name string, description string, maintMaintPAttr models.MaintPAttributes) (*models.MaintP, error) {
        rn := fmt.Sprintf("maintpol-%s",name)
        parentDn := fmt.Sprintf("uni/fabric")
        maintMaintP := models.NewMaintP(rn, parentDn, description, maintMaintPAttr)
        err := sm.Save(maintMaintP)
        return maintMaintP, err
}

func (sm *ServiceManager) ReadMaintP(name string) (*models.MaintP, error) {
        dn := fmt.Sprintf("uni/fabric/maintpol-%s", name )
        cont, err := sm.Get(dn)
        if err != nil {
                return nil, err
        }
        maintMaintP := models.MaintPFromContainer(cont)
        return maintMaintP, nil
}

func (sm *ServiceManager) DeleteMaintP(name string ) error {
        dn := fmt.Sprintf("uni/fabric/maintpol-%s", name)
        return sm.DeleteByDn(dn, models.MaintMaintPClassName)
}

func (sm *ServiceManager) UpdateMaintP(name string, description string, maintMaintPAttr models.MaintPAttributes) (*models.MaintP, error) {
        rn := fmt.Sprintf("maintpol-%s",name)
        parentDn := fmt.Sprintf("uni/fabric")
        maintMaintP := models.NewMaintP(rn, parentDn, description, maintMaintPAttr)
        maintMaintP.Status = "modified"
        err := sm.Save(maintMaintP)
        return maintMaintP, err

}

func (sm *ServiceManager) ListMaintP() ([]*models.MaintP, error) {
        baseurlStr := "/api/node/class"
        dnUrl := fmt.Sprintf("%s/maintMaintP.json", baseurlStr )
        cont, err := sm.GetViaURL(dnUrl)
        list := models.MaintPListFromContainer(cont)
        return list, err
}

