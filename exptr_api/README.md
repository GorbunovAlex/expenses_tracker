## TODO
- [ ] Add Makefile commands for migrations
    ```
    migration_up: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
    
    migration_down: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verb    ose down
    
    migration_fix: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force     VERSION
    ```
## Migrations

### Run migration
```
$ migrate -path internal/storage/postgres/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up

```

### Rollback migration
```
$ migrate -path internal/storage/postgres/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down
```

### Resolve migration error
```
$ migrate -path internal/storage/postgres/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force <VERSION>
```