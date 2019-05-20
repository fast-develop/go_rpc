cur_dir=`pwd`
echo "cur_dicr:" $cur_dir
project_dir=`pwd`
root_dir=`cd ../../;pwd`
project_name=recall_cf
echo "root_dir" $root_dir

export GOPATH=$root_dir

protoc --go_out=plugins=grpc:$project_dir/proto/recall_cf --proto_path=$project_dir/proto/recall_cf $project_dir/proto/recall_cf/recall_cf.proto

find $project_dir/proto -name ".libs" | xargs rm -rf

mkdir -p $project_dir/_publish_dir/recall_cf/log
cp -r $project_dir/release/* $project_dir/_publish_dir/recall_cf
go build -gcflags '-N -l' -o $project_dir/_publish_dir/recall_cf/bin/$project_name $project_dir/main/main.go
