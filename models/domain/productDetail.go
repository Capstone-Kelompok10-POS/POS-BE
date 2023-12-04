package domain

type ProductDetail struct {
<<<<<<< HEAD
	ID         uint    `json:"id"`
	ProductID  uint    `json:"productID"`
	Price      float64 `json:"price"`
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
=======
	ID         uint
	ProductID  uint
	Product    Product
	Price      float64
	TotalStock int
	Size       string
>>>>>>> a700314b5a44a7e1a732995fa2aa6dd7907760e8
}
