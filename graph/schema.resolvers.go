package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/insta-code1/go-redis-replica/graph/generated"
	"github.com/insta-code1/go-redis-replica/graph/model"
)

func (r *mutationResolver) CreateRedisMasterDeployment(ctx context.Context, input model.NewDeployment) (*model.Message, error) {
	deployment := r.Blueprints.RedisMasterDeployment()
	if input.Name != "" {
		deployment.ObjectMeta.Name = input.Name
	}
	if input.Image != "" {
		deployment.Spec.Template.Spec.Containers[0].Image = input.Image
	}

	deploymentName := r.Svc.CreateDeployment(deployment)

	return &model.Message{
		Text: fmt.Sprintf("Created deployment %q.\n", deploymentName),
	}, nil
}

func (r *mutationResolver) CreateRedisMasterService(ctx context.Context, input model.NewService) (*model.Message, error) {
	service := r.Blueprints.RedisMasterService()
	if input.Name != "" {
		service.ObjectMeta.Name = input.Name
		service.Spec.ExternalName = input.Name
	}

	serviceName := r.Svc.CreateService(service)

	return &model.Message{
		Text: fmt.Sprintf("Created service %q.\n", serviceName),
	}, nil
}

func (r *mutationResolver) CreateRedisSlaveDeployment(ctx context.Context, input model.NewDeployment) (*model.Message, error) {
	deployment := r.Blueprints.RedisReplicaDeployment()
	if input.Name != "" {
		deployment.ObjectMeta.Name = input.Name
	}
	if input.Image != "" {
		deployment.Spec.Template.Spec.Containers[0].Image = input.Image
	}

	deploymentName := r.Svc.CreateDeployment(deployment)

	return &model.Message{
		Text: fmt.Sprintf("Created deployment %q.\n", deploymentName),
	}, nil
}

func (r *mutationResolver) CreateRedisSlaveService(ctx context.Context, input model.NewService) (*model.Message, error) {
	service := r.Blueprints.RedisReplicaService()
	if input.Name != "" {
		service.ObjectMeta.Name = input.Name
		service.Spec.ExternalName = input.Name
	}

	serviceName := r.Svc.CreateService(service)

	return &model.Message{
		Text: fmt.Sprintf("Created service %q.\n", serviceName),
	}, nil
}

func (r *mutationResolver) UpdateDeployment(ctx context.Context, input model.UpdateDeployment) (*model.Message, error) {
	msg, err := r.Svc.UpdateDeployment(input.Name, input.Image)
	if err != nil {
		return nil, err
	}
	return &model.Message{
		Text: msg,
	}, nil
}

func (r *mutationResolver) DeleteDeployment(ctx context.Context, input model.DeleteDeployment) (*model.Message, error) {
	msg, err := r.Svc.DeleteDeployment(input.Name)
	if err != nil {
		return nil, err
	}
	return &model.Message{
		Text: msg,
	}, nil
}

func (r *queryResolver) ListDeployments(ctx context.Context) ([]*model.Message, error) {
	response := []*model.Message{}
	list, err := r.Svc.ListDeployments()
	if err != nil {
		return nil, err
	}

	for _, i := range list {
		response = append(response, &model.Message{
			Text: i,
		})
	}

	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
