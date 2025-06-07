# PropertyDetectionGo
初学了一个星期golang后，做了个简单的知产检测系统，希望大家对新人多多支持，谢谢。
## avue
这个是前端工程，用的vue2，拉项目后用npm i下载依赖就行了，node我用的是v14.16版本。
## PropertyDetection
这个是后端工程，用golang写的，数据库用mysql8.0及以上，minio，然后再config.yaml配置就行了，记得本地调试记得把env改成dev，部署改成其他的就行了。
后端包含了基本的角色、用户，菜单管理和菜单权限分配。
## 使用方法
其实做的非常简单，在知产数据库添加数据，后续在知产检索会从知产数据库进行对比。
