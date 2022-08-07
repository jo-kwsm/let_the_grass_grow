# let the grass grow

## Endpoint

- `/v1/training/?user=jo-kwsm`: 全件取得

`{"trainings":[{"user":"jo-kwsm","date":"2020-01-01","walking":0,"running":0,"cycling":0,"swimming":0},{"user":"jo-kwsm","date":"2020-01-02","walking":0,"running":0,"cycling":0,"swimming":4915},{"user":"jo-kwsm","date":"2020-01-03","walking":0,"running":374,"cycling":7500,"swimming":0},{"user":"jo-kwsm","date":"2020-01-04","walking":0,"running":0,"cycling":0,"swimming":7510}]}`

- `/v1/training/?user=jo-kwsm&start=2022-08-02&end=2022-08-03`: 期間指定取得

`{"trainings":[{"user":"jo-kwsm","date":"2022-08-02","walking":0,"running":0,"cycling":0,"swimming":0},{"user":"jo-kwsm","date":"2022-08-03","walking":7023,"running":0,"cycling":0,"swimming":0}]}`

- `/v1/training/2022-08-02?user=jo-kwsm`: 日付指定取得

`{"training":{"user":"jo-kwsm","date":"2022-08-02","walking":0,"running":0,"cycling":0,"swimming":0}}`
