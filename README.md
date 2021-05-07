# Sample Event Hubs

## 事前準備

### クライアント

```bash
$ cd test
$ go build
$ export EVENTHUB_CONNECTION_STRING="Endpoint=sb://evhns-tsunomurpoc-southeastasia-2qe0fs.servicebus.windows.net/;SharedAccessKeyName=test;SharedAccessKey=key;EntityPath=evh-tsunomurpoc-southeastasia-2qe0fs"
$ export  EVENTHUB_NAMESPACE=evhns-tsunomurpoc-southeastasia-wcqps5
$ export  EVENTHUB_NAME=evh-tsunomurpoc-southeastasia-wcqps5
$ ./eventhub-sample-client  -h
Usage of ./eventhub-sample-client:
  -count int
        合計のリクエスト数 (default 1)
  -prefix string
        テストをユニークにするためのプレフィックス
  -term int
        リクエスト送信完了までの期待する時間(秒)
        (クライアントの性能によって達成できない可能性があります) (default 1)
  -verbose int
        デバッグモード(数字が大きい方が詳細) (default 1)
```

### Functions

#### local.settings.json

必要に応じて編集。初期デプロイ後に Azure 上の Functions からダウンロードする。

ローカルでデバッグするときは、`Address already in use` が発生するため、`FUNCTIONS_WORKER_PROCESS_COUNT` を `1` にする。

```json
{
  "IsEncrypted": false,
  "Values": {
    "EVENTHUB": "Endpoint=sb://evhns-tsunomurpoc-southeastasia-c4mbe6.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=xxxxxxxxx=",
    "FUNCTIONS_WORKER_PROCESS_COUNT": "10",
    "PYTHON_THREADPOOL_THREAD_COUNT": "1",
    "SQL_SERVER": "sql-tsunomurpoc-southeastasia-c4mbe6.database.windows.net",
    "SQL_DATABASE": "sqldb-tsunomurpoc-southeastasia-c4mbe6",
    "SQL_USERNAME": "sqladmin",
    "SQL_PASSWORD": "password",
    "AzureWebJobs.EventHubTrigger1.Disabled": "false",
    "FUNCTIONS_WORKER_RUNTIME": "python",
    "APPINSIGHTS_INSTRUMENTATIONKEY": "xxxxxxxx-50b5-4a77-a087-9efdc13ded9c",
    "APPLICATIONINSIGHTS_CONNECTION_STRING": "InstrumentationKey=xxxxxxxx-50b5-4a77-a087-9efdc13ded9c;IngestionEndpoint=https://southeastasia-0.in.applicationinsights.azure.com/",
    "AzureWebJobsDashboard": "DefaultEndpointsProtocol=https;AccountName=sttsunomurpocc4mbe6;AccountKey=W/xxxxxxxxxxxxxxxxx==;EndpointSuffix=core.windows.net",
    "AzureWebJobsStorage": "DefaultEndpointsProtocol=https;AccountName=sttsunomurpocc4mbe6;AccountKey=W/xxxxxxxxx;EndpointSuffix=core.windows.net",
    "FUNCTIONS_EXTENSION_VERSION": "~3"
  }
}
```

#### function.json

`eventHubName` を実際の環境に合わせる。

```json
{
  "scriptFile": "__init__.py",
  "bindings": [
    {
      "type": "eventHubTrigger",
      "name": "events",
      "direction": "in",
      "eventHubName": "evh-tsunomurpoc-southeastasia-wcqps5",
      "connection": "EVENTHUB",
      "consumerGroup": "$Default",
      "cardinality": "many"
    }
  ]
}
```

### SQL Database

table.sql の CREATE 文を実行