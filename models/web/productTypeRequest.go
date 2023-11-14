package web

type ProductTypeCreate struct {
	TypeName        string `json:"typeName" validate:"required"`
	TypeDescription string `json:"typeDescription" validate:"required"`
}

type ProductTypeUpdate struct {
	TypeName        string `json:"typeName" validate:"required"`
	TypeDescription string `json:"typeDescription" validate:"required"`
}
