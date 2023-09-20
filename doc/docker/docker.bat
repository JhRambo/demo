
@echo off

docker stop my-centos

docker run -it --rm --name my-centos --privileged=true -v D:\code:/code  my-centos:v0.1 /bin/bash

