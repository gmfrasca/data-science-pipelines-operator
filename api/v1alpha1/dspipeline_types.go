/*
Copyright 2023.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DSPASpec struct {
	// DS Pipelines API Server configuration.
	// +kubebuilder:default:={deploy: true}
	*APIServer `json:"apiServer,omitempty"`
	// DS Pipelines PersistenceAgent configuration.
	// +kubebuilder:default:={deploy: true}
	*PersistenceAgent `json:"persistenceAgent,omitempty"`
	// DS Pipelines Scheduled Workflow configuration.
	// +kubebuilder:default:={deploy: true}
	*ScheduledWorkflow `json:"scheduledWorkflow,omitempty"`
	// Database specifies database configurations, used for DS Pipelines metadata tracking. Specify either the default MariaDB deployment, or configure your own External SQL DB.
	// +kubebuilder:default:={mariaDB: {deploy: true}}
	*Database `json:"database,omitempty"`
	// Deploy the KFP UI with DS Pipelines UI. This feature is unsupported, and primarily used for exploration, testing, and development purposes.
	// +kubebuilder:validation:Optional
	*MlPipelineUI `json:"mlpipelineUI"`
	// ObjectStorage specifies Object Store configurations, used for DS Pipelines artifact passing and storage. Specify either the your own External Storage (e.g. AWS S3), or use the default Minio deployment (unsupported, primarily for development, and testing) .
	// +kubebuilder:validation:Required
	*ObjectStorage `json:"objectStorage"`
	// +kubebuilder:default:={deploy: true}
	*MLMD `json:"mlmd,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:={deploy: false}
	*CRDViewer `json:"crdviewer"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:={deploy: false}
	*VisualizationServer `json:"visualizationServer"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="v1"
	DSPVersion string `json:"dspVersion,omitempty"`
	// DS Pipelines Argo Workflow Controller Configuration.
	// +kubebuilder:default:={deploy: false}
	*WorkflowController `json:"workflowController,omitempty"`
}

type APIServer struct {
	// Enable DS Pipelines Operator management of DSP API Server. Setting Deploy to false disables operator reconciliation. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy bool `json:"deploy"`
	// Specify a custom image for DSP API Server.
	Image string `json:"image,omitempty"`
	// Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	ApplyTektonCustomResource bool `json:"applyTektonCustomResource"`
	// Default: false
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	ArchiveLogs   bool   `json:"archiveLogs"`
	ArtifactImage string `json:"artifactImage,omitempty"`
	CacheImage    string `json:"cacheImage,omitempty"`
	// Image used for internal artifact passing handling within Tekton taskruns. This field specifies the image used in the 'move-all-results-to-tekton-home' step.
	MoveResultsImage         string `json:"moveResultsImage,omitempty"`
	*ArtifactScriptConfigMap `json:"artifactScriptConfigMap,omitempty"`
	// Inject the archive step script. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	InjectDefaultScript bool `json:"injectDefaultScript"`
	// Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	StripEOF bool `json:"stripEOF"`
	// Default: "Cancelled" - Allowed Values: "Cancelled", "StoppedRunFinally", "CancelledRunFinally"
	// +kubebuilder:validation:Enum=Cancelled;StoppedRunFinally;CancelledRunFinally
	// +kubebuilder:default:=Cancelled
	TerminateStatus string `json:"terminateStatus,omitempty"`
	// Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	TrackArtifacts bool `json:"trackArtifacts"`
	// Default: 120
	// +kubebuilder:default:=120
	DBConfigConMaxLifetimeSec int `json:"dbConfigConMaxLifetimeSec,omitempty"`
	// Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	CollectMetrics bool `json:"collectMetrics"`
	// Create an Openshift Route for this DSP API Server. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	EnableRoute bool `json:"enableOauth"`
	// Include sample pipelines with the deployment of this DSP API Server. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	EnableSamplePipeline bool `json:"enableSamplePipeline"`
	// Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	AutoUpdatePipelineDefaultVersion bool `json:"autoUpdatePipelineDefaultVersion"`
	// Specify custom Pod resource requirements for this component.
	Resources *ResourceRequirements `json:"resources,omitempty"`
}

type ArtifactScriptConfigMap struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type PersistenceAgent struct {
	// Enable DS Pipelines Operator management of Persisence Agent. Setting Deploy to false disables operator reconciliation. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy bool `json:"deploy"`
	// Specify a custom image for DSP PersistenceAgent.
	Image string `json:"image,omitempty"`
	// Number of worker for Persistence Agent sync job. Default: 2
	// +kubebuilder:default:=2
	NumWorkers int `json:"numWorkers,omitempty"`
	// Specify custom Pod resource requirements for this component.
	Resources *ResourceRequirements `json:"resources,omitempty"`
}

type ScheduledWorkflow struct {
	// Enable DS Pipelines Operator management of ScheduledWorkflow. Setting Deploy to false disables operator reconciliation. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy bool `json:"deploy"`
	// Specify a custom image for DSP ScheduledWorkflow controller.
	Image string `json:"image,omitempty"`
	// Specify the Cron timezone used for ScheduledWorkflow PipelineRuns. Default: UTC
	// +kubebuilder:default:=UTC
	CronScheduleTimezone string `json:"cronScheduleTimezone,omitempty"`
	// Specify custom Pod resource requirements for this component.
	Resources *ResourceRequirements `json:"resources,omitempty"`
}

type MlPipelineUI struct {
	// Enable DS Pipelines Operator management of KFP UI. Setting Deploy to false disables operator reconciliation. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy        bool   `json:"deploy"`
	ConfigMapName string `json:"configMap,omitempty"`
	// Specify custom Pod resource requirements for this component.
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// Specify a custom image for KFP UI pod.
	// +kubebuilder:validation:Required
	Image string `json:"image"`
}

type Database struct {
	*MariaDB    `json:"mariaDB,omitempty"`
	*ExternalDB `json:"externalDB,omitempty"`
	// Default: false
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	DisableHealthCheck bool `json:"disableHealthCheck"`
}

type MariaDB struct {
	// Enable DS Pipelines Operator management of MariaDB. Setting Deploy to false disables operator reconciliation. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy bool `json:"deploy"`
	// Specify a custom image for DSP MariaDB pod.
	Image string `json:"image,omitempty"`
	// The MariadB username that will be created. Should match `^[a-zA-Z0-9_]+`. Default: mlpipeline
	// +kubebuilder:default:=mlpipeline
	// +kubebuilder:validation:Pattern=`^[a-zA-Z0-9_]+$`
	Username       string          `json:"username,omitempty"`
	PasswordSecret *SecretKeyValue `json:"passwordSecret,omitempty"`
	// +kubebuilder:default:=mlpipeline
	// The database name that will be created. Should match `^[a-zA-Z0-9_]+`. // Default: mlpipeline
	// +kubebuilder:validation:Pattern=`^[a-zA-Z0-9_]+$`
	DBName string `json:"pipelineDBName,omitempty"`
	// Customize the size of the PVC created for the default MariaDB instance. Default: 10Gi
	// +kubebuilder:default:="10Gi"
	PVCSize resource.Quantity `json:"pvcSize,omitempty"`
	// Specify custom Pod resource requirements for this component.
	Resources *ResourceRequirements `json:"resources,omitempty"`
}

type ExternalDB struct {
	// +kubebuilder:validation:Required
	Host           string          `json:"host"`
	Port           string          `json:"port"`
	Username       string          `json:"username"`
	DBName         string          `json:"pipelineDBName"`
	PasswordSecret *SecretKeyValue `json:"passwordSecret"`
}

type ObjectStorage struct {
	// Enable DS Pipelines Operator management of Minio. Setting Deploy to false disables operator reconciliation.
	*Minio           `json:"minio,omitempty"`
	*ExternalStorage `json:"externalStorage,omitempty"`
	// Default: false
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	DisableHealthCheck bool `json:"disableHealthCheck"`
}

type Minio struct {
	// Enable DS Pipelines Operator management of Minio. Setting Deploy to false disables operator reconciliation. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy bool `json:"deploy"`
	// Provide the Bucket name that will be used to store artifacts in S3. If provided bucket does not exist, DSP Apiserver will attempt to create it. As such the credentials provided should have sufficient permissions to do create buckets. Default: mlpipeline
	// +kubebuilder:default:=mlpipeline
	Bucket string `json:"bucket,omitempty"`
	// Credentials for the S3 user (e.g. IAM user cred stored in a k8s secret.). Note that the S3 user should have the permissions to create a bucket if the provided bucket does not exist.
	*S3CredentialSecret `json:"s3CredentialsSecret,omitempty"`
	// Customize the size of the PVC created for the Minio instance. Default: 10Gi
	// +kubebuilder:default:="10Gi"
	PVCSize resource.Quantity `json:"pvcSize,omitempty"`
	// Specify custom Pod resource requirements for this component.
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// Specify a custom image for Minio pod.
	// +kubebuilder:validation:Required
	Image string `json:"image"`
}

type MLMD struct {
	// Enable DS Pipelines Operator management of MLMD. Setting Deploy to false disables operator reconciliation. Default: true
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy  bool `json:"deploy"`
	*Envoy  `json:"envoy,omitempty"`
	*GRPC   `json:"grpc,omitempty"`
	*Writer `json:"writer,omitempty"`
}

type Envoy struct {
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// +kubebuilder:validation:Required
	Image string `json:"image"`
}

type GRPC struct {
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// +kubebuilder:validation:Required
	Image string `json:"image"`
	// +kubebuilder:validation:Optional
	Port string `json:"port"`
}

type Writer struct {
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// +kubebuilder:validation:Required
	Image string `json:"image"`
}

type CRDViewer struct {
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy bool   `json:"deploy"`
	Image  string `json:"image,omitempty"`
}

type VisualizationServer struct {
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy bool   `json:"deploy"`
	Image  string `json:"image,omitempty"`
}

type WorkflowController struct {
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Deploy bool   `json:"deploy"`
	Image  string `json:"image,omitempty"`
}

// ResourceRequirements structures compute resource requirements.
// Replaces ResourceRequirements from corev1 which also includes optional storage field.
// We handle storage field separately, and should not include it as a subfield for Resources.
type ResourceRequirements struct {
	Limits   *Resources `json:"limits,omitempty"`
	Requests *Resources `json:"requests,omitempty"`
}

type Resources struct {
	CPU    resource.Quantity `json:"cpu,omitempty"`
	Memory resource.Quantity `json:"memory,omitempty"`
}

type ExternalStorage struct {
	// +kubebuilder:validation:Required
	Host                string `json:"host"`
	Bucket              string `json:"bucket"`
	Scheme              string `json:"scheme"`
	*S3CredentialSecret `json:"s3CredentialsSecret"`
	// +kubebuilder:validation:Optional
	Secure *bool `json:"secure"`
	// +kubebuilder:validation:Optional
	Port string `json:"port"`
}

type S3CredentialSecret struct {
	// +kubebuilder:validation:Required
	SecretName string `json:"secretName"`
	// The "Keys" in the k8sSecret key/value pairs. Not to be confused with the values.
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type SecretKeyValue struct {
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	Key  string `json:"key"`
}

type DSPAStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=dspa

type DataSciencePipelinesApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DSPASpec   `json:"spec,omitempty"`
	Status            DSPAStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

type DataSciencePipelinesApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSciencePipelinesApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataSciencePipelinesApplication{}, &DataSciencePipelinesApplicationList{})
}
