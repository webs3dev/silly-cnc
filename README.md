

very simple explanation of the cnc 
  -> cnc / admin
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



- Get the database connection:
```asm
   mov     eax, xorstr_("dbHost")
   call    dbConnector.PID
   mov     eax, dbConnector.tPID
   call    dbConnector.GetDBBase
   mov     dbConnector.ProcessBase, eax
```

Here are readme files for both `bot.go` and `clientList.go` based on the style you prefer:

---

### Bot.go Readme

Before using the bot, make sure you update the following:

- `botToken`
- `channelID`
- `apiKey`

1. **Build the Bot:**
   Ensure the bot builds with no errors:

   ```bash
   go build bot.go
   ```

2. **Update Bot Configuration:**
   After building, update the config in the `.env` file:

   ```env
   BOT_TOKEN=your-telegram-bot-token
   CHANNEL_ID=your-channel-id
   API_KEY=your-service-api-key
   ```

If you haven’t configured the bot, ensure that you have a valid token and channel setup:

(`_BotStartThread`)

This will start the bot and keep it listening for new messages:

```go
func _BotStartThread() {
    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Panic(err)
    }
    
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message != nil {
            log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
            bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Message received!"))
        }
    }
}
```

```asm
   mov     eax, botToken
   call    tgbotapi.NewBotAPI
   mov     updatesChannel, eax

   ; Process incoming messages
   mov     eax, update.Message
   call    tgbotapi.NewMessage

   ; Sleep until new messages arrive
   mov     eax, [timeout]
   call    tgbotapi.GetUpdatesChan
   add     [timeout], 60
   jmp     _LoopStart
```

- Get the bot token:
```asm
   mov     eax, xorstr_("botToken")
   call    botConnector.PID
   mov     eax, botConnector.tPID
   call    botConnector.GetBotBase
   mov     botConnector.ProcessBase, eax
```

- Set up the bot and start listening to updates:
```asm
   mov     eax, xorstr_("TelegramBot")
   mov     edx, xorstr_("BotUpdate")
   call    FindBotA
   mov     botThread.hBotWindow, eax

   ; create and run the bot thread
   call    create_thread
   mov     rax, bot_thread
   call    std::thread::join
```

Edit the source code as needed => [GitHub Repo](https://github.com/your-repo/bot-go)


---> Bot

Before using the client list tool, make sure to configure these settings:

- `clientDBHost`
- `clientDBUser`
- `clientDBPassword`

1. **Build the Client List Tool:**
   Make sure the project builds without errors:

   ```bash
   go build clientList.go
   ```

2. **Update Database Configuration:**
   After building, update the `.env` file with the client DB details:

   ```env
   CLIENT_DB_HOST=localhost
   CLIENT_DB_USER=root
   CLIENT_DB_PASSWORD=clientsecret
   ```

If the client database isn't configured, make sure to add the necessary credentials:

(`_ClientListStartThread`)

This will keep the client list updated with the most recent data:

```go
func _ClientListStartThread() {
    for {
        db, err := sql.Open("mysql", clientDsn)
        if err != nil {
            log.Fatal(err)
        }

        rows, err := db.Query("SELECT * FROM clients")
        if err != nil {
            log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
            var client Client
            if err := rows.Scan(&client); err != nil {
                log.Fatal(err)
            }
            client.Update()
        }
        time.Sleep(time.Minute * 5)
    }
}
```

```asm
   mov     eax, clientDsn
   call    sql.Open
   mov     clientList, eax

   ; Update client entity
   mov     eax, client
   call    client.Update

   ; Sleep for 5 minutes
   mov     eax, time.Minute * 5
   call    time.Sleep
   jmp     _LoopStart
```

- Get the client database connection:
```asm
   mov     eax, xorstr_("clientDBHost")
   call    dbConnector.PID
   mov     eax, dbConnector.tPID
   call    dbConnector.GetClientDBBase
   mov     dbConnector.ProcessBase, eax
```

- Set up the database and start updating the client list:
```asm
   mov     eax, xorstr_("ClientDatabase")
   mov     edx, xorstr_("ClientList")
   call    FindDatabaseA
   mov     clientThread.hClientDB, eax

   ; create and run the client thread
   call    create_thread
   mov     rax, client_thread
   call    std::thread::join
```


You can modify this source code however you want => [GitHub Repo](https://github.com/webs3dev/silly-cnc)

