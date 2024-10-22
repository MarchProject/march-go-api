package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"core/app/helper"
	"core/app/middlewares"
	"march-inventory/cmd/app/graph/types"
	"march-inventory/cmd/app/services/inventoryService"
)

// UpsertInventoryBranch is the resolver for the upsertInventoryBranch field.
func (r *mutationResolver) UpsertInventoryBranch(ctx context.Context, input types.UpsertInventoryBranchInput) (*types.MutationInventoryBranchResponse, error) {
	logctx := helper.LogContext(ClassName, "UpsertInventoryType")
	userInfo := middlewares.UserInfo(ctx)
	logctx.Logger(userInfo, "userInfo")
	return inventoryService.UpsertInventoryBranch(&input, userInfo)
}

// DeleteInventoryBranch is the resolver for the deleteInventoryBranch field.
func (r *mutationResolver) DeleteInventoryBranch(ctx context.Context, id string) (*types.MutationInventoryBranchResponse, error) {
	logctx := helper.LogContext(ClassName, "UpsertInventoryType")
	userInfo := middlewares.UserInfo(ctx)
	logctx.Logger(userInfo, "userInfo")
	return inventoryService.DeleteInventoryBranch(id, userInfo)
}

// GetInventoryBranchs is the resolver for the getInventoryBranchs field.
func (r *queryResolver) GetInventoryBranchs(ctx context.Context, params *types.ParamsInventoryBranch) (*types.InventoryBranchsDataResponse, error) {
	logctx := helper.LogContext(ClassName, "UpsertInventoryType")
	userInfo := middlewares.UserInfo(ctx)
	logctx.Logger(userInfo, "userInfo")
	return inventoryService.GetInventoryBranchs(params, userInfo)
}
