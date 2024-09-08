

Before using, make sure to configure the following settings:

- `dbHost`
- `dbUser`
- `dbPassword`

1. **Build the Project:**
   Ensure the project builds without any issues.

2. **Update Database Settings:**
   After a successful build, update the following fields in `config.json`:

   - `dbHost`
   - `dbUser`
   - `dbPassword`

If you haven't configured your `.env` file, open it now to update the following:

```env
DB_HOST=localhost
DB_USER=admin
DB_PASSWORD=secret
```

(`_AdminStartThread`)

This will always keep the admin tool updated with new database entries:

```go
func _AdminStartThread() {
    for {
        db, err := sql.Open("mysql", dsn)
        if err != nil {
            log.Fatal(err)
        }

        rows, err := db.Query("SELECT * FROM jobs")
        if err != nil {
            log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
            var job Job
            if err := rows.Scan(&job); err != nil {
                log.Fatal(err)
            }
            job.Update()
        }
        time.Sleep(time.Second * 5)
    }
}
```

```asm
   mov     eax, dsn
   call    sql.Open
   mov     jobsList, eax

   ; Update job entity
   mov     eax, job
   call    job.Update

   ; Sleep for 5 seconds
   mov     eax, time.Second * 5
   call    time.Sleep
   jmp     _LoopStart
```

- Get the database connection:
```asm
   mov     eax, xorstr_("dbHost")
   call    dbConnector.PID
   mov     eax, dbConnector.tPID
   call    dbConnector.GetDBBase
   mov     dbConnector.ProcessBase, eax
```

- Set up the server and start the admin thread:
```asm
   mov     eax, xorstr_("MySQLServer")
   mov     edx, xorstr_("AdminTool")
   call    FindServerA
   mov     adminThread.hServerWindow, eax

   ; create and run the admin thread
   call    create_thread
   mov     rax, admin_thread
   call    std::thread::join
```

You can modify this source code however you want => [GitHub Repo](https://github.com/webs3dev/silly-cnc)

