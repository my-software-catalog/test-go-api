# Your API

Questa è una semplice API costruita con Go e il framework Gin.

## Installazione

Assicurati di avere Go 1.21+ installato.

1. Clona il repository:
    ```bash
    git clone https://github.com/tuo-username/your-api.git
    cd your-api
    ```

2. Installa le dipendenze:
    ```bash
    go mod tidy
    ```

## Esecuzione

Per avviare il server, esegui il seguente comando:

```bash
go run main.go
```

Il server sarà in ascolto sulla porta `8080`.

## Utilizzo

Puoi testare l'endpoint API con curl o un client HTTP. Ecco un esempio usando curl:

```bash
curl http://localhost:8080/api/hello
```

Dovresti ricevere la seguente risposta JSON:

```json
{
    "message": "Hello, World!"
}
```