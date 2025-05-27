package services

import (
	"fmt"
	"reflect"
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
			"options": variant.Options,
			"price":   variant.Price,
			"stock":   variant.Stock,
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
	var variants []models.ProductVariant
	for _, variant := range p.Variants {
		// options := variant["options"].(map[string]interface{})
		price := variant["price"].(float64)
		stock := int(variant["stock"].(float64))

		var variantOptions models.IntSlice
		// fmt.Println("variant['options']", variant["options"])
		// variantOptions := make(map[string]string)
		for _, value := range variant["options"].([]interface{}) {
			// 在这里处理每个 value
			if val, ok := value.(float64); ok {
				// 处理 int 类型的 value
				variantOptions = append(variantOptions, int(val))
			} else {
				fmt.Println("value 無法轉爲 int:", reflect.TypeOf(value))
			}
		}

		// fmt.Println("variant['options']:", reflect.TypeOf(variant["options"]))
		// fmt.Println("variantOptions:", variantOptions)

		variants = append(variants, models.ProductVariant{
			ProductID: p.ID,
			Options:   variantOptions,
			Price:     price,
			Stock:     stock,
		})
	}

	// fmt.Println("variants:", variants)

	// fmt.Println("input:", p.DetailImages)
	p.UpdateTime = time.Now()
	p.CreateTime = time.Now()
	p.IsDeleted = false

	// fmt.Println("product:", p)
	return s.productRepo.AddProduct(p, options, specifications, variants)

}

// DeleteProduct deletes a product by its ID from the repository
func (s *ProductService) DeleteProduct(id int) error {
	return s.productRepo.DeleteProduct(id)
}

// DeleteProducts deletes multiple products by their IDs from the repository
func (s *ProductService) DeleteProducts(ids []int) error {
	return s.productRepo.DeleteProducts(ids)
}

func (s *ProductService) CopyProduct(id int) (*models.Product, error) {
	return s.productRepo.CopyProduct(id)
}

func (s *ProductService) UpdateProduct(p *models.Product) error {
	// 处理 Options
	optionMap := make(map[string][]string)
	for _, option := range p.Options {
		name := option["name"].(string)
		values := option["values"].([]interface{})
		for _, value := range values {
			optionMap[name] = append(optionMap[name], value.(string))
		}
	}

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
	var variants []models.ProductVariant
	for _, variant := range p.Variants {
		price := variant["price"].(float64)
		stock := int(variant["stock"].(float64))

		var variantOptions models.IntSlice
		for _, value := range variant["options"].([]interface{}) {
			if val, ok := value.(float64); ok {
				// 将 float64 转换为 int
				variantOptions = append(variantOptions, int(val))
			} else {
				fmt.Println("value 無法轉爲 int:", reflect.TypeOf(value))
			}
		}

		variants = append(variants, models.ProductVariant{
			ProductID: p.ID,
			Options:   variantOptions,
			Price:     price,
			Stock:     stock,
		})
	}

	p.UpdateTime = time.Now()
	// p.CreateTime = time.Now()
	// p.IsDeleted = false

	return s.productRepo.UpdateProduct(p, options, specifications, variants)
}

// 根據CategoryID獲取商品
func (s *ProductService) GetProductsByCategoryID(id int) ([]*models.Product, error) {
	return s.productRepo.GetProductsByCategoryID(id)
}

// 根據BrandID獲取商品
func (s *ProductService) GetProductsByBrandID(id int) ([]*models.Product, error) {
	return s.productRepo.GetProductsByBrandID(id)
}

// 根據BrandID獲取商品
func (s *ProductService) GetProductsByName(name string) ([]*models.Product, error) {
	return s.productRepo.GetProductsByName(name)
}
