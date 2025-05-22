package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductsService(engine *gorm.DB) *ProductService {
	return &ProductService{productRepo: repositories.NewProductRepository(engine)}
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	return s.productRepo.GetAllProducts()
}

func (s *ProductService) GetProductDetails(id int) (*models.Product, error) {
	product, specifications, options, variants, err := s.productRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	// 将 specifications 按 category 分组
	specMap := make(map[string][]map[string]string)
	for _, spec := range specifications {
		item := map[string]string{"name": spec.Name, "value": spec.Value}
		specMap[spec.Category] = append(specMap[spec.Category], item)
	}

	// 将分组后的 specifications 转换为所需格式
	var formattedSpecifications []map[string]interface{}
	for category, items := range specMap {
		formattedSpecifications = append(formattedSpecifications, map[string]interface{}{
			"category": category,
			"items":    items,
		})
	}

	// 将 options 按类型分组
	optionMap := make(map[string][]string)
	for _, option := range options {
		for _, value := range option.Values {
			optionMap[option.Name] = append(optionMap[option.Name], value)
		}
	}

	// 将 optionMap 转换为所需格式
	var formattedOptions []map[string]interface{}
	for name, values := range optionMap {
		formattedOptions = append(formattedOptions, map[string]interface{}{
			"name":   name,
			"values": values,
		})
	}

	// 将 variants 转换为所需格式
	var formattedVariants []map[string]interface{}
	for _, variant := range variants {
		formattedVariants = append(formattedVariants, map[string]interface{}{
			"option": variant.Options,
			"price":  variant.Price,
			"stock":  variant.Stock,
		})
	}

	product.Specifications = formattedSpecifications
	product.Options = formattedOptions
	product.Variants = formattedVariants

	return product, nil
}

func (s *ProductService) GetProductById(id int) (*models.Product, error) {
	product, err := s.GetProductDetails(id)

	if err != nil {
		return nil, err
	}

	// fmt.Println("product:", newProduct.Specifications)
	// return us.productRepo.GetProductById(id)
	return product, nil
}

func (s *ProductService) AddProduct(p *models.Product) error {

	// 处理 Options
	optionMap := make(map[string][]string)
	for _, option := range p.Options {
		name := option["name"].(string)
		values := option["values"].([]interface{})
		for _, value := range values {
			optionMap[name] = append(optionMap[name], value.(string))
		}
	}

	// fmt.Println(optionMap)
	// 将 optionMap 转换为所需格式并存储到数据库中
	var options []*models.Option
	for name, values := range optionMap {
		var stringSlice models.StringSlice
		stringSlice.FromString(strings.Join(values, ","))
		options = append(options, &models.Option{
			ProductID: p.ID,
			Name:      name,
			Values:    stringSlice,
		})
	}

	// 处理 Specifications
	specMap := make(map[string][]map[string]string)
	for _, spec := range p.Specifications {
		category := spec["category"].(string)
		items := spec["items"].([]interface{})
		for _, item := range items {
			itemMap := item.(map[string]interface{})
			name := itemMap["name"].(string)
			value := itemMap["value"].(string)
			specMap[category] = append(specMap[category], map[string]string{"name": name, "value": value})
		}
	}

	// 将 specMap 转换为所需格式并存储到数据库中
	var specifications []*models.Specification
	for category, items := range specMap {
		for _, item := range items {
			specifications = append(specifications, &models.Specification{
				ProductID: p.ID,
				Category:  category,
				Name:      item["name"],
				Value:     item["value"],
			})
		}
	}

	// 处理 Variants
	// var variants []*models.ProductVariant
	// for _, variant := range p.Variants {
	// 	options := variant["option"].(map[string]interface{})
	// 	price := variant["price"].(float64)
	// 	stock := int(variant["stock"].(float64))
	// }

	// fmt.Println("input:", p.DetailImages)
	p.UpdateTime = time.Now()
	p.CreateTime = time.Now()
	p.IsDeleted = false

	// fmt.Println("product:", p)
	return s.productRepo.AddProduct(p, options, specifications)

	// return s.productRepo.AddProduct(&input.Product, options, specifications)
}
