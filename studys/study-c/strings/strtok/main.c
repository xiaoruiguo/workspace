/*
 * main.c
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[])
{
	char str[100];
	sprintf(str,"a b c d e f");
	char *delim = " ";
	char *token = NULL;

	token = strtok(str, delim);
	while (NULL != token) {
		printf("%s\n", token);
		token = strtok(NULL, delim);
	}

	return 0;
}
