package common

import (
	"context"
	"net/http"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type BaseInputs interface {
	GetBase() BaseInputStruct
}

type BaseInputStruct struct {
	Name        string            `pulumi:"name" json:"name"`
	Description *string           `pulumi:"description,optional" computed:"true" json:"description,omitempty"`
	Tags        []string          `pulumi:"tags,optional" computed:"true" json:"tags,omitempty"`
	Metadata    map[string]string `pulumi:"metadata,optional" computed:"true" json:"metadata,omitempty"`
}

func (inputs BaseInputStruct) GetBase() BaseInputStruct {
	return inputs
}

func (inputs BaseInputStruct) ToOutput() BaseOutputStruct {
	return BaseOutputStruct{
		Name:        inputs.Name,
		Description: *inputs.Description,
		Tags:        inputs.Tags,
		Metadata:    inputs.Metadata,
	}
}

type BaseOutput interface {
	GetBase() BaseOutputStruct
}

type BaseOutputStruct struct {
	ID          string            `json:"id"`
	Name        string            `pulumi:"name" json:"name"`
	Description string            `pulumi:"description" json:"description"`
	Tags        []string          `pulumi:"tags" json:"tags"`
	Metadata    map[string]string `pulumi:"metadata,optional" json:"metadata,omitempty"`
}

func (output BaseOutputStruct) GetBase() BaseOutputStruct {
	return output
}

type BaseResource[I BaseInputs, O BaseOutput] struct {
	Path          string
	CreateOutput  func() O
	WithDefaults  func(I) I
	ExtraToOutput func(I, *O)
	ExtraDiff     func(oldValue O, newValue O, diffs map[string]provider.PropertyDiff)
}

func (b BaseResource[I, O]) toOutput(inputs I) O {
	var output O
	if b.ExtraToOutput != nil {
		b.ExtraToOutput(inputs, &output)
	}
	return output
}

func (b BaseResource[I, O]) diffOutput(oldValue O, newValue O) map[string]provider.PropertyDiff {
	oldBase := oldValue.GetBase()
	newBase := newValue.GetBase()
	diffs := map[string]provider.PropertyDiff{
		"name":        DiffString(oldBase.Name, newBase.Name),
		"description": DiffString(oldBase.Description, newBase.Description),
		"tags":        DiffSlice(oldBase.Tags, newBase.Tags),
		"metadata":    DiffMap(oldBase.Metadata, newBase.Metadata),
	}
	if b.ExtraDiff != nil {
		b.ExtraDiff(oldValue, newValue, diffs)
	}
	return diffs
}

func (b BaseResource[I, O]) Create(ctx context.Context, req infer.CreateRequest[I]) (infer.CreateResponse[O], error) {
	inputs := b.WithDefaults(req.Inputs)
	if req.DryRun {
		return infer.CreateResponse[O]{ID: req.Name, Output: b.toOutput(inputs)}, nil
	}
	output, err := b.doRequestAndGetResponse(ctx, http.MethodPost, "", inputs)
	if err != nil {
		return infer.CreateResponse[O]{}, err
	}
	return infer.CreateResponse[O]{ID: output.GetBase().ID, Output: output}, nil
}

func (b BaseResource[I, O]) Read(ctx context.Context, req infer.ReadRequest[I, O]) (infer.ReadResponse[I, O], error) {
	output, err := b.doRequestAndGetResponse(ctx, http.MethodGet, req.ID, nil)
	if err != nil {
		return infer.ReadResponse[I, O]{}, err
	}
	if output.GetBase().ID == "" {
		return infer.ReadResponse[I, O]{}, nil
	}
	return infer.ReadResponse[I, O]{ID: req.ID, Inputs: b.WithDefaults(req.Inputs), State: output}, nil
}

func (b BaseResource[I, O]) Update(ctx context.Context, req infer.UpdateRequest[I, O]) (infer.UpdateResponse[O], error) {
	inputs := b.WithDefaults(req.Inputs)
	if req.DryRun {
		return infer.UpdateResponse[O]{Output: b.toOutput(inputs)}, nil
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

func (b BaseResource[I, O]) Diff(_ context.Context, req infer.DiffRequest[I, O]) (infer.DiffResponse, error) {
	diffs := b.diffOutput(req.State, b.toOutput(b.WithDefaults(req.Inputs)))
	hasChanges := false
	for _, diff := range diffs {
		if diff.InputDiff {
			hasChanges = true
			break
		}
	}
	return infer.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          hasChanges,
		DetailedDiff:        diffs,
	}, nil
}
