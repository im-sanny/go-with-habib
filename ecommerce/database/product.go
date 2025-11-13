package database

var ProductList []Product

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Fresh Apples",
		Description: "Crisp and sweet red delicious apples",
		Price:       3.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod2 := Product{
		ID:          2,
		Title:       "Ripe Bananas",
		Description: "Fresh yellow bananas, perfect for smoothies",
		Price:       1.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod3 := Product{
		ID:          3,
		Title:       "Sweet Oranges",
		Description: "Juicy oranges packed with vitamin C",
		Price:       4.50,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}

	ProductList = append(ProductList, prod1, prod2, prod3)
}
