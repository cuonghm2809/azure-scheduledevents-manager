package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"time"
)

type (
	Opts struct {
		// logger
		Logger struct {
			Debug   bool `           long:"debug"        env:"DEBUG"    description:"debug mode"`
			Verbose bool `short:"v"  long:"verbose"      env:"VERBOSE"  description:"verbose mode"`
			LogJson bool `           long:"log.json"     env:"LOG_JSON" description:"Switch log output to json format"`
		}

		General struct {
			// general options
			ServerBind string        `long:"bind"                env:"SERVER_BIND"   description:"Server address"                default:":8080"`
			ScrapeTime time.Duration `long:"scrape-time"         env:"SCRAPE_TIME"   description:"Scrape time in seconds"        default:"1m"`
		}

		// Api option
		Azure struct {
			InstanceApiUrl        string        `long:"azure.metadatainstance-url"    env:"AZURE_METADATAINSTANCE_URL"    description:"Azure ScheduledEvents API URL" default:"http://169.254.169.254/metadata/instance?api-version=2019-08-01"`
			ScheduledEventsApiUrl string        `long:"azure.scheduledevents-url"     env:"AZURE_SCHEDULEDEVENTS_URL"     description:"Azure ScheduledEvents API URL" default:"http://169.254.169.254/metadata/scheduledevents?api-version=2019-08-01"`
			Timeout               time.Duration `long:"azure.timeout"                 env:"AZURE_TIMEOUT"                 description:"Azure API timeout (seconds)"   default:"30s"`
			ErrorThreshold        int           `long:"azure.error-threshold"         env:"AZURE_ERROR_THRESHOLD"         description:"Azure API error threshold (after which app will panic)"   default:"0"`
			ApproveScheduledEvent bool          `long:"azure.approve-scheduledevent"  env:"AZURE_APPROVE_SCHEDULEDEVENT"  description:"Approve ScheduledEvent and start (if possible) start them ASAP"`
		}

		Instance struct {
			VmNodeName string `long:"vm.nodename"    env:"VM_NODENAME"     description:"VM node name"`
		}

		Drain struct {
			Enable    bool          `long:"drain.enable"             env:"DRAIN_ENABLE"                description:"Enable drain handling"`
			Mode      string        `long:"drain.mode"               env:"DRAIN_MODE"                  description:"Mode" choice:"kubernetes" choice:"command" choice:"noop"` //nolint:golint,staticcheck
			NotBefore time.Duration `long:"drain.not-before"         env:"DRAIN_NOT_BEFORE"            description:"Dont drain before this time" default:"5m"`
			Events    []string      `long:"drain.events"             env:"DRAIN_EVENTS" env-delim:" "  description:"Enable drain handling" default:"reboot" default:"redeploy" default:"preempt" default:"terminate"` //nolint:staticcheck
		}

		Command struct {
			Test struct {
				Cmd string `long:"command.test.cmd"  env:"COMMAND_TEST_CMD"   description:"Test command in command mode"`
			}
			Drain struct {
				Cmd string `long:"command.drain.cmd"  env:"COMMAND_DRAIN_CMD"   description:"Drain command in command mode"`
			}
			Uncordon struct {
				Cmd string `long:"command.uncordon.cmd"  env:"COMMAND_UNCORDON_CMD"   description:"Uncordon command in command mode"`
			}
		}

		Kubernetes struct {
			NodeName string `long:"kube.nodename"  env:"KUBE_NODENAME"   description:"Kubernetes node name"`

			Drain struct {
				DeleteLocalData  bool          `long:"kube.drain.delete-local-data"  env:"KUBE_DRAIN_DELETE_LOCAL_DATA"     description:"Continue even if there are pods using emptyDir (local data that will be deleted when the node is drained)"`
				Force            bool          `long:"kube.drain.force"              env:"KUBE_DRAIN_FORCE"                 description:"Continue even if there are pods not managed by a ReplicationController, ReplicaSet, Job, DaemonSet or StatefulSet"`
				GracePeriod      int64         `long:"kube.drain.grace-period"       env:"KUBE_DRAIN_GRACE_PERIOD"          description:"Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be used."`
				IgnoreDaemonsets bool          `long:"kube.drain.ignore-daemonsets"  env:"KUBE_DRAIN_IGNORE_DAEMONSETS"     description:"Ignore DaemonSet-managed pods."`
				PodSelector      string        `long:"kube.drain.pod-selector"       env:"KUBE_DRAIN_POD_SELECTOR"          description:"Label selector to filter pods on the node"`
				Timeout          time.Duration `long:"kube.drain.timeout"            env:"KUBE_DRAIN_TIMEOUT"               description:"The length of time to wait before giving up, zero means infinite" default:"0s"`
				DryRun           bool          `long:"kube.drain.dry-run"            env:"KUBE_DRAIN_DRY_RUN"               description:"Do not drain, uncordon or label any node"`
			}
		}

		Notification struct {
			List        []string `long:"notification"                 env:"NOTIFICATION"              description:"Shoutrrr url for notifications (https://containrrr.github.io/shoutrrr/)" env-delim:" "`
			MsgTemplate string   `long:"notification.messagetemplate" env:"NOTIFICATION_MESSAGE_TEMPLATE"  description:"Notification template" default:"%v"`
		}

		Metrics struct {
			RequestStats bool `long:"metrics-requeststats" env:"METRICS_REQUESTSTATS" description:"Enable request stats metrics"`
		}
	}
)

func (o *Opts) GetJson() []byte {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return jsonBytes
}
