package repositories

import (
	"context"
	"errors"
	"fmt"
	productsimagesmodel "golang-starter/src/modules/product/entities"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/nurcahyaari/sqlabst"
)

type RepositoryProductsImagesQuery interface {
	SelectProductsImages(fields ...ProductsImagesField) RepositoryProductsImagesQuery
	ExcludeProductsImages(excludedFields ...ProductsImagesField) RepositoryProductsImagesQuery
	FilterProductsImages(filter Filter) RepositoryProductsImagesQuery
	PaginationProductsImages(pagination Pagination) RepositoryProductsImagesQuery
	OrderByProductsImages(orderBy []Order) RepositoryProductsImagesQuery
	GetProductsImagesCount(ctx context.Context) (int, error)
	GetProductsImages(ctx context.Context) (*productsimagesmodel.ProductsImages, error)
	GetProductsImagesList(ctx context.Context) (productsimagesmodel.ProductsImagesList, error)
}

type RepositoryProductsImagesQueryImpl struct {
	db         *sqlabst.SqlAbst
	query      string
	filter     Filter
	orderBy    []Order
	pagination Pagination
	fields     ProductsImagesFieldList
}

func (repo *RepositoryProductsImagesQueryImpl) SelectProductsImages(fields ...ProductsImagesField) RepositoryProductsImagesQuery {
	return &RepositoryProductsImagesQueryImpl{
		db:         repo.db,
		filter:     repo.filter,
		orderBy:    repo.orderBy,
		pagination: repo.pagination,
		fields:     fields,
	}
}

func (repo *RepositoryProductsImagesQueryImpl) ExcludeProductsImages(excludedFields ...ProductsImagesField) RepositoryProductsImagesQuery {
	selectedFieldsStr := excludeFields(ProductsImagesFieldList(excludedFields).toString(),
		ProductsImagesSelectFields{}.All().toString())

	var selectedFields []ProductsImagesField
	for _, sel := range selectedFieldsStr {
		selectedFields = append(selectedFields, ProductsImagesField(sel))
	}

	return &RepositoryProductsImagesQueryImpl{
		db:         repo.db,
		filter:     repo.filter,
		orderBy:    repo.orderBy,
		pagination: repo.pagination,
		fields:     selectedFields,
	}
}

func (repo *RepositoryProductsImagesQueryImpl) FilterProductsImages(filter Filter) RepositoryProductsImagesQuery {
	return &RepositoryProductsImagesQueryImpl{
		db:         repo.db,
		filter:     filter,
		orderBy:    repo.orderBy,
		pagination: repo.pagination,
		fields:     repo.fields,
	}
}

func (repo *RepositoryProductsImagesQueryImpl) PaginationProductsImages(pagination Pagination) RepositoryProductsImagesQuery {
	return &RepositoryProductsImagesQueryImpl{
		db:         repo.db,
		filter:     repo.filter,
		orderBy:    repo.orderBy,
		pagination: pagination,
		fields:     repo.fields,
	}
}

func (repo *RepositoryProductsImagesQueryImpl) OrderByProductsImages(orderBy []Order) RepositoryProductsImagesQuery {
	return &RepositoryProductsImagesQueryImpl{
		db:         repo.db,
		filter:     repo.filter,
		orderBy:    orderBy,
		pagination: repo.pagination,
		fields:     repo.fields,
	}
}

func (repo *RepositoryProductsImagesQueryImpl) GetProductsImagesList(ctx context.Context) (productsimagesmodel.ProductsImagesList, error) {
	var (
		productsImagesList productsimagesmodel.ProductsImagesList
		values             []interface{}
	)

	if len(repo.fields) == 0 {
		repo.fields = ProductsImagesSelectFields{}.All()
	}

	query := fmt.Sprintf("SELECT %s FROM products_images", strings.Join(repo.fields.toString(), ","))
	if repo.filter != nil {
		query += " WHERE " + repo.filter.Query()
		values = append(values, repo.filter.Values()...)
	}

	if len(repo.orderBy) > 0 {
		var orderStr []string
		for _, order := range repo.orderBy {
			orderStr = append(orderStr, order.Value()+" "+order.Direction())
		}
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(orderStr, ","))
	}

	if repo.pagination != nil {
		offset := (repo.pagination.GetPage() - 1) * repo.pagination.GetSize()
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", repo.pagination.GetSize(), offset)
	}

	err := repo.db.SelectContext(ctx, &productsImagesList, query, values...)
	if err != nil {
		return nil, err
	}
	return productsImagesList, nil
}

func (repo *RepositoryProductsImagesQueryImpl) GetProductsImagesCount(ctx context.Context) (int, error) {
	var values []interface{}
	query := fmt.Sprintf("SELECT count(1) FROM products_images")
	if repo.filter != nil {
		query += " WHERE " + repo.filter.Query()
		values = append(values, repo.filter.Values()...)
	}

	var count int
	err := repo.db.QueryRowContext(ctx, query, values...).Scan(&count)
	return count, err
}

func (repo *RepositoryProductsImagesQueryImpl) GetProductsImages(ctx context.Context) (*productsimagesmodel.ProductsImages, error) {
	productsImagesList, err := repo.GetProductsImagesList(ctx)
	if err != nil {
		return nil, err
	}

	if len(productsImagesList) == 0 {
		return nil, errors.New("productsimages not found")
	}

	return productsImagesList[0], nil
}

func NewRepoProductsImagesQuery(db *sqlabst.SqlAbst) RepositoryProductsImagesQuery {
	return &RepositoryProductsImagesQueryImpl{
		db: db,
	}
}

type ProductsImagesField string
type ProductsImagesFieldList []ProductsImagesField

func (fieldList ProductsImagesFieldList) toString() []string {
	var fieldsStr []string
	for _, field := range fieldList {
		fieldsStr = append(fieldsStr, string(field))
	}
	return fieldsStr
}

type ProductsImagesSelectFields struct {
}

func (ProductsImagesSelectFields) ProductimagesId() ProductsImagesField {
	return ProductsImagesField("productimages_id")
}
func (ProductsImagesSelectFields) ProductFkid() ProductsImagesField {
	return ProductsImagesField("product_fkid")
}
func (ProductsImagesSelectFields) Images() ProductsImagesField {
	return ProductsImagesField("images")
}
func (ProductsImagesSelectFields) CreatedAt() ProductsImagesField {
	return ProductsImagesField("created_at")
}
func (ProductsImagesSelectFields) UpdatedAt() ProductsImagesField {
	return ProductsImagesField("updated_at")
}

func (ProductsImagesSelectFields) All() ProductsImagesFieldList {
	return []ProductsImagesField{
		ProductsImagesField("productimages_id"),
		ProductsImagesField("product_fkid"),
		ProductsImagesField("images"),
		ProductsImagesField("created_at"),
		ProductsImagesField("updated_at"),
	}
}

func NewProductsImagesSelectFields() ProductsImagesSelectFields {
	return ProductsImagesSelectFields{}
}

type ProductsImagesFilter struct {
	operator string
	query    []string
	values   []interface{}
}

func NewProductsImagesFilter(operator string) ProductsImagesFilter {
	if operator == "" {
		operator = "AND"
	}
	return ProductsImagesFilter{
		operator: operator,
	}
}

func (f ProductsImagesFilter) SetFilterByProductimagesId(value interface{}, operator string) ProductsImagesFilter {
	query := "productimages_id " + operator + " (?)"
	var values []interface{}
	if value == nil {
		query = "productimages_id " + operator
	} else {
		switch strings.ToUpper(operator) {
		case "IN", "NOT IN":
			query, values, _ = sqlx.In(query, value)
		default:
			values = append(values, value)
		}
	}
	return ProductsImagesFilter{
		operator: f.operator,
		query:    append(f.query, query),
		values:   append(f.values, values...),
	}
}
func (f ProductsImagesFilter) SetFilterByProductFkid(value interface{}, operator string) ProductsImagesFilter {
	query := "product_fkid " + operator + " (?)"
	var values []interface{}
	if value == nil {
		query = "product_fkid " + operator
	} else {
		switch strings.ToUpper(operator) {
		case "IN", "NOT IN":
			query, values, _ = sqlx.In(query, value)
		default:
			values = append(values, value)
		}
	}
	return ProductsImagesFilter{
		operator: f.operator,
		query:    append(f.query, query),
		values:   append(f.values, values...),
	}
}
func (f ProductsImagesFilter) SetFilterByImages(value interface{}, operator string) ProductsImagesFilter {
	query := "images " + operator + " (?)"
	var values []interface{}
	if value == nil {
		query = "images " + operator
	} else {
		switch strings.ToUpper(operator) {
		case "IN", "NOT IN":
			query, values, _ = sqlx.In(query, value)
		default:
			values = append(values, value)
		}
	}
	return ProductsImagesFilter{
		operator: f.operator,
		query:    append(f.query, query),
		values:   append(f.values, values...),
	}
}
func (f ProductsImagesFilter) SetFilterByCreatedAt(value interface{}, operator string) ProductsImagesFilter {
	query := "created_at " + operator + " (?)"
	var values []interface{}
	if value == nil {
		query = "created_at " + operator
	} else {
		switch strings.ToUpper(operator) {
		case "IN", "NOT IN":
			query, values, _ = sqlx.In(query, value)
		default:
			values = append(values, value)
		}
	}
	return ProductsImagesFilter{
		operator: f.operator,
		query:    append(f.query, query),
		values:   append(f.values, values...),
	}
}
func (f ProductsImagesFilter) SetFilterByUpdatedAt(value interface{}, operator string) ProductsImagesFilter {
	query := "updated_at " + operator + " (?)"
	var values []interface{}
	if value == nil {
		query = "updated_at " + operator
	} else {
		switch strings.ToUpper(operator) {
		case "IN", "NOT IN":
			query, values, _ = sqlx.In(query, value)
		default:
			values = append(values, value)
		}
	}
	return ProductsImagesFilter{
		operator: f.operator,
		query:    append(f.query, query),
		values:   append(f.values, values...),
	}
}

func (f ProductsImagesFilter) Query() string {
	return strings.Join(f.query, " "+f.operator+" ")
}

func (f ProductsImagesFilter) Values() []interface{} {
	return f.values
}

type ProductsImagesProductimagesIdOrder struct {
	direction string
}

func (o ProductsImagesProductimagesIdOrder) SetDirection(direction string) ProductsImagesProductimagesIdOrder {
	return ProductsImagesProductimagesIdOrder{
		direction: direction,
	}
}
func (o ProductsImagesProductimagesIdOrder) Value() string {
	return "productimages_id"
}
func (o ProductsImagesProductimagesIdOrder) Direction() string {
	return o.direction
}
func NewProductsImagesProductimagesIdOrder() ProductsImagesProductimagesIdOrder {
	return ProductsImagesProductimagesIdOrder{}
}

type ProductsImagesProductFkidOrder struct {
	direction string
}

func (o ProductsImagesProductFkidOrder) SetDirection(direction string) ProductsImagesProductFkidOrder {
	return ProductsImagesProductFkidOrder{
		direction: direction,
	}
}
func (o ProductsImagesProductFkidOrder) Value() string {
	return "product_fkid"
}
func (o ProductsImagesProductFkidOrder) Direction() string {
	return o.direction
}
func NewProductsImagesProductFkidOrder() ProductsImagesProductFkidOrder {
	return ProductsImagesProductFkidOrder{}
}

type ProductsImagesImagesOrder struct {
	direction string
}

func (o ProductsImagesImagesOrder) SetDirection(direction string) ProductsImagesImagesOrder {
	return ProductsImagesImagesOrder{
		direction: direction,
	}
}
func (o ProductsImagesImagesOrder) Value() string {
	return "images"
}
func (o ProductsImagesImagesOrder) Direction() string {
	return o.direction
}
func NewProductsImagesImagesOrder() ProductsImagesImagesOrder {
	return ProductsImagesImagesOrder{}
}

type ProductsImagesCreatedAtOrder struct {
	direction string
}

func (o ProductsImagesCreatedAtOrder) SetDirection(direction string) ProductsImagesCreatedAtOrder {
	return ProductsImagesCreatedAtOrder{
		direction: direction,
	}
}
func (o ProductsImagesCreatedAtOrder) Value() string {
	return "created_at"
}
func (o ProductsImagesCreatedAtOrder) Direction() string {
	return o.direction
}
func NewProductsImagesCreatedAtOrder() ProductsImagesCreatedAtOrder {
	return ProductsImagesCreatedAtOrder{}
}

type ProductsImagesUpdatedAtOrder struct {
	direction string
}

func (o ProductsImagesUpdatedAtOrder) SetDirection(direction string) ProductsImagesUpdatedAtOrder {
	return ProductsImagesUpdatedAtOrder{
		direction: direction,
	}
}
func (o ProductsImagesUpdatedAtOrder) Value() string {
	return "updated_at"
}
func (o ProductsImagesUpdatedAtOrder) Direction() string {
	return o.direction
}
func NewProductsImagesUpdatedAtOrder() ProductsImagesUpdatedAtOrder {
	return ProductsImagesUpdatedAtOrder{}
}
