package transformtask

import (
	taskpb "grpc_test/proto/task"
	"grpc_test/src/models"
)

func TaskToTaskPB(task models.Task) *taskpb.Task {
	return &taskpb.Task{
		Id:   int32(task.Id),
		Name: task.Name,
	}
}

func ArrayTaskToArrayTaskPB(tasks []models.Task) []*taskpb.Task {
	var result []*taskpb.Task

	for _, item := range tasks {
		result = append(result, TaskToTaskPB(item))
	}

	return result
}
