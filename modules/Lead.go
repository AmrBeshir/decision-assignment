package modules

import "reflect"

type Lead struct {
	lead_type         string
	method_of_contact string
}

func CreateLead(lead_type string, method_of_contact string) Lead {
	return Lead{lead_type: lead_type, method_of_contact: method_of_contact}
}

func (lead *Lead) GetField(field string) string {
	r := reflect.ValueOf(lead)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
