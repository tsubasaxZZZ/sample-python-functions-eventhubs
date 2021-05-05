# Sample Event Hubs

## Go

### 実行方法

```bash
$ cd test
$ go build
$ export EVENTHUB_CONNECTION_STRING="Endpoint=sb://evhns-tsunomurpoc-southeastasia-2qe0fs.servicebus.windows.net/;SharedAccessKeyName=test;SharedAccessKey=key;EntityPath=evh-tsunomurpoc-southeastasia-2qe0fs"
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

## Functions

```json
{
  "IsEncrypted": false,
  "Values": {
    "AzureWebJobsStorage": "DefaultEndpointsProtocol=https;AccountName=example;AccountKey=ik8traxWHcgIBWG7Jf+WrjDRV97wuKZ8mmqSGevtXDfS5LL3hZqJ7I9oqDNShWo0u8g3SRngGLPk9BQUMkYEpA==;EndpointSuffix=core.windows.net",
    "FUNCTIONS_WORKER_RUNTIME": "python",
    "evhnstsunomurpocsoutheastasia2qe0fs_RootManageSharedAccessKey_EVENTHUB": "Endpoint=sb://evhns-tsunomurpoc-southeastasia-2qe0fs.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=example;EntityPath=evh-tsunomurpoc-southeastasia-2qe0fs",
    "APPINSIGHTS_INSTRUMENTATIONKEY": "40e24665-ac45-4abc-9d77-xxxxxxxx",
    "APPLICATIONINSIGHTS_CONNECTION_STRING": "InstrumentationKey=40e24665-ac45-4abc-9d77-db2b7e414e6d;IngestionEndpoint=https://southeastasia-0.in.applicationinsights.azure.com/",
    "AzureWebJobsDashboard": "DefaultEndpointsProtocol=https;AccountName=sttsunomurpoc2qe0fs;AccountKey=example;EndpointSuffix=core.windows.net",
    "FUNCTIONS_EXTENSION_VERSION": "~3",
    "SQL_SERVER": "sql-tsunomurpoc-southeastasia-2qe0fs.database.windows.net",
    "SQL_DATABASE": "sqldb-tsunomurpoc-southeastasia-2qe0fs",
    "SQL_USERNAME": "admin",
    "SQL_PASSWORD": "password!",
    "AzureWebJobs.EventHubTrigger1.Disabled": "false"
  }
}

```