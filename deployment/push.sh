DOMAIN=$1
VERSION=$2

docker tag coolcar/$DOMAIN  registry.cn-beijing.aliyuncs.com/coolcar-code/$DOMAIN:$VERSION

docker push registry.cn-beijing.aliyuncs.com/coolcar-code/$DOMAIN:$VERSION