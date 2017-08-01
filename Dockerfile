FROM golang:onbuild

ENV PATH "$PATH:/usr/local/ncbi/blast/bin"

ENTRYPOINT app run -a "192.168.64.2" -p "32548"
