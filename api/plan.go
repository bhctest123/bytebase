package api

import (
	"fmt"
	"math"
)

// PlanType is the type for a plan.
type PlanType int

const (
	// FREE is the plan type for FREE.
	FREE PlanType = iota
	// TEAM is the plan type for TEAM.
	TEAM
	// ENTERPRISE is the plan type for ENTERPRISE.
	ENTERPRISE
)

// String returns the string format of plan type.
func (p PlanType) String() string {
	switch p {
	case FREE:
		return "FREE"
	case TEAM:
		return "TEAM"
	case ENTERPRISE:
		return "ENTERPRISE"
	}
	return ""
}

// Priority returns the priority of the plan type.
// Higher priority means the plan supports more features.
func (p PlanType) Priority() int {
	switch p {
	case FREE:
		return 1
	case TEAM:
		return 2
	case ENTERPRISE:
		return 3
	}
	return 0
}

// FeatureType is the type of a feature.
type FeatureType string

const (
	// Admin & Security.

	// Feature3rdPartyAuth allows user to authenticate (login) and authorize (sync project member)
	//
	// Currently, we only support GitLab EE/CE auth.
	Feature3rdPartyAuth FeatureType = "bb.feature.3rd-party-auth"
	// FeatureRBAC enables RBAC.
	//
	// - Workspace level RBAC
	// - Project level RBAC.
	FeatureRBAC FeatureType = "bb.feature.rbac"

	// Branding.

	// FeatureBranding enables customized branding.
	//
	// Currently, we only support customizing the logo.
	FeatureBranding FeatureType = "bb.feature.branding"

	// Change Workflow.

	// FeatureDataSource exposes the data source concept.
	//
	// Currently, we DO NOT expose this feature.
	//
	// Internally Bytebase stores instance username/password in a separate data source model.
	// This allows a single instance to have multiple data sources (e.g. one RW and one RO).
	// And from the user's perspective, the username/password
	// look like the property of the instance, which are not. They are the property of data source which
	// in turns belongs to the instance.
	// - Support defining extra data source for a database and exposing the related data source UI.
	FeatureDataSource FeatureType = "bb.feature.data-source"
	// FeatureDBAWorkflow enforces the DBA workflow.
	//
	// - Developers can't create and view instances since they are exclusively by DBA, they can
	//   only access database.
	// - Developers can submit troubleshooting issue.
	FeatureDBAWorkflow FeatureType = "bb.feature.dba-workflow"
	// FeatureLGTM checks LGTM comments.
	FeatureLGTM FeatureType = "bb.feature.lgtm"
	// FeatureIMApproval integrates IM approval into Bytebase, allowing users approve Bytebase issues on the IM.
	FeatureIMApproval FeatureType = "bb.feature.im.approval"
	// FeatureMultiTenancy allows user to enable tenant mode for the project.
	//
	// Tenant mode allows user to track a group of homogeneous database changes together.
	// e.g. A game studio may deploy many servers, each server is fully isolated with its
	// own database. When a new game version is released, it may require to upgrade the
	// underlying database schema, then tenant mode will help the studio to track the
	// schema change across all databases.
	FeatureMultiTenancy FeatureType = "bb.feature.multi-tenancy"
	// FeatureOnlineMigration allows user to perform online-migration.
	FeatureOnlineMigration FeatureType = "bb.feature.online-migration"
	// FeatureSchemaDrift detects if there occurs schema drift.
	// See https://bytebase.com/docs/features/drift-detection
	FeatureSchemaDrift FeatureType = "bb.feature.schema-drift"
	// FeatureSQLReview allows user to specify schema policy for the environment
	//
	// e.g. One can configure rules for database schema or SQL query.
	FeatureSQLReview FeatureType = "bb.feature.sql-review"
	// FeatureTaskScheduleTime allows user to run task at a scheduled time.
	FeatureTaskScheduleTime FeatureType = "bb.feature.task-schedule-time"
	// FeatureVCSSQLReviewWorkflow allows user to enable the SQL review CI in VCS workflow.
	FeatureVCSSQLReviewWorkflow FeatureType = "bb.feature.vcs-sql-review"

	// Database management.

	// FeaturePITR allows user to perform point-in-time recovery for databases.
	FeaturePITR FeatureType = "bb.feature.pitr"
	// FeatureReadReplicaConnection allows user to set a read replica connection
	// including host and port to data source.
	FeatureReadReplicaConnection FeatureType = "bb.feature.read-replica-connection"
	// FeatureSyncSchemaAllVersions allows user to sync the base database schema all versions into target database.
	FeatureSyncSchemaAllVersions FeatureType = "bb.feature.sync-schema-all-versions"

	// Policy Control.

	// FeatureApprovalPolicy allows user to specify approval policy for the environment
	//
	// e.g. One can configure to NOT require approval for dev environment while require
	//      manual approval for production.
	FeatureApprovalPolicy FeatureType = "bb.feature.approval-policy"
	// FeatureBackupPolicy allows user to specify backup policy for the environment
	//
	// e.g. One can configure to NOT require backup for dev environment while require
	//      weekly backup for staging and daily backup for production.
	FeatureBackupPolicy FeatureType = "bb.feature.backup-policy"
	// FeatureEnvironmentTierPolicy allows user to set the tier of an environment.
	//
	// e.g. set the tier to "PROTECTED" for the production environment.
	FeatureEnvironmentTierPolicy FeatureType = "bb.feature.environment-tier-policy"
	// FeatureSensitiveData allows user to annotate and protect sensitive data.
	FeatureSensitiveData FeatureType = "bb.feature.sensitive-data"
)

// Name returns a readable name of the feature.
func (e FeatureType) Name() string {
	switch e {
	// Admin & Security
	case Feature3rdPartyAuth:
		return "3rd party auth"
	case FeatureRBAC:
		return "RBAC"
	// Branding
	case FeatureBranding:
		return "Branding"
	// Change Workflow
	case FeatureDataSource:
		return "Data source"
	case FeatureDBAWorkflow:
		return "DBA workflow"
	case FeatureLGTM:
		return "LGTM"
	case FeatureIMApproval:
		return "IM approval integration"
	case FeatureMultiTenancy:
		return "Multi-tenancy"
	case FeatureOnlineMigration:
		return "Online schema migration"
	case FeatureSchemaDrift:
		return "Schema drift"
	case FeatureSQLReview:
		return "SQL review"
	case FeatureTaskScheduleTime:
		return "Task schedule time"
	case FeatureVCSSQLReviewWorkflow:
		return "VCS SQL review workflow"
	// Database management
	case FeaturePITR:
		return "Point-in-time Recovery"
	case FeatureReadReplicaConnection:
		return "Read replica connection"
	case FeatureSyncSchemaAllVersions:
		return "Synchronize schema all versions"
	// Policy Control
	case FeatureApprovalPolicy:
		return "Approval policy"
	case FeatureBackupPolicy:
		return "Backup policy"
	case FeatureEnvironmentTierPolicy:
		return "Environment tier"
	case FeatureSensitiveData:
		return "Sensitive data"
	}
	return ""
}

// AccessErrorMessage returns a error message with feature name and minimum supported plan.
func (e FeatureType) AccessErrorMessage() string {
	plan := e.minimumSupportedPlan()
	return fmt.Sprintf("%s is a %s feature, please upgrade to access it.", e.Name(), plan.String())
}

// minimumSupportedPlan will find the minimum plan which supports the target feature.
func (e FeatureType) minimumSupportedPlan() PlanType {
	for i, enabled := range featureMatrix[e] {
		if enabled {
			return PlanType(i)
		}
	}

	return ENTERPRISE
}

// featureMatrix is a map from the a particular feature to the respective enablement of a particular
// plan in [FREE, TEAM, Enterprise].
var featureMatrix = map[FeatureType][3]bool{
	// Admin & Security
	Feature3rdPartyAuth: {false, true, true},
	FeatureRBAC:         {false, true, true},
	// Branding
	FeatureBranding: {false, false, true},
	// Change Workflow
	FeatureDataSource:           {false, false, false},
	FeatureDBAWorkflow:          {false, false, true},
	FeatureLGTM:                 {false, false, true},
	FeatureIMApproval:           {false, false, true},
	FeatureMultiTenancy:         {false, false, true},
	FeatureOnlineMigration:      {false, true, true},
	FeatureSchemaDrift:          {false, true, true},
	FeatureSQLReview:            {true, true, true},
	FeatureTaskScheduleTime:     {false, true, true},
	FeatureVCSSQLReviewWorkflow: {false, false, true},
	// Database management
	FeaturePITR:                  {false, true, true},
	FeatureReadReplicaConnection: {false, false, true},
	FeatureSyncSchemaAllVersions: {false, false, true},
	// Policy Control
	FeatureApprovalPolicy:        {false, true, true},
	FeatureBackupPolicy:          {false, true, true},
	FeatureEnvironmentTierPolicy: {false, false, true},
	FeatureSensitiveData:         {false, false, true},
}

// Plan is the API message for a plan.
type Plan struct {
	Type PlanType `jsonapi:"attr,type"`
}

// PlanPatch is the API message for patching a plan.
type PlanPatch struct {
	Type PlanType `jsonapi:"attr,type"`
}

// TrialPlanCreate is the API message for creating a trial plan.
type TrialPlanCreate struct {
	Type          PlanType `jsonapi:"attr,type"`
	Days          int      `jsonapi:"attr,days"`
	InstanceCount int      `jsonapi:"attr,instanceCount"`
}

// PlanLimit is the type for plan limits.
type PlanLimit int

const (
	// PlanLimitMaximumTask is the key name for maximum number of tasks for a plan.
	PlanLimitMaximumTask PlanLimit = iota
	// PlanLimitMaximumEnvironment is the key name for maximum number of environments for a plan.
	PlanLimitMaximumEnvironment
)

// PlanLimitValues is the plan limit value mapping.
var PlanLimitValues = map[PlanLimit][3]int64{
	PlanLimitMaximumTask:        {math.MaxInt64, math.MaxInt64, math.MaxInt64},
	PlanLimitMaximumEnvironment: {math.MaxInt64, math.MaxInt64, math.MaxInt64},
}

// Feature returns whether a particular feature is available in a particular plan.
func Feature(feature FeatureType, plan PlanType) bool {
	return featureMatrix[feature][plan]
}
