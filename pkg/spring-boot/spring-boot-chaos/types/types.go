package types

import (
	"k8s.io/api/core/v1"
	clientTypes "k8s.io/apimachinery/pkg/types"
)

// ExperimentDetails is for collecting all the experiment-related details
type ExperimentDetails struct {
	ExperimentName     string
	EngineName         string
	ChaosDuration      int
	ChaosInterval      int
	RampTime           int
	ChaosLib           string
	AppNS              string
	AppLabel           string
	AppKind            string
	ChaosUID           clientTypes.UID
	InstanceID         string
	ChaosNamespace     string
	ChaosPodName       string
	Timeout            int
	Delay              int
	TargetContainer    string
	PodsAffectedPerc   int
	TargetPods         string
	LIBImagePullPolicy string
	Sequence           string
	TargetPodList      v1.PodList

	// Chaos monkey parameters
	ChaosMonkeyAssault  ChaosMonkeyAssault
	ChaosMonkeyWatchers ChaosMonkeyWatchers
	ChaosMonkeyPath     string
	ChaosMonkeyPort     string
}

type ChaosMonkeyAssaultRevert struct {
	LatencyActive         bool `json:"latencyActive"`
	KillApplicationActive bool `json:"killApplicationActive"`
	MemoryActive          bool `json:"memoryActive"`
	CPUActive             bool `json:"cpuActive"`
	ExceptionsActive      bool `json:"exceptionsActive"`
}

type ChaosMonkeyAssault struct {
	Level                              int              `json:"level"`
	Deterministic                      bool             `json:"deterministic"`
	LatencyRangeStart                  int              `json:"latencyRangeStart"`
	LatencyRangeEnd                    int              `json:"latencyRangeEnd"`
	LatencyActive                      bool             `json:"latencyActive"`
	ExceptionsActive                   bool             `json:"exceptionsActive"`
	Exception                          AssaultException `json:"exceptions"`
	KillApplicationActive              bool             `json:"killApplicationActive"`
	KillApplicationCron                string           `json:"killApplicationCronExpression"`
	WatchedCustomServices              []string         `json:"watchedCustomServices"`
	MemoryActive                       bool             `json:"memoryActive"`
	MemoryMillisecondsHoldFilledMemory int              `json:"memoryMillisecondsHoldFilledMemory"`
	MemoryMillisecondsWaitNextIncrease int              `json:"memoryMillisecondsWaitNextIncrease"`
	MemoryFillIncrementFraction        float64          `json:"memoryFillIncrementFraction"`
	MemoryFillTargetFraction           float64          `json:"memoryFillTargetFraction"`
	MemoryCron                         string           `json:"memoryCronExpression"`
	CPUActive                          bool             `json:"cpuActive"`
	CPUMillisecondsHoldLoad            int              `json:"cpuMillisecondsHoldLoad"`
	CPULoadTargetFraction              float64          `json:"cpuLoadTargetFraction"`
	CPUCron                            string           `json:"cpuCronExpression"`
}

type ChaosMonkeyWatchers struct {
	Controller     bool `json:"controller"`
	RestController bool `json:"restController"`
	Service        bool `json:"service"`
	Repository     bool `json:"repository"`
	Component      bool `json:"component"`
	RestTemplate   bool `json:"restTemplate"`
	WebClient      bool `json:"webClient"`
}

type AssaultException struct {
	Type      string                     `json:"type"`
	Arguments []AssaultExceptionArgument `json:"arguments"`
}

type AssaultExceptionArgument struct {
	ClassName string `json:"className"`
	Value     string `json:"value"`
}
