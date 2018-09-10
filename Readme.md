# PGBatch

Executes a batch of PostgreSQL SQL commands in a single transaction.

## Install

```bash
go get -u github.com/behrang/go-pgbatch
```

## Usage

First create a connection pool using PGSERVICE name of your database:

```go
dbHandler := pgbatch.New("db_service_name", "")
```

Then to execute a batch of commands:

```go
people := make([]Person, 0)
err := dbHandler.Batch([]Command{
    {
        Query: `A SQL query`,
    },
    {
        Query: `Another SQL query with $1 and $2 parameters`,
        Args: []interface{}{param1, param2},
        Affect: 1,
    },
    {
        Query: `Yet another SQL query to scan results with a dynamic param $1`,
        ArgsFunc: func() []interface{} {
            return []interface{}{dynamicParam}
        },
        Scan: func(scan func(...interface{}) error) error {
            p := Person{}
            err := scan(&p.Name, &p.Age)
            people = append(people, person)
            return err
        },
    }
})
```

If any of the queries fail, transaction will be rollbacked.

`Affect` checks the affected rows. If it is a positive number, the number of affected rows should be equal to it or transaction will be rolled back. If is a negative number, no rows should be affected or transaction will be rolled back. If it is 0 or omitted, affected rows will not be checked.

`Args` provides arguments to the SQL prepared statement. These args should be final at the time of batch command creation.

`ArgsFunc` provides dynamic parameters to the SQL prepared statement. If args for a query need to be calculated from other queries in the same batch, use this function.

`Scan` scans each row of the results. If an error is returned, transaction will be rolled back.

`ScanOnce` scans at most one row. More rows will be ignored. If no row exists, nothing will be scanned. If an error is returned, transaction will be rolled back.

If no error occurs, the transaction will be committed. Intermediary rows and result sets are cleaned up.

## License

MIT