DESTDIR = .
PROGRAM = cApp
SRCS    = $(shell find $(PWD) -name "*.c" -type f)
OBJS    = $(SRCS:%.c=%.o)
CC      = gcc
INCDIRS = 
CFLAGS  = -g -O0 -Wall -Wextra -std=c17
LDFLAGS = -L../lib
LDLIBS  = -lmylib

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@ $(INCDIRS)

$(PROGRAM): $(OBJS)
	$(CC) $(CFLAGS) $(LDFLAGS) -o $(DESTDIR)/$(PROGRAM) $(OBJS) $(LDLIBS)

.PHONY: build
build: $(PROGRAM)

.PHONY: clean
clean:
	$(RM) $(OBJS) $(DESTDIR)/$(PROGRAM)
