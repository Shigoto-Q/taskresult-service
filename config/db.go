package config

import (
  "fmt"
  "log"
  // "github.com/georgysavva/scany/pgxscan"
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)
var DBUrl string = "postgres://debug:debug@localhost:5432/shigoto_q"

var Query  = `SELECT tasks_taskresult.task_name, tasks_taskresult.task_id, 
                      tasks_taskresult.status, tasks_taskresult.date_done, 
                      tasks_taskresult.date_created, tasks_taskresult.user_id
              FROM tasks_taskresult WHERE tasks_taskresult.user_id = %d 
              ORDER BY tasks_taskresult.date_done DESC`

type TaskResult struct {
  Task_name string
  Task_id string
  Status string
  Date_done string
  Date_created string
  User_id string
}


func FormatQuery(userid int) string {
  return fmt.Sprintf(Query, userid)
}

func SetupDb() *sqlx.DB {
    db, err := sqlx.Connect("postgres", "user=debug password=debug dbname=shigoto_q sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }
    return db
}


func FetchResults(userID int, db *sqlx.DB) TaskResult {
  rows, err := db.Queryx(fmt.Sprintf(Query, userID))
  if err != nil {
    log.Fatalln(err)
  }
  taskresult := TaskResult{}
  for rows.Next() {
    err := rows.StructScan(&taskresult)
    if err != nil {
      log.Fatalln(err)
    }
    fmt.Printf("%#v\n", taskresult)
  } 
  return taskresult
}
