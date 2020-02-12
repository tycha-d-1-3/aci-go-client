package models


import (
    "fmt"
    "strconv"

    "github.com/ciscoecosystem/aci-go-client/container"
)

const MaintMaintGrpClassName = "maintMaintGrp"

type MaintGrp struct {
    BaseAttributes
    MaintGrpAttributes
}

type MaintGrpAttributes struct {
    Annotation      string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    Fwtype          string `json:",omitempty"`
    Type            string `json:",omitempty"`

}


func NewMaintGrp(maintMaintGrpRn, parentDn, description string, maintMaintGrpattr MaintGrpAttributes) *MaintGrp {
    dn := fmt.Sprintf("%s/%s", parentDn, maintMaintGrpRn)
    return &MaintGrp{
        BaseAttributes: BaseAttributes{
            DistinguishedName: dn,
            Description:       description,
            Status:            "",
            ClassName:         MaintMaintGrpClassName,
            Rn:                maintMaintGrpRn,
        },

        MaintGrpAttributes: maintMaintGrpattr,

    }
}

func (maintMaintGrp *MaintGrp) ToMap() (map[string]string, error) {
    maintMaintGrpMap, err := maintMaintGrp.BaseAttributes.ToMap()
    if err != nil {
        return nil, err
    }

    A(maintMaintGrpMap, "annotation",maintMaintGrp.Annotation)
    A(maintMaintGrpMap, "nameAlias",maintMaintGrp.NameAlias)
    A(maintMaintGrpMap, "fwtype",maintMaintGrp.Fwtype)
    A(maintMaintGrpMap, "type",maintMaintGrp.Type)

    return maintMaintGrpMap, err
}

func MaintGrpFromContainerList(cont *container.Container, index int) *MaintGrp {

    MaintGrpCont := cont.S("imdata").Index(index).S(MaintMaintGrpClassName, "attributes")
    return &MaintGrp{
        BaseAttributes{
            DistinguishedName: G(MaintGrpCont, "dn"),
            Description:       G(MaintGrpCont, "descr"),
            Status:            G(MaintGrpCont, "status"),
            ClassName:         MaintMaintGrpClassName,
            Rn:                G(MaintGrpCont, "rn"),
        },

        MaintGrpAttributes{
        Annotation : G(MaintGrpCont, "annotation"),
        NameAlias : G(MaintGrpCont, "nameAlias"),
        Fwtype : G(MaintGrpCont, "fwtype"),
        Type : G(MaintGrpCont, "type"),


        },

    }
}

func MaintGrpFromContainer(cont *container.Container) *MaintGrp {

    return MaintGrpFromContainerList(cont, 0)
}

func MaintGrpListFromContainer(cont *container.Container) []*MaintGrp {
    length, _ := strconv.Atoi(G(cont, "totalCount"))

    arr := make([]*MaintGrp, length)

    for i := 0; i < length; i++ {

        arr[i] = MaintGrpFromContainerList(cont, i)
    }

    return arr
}
