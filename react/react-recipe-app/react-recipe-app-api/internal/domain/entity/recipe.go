package entity

//レシピ情報
type Recipe struct {
	Title       string
	Ingredients []Ingredient
	Calory      int
}

//材料
type Ingredient string
