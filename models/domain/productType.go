package domain

type ProductType struct {
	ID              uint   `json:"id"`
	TypeName        string `json:"typeName"`
	TypeDescription string `json:"typeDescription"`
}
