package repositories

import (
	"context"
	"database/sql"
	"fmt"
	productsimagesmodel "golang-starter/src/modules/product/entities"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/nurcahyaari/sqlabst"
)

type RepositoryProductsImagesCommand interface {
	InsertProductsImagesList(ctx context.Context, productsImagesList productsimagesmodel.ProductsImagesList) (*InsertResult, error)
	InsertProductsImages(ctx context.Context, productsImages *productsimagesmodel.ProductsImages) (*InsertResult, error)
	UpdateProductsImagesByFilter(ctx context.Context, productsImages *productsimagesmodel.ProductsImages, filter Filter, updatedFields ...ProductsImagesField) error
	UpdateProductsImages(ctx context.Context, productsImages *productsimagesmodel.ProductsImages, productimagesid int32, updatedFields ...ProductsImagesField) error
	DeleteProductsImagesList(ctx context.Context, filter Filter) error
	DeleteProductsImages(ctx context.Context, productimagesid int32) error
}

type RepositoryProductsImagesCommandImpl struct {
	db *sqlabst.SqlAbst
}

func (repo *RepositoryProductsImagesCommandImpl) InsertProductsImagesList(ctx context.Context, productsImagesList productsimagesmodel.ProductsImagesList) (*InsertResult, error) {
	command := `INSERT INTO products_images (product_fkid,
	images,
	created_at,
	updated_at) VALUES
		`

	var (
		placeholders []string
		args         []interface{}
	)
	for _, productsImages := range productsImagesList {
		placeholders = append(placeholders, `(?,
	?,
	?,
	?)`)
		args = append(args,
			productsImages.ProductFkid,
			productsImages.Images,
			productsImages.CreatedAt,
			productsImages.UpdatedAt,
		)
	}
	command += strings.Join(placeholders, ",")

	sqlResult, err := repo.exec(ctx, command, args)
	if err != nil {
		return nil, err
	}

	return &InsertResult{Result: sqlResult}, nil
}

func (repo *RepositoryProductsImagesCommandImpl) InsertProductsImages(ctx context.Context, productsImages *productsimagesmodel.ProductsImages) (*InsertResult, error) {
	return repo.InsertProductsImagesList(ctx, productsimagesmodel.ProductsImagesList{productsImages})
}

func (repo *RepositoryProductsImagesCommandImpl) UpdateProductsImagesByFilter(ctx context.Context, productsImages *productsimagesmodel.ProductsImages, filter Filter, updatedFields ...ProductsImagesField) error {
	updatedFieldQuery, values := buildUpdateFieldsProductsImagesQuery(updatedFields, productsImages)
	command := fmt.Sprintf(`UPDATE products_images 
			SET %s 
		WHERE %s
		`, strings.Join(updatedFieldQuery, ","), filter.Query())
	values = append(values, filter.Values()...)
	_, err := repo.exec(ctx, command, values)
	return err
}

func (repo *RepositoryProductsImagesCommandImpl) UpdateProductsImages(ctx context.Context, productsImages *productsimagesmodel.ProductsImages, productimagesid int32, updatedFields ...ProductsImagesField) error {
	updatedFieldQuery, values := buildUpdateFieldsProductsImagesQuery(updatedFields, productsImages)
	command := fmt.Sprintf(`UPDATE products_images 
			SET %s 
		WHERE productimages_id = ?
		`, strings.Join(updatedFieldQuery, ","))
	values = append(values, productimagesid)
	_, err := repo.exec(ctx, command, values)
	return err
}

func (repo *RepositoryProductsImagesCommandImpl) DeleteProductsImagesList(ctx context.Context, filter Filter) error {
	command := "DELETE FROM products_images WHERE " + filter.Query()
	_, err := repo.exec(ctx, command, filter.Values())
	return err
}

func (repo *RepositoryProductsImagesCommandImpl) DeleteProductsImages(ctx context.Context, productimagesid int32) error {
	command := "DELETE FROM products_images WHERE productimages_id = ?"
	_, err := repo.exec(ctx, command, []interface{}{productimagesid})
	return err
}

func NewRepoProductsImagesCommand(db *sqlabst.SqlAbst) RepositoryProductsImagesCommand {
	return &RepositoryProductsImagesCommandImpl{
		db: db,
	}
}

func (repo *RepositoryProductsImagesCommandImpl) exec(ctx context.Context, command string, args []interface{}) (sql.Result, error) {
	var (
		stmt *sqlx.Stmt
		err  error
	)
	stmt, err = repo.db.PreparexContext(ctx, command)

	if err != nil {
		return nil, err
	}

	return stmt.ExecContext(ctx, args...)
}

func buildUpdateFieldsProductsImagesQuery(updatedFields ProductsImagesFieldList, productsImages *productsimagesmodel.ProductsImages) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
	)

	for _, field := range updatedFields {
		switch field {
		case "productimages_id":
			updatedFieldsQuery = append(updatedFieldsQuery, "productimages_id = ?")
			args = append(args, productsImages.ProductimagesId)
		case "product_fkid":
			updatedFieldsQuery = append(updatedFieldsQuery, "product_fkid = ?")
			args = append(args, productsImages.ProductFkid)
		case "images":
			updatedFieldsQuery = append(updatedFieldsQuery, "images = ?")
			args = append(args, productsImages.Images)
		case "created_at":
			updatedFieldsQuery = append(updatedFieldsQuery, "created_at = ?")
			args = append(args, productsImages.CreatedAt)
		case "updated_at":
			updatedFieldsQuery = append(updatedFieldsQuery, "updated_at = ?")
			args = append(args, productsImages.UpdatedAt)
		}
	}

	return updatedFieldsQuery, args
}
