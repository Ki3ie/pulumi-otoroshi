package common

import (
	"context"
	"net/http"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type BaseInputs interface{}

type BaseOutput interface {
	GetID() string
}

type BaseOutputStruct struct {
	ID string `json:"id"`
}

func (b BaseOutputStruct) GetID() string {
	return b.ID
}

type BaseResource[I BaseInputs, O BaseOutput] struct {
	Path         string
	WithDefaults func(I) I
	ToOutput     func(I) O
	DiffOutput   func(O, O) map[string]provider.PropertyDiff
}

func (b BaseResource[I, O]) Create(ctx context.Context, req infer.CreateRequest[I]) (infer.CreateResponse[O], error) {
	inputs := b.WithDefaults(req.Inputs)
	if req.DryRun {
		return infer.CreateResponse[O]{ID: req.Name, Output: b.ToOutput(inputs)}, nil
	}
	output, err := b.doRequestAndGetResponse(ctx, http.MethodPost, "", inputs)
	if err != nil {
		return infer.CreateResponse[O]{}, err
	}
	return infer.CreateResponse[O]{ID: output.GetID(), Output: output}, nil
}

func (b BaseResource[I, O]) Read(ctx context.Context, req infer.ReadRequest[I, O]) (infer.ReadResponse[I, O], error) {
	output, err := b.doRequestAndGetResponse(ctx, http.MethodGet, req.ID, nil)
	if err != nil {
		return infer.ReadResponse[I, O]{}, err
	}
	if output.GetID() == "" {
		return infer.ReadResponse[I, O]{}, nil
	}
	return infer.ReadResponse[I, O]{ID: req.ID, Inputs: b.WithDefaults(req.Inputs), State: output}, nil
}

func (b BaseResource[I, O]) Update(ctx context.Context, req infer.UpdateRequest[I, O]) (infer.UpdateResponse[O], error) {
	inputs := b.WithDefaults(req.Inputs)
	if req.DryRun {
		return infer.UpdateResponse[O]{Output: b.ToOutput(inputs)}, nil
	}
	output, err := b.doRequestAndGetResponse(ctx, http.MethodPut, req.ID, inputs)
	if err != nil {
		return infer.UpdateResponse[O]{}, err
	}
	return infer.UpdateResponse[O]{Output: output}, nil
}

func (b BaseResource[I, O]) Delete(ctx context.Context, req infer.DeleteRequest[O]) (infer.DeleteResponse, error) {
	_, err := b.doRequestAndGetResponse(ctx, http.MethodDelete, req.ID, nil)
	if err != nil {
		return infer.DeleteResponse{}, err
	}
	return infer.DeleteResponse{}, nil
}

func (b BaseResource[I, O]) Diff(ctx context.Context, req infer.DiffRequest[I, O]) (infer.DiffResponse, error) {
	diffs := b.DiffOutput(req.State, b.ToOutput(b.WithDefaults(req.Inputs)))
	return infer.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diffs) > 0,
		DetailedDiff:        diffs,
	}, nil
}
