package models


import (
    "fmt"
    "strconv"

    "github.com/ciscoecosystem/aci-go-client/container"
)

const MaintMaintPClassName = "maintMaintP"

type MaintP struct {
    BaseAttributes
    MaintPAttributes
}

type MaintPAttributes struct {
    AdminSt string `json:",omitempty"`
    Annotation string `json:",omitempty"`
    Graceful string `json:",omitempty"`
    IgnoreCompat string `json:",omitempty"`
    InternalLabel string `json:",omitempty"`
    InternalSource string `json:",omitempty"`
    ModTs string `json:",omitempty"`
    Name string `json:",omitempty"`
    NameAlias string `json:",omitempty"`
    NotifCond string `json:",omitempty"`
    RunMode string `json:",omitempty"`
    SeqNum string `json:",omitempty"`
    SrUpgrade string `json:",omitempty"`
    SrVersion string `json:",omitempty"`
    TriggerTime string `json:",omitempty"`
    Version string `json:",omitempty"`
    VersionCheckOverride string `json:",omitempty"`
}

func NewMaintP(maintMaintPRn, parentDn, description string, maintMaintPattr MaintPAttributes) *MaintP {
    dn := fmt.Sprintf("%s/%s", parentDn, maintMaintPRn)
    return &MaintP{
        BaseAttributes: BaseAttributes{
            DistinguishedName: dn,
            Description:       description,
            Status:            "",
            ClassName:         MaintMaintPClassName,
            Rn:                maintMaintPRn,
        },

        MaintPAttributes: maintMaintPattr,

    }
}

func (maintMaintP *MaintP) ToMap() (map[string]string, error) {
    maintMaintPMap, err := maintMaintP.BaseAttributes.ToMap()
    if err != nil {
        return nil, err
    }

    A(maintMaintPMap, "adminSt", maintMaintP.AdminSt)
    A(maintMaintPMap, "annotation", maintMaintP.Annotation)
    A(maintMaintPMap, "graceful", maintMaintP.Graceful)
    A(maintMaintPMap, "ignoreCompat", maintMaintP.IgnoreCompat)
    A(maintMaintPMap, "internalLabel", maintMaintP.InternalLabel)
    A(maintMaintPMap, "internalSource", maintMaintP.InternalSource)
    A(maintMaintPMap, "modTs", maintMaintP.ModTs)
    A(maintMaintPMap, "name", maintMaintP.Name)
    A(maintMaintPMap, "nameAlias", maintMaintP.NameAlias)
    A(maintMaintPMap, "notifCond", maintMaintP.NotifCond)
    A(maintMaintPMap, "runMode", maintMaintP.RunMode)
    A(maintMaintPMap, "seqNum", maintMaintP.SeqNum)
    A(maintMaintPMap, "srUpgrade", maintMaintP.SrUpgrade)
    A(maintMaintPMap, "srVersion", maintMaintP.SrVersion)
    A(maintMaintPMap, "triggerTime", maintMaintP.TriggerTime)
    A(maintMaintPMap, "version", maintMaintP.Version)
    A(maintMaintPMap, "versionCheckOverride", maintMaintP.VersionCheckOverride)

    return maintMaintPMap, err
}

func MaintPFromContainerList(cont *container.Container, index int) *MaintP {

    MaintPCont := cont.S("imdata").Index(index).S(MaintMaintPClassName, "attributes")
    return &MaintP{
        BaseAttributes{
            DistinguishedName: G(MaintPCont, "dn"),
            Description:       G(MaintPCont, "descr"),
            Status:            G(MaintPCont, "status"),
            ClassName:         MaintMaintPClassName,
            Rn:                G(MaintPCont, "rn"),
        },

        MaintPAttributes{
            AdminSt : G(MaintPCont,"adminSt"),
            Annotation : G(MaintPCont,"annotation"),
            Graceful : G(MaintPCont,"graceful"),
            IgnoreCompat : G(MaintPCont,"ignoreCompat"),
            InternalLabel : G(MaintPCont,"internalLabel"),
            InternalSource : G(MaintPCont,"internalSource"),
            ModTs : G(MaintPCont,"modTs"),
            Name : G(MaintPCont,"name"),
            NameAlias : G(MaintPCont,"nameAlias"),
            NotifCond : G(MaintPCont,"notifCond"),
            RunMode : G(MaintPCont,"runMode"),
            SeqNum : G(MaintPCont,"seqNum"),
            SrUpgrade : G(MaintPCont,"srUpgrade"),
            SrVersion : G(MaintPCont,"srVersion"),
            TriggerTime : G(MaintPCont,"triggerTime"),
            Version : G(MaintPCont,"version"),
            VersionCheckOverride : G(MaintPCont,"versionCheckOverride"),
       },

    }
}

func MaintPFromContainer(cont *container.Container) *MaintP {

    return MaintPFromContainerList(cont, 0)
}

func MaintPListFromContainer(cont *container.Container) []*MaintP {
    length, _ := strconv.Atoi(G(cont, "totalCount"))

    arr := make([]*MaintP, length)

    for i := 0; i < length; i++ {

        arr[i] = MaintPFromContainerList(cont, i)
    }

    return arr
}
