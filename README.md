# eki-concierge
2駅から両方からちょうどよい距離の駅を教えてくれるAPI
## Endpoint
### `GET /v1/search?station=&another=`
`station`: 1駅目の駅名  
`another`: 2駅名の駅名

## 使用api
内部で利用しているAPIは[Tokyo乗り換えAPI](https://api.trip2.jp)
