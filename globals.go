package evergreen

import (
	"os"
	"time"

	"github.com/mongodb/grip"
)

const (
	User            = "mci"
	GithubPatchUser = "github_pull_request"

	HostRunning         = "running"
	HostTerminated      = "terminated"
	HostUninitialized   = "initializing"
	HostBuilding        = "building"
	HostStarting        = "starting"
	HostProvisioning    = "provisioning"
	HostProvisionFailed = "provision failed"
	HostQuarantined     = "quarantined"
	HostDecommissioned  = "decommissioned"

	HostStatusSuccess = "success"
	HostStatusFailed  = "failed"

	// Task Statuses used in the database models

	// TaskInactive is not assigned to any new tasks, but can be found
	// in the database and is used in the UI.
	TaskInactive = "inactive"

	// TaskUnstarted is assigned to a display task after cleaning up one of
	// its execution tasks. This indicates that the display task is
	// pending a rerun
	TaskUnstarted = "unstarted"

	// TaskUndispatched indicates either
	//  1. a task is not scheduled to run (when Task.Activated == false)
	//  2. a task is scheduled to run (when Task.Activated == true)
	TaskUndispatched = "undispatched"

	// TaskStarted indicates a task is running on an agent
	TaskStarted = "started"

	// TaskDispatched indicates that an agent has received the task, but
	// the agent has not yet told Evergreen that it's running the task
	TaskDispatched = "dispatched"

	// The task statuses below indicate that a task has finished.
	TaskSucceeded = "success"

	// These statuses indicate the types of failures that are stored in
	// Task.Status field, build TaskCache and TaskEndDetails.
	TaskFailed       = "failed"
	TaskSystemFailed = "system-failed"
	TaskTestTimedOut = "test-timed-out"
	TaskSetupFailed  = "setup-failed"

	// Task Command Types
	CommandTypeTest   = "test"
	CommandTypeSystem = "system"
	CommandTypeSetup  = "setup"

	// Task Statuses that are currently used only by the UI, and in tests
	// (these may be used in old tasks)
	TaskSystemUnresponse = "system-unresponsive"
	TaskSystemTimedOut   = "system-timed-out"
	TaskTimedOut         = "task-timed-out"

	// TaskConflict is used only in communication with the Agent
	TaskConflict = "task-conflict"

	TestFailedStatus         = "fail"
	TestSilentlyFailedStatus = "silentfail"
	TestSkippedStatus        = "skip"
	TestSucceededStatus      = "pass"

	BuildStarted   = "started"
	BuildCreated   = "created"
	BuildFailed    = "failed"
	BuildSucceeded = "success"

	VersionStarted   = "started"
	VersionCreated   = "created"
	VersionFailed    = "failed"
	VersionSucceeded = "success"

	PatchCreated   = "created"
	PatchStarted   = "started"
	PatchSucceeded = "succeeded"
	PatchFailed    = "failed"

	PushLogPushing = "pushing"
	PushLogSuccess = "success"

	HostTypeStatic = "static"

	CompileStage = "compile"
	PushStage    = "push"

	// maximum task (zero based) execution number
	MaxTaskExecution = 3

	// maximum task priority
	MaxTaskPriority = 100

	// LogMessage struct versions
	LogmessageFormatTimestamp = 1
	LogmessageCurrentVersion  = LogmessageFormatTimestamp

	EvergreenHome = "EVGHOME"
	MongodbUrl    = "MONGO_URL"

	// Special logging output targets
	LocalLoggingOverride          = "LOCAL"
	StandardOutputLoggingOverride = "STDOUT"

	DefaultTaskActivator   = ""
	StepbackTaskActivator  = "stepback"
	APIServerTaskActivator = "apiserver"

	RestRoutePrefix = "rest"
	APIRoutePrefix  = "api"

	AgentAPIVersion  = 2
	APIRoutePrefixV2 = "/rest/v2"

	DegradedLoggingPercent = 10

	SetupScriptName    = "setup.sh"
	TeardownScriptName = "teardown.sh"

	RoutePaginatorNextPageHeaderKey = "Link"
)

func IsFinishedTaskStatus(status string) bool {
	if status == TaskSucceeded ||
		IsFailedTaskStatus(status) {
		return true
	}

	return false
}

func IsFailedTaskStatus(status string) bool {
	return status == TaskFailed ||
		status == TaskSystemFailed ||
		status == TaskSystemTimedOut ||
		status == TaskSystemUnresponse ||
		status == TaskTestTimedOut ||
		status == TaskSetupFailed
}

// evergreen package names
const (
	UIPackage      = "EVERGREEN_UI"
	RESTV2Package  = "EVERGREEN_REST_V2"
	MonitorPackage = "EVERGREEN_MONITOR"
)

const (
	AuthTokenCookie     = "mci-token"
	TaskSecretHeader    = "Task-Secret"
	HostHeader          = "Host-Id"
	HostSecretHeader    = "Host-Secret"
	ContentTypeHeader   = "Content-Type"
	ContentTypeValue    = "application/json"
	ContentLengthHeader = "Content-Length"
	APIUserHeader       = "Api-User"
	APIKeyHeader        = "Api-Key"
)

// cloud provider related constants
const (
	ProviderNameEc2Auto     = "ec2-auto"
	ProviderNameEc2OnDemand = "ec2-ondemand"
	ProviderNameEc2Spot     = "ec2-spot"
	ProviderNameDocker      = "docker"
	ProviderNameDockerMock  = "docker-mock"
	ProviderNameGce         = "gce"
	ProviderNameStatic      = "static"
	ProviderNameOpenstack   = "openstack"
	ProviderNameVsphere     = "vsphere"
	ProviderNameMock        = "mock"

	// TODO: This can be removed when no more hosts with provider ec2 are running.
	ProviderNameEc2Legacy = "ec2"
)

var (
	// Providers where hosts can be created and terminated automatically.
	ProviderSpawnable = []string{
		ProviderNameDocker,
		ProviderNameEc2Legacy,
		ProviderNameEc2OnDemand,
		ProviderNameEc2Spot,
		ProviderNameEc2Auto,
		ProviderNameGce,
		ProviderNameOpenstack,
		ProviderNameVsphere,
		ProviderNameMock,
	}
	SystemVersionRequesterTypes = []string{
		RepotrackerVersionRequester,
		TriggerRequester,
	}
)

const (
	DefaultServiceConfigurationFileName = "/etc/mci_settings.yml"
	DefaultDatabaseUrl                  = "localhost:27017"
	DefaultDatabaseName                 = "mci"

	// database and config directory, set to the testing version by default for safety
	NotificationsFile = "mci-notifications.yml"
	ClientDirectory   = "clients"

	// version requester types
	PatchVersionRequester       = "patch_request"
	GithubPRRequester           = "github_pull_request"
	RepotrackerVersionRequester = "gitter_request"
	TriggerRequester            = "trigger_request"
	AdHocRequester              = "ad_hoc"
)

const (
	GenerateTasksCommandName = "generate.tasks"
	CreateHostCommandName    = "host.create"
)

type SenderKey int

const (
	SenderGithubStatus = SenderKey(iota)
	SenderEvergreenWebhook
	SenderSlack
	SenderJIRAIssue
	SenderJIRAComment
	SenderEmail
)

func (k SenderKey) String() string {
	switch k {
	case SenderGithubStatus:
		return "github-status"
	case SenderEmail:
		return "email"
	case SenderEvergreenWebhook:
		return "webhook"
	case SenderSlack:
		return "slack"
	case SenderJIRAComment:
		return "jira-comment"
	case SenderJIRAIssue:
		return "jira-issue"
	default:
		return "<error:unkwown>"
	}
}

const (
	defaultLogBufferingDuration  = 20
	defaultMgoDialTimeout        = 5 * time.Second
	defaultAmboyPoolSize         = 2
	defaultAmboyLocalStorageSize = 1024
	defaultAmboyQueueName        = "evg.service"
	defaultAmboyDBName           = "amboy"
	maxNotificationsPerSecond    = 100
)

// NameTimeFormat is the format in which to log times like instance start time.
const NameTimeFormat = "20060102150405"

var (
	PatchRequesters = []string{
		PatchVersionRequester,
		GithubPRRequester,
	}

	// UphostStatus is a list of all host statuses that are considered "up."
	// This is used for query building.
	UphostStatus = []string{
		HostRunning,
		HostUninitialized,
		HostBuilding,
		HostStarting,
		HostProvisioning,
		HostProvisionFailed,
	}

	// Hosts in "initializing" status aren't actually running yet:
	// they're just intents, so this list omits that value.
	ActiveStatus = []string{
		HostRunning,
		HostStarting,
		HostProvisioning,
		HostProvisionFailed,
	}

	// constant arrays for db update logic
	AbortableStatuses = []string{TaskStarted, TaskDispatched}
	CompletedStatuses = []string{TaskSucceeded, TaskFailed}

	ValidCommandTypes = []string{CommandTypeSetup, CommandTypeSystem, CommandTypeTest}
)

// FindEvergreenHome finds the directory of the EVGHOME environment variable.
func FindEvergreenHome() string {
	// check if env var is set
	root := os.Getenv(EvergreenHome)
	if len(root) > 0 {
		return root
	}

	grip.Errorf("%s is unset", EvergreenHome)
	return ""
}

// IsSystemActivator returns true when the task activator is Evergreen.
func IsSystemActivator(caller string) bool {
	return caller == DefaultTaskActivator || caller == APIServerTaskActivator
}

func IsPatchRequester(requester string) bool {
	return requester == PatchVersionRequester || requester == GithubPRRequester
}
