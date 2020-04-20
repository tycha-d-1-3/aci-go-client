package models


import (
        "fmt"
        "strconv"

        "github.com/ciscoecosystem/aci-go-client/container"
)

const FirmwareOSourceClassName = "firmwareOSource"

type OSource struct {
        BaseAttributes
    OSourceAttributes
}

type OSourceAttributes struct {
    Name       string `json:",omitempty"`
    Url       string `json:",omitempty"`
    Proto       string `json:",omitempty"`
    User       string `json:",omitempty"`
    AuthType       string `json:",omitempty"`
    AuthPass       string `json:",omitempty"`
}

func NewOSource(firmwareOSourceRn, parentDn, description string, firmwareOSourceAttr OSourceAttributes) *OSource {
        dn := fmt.Sprintf("%s/%s", parentDn, firmwareOSourceRn)
        return &OSource{
                BaseAttributes: BaseAttributes{
                        DistinguishedName: dn,
                        Description:       description,
                        Status:            "created, modified",
                        ClassName:         FirmwareOSourceClassName,
                        Rn:                firmwareOSourceRn,
                },
                OSourceAttributes: firmwareOSourceAttr,
        }
}

func (firmwareOSource *OSource) ToMap() (map[string]string, error) {
        firmwareOSourceMap, err := firmwareOSource.BaseAttributes.ToMap()
        if err != nil {
                return nil, err
        }

    A(firmwareOSourceMap, "name", firmwareOSource.Name)
    A(firmwareOSourceMap, "url", firmwareOSource.Url)
    A(firmwareOSourceMap, "proto", firmwareOSource.Proto)
    A(firmwareOSourceMap, "user", firmwareOSource.User)
    A(firmwareOSourceMap, "authType", firmwareOSource.AuthType)
    A(firmwareOSourceMap, "authPass", firmwareOSource.AuthPass)

    return firmwareOSourceMap, err
}

func OSourceFromContainerList(cont *container.Container, index int) *OSource {

        OSourceCont := cont.S("imdata").Index(index).S(FirmwareOSourceClassName, "attributes")
        return &OSource{
                BaseAttributes{
                        DistinguishedName: G(OSourceCont, "dn"),
                        Description:       G(OSourceCont, "descr"),
                        Status:            G(OSourceCont, "status"),
                        ClassName:         FirmwareOSourceClassName,
                        Rn:                G(OSourceCont, "rn"),
                },

                OSourceAttributes{
        Name : G(OSourceCont, "name"),
        Url : G(OSourceCont, "url"),

        Proto : G(OSourceCont, "proto"),
        User : G(OSourceCont, "user"),
        AuthType : G(OSourceCont, "authType"),
        AuthPass : G(OSourceCont, "AuthPass"),
        },

        }
}


func OSourceFromContainer(cont *container.Container) *OSource {

        return OSourceFromContainerList(cont, 0)
}

func OSourceListFromContainer(cont *container.Container) []*OSource {
        length, _ := strconv.Atoi(G(cont, "totalCount"))

        arr := make([]*OSource, length)

        for i := 0; i < length; i++ {

                arr[i] = OSourceFromContainerList(cont, i)
        }

        return arr
}

