FROM golang:onbuild

ENV PATH "$PATH:/usr/local/ncbi/blast/bin"


RUN tar -xzvf ncbi-blast-2.6.0+-src.tar.gz
RUN cd ncbi-blast-2.6.0+-src/c++ && ./configure && make && sudo make install
RUN rm -rf ncbi-blast-2.6.0+-src.tar.gz*
RUN sudo ldconfig


ENTRYPOINT app run -a "192.168.64.2" -p "32548"
