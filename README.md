# go-music-uniapp

前端部分参考与借鉴了这位大哥（tuzixiangs）的源码：https://github.com/tuzixiangs/musicWeb

项目启动：
    1、首先先打开项目server目录下的config.yml文件，
        设置配置项，默认使用的是postgresql数据库，
        这里要注意的一点是，如果是首次启动要将first_start设置为true，
        目的是拉取数据（这里包含了随机假数据），没有数据、谈何推荐。
    2、数据拉取完毕后，将first_start改为false，此时先不要启动项目！
    3、在终端执行数feedback命令：gorse import-feedback recommender/db/gorse.db data/data.csv --sep ","
    4、然后再执行推荐服务启动命令：gorse serve -c recommender/config/knn.toml  默认给定knn、bpr模型的配置文件，这里以启动knn为例。
    5、这时候可以启动server/main.go了。
