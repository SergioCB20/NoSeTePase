package models

type Category struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Color string `json:"color"` // para mostrar en UI
}
