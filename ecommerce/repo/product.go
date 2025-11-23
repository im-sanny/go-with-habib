package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList { // [a, b, c]
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	prod1 := &Product{
		ID:          1,
		Title:       "Fresh Apples",
		Description: "Crisp and sweet red delicious apples",
		Price:       3.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod2 := &Product{
		ID:          2,
		Title:       "Ripe Bananas",
		Description: "Fresh yellow bananas, perfect for smoothies",
		Price:       1.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod3 := &Product{
		ID:          3,
		Title:       "Sweet Oranges",
		Description: "Juicy oranges packed with vitamin C",
		Price:       4.50,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}

	r.productList = append(r.productList, prod1, prod2, prod3)
}
