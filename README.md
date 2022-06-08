# go-api
organize golang  apple pay ,google pay  and other api 


## receipt-data凭证请求后返回的内容字段翻译 ~

"receipt_type":"Production", //凭证类型 沙箱或正式

"adam_id":1111111111, //APPID

"app_item_id":1111111111,  //APPID

"bundle_id":"com.xxxx.www", //包名

"application_version”:”1.0.2”,//应用版本号

"download_id":333333333333333, //应用下载交易的唯一标识符。

"version_external_identifier":834324799,//版本外部的标识，沙箱环境下其值为：0正式环境其值为一个数字，会变，原因未知

"receipt_creation_date":"2020-05-05 08:44:20 Etc/GMT", //receipt创建日期(ms)

"receipt_creation_date_ms":"1588668260000”,//receipt创建日期的时间戳 转行可得到北京时间 （毫秒）

"receipt_creation_date_pst":"2020-05-05 01:44:20 America/Los_Angeles",//receipt创建日期的时间戳 美国洛杉矶时

"request_date":"2020-05-05 08:44:21 Etc/GMT", //该凭证请求验证的时间  格林威治时间（与北京时间相差8小时）

"request_date_ms":"1588668261849",//该凭证请求验证的时间戳 转行可得到北京时间）（毫秒）

"request_date_pst":"2020-05-05 01:44:21 America/Los_Angeles",//该凭证请求验证的时间  美国洛杉矶时间

"original_purchase_date":"2020-03-25 14:06:44 Etc/GMT", //原始购买的格林威治时间（与北京时间相差8小时）

"original_purchase_date_ms":"1585145204000", //（原始购买时间进行时间戳转行可得到北京时间）（毫秒）

"original_purchase_date_pst":"2020-03-25 07:06:44 America/Los_Angeles", //原始购买时间,美国洛杉矶时间

"original_application_version":"1.0.2", //原始应用版本号

"preorder_date":"2020-03-05 17:07:05 Etc/GMT", //预计购买的格林威治时间（与北京时间相差8小时

"preorder_date_ms":"1583428025000", //预计购买时间时间戳 转行可得到北京时间）（毫秒）

"preorder_date_pst":"2020-03-05 09:07:05 America/Los_Angeles", //预计购购时间 美国洛杉矶时间

"in_app":[

{

"quantity":"1",  //购买商品的数量

"product_id":"com.xxx.01",//商品ID

"transaction_id":"100000000000000",  //交易的标识

"original_transaction_id":"100000000000000",  //原始交易ID

"purchase_date":"2020-05-05 08:44:19 Etc/GMT", //购买的格林威治时间（与北京时间相差8小时）

"purchase_date_ms":"1588668259000", //购买时间毫秒（进行时间戳转行可得到北京时间）（毫秒）

"purchase_date_pst":"2020-05-05 01:44:19 America/Los_Angeles",  //购买时间,美国洛杉矶时间

"original_purchase_date":"2020-05-05 08:44:19 Etc/GMT",//原始购买的格林威治时间（与北京时间相差8小时）

"original_purchase_date_ms":"1588668259000",//（原始购买时间进行时间戳转行可得到北京时间）（毫秒）

"original_purchase_date_pst":"2020-05-05 01:44:19 America/Los_Angeles",//原始购买的格林威治时间（与北京时间相差8小时）

"is_trial_period":"false" //是否是免费试用

}

]

},

"status”:0,//验证状态0表示成功，其他表示不成功

"environment":"Production" //支付环境 沙箱或正式环境
