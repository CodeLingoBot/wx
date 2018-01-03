# 微信公众号自动回复

## 申请公众号步骤

1. [注册公众号](https://mp.weixin.qq.com/cgi-bin/registermidpage?action=index&lang=zh_CN)
2. 获取API的 appid + token （管理员界面左下角“开发”部分）
3. 设置服务器回调地址 “基本配置” 
    https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421135319
    
4. 配置nginx做proxy，直接用程序侦听80端口也不是不行，只是姿势好像不太正确

    ```
    upstream wx {
        server 127.0.0.1:7400;
    }

    server {
    listen 80 default_server;
    listen [::]:80 default_server;
    location /wx {
        proxy_pass http://wx;
        proxy_set_header Host $host:$server_port;
    }

    ```

4. 编译代码 侦听7400端口
