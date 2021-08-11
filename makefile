CC        = gcc 
CFLAGS    = -Wall -O -g
LIB       = libhello.a
OBJS	  = hello.o

all: ${LIB}

# .o的文件名对应.c的文件名
%.o: %.c
	${CC} ${CFLAGS} -o $@ -c $<

# 依赖的OBJS，由以上规则生成
${LIB}: ${OBJS}
	rm -f $@
	ar -cru $@ ${OBJS}
	rm -f ${OBJS}

# 无视clean文件是否存在
.PHONY: clean
clean:
	rm hello.o ${LIB}