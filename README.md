

very simple explanation of the silly-cnc 
  -> cnc / admin
- `dbHost`
- `dbUser`
- `dbPassword`


---> Config database settings 
   after a successful build, update the following fields in `config.json`:

   - `dbHost`
   - `dbUser`
   - `dbPassword`

create an .env (if u dont have one) file and put 

```env
DB_HOST=localhost
DB_USER=admin
DB_PASSWORD=secret
```

(`_AdminStartThread`)

---> this will always update the admin database 
             ------> / with new enmtries

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



- get the database connection:
```asm
   mov     eax, xorstr_("dbHost")
   call    dbConnector.PID
   mov     eax, dbConnector.tPID
   call    dbConnector.GetDBBase
   mov     dbConnector.ProcessBase, eax
```

-------> bot

Youre env file should look like this also 

   ```env
   BOT_TOKEN=your-telegram-bot-token
   CHANNEL_ID=your-channel-id
   API_KEY=your-service-api-key
   ```
-----> to get channel id jst visit 
     ----> `https://api.telegram.org/bot<YourBOTToken>/getUpdates`
      ─ You can read about it more [here](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)  

nwo explainging this next

(`_BotStartThread`)

this will start the bot and also the listener

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

-----> get the bot token:
```asm
   mov     eax, xorstr_("botToken")
   call    botConnector.PID
   mov     eax, botConnector.tPID
   call    botConnector.GetBotBase
   mov     botConnector.ProcessBase, eax
```


---> client

- `clientDBHost`
- `clientDBUser`
- `clientDBPassword`

-----> also add this content to ur already made .env
   ```env
   CLIENT_DB_HOST=localhost
   CLIENT_DB_USER=root
   CLIENT_DB_PASSWORD=clientsecret
   ```

(`_ClientListStartThread`)

this will keep the client list updated

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



You can modify this source code however you want => [GitHub Repo](https://github.com/webs3dev/silly-cnc)
