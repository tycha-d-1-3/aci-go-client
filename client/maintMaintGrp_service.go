package client
  
import (
        "fmt"
        "github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateMaintGrp(name string, description string, maintMaintGrpAttr models.MaintGrpAttributes) (*models.MaintGrp, error) {
        rn := fmt.Sprintf("maintgrp-%s",name)
        parentDn := fmt.Sprintf("uni/fabric")
        maintMaintGrp := models.NewMaintGrp(rn, parentDn, description, maintMaintGrpAttr)
        err := sm.Save(maintMaintGrp)
        return maintMaintGrp, err
}

func (sm *ServiceManager) ReadMaintGrp(name string) (*models.MaintGrp, error) {
        dn := fmt.Sprintf("uni/fabric/maintgrp-%s", name )
        cont, err := sm.Get(dn)
        if err != nil {
                return nil, err
        }
        maintMaintGrp := models.MaintGrpFromContainer(cont)
        return maintMaintGrp, nil
}

func (sm *ServiceManager) DeleteMaintGrp(name string ) error {
        dn := fmt.Sprintf("uni/fabric/maintgrp-%s", name)
        return sm.DeleteByDn(dn, models.MaintMaintGrpClassName)
}

func (sm *ServiceManager) UpdateMaintGrp(name string, description string, maintMaintGrpAttr models.MaintGrpAttributes) (*models.MaintGrp, error) {
        rn := fmt.Sprintf("maintgrp-%s",name)
        parentDn := fmt.Sprintf("uni/fabric")
        maintMaintGrp := models.NewMaintGrp(rn, parentDn, description, maintMaintGrpAttr)
        maintMaintGrp.Status = "modified"
        err := sm.Save(maintMaintGrp)
        return maintMaintGrp, err

}

func (sm *ServiceManager) ListMaintGrp() ([]*models.MaintGrp, error) {
        baseurlStr := "/api/node/class"
        dnUrl := fmt.Sprintf("%s/maintMaintGrp.json", baseurlStr )
        cont, err := sm.GetViaURL(dnUrl)
        list := models.MaintGrpListFromContainer(cont)
        return list, err
}

