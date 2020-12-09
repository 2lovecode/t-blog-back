package category

import (
	"context"
	"t-blog-back/models"
)

func CategoryList(ctx context.Context) (categories []models.Category, err error) {
	cg := &models.Category{}
	return cg.FindAll(ctx)
}
