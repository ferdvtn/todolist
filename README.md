# todolist

## create new migration
```migrate create -ext sql -dir db/migrations -seq create_todos_table```

## run migration
```migrate -source "file://db/migrations" -database "mysql://root:root@tcp(127.0.0.1:3306)/todolist" up```

## undo migration
```migrate -source "file://db/migrations" -database "mysql://root:root@tcp(127.0.0.1:3306)/todolist" down 1```