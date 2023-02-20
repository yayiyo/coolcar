# coolcar

## 如何编译以及运行小程序
1. `cd wx/miniprogram`
2. `npm install`
3. 打开小程序开发工具，导入项目
4. 点击工具 --> 构建npm
5. 点击编译


#### protocbuf 生成

server: `./gen.sh ./${service_name}/api ./${service_name}/api/gen/v1 ${service_name}`

wx: `./gen.sh ../../../server/${service_name}/api ${service_name}`