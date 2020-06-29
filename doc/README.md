# Swagger 文档
## 安装redoc
```
npm(cnpm) install -g redoc-cli
# 有时redoc-cli无法直接使用， 需要安装npx.
npm install -g npx
```

## 调试使用
```
redoc-cli serve -w --options.expandResponses="all" --options.requiredPropsFirst=true index.yaml
```

## 生产使用
```
redoc-cli bundle -o api.html index.yaml 
或
npx redoc-cli bundle -o api.html index.yaml
```