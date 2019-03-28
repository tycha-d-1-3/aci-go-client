package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateCloudContextProfile(name string, tenant string, description string, cloudCtxProfileattr models.CloudContextProfileAttributes, primaryCidr string, region string) (*models.CloudContextProfile, error) {
	rn := fmt.Sprintf("ctxprofile-%s", name)
	parentDn := tenant
	cloudCtxProfile := models.NewCloudContextProfile(rn, parentDn, description, cloudCtxProfileattr)
	jsonPayload, _, err := sm.PrepareModel(cloudCtxProfile)

	cidrJSON := []byte(fmt.Sprintf(`
	{
		"cloudCidr": {
			"attributes": {
				"addr": "%s",
				"primary": "yes",
				"status": "created,modified"
			}
		}
	}
	`, primaryCidr))

	regionAttach := []byte(fmt.Sprintf(`
	{
		"cloudRsCtxProfileToRegion": {
			"attributes": {
				"status": "created,modified",
				"tDn": "uni/clouddomp/provp-aws/region-%s"
			}
		}
	}
	`, region))

	cidrCon, err := container.ParseJSON(cidrJSON)
	regionCon, err := container.ParseJSON(regionAttach)
	if err != nil {
		fmt.Println("error occured in parsing")
		fmt.Println(err)
	}
	jsonPayload.Array(cloudCtxProfile.ClassName, "children")
	jsonPayload.ArrayAppend(cidrCon.Data(), cloudCtxProfile.ClassName, "children")
	jsonPayload.ArrayAppend(regionCon.Data(), cloudCtxProfile.ClassName, "children")
	fmt.Println(jsonPayload)
	jsonPayload.Set(name, cloudCtxProfile.ClassName, "attributes", "name")
	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("/api/node/mo/%s/%s.json", parentDn, rn), jsonPayload, true)
	if err != nil {
		return nil, err
	}
	cont, _, err := sm.client.Do(req)
	if err != nil {
		return nil, err
	}

	//err := sm.Save(cloudCtxProfile)
	return cloudCtxProfile, CheckForErrors(cont, "POST")
}

func (sm *ServiceManager) ReadCloudContextProfile(name string, tenant string) (*models.CloudContextProfile, error) {
	dn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s", tenant, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudCtxProfile := models.CloudContextProfileFromContainer(cont)
	return cloudCtxProfile, nil
}

func (sm *ServiceManager) DeleteCloudContextProfile(name string, tenant string) error {
	dn := fmt.Sprintf("uni/tn-%s/ctxprofile-%s", tenant, name)
	return sm.DeleteByDn(dn, models.CloudctxprofileClassName)
}

func (sm *ServiceManager) UpdateCloudContextProfile(name string, tenant string, description string, cloudCtxProfileattr models.CloudContextProfileAttributes, primaryCidr string, region string) (*models.CloudContextProfile, error) {
	rn := fmt.Sprintf("ctxprofile-%s", name)
	parentDn := tenant
	cloudCtxProfile := models.NewCloudContextProfile(rn, parentDn, description, cloudCtxProfileattr)
	cloudCtxProfile.Status = "modified"
	jsonPayload, _, err := sm.PrepareModel(cloudCtxProfile)

	cidrJSON := []byte(fmt.Sprintf(`
	{
		"cloudCidr": {
			"attributes": {
				"addr": "%s",
				"primary": "yes",
				"status": "modified"
			}
		}
	}
	`, primaryCidr))

	regionAttach := []byte(fmt.Sprintf(`
	{
		"cloudRsCtxProfileToRegion": {
			"attributes": {
				"status": "modified",
				"tDn": "uni/clouddomp/provp-aws/region-%s"
			}
		}
	}
	`, region))

	cidrCon, err := container.ParseJSON(cidrJSON)
	regionCon, err := container.ParseJSON(regionAttach)

	jsonPayload.Array(cloudCtxProfile.ClassName, "children")
	jsonPayload.ArrayAppend(cidrCon.Data(), cloudCtxProfile.ClassName, "children")
	jsonPayload.ArrayAppend(regionCon.Data(), cloudCtxProfile.ClassName, "children")
	fmt.Println(jsonPayload)
	jsonPayload.Set(name, cloudCtxProfile.ClassName, "attributes", "name")
	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("/api/node/mo/%s/%s.json", parentDn, rn), jsonPayload, true)
	if err != nil {
		return nil, err
	}
	fmt.Println("*******payload*******", jsonPayload.String())
	cont, _, err := sm.client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", cont)

	return cloudCtxProfile, CheckForErrors(cont, "POST")

}

func (sm *ServiceManager) ListCloudContextProfile(tenant string) ([]*models.CloudContextProfile, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudCtxProfile.json", baseurlStr, tenant)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudContextProfileListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationcloudRsCtxToFlowLogFromCloudContextProfile(parentDn, tnCloudAwsFlowLogPolName string) error {
	dn := fmt.Sprintf("%s/rsctxToFlowLog", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnCloudAwsFlowLogPolName": "%s"
								
			}
		}
	}`, "cloudRsCtxToFlowLog", dn, tnCloudAwsFlowLogPolName))

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

func (sm *ServiceManager) DeleteRelationcloudRsCtxToFlowLogFromCloudContextProfile(parentDn string) error {
	dn := fmt.Sprintf("%s/rsctxToFlowLog", parentDn)
	return sm.DeleteByDn(dn, "cloudRsCtxToFlowLog")
}
func (sm *ServiceManager) CreateRelationcloudRsToCtxFromCloudContextProfile(parentDn, tnFvCtxName string) error {
	dn := fmt.Sprintf("%s/rstoCtx", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvCtxName": "%s"
								
			}
		}
	}`, "cloudRsToCtx", dn, tnFvCtxName))

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
func (sm *ServiceManager) CreateRelationcloudRsCtxProfileToRegionFromCloudContextProfile(parentDn, tnCloudRegionName string) error {
	dn := fmt.Sprintf("%s/rsctxProfileToRegion", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnCloudRegionName": "%s"
								
			}
		}
	}`, "cloudRsCtxProfileToRegion", dn, tnCloudRegionName))

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

func (sm *ServiceManager) DeleteRelationcloudRsCtxProfileToRegionFromCloudContextProfile(parentDn string) error {
	dn := fmt.Sprintf("%s/rsctxProfileToRegion", parentDn)
	return sm.DeleteByDn(dn, "cloudRsCtxProfileToRegion")
}