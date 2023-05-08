# nginx

nginx 可以用来部署静态网页：

- 其中 nginx 的总配置文件位于 `/etc/nginx/nginx.conf` 文件，引入了包括可用网站等其他配置文件
- 可用网站的配置文件在 `/etc/nginx/sites-available/default` 进行了定义
  - `root` 定义了访问根路径映射的文件夹
  - `index` 定义了没有文件对应的 url 所访问的页面

重新配置 ngnix 的指令：`nginx -s reload`
