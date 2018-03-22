# eki-concierge
[![Build Status](https://travis-ci.org/kutsuzawa/eki-concierge.svg?branch=master)](https://travis-ci.org/kutsuzawa/eki-concierge) [![Maintainability](https://api.codeclimate.com/v1/badges/cd8080baad28f271fad7/maintainability)](https://codeclimate.com/github/kutsuzawa/eki-concierge/maintainability)  
2駅から両方からちょうどよい距離の駅を教えてくれるAPI
## Endpoint
### `GET /v1/search?station=&another=`
`station`: 1駅目の駅名  
`another`: 2駅名の駅名

## 使用api
内部で利用しているAPIは[Tokyo乗り換えAPI](https://api.trip2.jp)
