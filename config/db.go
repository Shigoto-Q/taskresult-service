package config

var DBUrl string = "postgres://debug:debug@localhost:5432/shigoto_q"

var Query  = `SELECT tasks_taskresult.task_name" FROM "tasks_taskresult" WHERE "tasks_taskresult"."user_id" = 1 ORDER BY "tasks_taskresult"."date_done" DESC`

type TaskResult struct {
  task_name string
}

