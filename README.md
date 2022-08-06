# let the grass grow

## Endpoint

- `/v1/training/`: 全件取得

`{"trainings":[{"date":"2022-08-01","count":3},{"date":"2022-08-02","count":4},{"date":"2022-08-03","count":1}]}`

- `/v1/training/?start=2022-08-02&end=2022-08-03`: 期間指定取得

`{"trainings":[{"date":"2022-08-02","count":4},{"date":"2022-08-03","count":1}]}`

- `/v1/training/2022-08-02`: 日付指定取得

`{"training":{"date":"2022-08-02","count":4}}`
