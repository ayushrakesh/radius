// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package extenders

import (
	"context"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/linkrp/datamodel"
	"github.com/project-radius/radius/pkg/linkrp/renderers"
	rpv1 "github.com/project-radius/radius/pkg/rp/v1"
)

var _ renderers.Renderer = (*Renderer)(nil)

type Renderer struct {
}

func (r *Renderer) Render(ctx context.Context, dm v1.ResourceDataModel, options renderers.RenderOptions) (renderers.RendererOutput, error) {
	resource, ok := dm.(*datamodel.Extender)
	if !ok {
		return renderers.RendererOutput{}, v1.ErrInvalidModelConversion
	}

	properties := resource.Properties

	_, err := renderers.ValidateApplicationID(properties.Application)
	if err != nil {
		return renderers.RendererOutput{}, err
	}
	computedValues, secretValues := MakeSecretsAndValues(resource.Name, properties)

	return renderers.RendererOutput{
		Resources:      []rpv1.OutputResource{},
		ComputedValues: computedValues,
		SecretValues:   secretValues,
	}, nil
}

func MakeSecretsAndValues(name string, properties datamodel.ExtenderProperties) (map[string]renderers.ComputedValueReference, map[string]rpv1.SecretValueReference) {
	computedValueReferences := map[string]renderers.ComputedValueReference{}
	for k, v := range properties.AdditionalProperties {
		computedValueReferences[k] = renderers.ComputedValueReference{
			Value: v,
		}
	}
	if properties.Secrets == nil {
		return computedValueReferences, nil
	}

	// Create secret value references to point to the secret output resources created above
	secretValues := map[string]rpv1.SecretValueReference{}
	for k, v := range properties.Secrets {
		secretValues[k] = rpv1.SecretValueReference{
			Value: v.(string),
		}
	}

	return computedValueReferences, secretValues
}