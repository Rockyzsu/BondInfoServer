# 基于Gin的http接口服务，提供可转债基础和扩展数据给Ptrade


[公众号](http://xximg.30daydo.com/webupload/ZSXQ_20220616_181142268.png)
----
## API
```
POST /bondinfo
{
    "sign":"id"
}

```
返回值:
```
`可转债代码` VARCHAR(10) NULL,
  `可转债名称` TEXT NULL,
  `可转债价格` DOUBLE NULL,
  `正股名称` TEXT NULL,
  `正股代码` TEXT NULL,
  `正股现价` DOUBLE NULL,
  `正股涨跌幅` DOUBLE NULL,
  `最新转股价` DOUBLE NULL,
  `溢价率` DOUBLE NULL,
  `可转债涨幅` DOUBLE NULL,
  `转股价值` DOUBLE NULL,
  `双低` DOUBLE NULL,
  `回售触发价` DOUBLE NULL,
  `转股起始日` BIGINT NULL,
  `到期时间` TEXT NULL,
  `成交额(万元)` DOUBLE NULL,
  `强赎价格` DOUBLE NULL,
  `剩余时间` DOUBLE NULL,
  `评级` TEXT NULL,
  `下修条件` TEXT NULL,
  `换手率` DOUBLE NULL,
  `下修提示` TEXT NULL,
  `提出下调次数` BIGINT NULL,
  `正股成交量` DOUBLE NULL,
  `转债剩余占总市值比` DOUBLE NULL,
  `剩余规模` DOUBLE NULL,
  `发行规模` DOUBLE NULL,
  `期权价值` TEXT NULL,
  `强赎提示` TEXT NULL,
  `强赎日期` TEXT NULL,
  `上市日期` TEXT NULL,
  `到期收益率` DOUBLE NULL,
  `强赎标志` TEXT NULL,
  `是否两融标的` TEXT NULL,
  `下修成功次数` TEXT NULL,
  `转股日期提示` TEXT NULL,
  `参考YTM` TEXT NULL,
  `更新日期` TEXT NULL
```
