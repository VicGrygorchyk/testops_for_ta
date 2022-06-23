FROM golang:latest

WORKDIR /home/tests
COPY . /home/tests/
RUN git clone https://github.com/VicGrygorchyk/testops_for_ta.git && cd testops_for_ta \
&& go mod tidy

CMD [ "go", "test", "-v", "./" ]
