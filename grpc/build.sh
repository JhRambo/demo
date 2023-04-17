#!/bin/bash
dir=/code/demo/grpc/proto/          #proto文件所在目录
scDir=/code/demo/grpc/proto/        #生成的文件所在目录

cd $dir

for filePath in $(ls $dir) 
do 
    if [ -f $dir"/"$filePath ]; then
        if [ ${filePath##*.} = "proto" ]; then
            #echo "create and cp ${filePath%.*} proto"
            if [ ! -d $scDir${filePath%.*} ];then
                mkdir $scDir${filePath%.*}
            fi
            protoc -I . --go_out=$scDir${filePath%.*}/ --go_opt paths=source_relative --go-grpc_out=$scDir${filePath%.*}/ --go-grpc_opt paths=source_relative --grpc-gateway_out=$scDir${filePath%.*}/ --grpc-gateway_opt logtostderr=true  --grpc-gateway_opt paths=source_relative $filePath
        fi
	fi
done

echo "all success ！"