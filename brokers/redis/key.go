package redis

import "fmt"

const (
	taskQueueKey     = "%s:tasks:queue"
	taskQueueExecKey = taskQueueKey + ":exec"
	taskDistinctKey  = taskQueueKey + ":%s:dis"
	taskPrepareKey   = "task:prepare:queue"
	taskChannelsKey  = "task:channels"
)

// TaskQueueKey function return the key for task queue
func TaskQueueKey(channel string) string {
	return fmt.Sprintf(taskQueueKey, channel)
}

// TaskQueueExecKey function return the key for the queue in execution
func TaskQueueExecKey(channel string) string {
	return fmt.Sprintf(taskQueueExecKey, channel)
}

// TaskQueueDistinctKey function  is used for distinct queue
func TaskQueueDistinctKey(channel string, command string) string {
	return fmt.Sprintf(taskDistinctKey, channel, command)
}

// TaskPrepareQueueKey return the prepare key for queue
func TaskPrepareQueueKey() string {
	return taskPrepareKey
}

// TaskChannelsKey return the key for channels storage
func TaskChannelsKey() string {
	return taskChannelsKey
}
