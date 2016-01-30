package main
import (
	"github.com/alexjlockwood/gcm"
)

type GcmError struct {
	Result gcm.Result
	OldToken string
	MulticastId int64
	Worker int
}
type CustomErrorLog struct {
	TimeStamp string
	Type string
	Data interface{}
}

type GcmTokenUpdateMsg struct {
	OldToken string
	NewToken string
}
type GcmStatusInactiveMsg struct {
	Token string
}
type Configuration struct {
	DebugMode bool `json:"DebugMode"`
	Rabbit struct {
				   Username string `json:"Username"`
				   Password string `json:"Password"`
				   Host string `json:"Host"`
				   Port int `json:"Port"`
				   Vhost string `json:"Vhost"`
				   ReconnectWaitTimeSec int `json:"ReconnectWaitTimeSec"`
				   CreateQueues bool `json:"CreateQueues"`
			   } `json:"Rabbit"`
	GcmQueues []GcmQueue `json:"GcmQueues"`
	GCM struct {
				   ApiKey	string `json:"ApiKey"`
			   } `json:"GCM"`
	Logging struct {
				GcmErr struct {
						   RootPath string `json:"RootPath"`
						   LogSuccess bool `json:"LogSuccess"`
					   } `json:"GcmErr"`
				AppErr struct {
						   FilePath string `json:"FilePath"`
					   } `json:"AppErr"`
			   } `json:"Logging"`
	Db struct {
		   DbHost string `json:"DbHost"`
		   DbPort int `json:"DbPort"`
		   DbUser string `json:"DbUser"`
		   DbPassword string `json:"DbPassword"`
		   DbDatabase string `json:"DbDatabase"`
		   TransactionMinCount struct {
					  TokenUpdate int `json:"TokenUpdate"`
					  StatusInactive int `json:"StatusInactive"`
				  } `json:"TransactionMinCount"`
	   } `json:"Db"`
}

type Message struct {
	Token []string `json:"Token"`
	Body map[string]interface{} `json:"Body"`
}

type GcmQueue struct {
	Identifier string `json:"identifier"`
	Name string `json:"Name"`
	Numworkers int `json:"NumWorkers"`
	ApiKey string `json:"ApiKey"`
	GcmTokenUpdateQueue string `json:"GcmTokenUpdateQueue"`
	GcmStatusInactiveQueue string `json:"GcmStatusInactiveQueue"`
	Queries struct {
				TokenUpdate string `json:"TokenUpdate"`
				StatusInactive string `json:"StatusInactive"`
			} `json:"Queries"`
}


const NeedAck = 1
const NoAckNeeded = 2

// TODO: Make configurable
// TODO: Refactor these consts
const StatusErrTokenUpdateTransaction = "ErrTokenUpdateTransaction"
const StatusErrStatusInactiveTransaction = "ErrStatusInactiveTransaction"
const StatusErrGcmError = "ErrGcmError"
const StatusSuccessGcmRequest = "SuccessGcm"
const StatusSuccessTokenUpdateTransaction = "SuccessTokenUpdateTransaction"
const StatusSuccessStatusInactiveTransaction = "SuccessStatusInactiveTransaction"