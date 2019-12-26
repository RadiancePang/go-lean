#基础镜像
FROM centos:latest

RUN yum install wget -y

RUN wget https://bootstrap.pypa.io/get-pip.py
RUN python get-pip.py
RUN yum install gcc-c++ -y
RUN yum install gcc libffi-devel python-devel openssl-devel -y
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ gevent
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ shap
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ lightgbm


RUN rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO
RUN curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo

RUN yum install golang -y

#工作目录
WORKDIR /root/go/src

#工作目录内容copy到镜像目录
COPY src /root/go/src

COPY summary.py /usr/lib64/python2.7/site-packages/shap/plots

RUN go build inference/inference.go

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

EXPOSE 8078
#项目上下文
ENTRYPOINT ["./inference","--config","beta"]