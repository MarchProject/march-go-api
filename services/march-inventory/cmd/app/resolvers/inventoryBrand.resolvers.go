package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"core/app/helper"
	"core/app/middlewares"
	graph "march-inventory/cmd/app/graph/generated"
	"march-inventory/cmd/app/graph/types"
	"march-inventory/cmd/app/services/inventoryService"
)

// UpsertInventoryBrand is the resolver for the upsertInventoryBrand field.
func (r *mutationResolver) UpsertInventoryBrand(ctx context.Context, input types.UpsertInventoryBrandInput) (*types.MutationInventoryBrandResponse, error) {
	logctx := helper.LogContext(ClassName, "UpsertInventoryBrand")
	// userInfo := middlewares.UserInfo(ctx)
	userInfo := middlewares.UserInfo(ctx)
	logctx.Logger(userInfo, "userInfo")
	return inventoryService.UpsertInventoryBrand(&input, userInfo)
}

// DeleteInventoryBrand is the resolver for the deleteInventoryBrand field.
func (r *mutationResolver) DeleteInventoryBrand(ctx context.Context, id string) (*types.MutationInventoryBrandResponse, error) {
	logctx := helper.LogContext(ClassName, "DeleteInventoryBrand")
	userInfo := middlewares.UserInfo(ctx)
	logctx.Logger(userInfo, "userInfo")
	return inventoryService.DeleteInventoryBrand(id, userInfo)
}

// GetInventoryBrand is the resolver for the getInventoryBrand field.
func (r *queryResolver) GetInventoryBrand(ctx context.Context, id *string) (*types.InventoryBrand, error) {
	logctx := helper.LogContext(ClassName, "UpsertInventoryType")
	// userInfo := middlewares.UserInfo(ctx)
	logctx.Logger([]interface{}{}, "")
	return inventoryService.GetInventoryBrand(id)
}

// GetInventoryBrands is the resolver for the getInventoryBrands field.
func (r *queryResolver) GetInventoryBrands(ctx context.Context, params *types.ParamsInventoryBrand) (*types.InventoryBrandsDataResponse, error) {
	logctx := helper.LogContext(ClassName, "UpsertInventoryType")
	userInfo := middlewares.UserInfo(ctx)
	logctx.Logger(userInfo, "userInfo")
	return inventoryService.GetInventoryBrands(params, userInfo)
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
