# CHANGELOG 

## 0.1.0

* Server creator
* Route creator
* Add method to route creator
* Run server feature 

## 0.2.0

* Server does not accept two methods at the same route. `panic.Err()` is generated
* `NewJSend` method renamed
* `JSendMessage` methods to add data and message

## 0.3.0

* Headers has been added into the **request** and the **response**
* Creating responses now is `NewResponse` instead of `NewJSend`

### 0.3.1

* Fix the error report on the channel received from `RunServer()`

### 0.4.0

* Only `NewResponse()` method is required on the response when is added a new method