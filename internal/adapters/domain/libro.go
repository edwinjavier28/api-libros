package domain

type ResponseLibro struct {
	ID         string
	Nombre     string
	Editorial  string
	Autor      string
	Precio     float64
	Cantidades int
	Edicion    int
	BestSeller bool
}

type LibroRequest struct {
	Nombre     string  `json:"Nombre"`
	Editorial  string  `json:"Editorial"`
	Autor      string  `json:"Autor"`
	Precio     float64 `json:"Precio"`
	Cantidades int     `json:"Cantidades"`
	Edicion    int     `json:"Edicion"`
	Bestseller bool    `json:"Bestseller"`
}
