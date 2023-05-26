/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	dspav1alpha1 "github.com/opendatahub-io/data-science-pipelines-operator/api/v1alpha1"
)

var mlmdTemplates = []string{
	"ml-metadata/metadata-envoy.deployment.yaml.tmpl",
	"ml-metadata/metadata-envoy.service.yaml.tmpl",
	"ml-metadata/metadata-grpc.deployment.yaml.tmpl",
	"ml-metadata/metadata-grpc.service.yaml.tmpl",
	"ml-metadata/metadata-grpc.serviceaccount.yaml.tmpl",
	"ml-metadata/metadata-writer.deployment.yaml.tmpl",
	"ml-metadata/metadata-writer.role.yaml.tmpl",
	"ml-metadata/metadata-writer.rolebinding.yaml.tmpl",
	"ml-metadata/metadata-writer.serviceaccount.yaml.tmpl",
}

func (r *DSPAReconciler) ReconcileMLMD(dsp *dspav1alpha1.DataSciencePipelinesApplication,
	params *DSPAParams) error {

	log := r.Log.WithValues("namespace", dsp.Namespace).WithValues("dspa_name", dsp.Name)

	if params.UsingMLMD() {
		log.Info("Applying ML-Metadata (MLMD) Resources")

		for _, template := range mlmdTemplates {
			err := r.Apply(dsp, params, template)
			if err != nil {
				return err
			}
		}
		log.Info("Finished applying MLMD Resources")
	}
	return nil
}
