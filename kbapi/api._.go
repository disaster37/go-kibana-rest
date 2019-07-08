package kbapi

import (
	"github.com/go-resty/resty"
)

type API struct {
	KibanaSpaces         *KibanaSpacesAPI
	KibanaRoleManagement *KibanaRoleManagementAPI
	KibanaDashboard      *KibanaDashboardAPI
	KibanaSavedObject    *KibanaSavedObjectAPI
}

//KibanaSpaces contain the Kibana spaces API
type KibanaSpacesAPI struct {
	Get    KibanaSpaceGet
	List   KibanaSpaceList
	Create KibanaSpaceCreate
	Delete KibanaSpaceDelete
	Update KibanaSpaceUpdate
}

type KibanaRoleManagementAPI struct {
	Get            KibanaRoleManagementGet
	List           KibanaRoleManagementList
	CreateOrUpdate KibanaRoleManagementCreateOrUpdate
	Delete         KibanaRoleManagementDelete
}

type KibanaDashboardAPI struct {
	Export KibanaDashboardExport
	Import KibanaDashboardImport
}

type KibanaSavedObjectAPI struct {
	Get    KibanaSavedObjectGet
	Create KibanaSavedObjectCreate
	Update KibanaSavedObjectUpdate
	Delete KibanaSavedObjectDelete
}

func New(c *resty.Client) *API {
	return &API{
		KibanaSpaces: &KibanaSpacesAPI{
			Get:    newKibanaSpaceGetFunc(c),
			List:   newKibanaSpaceListFunc(c),
			Create: newKibanaSpaceCreateFunc(c),
			Update: newKibanaSpaceUpdateFunc(c),
			Delete: newKibanaSpaceDeleteFunc(c),
		},
		KibanaRoleManagement: &KibanaRoleManagementAPI{
			Get:            newKibanaRoleManagementGetFunc(c),
			List:           newKibanaRoleManagementListFunc(c),
			CreateOrUpdate: newKibanaRoleManagementCreateOrUpdateFunc(c),
			Delete:         newKibanaRoleManagementDeleteFunc(c),
		},
		KibanaDashboard: &KibanaDashboardAPI{
			Export: newKibanaDashboardExportFunc(c),
			Import: newKibanaDashboardImportFunc(c),
		},
		KibanaSavedObject: &KibanaSavedObjectAPI{
			Get:    newKibanaSavedObjectGetFunc(c),
			Create: newKibanaSavedObjectCreateFunc(c),
			Update: newKibanaSavedObjectUpdateFunc(c),
			Delete: newKibanaSavedObjectDeleteFunc(c),
		},
	}
}
