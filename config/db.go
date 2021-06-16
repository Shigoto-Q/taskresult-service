package config

import (
  "fmt"
  "log"
  // "github.com/georgysavva/scany/pgxscan"
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)

var Query  = `SELECT tasks_taskresult.task_name, tasks_taskresult.task_id, 
                      tasks_taskresult.status, tasks_taskresult.date_done, 
                      tasks_taskresult.date_created, tasks_taskresult.user_id
              FROM tasks_taskresult WHERE tasks_taskresult.user_id = %d 
              ORDER BY tasks_taskresult.date_done DESC`
var tasksQuery = `SELECT
                      COUNT(tasks_taskresult.id) FILTER (WHERE tasks_taskresult.status = 'SUCCESS') AS "success",
                      COUNT(tasks_taskresult.id) FILTER (WHERE tasks_taskresult.status = 'FAILURE') AS "failure",
                      COUNT(tasks_taskresult.id) FILTER (WHERE tasks_taskresult.status = 'PENDING') AS "pending"
                  FROM tasks_taskresult
                  WHERE tasks_taskresult.user_id = %d`
type TaskResult struct {
  Task_name string
  Task_id string
  Status string
  Date_done string
  Date_created string
  User_id string
}

type TaskStatus struct {
  Success int
  Failure int
  Pending int
}

func SetupDb() *sqlx.DB {
    db, err := sqlx.Connect("postgres", "user=debug password=debug dbname=shigoto_q sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }
    return db
}

func FetchResultStatus(userID int, db *sqlx.DB) *[]TaskStatus {
    taskstatuses := []TaskStatus{}
    err := db.Select(&taskstatuses, fmt.Sprintf(tasksQuery, userID))
    if err != nil {
      log.Fatalln(err)
    }
    return &taskstatuses
}

func FetchResults(userID int, db *sqlx.DB) *[]TaskResult {
  taskresult := []TaskResult{}
  err := db.Select(&taskresult, fmt.Sprintf(Query, userID))
  if err != nil {
      log.Fatalln(err)
    }
  return &taskresult
}
