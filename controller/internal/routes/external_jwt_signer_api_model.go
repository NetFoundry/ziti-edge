/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package routes

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/env"
	"github.com/openziti/edge/controller/model"
	"github.com/openziti/edge/controller/response"
	"github.com/openziti/edge/rest_model"
	"github.com/openziti/fabric/controller/models"
	"github.com/openziti/foundation/util/stringz"
)

const EntityNameExternalJwtSigner = "external-jwt-signers"

var ExternalJwtSignerLinkFactory = NewBasicLinkFactory(EntityNameExternalJwtSigner)

func MapExternalJwtSignerToRestEntity(_ *env.AppEnv, _ *response.RequestContext, ExternalJwtSignerModel models.Entity) (interface{}, error) {
	externalJwtSigner, ok := ExternalJwtSignerModel.(*model.ExternalJwtSigner)

	if !ok {
		err := fmt.Errorf("entity is not an identity type \"%s\"", ExternalJwtSignerModel.GetId())
		log := pfxlog.Logger()
		log.Error(err)
		return nil, err
	}

	restModel := MapExternalJwtSignerToRestModel(externalJwtSigner)

	return restModel, nil
}

func MapExternalJwtSignerToRestModel(externalJwtSigner *model.ExternalJwtSigner) *rest_model.ExternalJwtSignerDetail {
	notAfter := strfmt.DateTime(externalJwtSigner.NotAfter)
	notBefore := strfmt.DateTime(externalJwtSigner.NotBefore)

	ret := &rest_model.ExternalJwtSignerDetail{
		BaseEntity:  BaseEntityToRestModel(externalJwtSigner, ExternalJwtSignerLinkFactory),
		CertPem:     &externalJwtSigner.CertPem,
		CommonName:  &externalJwtSigner.CommonName,
		Enabled:     &externalJwtSigner.Enabled,
		Fingerprint: &externalJwtSigner.Fingerprint,
		Name:        &externalJwtSigner.Name,
		NotAfter:    &notAfter,
		NotBefore:   &notBefore,
	}
	return ret
}

func MapCreateExternalJwtSignerToModel(signer *rest_model.ExternalJwtSignerCreate) *model.ExternalJwtSigner {
	return &model.ExternalJwtSigner{
		BaseEntity: models.BaseEntity{},
		Name:       *signer.Name,
		CertPem:    *signer.CertPem,
		Enabled:    *signer.Enabled,
	}
}

func MapUpdateExternalJwtSignerToModel(id string, signer *rest_model.ExternalJwtSignerUpdate) *model.ExternalJwtSigner {
	var tags map[string]interface{}
	if signer.Tags != nil && signer.Tags.SubTags != nil {
		tags = signer.Tags.SubTags
	}

	return &model.ExternalJwtSigner{
		BaseEntity: models.BaseEntity{
			Id:       id,
			Tags:     tags,
			IsSystem: false,
		},
		Name:    *signer.Name,
		CertPem: *signer.CertPem,
		Enabled: *signer.Enabled,
	}
}

func MapPatchExternalJwtSignerToModel(id string, signer *rest_model.ExternalJwtSignerPatch) *model.ExternalJwtSigner {
	enabled := false
	if signer.Enabled != nil {
		enabled = *signer.Enabled
	}

	var tags map[string]interface{}
	if signer.Tags != nil && signer.Tags.SubTags != nil {
		tags = signer.Tags.SubTags
	}

	return &model.ExternalJwtSigner{
		BaseEntity: models.BaseEntity{
			Id:       id,
			Tags:     tags,
			IsSystem: false,
		},
		Name:    stringz.OrEmpty(signer.Name),
		CertPem: stringz.OrEmpty(signer.CertPem),
		Enabled: enabled,
	}
}
